package models

import (
	"github.com/vinitparekh17/project-x/database"
	"github.com/vinitparekh17/project-x/utility"
)

/*
	Identity model that describes how our Identity table looks like
	This table is contains auth specific data i.e. email and password hash
	I did like to keep it minimal and simple to avoid any security issues
*/

type IdentityModel struct {
	UID      int64  `json:"id,omitempty" unique:"true"`
	Email    string `json:"email" unique:"true"`
	Password string `json:"password" min:"8"`
}

func (*IdentityModel) Create(usr IdentityModel) error {
	db := database.Connect()
	defer database.Disconnect(db)
	query := database.Insert{
		Table:  `"user".identity`,
		Fields: []string{"email", "password"},
	}

	smt, err := db.Prepare(query.Build())
	utility.ErrorHandler(err)
	defer smt.Close()
	_, er := smt.Exec(usr.Email, usr.Password)
	utility.ErrorHandler(er)
	return nil
}
