package str

func Strip(s string) {
}

func LeftStrip(s string) {

}

func RightStrip(s string) {

}

// zero fill
func Zfill(text string, length int) string {
	newText := text
	length = length - len(text)
	for i := 0; i < length; i++ {
		if len(text) <= length {
			newText = "0" + newText
		}
	}
	return newText
}
