package models

// User public model generated by bindu
type User struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required" json:"-"`
	DefaultProperties
}