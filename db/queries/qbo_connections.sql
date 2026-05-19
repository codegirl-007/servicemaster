-- name: CreateQBOConnection :one
INSERT INTO qbo_connections (
    id,
    tenant_id,
    realm_id,
    company_name,
    state
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING id, tenant_id, realm_id, company_name, state, created_at, updated_at;

-- name: GetQBOConnection :one
SELECT id, tenant_id, realm_id, company_name, state, created_at, updated_at
FROM qbo_connections
WHERE id = $1;

-- name: GetQBOConnectionByTenant :one
SELECT id, tenant_id, realm_id, company_name, state, created_at, updated_at
FROM qbo_connections
WHERE tenant_id = $1;

-- name: UpdateQBOConnectionState :one
UPDATE qbo_connections
SET
    state = $2,
    updated_at = now()
WHERE id = $1
RETURNING id, tenant_id, realm_id, company_name, state, created_at, updated_at;

-- name: UpdateQBOConnectionCompanyName :one
UPDATE qbo_connections
SET
    company_name = $2,
    updated_at = now()
WHERE id = $1
RETURNING id, tenant_id, realm_id, company_name, state, created_at, updated_at;

-- name: CreateQBOConnectionEvent :one
INSERT INTO qbo_connection_events (
    id,
    qbo_connection_id,
    tenant_id,
    event_type,
    message,
    metadata
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id, qbo_connection_id, tenant_id, event_type, message, metadata, occurred_at;

-- name: ListQBOConnectionEvents :many
SELECT id, qbo_connection_id, tenant_id, event_type, message, metadata, occurred_at
FROM qbo_connection_events
WHERE qbo_connection_id = $1
ORDER BY occurred_at DESC;

-- name: ListQBOConnectionEventsByTenant :many
SELECT id, qbo_connection_id, tenant_id, event_type, message, metadata, occurred_at
FROM qbo_connection_events
WHERE tenant_id = $1
ORDER BY occurred_at DESC;

-- name: GetLatestQBOConnectionEventByType :one
SELECT id, qbo_connection_id, tenant_id, event_type, message, metadata, occurred_at
FROM qbo_connection_events
WHERE qbo_connection_id = $1
  AND event_type = $2
ORDER BY occurred_at DESC
LIMIT 1;
