package function

import (
	"fmt"
)

// provisorische Datenbank dient nur als Beispiel

func Datenbank(firmen_nummer int64, firmenpersonal_nummer int64) int {

	datenbank_firma := [4]int64{12345678, 23456789, 10382457, 234542}
	datenbank_personalnummer := [6]int64{12345678, 1, 2, 3, 4, 5}

	index := 0

	for ; index < len(datenbank_firma); index++ {
		if datenbank_firma[index] == firmen_nummer && datenbank_personalnummer[0] == firmen_nummer {
			dex := 1
			for ; dex < len(datenbank_personalnummer); dex++ {
				if datenbank_personalnummer[dex] == firmenpersonal_nummer {
					return 0
				}
			}

			fmt.Printf("die Firma ist vorhanden aber der Mitarbeiter ist noch nicht angelegt worden, Personalnummer %ld \n", firmenpersonal_nummer)
			return 1
		}
	}

	fmt.Printf("die Firma ist in der Dankenbank nicht angelegt Firmennummer %ld - Programm wurde beendet - \n", firmen_nummer)
	return 1

}
