package mydb

type User struct{
	Email	string
	Password string
	Username string
}

type File struct{
	Username string
	Filename string
	Type string
	Id string
}

type FileGroup struct{
	Files []File
}