package myConcat

func Concat(inputSlicOfString []string) string {

	returnString := inputSlicOfString[0]

	for _, v := range inputSlicOfString[1:] {
		returnString += " "
		returnString += v
	}

	return returnString
}
