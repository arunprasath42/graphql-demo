package database

import (
	"context"
	"log"
	"time"

	"github.com/arunprasath42/graphql-live/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

func (db *DB) CreateEmployee(employee model.NewEmployee) (*model.Employee, error) {
	collection := db.client.Database("company").Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, employee)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Employee{
		ID:         res.InsertedID.(primitive.ObjectID).Hex(),
		Name:       employee.Name,
		IsTeamLead: employee.IsTeamLead,
	}, nil
}

//FindByID
func (db *DB) GetEmployee(id string) (*model.Employee, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("company").Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	employee := model.Employee{}
	res.Decode(&employee)
	return &employee, nil
}

//updateEmployee
func (db *DB) UpdateEmployee(id string, employee model.NewEmployee) (*model.Employee, error) {
	collection := db.client.Database("company").Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	_, err = collection.UpdateByID(ctx, ObjectID, bson.M{"$set": employee})
	if err != nil {
		log.Fatal(err)
	}

	return &model.Employee{
		ID:         id,
		Name:       employee.Name,
		IsTeamLead: employee.IsTeamLead,
	}, nil
}

//deleteEmployee
func (db *DB) DeleteEmployee(id string) (*model.Employee, error) {
	collection := db.client.Database("company").Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	res := collection.FindOneAndDelete(ctx, bson.M{"_id": ObjectID})
	employee := model.Employee{}
	res.Decode(&employee)
	return &employee, nil
}

//getEmployees
func (db *DB) GetEmployees() ([]*model.Employee, error) {
	collection := db.client.Database("company").Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var employees []*model.Employee
	if err = cursor.All(ctx, &employees); err != nil {
		log.Fatal(err)
	}
	return employees, nil
}
