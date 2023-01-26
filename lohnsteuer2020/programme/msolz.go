package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Msolz( user *user.User ) {
	user.Solzfrei = user.Solzfrei * float64( user.Kztab );
	if user.Jbmg > user.Solzfrei {
		user.Solzj = function.Abrunden_double( 2, ( user.Jbmg * 5.5 * 0.01 ) ); // abrunden auf cent
		user.Solzmin = ( user.Jbmg - user.Solzfrei ) * 0.20;
		if user.Solzmin < user.Solzj {
			user.Solzj = user.Solzmin;
		}
		user.Jw = user.Solzj * 100;
		Upanteil( user );
		user.Solzlzz = user.Anteil1;
	} else {
		user.Solzlzz = 0;
	}
	if user.R > 0 {
		user.Jw = user.Jbmg * 100;
		Upanteil( user );
		user.Bk = user.Anteil1;
	} else {
		user.Bk = 0;
	}
}
