-- name: CreateTask :one
INSERT INTO tasks (
    user_uuid,
    name
) VALUES (
    @user_uuid,
    @name
) RETURNING uuid;

-- name: GetTaskByUUID :one
SELECT *
FROM tasks
WHERE uuid = @task_uuid;

-- name: GetTasksByUserUUID :many
SELECT *
FROM tasks
WHERE user_uuid = @user_uuid;

-- name: UpdateTaskByUUID :one
UPDATE tasks
SET name = coalesce(sqlc.narg('name'), name),
    duration = coalesce(sqlc.narg('duration'), duration),
    is_completed = coalesce(sqlc.narg('is_completed'), is_completed)
WHERE uuid = @task_uuid
RETURNING *;
