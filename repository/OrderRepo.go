package repository

import (
	"bookCRUD/models"
	"github.com/go-pg/pg/v10/orm"
)

type orderRepo struct {
	db orm.DB
}

func NewOrderRepo(db orm.DB) models.OrderRepo {
	return &orderRepo{db: db}
}

func (o orderRepo) Create(order *models.Order) error {
	_, err := o.db.Model(order).Insert()
	return err
}

func (o orderRepo) GetByID(id string) (*models.Order, error) {
	order := &models.Order{}
	err := o.db.Model(order).Where("id = ?", id).First()
	return order, err
}

func (o orderRepo) GetAll() ([]*models.Order, error) {
	orders := make([]*models.Order, 0)
	err := o.db.Model(&orders).Select()
	return orders, err
}

func (o orderRepo) Update(order *models.Order) error {
	_, err := o.db.Model(order).Where("id = ?", order.ID).Update()
	return err
}

func (o orderRepo) Delete(id string) error {
	order := &models.Order{}
	_, err := o.db.Model(order).Where("id = ?", id).Delete()
	return err
}
