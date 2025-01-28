-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens (
    refreshToken, user_id, expires_at, revoked_at, created_at, updated_at
) VALUES (
    $1, $2, NOW() + INTERVAL '60 days', NULL, NOW(), NOW()
)
RETURNING *;

-- name: GetUserFromRefreshToken :one
SELECT users.* FROM users
JOIN refresh_tokens ON users.id = refresh_tokens.user_id
WHERE refreshToken = $1
    AND expires_at > NOW()
    AND revoked_at IS NULL;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET
    updated_at = NOW(),
    revoked_at = NOW()
WHERE refreshToken = $1;

-- name: RevokeTokenByID :exec
UPDATE refresh_tokens
SET
    updated_at = NOW(),
    revoked_at = NOW()
WHERE user_id = $1;