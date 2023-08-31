package repository

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/models"
)

type CategoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{}
}

func (c CategoryRepository) Create(name string) (*models.Category, error) {
	category := models.Category{Name: name}
	err := common.DB.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) Update(category models.Category, name string) (*models.Category, error) {
	err := common.DB.Model(&category).Update("name", name).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) SelectById(id int) (*models.Category, error) {
	var category = models.Category{}
	err := common.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) DeleteById(id int) error {
	err := common.DB.Delete(&models.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
