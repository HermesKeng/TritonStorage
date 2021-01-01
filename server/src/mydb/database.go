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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
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
	log.Println(user.Username)
	return true, user, nil
}

func AddFile(filename string, username string, c *mongo.Collection) bool {
	file := File{Filename: filename,Username: username} 
	var result bson.M
	err := c.FindOne(context.TODO(), bson.D{{"username", username},{"filename", filename}}).Decode(&result)
	if err != nil{
		if err == mongo.ErrNoDocuments {
			c.InsertOne(context.TODO(), file)
		}else{
			return false
		}
	}
	return true
}

func GetAllFilesByUsername(username string, c *mongo.Collection) (bool ,[]File) {
	cursor, err := c.Find(context.TODO(), bson.D{{"username", username}})
	if err != nil {
		log.Println(username+" Get All File Error!")
		return false, nil;
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err!=nil {
		log.Fatal(err)
	}
	files := make([]File, 0)
	for _, result := range results {
		id := result["_id"].(primitive.ObjectID).Hex()
		filename := result["filename"].(string)
		filetype := strings.Split(filename, ".")[1]
		newFile := File{Id: id, Filename:filename, Type:filetype, Username:result["username"].(string)}
		files = append(files, newFile)
	}

	return true, files
}
