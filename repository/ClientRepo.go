package repository

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/sklad/models"
)

type clientRepo struct {
	db orm.DB
}

func NewClientRepo(db orm.DB) models.ClientRepo {
	return &clientRepo{db: db}
}

func (c clientRepo) Create(client *models.Client) error {
	_, err := c.db.Model(client).Insert()
	return err
}

func (c clientRepo) GetByID(id string) (*models.Client, error) {
	client := &models.Client{}
	err := c.db.Model(client).Where("id = ?", id).Select()
	return client, err
}

func (c clientRepo) GetByEmail(email string) (*models.Client, error) {
	client := &models.Client{}
	err := c.db.Model(client).Where("email = ?", email).Select()
	return client, err
}

func (c clientRepo) GetAll() ([]*models.Client, error) {
	clients := make([]*models.Client, 0)
	err := c.db.Model(&clients).Select()
	return clients, err
}

func (c clientRepo) Update(client *models.Client) error {
	_, err := c.db.Model(client).Where("id = ?", client.ID).UpdateNotZero()
	return err
}

func (c clientRepo) Delete(id string) error {
	_, err := c.db.Model(&models.Client{}).Where("id = ?", id).Delete()
	return err
}
