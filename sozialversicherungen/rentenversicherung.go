package sozialversicherungen

import (
	"go_bruttoRechner/user"
)

func Rentenversicherung(user *user.User) {
	const rv float64 = 0.093 // in % Rentenversicherung
	user.Rv = user.Re4 * rv
}
