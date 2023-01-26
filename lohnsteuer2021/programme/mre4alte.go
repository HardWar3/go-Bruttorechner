package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Mre4alte( user *user.User) {
	if user.Alter1 == 0 {
		user.Alte = 0;
	} else {
		if user.Ajahr < 2006 {
			user.K = 1;
		} else if user.Ajahr < 2040 {
			user.K = user.Ajahr - 2004;
		} else {
			user.K = 36;
		}
		user.Bmg = user.Zre4j - user.Zvbezj;
		user.Alte = function.Aufrunden_double( 0 ,( user.Bmg * function.Tab( 4, user.K ) ) ); // aufrunden auf euro
		user.Hbalte = function.Tab( 5, user.K );
		if user.Alte > user.Hbalte {
			user.Alte = user.Hbalte;
		}
	}
}
