package programme

import (
	"go_bruttoRechner/user"
)

func Mberech( user *user.User ) {
	Mztabfb( user );
	user.Vfrb = ( user.Anp + user.Fvb + user.Fvbz ) * 100;
	Mlstjahr( user );
	user.Wvfrb = ( user.Zve - user.Gfb ) * 100;

	if user.Wvfrb < 0 {
		user.Wvfrb = 0;
	}

	user.Lstjahr = user.St * user.F;

	Uplstlzz( user );
	Upvkvlzz( user );

	if user.Zkf > 0 {
		user.Ztabfb = user.Ztabfb + user.Kfb;
		Mre4abz( user );
		Mlstjahr( user );
		user.Jbmg = user.St * user.F;
	} else {
		user.Jbmg = user.Lstjahr;
	}
	Msolz( user );
}
