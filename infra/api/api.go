package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type user struct {
	Name  string
	Vip   int
	Point int
	Money int
}

type product struct {
	Name   string
	Price  int
	Number int
}

var users = []user{
	{Name: "john", Vip: 1, Point: 200, Money: 5000},
	{Name: "bruch", Vip: 2, Point: 300, Money: 6000},
	{Name: "lisa", Vip: 0, Point: 50, Money: 1500},
}

var products = []product{
	{Name: "cat", Price: 1000, Number: 3},
	{Name: "dog", Price: 1500, Number: 5},
	{Name: "horse", Price: 5000, Number: 2},
}

var discount float32 = 0.05
var convert_ratio float32 = 1.

func GetUsers(c *gin.Context) {
	for _, result := range users {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			fmt.Println("something went wrong", zap.Error(err))
			panic(err)
		}
		fmt.Println(string(output))
	}
}

func GetProducts(c *gin.Context) {
	for _, result := range products {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			fmt.Println("something went wrong", zap.Error(err))
			panic(err)
		}
		fmt.Println(string(output))
	}
}

func BuyProduct(c *gin.Context) {
	product_name := c.Param("product")
	user_name := c.Param("user")
	rcv_point := c.Param("point")
	point, _ := strconv.Atoi(rcv_point)

	for i := 0; i < len(products); i++ {
		if products[i].Name == product_name {
			for j := 0; j < len(users); j++ {
				if users[j].Name == user_name {
					if point > users[j].Point {
						fmt.Println("[RESPONSE] not enough point")
						return
					}

					price := products[i].Price
					vip_price := float32(price) * (1. - discount*float32(users[j].Vip))
					estimated_payment := int(vip_price - convert_ratio*float32(point))
					if users[j].Money >= estimated_payment {
						users[j].Money = users[j].Money - estimated_payment
						users[j].Point = users[j].Point - point
						products[i].Number -= 1
						fmt.Println("[RESPONSE] buy action succeful")
					} else {
						fmt.Println("[RESPONSE] so sad, not enough money")
						return
					}
				}
			}
			return
		}
	}
	fmt.Println("[RESPONSE] No such thing")
}

func ChangeActivity(c *gin.Context) {
	state := c.Param("state")
	if state == "normal" {
		discount = 0.05
		convert_ratio = 1.
	} else if state == "festival" {
		discount = 0.06
		convert_ratio = 1.2
	} else if state == "big_festival" {
		discount = 0.08
		convert_ratio = 2.
	} else {
		fmt.Println("[RESPONSE] no such state")
	}
	fmt.Println("[RESPONSE]", state, "activated")
}
