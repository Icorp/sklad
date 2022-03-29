package repository

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/sklad/models"
)

type providerRepo struct {
	db orm.DB
}

func NewProviderRepo(db orm.DB) models.ProviderRepo {
	return &providerRepo{db: db}
}

func (p providerRepo) Create(provider *models.Provider) error {
	_, err := p.db.Model(provider).Insert()
	return err
}

func (p providerRepo) GetByID(id string) (*models.Provider, error) {
	provider := &models.Provider{}
	err := p.db.Model(provider).Where("id = ?", id).Select()
	return provider, err
}

func (p providerRepo) GetAll() ([]*models.Provider, error) {
	providers := make([]*models.Provider, 0)
	err := p.db.Model(&providers).Select()
	return providers, err
}

func (p providerRepo) Update(provider *models.Provider) error {
	_, err := p.db.Model(provider).Where("id = ?", provider.ID).Update()
	return err
}

func (p providerRepo) Delete(id string) error {
	_, err := p.db.Model(&models.Provider{}).Where("id = ?", id).Delete()
	return err
}
