package sozialversicherungen

import (
	"go_bruttoRechner/user"
)

func Arbeitslosenversicherung(user *user.User) {
	const av float64 = 0.0125
	user.Av = user.Re4 * av
}
