package main

import(
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	rice "github.com/GeertJohan/go.rice"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"mydb"
)



func main(){
	appBox, err := rice.FindBox("../build")
	if err != nil{
		log.Fatal(err)
	}
	log.Println("Connect to MongoDb")
	ctx, cancel, client  := mydb.NewDatabaseClient()
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	/*databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil{
		log.Fatal(err)
	}
	log.Println(databases)*/
	
	http.Handle("/static/", http.FileServer(appBox.HTTPBox()))
	http.HandleFunc("/", serveAppHandler(appBox))
	http.HandleFunc("/users", serveUsers(appBox, client))
	http.HandleFunc("/newuser", registerNewUser(appBox, client))
	log.Println("Server start at port 8080")
	if err:= http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}

func registerNewUser(app *rice.Box, c *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexFile, err := app.Open("index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		reqStr, err := ioutil.ReadAll(r.Body)
		if err != nil{
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		u := mydb.User{}
		log.Println(string(reqStr))
		err = json.Unmarshal(reqStr, &u)

		if err != nil {
			log.Println(err)
		}
		collection := c.Database("tritonstorage").Collection("users")

		switch r.Method {
			case "POST":
				log.Println("register")
				isExist,_, err := mydb.GetUser(u.Email, collection)
				if !isExist {
					err = mydb.AddNewUser(u, collection)
					if err != nil{
						log.Println(err)
					}
				}else{
					log.Println("Users has already existed")
				}
		}
		http.ServeContent(w, r, "index.html", time.Time{}, indexFile)
	}
}
func serveUsers(app *rice.Box, c *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexFile, err := app.Open("index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		reqStr, err := ioutil.ReadAll(r.Body)
		if err != nil{
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		u := mydb.User{}
		log.Println(string(reqStr))
		err = json.Unmarshal(reqStr, &u)

		if err != nil {
			log.Println(err)
		}
		collection := c.Database("tritonstorage").Collection("users")

		switch r.Method {
			case "POST":
				log.Println("check")
				isExist, user, err := mydb.GetUser(u.Email, collection)
				
				if err != nil{
					log.Println(err)
				}else if isExist && user.Password == u.Password{
					log.Println("The user is verified by server")
				}else if !isExist {
					log.Println("data is not exist")
				}
		}
		http.ServeContent(w, r, "index.html", time.Time{}, indexFile)
	}
}


func serveAppHandler(app *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexFile, err := app.Open("index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Println("get request")
		http.ServeContent(w, r, "index.html", time.Time{}, indexFile)
	}
}