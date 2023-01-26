package sozialversicherungen

import (
	"go_bruttoRechner/user"
)

func Sozialversicherungen(user *user.User) {

	Arbeitslosenversicherung(user)
	Krankenversicherung(user)
	Pflegeversicherung(user)
	Rentenversicherung(user)

}
