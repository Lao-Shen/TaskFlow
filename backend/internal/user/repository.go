package user

// 数据库操作

func InsertUser(user User) (User, error) {
	return user, nil
}

func UpdateUser(user User) (User, error) {
	return user, nil
}

func DeleteUser(user User) error {
	return nil
}

func GetUser(user User) (User, error) {
	return user, nil
}

func GetUserByUsername(username string) (User, error) {
	return User{}, nil
}

func GetUserByEmail(email string) (User, error) {
	return User{}, nil
}

func GetUserById(userId string) (User, error) {
	return User{}, nil
}
