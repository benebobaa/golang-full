// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateStep(ctx context.Context, arg CreateStepParams) (Step, error)
	CreateWorkflow(ctx context.Context, arg CreateWorkflowParams) (Workflow, error)
	CreateWorkflowInstance(ctx context.Context, arg CreateWorkflowInstanceParams) (WorkflowInstance, error)
	CreateWorkflowInstanceStep(ctx context.Context, arg CreateWorkflowInstanceStepParams) (WorkflowInstanceStep, error)
	CreateWorkflowStep(ctx context.Context, arg CreateWorkflowStepParams) (WorkflowStep, error)
	CreateWorkflowType(ctx context.Context, name string) (WorkflowType, error)
}

var _ Querier = (*Queries)(nil)
