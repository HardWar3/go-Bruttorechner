package programme

import (
	"go_bruttoRechner/user"
)

func Mre4jl( user *user.User ) {

	if user.Lzz == 1 {
		user.Zre4j = user.Re4 * 0.01;
		user.Zvbezj = user.Vbez * 0.01;
		user.Jlfreib = user.Lzzfreib * 0.01;
		user.Jlhinzu = user.Lzzhinzu * 0.01;
	} else if user.Lzz == 2 {
		user.Zre4j = ( user.Re4 * 12 ) * 0.01;
		user.Zvbezj = ( user.Vbez * 12 ) * 0.01;
		user.Jlfreib = ( user.Lzzfreib * 12 ) * 0.01;
		user.Jlhinzu = ( user.Lzzhinzu * 12 ) * 0.01;
	} else if user.Lzz == 3 {
		user.Zre4j = ( user.Re4 * 360 / 7 ) * 0.01;
		user.Zvbezj = ( user.Vbez * 360 / 7 ) * 0.01;
		user.Jlfreib = ( user.Lzzfreib * 360 / 7 ) * 0.01;
		user.Jlhinzu = ( user.Lzzhinzu * 360 / 7 ) * 0.01;
	} else if user.Lzz == 4 {
		user.Zre4j = ( user.Re4 * 360 ) * 0.01;
		user.Zvbezj = ( user.Vbez * 360 ) * 0.01;
		user.Jlfreib = ( user.Lzzfreib * 360 ) * 0.01;
		user.Jlhinzu = ( user.Lzzhinzu * 360 ) * 0.01;
	}

	if user.Af == 0 {
		user.F = 1;
	}

}
