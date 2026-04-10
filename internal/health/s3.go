package health

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"

	appconfig "Derzhavnaya/internal/config"
)

func CheckS3(cfg appconfig.S3Config) error {
	var logMode aws.ClientLogMode
	if cfg.DebugS3 {
		logMode = aws.LogSigning | aws.LogRetries | aws.LogRequestWithBody | aws.LogResponseWithBody
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		),
		config.WithClientLogMode(logMode),
	)
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.Endpoint)
		o.UsePathStyle = true // обязательно для Garage и MinIO
	})

	out, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:  aws.String(cfg.Bucket),
		MaxKeys: aws.Int32(10),
	})
	if err != nil {
		return err
	}

	log.Info().
		Str("endpoint", cfg.Endpoint).
		Str("bucket", cfg.Bucket).
		Int("objects_shown", len(out.Contents)).
		Bool("truncated", out.IsTruncated != nil && *out.IsTruncated).
		Msg("S3 OK")

	for _, obj := range out.Contents {
		log.Debug().
			Str("key", aws.ToString(obj.Key)).
			Int64("size", aws.ToInt64(obj.Size)).
			Msg("  object")
	}

	return nil
}
