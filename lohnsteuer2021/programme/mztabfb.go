package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Mztabfb( user *user.User ) {
	user.Anp = 0;
	if user.Zvbez >= 0 {
		if user.Zvbez < user.Fvbz {
			user.Fvbz = user.Zvbez;
		}
	}
	if user.Stkl < 6 {
		if user.Zvbez > 0 {
			if user.Zvbez - user.Fvbz < 102 {
				user.Anp = function.Aufrunden_double( 0, ( user.Zvbez - user.Fvbz ) ); // aufrunden auf ganze euro
			} else {
				user.Anp = 102;
			}
		}
	} else {
		user.Fvbz = 0;
		user.Fvbzso = 0;
	}
	if user.Stkl < 6 {
		if user.Zre4 > user.Zvbez {
			if user.Zre4 - user.Zvbez < 1000 {
				user.Anp = function.Aufrunden_double( 0, ( user.Anp + user.Zre4 - user.Zvbez ) ); // aufrunden auf ganze euro
			} else {
				user.Anp = user.Anp + 1000;
			}
		}
	}
	user.Kztab = 1;
	if user.Stkl == 1 {
		user.Sap = 36;
		user.Kfb = user.Zkf * 8388;
	} else if user.Stkl == 2 {
		user.Efa = 1908;
		user.Sap = 36;
		user.Kfb = user.Zkf * 8388;
	} else if user.Stkl == 3 {
		user.Kztab = 2;
		user.Sap = 36;
		user.Kfb = user.Zkf * 8388;
	} else if user.Stkl == 4 {
		user.Sap = 36;
		user.Kfb = user.Zkf * 4194;
	} else if  user.Stkl == 5 {
		user.Sap = 36;
		user.Kfb = 0;
	} else {
		user.Kfb = 0;
	}
	user.Ztabfb = user.Efa + user.Anp + user.Sap + user.Fvbz;
}
