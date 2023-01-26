package eingabe

import (
	"strings"
)

func Ingabe( eingabe string ) ( string, bool ) {

	var only_once_komma bool = false;

	for index := 0; index < len(eingabe); index++ {

		eingabeChar := string((eingabe)[index]);

		if eingabeChar == "," || eingabeChar == "." {
			if only_once_komma == false {
				if eingabeChar == "," {

					eingabe = strings.Replace( eingabe, ",", ".", 1 );
					only_once_komma = true;

				} else {

					only_once_komma = true;

				}
			} else {

				return eingabe, true;

			}
		} else if eingabeChar < "0" || eingabeChar > "9" {

			return eingabe, true;

		}
	}

	return eingabe, false;

}
