package function

import (
	"encoding/csv"
);

func WriteCsvDatei( inhalt_array []string, error_meldung string, file_write *csv.Writer ) int {

	err := file_write.Write(inhalt_array);
	if ErrorOut( err, error_meldung ) == 1 {
		return 1;
	}
	return 0;
}
