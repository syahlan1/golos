package masterTableGroupService

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

// func ShowMenuIcons() (result []models.Dropdown) {
// 	return
// }

func CreateMasterTableGroup(claims string, data models.MasterTableGroup) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	data.CreatedBy = user.Username

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Master Table Group")
	}

	return nil
}

func ShowMasterTableGroup() (result []models.MasterTableGroup) {
	var masterTableGroup []models.MasterTableGroup

	connection.DB.Find(&masterTableGroup)

	return masterTableGroup
}

func ShowMasterTableGroupDetail(masterTableGroupId string) (result models.MasterTableGroup, err error) {
	var masterTableGroup models.MasterTableGroup

	if err := connection.DB.First(&masterTableGroup, masterTableGroupId).Error; err != nil {
		return result, errors.New("MasterTableGroup not found")
	}

	return masterTableGroup, nil
}

func UpdateMasterTableGroup(claims, masterTableGroupId string, updatedMasterTableGroup models.MasterTableGroup) (result models.MasterTableGroup, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterTableGroup models.MasterTableGroup
	if err := connection.DB.First(&masterTableGroup, masterTableGroupId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	masterTableGroup.UpdatedBy = user.Username
	masterTableGroup.UpdatedAt = time.Now()
	masterTableGroup.Description = updatedMasterTableGroup.Description
	masterTableGroup.EnglishDescription = updatedMasterTableGroup.EnglishDescription
	masterTableGroup.MenuIcon = updatedMasterTableGroup.MenuIcon

	if err := connection.DB.Save(&masterTableGroup).Error; err != nil {
		return result, errors.New("failed to update Master Table Group")
	}

	return masterTableGroup, nil
}

func DeleteMasterTableGroup(claims, masterTableGroupId string) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	var masterTableGroup models.MasterTableGroup

	masterTableGroup.ModelMasterForm = utils.SoftDelete(user.Username)
	return connection.DB.Model(&masterTableGroup).Where("id = ?", masterTableGroupId).Updates(&masterTableGroup).Error
}

// Master Table Item
func CreateMasterTableItem(claims string, data models.MasterTableItem) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	data.CreatedBy = user.Username

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Master Table Item")
	}

	return nil
}

func ShowMasterTableItem(groupId string) (result []models.MasterTableItem) {
	var masterTableItem []models.MasterTableItem

	connection.DB.Where("group_id = ?", groupId).Find(&masterTableItem)

	return masterTableItem
}

func ShowMasterTableItemDetail(masterTableItemId string) (result models.MasterTableItem, err error) {
	var masterTableItem models.MasterTableItem

	if err := connection.DB.First(&masterTableItem, masterTableItemId).Error; err != nil {
		return result, errors.New("MasterTableItem not found")
	}

	return masterTableItem, nil
}

func UpdateMasterTableItem(claims, masterTableItemId string, updatedMasterTableItem models.MasterTableItem) (result models.MasterTableItem, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterTableItem models.MasterTableItem
	if err := connection.DB.First(&masterTableItem, masterTableItemId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	masterTableItem.UpdatedBy = user.Username
	masterTableItem.UpdatedAt = time.Now()
	masterTableItem.Name = updatedMasterTableItem.Name
	masterTableItem.Sequence = updatedMasterTableItem.Sequence
	masterTableItem.Type = updatedMasterTableItem.Type
	masterTableItem.IsMaster = updatedMasterTableItem.IsMaster

	if err := connection.DB.Save(&masterTableItem).Error; err != nil {
		return result, errors.New("failed to update Master Table Item")
	}

	return masterTableItem, nil
}

func DeleteMasterTableItem(claims, masterTableItemId string) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	var masterTableItem models.MasterTableItem

	masterTableItem.ModelMasterForm = utils.SoftDelete(user.Username)
	return connection.DB.Model(&masterTableItem).Where("id = ?", masterTableItemId).Updates(&masterTableItem).Error
}

func GenerateTableGroup(tableID string) (err error) {
	db := connection.DB

	var masterTable models.MasterTable
	if err := db.Select("master_tables.*, md.database_name as module_name").
		Joins("JOIN master_modules md ON md.id = master_tables.module_id").
		First(&masterTable, tableID).Error; err != nil {
		return errors.New("data not found")
	}

	groupColumn := models.MasterColumn{
		FieldName: masterTable.ModuleName + "_group_id",
		FieldType: "INTEGER",
	}

	var checkTable bool
	err = db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_schema = ? AND table_name = ? AND column_name = ?)", masterTable.ModuleName, masterTable.TableName, groupColumn.FieldName).Scan(&checkTable).Error
	if err != nil {
		return err
	}

	if !checkTable {
		query := fmt.Sprintf(`ALTER TABLE "%s"."%s" ADD COLUMN %s %s`, masterTable.ModuleName, masterTable.TableName, groupColumn.FieldName, groupColumn.FieldType)
		if err := db.Exec(query).Error; err != nil {
			return err
		}
	}

	return
}
