package kirchensteuer

import (
	"go_bruttoRechner/user"
)

func Kirchensteuer( user *user.User) {
	const kst float64 = 0.09;	// in % Kirchensteuer
	user.Kst = user.Lstlzz * kst;
}
