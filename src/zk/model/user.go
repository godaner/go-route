package model
type User struct {
	Username        string  `json:"username" bson:"username"`
	Password	string `json:"password" bson:"password"`
}
var UserCol = "user"
var DB = "gowebdemo"