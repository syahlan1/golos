package masterFileService

import (
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func UploadFile(file *multipart.FileHeader, groupId, moduleId, username string) (result models.MasterFile, err error) {

	var paramPath models.MasterParameter
	if groupId != "" {
		connection.DB.
			Joins("JOIN master_table_groups mtg ON mtg.module_id = master_parameters.module_id").
			Where("param_key = ? AND mtg.id = ? AND mtg.deleted_at is null", "DOC_PATH", groupId).First(&paramPath)
	} else {
		connection.DB.Where("param_key = ? AND module_id = ?", "DOC_PATH", moduleId).First(&paramPath)
	}

	if paramPath.Id == 0 {
		connection.DB.Where("param_key = ? AND module_id = 0", "DOC_PATH").First(&paramPath)
	}

	if paramPath.Id == 0 {
		return result, errors.New("DOC_PATH must be configured")
	}

	if paramPath.IsEncrypted == 1 {
		hashed, err := utils.Decrypt(paramPath.ParamValue)
		if err != nil {
			return result, err
		}
		paramPath.ParamValue = string(hashed)
	}

	result.FileName, result.FilePath, err = utils.UploadFile(file, paramPath.ParamValue)
	if err != nil {
		return result, err
	}
	result.File = filepath.Base(result.FilePath)
	result.FileType = utils.GetFileExtension(result.FileName)
	result.CreatedBy = username

	err = connection.DB.Create(&result).Error
	if err != nil {
		return result, errors.New("failed to create the file : " + err.Error())
	}

	return
}
