package eingabe

import (
	"fmt"
)

func Eingabe() string {

	var test string;

	fmt.Scanf( "%s", &test );

	fmt.Printf( "--- %s --- \n", test );;

	return test;

}
