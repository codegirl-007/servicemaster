-- name: CreateTenant :one
INSERT INTO tenants (
    id,
    name
) VALUES (
    $1,
    $2
)
RETURNING id, name, created_at, updated_at;

-- name: GetTenant :one
SELECT id, name, created_at, updated_at
FROM tenants
WHERE id = $1;

-- name: ListTenants :many
SELECT id, name, created_at, updated_at
FROM tenants
ORDER BY created_at DESC;

-- name: UpdateTenantName :one
UPDATE tenants
SET
    name = $2,
    updated_at = now()
WHERE id = $1
RETURNING id, name, created_at, updated_at;
