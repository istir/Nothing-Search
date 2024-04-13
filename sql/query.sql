-- name: CreateFile :one
INSERT INTO files (
  name, path, date_modified, date_created
) VALUES (
  ?,
  ?,
  ?,
  ?
) ON CONFLICT(path) DO UPDATE SET
  name = excluded.name,
  date_modified = excluded.date_modified,
  date_created = excluded.date_created
RETURNING *;

-- name: GetFile :one
SELECT * FROM files WHERE path = ? LIMIT 1;
