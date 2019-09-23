package sum

func Getsum(p int) (int) {

	var sum int
	// para tales que 0 <= v <= p
	for v := 0; v <= p; v++ {
		// calcule la suma de los valores v
		// siendo v un número divisible por 3 o por 5
		if (v%3 == 0) || (v%5 == 0) {
			sum += v
		}
	}

	return sum
}
