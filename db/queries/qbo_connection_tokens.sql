-- name: UpsertQBOConnectionTokens :one
INSERT INTO qbo_connection_tokens (
    qbo_connection_id,
    tenant_id,
    encrypted_access_token,
    encrypted_refresh_token,
    access_token_expires_at,
    refresh_token_expires_at,
    token_type,
    scope
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
ON CONFLICT (qbo_connection_id) DO UPDATE
SET
    encrypted_access_token = EXCLUDED.encrypted_access_token,
    encrypted_refresh_token = EXCLUDED.encrypted_refresh_token,
    access_token_expires_at = EXCLUDED.access_token_expires_at,
    refresh_token_expires_at = EXCLUDED.refresh_token_expires_at,
    token_type = EXCLUDED.token_type,
    scope = EXCLUDED.scope,
    version = qbo_connection_tokens.version + 1,
    updated_at = now()
RETURNING *;

-- name: GetQBOConnectionTokens :one
SELECT *
FROM qbo_connection_tokens
WHERE qbo_connection_id = $1;

-- name: GetQBOConnectionTokensForUpdate :one
SELECT *
FROM qbo_connection_tokens
WHERE qbo_connection_id = $1
FOR UPDATE;

-- name: ReplaceQBOConnectionTokensIfVersion :one
UPDATE qbo_connection_tokens
SET
    encrypted_access_token = $3,
    encrypted_refresh_token = $4,
    access_token_expires_at = $5,
    refresh_token_expires_at = $6,
    token_type = $7,
    scope = $8,
    version = version + 1,
    updated_at = now()
WHERE qbo_connection_id = $1
  AND version = $2
RETURNING *;
