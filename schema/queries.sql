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

-- name: GetActiveMenuItems :many
SELECT * FROM web.menu_item
WHERE is_active = true
ORDER BY position ASC;

-- name: DeleteSession :exec
DELETE FROM  web.sessions WHERE id = $1;


-- name: GetAnsweredQuestionsPaginated :many
SELECT *
FROM public.hram_talk
WHERE flag = 1
  AND (data_a , id) < ($2, $3)
ORDER BY data_a DESC, id DESC
LIMIT $1;