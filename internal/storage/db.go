package storage

import (
	"barbershop/creativo/pkg/types"
	"database/sql"
)

func OpenDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/creativo?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func UserInserter(db *sql.DB) types.CreateUserFn {
	return func(u types.User) (types.User, error) {
		query := `
			INSERT INTO users (first_name, last_name, email, phone, status, roles, password, created_at, created_by, updated_at, updated_by)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
			RETURNING id
		`

		var id types.UserID

		err := db.QueryRow(
			query,
			u.Name.FirstName,
			u.Name.LastName,
			u.Contact.Email,
			u.Contact.Phone,
			u.Status,
			u.Roles,
			u.Password,
			u.Audit.CreatedAt,
			u.Audit.CreatedBy,
			u.Audit.UpdatedAt,
			u.Audit.UpdatedBy,
		).Scan(&id)

		if err != nil {
			return types.User{}, err
		}

		newUser := types.User{
			ID:       id,
			Name:     u.Name,
			Contact:  u.Contact,
			Status:   u.Status,
			Roles:    u.Roles,
			Password: u.Password,
			Audit:    u.Audit,
		}

		return newUser, nil
	}
}

func scanUserRow(row *sql.Row) (*types.User, error) {
	var u types.User

	err := row.Scan(
		&u.ID,
		&u.Name.FirstName,
		&u.Name.LastName,
		&u.Contact.Email,
		&u.Contact.Phone,
		&u.Status,
		&u.Roles,
		&u.Password,
		&u.Audit.CreatedAt,
		&u.Audit.CreatedBy,
		&u.Audit.UpdatedAt,
		&u.Audit.UpdatedBy,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func UserByIDGetter(db *sql.DB) types.GetUserByIDFn {
	return func(id types.UserID) (*types.User, error) {
		query := `
			SELECT id, first_name, last_name, email, phone, status, roles, password, created_at, created_by, updated_at, updated_by
			FROM users
			WHERE id = $1
		`

		return scanUserRow(db.QueryRow(query, id))
	}
}

func UserByEmailGetter(db *sql.DB) types.GetUserByEmailFn {
	return func(email types.EmailAddress) (*types.User, error) {
		query := `
			SELECT id, first_name, last_name, email, phone, status, roles, password, created_at, created_by, updated_at, updated_by
			FROM users
			WHERE email = $1
		`

		return scanUserRow(db.QueryRow(query, email))
	}
}

func MakeUserByEmailOrPhoneExists(db *sql.DB) types.CheckIfUserExistsByEmailOrPhoneFn {
	return func(email types.EmailAddress, phone types.PhoneNumber) (bool, error) {
		query := `
			SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 OR phone = $2)
		`

		var exists bool

		err := db.QueryRow(query, email, phone).Scan(&exists)
		if err != nil {
			return false, err
		}

		return exists, nil
	}
}
