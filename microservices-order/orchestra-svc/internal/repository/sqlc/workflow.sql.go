// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: workflow.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createStep = `-- name: CreateStep :one
INSERT INTO steps (name, description) 
VALUES 
    ($1, $2) RETURNING id, name, description, created_at, updated_at
`

type CreateStepParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateStep(ctx context.Context, arg CreateStepParams) (Step, error) {
	row := q.db.QueryRowContext(ctx, createStep, arg.Name, arg.Description)
	var i Step
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createWorkflow = `-- name: CreateWorkflow :one
INSERT INTO workflows (workflow_type_id, name, description) 
VALUES 
    ($1, $2, $3) RETURNING id, workflow_type_id, name, description, created_at, updated_at
`

type CreateWorkflowParams struct {
	WorkflowTypeID int32          `json:"workflow_type_id"`
	Name           string         `json:"name"`
	Description    sql.NullString `json:"description"`
}

func (q *Queries) CreateWorkflow(ctx context.Context, arg CreateWorkflowParams) (Workflow, error) {
	row := q.db.QueryRowContext(ctx, createWorkflow, arg.WorkflowTypeID, arg.Name, arg.Description)
	var i Workflow
	err := row.Scan(
		&i.ID,
		&i.WorkflowTypeID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createWorkflowStep = `-- name: CreateWorkflowStep :one
INSERT INTO workflow_steps (workflow_id, step_id, order_index)
VALUES 
    ($1, $2, $3) RETURNING id, workflow_id, step_id, order_index
`

type CreateWorkflowStepParams struct {
	WorkflowID int32 `json:"workflow_id"`
	StepID     int32 `json:"step_id"`
	OrderIndex int32 `json:"order_index"`
}

func (q *Queries) CreateWorkflowStep(ctx context.Context, arg CreateWorkflowStepParams) (WorkflowStep, error) {
	row := q.db.QueryRowContext(ctx, createWorkflowStep, arg.WorkflowID, arg.StepID, arg.OrderIndex)
	var i WorkflowStep
	err := row.Scan(
		&i.ID,
		&i.WorkflowID,
		&i.StepID,
		&i.OrderIndex,
	)
	return i, err
}

const createWorkflowType = `-- name: CreateWorkflowType :one
INSERT INTO workflow_types (name) 
VALUES 
    ($1) RETURNING id, name
`

func (q *Queries) CreateWorkflowType(ctx context.Context, name string) (WorkflowType, error) {
	row := q.db.QueryRowContext(ctx, createWorkflowType, name)
	var i WorkflowType
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
