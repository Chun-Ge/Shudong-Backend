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

// GetCategoryNameByID ...
func GetCategoryNameByID(categoryID int64) (string, error) {
	category := &entity.Category{
		ID: categoryID,
	}
	_, err := database.Orm.Table("category").Get(category)
	return category.Name, err
}

// GetCategoryIDByName ...
func GetCategoryIDByName(categoryName string) (int64, error) {
	category := &entity.Category{
		Name: categoryName,
	}
	has, err := database.Orm.Table("category").Get(category)
	if !has {
		return -1, err
	}
	return category.ID, err
}
