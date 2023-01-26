package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Uptab22( user *user.User ) {
	if user.X < user.Gfb + 1 {
		user.St = 0;
	} else if user.X < 14927 {
		user.Y = ( user.X - user.Gfb ) * 0.0001;
		user.Rw = user.Y * 1088.67;
		user.Rw = user.Rw + 1400;
		user.St = function.Abrunden_double( 0, ( user.Rw * user.Y ) ); // abrunden auf euro
	} else if user.X < 58597 {
		user.Y = ( user.X - 14926 ) * 0.0001;
		user.Rw = user.Y * 206.43;
		user.Rw = user.Rw + 2397;
		user.Rw = user.Rw * user.Y;
		user.St = function.Abrunden_double( 0, ( user.Rw + 869.32 ) ); // abrunden auf ganze euro
	} else if user.X < 277826 {
		user.St = function.Abrunden_double( 0,( user.X * 0.42 ) - 9336.45 ); // abrunden auf ganze euro
	} else {
		user.St = function.Abrunden_double( 0, ( user.X * 0.45 ) - 17671.20 ); // abrunden auf ganze euro
	}
	user.St = user.St * float64( user.Kztab );
}
