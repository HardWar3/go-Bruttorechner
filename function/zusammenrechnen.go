package function

import (
	"go_bruttoRechner/user"
)

func Zusammenrechnen(user *user.User) {
	user.Gesamt_st = user.Lstlzz + user.Solzlzz + user.Kst
	user.Sum_sozialv_an = user.Rv + user.Av + user.Kv + user.Pv_an
	user.Sum_sozialv_ag = user.Rv + user.Av + user.Kv + user.Pv_ag
	user.Netto_lohn = user.Re4 - (user.Gesamt_st + user.Sum_sozialv_an)
	user.Gesamt_belast_ag = user.Re4 + user.Sum_sozialv_ag

	user.Lstlzz = Abrunden_double(0, user.Lstlzz)
	user.Solzlzz = Abrunden_double(0, user.Solzlzz)
	user.Kst = Abrunden_double(0, user.Kst)
	user.Gesamt_st = Abrunden_double(0, user.Gesamt_st)
	user.Netto_lohn = Abrunden_double(0, user.Netto_lohn)
}
