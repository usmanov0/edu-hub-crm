package repository

import (
	"context"
	"database/sql"
	"edu-sphere-crm/entity"
	"edu-sphere-crm/pkg/customErrors"
	"errors"
	"github.com/jackc/pgx"
	"time"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	GetAllUsers(ctx context.Context, offset, limit int) ([]*entity.User, error)
	GetAllByCourse(ctx context.Context, courseId int) ([]entity.User, error)
	GetPasswordByEmail(ctx context.Context, email string) (bool, string, error)
	FindById(ctx context.Context, userId int) (*entity.User, error)
	FindOneByEmail(ctx context.Context, email string) (*entity.User, error)
	ExistsByMail(ctx context.Context, email string) (bool, error)
	Search(ctx context.Context, criteria string) ([]entity.User, error)
	UpdateRole(ctx context.Context, id int, user *entity.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Save(ctx context.Context, user *entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
	INSERT INTO users (first_name, last_name, role, email, password, phone_number, address, registration_date, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	ON CONFLICT (id) DO UPDATE
	SET first_name = excluded.first_name, 
	    last_name = excluded.last_name, 
	    role = excluded.role, 
	    email = excluded.email, 
	    password = excluded.password, 
	    phone_number = excluded.phone_number, 
	    address = excluded.address, 
	    updated_at = excluded.updated_at
	RETURNING id`

	err := u.db.QueryRow(query, user.FirstName, user.LastName, user.Role, user.Email, user.Password, user.PhoneNumber, user.Address, user.RegistrationDate, user.UpdatedAt).
		Scan(&user.Id)
	if err != nil {
		return customErrors.ErrIdScanFailed
	}
	return nil
}

func (u *userRepository) GetAllUsers(ctx context.Context, offset, limit int) ([]*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		SELECT u.first_name, u.last_name, u.role, u.email, u.password, u.phone_number, u.address, u.registration_date, u.updated_at 
        FROM users u
        ORDER BY u.registration_date 
        OFFSET $1 
        LIMIT $2`

	rows, err := u.db.Query(query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.FirstName,
			&user.LastName,
			&user.Role,
			&user.Email,
			&user.Password,
			&user.PhoneNumber,
			&user.Address,
			&user.RegistrationDate,
			&user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetAllByCourse(ctx context.Context, courseId int) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
			SELECT u.first_name, 
		       u.last_name, 
		       u.role, 
		       u.email, 
		       u.password, 
		       u.phone_number, 
		       u.address, 
		       u.registration_date,
		       u.updated_at 
			 FROM users u 
			INNER JOIN courses c ON u.id = c.user_id 
			WHERE c.course_id = $1`

	rows, err := u.db.Query(query, courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Role, &user.Email, &user.Password,
			&user.PhoneNumber, &user.Address, &user.RegistrationDate, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) GetPasswordByEmail(ctx context.Context, email string) (exists bool, hashedPassword string, err error) {
	query := "SELECT password FROM users WHERE email = ?"
	row := u.db.QueryRow(query, email)

	err = row.Scan(&hashedPassword)

	if errors.Is(err, sql.ErrNoRows) {
		return false, "", nil
	} else if err != nil {
		// Other error occurred
		return false, "", err
	}
	return true, hashedPassword, nil
}

func (u *userRepository) FindById(ctx context.Context, userId int) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
			SELECT u.first_name, 
		       u.last_name, 
		       u.role, 
		       u.email, 
		       u.password, 
		       u.phone_number, 
		       u.address, 
		       u.registration_date,
		       u.updated_at 
			FROM users u 
			WHERE u.id = $1`

	var user entity.User
	err := u.db.QueryRow(query, userId).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Address,
		&user.RegistrationDate,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, customErrors.ErrScanRows
	}
	return &user, nil

}

func (u *userRepository) FindOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
			SELECT u.first_name, 
		       u.last_name, 
		       u.role, 
		       u.email, 
		       u.password, 
		       u.phone_number, 
		       u.address, 
		       u.registration_date,
		       u.updated_at 
			FROM users u 
			WHERE u.email = $1`

	var user entity.User
	err := u.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Address,
		&user.RegistrationDate,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, customErrors.ErrScanRows
	}
	return &user, nil
}

func (u *userRepository) ExistsByMail(ctx context.Context, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
			SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	var exists bool
	err := u.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, customErrors.ErrIdScanFailed
	}
	return exists, nil
}

func (u *userRepository) Search(ctx context.Context, criteria string) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		SELECT u.first_name, 
		       u.last_name, 
		       u.role, 
		       u.email, 
		       u.password, 
		       u.phone_number, 
		       u.address, 
		       u.registration_date,
		       u.updated_at 
		FROM users u 
		WHERE u.first_name ILIKE $1 
		    OR u.last_name ILIKE $1 
		    OR u.email ILIKE $1`

	rows, err := u.db.Query(query, "%%"+criteria+"%%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Role,
			&user.Email,
			&user.Password,
			&user.PhoneNumber,
			&user.Address,
			&user.RegistrationDate,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, customErrors.ErrScanRows
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) UpdateRole(ctx context.Context, id int, user *entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `UPDATE users 
			  SET role = $1 
			  WHERE id = $2`

	_, err := u.db.Exec(query, user.Role, id)
	if err != nil {
		return customErrors.ErrFailedExecuteQuery
	}
	return nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
			DELETE FROM users 
			WHERE id = $1`

	_, err := u.db.Exec(query, id)
	if err != nil {
		return customErrors.ErrFailedDeleteAccount
	}
	return nil
}
