package lib

import (
	"errors"

	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/BambuSolar/GoDirector/models"
	"fmt"
)

func SignupUser(u *models.User) (int64, error) {
	var (
		err error
		msg string
	)

	if models.Users().Filter("email", u.Email).Exist() {
		msg = "was already regsitered input email address."
		return 0, errors.New(msg)
	}

	u.Password = convert.StrTo(u.Password).Md5()

	err = u.Insert()

	fmt.Println(err.Error())

	if err != nil {
		return 0, err
	}

	return u.Id, err
}
