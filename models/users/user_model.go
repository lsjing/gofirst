package user_model

import (
	"fmt"
	"github.com/lsjing/gofirst/entitys"
	"github.com/lsjing/gofirst/models"
)

func UserAdd(user *entitys.User) bool {
	orm:= models.GetMaster()
	_,err:=orm.Insert(user)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func UserOne() (*entitys.User, bool) {
	orm:= models.GetSlave()
	user := new(entitys.User)
	has, err := orm.Where("name=?", "laodeng").Get(user)
	if err!=nil {
		fmt.Println(err)
	}
	return user,has
}

func UserOneById(id string) (*entitys.User, bool, error) {
	orm := models.GetSlave()
	user := new(entitys.User)
	has, err := orm.Where("user_id=?", id).Get(user)
	if err != nil {
		fmt.Println(err)
	}

	return user, has, err
}

func UserOneByName(name string) (*entitys.User, bool, error) {
	orm := models.GetSlave()
	user := new(entitys.User)
	has, err := orm.Where("name=?", name).Get(user)
	if err != nil {
		fmt.Println(err)
	}
	return user, has, err
}

func UserList() ([]entitys.User,error) {
	orm:= models.GetSlave()
	users := make([]entitys.User, 0)
	err := orm.Where("user_id>0").Find(&users)
	return users,err
}


