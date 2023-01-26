package function

import (
	"fmt"
);

func ErrorOut ( err error, msg string ) int {

	if err != nil {
		fmt.Println( msg );
		return 1;
	} else {
		return(0);
	}

}
