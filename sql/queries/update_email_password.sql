-- name: UpdateEmailPassword :one

UPDATE users u
SET email = $2, password=$3, updated_at=$4
FROM refresh_tokens t
WHERE u.id = t.user_id
AND u.id = $1
RETURNING u.*, t.token;
