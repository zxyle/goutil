package str

import "strings"

func Capitalize(s string) {
	return

}

func Casefold() {

}

func Center() {

}

func Count() {

}

func Encode() {

}

func Endswith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func Expandtabs() {

}

func Find() {

}

func Format() {

}

func Format_map() {

}

func Index() {

}

func Isalnum() {

}

func Isalpha() {

}

func Isascii() {

}

func Isdecimal() {

}

func Isdigit() {

}

func Isidentifier() {

}

func Islower() {

}

func Isnumeric() {

}

func Isprintable() {

}

func Isspace() {

}

func Istitle() {

}

func Isupper() {

}

func Join() {

}

func Ljust() {

}

func Lower(s string) string {
	return strings.ToLower(s)
}

func Lstrip() {

}

func Maketrans() {

}

func Partition() {

}

func Replace() {

}
func Rfind() {

}

func Rindex() {

}

func Rjust() {

}

func Rpartition() {

}

func Rsplit() {

}

func Rstrip() {

}

func Split() {

}

func Splitlines() {

}

func Startswith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func Strip(s string) {
}

func Swapcase() {

}

func Title() {

}

func Translate() {

}

func Upper(s string) string {
	return strings.ToUpper(s)
}

// zero fill
func Zfill(s string, width int) string {
	width = width - len(s)
	for i := 0; i < width; i++ {
		if len(s) <= width {
			s = "0" + s
		}
	}
	return s
}
