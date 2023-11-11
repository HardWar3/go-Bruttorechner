package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Stsmin(user *user.User) {
	if user.Sts < 0 {
		if user.Mbv != 0 {
			user.Lstlzz = user.Lstlzz + user.Sts
			if user.Lstlzz < 0 {
				user.Lstlzz = 0
			}
			user.Solzlzz = function.Abrunden_double(2, (user.Solzlzz + user.Sts*5.5*0.01))
			if user.Solzlzz < 0 {
				user.Solzlzz = 0
			}
			user.Bk = user.Bk + user.Sts
			if user.Bk < 0 {
				user.Bk = 0
			}
		}
		user.Sts = 0
		user.Solzs = 0
	} else {
		Msolzsts(user)
	}
	if user.R > 0 {
		user.Bks = user.Sts
	} else {
		user.Bks = 0
	}
}
