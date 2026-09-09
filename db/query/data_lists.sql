-- name: ListDataLists :many
SELECT id, label, kind, key, value, item_order, create_time, update_time
FROM data_lists
WHERE (sqlc.arg(kind)::text = '' OR kind = sqlc.arg(kind)::text)
ORDER BY id;

-- name: CountDataListsByKind :one
SELECT count(*) FROM data_lists
WHERE (sqlc.arg(kind)::text = '' OR kind = sqlc.arg(kind)::text);

-- name: ListDataListsByKindPage :many
SELECT id, label, kind, key, value, item_order, create_time, update_time
FROM data_lists
WHERE (sqlc.arg(kind)::text = '' OR kind = sqlc.arg(kind)::text)
ORDER BY id
LIMIT sqlc.arg(page_size)::integer OFFSET sqlc.arg(page_offset)::integer;

-- name: ListDataListsByKindsAsc :many
SELECT id, label, kind, key, value, item_order, create_time, update_time
FROM data_lists
WHERE kind = ANY(sqlc.arg(kinds)::text[])
ORDER BY item_order ASC;

-- name: ListDataListsByKindsDesc :many
SELECT id, label, kind, key, value, item_order, create_time, update_time
FROM data_lists
WHERE kind = ANY(sqlc.arg(kinds)::text[])
ORDER BY item_order DESC;

-- name: ListDataListsByKinds :many
SELECT id, label, kind, key, value, item_order, create_time, update_time
FROM data_lists
WHERE kind = ANY(sqlc.arg(kinds)::text[])
ORDER BY id;

-- name: ListDataListSortData :many
SELECT id, label, kind, key, value, item_order, create_time, update_time
FROM data_lists WHERE kind = $1 ORDER BY item_order DESC;

-- name: GetDataList :one
SELECT id, label, kind, key, value, item_order, create_time, update_time
FROM data_lists WHERE id = $1;

-- name: CreateDataList :one
INSERT INTO data_lists (label, kind, key, value, item_order)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, label, kind, key, value, item_order, create_time, update_time;

-- name: UpdateDataList :one
UPDATE data_lists
SET key = sqlc.arg(key), value = sqlc.arg(value),
    item_order = COALESCE(sqlc.narg(item_order)::integer, item_order),
    update_time = now()
WHERE id = sqlc.arg(id)
RETURNING id, label, kind, key, value, item_order, create_time, update_time;

-- name: UpdateDataListOrder :execrows
UPDATE data_lists SET item_order = $2, update_time = now() WHERE id = $1;

-- name: DeleteDataList :execrows
DELETE FROM data_lists WHERE id = $1;
