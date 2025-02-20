package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) (*User, error)
}

type ProductStore interface {
	GetProduct() ([]Product, error)
	GetProductById(id int) (*Product, error)
	CreateProduct(Product) (*Product, error)
	UpdateProduct(Product, id int) (*Product, error)
}

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
}

type RegisterPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginPayload struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
