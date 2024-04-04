package repository

import (
	"edu-sphere-crm/entity"
	"github.com/jackc/pgx"
)

type UserRepository interface {
	Save(user *entity.User) error
	FindById(userId int) (*entity.User, error)
	FindOneByEmail(email string) (*entity.User, error)
	ExistsByMail(email string) (bool, error)
	Search(criteria string) ([]entity.User, error)
	UpdateUser(userId int) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Save(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindById(userId int) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindOneByEmail(email string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) ExistsByMail(email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Search(criteria string) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) UpdateUser(userId int) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}
