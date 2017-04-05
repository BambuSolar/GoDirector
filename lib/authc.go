package lib

import (
	"errors"
	"time"

	"github.com/BambuSolar/GoDirector/models"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/BambuSolar/GoDirector/services"
)

/*
 Get authenticated user and update logintime
*/
func Authenticate(email string, password string, g_recaptcha_response string) (user *models.User, err error) {
	msg := "invalid email or password."
	user = &models.User{Email: email}

	recaptcha := services.Recaptcha{}

	if (recaptcha.Check(g_recaptcha_response)) {

		if err := user.Read("Email"); err != nil {
			if err.Error() == "<QuerySeter> no row found" {
				err = errors.New(msg)
			}
			return user, err
		} else if user.Id < 1 {
			// No user
			return user, errors.New(msg)
		} else if user.Password != convert.StrTo(password).Md5() {
			// No matched password
			return user, errors.New(msg)
		} else {
			user.Lastlogintime = time.Now()
			user.Update("Lastlogintime")
			return user, nil
		}

	}else{
		return nil, errors.New("Recaptcha Failed")
	}
}
