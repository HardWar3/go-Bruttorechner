package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Upevp(user *user.User) {
	if user.Krv > 1 {
		user.Vsp1 = 0
	} else {
		if user.Zre4vp > user.Bbgrv {
			user.Zre4vp = user.Bbgrv
		}
		user.Vsp1 = user.Zre4vp * user.Rvsatzan
	}
	user.Vsp2 = 0.12 * user.Zre4vp
	if user.Stkl == 3 {
		user.Vhb = 3000
	} else {
		user.Vhb = 1900
	}
	if user.Vsp2 > user.Vhb {
		user.Vsp2 = user.Vhb
	}
	user.Vspn = function.Aufrunden_double(0, (user.Vsp1 + user.Vsp2))
	Mvsp(user)
	if user.Vspn > user.Vsp {
		user.Vsp = user.Vspn
	}
}
