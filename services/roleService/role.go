package roleService

import (
	"errors"
	"log"

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

func ShowRoleModules(roleId string) (result []models.RoleModules, err error) {
	
	return
}
