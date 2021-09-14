package databases

import (
	"project/pbm/config"
	"project/pbm/middlewares"
	"project/pbm/models"
)

func CheckEmail(email string) (bool, error) {
	var user models.User

	if err := config.DB.Model(&user).Where("email=?", email).First(&user).Error; err != nil {
		return false, err
	}

	if user.Email == email {
		return true, nil
	} else {
		return false, nil
	}
}

func GetLevel(id int) (models.Level, error) {
	var level models.Level
	if err := config.DB.Where("id=?", id).First(&level).Error; err != nil {
		return level, err
	}

	return level, nil
}

func GetPassword(email string) (string, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user.Password, err
	}
	return user.Password, nil
}

func GetUserById(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Where("id=?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func Register(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func Login(email string) (models.User, error) {
	var user models.User
	var err error
	if err = config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return user, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
