package main

import (
	"context"
	"log"
	// "net"
	"fmt"
	"encoding/json"

	"github.com/mragungsetiaji/learn-go/grpc/common/config"
	"github.com/mragungsetiaji/learn-go/grpc/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}
func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}
func main() {
	user1 := model.User{
		Id:       "n001",
		Name:     "Tuk tuk",
		Password: "tuktuk123",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	user2 := model.User{
		Id:       "n002",
		Name:     "Simba",
		Password: "simba321",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
  
	  garage1 := model.Garage{
		  Id:   "q001",
		  Name: "tuktuk tempe",
		  Coordinate: &model.GarageCoordinate{
			  Latitude:  45.123123123,
			  Longitude: 54.1231313123,
		  },
	  }

	  garage2 := model.Garage{
		Id:   "x02",
		Name: "Simba Ubi",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}
	
	  
	user := serviceUser()

	fmt.Println("\n", "===========> user test")

	// register user1
	user.Register(context.Background(), &user1)

	// register user2
	user.Register(context.Background(), &user2)
	// ...

	// show all registered users
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	garage := serviceGarage()

	fmt.Println("\n", "===========> garage test A")

	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	// add garage2 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage2,
	})

	// show all garages of user1
	res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))
  }