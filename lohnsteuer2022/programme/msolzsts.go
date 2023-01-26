package programme

import (
	"go_bruttoRechner/function"
	"go_bruttoRechner/user"
)

func Msolzsts( user *user.User ) {
	if user.Zkf > 0 {
		user.Solzszve = user.Zve - user.Kfb;
	} else {
		user.Solzszve = user.Zve;
	}

	if user.Solzszve < 1 {
		user.Solzszve = 0;
		user.X = 0;
	} else {
		user.X = function.Abrunden_double(0, user.Solzszve / float64(user.Kztab)); // abrunden auf ganze euro
	}

	if user.Stkl < 5 {
		Uptab22( user );
	} else {
		Mst5_6( user );
	}

	user.Solzsbmg = user.St * user.F;
	
	if user.Solzsbmg > user.Solzfrei {
		user.Solzs = function.Abrunden_double( 2, ( user.Sts * 5.5 * 0.01 ) ); // abrunden auf ganze cent
	} else {
		user.Solzs = 0;
	}
}
