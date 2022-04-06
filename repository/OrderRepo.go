package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/sklad/models"
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

func (o orderRepo) GetAll(request *models.OrderRequest) ([]*models.Order, error) {
	orders := make([]*models.Order, 0)
	query := o.db.Model(&orders)

	if request != nil {
		if request.From != nil {
			query.Where("created_at > ?", request.From)
		}

		if request.To != nil {
			query.Where("created_at < ?", request.To)
		}
	}

	if err := query.Select(); err != nil && err != pg.ErrNoRows {
		return nil, err
	}

	return orders, nil
}

func (o orderRepo) List() ([]*models.Order, error) {
	orders := make([]*models.Order, 0)
	err := o.db.Model(&orders).
		Relation("Product").
		Relation("Provider").
		Relation("Employee").
		Relation("ProductCategory").
		Select()

	return orders, err
}

func (o orderRepo) Update(order *models.Order) error {
	_, err := o.db.Model(order).Where("id = ?", order.ID).UpdateNotZero()
	return err
}

func (o orderRepo) Delete(id string) error {
	order := &models.Order{}
	_, err := o.db.Model(order).Where("id = ?", id).Delete()
	return err
}
