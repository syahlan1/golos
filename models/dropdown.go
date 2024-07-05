package models

type Dropdown struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DropdownEn struct {
	Code               string `json:"code"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
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
