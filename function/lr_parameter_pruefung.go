package function

import (
	"fmt"
	"strconv"

	"go_bruttoRechner/eingabe"
	"go_bruttoRechner/user"
)

func Lr_parameter_pruefung( users []user.User, inhalt_vom_file []string ) int {

	//for index := 0; index < user_length; index++ {
	for index := range users {

		//user := users[index:(index+1)]
		user := &users[index]

		inhalt_index := index * 13;

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Id_nummer \n", inhalt_index );
			return 1;
		}
		if ( eingabe.Kontrolle_des_autofiles( &inhalt_vom_file[ inhalt_index ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Id_nummer \n", inhalt_index );
			return 1;
		}

		inhalt_i, err := strconv.ParseInt( inhalt_vom_file[ inhalt_index ], 10, 64 );
		if ( err != nil) {
			fmt.Println( err );
		}

		user.Id_nummer = inhalt_i;

		//fmt.Printf( "XXXxxxXXX ---- %d ---- \n", user.Id_nummer);

		//id checken hier
		var id_nummer string;

		// keine ahnung ob der substring so stehen bleibt
		id_nummer = strconv.FormatInt(user.Id_nummer, 10);

		firmen_nummer_cp, err := strconv.ParseInt( id_nummer[0:8], 10, 64 );
		if ( err != nil) {
			fmt.Println(err);
		}

		firmenpersonal_nummer_cp, err := strconv.ParseInt( id_nummer[8:len(id_nummer)], 10, 64 );
		if ( err != nil) {
			fmt.Println(err);
		}

		// das ist kein Datenbank dies soll eine Art DB sein
		// es dient nur symbolisch "datenbank"
		datenbank_check := Datenbank( firmen_nummer_cp, firmenpersonal_nummer_cp );

		if ( datenbank_check == 1 ) {

			return 1;

		}

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+1 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Geburtsjahr \n", inhalt_index+1 );
			return 1;
		}
		if ( eingabe.Kontrolle_des_autofiles( &inhalt_vom_file[ inhalt_index+1 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Geburtsjahr \n", inhalt_index+1 );
			return 1;
		}

		inhalt_i, err = strconv.ParseInt( inhalt_vom_file[inhalt_index+1], 10, 64 );
		if err != nil {
			 fmt.Println(err);
		}

		user.Ajahr = inhalt_i;

		// das alter checken
		if 2019 - user.Ajahr > 64 {
			user.Alter1 = 1;
		} else {
			user.Alter1 = 0;
		}

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+2 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Berechungsart LZZ \n", inhalt_index+2 );
			return 1;
		}
		if ( eingabe.Kontrolle_des_autofiles( &inhalt_vom_file[ inhalt_index+2 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Berechungsart LZZ \n", inhalt_index+2 );
			return 1;
		}

		inhalt_i, err = strconv.ParseInt( inhalt_vom_file[inhalt_index+2], 10, 64 );
		if err != nil {
			fmt.Println(err);
		}

		user.Lzz = inhalt_i;

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+3 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Gehalt \n", inhalt_index+3 );
			return 1;
		}

		inhalt_f, err := strconv.ParseFloat( inhalt_vom_file[inhalt_index+3], 64);
		if err != nil {
			fmt.Println(err);
		}

		user.Re4 = inhalt_f;

		if user.Re4 < 0 {

			fmt.Printf( "ERROR %f : - Feld RE4 ist im Negativen bereich\n", user.Id_nummer );
			return 1;

		} else if user.Re4 > 0 {

			// wie soll abgerechnet werden LZZ

			if user.Lzz == 1 {
				// Jahr
				user.Re4 = inhalt_f * 100;
				user.Jre4 = user.Re4;
			} else if user.Lzz == 2 {
				// Monat
				user.Re4 = inhalt_f * 100;
				user.Jre4 = user.Re4 * 12;
			} else if user.Lzz == 3 {
				// Wochen
				user.Re4 = inhalt_f * 100;
				user.Jre4 = user.Re4 * 360 / 7.0;
			} else if user.Lzz == 4 {
				// Tag
				user.Re4 =  inhalt_f * 100;
				user.Jre4 = user.Re4 * 360;
			} else {
				// abbruch des Programms an dieser Stelle wegen falscher LZZ
				fmt.Printf( "ERROR %f : - Feld LZZ hat ein Problem\n", user.Id_nummer );
				return 1;
			}

		} else {

			fmt.Printf( "ERROR %f : - Feld RE4 hat ein Problem\n", user.Id_nummer );
			return 1;

		}

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+4 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Versorgungsbezüge \n", inhalt_index+4 );
			return 1;
		}

		inhalt_f, err = strconv.ParseFloat( inhalt_vom_file[inhalt_index+4], 64 );
		if err != nil {
			fmt.Println(err);
		}

		user.Vbez = inhalt_f;

		if user.Vbez < 0 {

			fmt.Printf( "ERROR %f : - Feld VBEZ ist im negativen Wertebereich\n", user.Id_nummer );
			return 1;

		}

		// vbez darf nicht größer re4 sein
		if user.Vbez >= user.Re4 {
			fmt.Printf( "ERROR %f : - der VBEZ ist größer RE4 - VBEZ muss kleiner RE4 sein\n", user.Id_nummer );
			return 1;
		}

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+5 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Steuerklasse \n", inhalt_index+5 );
			return 1;
		}
		if ( eingabe.Kontrolle_des_autofiles( &inhalt_vom_file[ inhalt_index+5 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Steuerklasse \n", inhalt_index+5 );
			return 1;
		}

		inhalt_i, err = strconv.ParseInt( inhalt_vom_file[inhalt_index+5], 10, 64 );
		if err != nil {
			fmt.Println(err);
		}

		user.Stkl = inhalt_i;

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+6 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Kinderfreibetrag \n", inhalt_index+6 );
			return 1;
		}

		inhalt_f, err = strconv.ParseFloat( inhalt_vom_file[inhalt_index+6], 64 );
		if err != nil {
			fmt.Println(err);
		}

		user.Zkf = inhalt_f;

		// steuerklasse darf nicht größer 6 oder kleiner 1 gehen und in stkl 2 muss zkf angegeben werden
		// und Kinderfreibeträge checken stkl 1 und stkl 2 wichtig
		if user.Stkl < 1 || user.Stkl > 6 {
			fmt.Printf( "ERROR %f : - es Stimmt etwas nicht mit dem FELD STKL\n", user.Id_nummer );
			return 1;
		} else if user.Stkl == 2 && user.Zkf < 0.5 {
			fmt.Printf( "ERROR %f : - STKL 2 müssen Kinder angegeben werden\n", user.Id_nummer );
			return 1;
		} else if user.Stkl == 1 && user.Zkf > 0 {
			fmt.Printf( "ERROR %f : - STKL 1 kann keine Kinder haben\n", user.Id_nummer );
			return 1;
		}

		// pvz abchecken
		if user.Zkf < 0.5 && user.Ajahr > 23 {
			user.Pvz = 1;
		}

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+7 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Religion \n", inhalt_index+7 );
			return 1;
		}
		if ( eingabe.Kontrolle_des_autofiles( &inhalt_vom_file[ inhalt_index+7 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Religion \n", inhalt_index+7 );
			return 1;
		}

		inhalt_i, err = strconv.ParseInt( inhalt_vom_file[inhalt_index+7], 10, 64 );
		if err != nil {
			fmt.Println(err);
		}

		user.R = inhalt_i;

		// religion
		if user.R < 0 || user.R > 1 {
			fmt.Printf( "ERROR %f : - Religion kann nur 0 oder 1 sein\n", user.Id_nummer );
			return 1;
		}

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+8 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Rentenversicherung \n", inhalt_index+8 );
			return 1;
		}
		if ( eingabe.Kontrolle_des_autofiles( &inhalt_vom_file[ inhalt_index+8 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Rentenversicherung \n", inhalt_index+8 );
			return 1;
		}

		inhalt_i, err = strconv.ParseInt( inhalt_vom_file[inhalt_index+8], 10, 64 );
		if err != nil {
			fmt.Println(err);
		}

		user.Krv = inhalt_i;

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+9 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Pflegeversicherung Sachen \n", inhalt_index+9 );
			return 1;
		}
		if ( eingabe.Kontrolle_des_autofiles( &inhalt_vom_file[ inhalt_index+9 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Pflegeversicherung Sachen \n", inhalt_index+9 );
			return 1;
		}

		inhalt_i, err = strconv.ParseInt( inhalt_vom_file[inhalt_index+9], 10, 64 );
		if err != nil {
			fmt.Println(err);
		}

		user.Pvs = inhalt_i;

		// rentenversicherung und pvs
		if user.Krv < 0 || user.Krv > 2 {
			fmt.Printf( "ERROR %f : - Rentenversicherung kann nur 0, 1 oder 2 sein\n", user.Id_nummer );
			return 1;
		} else if user.Krv == 1 {
			if user.Pvs < 0 || user.Pvs > 1 {
				fmt.Printf( "ERROR %f : - der PVS kann nur 0 oder 1 sein\n", user.Id_nummer );
				return 1;
			}
		} else {
			if user.Pvs != 0 {
				fmt.Printf( "ERROR %f : - wenn KRV 0 oder 2 ist muss PVS 0 sein\n", user.Id_nummer );
				return 1;
			}
		}

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+10 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Krankenversicherungs Art \n", inhalt_index+10 );
			return 1;
		}
		if ( eingabe.Kontrolle_des_autofiles( &inhalt_vom_file[ inhalt_index+10 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! Krankenversicherungs Art \n", inhalt_index+10 );
			return 1;
		}

		inhalt_i, err = strconv.ParseInt( inhalt_vom_file[inhalt_index+10], 10, 64 );
		if err != nil {
			fmt.Println(err);
		}

		user.Pkv = inhalt_i;

		if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+11 ] ) ) {
			fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! zusatzbeitrag Krankenkasse \n", inhalt_index+11 );
			return 1;
		}

		inhalt_f, err = strconv.ParseFloat( inhalt_vom_file[inhalt_index+11], 64 );
		if err != nil {
			fmt.Println(err);
		}

		user.Kvz = inhalt_f;

		// pkv Krankenversicherung 0 1 2 privat und privat mit schuss
		// pkv 1 2 braucht pkpv als feld !!!
		if user.Pkv < 0 || user.Pkv > 2 {
			fmt.Printf( "ERROR %f : - wenn PKV kann nur 0, 1 oder 2 sein\n", user.Id_nummer );
			return 1;
		} else if user.Pkv != 0 {
			if user.Kvz != 0 {
				fmt.Printf( "ERROR %f : - der KVZ muss 0 sein bei PKV 1 und 2\n" );
				return 1;
			}

			if ( eingabe.Kontrolle_der_eingabe( &inhalt_vom_file[ inhalt_index+12 ] ) ) {
				fmt.Printf( "Die Zeile %d beinhaltet ein Falsches Format! PrivateKrankenversicherungsBeitrag \n", inhalt_index+12 );
				return 1;
			}

			inhalt_f, err = strconv.ParseFloat( inhalt_vom_file[inhalt_index+12], 64 );
			if err != nil {
				fmt.Println(err);
			}

		} else {
			if user.Kvz == 0 {
				fmt.Println( "ERROR %f : - der KVZ muss gesetzt sein bei PKV 0\n", user.Id_nummer );
				return 1;
			}
		}

		fmt.Println(inhalt_vom_file[inhalt_index]);
		fmt.Println(inhalt_vom_file[inhalt_index+1]);
		fmt.Println(inhalt_vom_file[inhalt_index+2]);
		fmt.Println(inhalt_vom_file[inhalt_index+3]);
		fmt.Println(inhalt_vom_file[inhalt_index+4]);
		fmt.Println(inhalt_vom_file[inhalt_index+5]);
		fmt.Println(inhalt_vom_file[inhalt_index+6]);
		fmt.Println(inhalt_vom_file[inhalt_index+7]);
		fmt.Println(inhalt_vom_file[inhalt_index+8]);
		fmt.Println(inhalt_vom_file[inhalt_index+9]);
		fmt.Println(inhalt_vom_file[inhalt_index+10]);
		fmt.Println(inhalt_vom_file[inhalt_index+11]);
		fmt.Println(inhalt_vom_file[inhalt_index+12]);
		fmt.Println("--------------------------------------------");

		fmt.Println( user.Id_nummer );
		fmt.Println( user.Lzz );
		fmt.Println( user.Re4 );
		fmt.Println( user.Jre4 );
		fmt.Println( user.Stkl );
		fmt.Println( user.Zkf );
		fmt.Println("--------------------------------------------");

		//users = append( users, user );

	}

	return 0;

}
