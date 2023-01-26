package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Msonst( user *user.User ) {

	user.Lzz = 1;
	if user.Zmvb == 0 {
		user.Zmvb = 12;
	}

	if user.Sonstb == 0 && user.Mbv == 0 {
		user.Vkvsonst = 0;
		user.Lstso = 0;
		user.Sts = 0;
		user.Solzs = 0;
		user.Bks = 0;
	} else {
		Mosonst( user );
		Upvkv( user );
		user.Vkvsonst = user.Vkv;
		user.Zre4j = ( user.Jre4 + user.Sonstb ) * 0.01;
		user.Zvbezj = ( user.Jvbez + user.Vbs ) * 0.01;
		user.Vbezbso = user.Sterbe;
		Mre4sonst( user );
		Mlstjahr( user );
		user.Wvfrbm = ( user.Zve - user.Gfb ) * 100.00;

		if user.Wvfrbm < 0 {
			user.Wvfrbm = 0;
		}

		Upvkv( user );
		user.Vkvsonst = user.Vkv - user.Vkvsonst;
		user.Lstso = user.St * 100.00;
		user.Sts = function.Abrunden_double( 0, ( user.Lstso - user.Lstoso ) * user.F ); // abrunden auf ganze euro
		Stsmin( user );
	}
}
