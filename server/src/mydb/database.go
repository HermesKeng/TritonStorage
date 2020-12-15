package mydb

import(
	"log"
	"time"
	"context"
	"errors"
	"fmt"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

const connectionString = "mongodb+srv://mainUser:asdfghjkl@cluster0.7braz.mongodb.net/<dbname>?retryWrites=true&w=majority"
func NewDatabaseClient() (context.Context, func(), *mongo.Client) {
	ctx, cancel  := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	return ctx, cancel, client
}
func AddNewUser(u User, c *mongo.Collection) error {
	insertResult, err := c.InsertOne(context.TODO(), u)
	if err != nil{
		log.Println(err)
		return errors.New("insertion wrong")
	}
	fmt.Println("InsertID:", insertResult.InsertedID)
	return nil
}

func GetUser(email string, c *mongo.Collection) (bool, User, error) {
	var user User
	err := c.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil{
		return false, User{},nil
	}
	log.Println(user.Email)
	log.Println(user.Password)
	return true, user, nil
}