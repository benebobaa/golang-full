// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
)

type Step struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type Workflow struct {
	ID             int32          `json:"id"`
	WorkflowTypeID int32          `json:"workflow_type_id"`
	Name           string         `json:"name"`
	Description    sql.NullString `json:"description"`
	CreatedAt      sql.NullTime   `json:"created_at"`
	UpdatedAt      sql.NullTime   `json:"updated_at"`
}

type WorkflowInstance struct {
	ID         int32        `json:"id"`
	WorkflowID int32        `json:"workflow_id"`
	Status     string       `json:"status"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type WorkflowInstanceStep struct {
	ID                 int32        `json:"id"`
	WorkflowInstanceID int32        `json:"workflow_instance_id"`
	WorkflowStepID     int32        `json:"workflow_step_id"`
	Status             string       `json:"status"`
	StartedAt          sql.NullTime `json:"started_at"`
	CompletedAt        sql.NullTime `json:"completed_at"`
}

type WorkflowStep struct {
	ID         int32 `json:"id"`
	WorkflowID int32 `json:"workflow_id"`
	StepID     int32 `json:"step_id"`
	OrderIndex int32 `json:"order_index"`
}

type WorkflowType struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
