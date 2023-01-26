package lohnsteuer2022

import (
	"go_bruttoRechner/user"

	"go_bruttoRechner/lohnsteuer2022/programme"
)

func Lohnsteuer2022(user *user.User) {

	programme.Mpara(user)
	programme.Mre4jl(user)
	user.Vbezbso = 0
	user.Kennvmt = 0
	programme.Mre4(user)
	programme.Mre4abz(user)
	programme.Mberech(user)
	programme.Msonst(user)
	programme.Mvmt(user)

}
