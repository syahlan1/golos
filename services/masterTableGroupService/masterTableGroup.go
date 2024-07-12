package masterTableGroupService

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/masterColumnService"
	"github.com/syahlan1/golos/services/masterTemplateService"
	"github.com/syahlan1/golos/utils"
	"gorm.io/gorm"
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
	data.ParentType = "C"

	err = connection.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return errors.New("failed to create Master Table Group")
		}

		var moduleName string

		if err := tx.
			Select("module_name").
			Model(&models.MasterModule{}).
			Where("id = ?", data.ModuleId).
			Scan(&moduleName).Error; err != nil {
			return err
		}

		var ParentMenu, ParentMenuAdmin models.Menu
		if err := tx.
			Select("*").
			Where("menu_code = ?", moduleName+"_menu").
			Find(&ParentMenu).
			Error; err != nil {
			return err
		}

		if err := tx.
			Select("*").
			Where("menu_code = ?", moduleName+"_admin").
			Find(&ParentMenuAdmin).
			Error; err != nil {
			return err
		}

		// var masterTableGroupMenu models.Menu
		if ParentMenu.Id == 0 || ParentMenuAdmin.Id == 0 {
			log.Println("Parent menu not found")
			var order int

			if err := tx.Select("order").
				Model(&models.Menu{}).
				Where("type = ?", "P").
				Order(`"order" desc`).
				Limit(1).
				Scan(&order).Error; err != nil {
				return err
			}

			if ParentMenu.Id == 0 {
				order++
				menuCode := moduleName + "_menu"
				ParentMenu = models.Menu{
					MenuCode: &menuCode,
					Icon:     "fa-table",
					Order:    order,
					Type:     "P",
					Label:    moduleName + " Menu",
					ModelMasterForm: models.ModelMasterForm{
						CreatedBy: "System",
					},
				}

				if err := tx.Create(&ParentMenu).Error; err != nil {
					return err
				}
			}

			if ParentMenuAdmin.Id == 0 {
				order++
				menuCode := moduleName + "_admin"
				ParentMenuAdmin = models.Menu{
					MenuCode: &menuCode,
					Icon:     "fa-table",
					Order:    order,
					Type:     "P",
					Label:    moduleName + " Admin",
					ModelMasterForm: models.ModelMasterForm{
						CreatedBy: "System",
					},
				}

				if err := tx.Create(&ParentMenuAdmin).Error; err != nil {
					return err
				}
			}
		}

		var ChildOrder int
		if err := tx.Select("order").
			Model(&models.Menu{}).
			Where("parent_id = ?", ParentMenu.Id).
			Order(`"order" desc`).
			Limit(1).
			Scan(&ChildOrder).Error; err != nil {
			return err
		}

		ChildMenu := models.Menu{
			ParentId: &ParentMenu.Id,
			Icon:     data.MenuIcon,
			Order:    ChildOrder + 1,
			Type:     "C",
			Label:    data.Description,
			MenuCode: &data.GroupName,
			ModelMasterForm: models.ModelMasterForm{
				CreatedBy: user.Username,
			},
		}

		ChildMenuFill := ChildMenu
		commandFill := "/" + utils.ToDashCase(moduleName) + "/fill/" + data.GroupName

		ChildMenuFill.Command = &commandFill

		if err := tx.Create(&ChildMenuFill).Error; err != nil {
			return err
		}

		ChildMenuAdmin := ChildMenu
		commandAdmin := "/" + utils.ToDashCase(moduleName) + "/admin/" + data.GroupName
		MenuCodeAdmin := data.GroupName + "_admin"

		ChildMenuAdmin.ParentId = &ParentMenuAdmin.Id
		ChildMenuAdmin.Command = &commandAdmin
		ChildMenuAdmin.MenuCode = &MenuCodeAdmin

		if err := tx.Create(&ChildMenuAdmin).Error; err != nil {
			return err
		}

		return nil
	})
	return err
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
	var masterTableGroup models.MasterTableGroup

	err = connection.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Where("id = ?", claims).First(&user).Error; err != nil {
			log.Println("Error retrieving user:", err)
			return err
		}

		if err := tx.First(&masterTableGroup, masterTableGroupId).Error; err != nil {
			return errors.New("data Not Found")
		}

		masterTableGroup.UpdatedBy = user.Username
		masterTableGroup.UpdatedAt = time.Now()
		masterTableGroup.Description = updatedMasterTableGroup.Description
		masterTableGroup.EnglishDescription = updatedMasterTableGroup.EnglishDescription
		masterTableGroup.MenuIcon = updatedMasterTableGroup.MenuIcon

		if err := tx.Save(&masterTableGroup).Error; err != nil {
			return errors.New("failed to update Master Table Group")
		}

		if err := tx.Table("menus").
			Where("menu_code = ? OR menu_code = ?", masterTableGroup.GroupName, masterTableGroup.GroupName+"_admin").
			Updates(map[string]interface{}{
				"icon":       updatedMasterTableGroup.MenuIcon,
				"updated_by": user.Username,
				"updated_at": time.Now(),
			}).Error; err != nil {
			return errors.New("failed to update Menu Icon")
		}

		return nil
	})
	if err != nil {
		return result, err
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

func ShowFormMasterTableGroup(groupName, username string) (result models.FormMasterTableGroupParent, err error) {

	if err = connection.DB.
		Select("master_table_groups.id, type, parent_type, description, english_description").
		Model(&models.MasterTableGroup{}).
		Where("group_name = ?", groupName).
		Find(&result).Error; err != nil {
		return result, err
	}

	if result.ParentType == "P" {
		result.Child, result.CanSubmit, err = scanRowsChild(result.Id, username, true)
	} else {
		result.Form, result.CanSubmit, err = scanRowsForm(result.Id, nil, username, true)
	}

	return
}

func scanRowsChild(parentId int, username string, canSubmitDefault bool) (result []models.FormMasterTableGroupParent, canSubmit bool, err error) {

	rows, err := connection.DB.
		Select("master_table_groups.id, parent_id, type, parent_type, description, english_description").
		Model(&models.MasterTableGroup{}).
		Where("parent_id = ?", parentId).
		Order(`"order"`).
		Rows()

	if err != nil {
		return result, canSubmit, err
	}

	defer rows.Close()

	var child models.FormMasterTableGroupParent
	child.CanSubmit = canSubmitDefault
	canSubmit = canSubmitDefault

	for rows.Next() {
		if err := connection.DB.ScanRows(rows, &child); err != nil {
			return result, canSubmit, err
		}
		child.Form, child.CanSubmit, err = scanRowsForm(child.Id, &parentId, username, child.CanSubmit)
		if err != nil {
			return result, canSubmit, err
		}

		if !child.CanSubmit {
			canSubmit = false
		}

		result = append(result, child)
	}

	return result, canSubmit, nil
}

func scanRowsForm(groupID int, parentId *int, username string, canSubmitDefault bool) (result []models.FormMasterTableItem, canSubmit bool, err error) {

	canSubmit = canSubmitDefault
	rows, err := connection.DB.
		Select("master_table_items.*, mt.id as table_id, mt.description as table_name").
		Joins("JOIN master_tables mt ON mt.id = master_table_items.table_id").
		Joins("JOIN master_table_groups mtg ON mtg.id = master_table_items.group_id").
		Model(&models.MasterTableItem{}).
		Where("mtg.id = ?", groupID).
		Order("master_table_items.sequence").
		Rows()

	if err != nil {
		return result, canSubmit, err
	}

	defer rows.Close()

	var item models.FormMasterTableItem

	for rows.Next() {
		if err := connection.DB.ScanRows(rows, &item); err != nil {
			return result, canSubmit, err
		}
		item.DataId, err = GetIdDataTableItem(groupID, item.Id, parentId, username)
		if err != nil {
			return result, canSubmit, err
		}
		FormList, err := masterColumnService.GetFormColumn(strconv.Itoa(item.TableId))
		if err != nil {
			return result, canSubmit, err
		}
		item.FormList = FormList.Form

		if len(item.DataId) == 0 || !(canSubmit) {
			canSubmit = false
		}

		result = append(result, item)
	}

	return result, canSubmit, nil
}

func GetIdDataTableItem(tableGroupId, tableItemId int, parentId *int, username string) (id []int, err error) {

	var schema, table string

	err = connection.DB.
		Select("mm.database_name, master_tables.table_name").
		Joins("JOIN master_table_items mti ON mti.table_id = master_tables.id").
		Joins("JOIN master_modules mm ON mm.id = master_tables.module_id").
		Joins("JOIN master_table_groups mtg ON mtg.id = mti.group_id").
		Model(&models.MasterTable{}).
		Where("mti.deleted_at is null AND mm.deleted_at is null").
		Where("mti.id = ? AND (mtg.id = ? OR mtg.parent_id = ?)", tableItemId, tableGroupId, parentId).Row().Scan(&schema, &table)
	if err != nil {
		return nil, err
	}

	if parentId != nil {
		tableGroupId = *parentId
	}

	err = connection.DB.
		Table(schema+"."+table).
		Where("deleted_date is null AND created_by = ? AND item_status_id is null AND "+schema+"_group_id = ?", username, tableGroupId).
		Order("id").
		Pluck("id", &id).Error

	return
}

func ShowDataMasterTableGroup(tableGroupId, tableItemId, username, id string) (data []map[string]interface{}, err error) {

	var schemaId, tableId string

	if err := connection.DB.
		Select("master_tables.module_id, master_tables.id").
		Joins("JOIN master_table_items mti ON mti.table_id = master_tables.id").
		Model(&models.MasterTable{}).
		Where("mti.deleted_at is null").
		Where("mti.id = ? AND mti.group_id = ?", tableItemId, tableGroupId).Row().Scan(&schemaId, &tableId); err != nil {
		return data, errors.New("data not found")
	}

	data, err = masterTemplateService.ShowMasterTemplate(schemaId, tableId, username, tableGroupId, "", "", id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func CreateDataMasterTableGroup(tableGroupId, tableItemId, username string, data map[string]interface{}) (err error, errValidation []models.Validate) {

	var schemaId, tableId string

	if err := connection.DB.
		Select("master_tables.module_id, master_tables.id").
		Joins("JOIN master_table_items mti ON mti.table_id = master_tables.id").
		Joins("JOIN master_table_groups mtg ON mtg.id = mti.group_id").
		Model(&models.MasterTable{}).
		Where("mti.deleted_at is null").
		Where("mti.id = ? AND (mtg.id = ? OR mtg.parent_id = ?)", tableItemId, tableGroupId, tableGroupId).Row().Scan(&schemaId, &tableId); err != nil {
		return errors.New("data not found"), nil
	}

	err, errValidation = masterTemplateService.CreateMasterTemplate(schemaId, tableId, username, tableGroupId, data)

	return
}

func SubmitTableGroupItem(username string, data models.TableGroupItemStatus) (err error) {

	db := connection.DB

	data.Status = "SUBMITTED"
	data.Username = username
	data.CreatedBy = username

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	var masterTable models.MasterTable
	rows, err := db.
		Select("master_tables.*, md.database_name as module_name").
		Joins("JOIN master_modules md ON md.id = master_tables.module_id").
		Joins("JOIN master_table_items mti ON mti.table_id = master_tables.id").
		Joins("JOIN master_table_groups mtg ON mtg.id = mti.group_id").
		Model(&models.MasterTable{}).
		Where("mti.deleted_at is null AND md.deleted_at is null AND (mtg.id = ? OR mtg.parent_id = ?)", data.GroupId, data.GroupId).
		Rows()
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

		if err := db.ScanRows(rows, &masterTable); err != nil {
			return err
		}

		err = db.
			Table(masterTable.ModuleName+"."+masterTable.TableName).
			Where("created_by = ? AND "+masterTable.ModuleName+"_group_id = ? AND item_status_id is null AND deleted_date is null", username, data.GroupId).
			Updates(map[string]interface{}{"item_status_id": data.Id}).Error
		if err != nil {
			return err
		}

	}

	return
}

func ShowApprovalTableGroupItem(groupName, username string) (data models.ShowApprovalTableGroup, err error) {

	var schemaId, tableId, tableGroupId, ParentType, ParentId string

	if err := connection.DB.
		Select("id, parent_type").
		Model(&models.MasterTableGroup{}).
		Where("group_name = ?", groupName).Row().Scan(&ParentId, &ParentType); err != nil {
		return data, errors.New("data not found")
	}

	db := connection.DB.
		Select("master_tables.module_id, master_tables.id, mti.group_id as table_group_id").
		Joins("JOIN master_table_items mti ON mti.table_id = master_tables.id").
		Model(&models.MasterTable{}).
		Where("mti.deleted_at is null").
		Where("mti.is_master = ?", true)

	if ParentType == "P" {

		db = db.
			Select("master_tables.module_id, master_tables.id, mtg.parent_id as table_group_id").
			Joins("JOIN master_table_groups mtg ON mtg.id = mti.group_id").
			Where("mtg.deleted_at is null").
			Where("mtg.parent_id = ?", ParentId).
			Order(`mtg.order`).
			Limit(1)

	} else {

		db = db.
			Where("mti.group_id = ?", ParentId)

	}

	if err = db.Row().Scan(&schemaId, &tableId, &tableGroupId); err != nil {
		return data, err
	}

	data.Submitted, err = masterTemplateService.ShowMasterTemplate(schemaId, tableId, "", tableGroupId, "SUBMITTED", "", "")
	if err != nil {
		return data, err
	}

	data.Rejected, err = masterTemplateService.ShowMasterTemplate(schemaId, tableId, "", tableGroupId, "REJECTED", "", "")
	if err != nil {
		return data, err
	}

	data.Approved, err = masterTemplateService.ShowMasterTemplate(schemaId, tableId, "", tableGroupId, "APPROVED", "", "")
	if err != nil {
		return data, err
	}

	return
}

func ShowDetailApprovalTableGroupItem(idApproval, username string) (data models.ShowDetailApprovalTableGroupParent, err error) {

	err = connection.DB.
		Select("table_group_item_statuses.*, mtg.description, mtg.english_description,mtg.parent_type").
		Model(&models.TableGroupItemStatus{}).
		Joins("JOIN master_table_groups mtg ON mtg.id = table_group_item_statuses.group_id").
		Where("table_group_item_statuses.id = ?", idApproval).
		First(&data).Error
	if err != nil {
		return
	}

	if data.ParentType == "P" {
		data.Child, err = showDetailApprovalTableGroupItemParent(data.GroupId, idApproval)
	} else {
		data.Data, err = showDetailApprovalTableGroupItemChild(data.GroupId, nil, idApproval)
	}

	return
}

func showDetailApprovalTableGroupItemParent(groupId int, idApproval string) (data []models.ShowDetailApprovalTableGroupParent, err error) {

	rows, err := connection.DB.
		Select("id,description, english_description,parent_type").
		Model(&models.MasterTableGroup{}).
		Where("parent_id = ?", groupId).
		Order(`"order"`).
		Rows()
	if err != nil {
		return
	}

	defer rows.Close()

	var child models.ShowDetailApprovalTableGroupParent
	for rows.Next() {

		if err := connection.DB.ScanRows(rows, &child); err != nil {
			return data, err
		}

		child.Data, err = showDetailApprovalTableGroupItemChild(child.Id, &groupId, idApproval)
		if err != nil {
			return
		}

		data = append(data, child)
	}

	return
}

func showDetailApprovalTableGroupItemChild(groupId int, parentId *int, idApproval string) (data []models.DataDetailApprovalTableGroup, err error) {
	var dataTable models.DataDetailApprovalTableGroup

	rows, err := connection.DB.
		Select("master_tables.description as table_name, mti.type, master_tables.module_id AS schema_id, master_tables.id as table_id").
		Joins("JOIN master_table_items mti ON mti.table_id = master_tables.id").
		Model(&models.MasterTable{}).
		Where("mti.deleted_at is null").
		Where("mti.group_id = ?", groupId).
		Order("mti.sequence").
		Rows()
	if err != nil {
		return
	}

	if parentId != nil {
		groupId = *parentId
	}

	defer rows.Close()

	for rows.Next() {
		if err := connection.DB.ScanRows(rows, &dataTable); err != nil {
			return data, err
		}

		dataTable.Data, err = masterTemplateService.ShowMasterTemplate(dataTable.SchemaId, dataTable.TableId, "", strconv.Itoa(groupId), "", idApproval, "")

		data = append(data, dataTable)

	}
	return
}

func ApprovalTableGroupItem(username string, data models.TableGroupItemStatus) (err error) {

	db := connection.DB

	data.Username = username
	data.UpdatedBy = username

	err = db.Model(&data).Where("id = ?", data.Id).Updates(data).Error

	return
}

func UpdateDataMasterTableGroup(tableGroupId, tableItemId, idData, username string, data map[string]interface{}) (err error, errValidation []models.Validate) {
	var schemaId, tableId string

	if err := connection.DB.
		Select("master_tables.module_id, master_tables.id").
		Joins("JOIN master_table_items mti ON mti.table_id = master_tables.id").
		Joins("JOIN master_table_groups mtg ON mtg.id = mti.group_id").
		Model(&models.MasterTable{}).
		Where("mti.deleted_at is null").
		Where("mti.id = ? AND (mtg.id = ? OR mtg.parent_id = ?)", tableItemId, tableGroupId, tableGroupId).Row().Scan(&schemaId, &tableId); err != nil {
		return errors.New("data not found"), nil
	}

	err, errValidation = masterTemplateService.UpdateMasterTemplate(schemaId, tableId, idData, username, data)

	return
}

func DeleteDataMasterTableGroup(tableGroupId, tableItemId, idData, username string) (err error) {

	var schemaId, tableId string

	if err := connection.DB.
		Select("master_tables.module_id, master_tables.id").
		Joins("JOIN master_table_items mti ON mti.table_id = master_tables.id").
		Joins("JOIN master_table_groups mtg ON mtg.id = mti.group_id").
		Model(&models.MasterTable{}).
		Where("mti.deleted_at is null").
		Where("mti.id = ? AND (mtg.id = ? OR mtg.parent_id = ?)", tableItemId, tableGroupId, tableGroupId).Row().Scan(&schemaId, &tableId); err != nil {
		return errors.New("data not found")
	}

	err = masterTemplateService.DeleteMasterTemplate(schemaId, tableId, idData, username)

	return
}

func GenerateTableGroup(tableID string) (err error) {
	db := connection.DB

	var masterTable models.MasterTable
	if err := db.Select("master_tables.*, md.database_name as module_name").
		Joins("JOIN master_modules md ON md.id = master_tables.module_id").
		First(&masterTable, tableID).Error; err != nil {
		return errors.New("data not found")
	}

	groupColumn := []models.MasterColumn{
		{
			FieldName: masterTable.ModuleName + "_group_id",
			FieldType: "INTEGER",
		},
		{
			FieldName: "item_status_id",
			FieldType: "INTEGER",
		},
	}

	for _, v := range groupColumn {
		var checkTable bool
		err = db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_schema = ? AND table_name = ? AND column_name = ?)", masterTable.ModuleName, masterTable.TableName, v.FieldName).Scan(&checkTable).Error
		if err != nil {
			return err
		}

		if !checkTable {
			query := fmt.Sprintf(`ALTER TABLE "%s"."%s" ADD COLUMN %s %s`, masterTable.ModuleName, masterTable.TableName, v.FieldName, v.FieldType)
			if err := db.Exec(query).Error; err != nil {
				return err
			}
		}
	}

	return
}

func GenerateTableGroupByGroupId(tableGroupId string) (err error) {

	db := connection.DB

	var masterTable models.MasterTable

	rows, err := db.Debug().
		Select("master_tables.*, md.database_name as module_name").
		Joins("JOIN master_modules md ON md.id = master_tables.module_id").
		Joins("JOIN master_table_items mti ON mti.table_id = master_tables.id").
		Model(&models.MasterTable{}).
		Where("mti.deleted_at is null AND md.deleted_at is null AND mti.group_id = ?", tableGroupId).
		Rows()
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

		if err := db.ScanRows(rows, &masterTable); err != nil {
			return err
		}

		groupColumn := []models.MasterColumn{
			{
				FieldName: masterTable.ModuleName + "_group_id",
				FieldType: "INTEGER",
			},
			{
				FieldName: "item_status_id",
				FieldType: "INTEGER",
			},
		}

		var alterTableQueries []string

		for _, v := range groupColumn {
			var checkTable bool
			err = db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_schema = ? AND table_name = ? AND column_name = ?)", masterTable.ModuleName, masterTable.TableName, v.FieldName).Scan(&checkTable).Error
			if err != nil {
				return err
			}

			if !checkTable {
				alterTableQueries = append(alterTableQueries, fmt.Sprintf("ADD COLUMN %s %s", v.FieldName, v.FieldType))
			}
		}

		if len(alterTableQueries) > 0 {
			alterTableSQL := fmt.Sprintf(`ALTER TABLE "%s"."%s"`+"\n"+`%s;`, masterTable.ModuleName, masterTable.TableName, strings.Join(alterTableQueries, ",\n"))
			if err := db.Exec(alterTableSQL).Error; err != nil {
				return err
			}
		}

	}

	return
}
