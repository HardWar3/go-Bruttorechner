package programme

import (
	"go_bruttoRechner/user"
)

func Mlstjahr( user *user.User ) {
	Upevp( user );
	if user.Kennvmt != 1 {
		user.Zve = user.Zre4 - user.Ztabfb - user.Vsp;
		Upmlst( user );
	} else {
		user.Zve = user.Zre4 - user.Ztabfb - user.Vsp - user.Vmt * 0.01 - user.Vkapa * 0.01;
		if user.Zve < 0 {
			user.Zve = ( user.Zve + user.Vmt * 0.01 + user.Vkapa * 0.01 ) * 5.00;
			Upmlst( user );
			user.St = user.St * 5.00;
		} else {
			Upmlst( user );
			user.Stovmt = user.St;
			user.Zve = user.Zve + ( user.Vmt + user.Vkapa ) * 0.05;
			Upmlst( user );
			user.St = ( user.St - user.Stovmt ) * 5.00 + user.Stovmt;
		}
	}
}
