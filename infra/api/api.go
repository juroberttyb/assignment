package api

import (
	"context"
	"encoding/json"
	"server/infra/api/db"
	"server/utils/log"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetUsers(c *gin.Context) {
	for _, result := range users {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Error("something went wrong", zap.Error(err))
			panic(err)
		}
		log.Info(string(output))
	}
}

func GetProducts(c *gin.Context) {
	for _, result := range products {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Error("something went wrong", zap.Error(err))
			panic(err)
		}
		log.Info(string(output))
	}
}

func CreateGroup(c *gin.Context) {
	group_name := c.Param("group_name")

	coll := db.Client.Database("groups").Collection("version1")
	doc := bson.D{
		{"group_name", group_name},
		{"total_member", 0},
		{"members", []string{}},
		{"start_time", time.Now()},
		{"active", true},
	}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}

	insertion_response := response{
		ID:      result.InsertedID.(primitive.ObjectID).Hex(),
		Message: "success",
	}

	c.IndentedJSON(http.StatusOK, insertion_response)
}

func BuyProduct(c *gin.Context) {
	product_name := c.Param("product")
	user_name := c.Param("user")

	for _, result := range products {
		if result["Name"] == product_name {
			result["Number"] -= 1
			return
		}
	}

}

func ChangeActivity(c *gin.Context) {
	state := c.Param("state")

	for _, result := range products {
		if result["Name"] == product_name {
			result["Number"] -= 1
			return
		}
	}

}

func Init() {
	log.Info("Connecting to database...")
	db.Connect()
	log.Info("Connection to database established.")
}
