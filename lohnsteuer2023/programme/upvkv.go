package programme

import (
	"go_bruttoRechner/user"
)

func Upvkv( user *user.User) {
	if user.Pkv > 0 {
		if user.Vsp2 > user.Vsp3 {
			user.Vkv = user.Vsp2 * 0.01;
		} else {
			user.Vkv = user.Vsp3 * 0.01;
		}
	} else {
		user.Vkv = 0;
	}
}
