package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Upanteil( user *user.User ) {
	if user.Lzz == 1 {
		user.Anteil1 = user.Jw;
	} else if user.Lzz == 2 {
		user.Anteil1 = function.Abrunden_double( 2, ( user.Jw / 12.00 ) ); // abrunden
	} else if user.Lzz == 3 {
		user.Anteil1 = function.Abrunden_double( 2, ( user.Jw * 7.00 / 360.00 ) ); // abrunden
	} else {
		user.Anteil1 = function.Abrunden_double( 2, ( user.Jw / 360.00 ) ); // abrunden
	}
}
