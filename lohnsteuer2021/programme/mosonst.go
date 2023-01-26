package programme

import (
	"go_bruttoRechner/user"
)

func Mosonst( user *user.User ) {
	user.Zre4j = user.Jre4 * 0.01;
	user.Zvbezj = user.Jvbez * 0.01;
	user.Jlfreib = user.Jfreib * 0.01;
	user.Jlhinzu = user.Jhinzu * 0.01;
	Mre4( user );
	Mre4abz( user );
	user.Zre4vp = user.Zre4vp - user.Jre4ent * 0.01;
	Mztabfb( user );
	user.Vfrbs1 = ( user.Anp + user.Fvb + user.Fvbz ) * 100;
	Mlstjahr( user );
	user.Wvfrbo = ( user.Zve - user.Gfb ) * 100;
	if user.Wvfrbo < 0 {
		user.Wvfrbo = 0;
	}
	user.Lstoso = user.St * 100;
}
