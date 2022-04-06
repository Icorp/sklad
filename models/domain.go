// models/book.go

package models

import (
	"time"
)

type Client struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	tableName struct{}  `pg:"clients"` //nolint
}

type Employee struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	tableName struct{}  `pg:"employees"` //nolint
}

type Order struct {
	ID              string           `json:"id"`
	ClientID        string           `json:"client_id"`
	ProductID       string           `json:"product_id"`
	EmployeeID      string           `json:"employee_id"`
	Count           int              `json:"count"`
	Address         string           `json:"address"`
	CurrentPrice    int              `json:"current_price"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	Order           *Order           `pg:"rel:has-one"`
	Provider        *Provider        `pg:"rel:has-one"`
	Product         *Product         `pg:"rel:has-one"`
	ProductCategory *ProductCategory `pg:"rel:has-one"`
	Employee        *Employee        `pg:"rel:has-one"`
	tableName       struct{}         `pg:"orders"` //nolint
}

type ProductCategory struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	tableName struct{} `pg:"product_category"` //nolint
}

type Product struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Price             int       `json:"price"`
	Count             int       `json:"count"`
	ProductCategoryID string    `json:"product_category_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	tableName         struct{}  `pg:"products"` //nolint
}

type Provider struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	tableName struct{}  `pg:"providers"` //nolint
}

type ClientRepo interface {
	Create(client *Client) error
	GetByID(id string) (*Client, error)
	GetByEmail(email string) (*Client, error)
	GetAll() ([]*Client, error)
	Update(client *Client) error
	Delete(id string) error
}

type EmployeeRepo interface {
	Create(employee *Employee) error
	GetByID(id string) (*Employee, error)
	GetAll() ([]*Employee, error)
	Update(employee *Employee) error
	Delete(id string) error
}

type OrderRequest struct {
	From *time.Time
	To   *time.Time
}

type OrderRepo interface {
	Create(order *Order) error
	GetByID(id string) (*Order, error)
	GetAll(request *OrderRequest) ([]*Order, error)
	List() ([]*Order, error)
	Update(order *Order) error
	Delete(id string) error
}

type ProductCategoryRepo interface {
	Create(productCategory *ProductCategory) error
	GetByID(id string) (*ProductCategory, error)
	GetAll() ([]*ProductCategory, error)
	Update(productCategory *ProductCategory) error
	Delete(id string) error
}

type ProductRepo interface {
	Create(product *Product) error
	GetByID(id string) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(id string) error
}

type ProviderRepo interface {
	Create(provider *Provider) error
	GetByID(id string) (*Provider, error)
	GetAll() ([]*Provider, error)
	Update(provider *Provider) error
	Delete(id string) error
}
