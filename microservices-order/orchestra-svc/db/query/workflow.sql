-- name: CreateWorkflowType :one
INSERT INTO workflow_types (name) 
VALUES 
    ($1) RETURNING *;

-- name: CreateWorkflow :one
INSERT INTO workflows (workflow_type_id, name, description) 
VALUES 
    ($1, $2, $3) RETURNING *;

-- name: CreateStep :one
INSERT INTO steps (name, description) 
VALUES 
    ($1, $2) RETURNING *;

-- name: CreateWorkflowStep :one
INSERT INTO workflow_steps (workflow_id, step_id, order_index)
VALUES 
    ($1, $2, $3) RETURNING *;
