package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Mvmt( user *user.User ) {
	if user.Vkapa < 0 {
		user.Vkapa = 0;
	}

	if user.Vmt + user.Vkapa > 0 {

		if user.Lstso == 0 {
			Mosonst( user );
			user.Lst1 = user.Lstoso;
		} else {
			user.Lst1 = user.Lstso;
		}

		user.Vbezbso = user.Sterbe + user.Vkapa;
		user.Zre4j = ( user.Jre4 + user.Sonstb + user.Vmt + user.Vkapa ) * 0.01;
		user.Zvbezj = ( user.Jvbez + user.Vbs + user.Vkapa ) * 0.01;
		user.Kennvmt = 2;
		Mre4sonst( user );
		Mlstjahr( user );
		user.Lst3 = user.St * 100;
		Mre4abz( user );
		user.Zre4vp = user.Zre4vp - user.Jre4ent * 0.01 - user.Sonstent * 0.01;
		user.Kennvmt = 1;
		Mlstjahr( user );
		user.Lst2 = user.St * 100;
		user.Stv = user.Lst2 - user.Lst1;
		user.Lst3 = user.Lst3 - user.Lst1;

		if user.Lst3 < user.Stv {
			user.Stv = user.Lst3;
		}

		if user.Stv < 0 {
			user.Stv = 0;
		} else {
			user.Stv = function.Abrunden_double( 0, ( user.Stv * user.F ) ); // abrunden auf ganze euro
		}

		user.Solzsbmg = user.Stv * 0.01 + user.Jbmg;

		if user.Solzsbmg > user.Solzfrei {
			user.Solzv = function.Abrunden_double( 2, ( user.Stv * 5.5 * 0.01) ); // abrunden auf ganzen cent
		} else {
			user.Solzv = 0;
		}

		if  user.R > 0 {
			user.Bkv = user.Stv;
		} else {
			user.Bkv = 0;
		}

	} else {
		user.Stv = 0;
		user.Solzv = 0;
		user.Bkv = 0;
	}

}
