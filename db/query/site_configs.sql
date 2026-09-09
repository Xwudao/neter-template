-- name: ListSiteConfigNames :many
SELECT name FROM site_configs ORDER BY name;

-- name: ListSiteConfigs :many
SELECT id, name, config, create_time, update_time FROM site_configs ORDER BY id;

-- name: GetSiteConfig :one
SELECT id, name, config, create_time, update_time FROM site_configs WHERE id = $1;

-- name: GetSiteConfigByName :one
SELECT id, name, config, create_time, update_time FROM site_configs WHERE name = $1;

-- name: CreateSiteConfig :one
INSERT INTO site_configs (name, config)
VALUES ($1, $2)
RETURNING id, name, config, create_time, update_time;

-- name: UpdateSiteConfig :execrows
UPDATE site_configs
SET config = $2, update_time = now()
WHERE name = $1;

-- name: DeleteSiteConfig :execrows
DELETE FROM site_configs WHERE id = $1;
