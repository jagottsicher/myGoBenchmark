package myConcat

func Concat(inputSlicOfString []string) string {

	rückgabeWert := inputSlicOfString[0]

	for _, v := range inputSlicOfString[1:] {
		rückgabeWert += " "
		rückgabeWert += v
	}

	return rückgabeWert
}
