package authService

import (
	"errors"
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

func UserPermission(userId string) (result models.UserPermission, err error) {

	if err := connection.DB.
		Model(models.Users{}).
		Where("id = ?", userId).First(&result).Error; err != nil {
		return result, err
	}

	if err := connection.DB.
		Model(models.Roles{}).
		Where("id = ?", result.RoleId).First(&result.Role).Error; err != nil {
		return result, err
	}

	if err := connection.DB.Raw(`
		WITH RECURSIVE ParentHierarchy AS (
			SELECT id, parent_id
			FROM menus
			WHERE id IN(select m.id 
						from role_menus rm 
						join menus m on m.id = rm.menu_id 
						where rm.role_id = ? and
						rm.deleted_at is null and m.deleted_at is null and (rm.read = true or rm.write = true or rm.update = true or rm.delete = true or rm.download = true))
			union all
			SELECT m.id, m.parent_id
			FROM menus m
			JOIN ParentHierarchy ph ON ph.parent_id = m.id)
		SELECT m.id, m.parent_id, m.icon, m.label as title, m.command as path
		FROM ParentHierarchy ph
		JOIN menus m ON ph.id = m.id
		WHERE m.deleted_at IS null
		and m.parent_id is null
		group by m.id
		order by "order"`, result.RoleId).
		Find(&result.Role.Menu).Error; err != nil {
		return result, err
	}

	for i, data := range result.Role.Menu {
		Child, err := findChild(data.Id, int(result.RoleId))
		if err != nil {
			return result, err
		}
		result.Role.Menu[i].Subnav = Child
	}

	// Get Module
	// if err := connection.DB.
	// 	Select("rm.*, master_modules.id AS module_id, master_modules.module_name AS module_name",
	// 		"master_modules.description AS description, master_modules.is_active AS is_active").
	// 	Joins("JOIN role_modules rm ON rm.module_id = master_modules.id").
	// 	Model(models.MasterModule{}).
	// 	Where("rm.roles_id = ? and rm.deleted_at is null", result.RoleId).
	// 	Find(&result.Role.Module).Error; err != nil {
	// 	return result, err
	// }

	// for i, data := range result.Role.Module {

	// 	var table []models.ShowRoleTables

	// 	if err := connection.DB.
	// 		Select("*, master_tables.table_name AS table").
	// 		Joins("JOIN role_tables rt ON rt.table_id = master_tables.id").
	// 		Model(models.MasterTable{}).
	// 		Where("rt.role_modules_id = ? and rt.deleted_at is null", data.Id).
	// 		Find(&table).Error; err != nil {
	// 		return result, err
	// 	}

	// 	result.Role.Module[i].Table = table
	// }

	return
}

func findChild(parentId, roleId int) ([]models.ShowMenu, error) {
	var child []models.ShowMenu
	if err := connection.DB.
		Select(`menus.id, parent_id, "type", icon, label as title, command as path`).
		Joins("LEFT JOIN role_menus rm ON rm.menu_id = menus.id").
		Model(models.Menu{}).
		Where(`rm.role_id = ? OR (menus."type" = ? AND (select count(*) from menus m 
				left join role_menus rm on rm.menu_id = m.id 
				where parent_id = menus.id 
				AND m."deleted_at" IS NULL 
				AND rm."deleted_at" IS NULL  
				AND (rm.read = true or rm.write = true or rm.update = true or rm.delete = true or rm.download = true)) > 0)`, roleId, "P").
		Where("parent_id = ? and rm.deleted_at is null", parentId).
		Order(`"order" asc`).
		Find(&child).Error; err != nil {
		return nil, err
	}


	for i, data := range child {
		subChild, err := findChild(data.Id, roleId)
		if err != nil {
			return nil, err
		}

		child[i].Subnav = subChild

	}

	return child, nil
}

func Logout(userId string) {
	var user models.Users

	connection.DB.Model(&user).Where("id = ?", userId).Update("is_login", 0)
}
