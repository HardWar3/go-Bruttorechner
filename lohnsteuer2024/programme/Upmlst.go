package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Upmlst(user *user.User) {
	if user.Zve < 1 {
		user.Zve = 0
		user.X = 0
	} else {
		user.X = function.Abrunden_double(0, (user.Zve / float64(user.Kztab)))
	}
	if user.Stkl < 5 {
		Uptab24(user)
	} else {
		Mst5_6(user)
	}
}
