package databases

import (
	"project/pbm/config"
	"project/pbm/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mock_user = models.User{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		LevelID:  1,
	}
	mock_user_login = models.User{
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mock_user_update = models.User{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		LevelID:  1,
	}
)

func TestRegisterSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	user, err := Register(mock_user)
	if assert.NoError(t, err) {
		assert.Equal(t, "Riska", user.Name)
		assert.Equal(t, "riska@gmail.com", user.Email)
		assert.Equal(t, "123", user.Password)
	}
}

func TestRegisterError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	_, err := Register(mock_user)
	assert.Error(t, err)
}

func TestCheckEmailSame(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	Register(mock_user)
	same, err := CheckEmail("riska@gmail.com")
	if assert.NoError(t, err) {
		assert.Equal(t, true, same)
	}
}

func TestCheckEmailError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	Register(mock_user)
	_, err := CheckEmail("riska@gmail.com")
	assert.Error(t, err)
}

func TestLoginSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	create_user, _ := Register(mock_user)
	user, err := Login(create_user.Email)
	if assert.NoError(t, err) {
		assert.Equal(t, "Riska", user.Name)
		assert.Equal(t, "riska@gmail.com", user.Email)
		assert.Equal(t, "123", user.Password)
	}
}

func TestLoginWrongEmail(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	create_user, _ := Register(mock_user)
	_, err := Login(create_user.Name)
	assert.Error(t, err)
}

func TestGetUserByIdSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	create_user, _ := Register(mock_user)
	user, err := GetUserById(int(create_user.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Riska", user.Name)
		assert.Equal(t, "riska@gmail.com", user.Email)
		assert.Equal(t, "123", user.Password)
	}
}

func TestGetUserByIdError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	Register(mock_user)
	_, err := GetUserById(1)
	assert.Error(t, err)
}

func TestGetLevelSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Level{})
	config.DB.Migrator().AutoMigrate(&models.Level{})
	new_level := models.Level{
		Name: "Bos",
	}
	if err := config.DB.Save(&new_level).Error; err != nil {
		t.Error(err)
	}
	create_user, _ := Register(mock_user)
	level, err := GetLevel(int(create_user.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Bos", level.Name)
	}
}

func TestGetLevelError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Level{})
	Register(mock_user)
	_, err := GetLevel(1)
	assert.Error(t, err)
}

func TestGetPasswordSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	create_user, _ := Register(mock_user)
	password, err := GetPassword(create_user.Email)
	if assert.NoError(t, err) {
		assert.Equal(t, "123", password)
	}
}

func TestGetPasswordError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	Register(mock_user)
	_, err := GetPassword("riska@gmail.com")
	assert.Error(t, err)
}

func TestUpdateUserSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	create_user, _ := Register(mock_user)
	user, err := UpdateUser(create_user)
	if assert.NoError(t, err) {
		assert.Equal(t, "Riska", user.Name)
		assert.Equal(t, "riska@gmail.com", user.Email)
		assert.Equal(t, "123", user.Password)
	}
}

func TestUpdateUserError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	create_user, _ := Register(mock_user)
	_, err := UpdateUser(create_user)
	assert.Error(t, err)
}
