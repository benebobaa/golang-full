-- name: CreateWorkflowInstance :one
INSERT INTO workflow_instances (workflow_id, status) 
VALUES 
    ($1, $2) RETURNING *;

-- name: CreateWorkflowInstanceStep :one
INSERT INTO workflow_instance_steps (workflow_instance_id, workflow_step_id, status, started_at, completed_at) 
VALUES 
    ($1, $2, $3, $4, $5) RETURNING *;
