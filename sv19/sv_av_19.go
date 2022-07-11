package sv19

import (
	"../user"
)

func Sv_av_19( user *user.User ) {
	const av float64  = 0.0125;
	user.Av = user.Re4 * av;
}
