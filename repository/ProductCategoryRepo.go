package repository

import (
	"bookCRUD/models"
	"github.com/go-pg/pg/v10/orm"
)

type productCategoryRepo struct {
	db orm.DB
}

func NewProductCategoryRepo(db orm.DB) models.ProductCategoryRepo {
	return &productCategoryRepo{db: db}
}

func (p productCategoryRepo) Create(productCategory *models.ProductCategory) error {
	_, err := p.db.Model(productCategory).Insert()
	return err
}

func (p productCategoryRepo) GetByID(id string) (*models.ProductCategory, error) {
	productCategory := &models.ProductCategory{}
	err := p.db.Model(productCategory).Where("id = ?", id).Select()
	return productCategory, err
}

func (p productCategoryRepo) GetAll() ([]*models.ProductCategory, error) {
	productCategories := make([]*models.ProductCategory, 0)
	err := p.db.Model(&productCategories).Select()
	return productCategories, err
}

func (p productCategoryRepo) Update(productCategory *models.ProductCategory) error {
	_, err := p.db.Model(productCategory).Where("id = ?", productCategory.ID).Update()
	return err
}

func (p productCategoryRepo) Delete(id string) error {
	_, err := p.db.Model(&models.ProductCategory{}).Where("id = ?", id).Delete()
	return err
}
