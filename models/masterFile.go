package models

type MasterFile struct {
	Id              int    `json:"id"`
	File            string `json:"file"`
	FileName        string `json:"file_name"`
	FilePath        string `json:"file_path"`
	FileType        string `json:"file_type"`
	ModelMasterForm `json:"-"`
}
