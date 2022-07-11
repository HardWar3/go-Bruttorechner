package programme

import (
	"../../user"
)

func Upvkv( user *user.User) {
	if user.Pkv > 0 {
		if user.Vsp2 > user.Vsp3 {
			user.Vkv = user.Vsp2 * 100;
		} else {
			user.Vkv = user.Vsp3 * 100;
		}
	} else {
		user.Vkv = 0;
	}
}
