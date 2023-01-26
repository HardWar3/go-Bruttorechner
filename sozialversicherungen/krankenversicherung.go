package sozialversicherungen

import (
	"go_bruttoRechner/user"
)

func Krankenversicherung(user *user.User) {
	const kv float64 = 0.0775 // in % Krankenversicherung
	user.Kv = user.Re4 * kv
}
