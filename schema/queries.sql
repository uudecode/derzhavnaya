-- name: GetUserByEmail :one
SELECT * FROM web.users
WHERE email = $1 LIMIT 1;

-- name: CreateSession :one
INSERT INTO web.sessions (id, user_id, expires_at)
VALUES ($1, $2, $3)
    RETURNING *;

-- name: GetUserBySessionID :one
SELECT u.*
  FROM web.users u
       JOIN web.sessions s
         ON s.user_id = u.id
WHERE s.id = $1 AND s.expires_at > NOW();

-- name: CreateAdminIfNotExist :exec
INSERT INTO web.users (email, password, role, full_name)
VALUES ($1, $2, 'admin', 'System Administrator')
    ON CONFLICT (email) DO NOTHING;