package auth

import (
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/models"
	"context"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

type limiterEntry struct {
	count     int
	updatedAt time.Time
}

type RateLimiter struct {
	attempts map[string]limiterEntry
	mu       sync.Mutex
}

func NewRateLimiter(lc fx.Lifecycle) *RateLimiter {
	rl := &RateLimiter{
		attempts: make(map[string]limiterEntry),
	}

	_, cancel := context.WithCancel(context.Background())

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info().Msg("Starting RateLimiter cleanup loop...")
			go rl.cleanupLoop(ctx)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Stopping RateLimiter cleanup loop...")
			cancel()
			return nil
		},
	})

	return rl
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	entry, exists := rl.attempts[ip]

	if exists && now.Sub(entry.updatedAt) > 15*time.Minute {
		entry.count = 0
	}

	if entry.count >= 5 {
		return false
	}

	entry.count++
	entry.updatedAt = now
	rl.attempts[ip] = entry

	return true
}

func (rl *RateLimiter) cleanupLoop(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			rl.mu.Lock()
			now := time.Now()
			for ip, entry := range rl.attempts {
				if now.Sub(entry.updatedAt) > 15*time.Minute {
					delete(rl.attempts, ip)
				}
			}
			rl.mu.Unlock()
		}
	}
}

func GetIP(r *http.Request) string {
	if ip := r.Header.Get("X-Real-Ip"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

func FilterMenuByRole(items []db.WebMenuItem, userRole string) []models.MenuItem {
	var filtered []models.MenuItem
	for _, item := range items {
		if !item.Role.Valid || item.Role.String == "" {
			filtered = append(filtered, models.NewMenuItemFromDB(item))
			continue
		}
		if strings.EqualFold(item.Role.String, userRole) || strings.EqualFold(userRole, "admin") {
			filtered = append(filtered, models.NewMenuItemFromDB(item))
		}
	}
	return filtered
}
