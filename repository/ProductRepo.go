package repository

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/sklad/models"
)

type productRepo struct {
	db orm.DB
}

func NewProductRepo(db orm.DB) models.ProductRepo {
	return &productRepo{db: db}
}

func (p productRepo) Create(product *models.Product) error {
	_, err := p.db.Model(product).Insert()
	return err
}

func (p productRepo) GetByID(id string) (*models.Product, error) {
	product := &models.Product{}
	err := p.db.Model(product).Where("id = ?", id).Select()
	return product, err
}

func (p productRepo) GetAll() ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	err := p.db.Model(&products).Select()
	return products, err
}

func (p productRepo) Update(product *models.Product) error {
	_, err := p.db.Model(product).Where("id = ?", product.ID).UpdateNotZero()
	return err
}

func (p productRepo) Delete(id string) error {
	_, err := p.db.Model(&models.Product{}).Where("id = ?", id).Delete()
	return err
}
