package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Mvsp( user *user.User ) {
	if user.Zre4vp > user.Bbgkvpv {
		user.Zre4vp = user.Bbgkvpv;
	}
	if user.Pkv > 0 {
		if user.Stkl == 6 {
			user.Vsp3 = 0;
		} else {
			user.Vsp3 = user.Pkpv * 12.00 * 0.01;
			if user.Pkv == 2 {
				user.Vsp3 = user.Vsp3 - user.Zre4vp * ( user.Kvsatzag + user.Pvsatzag );
			}
		}
	} else {
		user.Vsp3 = user.Zre4vp * ( user.Kvsatzan + user.Pvsatzan );
	}
	user.Vsp = function.Aufrunden_double( 0, ( user.Vsp3 + user.Vsp1 ) ); // aufrunden auf ganze euro
}
