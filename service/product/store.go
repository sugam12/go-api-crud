package product

import (
	"database/sql"

	types "github.com/sugam12/go-api-crud/payload"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProduct() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, 0)
	for rows.Next() {
		prod, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *prod)
	}
	return products, nil

}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
		&product.CreatedAt,
		&product.Category)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Store) GetProductById(id int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM product WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	prod := new(types.Product)
	prod, err = scanRowsIntoProduct(rows)
	if err != nil {
		return nil, err
	}
	return prod, nil

}
func (s *Store) CreateProduct(prod types.Product) (*types.Product, error) {
	product := new(types.Product)
	_, err := s.db.Exec("INSERT INTO product(name,description,image,price,quantity,category) values (?,?,?,?,?,?)", prod.Name, prod.Description, prod.Image, prod.Price, prod.Quantity, prod.Category)
	if err != nil {
		return nil, err
	}
	//todo get id and append in product
	return product, nil
}
func (s *Store) UpdateProduct(Product, id int) (*types.Product, error) {
	prod, err := s.GetProductById(id)
	if err != nil {
		return nil, err
	}

	// todo update statement
	return prod, nil

}
