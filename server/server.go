package main

import(
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	mongo "go.mongodb.org/mongo-driver/mongo"
	chi "github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"auth"
	"mydb"
	"bufio"
	"os"
	"time"
	"path/filepath"
	"strings"
	"bytes"
)

type JsonBody struct{
	Username string
	Token string
	IsSuccess bool
}


func main(){
	r := chi.NewRouter()
	appBox, err := rice.FindBox("../build")
	if err != nil{
		log.Println("something wrong for rice box")
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	
	log.Println("Connect to MongoDb")
	ctx, cancel, client  := mydb.NewDatabaseClient()
	defer cancel()
	defer client.Disconnect(ctx)
	/*databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil{
		log.Fatal(err)
	}
	log.Println(databases)*/
	r.Get("/", serveAppHandler(appBox,r))
	r.Get("/{userID}/files", getAllFile(client))
	r.Get("/{userID}/files/{id}", downloadFile(client))
	r.Get("/newfile", serveAppHandler(appBox,r))
	r.Post("/users", serveUsers(client))
	r.Post("/newuser", registerNewUser(client))
	r.Post("/newfile", uploadFile(client))
	log.Println("Server start at port 8080")
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "../build/static"))
	filesDir2 := http.Dir(filepath.Join(workDir, "../build"))
	FileServer(r, "/static", filesDir)
	FileServer(r, "/", filesDir2)
	http.ListenAndServe(":8080", r)
}

func downloadFile(c *mongo.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		log.Println(r.URL.Path)
		paths:= strings.Split(r.URL.Path, "/") 
		log.Println(paths[3])
		collection := c.Database("tritonstorage").Collection("fileinfo")
		isSuccess, filename := mydb.GetFilenameById(paths[3], collection)
		if !isSuccess{
			log.Println("database error cannot find the file or non exist")
		}
		targetFile, err:= os.Open("./filestorage/"+filename)

		if err != nil{
			log.Println(err)
		}
		reader := bufio.NewReader(targetFile)
		data := bytes.NewBuffer(make([]byte,0))
		for{
			buffer := make([]byte, 8092)
			count, err := reader.Read(buffer)
			if err != nil{
				log.Println("read file error")
			}
			data.Write(buffer[:count])
			if count < 8092 {
				break
			}
		}
		w.Write(data.Bytes())
	}
}
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
func getAllFile(c *mongo.Client) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("x-user")
		collection := c.Database("tritonstorage").Collection("fileinfo")
		isSuccess, files := mydb.GetAllFilesByUsername(username, collection)
		if !isSuccess {

		}
		jsonData, _ := json.Marshal(files)
		//log.Println(string(jsonData))
		fmt.Fprintf(w, string(jsonData))
	}
}
func uploadFile(c *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*reqStr, err := ioutil.ReadAll(r.Body)
		log.Println(string(reqStr))
		if err != nil{
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}*/
		collection := c.Database("tritonstorage").Collection("fileinfo")
		
		switch r.Method {
			case "POST":
				file, handler, err := r.FormFile("file")
				if err!= nil{
					log.Println("[file]Something wrong")
				}
				defer file.Close()
				fmt.Printf("Uploaded File: %+v\n", handler.Filename)
				fmt.Printf("File Size: %+v\n", handler.Size)
				fmt.Printf("MIME Header: %+v\n", handler.Header)
				
				username := r.Header.Get("x-user")
				isSuccess := mydb.AddFile(handler.Filename, username, collection)
				if !isSuccess {
					log.Println("data insert fail")
					fmt.Fprintf(w, "false")
				}
				
				fileContent := bufio.NewReader(file)
				newfile, err := os.OpenFile("./filestorage/"+handler.Filename, os.O_CREATE|os.O_RDWR, 0755)
				defer newfile.Close()
				if err!=nil{
					log.Println("file :"+handler.Filename+" Creation Error")
					fmt.Fprintf(w, "false")
					return
				}
				blockSize := 8152
				for {
					buffer := make([]byte, blockSize)
					count, err := fileContent.Read(buffer)
					if err!= nil{
						fmt.Fprintf(w, "false")
						return
					}
					newfile.Write(buffer[:count])
					if count < blockSize{
						break
					}
				}
				log.Println("Finish Write File: "+ handler.Filename)
				fmt.Fprintf(w, "true")
		}
	}
}
func registerNewUser(c *mongo.Client) http.HandlerFunc {
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
func serveUsers(c *mongo.Client) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("user")
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


func serveAppHandler(app *rice.Box, m *chi.Mux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("okok")
		indexFile, err := app.Open("index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.ServeContent(w, r, "index.html", time.Time{}, indexFile)
	}
}
