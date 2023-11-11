package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Uptab24(user *user.User) {
	if user.X < user.Gfb+1 {
		user.St = 0
	} else if user.X < 17006 {
		user.Y = (user.X - user.Gfb) / 10000.00
		user.Rw = user.Y * 922.98
		user.Rw = user.Rw + 1400
		user.St = function.Abrunden_double(0, (user.Rw * user.Y))
	} else if user.X < 66761 {
		user.Y = (user.X - 17005) / 10000.00
		user.Rw = user.Y * 181.19
		user.Rw = user.Rw + 2397
		user.Rw = user.Rw * user.Y
		user.St = function.Abrunden_double(0, (user.Rw + 1025.38))
	} else if user.X < 277826 {
		user.St = function.Abrunden_double(0, (user.X*0.42 - 10602.13))
	} else {
		user.St = function.Abrunden_double(0, (user.X*0.45 - 18936.88))
	}
	user.St = user.St * float64(user.Kztab)
}
