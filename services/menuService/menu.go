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

func ShowAllRoleMenu(roleId string) (result []models.ShowRoleMenu, err error) {
	if err := connection.DB.Raw(`
			SELECT * FROM (
				SELECT *, (ROW_NUMBER() OVER(PARTITION BY menu_id, menu ORDER BY id)) AS rn
				FROM (
					SELECT rm.id, menus.id as menu_id, menus.menu_code as menu, COALESCE(rm.read, FALSE) as read, 
						COALESCE(rm.delete, FALSE) as delete, COALESCE(rm.update, FALSE) as update, 
						COALESCE(rm.download, FALSE) as download, COALESCE(rm.write, FALSE) as write 
					FROM menus
					LEFT JOIN role_menu rm ON rm.menu_id = menus.id  
					JOIN roles r ON r.id = rm.role_id
					WHERE type = 'C' AND r.id = ? AND menus.deleted_at IS NULL
					UNION
					SELECT NULL as id, menus.id as menu_id, menus.menu_code as menu, 
						FALSE as read, FALSE as delete, FALSE as update, FALSE as download, FALSE as write 
					FROM menus
					WHERE type = 'C' AND menus.deleted_at IS NULL
				) AS combined
			) AS numbered
			WHERE rn = 1`, roleId).Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func CreateRoleMenu(claims string, data models.CreateRoleMenu) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	var RoleMenus []models.RoleMenu
	for _, value := range data.RoleMenu {

		RoleMenu := models.RoleMenu{
			Id:              value.Id,
			RoleId:          data.RoleId,
			MenuId:          value.MenuId,
			Read:            value.Read,
			Delete:          value.Delete,
			Update:          value.Update,
			Download:        value.Download,
			Write:           value.Write,
			ModelMasterForm: models.ModelMasterForm{CreatedBy: user.Username},
		}

		RoleMenus = append(RoleMenus, RoleMenu)
	}

	if err := connection.DB.Save(&RoleMenus).Error; err != nil {
		return err
	}

	return
}
