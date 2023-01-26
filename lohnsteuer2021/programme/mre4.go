package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Mre4( user *user.User ) {
	if user.Zvbezj == 1 {
		user.Fvbz = 0;
		user.Fvb = 0;
		user.Fvbzso = 0;
		user.Fvbso = 0;
	} else {
		if user.Vjahr < 2006 {
			user.J = 1;
		} else if user.Vjahr < 2040 {
			user.J = user.Vjahr - 2004;
		} else {
			user.J = 36;
		}

		if user.Lzz == 1 {
			user.Vbezb = ( user.Vbezm * float64( user.Zmvb ) ) + user.Vbezs;
			user.Hfvb = ( function.Tab( 2, user.J ) / 12.0 ) * float64( user.Zmvb );
			user.Fvbz = function.Aufrunden_double( 0, ( function.Tab( 3, user.J ) / 12 ) * float64( user.Zmvb ) ); // aufrunden auf euro
		} else {
			user.Vbezb = user.Vbezm * 12 + user.Vbezs;
			user.Hfvb = function.Tab( 2, user.J );
			user.Fvbz = function.Tab( 3, user.J );
		}

		user.Fvb = function.Aufrunden_double( 2, ( ( user.Vbezb * function.Tab( 1, user.J ) ) * 0.01 ) ); // aufrunden auf cent

		if user.Fvb > user.Hfvb {
			user.Fvb = user.Zvbezj;
		}

		if user.Fvb > user.Zvbezj {
			user.Fvb = user.Zvbezj;
		}

		user.Fvbso = function.Aufrunden_double( 2, ( user.Fvb + ( user.Vbezbso * function.Tab( 1, user.J ) ) * 0.01  ) ); // aufrunden auf cent

		if user.Fvbso > function.Tab( 2, user.J ) {
			user.Fvbso = function.Tab( 2, user.J );
		}

		user.Hfvbzso = ( user.Vbezb + user.Vbezbso ) * 0.01 - user.Fvbso;
		user.Fvbzso = function.Aufrunden_double( 0, ( user.Fvbz + user.Vbezbso * 0.01 ) ); // aufrunden auf euro

		if user.Fvbzso > user.Hfvbzso {
			user.Fvbzso = function.Aufrunden_double( 2, ( user.Hfvbzso ) ); // aufrunden auf euro
		}

		if user.Fvbzso > function.Tab( 3, user.J ) {
			user.Fvbzso = function.Tab( 3, user.J );
		}

		user.Hfvbz = user.Vbezb * 0.01 - user.Fvb;

		if user.Fvbz > user.Hfvbz {
			user.Fvbz = function.Aufrunden_double( 2, ( user.Hfvbz ) ); // aufrunden auf euro
		}
	}

	Mre4alte( user );

}
