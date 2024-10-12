-- name: CreateUser :one
INSERT INTO users (
    first_name, last_name
) VALUES (
    @first_name, @last_name
) RETURNING uuid;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByUUID :one
SELECT *
FROM users
WHERE uuid = @user_uuid;
