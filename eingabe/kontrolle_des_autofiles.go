package eingabe

import (
	"fmt"
)

func Kontrolle_des_autofiles( eingabe *string ) bool {

	for index := 0; index < len(*eingabe); index++ {

		eingabeChar := string((*eingabe)[index]);

		if eingabeChar == "," || eingabeChar == "." {
			fmt.Println( (*eingabe) );
			*eingabe = (*eingabe)[:index] + (*eingabe)[(index+1):];
			fmt.Println( (*eingabe) );
		} else if eingabeChar < "0" || eingabeChar > "9" {
			return true;
		}
	}

	return false;

}
