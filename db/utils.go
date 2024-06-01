package db

import (
	"budget-be/utils"
)

func mapUserToStruct(scanner utils.CustomScanner, user *utils.User) error {
	err := scanner.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.Password)

	if err != nil {
		return err
	}

	return nil
}
