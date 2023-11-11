package programme

import (
	"go_bruttoRechner/user"
)

func Mpara(user *user.User) {

	if user.Krv < 2 {
		if user.Krv == 0 {
			user.Bbgrv = 90600
		} else {
			user.Bbgrv = 89400
		}
		user.Rvsatzan = 0.093
	}

	user.Bbgkvpv = 62100
	user.Kvsatzan = user.Kvz/200 + 0.07
	user.Kvsatzag = 0.0085 + 0.07

	if user.Pvs == 1 {
		user.Pvsatzan = 0.022
		user.Pvsatzag = 0.012
	} else {
		user.Pvsatzan = 0.017
		user.Pvsatzag = 0.017
	}

	if user.Pvz == 1 {
		user.Pvsatzan = user.Pvsatzan + 0.006
	}

	user.W1stkl5 = 13279
	user.W2stkl5 = 33380
	user.W3stkl5 = 222260
	user.Gfb = 11604
	user.Solzfrei = 18130

}
