// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: workflow_instance.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createWorkflowInstance = `-- name: CreateWorkflowInstance :one
INSERT INTO workflow_instances (workflow_id, status) 
VALUES 
    ($1, $2) RETURNING id, workflow_id, status, created_at, updated_at
`

type CreateWorkflowInstanceParams struct {
	WorkflowID int32  `json:"workflow_id"`
	Status     string `json:"status"`
}

func (q *Queries) CreateWorkflowInstance(ctx context.Context, arg CreateWorkflowInstanceParams) (WorkflowInstance, error) {
	row := q.db.QueryRowContext(ctx, createWorkflowInstance, arg.WorkflowID, arg.Status)
	var i WorkflowInstance
	err := row.Scan(
		&i.ID,
		&i.WorkflowID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createWorkflowInstanceStep = `-- name: CreateWorkflowInstanceStep :one
INSERT INTO workflow_instance_steps (workflow_instance_id, workflow_step_id, status, started_at, completed_at) 
VALUES 
    ($1, $2, $3, $4, $5) RETURNING id, workflow_instance_id, workflow_step_id, status, started_at, completed_at
`

type CreateWorkflowInstanceStepParams struct {
	WorkflowInstanceID int32        `json:"workflow_instance_id"`
	WorkflowStepID     int32        `json:"workflow_step_id"`
	Status             string       `json:"status"`
	StartedAt          sql.NullTime `json:"started_at"`
	CompletedAt        sql.NullTime `json:"completed_at"`
}

func (q *Queries) CreateWorkflowInstanceStep(ctx context.Context, arg CreateWorkflowInstanceStepParams) (WorkflowInstanceStep, error) {
	row := q.db.QueryRowContext(ctx, createWorkflowInstanceStep,
		arg.WorkflowInstanceID,
		arg.WorkflowStepID,
		arg.Status,
		arg.StartedAt,
		arg.CompletedAt,
	)
	var i WorkflowInstanceStep
	err := row.Scan(
		&i.ID,
		&i.WorkflowInstanceID,
		&i.WorkflowStepID,
		&i.Status,
		&i.StartedAt,
		&i.CompletedAt,
	)
	return i, err
}