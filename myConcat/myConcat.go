package myConcat

func Concat(inputSlicOfString []string) string {

	rückgabeWert := inputSlicOfString[0]

	for _, v := range inputSlicOfString[1:] {
		rückgabeWert += " "
		rückgabeWert += v
	}

	return rückgabeWert
}

func Join(elems []string, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	b := make([]byte, n) 
	bp := copy(b, elems[0])
	for _, s := range elems[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}