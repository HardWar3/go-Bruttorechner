package programme

import (
	"go_bruttoRechner/user"
)

func Uplstlzz(user *user.User) {
	user.Jw = user.Lstjahr * 100
	Upanteil(user)
	user.Lstlzz = user.Anteil1
}
