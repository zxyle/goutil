package str

import (
	"strings"
)

type PyStr string

func (p PyStr) Capitalize() {
	panic("no implement")
}

func (p PyStr) CaseFold() {
	panic("no implement")
}

func (p PyStr) Center() {
	panic("no implement")
}

func (p PyStr) Count() {
	panic("no implement")
}

func (p PyStr) Encode() {
	panic("no implement")
}

func (p PyStr) EndsWith(suffix string) bool {
	return strings.HasSuffix(p.ToStr(), suffix)
}

func (p PyStr) ExpandTabs() {
	panic("no implement")
}

func (p PyStr) Find() {
	panic("no implement")
}

func (p PyStr) Format() {
	panic("no implement")
}

func (p PyStr) FormatMap() {
	panic("no implement")
}

func (p PyStr) Index() {
	panic("no implement")
}

func (p PyStr) Isalnum() {
	panic("no implement")
}

func (p PyStr) IsAlpha() {
	panic("no implement")
}

func (p PyStr) IsAscii() {
	panic("no implement")
}

func (p PyStr) IsDecimal() {
	panic("no implement")
}

func (p PyStr) IsDigit() {
	panic("no implement")
}

func (p PyStr) IsIdentifier() {
	panic("no implement")
}

func (p PyStr) IsLower() {
	panic("no implement")
}

func (p PyStr) IsNumeric() {
	panic("no implement")
}

func (p PyStr) IsPrintable() {
	panic("no implement")
}

func (p PyStr) IsSpace() {
	panic("no implement")
}

func (p PyStr) IsTitle() {
	panic("no implement")
}

func (p PyStr) IsUpper() {
	panic("no implement")
}

func (p PyStr) Join(array []string) (result string) {
	for k, v := range array {
		// Last element
		if (k + 1) == len(array) {
			result += v
		} else {
			result += v + p.ToStr()
		}
	}
	return
}

func (p PyStr) LJust() {
	panic("no implement")
}

func (p PyStr) Lower() string {
	return strings.ToLower(p.ToStr())
}

func (p PyStr) LStrip() string {
	spaceNum := 0
	for i := 0; i <= len(p); i++ {
		if p[i] == ' ' {
			spaceNum++
		} else {
			break
		}
	}
	newString := p[spaceNum:]
	return newString.ToStr()
}

func (p PyStr) MakeTrans() {
	panic("no implement")
}

func (p PyStr) Partition() {
	panic("no implement")
}

func (p PyStr) Replace() {
	panic("no implement")
}

func (p PyStr) RFind() {
	panic("no implement")
}

func (p PyStr) RIndex() {
	panic("no implement")
}

func (p PyStr) RJust() {
	panic("no implement")
}

func (p PyStr) RPartition() {
	panic("no implement")
}

func (p PyStr) RSplit() {
	panic("no implement")
}

func (p PyStr) RStrip() string {
	spaceNum := 0
	for i := len(p) - 1; i >= 0; i-- { // 去除字符串尾部的所有空格
		if p[i] == ' ' {
			spaceNum++
		} else {
			break
		}
	}
	newString := p[:len(p)-spaceNum]
	return newString.ToStr()
}

func (p PyStr) Split() {
	panic("no implement")
}

func (p PyStr) SplitLines() {
	panic("no implement")
}

func (p PyStr) StartsWith(prefix string) bool {
	return strings.HasPrefix(p.ToStr(), prefix)
}

func (p PyStr) Strip() string {
	newString := p.LStrip()
	p2 := PyStr(newString)
	newString = p2.RStrip()
	return newString
}

func (p PyStr) SwapCase() {
	panic("no implement")
}

func (p PyStr) Title() {
	panic("no implement")
}

func (p PyStr) Translate() {
	panic("no implement")
}

func (p PyStr) Upper() string {
	return strings.ToUpper(p.ToStr())
}

func (p PyStr) ZFill(width int) string {
	width = width - len(p)
	for i := 0; i < width; i++ {
		if len(p) <= width {
			p = "0" + p
		}
	}
	return p.ToStr()
}

func (p PyStr) ToStr() string {
	return string(p)
}
