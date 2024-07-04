package models

type Dropdown struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DropdownEn struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

type OwnershipDataDropdown struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	NoIdentity string `json:"no_identity"`
	NPWP       string `json:"npwp"`
	KeyPerson  bool   `json:"key_person"`
}

type RoleWorkflowDropdown struct {
	Id         int    `json:"id"`
	WorkflowId int    `json:"workflow_id"`
	Name       string `json:"name"`
	Selected   bool   `json:"selected"`
}
