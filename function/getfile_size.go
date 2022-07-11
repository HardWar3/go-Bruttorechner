package function

import (

	"fmt"
	"os"
	"bufio"

)

func Getfile_size ( file_name string ) int {

	inhalt_index := 0;

	file, err := os.Open( file_name );
	if err != nil {

		fmt.Println( err );

	}
	defer file.Close();

	scanner := bufio.NewScanner(file);
	for scanner.Scan() {

		inhalt_index++;

	}

	if err := scanner.Err(); err != nil {

		fmt.Println( err );

	}

	return inhalt_index;

}
