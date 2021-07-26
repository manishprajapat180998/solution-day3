package Models

import (
	"fmt"
	"github.com/exercise2/Config"
	_ "github.com/go-sql-driver/mysql"
)
//GetAllUsers Fetch all user data
func GetAllUsers(person *[]Person) (err error) {
	if err = Config.DB.Find(person).Error; err != nil {
		return err
	}
	return nil
}
//CreateUser ... Insert New data
func CreateUser(person *Person) (err error) {
	if err = Config.DB.Create(person).Error; err != nil {
		return err
	}
	return nil
}
//GetUserByID ... Fetch only one user by Id
func GetUserByID(person *Person, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(person).Error; err != nil {
		return err
	}
	return nil
}
//UpdateUser ... Update user
func UpdateUser(person *Person, id string) (err error) {
	fmt.Println(person)
	Config.DB.Save(person)
	return nil
}
//DeleteUser ... Delete user
func DeleteUser(person *Person, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(person)
	return nil
}