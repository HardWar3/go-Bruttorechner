package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Uptab19( user *user.User ) {
	if user.X < user.Gfb + 1 {
		user.St = 0;
	} else if user.X < 14255 {
		user.Y = ( user.X - user.Gfb ) * 0.0001;
		user.Rw = user.Y * 980.14;
		user.Rw = user.Rw + 1400;
		user.St = function.Abrunden_double( 0, ( user.Rw * user.Y ) ); // abrunden auf euro
	} else if user.X < 55961 {
		user.Y = ( user.X - 14254 ) * 0.0001;
		user.Rw = user.Y * 216.16;
		user.Rw = user.Rw + 2397;
		user.Rw = user.Rw * user.Y;
		user.St = function.Abrunden_double( 0, ( user.Rw + 965.58 ) ); // abrunden auf ganze euro
	} else if user.X < 265327 {
		user.St = function.Abrunden_double( 0,( user.X * 0.42 ) - 8780.9 ); // abrunden auf ganze euro
	} else {
		user.St = function.Abrunden_double( 0, ( user.X * 0.45 ) - 16740.68 ); // abrunden auf ganze euro
	}
	user.St = user.St * float64( user.Kztab );
}
