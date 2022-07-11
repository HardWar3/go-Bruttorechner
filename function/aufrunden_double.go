package function

func Aufrunden_double( wo int64, was float64 ) float64 {
	var result float64;
	var vorgang_1 float64;
	var vorgang_2 float64;

	komma := []float64{ 1.0,10.0,100.0,1000.0,10000.0,100000.0,10000000.0 };
	if wo < 7 {
		vorgang_1 = float64(int64( ( was + 0.99 ) * komma[wo] ));
		vorgang_2 = ( vorgang_1 / komma[wo] );
		result = vorgang_2;
		return result;
	} else {
		return was;
	}
}
