package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Uptab21( user *user.User ) {
	if user.X < user.Gfb + 1 {
		user.St = 0;
	} else if user.X < 14754 {
		user.Y = ( user.X - user.Gfb ) * 0.0001;
		user.Rw = user.Y * 995.21;
		user.Rw = user.Rw + 1400;
		user.St = function.Abrunden_double( 0, ( user.Rw * user.Y ) ); // abrunden auf euro
	} else if user.X < 57919 {
		user.Y = ( user.X - 14753 ) * 0.0001;
		user.Rw = user.Y * 208.85;
		user.Rw = user.Rw + 2397;
		user.Rw = user.Rw * user.Y;
		user.St = function.Abrunden_double( 0, ( user.Rw + 950.96 ) ); // abrunden auf ganze euro
	} else if user.X < 274613 {
		user.St = function.Abrunden_double( 0,( user.X * 0.42 ) - 9136.63 ); // abrunden auf ganze euro
	} else {
		user.St = function.Abrunden_double( 0, ( user.X * 0.45 ) - 17374.99 ); // abrunden auf ganze euro
	}
	user.St = user.St * float64( user.Kztab );
}
