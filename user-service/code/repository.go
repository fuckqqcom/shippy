package code

import (
	"fmt"

	pb "shippy/user-service/proto/user"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	Create(*pb.User) error
	GetByEmailAndPassword(*pb.User) (*pb.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var u *pb.User
	u.Uid = id
	if err := repo.DB.Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Create(u *pb.User) error {
	if err := repo.DB.Create(u).Error; err != nil {
		fmt.Printf("insert error--->%s,--->%s", err, u)
		return err
	}
	return nil
}

// bug: 修改u 导致原始密码字符串被修改为数据表中的hash字串，进而影响func CompareHashAndPassword(hashedPassword, password []byte)  判断。
// func (repo *UserRepository) GetByEmailAndPassword(u *pb.User) (*pb.User, error) {
// 	if err := repo.DB.Find(&u).Where("name=?", u.Name).Where("email=?", u.Email).Error; err != nil {
// 		return nil, err
// 	}
// 	return u, nil
// }

// 正确获取数据中的记录，一定要new 新的对象，如果使用传入的对象将导致bug
func (repo *UserRepository) GetByEmailAndPassword(u *pb.User) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.DB.Find(&user).Where("name=?", u.Name).Where("email=?", u.Email).Error; err != nil {
		return nil, err
	}
	return user, nil
}
