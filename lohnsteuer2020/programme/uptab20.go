package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Uptab20( user *user.User ) {
	if user.X < user.Gfb + 1 {
		user.St = 0;
	} else if user.X < 14533 {
		user.Y = ( user.X - user.Gfb ) * 0.0001;
		user.Rw = user.Y * 972.87;
		user.Rw = user.Rw + 1400;
		user.St = function.Abrunden_double( 0, ( user.Rw * user.Y ) ); // abrunden auf euro
	} else if user.X < 57052 {
		user.Y = ( user.X - 14532 ) * 0.0001;
		user.Rw = user.Y * 212.02;
		user.Rw = user.Rw + 2397;
		user.Rw = user.Rw * user.Y;
		user.St = function.Abrunden_double( 0, ( user.Rw + 972.79 ) ); // abrunden auf ganze euro
	} else if user.X < 270501 {
		user.St = function.Abrunden_double( 0,( user.X * 0.42 ) - 8963.74 ); // abrunden auf ganze euro
	} else {
		user.St = function.Abrunden_double( 0, ( user.X * 0.45 ) - 17078.74 ); // abrunden auf ganze euro
	}
	user.St = user.St * float64( user.Kztab );
}
