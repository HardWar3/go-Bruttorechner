package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Uptab23( user *user.User ) {
	if user.X < user.Gfb + 1 {
		user.St = 0;
	} else if user.X < 16000 {
		user.Y = ( user.X - user.Gfb ) * 0.0001;
		user.Rw = user.Y * 979.18;
		user.Rw = user.Rw + 1400;
		user.St = function.Abrunden_double( 0, ( user.Rw * user.Y ) ); // abrunden auf euro
	} else if user.X < 62810 {
		user.Y = ( user.X - 15999 ) * 0.0001;
		user.Rw = user.Y * 192.59;
		user.Rw = user.Rw + 2397;
		user.Rw = user.Rw * user.Y;
		user.St = function.Abrunden_double( 0, ( user.Rw + 966.53 ) ); // abrunden auf ganze euro
	} else if user.X < 277826 {
		user.St = function.Abrunden_double( 0,( user.X * 0.42 ) - 9972.98 ); // abrunden auf ganze euro
	} else {
		user.St = function.Abrunden_double( 0, ( user.X * 0.45 ) - 18307.73 ); // abrunden auf ganze euro
	}
	user.St = user.St * float64( user.Kztab );
}
