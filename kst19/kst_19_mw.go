package kst19

import (
	"../user"
)

func Kst_19_MW( user *user.User) {
	const kst float64 = 0.09;	// in % Kirchensteuer
	user.Kst = user.Lstlzz * kst;
}
