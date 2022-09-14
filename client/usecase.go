package client

import (
	"client_tutorial/entity"
	"client_tutorial/prototype"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type user struct {
	conn *grpc.ClientConn
}

type User interface {
	GetByID(id int32) (*entity.User, error)
	GetByName(name string) ([]*entity.User, error)
	CreateUser(name, email string, age int32) (string, error)
	UpdateUser(name, email string, age, id int32) (string, error)
	DeleteUser(id int32) (string, error)
}

func NewUser(con *grpc.ClientConn) User {
	return user{conn: con}
}

func (u user) CreateUser(name, email string, age int32) (string, error) {
	c := prototype.NewTutorialClient(u.conn)

	req := &prototype.CreateData{
		Name:  name,
		Email: email,
		Age:   age,
	}

	res, err := c.CreateUser(context.Background(), req)
	if err != nil {
		return "Failed Create User", err
	}
	return string(res.Message), err
}
func (u user) UpdateUser(name, email string, age, id int32) (string, error) {
	c := prototype.NewTutorialClient(u.conn)

	req := &prototype.UpdateData{
		Name:  name,
		Email: email,
		Age:   age,
		Id:    id,
	}

	res, err := c.UpdateUser(context.Background(), req)
	if err != nil {
		return "Failed Update User", err
	}
	return string(res.Message), err
}

func (u user) DeleteUser(id int32) (string, error) {
	c := prototype.NewTutorialClient(u.conn)

	req := &prototype.DeleteData{
		Id: id,
	}

	res, err := c.DeleteUser(context.Background(), req)
	if err != nil {
		return "Failed Delete User", err
	}
	return string(res.Message), err
}

func (u user) GetByID(id int32) (*entity.User, error) {
	c := prototype.NewTutorialClient(u.conn)

	req := &prototype.GetDatabyID{
		Id: id,
	}

	res, err := c.GetByID(context.Background(), req)
	fmt.Println(res)
	user := &entity.User{
		Age:     res.User.GetAge(),
		User_id: res.User.GetId(),
		Email:   res.User.Email,
		Name:    res.User.Name,
	}
	return user, err
}

func (u user) GetByName(name string) ([]*entity.User, error) {
	c := prototype.NewTutorialClient(u.conn)

	req := &prototype.GetDatabyName{
		Name: name,
	}

	res, err := c.GetByName(context.Background(), req)
	users := []*entity.User{}

	for _, user := range res.Data {
		temp := &entity.User{
			User_id: user.Id,
			Name:    user.Name,
			Age:     user.Age,
			Email:   user.Email,
		}
		users = append(users, temp)
	}
	return users, err
}
