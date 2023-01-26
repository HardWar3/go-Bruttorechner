package programme

import (
	"go_bruttoRechner/user"
)

func Mpara( user *user.User ) {

	if user.Krv < 2 {
		if user.Krv == 0 {
			user.Bbgrv = 85200;
		} else {
			user.Bbgrv = 80400;
		}

		user.Rvsatzan = 0.0930;
		user.Tbsvorv = 0.84;

	}

	user.Bbgkvpv = 58050;
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

	user.W1stkl5 = 11237;
	user.W1stkl5 = 28959;
	user.W1stkl5 = 219690;
	user.Gfb = 9744;
	user.Solzfrei = 16956;

}
