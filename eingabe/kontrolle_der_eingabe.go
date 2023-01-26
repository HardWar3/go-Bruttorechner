package eingabe

import (
	"fmt"
)

func Kontrolle_der_eingabe( eingabe *string ) bool {

	var only_once_komma bool = false;

	for index := 0; index < len(*eingabe); index++ {

		eingabeChar := string((*eingabe)[index]);

		if eingabeChar == "," || eingabeChar == "." {
			if only_once_komma == false {
				if eingabeChar == "," {
					fmt.Println( (*eingabe) );
					*eingabe = (*eingabe)[:index] + "." + (*eingabe)[(index+1):];
					only_once_komma = true;
					fmt.Println( (*eingabe) );
				} else {

					only_once_komma = true;

				}
			} else {

				return true;

			}
		} else if eingabeChar < "0" || eingabeChar > "9" {

			return true;

		}
	}

	return false;

}
