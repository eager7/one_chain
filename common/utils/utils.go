package utils

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

func JsonString(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(data)
}

func JsonObject(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}

func HexToString(h string) string {
	n, err := hex.DecodeString(HexFormat(h))
	if err != nil {
		fmt.Println(err)
	}
	return TrimZero(string(n))
}

func TrimZero(s string) string {
	str := make([]rune, 0, len(s))
	for _, v := range []rune(s) {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			continue
		}
		str = append(str, v)
	}
	return string(str)
}

func ErrorString(errs ...error) string {
	var s string
	for _, err := range errs {
		s += err.Error()
	}
	return s
}

func FromHex(s string) []byte {
	s = HexFormat(s)
	h, _ := hex.DecodeString(s)
	return h
}

func HexFormat(s string) string {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return s
}

func HexToUint64(hex string) (uint64, error) {
	n, err := strconv.ParseUint(HexFormat(hex), 16, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func HexToInt64(hex string) (int64, error) {
	n, err := strconv.ParseInt(HexFormat(hex), 16, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func HexToDebugNumber(hex string) uint64 {
	n, _ := HexToUint64(hex)
	return n
}

func HexToUint32(hex string) (uint32, error) {
	n, err := strconv.ParseUint(HexFormat(hex), 16, 32)
	if err != nil {
		return 0, err
	}
	return uint32(n), nil
}

func CRC32(s string) uint32 {
	ieee := crc32.NewIEEE()
	_, _ = ieee.Write([]byte(s))
	return ieee.Sum32()
}

func BigIntToHex(n *big.Int) string {
	return fmt.Sprintf("%x", n)
}

func BigIntFromHex(h string) *big.Int {
	h = HexFormat(h)
	b, _ := new(big.Int).SetString(h, 16)
	return b
}

func ByteToHex(data []byte) string {
	return hex.EncodeToString(data)
}

func HexToByte(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func LimitInput(input string) string {
	if len(input) > 15360 {
		return input[:15359]
	}
	return input
}

func MysqlFormat(s string) string {
	if !strings.ContainsAny(s, `;'\"&<>`) {
		return s
	}
	s = strings.Replace(s, `\`, `\\`, -1)
	s = strings.Replace(s, `'`, `\'`, -1)
	s = strings.Replace(s, `;`, `\;`, -1)
	s = strings.Replace(s, `"`, `\"`, -1)
	s = strings.Replace(s, `&`, `\&`, -1)
	s = strings.Replace(s, `<`, `\<`, -1)
	s = strings.Replace(s, `>`, `\>`, -1)
	return s
}

type Writer struct {
	Body bytes.Buffer
}

func (w *Writer) Write(b []byte) (int, error) {
	return w.Body.Write(b)
}
