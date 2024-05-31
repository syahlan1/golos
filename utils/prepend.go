package utils

import "github.com/syahlan1/golos/models"

func Prepend(slice []models.Dropdown, value models.Dropdown) []models.Dropdown {
	return append([]models.Dropdown{value}, slice...)
}