package code

import (
	"fmt"
	"github.com/jinzhu/gorm"
	pb "shippy/user-service/proto/user"
)

type Repository interface {
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	Create(*pb.User) error
	GetByEmailAndPassword(*pb.User) (*pb.User, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var u *pb.User
	u.Uid = id
	if err := repo.Db.Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Create(u *pb.User) error {
	u.Name = "111"
	u.Password = "2222"
	u.Company = "2222"

	if err := repo.Db.Create(u).Error; err != nil {
		fmt.Printf("insert error--->%s,--->%s", err, u)
		return err
	}
	return nil
}

func (repo *UserRepository) GetByEmailAndPassword(u *pb.User) (*pb.User, error) {
	if err := repo.Db.Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
