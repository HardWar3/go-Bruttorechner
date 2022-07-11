package sv19

import (
	"../user"
)

func Sv_rv_19( user *user.User ) {
	const rv float64 = 0.093;	// in % Rentenversicherung
	user.Rv = user.Re4 * rv;
}
