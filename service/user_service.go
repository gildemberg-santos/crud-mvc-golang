package service

import (
	"github.com/gildemberg-santos/crud-mvc-golang/entity"
	"github.com/gildemberg-santos/crud-mvc-golang/model"
)

func UserCreate(user_request entity.User) error {
	var user model.User
	db, err := model.NewDB(&user)
	if err != nil {
		return err
	}

	user = model.User{
		Fullname:      user_request.Fullname,
		TaxIdentifier: user_request.TaxIdentifier,
		Email:         user_request.Email,
		Password:      user_request.Password,
	}
	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UserUpdate(user_request entity.User) error {
	var user model.User
	db, err := model.NewDB(&user)
	if err != nil {
		return err
	}

	result := db.Find(&user, user_request.ID)
	if result.Error != nil {
		return result.Error
	}

	user.Fullname = user_request.Fullname
	user.TaxIdentifier = user_request.TaxIdentifier
	user.Email = user_request.Email
	user.Password = user_request.Password

	result = db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UserFindAll() ([]model.User, error) {
	var user model.User
	db, err := model.NewDB(&user)
	if err != nil {
		return nil, err
	}

	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func UserFindByID(id string) (model.User, error) {
	var user model.User
	db, err := model.NewDB(&user)
	if err != nil {
		return user, err
	}

	result := db.Find(&user, id)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func UserDeleteByID(id string) error {
	var user model.User
	db, err := model.NewDB(&user)
	if err != nil {
		return err
	}

	result := db.Find(&user, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
