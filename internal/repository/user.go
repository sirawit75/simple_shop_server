package repository

import "sirawit/shop/internal/model"

func (u *userQuery) Register(input model.User) (*model.User, error) {
	if err := u.db.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (u *userQuery) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
