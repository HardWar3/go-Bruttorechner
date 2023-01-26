package programme

import (
	"go_bruttoRechner/user"
)

func Mpara(user *user.User) {

	if user.Krv < 2 {
		if user.Krv == 0 {
			user.Bbgrv = 87600
		} else {
			user.Bbgrv = 85200
		}

		user.Rvsatzan = 0.0930
		user.Tbsvorv = 1.00

	}

	user.Bbgkvpv = 59850
	user.Kvsatzan = user.Kvz/200 + 0.07
	user.Kvsatzag = 0.008 + 0.07

	if user.Pvs == 1 {
		user.Pvsatzan = 0.02025
		user.Pvsatzag = 0.01025
	} else {
		user.Pvsatzan = 0.01525
		user.Pvsatzag = 0.01525
	}

	if user.Pvz == 1 {
		user.Pvsatzan = user.Pvsatzan + 0.0035
	}

	user.W1stkl5 = 12485
	user.W1stkl5 = 31404
	user.W1stkl5 = 222260
	user.Gfb = 10908
	user.Solzfrei = 17543

}
