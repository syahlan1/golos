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

	// Get menu
	if err := connection.DB.Select("m.*").
		Joins("JOIN role_menu rm ON rm.menu_id = menus.id").
		Joins("JOIN menus m ON m.id = menus.parent_id").
		Model(models.Menu{}).
		Where("rm.role_id = ?", result.RoleId).
		Group("m.id").
		Order(`"order" asc`).
		Find(&result.Role.Menu).Error; err != nil {
		return result, err
	}

	for i, data := range result.Role.Menu {

		var child []models.ShowMenuPermission

		if err := connection.DB.Select("*").
			Joins("JOIN role_menu rm ON rm.menu_id = menus.id").
			Model(models.Menu{}).
			Where("parent_id = ? AND rm.role_id = ?", data.Id, result.RoleId).
			Order(`"order" asc`).
			Find(&child).Error; err != nil {
			return result, err
		}

		result.Role.Menu[i].Child = child
	}

	return
}

func Logout(userId string) {
	var user models.Users

	connection.DB.Model(&user).Where("id = ?", userId).Update("is_login", 0)
}
