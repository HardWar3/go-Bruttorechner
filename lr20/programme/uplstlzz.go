package programme

import (
	"../../user"
)

func Uplstlzz( user *user.User ) {
	user.Jw = user.Lstjahr * 100;
	Upanteil( user );
	user.Lstlzz = user.Anteil1;
}
