package code2box

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Bytes2HexStr(data []byte) string {
	return fmt.Sprintf("%x", data)
}

func Bytes2HexStrs(data []byte, ox ...bool) []string {
	strs := make([]string, len(data))
	for k, v := range data {
		if len(ox) > 0 && ox[0] {
			strs[k] = fmt.Sprintf("0x%02x", v)
			continue
		}
		strs[k] = fmt.Sprintf("%02x", v)
	}
	return strs
}

type B2SOpt struct {
	RowNum    bool
	Translate bool
	OX        bool
	ColSize   int
	ColSplit  string
}

// []byte 格式化为hex字符串，格式可指定
func Bytes2HexStrWithOpt(data []byte, opt ...B2SOpt) string {
	def := B2SOpt{
		true,
		true,
		false,
		16,
		" ",
	}
	if len(opt) > 0 {
		def = opt[0]
	}
	strs := Bytes2HexStrs(data, def.OX)
	l := len(strs)
	if def.ColSize == 0 {
		def.ColSize = l
	}
	tmp := ""
	// number format string
	NFS := fmt.Sprintf(" %%%dd    ", GetDigitsNum(l))
	for i := 0; i < l; i += def.ColSize {

		if i+def.ColSize > l {
			if def.RowNum {
				tmp += fmt.Sprintf(NFS, l)
			}
			tmp += strings.Join(strs[i:], def.ColSplit)
			space := ""
			for ii := 0; ii < def.ColSize-(l-i); ii++ {
				if def.OX {
					space += "     "
					continue
				}
				space += "   "
			}
			if def.Translate {
				tmp += space + "    " + replaceNT(string(data[i:])) + "\n"
			}
			tmp += "\n"
			break
		}
		if def.RowNum {
			tmp += fmt.Sprintf(NFS, i+def.ColSize)
		}

		tmp += strings.Join(strs[i:i+def.ColSize], def.ColSplit)

		if def.Translate {
			tmp += "    " + replaceNT(string(data[i:i+def.ColSize]))
		}
		tmp += "\n"
	}
	return tmp
}

// 输入16、10进制字符串，返回[]byte
// 需要指定是10还是16进制,需要指定分割符号,16进制还要指定是否携带0x
func HexStr2Byte(src string, baseSize int, split string) ([]byte, error) {

	src = strings.ReplaceAll(src, "0x", "")
	items := strings.Split(src, split)
	dst := make([]byte, len(items))
	var err error
	switch baseSize {
	case 10:
		for k, v := range items {
			u, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return dst, err
			}
			dst[k] = byte(u)
		}

	case 16:
		for k, v := range items {
			u, err := strconv.ParseUint(v, 16, 64)
			if err != nil {
				return dst, err
			}
			dst[k] = byte(u)
		}
	default:
		err = errors.New("base size must be 10 or 16")
	}
	return dst, err
}

// 输入16、10进制字符串，格式化为hex字符串，格式可指定
type S2HexOpt struct {
	B2SOpt
	BaseSize int
	Split    string
}

func FormatHexStr(src string, sh S2HexOpt) (string, error) {
	src = removeT(src)
	src = replaceN(src, " ")
	src = replace2S(src, " ")
	src = replace2S(src, " ")
	b, err := HexStr2Byte(src, sh.BaseSize, sh.Split)
	if err != nil {

		return "", err
	}
	return Bytes2HexStrWithOpt(b, sh.B2SOpt), nil
}

func Str2HexStr(src string, opt ...B2SOpt) string {

	data := []byte(src)

	if len(opt) > 0 {
		return Bytes2HexStrWithOpt(data, opt[0])
	}
	return fmt.Sprintf("%x\n", data)
}
