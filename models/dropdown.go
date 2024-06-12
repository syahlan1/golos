package models

type Dropdown struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OwnershipDataDropdown struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	NoIdentity string `json:"no_identity"`
	NPWP       string `json:"npwp"`
	KeyPerson  bool   `json:"key_person"`
}
