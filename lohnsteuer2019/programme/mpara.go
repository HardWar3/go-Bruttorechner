package programme

import (
	"go_bruttoRechner/user"
)

func Mpara( user *user.User ) {

	if user.Krv < 2 {
		if user.Krv == 0 {
			user.Bbgrv = 80400;
		} else {
			user.Bbgrv = 73800;
		}

		user.Rvsatzan = 0.0930;
		user.Tbsvorv = 0.76;

	}

	user.Bbgkvpv = 54450;
	user.Kvsatzan = user.Kvz / 200 + 0.07;
	user.Kvsatzag = 0.0045 + 0.07;

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

	user.W1stkl5 = 10635;
	user.W1stkl5 = 27980;
	user.W1stkl5 = 212261;
	user.Gfb = 9168;
	user.Solzfrei = 972;

}
