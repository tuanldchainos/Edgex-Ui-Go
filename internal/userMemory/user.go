package userMemory

import (
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/core"
)

var BasicUser User

type User struct {
	Name string
	Pass string
}

func SetUserPassword() {
	BasicUser = User{
		Name: core.UserName,
		Pass: core.UserPass,
	}
}

func UpdateUserPass(pass string) {
	BasicUser = User{
		Name: core.UserName,
		Pass: pass,
	}
}
