package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
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
		Db:     db, //connect to db
	}
	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error { //c for accessing res and req

		query := bson.D{{}}

		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var employees []Employee = make([]Employee, 0)

		if err := cursor.All(c.Context(), &employees); err != nil { //convert type of data
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(employees)

	})

	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		employee.ID = ""

		insertionResult, err := collection.InsertOne(c.Context(), employee)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}} //keep inserted data
		createdRecord := collection.FindOne(c.Context(), filter)          //check data

		createdEmployee := &Employee{}
		createdRecord.Decode(createdEmployee) //json data

		return c.Status(201).JSON(createdEmployee)

	})

	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")

		employeeID, err := primitive.ObjectIDFromHex(idParam)

		if err != nil {
			return c.SendStatus(400)
		}

		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: employeeID}}
		update := bson.D{
			{Key: "$set",
				Value: bson.D{
					{key: "name", Value: employee.Name},
					{key: "age", Value: employee.Age},
					{key: "salary", Value: employee.Salary},
				},
			},
		}

		mg.Db.Collection("employee").FindOneAndUpdate(c.Context(), query, update).Err()

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(400)
			}
			return c.SendStatus(500)
		}
		employee.ID = idParam

		return c.SendStatus(200).JSON(employee)
	})

	app.Delete("/employee/:id", func(c *fiber.Ctx) error {

		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))

		if err != nil {
			return c.SendStatus(400)
		}

		query := bson.D{{Key: "_id", Value: employeeID}}

		deleteResult, err := mg.Db.Collection("employee").DeleteOne(c.Context(), query)

		if err != nil {
			return c.SendStatus(500)
		}

		if deleteResult.DeletedCount < 1 { // nothing got delete
			return c.SendStatus(404)
		}

		return c.SendStatus(200).JSON("recoed deleted")
	})

	app.Listen(":3000")
	log.Fatal(app.Listen(":3000"))
}
