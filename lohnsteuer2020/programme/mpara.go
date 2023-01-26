package programme

import (
	"go_bruttoRechner/user"
)

func Mpara( user *user.User ) {

	if user.Krv < 2 {
		if user.Krv == 0 {
			user.Bbgrv = 82800;
		} else {
			user.Bbgrv = 77400;
		}

		user.Rvsatzan = 0.0930;
		user.Tbsvorv = 0.8;

	}

	user.Bbgkvpv = 56250;
	user.Kvsatzan = user.Kvz / 200 + 0.07;
	user.Kvsatzag = 0.0055 + 0.07;

	if user.Pvs == 1 {
		user.Pvsatzan = 0.02025;
		user.Pvsatzag = 0.01025;
	} else {
		user.Pvsatzan = 0.01525;
		user.Pvsatzag = 0.01525;
	}

	if user.Pvz == 1 {
		user.Pvsatzan = user.Pvsatzan + 0.0025;
	}

	user.W1stkl5 = 10898;
	user.W1stkl5 = 28526;
	user.W1stkl5 = 216400;
	user.Gfb = 9408;
	user.Solzfrei = 972;

}
