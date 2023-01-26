package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Up5_6( user *user.User ) {
	user.X = user.Zx * 1.25;
	Uptab20( user );
	user.X = user.Zx * 0.75;
	Uptab20( user );
	user.St2 = user.St;
	user.Diff = ( user.St1 - user.St2 ) * 2;
	user.Mist = function.Abrunden_double( 0, user.Zx * 0.14 ); // abrunden auf ganze euro
	if user.Mist > user.Diff {
		user.St = user.Mist;
	} else {
		user.St = user.Diff;
	}
}
