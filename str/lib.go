package str

import "strings"

func Capitalize(s string) {
	return

}

func CaseFold() {

}

func Center() {

}

func Count() {

}

func Encode() {

}

func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func ExpandTabs() {

}

func Find() {

}

func Format() {

}

func FormatMap() {

}

func Index() {

}

func Isalnum() {

}

func IsAlpha() {

}

func IsAscii() {

}

func IsDecimal() {

}

func IsDigit() {

}

func IsIdentifier() {

}

func IsLower() {

}

func IsNumeric() {

}

func IsPrintable() {

}

func IsSpace() {

}

func IsTitle() {

}

func IsUpper() {

}

func Join() {

}

func LJust() {

}

func Lower(s string) string {
	return strings.ToLower(s)
}

func LStrip() {

}

func MakeTrans() {

}

func Partition() {

}

func Replace() {

}
func RFind() {

}

func RIndex() {

}

func RJust() {

}

func RPartition() {

}

func RSplit() {

}

func RStrip() {

}

func Split() {

}

func SplitLines() {

}

func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func Strip(s string) {
}

func SwapCase() {

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
