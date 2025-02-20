package user

import (
	"database/sql"
	"fmt"

	types "github.com/sugam12/go-api-crud/payload"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email =  ?", email)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil

}

func (s *Store) CreateUser(user types.User) (*types.User, error) {
	_, err := s.db.Exec("INSERT INTO users(firstName, lastName, email, password) VALUES(?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	validUser := new(types.User)

	validUser, err = s.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	return validUser, nil
}
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.FirstName,
		&user.Id,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
