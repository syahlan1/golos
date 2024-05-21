package controllers

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
	"golang.org/x/crypto/bcrypt"
)

// var SecretKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Check apakah username sudah ada
	var existingUser models.Users
	if err := connection.DB.Where("username = ?", data["username"]).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create new user entry
	newUser := models.Users{
		Username: data["username"],
		Email:    data["email"],
		Password: hashedPassword,
		IsLogin:  0, // IsLogin default nya 0
		RoleId:   1,
	}
	if err := connection.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Return success response
	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.Users

	connection.DB.Where("username = ?", data["username"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if user.IsLogin == 1 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User is already active",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	var param models.MasterParameter
	connection.DB.Where("param_key = ?", "TOKEN_TTL").First(&param)

	tokenTTL, err := strconv.Atoi(param.ParamValue)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid token TTL value",
		})
	}

	token, err := utils.GenerateJWT(user.Id, tokenTTL)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could't create token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Second * time.Duration(tokenTTL)),
		HTTPOnly: true,
		Secure:   true,
	}

	c.Cookie(&cookie)

	if err := connection.DB.Model(&user).Update("is_login", 1).Error; err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Failed to update user status"})
	}

	time.AfterFunc(time.Second*time.Duration(tokenTTL), func() {
		connection.DB.Model(&user).Update("is_login", 0)
	})

	return c.JSON(fiber.Map{
		"message": "Login Succcessfully",
		"token":   token,
	})
}

// show username
func User(c *fiber.Ctx) error {
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	// Mendapatkan data pengguna (user) dari database
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).Preload("Role").First(&user).Error; err != nil {
		return err
	}

	// Preload izin-izin (permissions) dari peran (role) pengguna
	if err := connection.DB.Model(&user.Role).Association("Permissions").Find(&user.Role.Permissions); err != nil {
		return err
	}

	return c.JSON(user)
}

func ShowRole(c *fiber.Ctx) error {
	var roles []models.Roles
	if err := connection.DB.Find(&roles).Error; err != nil {
		return err
	}

	return c.JSON(roles)
}

func ShowAllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission
	if err := connection.DB.Find(&permissions).Error; err != nil {
		return err
	}
	return c.JSON(permissions)
}

func ShowPermissions(c *fiber.Ctx) error {
	roleID := c.Params("id")
	var role models.Roles
	if err := connection.DB.Preload("Permissions").Where("id = ?", roleID).First(&role).Error; err != nil {
		return err
	}

	return c.JSON(role)
}

// logout
func Logout(c *fiber.Ctx) error {
	userId := getUserIdFromToken(c)

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	var user models.Users

	connection.DB.Model(&user).Where("id = ?", userId).Update("is_login", 0)

	return c.JSON(fiber.Map{
		"message": "Logged out Successfully!",
	})
}

// Fungsi untuk mendapatkan ID pengguna dari token JWT
func getUserIdFromToken(c *fiber.Ctx) uint {
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return 0
	}

	// Mengonversi issuer token menjadi tipe data uint
	userId, _ := strconv.ParseUint(claims, 10, 64)

	return uint(userId)
}

func CreateRole(c *fiber.Ctx) error {
	// Parse request body
	var req struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Permissions []string `json:"permissions"`
	}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	// Cek apakah role dengan nama yang sama sudah ada
	var existingRole models.Roles
	if err := connection.DB.Where("name = ?", req.Name).First(&existingRole).Error; err == nil {
		// Role sudah ada, kirim respons konflik
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Role already exists"})
	}

	// Buat role baru
	newRole := models.Roles{Name: req.Name, Description: req.Description, CreatedBy: user.Username}

	// Simpan role baru ke dalam database
	if err := connection.DB.Create(&newRole).Error; err != nil {
		return err
	}

	// Dapatkan ID permissions berdasarkan nama permissions yang diberikan
	var permissions []models.Permission
	if err := connection.DB.Where("name IN ?", req.Permissions).Find(&permissions).Error; err != nil {
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

	return c.JSON(fiber.Map{"message": "Role created successfully"})
}

func DeleteRole(c *fiber.Ctx) error {
	roleID := c.Params("id")

	// Mulai transaksi
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

	// Jika sukses, kirim respons sukses
	return c.JSON(fiber.Map{"message": "Role deleted successfully"})
}

func UpdateRole(c *fiber.Ctx) error {
	// Parse request body
	var req struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Permissions []string `json:"permissions"`
	}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("status unauthorized")
	}
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	// Get role ID from URL parameter
	roleID := c.Params("id")

	// Find existing role by ID
	var existingRole models.Roles
	if err := connection.DB.Where("id = ?", roleID).First(&existingRole).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Role not found"})
	}

	// Update role name
	existingRole.Name = req.Name
	existingRole.Description = req.Description
	existingRole.UpdatedBy = user.Username

	// Save the updated role to database
	if err := connection.DB.Save(&existingRole).Error; err != nil {
		return err
	}

	// Get IDs of permissions from the input
	var permissions []models.Permission
	if err := connection.DB.Where("name IN ?", req.Permissions).Find(&permissions).Error; err != nil {
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
		if err != nil {
			// RolePermission doesn't exist, create a new one
			rolePermission := models.RolePermission{RolesId: existingRole.Id, PermissionId: permission.Id}
			if err := connection.DB.Create(&rolePermission).Error; err != nil {
				return err
			}
		}
	}

	// Delete existing RolePermission entries not present in req.Permissions
	if err := connection.DB.Where("roles_id = ? AND permission_id NOT IN ?", existingRole.Id, permissionIDs).Delete(&models.RolePermission{}).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Role updated successfully"})
}

// Helper function to get string value from map
func getStringValue(data map[string]interface{}, key string) string {
	value, ok := data[key].(string)
	if !ok {
		return "" // return empty string if not found or not a string
	}
	return value
}

// Helper function to get int value from map
func getIntValue(data map[string]interface{}, key string) int {
	value, ok := data[key].(float64)
	if !ok {
		return 0 // return 0 if not found or not a number
	}
	return int(value)
}
