package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Mst5_6( user *user.User ) {
	user.Zzx = user.X;
	if user.Zzx > user.W2stkl5 {
		user.Zx = user.W2stkl5;
		Up5_6( user );
		if user.Zzx > user.W3stkl5 {
			user.St = function.Abrunden_double( 0, user.St + ( user.W3stkl5 - user.W2stkl5 )  * 0.42 ); // abrunden auf ganze euro
			user.St = function.Abrunden_double( 0, user.St + ( user.Zzx - user.W3stkl5 ) * 0.45 ); // abrunden auf ganze euro
		} else {
			user.St = function.Abrunden_double( 0, user.St + ( user.Zzx - user.W2stkl5 ) * 0.42 ); // abrunden auf ganze euro
		}
	} else {
		user.Zx = user.Zzx;
		Up5_6( user );
		if user.Zzx > user.W1stkl5 {
			user.Vergl = user.St;
			user.Zx = user.W1stkl5;
			Up5_6( user );
			user.Hoch = function.Abrunden_double( 0, user.St + ( user.Zzx - user.W1stkl5 ) * 0.42 ); // abrunden auf ganz euro
			if user.Hoch < user.Vergl {
				user.St = user.Hoch;
			} else {
				user.St = user.Vergl;
			}
		}
	}
}
