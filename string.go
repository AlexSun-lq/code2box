package code2box

import "strings"

func replaceNT(src string) string {
	src = strings.ReplaceAll(src, "\n", "\\n")
	src = strings.ReplaceAll(src, "\t", "\\t")
	return src
}
func replaceN(src string, dst string) string {
	src = strings.ReplaceAll(src, "\n", dst)
	return src
}
func replace2S(src string, dst string) string {
	src = strings.ReplaceAll(src, "  ", dst)
	return src
}
func replaceT(src string) string {
	src = strings.ReplaceAll(src, "\t", "\\t")
	return src
}
func removeT(src string) string {
	src = strings.ReplaceAll(src, "\t", "")
	return src
}
