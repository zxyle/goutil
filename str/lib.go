package str

import (
	"strings"
)

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

func LStrip(s string) string {
	spaceNum := 0
	for i := 0; i <= len(s); i++ {
		if s[i] == ' ' {
			spaceNum++
		} else {
			break
		}
	}
	newString := s[spaceNum:]
	return newString

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

func RStrip(s string) string {
	spaceNum := 0
	for i := len(s) - 1; i >= 0; i-- { // 去除字符串尾部的所有空格
		if s[i] == ' ' {
			spaceNum++
		} else {
			break
		}
	}
	newString := s[:len(s)-spaceNum]
	return newString
}

func Split() {

}

func SplitLines() {

}

func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func Strip(s string) string {
	newString := LStrip(s)
	newString = RStrip(newString)
	return newString
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
