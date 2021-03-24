package model

type Task struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	EstimatedRequiredHours int `json:"estimatedRequiredHours"`
	Stack string `json:"stack"`
	MinimumRoles []string `json:"minimumRoles"`
	AssignedIdNumber int64 `json:"assignedIdNumber"`
	ParentTaskId int `json:"parentTaskId"`
}
