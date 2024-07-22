package models

type MasterWorkflow struct {
	Id                       int    `json:"id" gorm:"primaryKey"`
	StatusDescription        string `json:"status_description"`
	StatusEnglishDescription string `json:"status_english_description"`
	StatusName               string `json:"status_name"`
	ModelMasterForm          `json:"-"`
}

type MasterWorkflowStep struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	NextStep        *int   `json:"next_step"`
	Step            int    `json:"step"`
	WorkflowId      int    `json:"workflow_id"`
	GroupId         int    `json:"group_id"`
	ModelMasterForm `json:"-"`
	Workflow        MasterWorkflow   `json:"-" gorm:"foreignKey:WorkflowId"`
	Group           MasterTableGroup `json:"-" gorm:"foreignKey:GroupId"`
}
