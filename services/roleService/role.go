package roleService

import (
	"errors"
	"log"
	"strconv"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"gorm.io/gorm"
)

func ShowRole() (result []models.Roles, err error) {
	var roles []models.Roles
	if err := connection.DB.Find(&roles).Error; err != nil {
		return result, err
	}

	return roles, nil
}

func ShowAllPermissions() (result []models.Permission, err error) {
	var permissions []models.Permission
	if err := connection.DB.Find(&permissions).Error; err != nil {
		return result, err
	}

	return result, nil
}

func ShowPermissions(roleID string) (result models.Roles, err error) {
	var role models.Roles
	if err := connection.DB.Preload("Permissions").Where("id = ?", roleID).First(&role).Error; err != nil {
		return result, err
	}

	return role, nil
}

func CreateRole(userId string, data models.CreateRole) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	// Cek apakah role dengan nama yang sama sudah ada
	var existingRole models.Roles
	if err := connection.DB.Where("name = ?", data.Name).First(&existingRole).Error; err == nil {
		// Role sudah ada, kirim respons konflik
		return errors.New("role already exists")
	}

	// Buat role baru
	newRole := models.Roles{
		Name:            data.Name,
		Description:     data.Description,
		ModelMasterForm: models.ModelMasterForm{CreatedBy: user.Username},
	}

	// Simpan role baru ke dalam database
	if err := connection.DB.Create(&newRole).Error; err != nil {
		return err
	}

	// Dapatkan ID permissions berdasarkan nama permissions yang diberikan
	var permissions []models.Permission
	if err := connection.DB.Where("name IN ?", data.Permissions).Find(&permissions).Error; err != nil {
		return err
	}

	// Buat entri RolePermission untuk setiap permission yang terkait dengan role baru
	for _, permission := range permissions {
		// Periksa apakah RolePermission sudah ada
		var existingRolePermission models.RolePermission
		if err := connection.DB.Where("roles_id = ? AND permission_id = ?", newRole.Id, permission.Id).First(&existingRolePermission).Error; err != nil {
			// RolePermission belum ada, tambahkan entri baru
			rolePermission := models.RolePermission{RolesId: newRole.Id, PermissionId: permission.Id}
			if err := connection.DB.Create(&rolePermission).Error; err != nil {
				return err
			}
		}
	}

	return nil

}

func DeleteRole(roleID string) (err error) {
	tx := connection.DB.Begin()

	// Hapus role permission yang terkait dengan role yang akan dihapus
	if err := tx.Where("roles_id = ?", roleID).Delete(&models.RolePermission{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update role_id user menjadi 0 untuk user yang memiliki role yang akan dihapus
	if err := tx.Model(&models.Users{}).Where("role_id = ?", roleID).Update("role_id", 0).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Hapus role
	if err := tx.Where("id = ?", roleID).Delete(&models.Roles{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaksi
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func UpdateRole(userID, roleID string, data models.CreateRole) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	// Find existing role by ID
	var existingRole models.Roles
	if err := connection.DB.Where("id = ?", roleID).First(&existingRole).Error; err != nil {
		return errors.New("role not found")
	}

	// Update role name
	existingRole.Name = data.Name
	existingRole.Description = data.Description
	existingRole.UpdatedBy = user.Username

	// Save the updated role to database
	if err := connection.DB.Save(&existingRole).Error; err != nil {
		return err
	}

	// Get IDs of permissions from the input
	var permissions []models.Permission
	if err := connection.DB.Where("name IN ?", data.Permissions).Find(&permissions).Error; err != nil {
		return err
	}

	// Collect IDs of permissions from the input
	var permissionIDs []uint
	for _, permission := range permissions {
		permissionIDs = append(permissionIDs, permission.Id)
	}

	// Update RolePermission entries for the role
	// Update existing RolePermission entries based on input permissions
	for _, permission := range permissions {
		// Check if the RolePermission already exists
		var existingRolePermission models.RolePermission
		err := connection.DB.Where("roles_id = ? AND permission_id = ?", existingRole.Id, permission.Id).First(&existingRolePermission).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// RolePermission doesn't exist, create a new one
			rolePermission := models.RolePermission{RolesId: existingRole.Id, PermissionId: permission.Id}
			if err := connection.DB.Create(&rolePermission).Error; err != nil {
				return err
			}
		}
	}

	// Delete existing RolePermission entries not present in data.Permissions
	if err := connection.DB.Where("roles_id = ? AND permission_id NOT IN ?", existingRole.Id, permissionIDs).Delete(&models.RolePermission{}).Error; err != nil {
		return err
	}

	return nil
}

func CreateRoleModules(userId string, data models.CreateRoleModules) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	var roleModules []models.RoleModules
	for _, moduleId := range data.ModuleId {
		roleModule := models.RoleModules{
			RolesId:         data.RolesId,
			ModuleId:        moduleId,
			ModelMasterForm: models.ModelMasterForm{CreatedBy: user.Username},
		}
		roleModules = append(roleModules, roleModule)
	}

	if err := connection.DB.Create(&roleModules).Error; err != nil {
		return err
	}

	return
}

func ShowAllRoleMenu(roleId string) (result []models.ShowRoleMenu, err error) {

	if err := connection.DB.
		Select("rm.id, menus.id as menu_id, menus.menu_code as menu, COALESCE(rm.read, FALSE) as read",
			"COALESCE(rm.delete, FALSE) as delete, COALESCE(rm.update, FALSE) as update",
			"COALESCE(rm.download, FALSE) as download, COALESCE(rm.write, FALSE) as write").
		Joins("left join role_menus rm on rm.menu_id = menus.id and rm.role_id = ?", roleId).
		Where("menus.type = 'C' and rm.deleted_at is null").
		Model(models.Menu{}).
		Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func CreateRoleMenu(claims, roleId string, data []models.RoleMenu) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	roleIdInt, _ := strconv.Atoi(roleId)

	var RoleMenus []models.RoleMenu
	for _, value := range data {

		RoleMenu := models.RoleMenu{
			Id:              value.Id,
			RoleId:          roleIdInt,
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

func ShowRoleModules(roleId string) (result []models.ShowRoleModules, err error) {

	if err := connection.DB.Select("rm.id, master_modules.id AS module_id, master_modules.module_name AS module").
		Joins("left join role_modules rm on rm.module_id = master_modules.id AND rm.roles_id = ?", roleId).
		Where("rm.deleted_at is null").
		Model(models.MasterModule{}).
		Find(&result).Error; err != nil {
		return result, err
	}

	for i, value := range result {

		rows, err := connection.DB.Select("COALESCE(rt.id, 0), master_tables.id as table_id , master_tables.table_name as table",
			"COALESCE(rt.read, FALSE) as read, COALESCE(rt.delete, FALSE) as delete, COALESCE(rt.update, FALSE) as update",
			"COALESCE(rt.download, FALSE) as download, COALESCE(rt.write, FALSE) as write ",
			"(case when read is true or delete is true or update is true or write is true then true else false end) as selected").
			Joins("left join role_tables rt on rt.table_id = master_tables.id AND rt.role_modules_id = ?", value.Id).
			Model(models.MasterTable{}).
			Where("master_tables.module_id = ? and rt.deleted_at is null", value.ModuleId).
			Rows()

		if err != nil {
			return result, err
		}
		defer rows.Close()

		var data models.ShowRoleTables
		for rows.Next() {
			if err := rows.Scan(&data.Id, &data.TableId, &data.Table, &data.Read, &data.Delete, &data.Update, &data.Download, &data.Write, &data.Selected); err != nil {
				return result, err
			}

			if data.Selected {
				result[i].TableSelected++
			}

			result[i].Table = append(result[i].Table, data)
		}

	}

	return result, nil
}

func CreateRoleModuleTables(userId, roleId string, data []models.CreateRoleModuleTables) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	roleIdInt, _ := strconv.Atoi(roleId)

	err = connection.DB.Transaction(func(tx *gorm.DB) error {

		for _, value := range data {
			roleModule := models.RoleModules{
				Id:              value.Id,
				RolesId:         uint(roleIdInt),
				ModuleId:        value.ModuleId,
				ModelMasterForm: models.ModelMasterForm{CreatedBy: user.Username},
			}

			if err := tx.Save(&roleModule).Error; err != nil {
				return err
			}

			for i := range value.Tables {
				value.Tables[i].RoleModulesId = roleModule.Id
				value.Tables[i].CreatedBy = user.Username
			}

			if err := tx.Save(&value.Tables).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return
}

func ShowRoleWorkflows(roleId string) (result models.ShowRoleWorkflows, err error) {

	if err := connection.DB.Select("rw.id, master_workflows.id as workflow_id,CONCAT(master_workflows.status_name, ' - ', master_workflows.status_description) as name, COALESCE(rw.selected, FALSE) as selected").
		Joins("left join role_workflows rw on rw.workflow_id = master_workflows.id AND rw.roles_id = ?", roleId).
		Where("rw.deleted_at is null").
		Model(models.MasterWorkflow{}).
		Scan(&result.All).Error; err != nil {
		return result, err
	}

	if err := connection.DB.Select("rw.id, master_workflows.id as workflow_id,CONCAT(master_workflows.status_name, ' - ', master_workflows.status_description) as name, COALESCE(rw.selected, FALSE) as selected").
		Joins("left join role_workflows rw on rw.workflow_id = master_workflows.id AND rw.roles_id = ?", roleId).
		Where("rw.deleted_at is null and rw.selected = true").
		Model(models.MasterWorkflow{}).
		Scan(&result.Selected).Error; err != nil {
		return result, err
	}

	return
}

func CreateRoleWorkflows(userId, roleId string, data []models.RoleWorkflowDropdown) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	roleIdInt, _ := strconv.Atoi(roleId)

	err = connection.DB.Transaction(func(tx *gorm.DB) error {
		log.Println("data", roleIdInt)
		var roleWorkflows []models.RoleWorkflow
		err = connection.DB.Model(&models.RoleWorkflow{}).
			Where("roles_id = ?", roleIdInt).
			Update("selected", false).Error
		if err != nil {
			return err
		}

		for _, value := range data {
			roleWorkflow := models.RoleWorkflow{
				Id:              uint(value.Id),
				RolesId:         uint(roleIdInt),
				WorkflowId:      value.WorkflowId,
				Selected:        true,
				ModelMasterForm: models.ModelMasterForm{CreatedBy: user.Username},
			}

			roleWorkflows = append(roleWorkflows, roleWorkflow)
		}

		if err := tx.Save(&roleWorkflows).Error; err != nil {
			return err
		}
		return nil
	})

	return
}
