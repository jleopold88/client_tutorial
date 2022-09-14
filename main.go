package main

import (
	"client_tutorial/client"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":5435", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	user := client.NewUser(conn)
	res, _ := user.CreateUser("Jess", "hehehe@yahoo.com", 24)
	fmt.Println(res)
	res, _ = user.UpdateUser("James", "hihihi@yahoo.com", 10, 1)
	fmt.Println(res)
	res, _ = user.DeleteUser(4)
	fmt.Println(res)
	user_data, _ := user.GetByID(5)
	users_data, _ := user.GetByName("es")
	fmt.Println(user_data)
	fmt.Println(users_data)
}
