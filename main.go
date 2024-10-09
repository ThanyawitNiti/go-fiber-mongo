package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//create id for every record
)

// instance of a session with mongodb
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"

const mongoURI = "mongodb://localhost:27017" + dbName

// const mongoURI = "mongodb+srv://thanyawit:Q5rvmuP6FYHhCRaX@cluster0.zff25.mongodb.net/golang_db?retryWrites=true&w=majority&appName=Cluster0" + dbName

type Employee struct {
	ID     string  `json:"id" bson:"employeeId"`
	Name   string  `json:"name" bson:"name"`
	Salary float64 `json:"salary" bson:"salary"`
	Age    float64 `json:"age" bson:"age"`
}

func Connect() error { // connect golang to db
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error { //c for accessing res and req

	})
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
