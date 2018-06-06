package model

import (
	"database"
	"entity"
)

// GetAllCategoryNames ...
func GetAllCategoryNames() ([]string, error) {
	categories := make(entity.Categories, 0)
	err := database.Orm.Table("category").Find(&categories)

	retNames := make([]string, len(categories))
	for idx, category := range categories {
		retNames[idx] = category.Name
	}
	return retNames, err
}
