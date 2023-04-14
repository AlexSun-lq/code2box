package code2box

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestBytes2HexStr(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5, 6, 33, 44, 65, 234}
	fmt.Println(Bytes2HexStr(data))
}

func TestBytes2HexStrs(t *testing.T) {
	data := []byte{69,0,0,60,2,235,0,0,52,1,250,70,123,206,89,226,172,16,7,207,0,0,82,126,0,1,2,221,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,97,98,99,100,101,102,103,104,105}
	fmt.Println(Bytes2HexStrs(data))
}

func TestBytes2HexStrWithOpt(t *testing.T) {

	data := RandomBytes(64)
	fmt.Println(Bytes2HexStrWithOpt(data))

	data = RandomBytes(32)
	fmt.Println(Bytes2HexStrWithOpt(data, B2SOpt{
		RowNum:    true,
		Translate: false,
		OX:        true,
		ColSplit:  " ",
	}))
}

func TestHexStr2Byte(t *testing.T) {
	src := "1 2 3 44 123"
	b, err := HexStr2Byte(src, 10, " ")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(b)

	src = "1 02 03 44 123"
	b, err = HexStr2Byte(src, 16, " ")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(b)
	src = "191"
	b, err = HexStr2Byte(src, 10, " ")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(b)
}

func TestFormatHexStr(t *testing.T) {
	opt := S2HexOpt{
		B2SOpt: B2SOpt{
			RowNum:    true,
			Translate: false,
			OX:        true,
			ColSize:   16,
			ColSplit:  " ",
		},
		BaseSize: 16,
		Split:    ", ",
	}
	// src := `72 84 84 80 47 49 46 49 32 50 48 48 32 79 75 10 83 101 114 118 101 114 58 32 110 103 105 110 120 47 49 46 50 49 46 54 10 67 111 110 116 101 110 116 45 84 121 112 101 58 32 116 101 120 116 47 104 116 109 108 59 32 99 104 97 114 115 101 116 61 117 116 102 45 56 10 67 111 110 110 101 99 116 105 111 110 58 32 67 108 111 115 101 100 10 10 60 33 68 79 67 84 89 80 69 32 104 116 109 108 62 10 60 104 116 109 108 62 10 60 104 101 97 100 62 10 60 116 105 116 108 101 62 116 108 115 230 181 139 232 175 149 33 60 47 116 105 116 108 101 62 10 60 115 116 121 108 101 62 10 104 116 109 108 32 123 32 99 111 108 111 114 45 115 99 104 101 109 101 58 32 108 105 103 104 116 32 100 97 114 107 59 32 125 10 98 111 100 121 32 123 32 119 105 100 116 104 58 32 51 53 101 109 59 32 109 97 114 103 105 110 58 32 48 32 97 117 116 111 59 10 102 111 110 116 45 102 97 109 105 108 121 58 32 84 97 104 111 109 97 44 32 86 101 114 100 97 110 97 44 32 65 114 105 97 108 44 32 115 97 110 115 45 115 101 114 105 102 59 32 125 10 60 47 115 116 121 108 101 62 10 60 47 104 101 97 100 62 10 60 98 111 100 121 62 10 60 104 49 62 232 191 153 230 152 175 228 184 128 228 184 170 230 181 139 232 175 149 33 60 47 104 49 62 10 60 112 62 231 148 168 228 186 142 116 108 115 231 154 132 230 181 139 232 175 149 233 161 181 233 157 162 227 128 130 60 47 112 62 10 10 60 112 62 230 181 139 232 175 149 230 156 186 229 153 168 105 112 58 49 55 50 46 49 54 46 51 46 50 52 56 60 47 112 62 10 60 112 62 230 181 139 232 175 149 231 171 175 229 143 163 58 52 52 51 60 47 112 62 10 60 47 98 111 100 121 62 10 60 47 104 116 109 108 62 10 10`
	src :=`0x4, 0xa2, 0x50, 0x56, 0xd9, 0x5e, 0x7b, 0xa5, 0xa1, 0xdf, 0xff, 0x22, 0x6e, 0xe5, 0x20, 0x87, 0xee, 0xe0, 0x1a, 0x91, 0x31, 0x7d, 0x1f, 0x7b, 0x67, 0x53, 0x5f, 0x9e, 0xd1, 0xc2, 0x5f, 0x65, 0xf1, 0x37, 0x16, 0x21, 0x3e, 0xd1, 0xc9, 0xe0, 0x31, 0x52, 0x4d, 0x6e, 0xb1, 0x49, 0xd4, 0x1c, 0x53, 0xaa, 0x60, 0x47, 0xb1, 0x2e, 0x5a, 0xb0, 0x9, 0x75, 0xb1, 0x51, 0x1e, 0x2, 0x81, 0xfe, 0xbb`
	dst, err := FormatHexStr(src, opt)
	if err != nil {
		fmt.Println("err Str2Hex: ", err)
		return
	}
	fmt.Println(dst)
}

func TestStr2HexStr(t *testing.T) {
	src := `HTTP/1.1 200 OK\nServer: nginx/1.21.6\nContent-Type: text/html; charset=utf-8\nConnection: Closed\n\n<!DOCTYPE html>\n<html>\n<head>\n<title>tls测试!</title>\n<style>\nhtml { color-scheme: light dark; }\nbody { width: 35em; margin: 0 auto;\nfont-family: Tahoma, Verdana, Arial, sans-serif; }\n</style>\n</head>\n<body>\n<h1>这是一个测试!</h1>\n<p>用于tls的测试页面。</p>\n\n<p>测试机器ip:172.16.3.248</p>\n<p>测试端口:443</p>\n</body>\n</html>\n\n`
	fmt.Println(Str2HexStr(src, B2SOpt{
		RowNum:    true,
		Translate: true,
		OX:        false,
		ColSize:   16,
		ColSplit:  " ",
	}))
}

func TestXxx(t *testing.T) {

	sStr := `publicKey :  [4 162 247 109 246 163 85 162 174 215 203 132 76 247 41 76 98 242 73 57 205 107 22 29 140 126 113 78 59 132 173 180 5 240 219 77 180 133 229 105 111 70 60 120 220 18 233 62 216 38 47 0 193 201 226 93 142 226 20 210 248 158 225 155 148]
	ourPublicKey :  [4 126 33 110 192 142 77 57 255 58 247 0 173 105 176 208 87 184 91 148 9 141 75 75 230 190 6 179 37 228 58 165 223 7 138 233 45 231 74 254 130 123 119 40 143 148 201 130 235 245 69 241 27 255 196 119 207 12 111 118 186 112 155 200 183]
	preMasterSecret :  [18 236 232 103 98 209 79 208 157 55 46 53 91 63 219 100 62 241 74 63 181 232 203 43 58 212 223 197 132 189 112 39]
	hs.hello.random :  [35 105 38 41 230 234 226 76 30 95 208 170 110 221 232 165 126 63 157 3 100 100 187 113 204 129 25 107 88 75 206 168]
	hs.serverHello.random :  [18 243 81 15 136 195 158 103 128 40 31 88 159 225 221 203 36 94 247 140 85 7 112 219 125 253 61 247 44 249 114 91]
	masterSecret :  [170 92 205 127 61 47 185 105 250 145 74 200 18 84 166 8 90 105 169 166 113 222 224 84 194 242 123 93 255 51 138 93 45 111 71 32 130 76 67 96 244 202 167 33 82 87 2 178]
	clientMAC  :  []
	serverMAC  :  []
	clientKey  :  [97 183 69 100 137 11 137 195 118 33 102 69 41 246 158 114 119 147 164 104 11 96 178 110 89 107 190 154 41 221 116 56]
	serverKey  :  [56 211 59 255 35 176 196 133 228 109 48 224 7 144 94 6 38 152 38 141 190 73 57 212 5 31 53 179 189 90 192 198]
	clientIV  :  [46 240 241 202]`

	rows := strings.Split(sStr, "\n")
	for _, v := range rows {
		kv := strings.Split(v, ":")
		name := kv[0]
		val := strings.ReplaceAll(kv[1], "[", "")
		val = strings.ReplaceAll(val, "]", "")
		cols := strings.Split(val, " ")
		tmpB := []byte{}
		for _, vv := range cols {
			if vv != "" {
				tn, _ := strconv.Atoi(vv)
				tmpB = append(tmpB, byte(tn))
			}
		}
		fmt.Println(name)
		fmt.Println(Bytes2HexStrWithOpt(tmpB))
	}

}
