package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

// * implemented
func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	var categories []entity.Category
	query, err := r.db.Model(&entity.Category{}).Select("*").Where("user_id = ?", id).Rows()
	defer query.Close()
	for query.Next() {
		err = r.db.ScanRows(query, &categories)
	}
	if err != nil {
		return []entity.Category{}, err
	}
	return categories, nil
}

// *implemented
func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	res := r.db.Create(&category)
	if res.Error != nil {
		return 0, res.Error
	}
	return category.ID, nil
}

// ?need to check later for the return value
func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	res := r.db.Create(&categories)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// *implemented
func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	var category entity.Category
	query := r.db.Model(&entity.Category{}).First(&category, id)
	if query.Error != nil {
		return entity.Category{}, query.Error
	}
	return category, nil
}

// *implemented
func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	res := r.db.Model(&entity.Category{}).Where("id = ?", category.ID).Updates(&category)
	return res.Error
}

// *implemented
func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	res := r.db.Delete(&entity.Category{}, id)
	return res.Error
}
