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
		// data.ParentId = nil
		// data.MenuCode = nil
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
		Select("id, parent_id, icon, label as title, command as path").
		Model(models.Menu{}).
		Where("parent_id is null").
		Order(`"order" asc`).
		Find(&result).Error; err != nil {
		return result, err
	}

	for i, data := range result {
		child, err := findChild(data.Id)
		if err != nil {
			return result, err
		}
		result[i].Subnav = child
	}

	return result, nil
}

func findChild (parentId int) ([]models.ShowMenu, error) {
	var child []models.ShowMenu
	if err := connection.DB.Debug().
		Select("id, parent_id, icon, label as title, command as path").
		Model(models.Menu{}).
		Where("parent_id = ?", parentId).
		Order(`"order" asc`).
		Find(&child).Error; err != nil {
		return nil, err
	}

	for i, data := range child {
		subChild, err := findChild(data.Id)
		if err != nil {
			return nil, err
		}
		child[i].Subnav = subChild
	}

	return child, nil
}
