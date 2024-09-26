-- name: CreateUser :exec
INSERT INTO users (
    uuid, first_name, last_name
) VALUES (
    @uuid, @first_name, @last_name
);

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByUUID :one
SELECT *
FROM users
WHERE uuid = @user_uuid;
