package database

import (
	"github.com/galifornia/go-simple-hrms/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewEmployee(c *fiber.Ctx) error {
	coll := DB.Database.Collection("employees")
	employee := new(types.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString((err.Error()))
	}

	result, err := coll.InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(500).SendString((err.Error()))
	}

	// double check that insertion went well ??
	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	record := coll.FindOne(c.Context(), filter)
	createdEmployee := &types.Employee{}
	record.Decode(createdEmployee)

	return c.Status(201).JSON(createdEmployee)
}

func GetEmployees(c *fiber.Ctx) error {
	coll := DB.Database.Collection("employees")
	query := bson.D{{}}
	// make([]Employee, 0)
	var employees []types.Employee

	cursor, err := coll.Find(c.Context(), query)

	if err != nil {
		return c.Status(500).SendString((err.Error()))
	}

	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString((err.Error()))
	}

	return c.JSON(employees)
}

func GetEmployee(c *fiber.Ctx) error {
	coll := DB.Database.Collection("employees")

	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(500).SendString((err.Error()))
	}

	query := bson.D{{Key: "_id", Value: id}}
	result := coll.FindOne(c.Context(), query)
	if err := result.Err(); err != nil {
		return c.Status(500).SendString((err.Error()))
	}

	var employee types.Employee
	if err := result.Decode(&employee); err != nil {
		return c.Status(500).SendString((err.Error()))
	}

	return c.JSON(employee)
}

func DeleteEmployee(c *fiber.Ctx) error {
	coll := DB.Database.Collection("employees")
	idParam := c.Params("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString((err.Error()))
	}

	filter := bson.D{{Key: "_id", Value: id}}
	result, err := coll.DeleteOne(c.Context(), &filter)
	if err != nil {
		return c.Status(400).SendString((err.Error()))
	}

	if result.DeletedCount < 1 {
		return c.Status(404).SendString((err.Error()))
	}

	return c.Status(200).JSON("Employee record was deleted!")
}

func UpdateEmployee(c *fiber.Ctx) error {
	coll := DB.Database.Collection("employees")

	idParam := c.Params("id")
	employee := new(types.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString((err.Error()))
	}

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString((err.Error()))
	}

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: employee.Name},
		{Key: "manager", Value: employee.Manager},
		{Key: "position", Value: employee.Position},
		{Key: "salary", Value: employee.Salary},
		{Key: "age", Value: employee.Age},
	}}}

	err = coll.FindOneAndUpdate(c.Context(), filter, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(400).SendString(err.Error())
		}
		return c.Status(500).SendString((err.Error()))
	}

	return c.Status(200).JSON(employee)
}
