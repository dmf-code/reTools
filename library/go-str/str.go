package str

import (
	"regexp"
	"strconv"
	"strings"
)


func ToUint(str string) (b uint) {
	a, _ := strconv.ParseUint(str, 10, 64)
	b = uint(a)
	return
}

func ToInt(str string) (b int) {
	a, _ := strconv.ParseInt(str, 10, 64)
	b = int(a)
	return
}

func ToBool(str string) (b bool) {
	b, _ = strconv.ParseBool(str)

	return
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	// 去除图片
	re, _ = regexp.Compile("\\<img[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func SubString(str, dot string, start, length int) string {
	rs := []rune(str)
	end := start + length
	if end > len(rs) {
		end = len(rs)
	}
	subString := string(rs[start:end])
	if dot != "" {
		subString += dot
	}
	return subString
}
