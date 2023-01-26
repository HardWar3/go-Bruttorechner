package sozialversicherungen

import (
	"go_bruttoRechner/user"
)

func Pflegeversicherung(user *user.User) {
	const pv float64 = 1.525  // in % Pflegeversicherung
	const pvs float64 = 2.025 // in % Sachen Pflegeversicherung
	const pvz float64 = 0.25  // in % wenn ohne Kinder und älter 23

	var pv_zusammen float64
	var pv_ag float64

	if user.Pvs == 1 {
		if user.Pvz == 0 {
			pv_zusammen = pvs
		} else {
			pv_zusammen = pvs + pvz
			pv_ag = pvs
		}
	} else {
		if user.Pvz == 0 {
			pv_zusammen = pv
			pv_ag = pv
		} else {
			pv_zusammen = pv + pvz
			pv_ag = pv
		}
	}

	user.Pv_an = (user.Re4 * (pv_zusammen / 100.0))
	user.Pv_ag = (user.Re4 * (pv_ag / 100.0))

}
