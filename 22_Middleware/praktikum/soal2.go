package main

import (
"fmt"
"net/http"
"strconv"
"github.com/dgrijalva/jwt-go"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"github.com/labstack/echo"
"github.com/labstack/echo/middleware"

)
var DB *gorm.DB

func init() {
InitDB()
InitialMigration()
}

type Config struct {
DB_Username string
DB_Password string
DB_Port string
DB_Host string
DB_Name string
JWT_Secret string
}

func InitDB() {
config := Config{
DB_Username: "root",
DB_Password: "root123",
DB_Port: "3306",
DB_Host: "localhost",
DB_Name: "crud_go",
JWT_Secret: "mysecretkey",
}
connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	config.DB_Username,
	config.DB_Password,
	config.DB_Host,
	config.DB_Port,
	config.DB_Name,
)

var err error
DB, err = gorm.Open("mysql", connectionString)
if err != nil {
	panic(err)
}
}

type Book struct {
	gorm.Model
	Title string json:"title" form:"title"
	Author string json:"author" form:"author"
	Publisher string json:"publisher" form:"publisher"
	}
	
	type User struct {
	gorm.Model
	Username string json:"username" form:"username"
	Password string json:"password" form:"password"
	Token string json:"token" form:"token"
	}
	
	func InitialMigration() {
	DB.AutoMigrate(&Book{})
	DB.AutoMigrate(&User{})
	}
	func generateToken(username string, secret string) (string, error) {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
claims["username"] = username

tokenString, err := token.SignedString([]byte(secret))
if err != nil {
	return "", err
}

return tokenString, nil
	}

	// login user
func LoginUserController(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user User
if err := DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
	return echo.NewHTTPError(http.StatusBadRequest, "Invalid username or password")
}

tokenString, err := generateToken(user.Username, Config.JWT_Secret)
if err != nil {
	return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
}

user.Token = tokenString
if err := DB.Save(&user).Error; err != nil {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

return c.JSON(http.StatusOK, map[string]interface{}{
	"message": "success login",
	"user":    user,
})
}

// get all books
func GetBooksController(c echo.Context) error {
var books []Book
if err := DB.Find(&books).Error; err != nil {
return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
return c.JSON(http.StatusOK, books)
}

// get book by id
func GetBookByIDController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
var book Book
if err := DB.First(&book, bookID).Error;
