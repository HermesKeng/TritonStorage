package main

import(
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"auth"
	"mydb"
)

type JsonBody struct{
	Username string
	Token string
	IsSuccess bool
}


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
					_, tokenString, err := auth.CreateToken(u.Username, u.Password, u.Email)
					log.Println("new token: " + tokenString)
					if err > 0{
						log.Println("Internal Server Error")
					}

					body := JsonBody{Token: tokenString, Username: u.Username, IsSuccess: true}
					json.NewEncoder(w).Encode(body)
				}else{
					log.Println("Users has already existed")
					fmt.Fprintf(w, "false")
				}
		}
	}
}
/*	
http.SetCookie(w, &http.Cookie{Name:"token",Value: tokenString, Expires: expireTime,})
*/
func serveUsers(app *rice.Box, c *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqStr, err := ioutil.ReadAll(r.Body)
		log.Println(r.Cookie("token"))
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
		w.Header().Set("content-type", "application/json")
		switch r.Method {
			case "POST":
				isExist, user, err := mydb.GetUser(u.Email, collection)
				
				if err != nil{
					log.Println(err)
				}else if isExist && user.Password == u.Password{
					log.Println("The user is verified by server")
					expireTime, tokenString, err := auth.CreateToken(u.Username, u.Password, u.Email)
					if err > 0{
						log.Println("Internal Server Error")
					}
					http.SetCookie(w, &http.Cookie{
						Name:"token",
						Value: tokenString,
						Expires: expireTime,
					})
					body := JsonBody{Token: tokenString, Username: user.Username, IsSuccess: true}
					json.NewEncoder(w).Encode(body)
				}else {
					log.Println("Account is incorrect or password incorrect")
					
					fmt.Fprintf(w, "false")
				}
		}
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