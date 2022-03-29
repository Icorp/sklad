package repository

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/sklad/models"
)

type employeeRepo struct {
	db orm.DB
}

func NewEmployeeRepo(db orm.DB) models.EmployeeRepo {
	return &employeeRepo{db: db}
}

func (p employeeRepo) Create(employee *models.Employee) error {
	_, err := p.db.Model(employee).Insert()
	return err
}

func (p employeeRepo) GetByID(id string) (*models.Employee, error) {
	employee := &models.Employee{}
	err := p.db.Model(employee).Where("id = ?", id).Select()
	return employee, err
}

func (p employeeRepo) GetAll() ([]*models.Employee, error) {
	employees := make([]*models.Employee, 0)
	err := p.db.Model(&employees).Select()
	return employees, err
}

func (p employeeRepo) Update(employee *models.Employee) error {
	_, err := p.db.Model(employee).Where("id = ?", employee.ID).UpdateNotZero()
	return err
}

func (p employeeRepo) Delete(id string) error {
	_, err := p.db.Model(&models.Employee{}).Where("id = ?", id).Delete()
	return err
}
