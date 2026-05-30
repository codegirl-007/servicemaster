-- name: CreateOAuthState :one
INSERT INTO qbo_oauth_states (
    id,
    tenant_id,
    state_checksum,
    encrypted_state,
    expires_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, tenant_id, state_checksum, encrypted_state, expires_at, consumed_at, created_at;

-- name: GetActiveOAuthStateByChecksum :one
SELECT *
FROM qbo_oauth_states
WHERE state_checksum = $1
  AND consumed_at IS NULL
  AND expires_at > now();

-- name: ConsumeOAuthState :one
UPDATE qbo_oauth_states
SET consumed_at = now()
WHERE id = $1
  AND consumed_at IS NULL
RETURNING id;

-- name: DeleteExpiredOAuthStates :exec
DELETE FROM qbo_oauth_states
WHERE expires_at < now()
  AND consumed_at IS NULL;
