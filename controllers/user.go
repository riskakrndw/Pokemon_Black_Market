package controllers

import (
	"log"
	"net/http"
	"project/pbm/databases"
	"project/pbm/middlewares"
	"project/pbm/models"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserOutput struct {
	ID    uint   `json:"id"`
	Level string `json:"level"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserOutput1 struct {
	ID    uint   `json:"id"`
	Level string `json:"level"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func EncryptPwd(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func RegisterBos(c echo.Context) error {
	//get user's input
	input_user := models.User{}
	input_user.LevelID = 1
	c.Bind(&input_user)

	//check is data nil?
	if input_user.Email == "" || input_user.Password == "" || input_user.Name == "" || input_user.LevelID == 0 {
		return c.JSON(http.StatusBadRequest, "Please fill all data")
	}

	//check is email exists?
	is_email_exists, _ := databases.CheckEmail(input_user.Email)
	if is_email_exists {
		return c.JSON(http.StatusBadRequest, "Email already exists")
	}

	//encrypt pass user
	convert_pwd := []byte(input_user.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	input_user.Password = hashed_pwd //set new pass

	//create new user
	user, err := databases.Register(input_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot insert data")
	}

	//get level name
	level, err := databases.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	output := UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(http.StatusOK, output)
}

func RegisterOperasional(c echo.Context) error {
	//get user's input
	input_user := models.User{}
	input_user.LevelID = 2
	c.Bind(&input_user)

	//check is data nil?
	if input_user.Email == "" || input_user.Password == "" || input_user.Name == "" || input_user.LevelID == 0 {
		return c.JSON(http.StatusBadRequest, "Please fill all data")
	}

	//check is email exists?
	is_email_exists, _ := databases.CheckEmail(input_user.Email)
	if is_email_exists {
		return c.JSON(http.StatusBadRequest, "Email already exist")
	}

	//encrypt pass user
	convert_pwd := []byte(input_user.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	input_user.Password = hashed_pwd //set new pass

	//create new user
	user, err := databases.Register(input_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot insert data")
	}

	//get level name
	level, err := databases.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	output := UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(http.StatusOK, output)
}

func RegisterPengedar(c echo.Context) error {
	//get user's input
	input_user := models.User{}
	input_user.LevelID = 3
	c.Bind(&input_user)

	//check is data nil?
	if input_user.Email == "" || input_user.Password == "" || input_user.Name == "" || input_user.LevelID == 0 {
		return c.JSON(http.StatusBadRequest, "Please fill all data")
	}

	//check is email exists?
	is_email_exists, _ := databases.CheckEmail(input_user.Email)
	if is_email_exists {
		return c.JSON(http.StatusBadRequest, "Email already exist")
	}

	//encrypt pass user
	convert_pwd := []byte(input_user.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	input_user.Password = hashed_pwd //set new pass

	//create new user
	user, err := databases.Register(input_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot insert data")
	}

	//get level name
	level, err := databases.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	output := UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(http.StatusOK, output)
}

func Login(c echo.Context) error {
	//get user's input
	input_user := models.User{}
	c.Bind(&input_user)

	//check is data nil?
	if input_user.Email == "" || input_user.Password == "" {
		return c.JSON(http.StatusBadRequest, "Please fill all data")
	}

	//compare password on form with db
	get_pwd, x := databases.GetPassword(input_user.Email) //get password
	if x != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	err := bcrypt.CompareHashAndPassword([]byte(get_pwd), []byte(input_user.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "User Unauthorized. Email or Password not equal")
	}

	//login
	user, err := databases.Login(input_user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//get level name
	level, err := databases.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	output := UserOutput1{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
		Token: user.Token,
	}

	return c.JSON(http.StatusOK, output)
}

func GetProfile(c echo.Context) error {
	//get id user login
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}

	//get customer by id
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//get level name
	level, err := databases.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	output := UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(http.StatusOK, output)
}

func GetProfileTesting() echo.HandlerFunc {
	return GetProfile
}

func Logout(c echo.Context) error {
	//get id user login
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}

	//get customer by id
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	user.Token = ""
	c.Bind(&user)
	customer_updated, err := databases.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}

	//get level name
	level, err := databases.GetLevel(int(customer_updated.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	output := UserOutput1{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
		Token: user.Token,
	}

	return c.JSON(http.StatusOK, output)
}
