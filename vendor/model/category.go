package model

import (
	"database"
	"entity"
)

// GetCategoryNameByID ...
func GetCategoryNameByID(categoryID int64) (string, error) {
	category := &entity.Category{
		ID: categoryID,
	}
	_, err := database.Orm.Table("category").Get(category)
	return category.Name, err
}
