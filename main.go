package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/madecanggih/protobuf/model"
)

func main() {
	var myUser = &model.User{
		Id:       "u001",
		Name:     "John Doe",
		Password: "S3cur3",
		Gender:   model.UserGender_FEMALE,
	}

	// var userList = &model.UserList{
	// 	List: []*model.User{
	// 		myUser,
	// 	},
	// }

	var myGarage = &model.Garage{
		Id:   "g001",
		Name: "Shop Here",
		Coordinate: &model.GarageCoordinate{
			Latitude:  149.42,
			Longitude: 52.12,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			myGarage,
		},
	}

	// var garageListByUser = &model.GarageListByUser{
	// 	List: map[string]*model.GarageList{
	// 		myUser.Id: garageList,
	// 	},
	// }

	fmt.Printf("Original \n %#v \n", myUser)
	fmt.Printf("String \n %v \n", myUser.String())

	var buf1 bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf1, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	jsonResult := buf1.String()
	fmt.Printf("JSON \n %v \n", jsonResult)

	protoResult := new(model.GarageList)

	err2 := jsonpb.UnmarshalString(jsonResult, protoResult)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("Proto String \n %v \n", protoResult.String())
}
