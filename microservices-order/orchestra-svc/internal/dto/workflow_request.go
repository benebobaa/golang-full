package dto

type WorkflowRequest struct {
	Name     string   `json:"name" valo:"notblank"`
	Workflow Workflow `json:"workflow" valo:"notnil,valid"`
	Step     []Step   `json:"step" valo:"valid,notnil"`
}

type Workflow struct {
	Name        string `json:"name" valo:"notblank"`
	Description string `json:"description" valo:"notblank"`
}

type Step struct {
	Name        string `json:"name" valo:"notblank"`
	Description string `json:"description" valo:"notblank"`
}
