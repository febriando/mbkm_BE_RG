package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

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

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	categories := []entity.Category{}
	if err := r.db.WithContext(ctx).Table("categories").Where("user_id = ?", id).Find(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Category{}, nil
		}
		return nil, err
	}
	return categories, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	if err := r.db.WithContext(ctx).Create(&category).Error; err != nil {
		return 0, err
	}
	return category.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	if err := r.db.WithContext(ctx).Create(&categories).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	category := entity.Category{}
	if err := r.db.WithContext(ctx).Table("categories").Where("id = ?", id).Find(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Category{}, nil
		}
		return entity.Category{}, err
	}
	return category, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	if err := r.db.WithContext(ctx).Updates(&category).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(entity.Category{}).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
