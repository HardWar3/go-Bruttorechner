package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go_bruttoRechner/eingabe"
	"go_bruttoRechner/function"
	"go_bruttoRechner/kirchensteuer"
	"go_bruttoRechner/lohnsteuer2019"
	"go_bruttoRechner/lohnsteuer2020"
	"go_bruttoRechner/lohnsteuer2021"
	"go_bruttoRechner/lohnsteuer2022"
	"go_bruttoRechner/lohnsteuer2023"
	"go_bruttoRechner/lohnsteuer2024"
	"go_bruttoRechner/sozialversicherungen"
	"go_bruttoRechner/user"

	"github.com/jung-kurt/gofpdf"
)

var opt_manuell, opt_auto, opt_file int
var file_name, PGM, opt_lohn, opt_format string
var file_datei *os.File
var spooldir = "/tmp/"
var pdf *gofpdf.Fpdf

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] for bruttoRechner\n", PGM)
	fmt.Fprintf(os.Stderr, "Felder mit ! sind Pflichtfelder\n")
	fmt.Fprintf(os.Stderr, "Options:\n"+
		"\tm \t! Manueller Modus für Eingabe\n"+
		"\ta \t! Automatisierter Modus f ist dann pflicht\n"+
		"\tf \t  f Filename\n"+
		"\tS \t  spooldir ( default: '/tmp' )\n"+
		"\tlr \t! LohnRechnungsArt - lohn 20 \n"+
		"\tdf \t! DatenFormat - df csv/txt/xml/-pdf-\n")

}

func main() {

	argsWith := os.Args
	PGM = argsWith[0]

	if process_arguments(argsWith) == 1 {
		usage()
		os.Exit(1)
	}

	if process() == 1 {
		os.Exit(1)
	}

	os.Exit(0)

}

func process_arguments(args []string) int {

	if len(args) < 1 {
		return 1
	}

	index := 1

	for index < len(args) {

		var inhalt string = string(args[index])

		//andere möglichkeit noch möglich als so sondern mit zwei breaks in default
		switch inhalt {

		case "m":
			if opt_manuell == 0 {
				fmt.Println("Manueller Modus")
				opt_manuell = 1
			} else {
				fmt.Println("zweite mal auslösen von Manueller Modus")
				fmt.Println("Programm muss beenedet werden")
				return 1
			}
		case "a":
			if opt_auto == 0 {
				fmt.Println("Auto Modus")
				opt_auto = 1
			} else {
				fmt.Println("zweite mal auslösen von Auto Modus")
				fmt.Println("Programm muss beendet werden")
				return 1
			}
		case "f":
			if opt_file == 0 {
				fmt.Printf("Filename %s \n", args[index+1])
				file_name = args[index+1]
				index++
				opt_file = 1
			}
		case "lohn":
			if args[index+1] == "19" {
				opt_lohn = "19"
				index++
			} else if args[index+1] == "20" {
				opt_lohn = "20"
				index++
			} else if args[index+1] == "21" {
				opt_lohn = "21"
				index++
			} else if args[index+1] == "22" {
				opt_lohn = "22"
				index++
			} else if args[index+1] == "23" {
				opt_lohn = "23"
				index++
			} else if args[index+1] == "24" {
				opt_lohn = "24"
				index++
			} else {
				return 1
			}
		case "df":
			if opt_format == "" {
				opt_format = args[index+1]
				index++
			} else {
				return 1
			}
		case "S":
			if spooldir != "/tmp/" {
				spooldir = args[index+1]
				index++
			} else {
				return 1
			}
		default:
			fmt.Println("keine vorhandene Option")
			fmt.Printf("Ungültiger Parameter %s\n", inhalt)
			return 1
		}

		index++

	}

	if opt_manuell == 1 && opt_auto == 1 {

		fmt.Println("Es kann immer nur einer der beiden Modi aktiv sein (Manueller oder Auto Modi)")
		return 1

	} else if opt_manuell == 0 && opt_auto == 0 {

		fmt.Println("Es muss ein Parameter mitgegeben werden! m oder a")
		return 1

	} else if opt_lohn == "" {
		fmt.Println("Im Feld LR ist kein Wert zu finden LR ist ein Pflichtfeld!")
		return 1
	} else if opt_format == "" {
		fmt.Println("DAS IST EINE VIEW ANSICHT !!! (da kein df angegeben wurde)\nMit ctrl + c zum Beenden des Programms")
		fmt.Println()
		time.Sleep(10 * time.Second)
	}
	return 0

}

func process() int {

	if opt_auto == 1 {

		if opt_auto == 1 && file_name == "" {

			fmt.Println("Bitte geben Sie im Auto-Modus die Option -f Filename an")
			usage()
			return 1

		}

		file_users, err := os.Open(file_name)

		if function.ErrorOut(err, "Datei konnte NICHT geoeffnet werden.") == 1 {
			return 1
		}

		defer file_users.Close()

		file_size := function.Getfile_size(file_name)

		var inhalt_vom_file []string

		fmt.Printf("%d -erster-: \n", file_size)

		if file_size%13 != 0 {

			fmt.Println("Die Daten sind nicht im Format Programm wird Beendet")
			file_users.Close()
			return 1

		}

		file_index := 0
		scanner := bufio.NewScanner(file_users)

		fmt.Println("Datei konnte geoeffnet werden. ")
		for scanner.Scan() {

			inhalt := scanner.Text()
			inhalt_vom_file = append(inhalt_vom_file, inhalt)
			file_index++

		}

		err = scanner.Err()

		if function.ErrorOut(err, "Datei konnte NICHT geoeffnet werden.") == 1 {
			return 1
		}

		// Datei schliessen
		file_users.Close()

		user_length := (len(inhalt_vom_file)) / 13

		users := make([]user.User, user_length)

		fmt.Println(len(users))

		// file controlle
		var fehlercode int
		fehlercode = function.Lr_parameter_pruefung(users, inhalt_vom_file)

		if fehlercode == 1 {

			return 1

		}

		if opt_format != "pdf" {

			file_path := ""

			if opt_format == "txt" {
				file_path = file_path + spooldir + "/auto_modus.txt"
			} else if opt_format == "csv" {
				file_path = file_path + spooldir + "/auto_modus.csv"
			}

			fmt.Println(file_path)

			_, err = os.Stat(file_path)

			if os.IsNotExist(err) {
				if opt_format == "txt" || opt_format == "csv" {
					file_datei, err = os.Create(file_path)
				}
				if function.ErrorOut(err, "Fehler bei Create von Datei1") == 1 {
					return 1
				}
				defer file_datei.Close()
			} else {
				err = os.Remove(file_path)
				if function.ErrorOut(err, "Fehler bei Remove von Datei2") == 1 {
					return 1
				}
				if opt_format == "txt" || opt_format == "csv" {
					file_datei, err = os.Create(file_path)
				}
				if function.ErrorOut(err, "Fehler bei Create von Datei3") == 1 {
					return 1
				}
				defer file_datei.Close()
			}

			file_datei, err = os.OpenFile(file_path, os.O_RDWR, 0644)
			if function.ErrorOut(err, "Fehler bei Öffnen von Datei") == 1 {
				return 1
			}
			defer file_datei.Close()

		}

		for i := 0; i < len(users); i++ {
			user := users[i]

			if opt_lohn == "19" {
				lohnsteuer2019.Lohnsteuer2019(&user)
			} else if opt_lohn == "20" {
				lohnsteuer2020.Lohnsteuer2020(&user)
			} else if opt_lohn == "21" {
				lohnsteuer2021.Lohnsteuer2021(&user)
			} else if opt_lohn == "22" {
				lohnsteuer2022.Lohnsteuer2022(&user)
			} else if opt_lohn == "23" {
				lohnsteuer2023.Lohnsteuer2023(&user)
			} else if opt_lohn == "24" {
				lohnsteuer2024.Lohnsteuer2024(&user)
			}

			if user.R == 1 {
				kirchensteuer.Kirchensteuer(&user)
			}
			sozialversicherungen.Sozialversicherungen(&user)
			function.Zusammenrechnen(&user)

			if opt_format == "txt" {

				file_write := bufio.NewWriter(file_datei)

				// schreibvorgang mit fmt.Fprintf( "TEST" );
				fmt.Fprintf(file_write, " >>-->>-- - %.d - ID Nummer --<<--<<\n", user.Id_nummer)
				fmt.Fprintf(file_write, " >>-->>-- - %d - GB Jahr --<<--<<\n", user.Ajahr)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - LSTLZZ --<<--<<\n", user.Lstlzz)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - SOLI --<<--<<\n", user.Solzlzz)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Kirche --<<--<<\n", user.Kst)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Renten --<<--<<\n", user.Rv)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Arlose --<<--<<\n", user.Av)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Kranken --<<--<<\n", user.Kv)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Pflege AN --<<--<<\n", user.Pv_an)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Pflege AG --<<--<<\n", user.Pv_ag)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Gesamt_ST --<<--<<\n", user.Gesamt_st)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - sum_so_an --<<--<<\n", user.Sum_sozialv_an)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - sum_so_ag --<<--<<\n", user.Sum_sozialv_ag)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - netto_lohn --<<--<<\n", user.Netto_lohn)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - gesamt_belast --<<--<<\n", user.Gesamt_belast_ag)

				file_write.Flush()

			} else if opt_format == "csv" {

				file_write := csv.NewWriter(file_datei)

				inhalt_string := fmt.Sprintf("%d", user.Id_nummer)
				error_meldung := "DateiSchreibFehler bei der ID"
				inhalt_array := []string{"Firmane/Personal-Nummer : ", inhalt_string, " "}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%d", user.Ajahr)
				error_meldung = "DateiSchreibFehler bei dem Geburtsjahr"
				inhalt_array = []string{"Geburtsjahr : ", inhalt_string, " "}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Lstlzz/100)
				error_meldung = "DateiSchreibFehler bei der Lohnsteuer"
				inhalt_array = []string{"Lohnsteuer : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Solzlzz/100)
				error_meldung = "DateiSchreibFehler bei Soli"
				inhalt_array = []string{"Solidaritätszuschlag : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Kst/100)
				error_meldung = "DateiSchreibFehler bei Kirchensteuer"
				inhalt_array = []string{"Kirchensteuer : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Gesamt_st/100)
				error_meldung = "DateiSchreibFehler bei Summe der Steuern"
				inhalt_array = []string{"Summe der Steuern : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Rv/100)
				error_meldung = "DateiSchreibFehler bei Rentenversicherung AN"
				inhalt_array = []string{"9.3 % Rentenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Av/100)
				error_meldung = "DateiSchreibFehler bei Arbeitslosenversicherung AN"
				inhalt_array = []string{"1.25% Arbeitslosenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Kv/100)
				error_meldung = "DateiSchreibFehler bei Krankenversicherung AN"
				inhalt_array = []string{"7.8% Krankenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Pv_an/100)
				error_meldung = "DatieSchreibFehler bei Pflegeversicherung"
				inhalt_array = []string{"Pflegeversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Sum_sozialv_an/100)
				error_meldung = "DateiSchreibFehler bei Summe Sozialversicherung"
				inhalt_array = []string{"Summe Sozialversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Netto_lohn/100)
				error_meldung = "DateiSchreibFehler bei Netto Gehalt"
				inhalt_array = []string{"Netto Gehalt : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_array = []string{" ", " ", " "}
				error_meldung = "DateiSchreibFehler bei Empty"
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Rv/100)
				error_meldung = "DateiSchreibFehler bei Rentenversicherung AG"
				inhalt_array = []string{"9.3% Rentenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Av/100)
				error_meldung = "DateiSchreibFehler bei Arbeitslosenversicherung AG"
				inhalt_array = []string{"1.25% Arbeitslosenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Kv/100)
				error_meldung = "DateiSchreibFehler bei Krankenversicherung AG"
				inhalt_array = []string{"7.8% Krankenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Pv_ag/100)
				error_meldung = "DateiSchreibFehler bei Pflegeversicherung AG"
				inhalt_array = []string{"Pflegeversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Sum_sozialv_ag/100)
				error_meldung = "DateiSchreibFehler bei Summe Arbeitsgeberanteil"
				inhalt_array = []string{"Summe Arbeitgeberanteil : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user.Gesamt_belast_ag/100)
				error_meldung = "DateiSchreibFehler bei Gesamtbelastung Arbeitsgeber"
				inhalt_array = []string{"Gesamtbelastung Arbeitgeber : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_array = []string{" ", " ", " "}
				error_meldung = "DateiSchreibFehler bei Empty"
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				file_write.Flush()

			} else if opt_format == "pdf" {

				pdf := newReport(&user)

				if pdf.Err() {
					log.Fatalf("PDF Datei konnte nicht erstellt werden %s \n", pdf.Error())
					return 1
				}

				err := savePDF(pdf)
				if function.ErrorOut(err, "Speichervorgang Fehlgeschlagen PDF") == 1 {
					return 1
				}

			}
		}

	} else if opt_manuell == 1 {

		var user_daten user.User
		var _eingabe string
		var eingabeTest bool = true
		var inhalt_float float64
		var err error

		for {

			fmt.Printf("Bitte geben Sie Ihr Geburtsjahr ein : ---> ")
			_eingabe = eingabe.Eingabe()
			eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

			if !eingabeTest {
				break
			}

		}

		inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
		if function.ErrorOut(err, "Fehler bei eingabe Parsen in Float64 ( Geburtsjahr ) ") == 1 {
			return 1
		}

		user_daten.Ajahr = int64(function.Abrunden_double(0, inhalt_float))

		// ajahr muss noch weiter verarbeitet werden wegen ALTER1 etc
		user_daten.Alter1 = 0
		if 2019-user_daten.Ajahr > 64 {

			user_daten.Alter1 = 1

		}

		eingabeTest = true

		for {

			fmt.Printf("Soll in Monat( 1 ) oder in Jahren( 0 ) berechnet werden : ---> ")
			_eingabe = eingabe.Eingabe()
			eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

			if !eingabeTest {
				break
			}

		}

		inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
		if function.ErrorOut(err, "Fehler bei eingabe Parsen in Float64 ( LZZ )") == 1 {
			return 1
		}

		eingabeTest = true
		if int64(inhalt_float) == 1 {

			user_daten.Lzz = 2

			// Monat
			for {

				fmt.Printf("Geben Sie nun ihr BruttoLohn im Monat an : ---> ")
				_eingabe = eingabe.Eingabe()
				eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

				if !eingabeTest {
					break
				}

			}

			inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
			if function.ErrorOut(err, "Fehler bei eingabe Parsen in FLoat64 ( BruttoLohn )") == 1 {
				return 1
			}

			user_daten.Re4 = function.Abrunden_double(2, inhalt_float) * 100
			user_daten.Jre4 = function.Abrunden_double(2, (user_daten.Re4 * 12))

		} else if int64(inhalt_float) == 0 {

			user_daten.Lzz = 1

			// Jahr
			for {

				fmt.Printf("Geben Sie nun ihr BruttoLohn im Jahr an : ---> ")
				_eingabe = eingabe.Eingabe()
				eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

				if !eingabeTest {
					break
				}

			}

			inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
			if function.ErrorOut(err, "Fehler bei eingabe Parsen in Float64 ( BruttoLohn ) ") == 1 {
				return 1
			}

			user_daten.Re4 = function.Abrunden_double(2, inhalt_float) * 100
			user_daten.Jre4 = user_daten.Re4

		} else {

			fmt.Println("Im Manuellen Modus gibt es nur Monat oder Jahr deswegen muss das Programm Beendet werden")
			return 1

		}

		// wenn was tun dann hier für RE4

		eingabeTest = true
		for {

			fmt.Printf("davon Versorgungsbezüge : ---> ")
			_eingabe = eingabe.Eingabe()
			eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

			inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
			if function.ErrorOut(err, "Fehler bei eingabe Parsen in Float64 ( Versorgungsbezüge )") == 1 {
				return 1
			}

			if !(inhalt_float*100 >= user_daten.Re4) {
				break
			}

		}

		user_daten.Vbez = function.Abrunden_double(2, (inhalt_float * 100))

		// wenn was tun dann hier für VBEZ

		eingabeTest = true
		for {

			fmt.Printf("Welche Steuerklassen haben Sie : ---> ")
			_eingabe = eingabe.Eingabe()
			eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

			inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
			if function.ErrorOut(err, "Fehler bei eingabe Parsen in Float64 ( Steuerklasse )") == 1 {
				return 1
			}

			if !(inhalt_float < 1 || inhalt_float > 6) {
				break
			}

		}

		user_daten.Stkl = int64(inhalt_float)

		// wenn was tun dann hier für STKL

		if user_daten.Stkl != 1 {

			for {

				fmt.Printf("Wie viele Kinder haben Sie : ---> ")
				_eingabe := eingabe.Eingabe()
				eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

				inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
				if function.ErrorOut(err, "Fehler bei eingabe Parsen in Float64 ( KinderFreibetrag )") == 1 {
					return 1
				}

				if !(user_daten.Stkl == 2 && inhalt_float <= 0) {
					break
				} else if user_daten.Stkl != 2 {
					break
				}

			}

			user_daten.Zkf = function.Abrunden_double(1, inhalt_float)

		}

		// wenn was tun dann hier für ZKF

		user_daten.Pvz = 0

		if user_daten.Zkf == 0.0 && 2019-user_daten.Ajahr > 23 {

			user_daten.Pvz = 1

		}

		for {

			fmt.Printf("Sind Sie in der Kirche ( Ja = 1 / Nein = 0 ) : ---> ")
			_eingabe = eingabe.Eingabe()
			eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

			if !eingabeTest {
				break
			}

		}

		inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
		if function.ErrorOut(err, "Fehler bei Parsen in Float64 ( Kirche )") == 1 {
			return 1
		}

		user_daten.R = int64(function.Abrunden_double(0, inhalt_float))

		// wenn was run dann hier für R

		for {

			fmt.Printf("Welche gesetzliche Rentenversicherung sind Sie (für West 0 / für Ost 1 / für alles andere 2) : ---> ")
			_eingabe = eingabe.Eingabe()
			eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

			inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
			if function.ErrorOut(err, "Fehler bei Parsen in Float64 ( Rentenversicherung )") == 1 {
				return 1
			}

			if inhalt_float == 0 || inhalt_float == 1 || inhalt_float == 2 {
				break
			}

		}

		user_daten.Krv = int64(function.Abrunden_double(0, inhalt_float))

		// wenn was tun dann hier für KRV

		if user_daten.Krv == 1 {

			for {

				fmt.Printf("Wohnen Sie in Sachen ( für Ja 1 / für Nein 0 ) : ---> ")
				_eingabe = eingabe.Eingabe()
				eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

				inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
				if function.ErrorOut(err, "Fehler bei Parsen in Float64 ( Sachen )") == 1 {
					return 1
				}

				if inhalt_float == 0 || inhalt_float == 1 {
					break
				}

			}

			user_daten.Pvs = int64(function.Abrunden_double(0, inhalt_float))

		}

		for {

			fmt.Printf("Sind Sie gesetzlich Krankenversichert ( 0 ) / Privat ohne zuschuss ( 1 ) / Privat mit zuschuss ( 2 ) : ---> ")
			_eingabe = eingabe.Eingabe()
			eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

			inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
			if function.ErrorOut(err, "Fehler bei Parsen in Float64 ( Krankenversicherung )") == 1 {
				return 1
			}

			if inhalt_float == 0 || inhalt_float == 1 || inhalt_float == 2 {
				break
			}

		}

		user_daten.Pkv = int64(function.Abrunden_double(0, inhalt_float))

		// wenn was tun dann hier für PKV

		if user_daten.Pkv == 0 {

			for {

				fmt.Printf("Wie hoch ist der Zusatzbeitragssatz der gesetzlichen Krankenversicherung : ---> ")
				_eingabe = eingabe.Eingabe()
				eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

				if !eingabeTest {
					break
				}

			}

			inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
			if function.ErrorOut(err, "Fehler bei Parsen in Float64 ( Zbsatz Krankenversicherung )") == 1 {
				return 1
			}

			user_daten.Kvz = inhalt_float

		} else if user_daten.Pkv == 1 || user_daten.Pkv == 2 {

			for {

				fmt.Printf("Wie hoch ist der Privat Versicherungsbeitrag : ---> ")
				_eingabe = eingabe.Eingabe()
				eingabeTest = eingabe.Kontrolle_der_eingabe(&_eingabe)

				if !eingabeTest {
					break
				}

			}

			inhalt_float, err = strconv.ParseFloat(_eingabe, 64)
			if function.ErrorOut(err, "Fehler bei Parsen in Float64 ( Private Versicherungsbeirtag )") == 1 {
				return 1
			}

			user_daten.Pkpv = inhalt_float * 100

		}

		if opt_lohn == "19" {
			lohnsteuer2019.Lohnsteuer2019(&user_daten)
		} else if opt_lohn == "20" {
			lohnsteuer2020.Lohnsteuer2020(&user_daten)
		} else if opt_lohn == "21" {
			lohnsteuer2021.Lohnsteuer2021(&user_daten)
		} else if opt_lohn == "22" {
			lohnsteuer2022.Lohnsteuer2022(&user_daten)
		} else if opt_lohn == "23" {
			lohnsteuer2023.Lohnsteuer2023(&user_daten)
		} else if opt_lohn == "24" {
			lohnsteuer2024.Lohnsteuer2024(&user_daten)
		}

		if user_daten.R == 1 {
			kirchensteuer.Kirchensteuer(&user_daten)
		}
		sozialversicherungen.Sozialversicherungen(&user_daten)
		function.Zusammenrechnen(&user_daten)

		file_path := ""

		if opt_format == "txt" {
			file_path = file_path + spooldir + "/manu_modus.txt"
		} else if opt_format == "csv" {
			file_path = file_path + spooldir + "/manu_modus.csv"
		}

		if opt_format != "" {

			fmt.Println(file_path)

			_, err = os.Stat(file_path)

			if os.IsNotExist(err) {
				if opt_format == "txt" || opt_format == "csv" {
					file_datei, err = os.Create(file_path)
				}
				if function.ErrorOut(err, "Fehler bei Create von Datei4") == 1 {
					return 1
				}
				defer file_datei.Close()
			} else {
				err = os.Remove(file_path)
				if function.ErrorOut(err, "Fehler bei Remove von Date5") == 1 {
					return 1
				}
				if opt_format == "txt" || opt_format == "csv" {
					file_datei, err = os.Create(file_path)
				}
				if function.ErrorOut(err, "Fehler bei Create von Datei6") == 1 {
					return 1
				}
				defer file_datei.Close()
			}

			file_datei, err = os.OpenFile(file_path, os.O_RDWR, 0644)
			if function.ErrorOut(err, "Fehler bei Öffnen von Datei") == 1 {
				return 1
			}
			defer file_datei.Close()

			if opt_format == "txt" {

				file_write := bufio.NewWriter(file_datei)

				// schreibvorgang mit fmt.Fprintf( "TEST" );
				fmt.Fprintf(file_write, " >>-->>-- - %d - ID Nummer --<<--<<\n", user_daten.Id_nummer)
				fmt.Fprintf(file_write, " >>-->>-- - %d - GB Jahr --<<--<<\n", user_daten.Ajahr)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - LSTLZZ --<<--<<\n", user_daten.Lstlzz)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - SOLI --<<--<<\n", user_daten.Solzlzz)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Kirche --<<--<<\n", user_daten.Kst)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Renten --<<--<<\n", user_daten.Rv)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Arlose --<<--<<\n", user_daten.Av)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Kranken --<<--<<\n", user_daten.Kv)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Pflege AN --<<--<<\n", user_daten.Pv_an)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Pflege AG --<<--<<\n", user_daten.Pv_ag)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - Gesamt_ST --<<--<<\n", user_daten.Gesamt_st)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - sum_so_an --<<--<<\n", user_daten.Sum_sozialv_an)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - sum_so_ag --<<--<<\n", user_daten.Sum_sozialv_ag)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - netto_lohn --<<--<<\n", user_daten.Netto_lohn)
				fmt.Fprintf(file_write, " >>-->>-- - %.2f - gesamt_belast --<<--<<\n", user_daten.Gesamt_belast_ag)
				file_write.Flush()

			} else if opt_format == "csv" {

				file_write := csv.NewWriter(file_datei)

				inhalt_string := fmt.Sprintf("%d", user_daten.Id_nummer)
				error_meldung := "DateiSchreibFehler bei der ID"
				inhalt_array := []string{"Firmane/Personal-Nummer : ", inhalt_string, " "}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%d", user_daten.Ajahr)
				error_meldung = "DateiSchreibFehler bei dem Geburtsjahr"
				inhalt_array = []string{"Geburtsjahr : ", inhalt_string, " "}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Lstlzz/100)
				error_meldung = "DateiSchreibFehler bei der Lohnsteuer"
				inhalt_array = []string{"Lohnsteuer : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Solzlzz/100)
				error_meldung = "DateiSchreibFehler bei Soli"
				inhalt_array = []string{"Solidaritätszuschlag : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Kst/100)
				error_meldung = "DateiSchreibFehler bei Kirchensteuer"
				inhalt_array = []string{"Kirchensteuer : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Gesamt_st/100)
				error_meldung = "DateiSchreibFehler bei Summe der Steuern"
				inhalt_array = []string{"Summe der Steuern : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Rv/100)
				error_meldung = "DateiSchreibFehler bei Rentenversicherung AN"
				inhalt_array = []string{"9.3 % Rentenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Av/100)
				error_meldung = "DateiSchreibFehler bei Arbeitslosenversicherung AN"
				inhalt_array = []string{"1.25% Arbeitslosenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Kv/100)
				error_meldung = "DateiSchreibFehler bei Krankenversicherung AN"
				inhalt_array = []string{"7.8% Krankenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Pv_an/100)
				error_meldung = "DatieSchreibFehler bei Pflegeversicherung"
				inhalt_array = []string{"Pflegeversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Sum_sozialv_an/100)
				error_meldung = "DateiSchreibFehler bei Summe Sozialversicherung"
				inhalt_array = []string{"Summe Sozialversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Netto_lohn/100)
				error_meldung = "DateiSchreibFehler bei Netto Gehalt"
				inhalt_array = []string{"Netto Gehalt : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_array = []string{" ", " ", " "}
				error_meldung = "DateiSchreibFehler bei Empty"
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Rv/100)
				error_meldung = "DateiSchreibFehler bei Rentenversicherung AG"
				inhalt_array = []string{"9.3% Rentenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Av/100)
				error_meldung = "DateiSchreibFehler bei Arbeitslosenversicherung AG"
				inhalt_array = []string{"1.25% Arbeitslosenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Kv/100)
				error_meldung = "DateiSchreibFehler bei Krankenversicherung AG"
				inhalt_array = []string{"7.8% Krankenversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Pv_ag/100)
				error_meldung = "DateiSchreibFehler bei Pflegeversicherung AG"
				inhalt_array = []string{"Pflegeversicherung : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Sum_sozialv_ag/100)
				error_meldung = "DateiSchreibFehler bei Summe Arbeitsgeberanteil"
				inhalt_array = []string{"Summe Arbeitgeberanteil : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_string = fmt.Sprintf("%.2f", user_daten.Gesamt_belast_ag/100)
				error_meldung = "DateiSchreibFehler bei Gesamtbelastung Arbeitsgeber"
				inhalt_array = []string{"Gesamtbelastung Arbeitgeber : ", inhalt_string, "Euro"}
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				inhalt_array = []string{" ", " ", " "}
				error_meldung = "DateiSchreibFehler bei Empty"
				if function.WriteCsvDatei(inhalt_array, error_meldung, file_write) == 1 {
					return 1
				}

				file_write.Flush()

			}
		} else {

			fmt.Printf(" >>-->>-- - %d - ID Nummer --<<--<<\n", user_daten.Id_nummer)
			fmt.Printf(" >>-->>-- - %d - GB Jahr --<<--<<\n", user_daten.Ajahr)
			fmt.Printf(" >>-->>-- - %.2f - LSTLZZ --<<--<<\n", user_daten.Lstlzz/100.00)
			fmt.Printf(" >>-->>-- - %.2f - SOLI --<<--<<\n", user_daten.Solzlzz/100.00)
			fmt.Printf(" >>-->>-- - %.2f - Kirche --<<--<<\n", user_daten.Kst/100.00)
			fmt.Printf(" >>-->>-- - %.2f - Renten --<<--<<\n", user_daten.Rv/100.00)
			fmt.Printf(" >>-->>-- - %.2f - Arlose --<<--<<\n", user_daten.Av/100.00)
			fmt.Printf(" >>-->>-- - %.2f - Kranken --<<--<<\n", user_daten.Kv/100.00)
			fmt.Printf(" >>-->>-- - %.2f - Pflege AN --<<--<<\n", user_daten.Pv_an/100.00)
			fmt.Printf(" >>-->>-- - %.2f - Pflege AG --<<--<<\n", user_daten.Pv_ag/100.00)
			fmt.Printf(" >>-->>-- - %.2f - Gesamt_ST --<<--<<\n", user_daten.Gesamt_st/100.00)
			fmt.Printf(" >>-->>-- - %.2f - sum_so_an --<<--<<\n", user_daten.Sum_sozialv_an/100.00)
			fmt.Printf(" >>-->>-- - %.2f - sum_so_ag --<<--<<\n", user_daten.Sum_sozialv_ag/100.00)
			fmt.Printf(" >>-->>-- - %.2f - netto_lohn --<<--<<\n", user_daten.Netto_lohn/100.00)
			fmt.Printf(" >>-->>-- - %.2f - gesamt_belast --<<--<<\n", user_daten.Gesamt_belast_ag/100.00)

		}

	} else {
		return 1
	}

	return 0

}

func newReport(user *user.User) *gofpdf.Fpdf {

	if pdf == nil {

		pdf = gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

	}

	pdf = pdfDesign(pdf)

	pdf = pdfWrite(pdf, user)

	return pdf
}

func pdfDesign(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {

	pdf.SetTextColor(139, 0, 139)

	// Header

	pdf.SetLeftMargin(15)
	pdf.SetFont("Arial", "", 6)
	pdf.SetFillColor(255, 0, 255)
	pdf.CellFormat(55, 10, "", "0", 0, "L", false, 0, "")

	pdf.Ln(3.9)
	pdf.SetY(8)
	pdf.SetFont("Arial", "B", 14)
	pdf.SetFillColor(0, 100, 255)
	pdf.SetLeftMargin(70)
	pdf.SetTopMargin(5)
	pdf.CellFormat(78, 6, "Lohn- und Gehaltsabrechnung", "0", 0, "L", false, 0, "")

	pdf.Ln(2)
	pdf.SetLeftMargin(148)
	pdf.SetFillColor(0, 255, 100)
	pdf.SetFont("Arial", "", 6)
	pdf.SetTopMargin(20)

	pdf.CellFormat(55, 4, "Sorgfaeltig aufbewahren! Gilt als Entgeltbescheinigung.", "0", 0, "L", false, 0, "")
	pdf.SetLeftMargin(15)

	pdf.SetTextColor(0, 0, 0)

	// Main Background

	pdf.Ln(4.5)
	pdf.SetY(14)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetDrawColor(139, 0, 139)
	pdf.CellFormat(188, 262.988888888, "", "1", 0, "L", true, 0, "")

	// Feld oben links

	pdf.Ln(0)
	pdf.SetY(16)
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(98, 10, "", "1", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(18)
	pdf.SetX(15.1)
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(97.8, 5, "FKN :", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	// Feld oben rechts

	pdf.Ln(0)
	pdf.SetY(14)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetLeftMargin(113)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(90, 72, "", "1", 0, "L", true, 0, "")

	pdf.Ln(0)
	pdf.SetY(14.1)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(18.6)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(23.1)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(27.6)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(32.1)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(36.6)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(41.1)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(45.6)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(50.1)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(54.6)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(59.1)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(63.6)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 9, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(72.6)
	pdf.SetX(113.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(89.8, 9, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	// Adress Feld

	pdf.Ln(0)
	pdf.SetY(28)
	pdf.SetX(15)
	pdf.SetFillColor(255, 255, 255)
	pdf.CellFormat(98, 58, "", "1", 0, "L", true, 0, "")

	pdf.Ln(0)
	pdf.SetY(28)
	pdf.SetX(15)
	pdf.SetFillColor(255, 255, 255)
	pdf.CellFormat(98, 18, "", "1", 0, "L", true, 0, "")

	// Mitte Main Feld

	pdf.Ln(0)
	pdf.SetY(88)
	pdf.SetX(15)
	pdf.SetFillColor(255, 255, 255)
	pdf.CellFormat(188, 162, "", "1", 0, "L", true, 0, "")

	pdf.Ln(0)
	pdf.SetY(89.6)
	pdf.SetX(15.1)
	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(94.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(98.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(103.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(107.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(112.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(116.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(121.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(125.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(130.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(134.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(139.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(143.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(148.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(152.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(157.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(161.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(166.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(170.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(175.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(179.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(184.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(188.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(193.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(197.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(202.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(206.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(211.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(215.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(220.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(224.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(229.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(233.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(238.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(242.6)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(247.1)
	pdf.SetX(15.1)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 2, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Line(22, 88, 22, 250)
	pdf.Line(35, 88, 35, 250)
	pdf.Line(96, 88, 96, 250)
	pdf.Line(107.5, 88, 107.5, 250)
	pdf.Line(117.5, 88, 117.5, 250)
	pdf.Line(140.5, 88, 140.5, 250)
	pdf.Line(160.5, 88, 160.5, 250)
	pdf.Line(180.5, 88, 180.5, 250)

	// Footer

	pdf.Ln(0)
	pdf.SetY(252)
	pdf.SetX(15)
	pdf.SetFillColor(255, 255, 255)
	pdf.CellFormat(188, 22.9, "", "1", 0, "L", true, 0, "")

	pdf.Ln(0)
	pdf.SetY(254.7)
	pdf.SetX(15.1)
	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(259.7)
	pdf.SetX(15.1)
	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(264.7)
	pdf.SetX(15.1)
	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(0)
	pdf.SetY(269.6888)
	pdf.SetX(15.1)
	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 240, 255)
	pdf.SetTextColor(139, 0, 139)
	pdf.SetCellMargin(2)
	pdf.CellFormat(187.8, 4.5, "", "0", 0, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	// alle Felder die in Farbe Magenta sind
	// ObenRechts die Felde
	pdf.SetTextColor(139, 0, 139)
	pdf.SetFont("Arial", "", 6)

	// erste reihe vom Feld ObenRechts

	pdf.Text(114.5, 15.8, "Abrechnung")
	pdf.Text(114.5, 18, "Monat  Jahr")

	pdf.Text(129.5, 17, "Url. Anspr.-VJ")

	pdf.Text(153.7, 17, "Url. Anspr.-LJ")

	pdf.Text(174.5, 17, "Url. gen.")

	pdf.Text(191, 17, "Url. Rest")

	// zweite reihe vom Feld obenrechts

	pdf.Text(114.5, 26.2, "Geburtsdatum")

	pdf.Text(129.5, 25.2, "St.")
	pdf.Text(129.5, 27.2, "Tg.")

	pdf.Text(135.5, 25.2, "Steuer-")
	pdf.Text(135.5, 27.2, "klasse")

	pdf.Text(142.5, 25.2, "Kinder-")
	pdf.Text(142.5, 27.2, "Freibetrage")

	pdf.Text(153.7, 25.2, "Rel.")
	// pdf.Text(142.5, 27.2, "Freibetrage")

	pdf.Text(159.2, 25.2, "Faktor")
	// pdf.Text(142.5, 27.2, "Freibetrage")

	// pdf.Text(159.2, 25.2, "Faktor")
	pdf.Text(174.5, 27.2, "jahrlich")

	pdf.Text(181.5, 25.2, "Freibetrag")
	// pdf.Text(174.5, 27.2, "jahrlich")

	// pdf.Text(181.5, 25.2, "Freibetrag")
	pdf.Text(191, 27.2, "monatlich")

	// dritte reihe vom Feld obenrechts

	pdf.Text(114.5, 35.2, "Eintrittsdatum")

	pdf.Text(129.5, 34.2, "Sv.")
	pdf.Text(129.5, 36.2, "Tg.")

	pdf.Text(135.5, 34.2, "Sozialversicherung")
	pdf.Text(135.5, 36.2, "KV")
	pdf.Text(140.1, 36.2, "RV")
	pdf.Text(144.8, 36.2, "AV")
	pdf.Text(149.4, 36.2, "PV")

	pdf.Text(153.7, 34.2, "Kind")
	pdf.Text(153.7, 36.2, "(PV)")

	pdf.Text(159.2, 34.2, "Bundes-")
	pdf.Text(159.2, 36.2, "land")

	pdf.Text(174.5, 34.2, "Um-")
	pdf.Text(174.5, 36.2, "lage")

	pdf.Text(182, 34.2, "Pers.Gr.")
	pdf.Text(182, 36.2, "Schl.")

	pdf.Text(191, 34.2, "Basis-")
	pdf.Text(191, 36.2, "absicherung")

	// vierte reihe vom Feld obenrechts

	pdf.Text(114.5, 44.2, "Austrittsdatum")

	pdf.Text(129.5, 43.2, "Kranken-")
	pdf.Text(129.5, 45.2, "kasse")

	pdf.Text(142.5, 43.2, "Name")
	// pdf.Text(129.5, 45.2, "")

	// fünfte reihe vom Feld obenrechts

	pdf.Text(114.5, 52.2, "Lohnsteuer-")
	pdf.Text(114.5, 54.2, "identifikationsnummer")

	pdf.Text(142.5, 52.2, "Sozial-")
	pdf.Text(142.5, 54.2, "versicherungsnummer")

	pdf.Text(182, 52.2, "Mehrfach-")
	pdf.Text(182, 54.2, "beschaeftigung")

	// six reihe vom Feld obenrechts

	// pdf.Text(114.5, 61.2, "Lohnsteuer-")
	pdf.Text(114.5, 63.2, "SL")

	pdf.Text(122.5, 61.2, "IBAN")
	pdf.Text(122.5, 63.2, "Bank")

	pdf.Text(174.5, 61.2, "BIC")
	// pdf.Text(114.5, 63.2, "SL")

	pdf.Text(9, 273, "Blatt:") // Wie viele Blätter

	pdf.SetFont("Arial", "", 8)

	pdf.Text(16, 258, "Aufrechnung: ") // Footer Aufrechnung

	// AufrechungsFelder

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(0, 0, 0)

	//pdf.Text(16,258,"Blatt:")
	pdf.Text(16, 263, "ST-Tage")
	pdf.Text(16, 268, "SV-Tage")
	pdf.Text(16, 273, "0001")

	pdf.Text(40, 258, "Gesamtbrutto:")
	pdf.Text(40, 263, "Steuerbrutto  :")
	pdf.Text(40, 268, "RV/AV Brutto:")
	pdf.Text(40, 273, "KV/PV Brutto:")

	pdf.Text(90, 258, "Sonst.Bezug   :")
	pdf.Text(90, 263, "Lohnsteuer     :")
	pdf.Text(90, 268, "Kirchensteuer :")
	pdf.Text(90, 273, "Solid.Zuschlag:")

	pdf.Text(145, 258, "SV-Anteil KV/AV/PV:")
	pdf.Text(145, 263, "SV-Anteil RV           :")
	pdf.Text(145, 268, "KV/PV Erstattung     :")
	pdf.Text(145, 273, "RV Erstattung          :")

	pdf.SetFont("Arial", "", 6)

	pdf.Text(18, 280, time.Now().Format("Mon Jan 2, 2006")) // Erstellungsdatum

	return pdf

}

func pdfWrite(pdf *gofpdf.Fpdf, user *user.User) *gofpdf.Fpdf {

	// ---------------------------------------------- WRITE - W

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", 10)

	// W erste reihe vom feld obenrechts
	pdf.Text(114.5, 22.2, time.Now().Format("Jan.06"))
	pdf.Text(129.5, 22.2, "--,--")
	pdf.Text(153.7, 22.2, "--,--")
	pdf.Text(174.5, 22.2, "--,--")
	pdf.Text(191, 22.2, "--,--")

	// W zweite reihe vom Feld obenrechts

	gb := fmt.Sprintf("%d", user.Ajahr)
	r := fmt.Sprintf("%d", user.R)
	f := fmt.Sprintf("%.0f", user.F)
	stkl := fmt.Sprintf("%d", user.Stkl)
	zkf := fmt.Sprintf("%.1f", user.Zkf)

	pdf.Text(114.5, 31.2, gb)
	pdf.Text(129.5, 31.2, "30")
	pdf.Text(135.5, 31.2, stkl)
	pdf.Text(142.5, 31.2, zkf)
	pdf.Text(153.7, 31.2, r)
	pdf.Text(159.2, 31.2, f)
	pdf.Text(174.5, 31.2, "")
	pdf.Text(181.5, 31.2, "")
	pdf.Text(191, 31.2, "")

	// W dritte reihe vom Feld obenrechts
	pdf.Text(114.5, 40.2, "--.--.--")
	pdf.Text(135.5, 40.2, "1")
	pdf.Text(140.1, 40.2, "1")
	pdf.Text(144.8, 40.2, "1")
	pdf.Text(149.4, 40.2, "1")
	pdf.Text(153.7, 40.2, "")
	pdf.Text(159.2, 40.2, "---")
	pdf.Text(174.5, 40.2, "-")
	pdf.Text(182, 40.2, "---")
	pdf.Text(191, 40.2, "")

	// W vierte reihe vom Feld obenrechts
	pdf.Text(114.5, 49.2, "--.--.--")
	pdf.Text(129.5, 49.2, "000")
	pdf.Text(142.5, 49.2, "--- UnK ---")

	// W fünfte reihe vom Feld obenrechts
	pdf.Text(114.5, 58.2, "----")
	pdf.Text(142.5, 58.2, "----")
	pdf.Text(182, 58.2, "----")

	// W six reihe vom Feld obenrechts
	pdf.Text(114.5, 67.2, "010")
	pdf.Text(122.5, 67.2, "DE54622200000005589044")
	pdf.Text(122.5, 71.7, "Bauspk Schwabisch Hall")

	pdf.Text(174.5, 67.2, "BSHHDE61XXX")
	// pdf.Text(114.5, 71.7, "31.05.88")

	pdf.SetFont("Arial", "", 10)

	// W FKN Feld
	pdf.Text(25, 21.4, "999990 Musterfirma Angebot")

	// W Adress Feld

	pdf.SetFont("Arial", "B", 10)

	pdf.Text(24, 45, "Personlich/Vertraulich!")

	pdf.SetFont("Arial", "", 8)

	pdf.Text(24, 33.5, "Musterfirma Angebot")
	pdf.Text(24, 36.5, "Mustertr.38")
	pdf.Text(24, 39.5, "80469 Musterstadt")

	pdf.Text(87.2, 33.5, "999990")
	pdf.Text(87.2, 36.5, "934B17")

	pdf.Text(75.2, 50.5, "Pers.Nr.")

	pdf.Text(87.2, 50.5, "000003")
	pdf.Text(87.2, 53.5, "11/19")
	pdf.Text(87.2, 56.5, "1000")

	pdf.SetFont("Arial", "", 10)

	pdf.Text(24, 52, "Frau")
	pdf.Text(24, 56, "Muster, Johanna")
	pdf.Text(24, 60, "Musterallee 12")
	pdf.Text(24, 64, "80809 Munchen")

	// W Mittel Feld

	pdf.Text(15.6, 93.2, "")
	pdf.Text(23, 93.2, "")
	pdf.Text(36, 93.2, ("Abrechnung                           " + time.Now().Format("Jan.06")))
	pdf.Text(98, 93.2, "")
	pdf.Text(109, 93.2, "")
	pdf.Text(118, 93.2, "")
	pdf.Text(141, 93.2, "")
	pdf.Text(161, 93.2, "")
	pdf.Text(181, 93.2, "")

	pdf.Text(15.6, 97.7, "LA")
	pdf.Text(23, 97.7, "Anzahl")
	pdf.Text(36, 97.7, "Bezeichnung")
	pdf.Text(98, 97.7, "EUR")
	pdf.Text(109, 97.7, "  %  ")
	pdf.Text(118, 97.7, "Kostenstelle")
	pdf.Text(141, 97.7, "SV.Btto.")
	pdf.Text(161, 97.7, "ST.Btto.")
	pdf.Text(181, 97.7, "Ges.Btto.")

	pdf.Text(15.6, 102.2, "-----")
	pdf.Text(23, 102.2, "----------")
	pdf.Text(36, 102.2, "--------------------------------------------------")
	pdf.Text(98, 102.2, "-------")
	pdf.Text(109, 102.2, "------")
	pdf.Text(118, 102.2, "------------------")
	pdf.Text(141, 102.2, "---------------")
	pdf.Text(161, 102.2, "---------------")
	pdf.Text(181, 102.2, "------------------")

	// 32 zeilen
	zeile := []float64{106.7, 111.2, 115.7, 120.2, 124.7, 129.2, 133.7, 138.2, 142.7, 147.2, 151.7, 156.2, 160.7, 165.2,
		169.7, 174.2, 178.7, 183.2, 187.7, 192.2, 196.7, 201.2, 205.7, 210.2, 214.7, 219.2, 223.7,
		228.2, 232.7, 237.2, 241.7, 246.2}
	// 9 splaten
	spalte := []float64{15.6, 23, 36, 98, 109, 118, 141, 161, 181}

	leerBalken := []string{"-----", "----------", "--------------------------------------------------", "-------",
		"------", "------------------", "---------------", "---------------",
		"------------------"}

	// hier muss die forschleife hin

	for index := 0; index < 17; index++ {

		for ind := 0; ind < 9; ind++ {

			if index == 1 || index == 5 || index == 7 || index == 12 || index == 14 || index == 16 {

				if ind != 0 && ind != 1 {

					pdf.Text(spalte[ind], zeile[index], leerBalken[ind])

				}

			} else if index == 0 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Gehalt")

				} else if ind == 5 && false {

					wert := fmt.Sprintf("%d", user.Id_nummer)
					pdf.Text(spalte[ind], zeile[index], wert)

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Re4 * 0.01))
					pdf.Text(spalte[ind], zeile[index], wert)

				}

			} else if index == 2 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Lohnsteuer")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Lstlzz * 0.01))
					pdf.Text(spalte[ind], zeile[index], wert)

				}

			} else if index == 3 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Solidaritaetszuschlag")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Solzlzz / 100))
					pdf.Text(spalte[ind], zeile[index], wert)

				}

			} else if index == 4 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Kirchensteuer")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Kst / 100))
					pdf.Text(spalte[ind], zeile[index], wert)

				}

			} else if index == 6 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Summe der Steuern")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Gesamt_st / 100))
					pdf.Text(spalte[ind], zeile[index], wert)

				}

			} else if index == 8 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Rentenversicherung")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Rv / 100))
					pdf.Text(spalte[ind], zeile[index], wert)

				}
			} else if index == 9 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Arbeitslosenversicherung")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Av / 100))
					pdf.Text(spalte[ind], zeile[index], wert)

				}

			} else if index == 10 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Krankenversicherung")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Kv / 100))
					pdf.Text(spalte[ind], zeile[index], wert)

				}
			} else if index == 11 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Pflegeversicherung")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Pv_an / 100))
					pdf.Text(spalte[ind], zeile[index], wert)

				}

			} else if index == 13 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Summe der Sozialversicherungen")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Sum_sozialv_an / 100))
					pdf.Text(spalte[ind], zeile[index], wert)

				}

			} else if index == 15 {

				if ind == 2 {

					pdf.Text(spalte[ind], zeile[index], "Netto Gehalt")

				} else if ind == 8 {

					wert := fmt.Sprintf("%.2f", (user.Netto_lohn / 100))
					pdf.Text(spalte[ind], zeile[index], wert)

				}

			}
		}
	}

	// W AufrechungsFelder

	pdf.SetFont("Arial", "", 10)

	//pdf.Text(16,258,"Blatt:")
	pdf.Text(32, 263, "330")
	pdf.Text(32, 268, "330")
	//pdf.Text(16,273,"")

	pdf.Text(72, 258, "18.320,00")
	pdf.Text(72, 263, "17.220,00")
	pdf.Text(72, 268, "17.220,00")
	pdf.Text(72, 273, "17.220,00")

	pdf.Text(127, 258, "00.000,00")
	pdf.Text(127, 263, "00.971,38")
	pdf.Text(127, 268, "00.026,31")
	pdf.Text(127, 273, "00.018,09")

	pdf.Text(185, 258, "00.000,00")
	pdf.Text(185, 263, "00.971,38")
	pdf.Text(185, 268, "00.026,31")
	pdf.Text(185, 273, "00.018,09")

	return pdf

}

func savePDF(pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose("report.pdf")
}
