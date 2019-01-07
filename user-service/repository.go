package main

import (
	"errors"
	"github.com/xormplus/xorm"
	pb "shippy/user-service/proto/user"
)

type Repository interface {
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *xorm.Engine
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User

	if err := repo.db.Find(&users).Error; err != nil {
		return nil, errors.New(err())
	}

	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id

	if err := repo.db.Find(&user).Error; err != nil {
		return nil, errors.New(err())
	}

	return user, nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.Find(&user).Error; err != nil {
		return nil, errors.New(err())
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if _, err := repo.db.Insert(user); err != nil {
		return err
	}

	return nil
}
