package sv19

import (
	"../user"
)

func Sv_kv_19( user *user.User ) {
	const kv float64 = 0.0775;	// in % Krankenversicherung
	user.Kv = user.Re4 * kv;
}
