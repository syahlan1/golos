package menuService

import (
	"errors"
	"log"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func CreateMenu(claims string, data models.Menu) (err error) {

	if data.Type == "C" {

	} else if data.Type == "P" {
		data.Command = nil
		data.ParentId = nil
		data.MenuCode = nil
	} else {
		return errors.New("invalid type")
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	data.CreatedBy = user.Username

	if err := connection.DB.Create(&data).Error; err != nil {
		return err
	}

	return
}

func ShowMenu() (result []models.ShowMenu, err error) {

	if err := connection.DB.
		Model(models.Menu{}).
		Where("type = ?", "P").
		Order(`"order" asc`).
		Find(&result).Error; err != nil {
		return result, err
	}

	for i, data := range result {

		var child []models.Menu

		if err := connection.DB.
			Where("parent_id = ?", data.Id).
			Order(`"order" asc`).
			Find(&child).Error; err != nil {
			return result, err
		}

		result[i].Child = child
	}

	return result, nil
}


