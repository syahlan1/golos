package documentService

import (
	"errors"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func CreateDocument(data *models.Document) (err error) {
	return connection.DB.Create(&data).Error
}

func ShowDocumentById(id int) (result models.Document, err error) {

	if err = connection.DB.First(&result, "id = ?", id).Error; err != nil {
		return
	}
	return
}

func UpdateDocument(id int, data models.Document) (err error) {

	var document models.Document
	if err := connection.DB.First(&document, id).Error; err != nil {
		return errors.New("document not found")
	}
	updatedDocument := data

	document.DocumentFile = updatedDocument.DocumentFile
	document.DocumentPath = updatedDocument.DocumentPath
	document.Status = updatedDocument.Status
	document.NoCreditSalesForm = updatedDocument.NoCreditSalesForm
	document.DateOfLetter = updatedDocument.DateOfLetter
	document.DateOfReceipt = updatedDocument.DateOfReceipt

	if err := connection.DB.Save(&document).Error; err != nil {
		return errors.New("failed to update the document data")
	}
	return
}
