package authService

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func getParameterValue(key string) (string, error) {
	var param models.MasterParameter
	if err := connection.DB.Where("param_key = ? AND status = ?", key, "L").First(&param).Error; err != nil {
		return "", err
	}
	return param.ParamValue, nil
}

func Register(data models.Register) (err error) {
	minLengthStr, err := getParameterValue("USR_MIN")
	if err != nil {
		return errors.New("failed to get USR_MIN parameter")
	}

	minLength, err := strconv.Atoi(minLengthStr)
	if err != nil {
		return errors.New("invalid USR_MIN parameter value")
	}

	validChars, err := getParameterValue("USR_CHAR")
	if err != nil {
		return errors.New("failed to get USR_CHAR parameter")
	}

	validNums, err := getParameterValue("USR_NUM")
	if err != nil {
		return errors.New("failed to get USR_NUM parameter")
	}

	// Validasi username
	if !isValidUsername(data.Username, minLength, validChars, validNums) {
		return errors.New("username must be at least " + minLengthStr + " characters long and contain at least one letter and one number")
	}

	// Check apakah username sudah ada
	var existingUser models.Users
	if err := connection.DB.Where("username = ?", data.Username).First(&existingUser).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("username already exists")
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create new user entry
	newUser := models.Users{
		Username:           data.Username,
		Email:              data.Email,
		FirstName:          data.FirstName,
		LastName:           data.LastName,
		IsActive:           data.IsActive,
		Password:           hashedPassword,
		LastPasswordChange: time.Now(),
		IsLogin:            0, // IsLogin default nya 0
		RoleId:             1,
		Status:             "L",
	}
	if err := connection.DB.Create(&newUser).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func Login(data models.Login) (token string, tokenTTL int, err error) {
	var user models.Users

	connection.DB.Where("username = ?", data.Username).First(&user)

	if user.Id == 0 {
		return token, tokenTTL, errors.New("user not found")
	}

	maxFailedAttemptsStr, err := getParameterValue("AUTH_ATM")
	if err != nil {
		return token, tokenTTL, errors.New("failed to retrieve AUTH_ATM parameter")
	}

	maxFailedAttempts, err := strconv.Atoi(maxFailedAttemptsStr)
	if err != nil {
		return token, tokenTTL, errors.New("invalid AUTH_ATM value")
	}

	if user.FailedAttempts >= maxFailedAttempts {
		return token, tokenTTL, errors.New("account locked due to too many failed login attempts")
	}

	if user.IsLogin == 1 {
		return token, tokenTTL, errors.New("user is already active")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		user.FailedAttempts++
		connection.DB.Save(&user)
		return token, tokenTTL, errors.New("incorrect password")
	}

	authTtlStr, err := getParameterValue("AUTH_TTL")
	if err != nil {
		return token, tokenTTL, errors.New("failed to retrieve AUTH_TTL parameter")
	}

	tokenTTL, err = strconv.Atoi(authTtlStr)
	if err != nil {
		return token, tokenTTL, errors.New("invalid token AUTH_TTL value")
	}

	token, err = utils.GenerateJWT(user.Id, tokenTTL)
	if err != nil {
		return token, tokenTTL, errors.New("couldn't create token")
	}

	user.IsLogin = 1
	user.FailedAttempts = 0
	user.LastLogin = time.Now()

	if err := connection.DB.Save(&user).Error; err != nil {
		return token, tokenTTL, errors.New("failed to update user status")
	}

	time.AfterFunc(time.Second*time.Duration(tokenTTL), func() {
		connection.DB.Model(&user).Update("is_login", 0)
	})

	return token, tokenTTL, nil
}

func UpdateUser(UserID string, data models.Users) (user models.Users, err error) {

	if err := connection.DB.First(&user, UserID).Error; err != nil {
		return user, errors.New("user not found")
	}

	user.Username = data.Username
	user.Email = data.Email
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.IsActive = data.IsActive

	if err := connection.DB.Save(&user).Error; err != nil {
		return user, errors.New("failed to update user")
	}

	return user, nil
}

func DeleteUser(UserID string) (err error) {

	var user models.Users
	if err := connection.DB.First(&user, UserID).Error; err != nil {
		return errors.New("user not found")
	}

	user.Status = "D"

	if err := connection.DB.Save(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func ChangePassword(UserID string, data models.Register) (user models.Users, err error) {

	if err := connection.DB.First(&user, UserID).Error; err != nil {
		return user, errors.New("user not found")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = hashedPassword
	user.LastPasswordChange = time.Now()
	user.Status = "L"

	if err := connection.DB.Save(&user).Error; err != nil {
		return user, errors.New("failed to update user")
	}

	return user, nil
}

// func CheckAndUpdateUserStatus() error {
// 	var users []models.Users
// 	ninetyDaysAgo := time.Now().AddDate(0, 0, -1)
// 	fiftyDaysAgo := time.Now().AddDate(0, 0, -1)

// 	// Mendapatkan pengguna yang statusnya "L" dan belum login selama 90 hari
// 	if err := connection.DB.Where("status = ? AND last_login <= ?", "L", ninetyDaysAgo).Find(&users).Error; err != nil {
// 		return fmt.Errorf("failed to fetch users who haven't logged in: %w", err)
// 	}

// 	for _, user := range users {
// 		user.Status = "D" // Set status to "D" for blocked
// 		if err := connection.DB.Save(&user).Error; err != nil {
// 			log.Printf("Failed to update user status for user ID %d: %v\n", user.Id, err)
// 		}
// 	}

// 	// Mendapatkan pengguna yang statusnya "L" dan belum mengganti kata sandi selama 50 hari
// 	if err := connection.DB.Where("status = ? AND last_password_change <= ?", "L", fiftyDaysAgo).Find(&users).Error; err != nil {
// 		return fmt.Errorf("failed to fetch users who haven't changed password: %w", err)
// 	}

// 	for _, user := range users {
// 		user.Status = "D" // Set status to "D" for blocked
// 		if err := connection.DB.Save(&user).Error; err != nil {
// 			log.Printf("Failed to update user status for user ID %d: %v\n", user.Id, err)
// 		}
// 	}

// 	return nil
// }

// func StartStatusCheckScheduler() {
// 	ticker := time.NewTicker(24 * time.Hour)
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				if err := CheckAndUpdateUserStatus(); err != nil {
// 					log.Printf("Error updating user status: %v\n", err)
// 				}
// 			}
// 		}
// 	}()
// }

func User(userId string) (result models.Users, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", userId).Preload("Role").First(&user).Error; err != nil {
		return user, err
	}

	// Preload izin-izin (permissions) dari peran (role) pengguna
	if err := connection.DB.Model(&user.Role).Association("Permissions").Find(&user.Role.Permissions); err != nil {
		return user, err
	}

	return user, nil
}

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

func Logout(userId string) {
	var user models.Users

	connection.DB.Model(&user).Where("id = ?", userId).Update("is_login", 0)
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
		Name:        data.Name,
		Description: data.Description,
		CreatedBy:   user.Username,
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
