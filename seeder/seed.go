package seeder

import (
	"os"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"golang.org/x/crypto/bcrypt"
)

func UserSeed() {
	var adminPassword = os.Getenv("ADMIN_PASSWORD")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	var users = []models.Users{
		{
			Username: os.Getenv("ADMIN_USERNAME"),
			Email:    os.Getenv("ADMIN_EMAIL"),
			Password: hashedPassword,
			IsLogin:  0,
			RoleId:   1,
		},
	}

	var roles = []models.Roles{
		{
			Name:        "Superadmin",
			Description: "High Tier Admin",
		},
	}

	var permission = []models.Permission{
		{
			Name: "create",
		},
		{
			Name: "update",
		},
		{
			Name: "delete",
		},
		{
			Name: "approve",
		},
		{
			Name: "create_role",
		},
		{
			Name: "delete_role",
		},
	}

	var rolePermissions = []models.RolePermission{
		{
			RolesId:      1,
			PermissionId: 1,
		},
		{
			RolesId:      1,
			PermissionId: 2,
		},
		{
			RolesId:      1,
			PermissionId: 3,
		},
		{
			RolesId:      1,
			PermissionId: 4,
		},
		{
			RolesId:      1,
			PermissionId: 5,
		},
		{
			RolesId:      1,
			PermissionId: 6,
		},
	}

	connection.DB.Save(&permission)
	connection.DB.Save(&roles)
	connection.DB.Save(&users)
	connection.DB.Save(&rolePermissions)
}


func CreditSeed() {
	var credit = []models.CreditType{
		{
			Code: "0213",
			Name: "MIKRO INVESTASI",
		},
		{
			Code: "0216",
			Name: "MIKRO MODAL KERJA",
		},
		{
			Code: "0521",
			Name: "PEMILIKAN KENDARAAN BERMOTOR",
		},
		{
			Code: "0527",
			Name: "MULTIGUNA UMUM",
		},
		{
			Code: "0529",
			Name: "KREDIT TANPA AGUNAN",
		},

	}

	var usage = []models.Usage{
		{
			Code: "1",
			Name: "Modal Kerja",
		},
		{
			Code: "2",
			Name: "Investasi",
		},
		{
			Code: "3",
			Name: "Konsumsi",
		},
	}

	var currency = []models.Currency{
		{
			Code: "IDR",
			Name: "Rupiah",
		},
		{
			Code: "USD",
			Name: "Dollar",
		},
		{
			Code: "EUR",
			Name: "Euro",
		},
	}

	connection.DB.Save(&credit)
	connection.DB.Save(&usage)
	connection.DB.Save(&currency)
}