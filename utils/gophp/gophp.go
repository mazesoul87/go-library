package gophp

import (
	"github.com/mazesoul87/go-library/utils/gophp/serialize"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// Serialize 序列
func Serialize(value any) ([]byte, error) {
	return serialize.Marshal(value)
}

// Unserialize 反序列
func Unserialize(data []byte) (any, error) {
	return serialize.UnMarshal(data)
}

func BaseConvert(number string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(number, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// ArrayColumn array_column()
func ArrayColumn(input map[string]map[string]any, columnKey string) []any {
	columns := make([]any, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}

// Strtr strtr()
//
// If the parameter length is 1, type is: map[string]string
// Strtr("baab", map[string]string{"ab": "01"}) will return "ba01"
// If the parameter length is 2, type is: string, string
// Strtr("baab", "ab", "01") will return "1001", a => 0; b => 1.
func Strtr(haystack string, params ...interface{}) string {
	ac := len(params)
	if ac == 1 {
		pairs := params[0].(map[string]string)
		length := len(pairs)
		if length == 0 {
			return haystack
		}
		oldnew := make([]string, length*2)
		for o, n := range pairs {
			if o == "" {
				return haystack
			}
			oldnew = append(oldnew, o, n)
		}
		return strings.NewReplacer(oldnew...).Replace(haystack)
	} else if ac == 2 {
		from := params[0].(string)
		to := params[1].(string)
		trlen, lt := len(from), len(to)
		if trlen > lt {
			trlen = lt
		}
		if trlen == 0 {
			return haystack
		}
		str := make([]uint8, len(haystack))
		var xlat [256]uint8
		var i int
		var j uint8
		if trlen == 1 {
			for i = 0; i < len(haystack); i++ {
				if haystack[i] == from[0] {
					str[i] = to[0]
				} else {
					str[i] = haystack[i]
				}
			}
			return string(str)
		}
		// trlen != 1
		for {
			xlat[j] = j
			if j++; j == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[from[i]] = to[i]
		}
		for i = 0; i < len(haystack); i++ {
			str[i] = xlat[haystack[i]]
		}
		return string(str)
	}
	return haystack
}

func Rtrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	return strings.TrimRight(str, characterMask[0])
}

func StrPad(input string, padLength int, padString string) string {
	output := padString
	for padLength > len(output) {
		output += output
	}
	if len(input) >= padLength {
		return input
	}
	return output[:padLength-len(input)] + input
}

// Empty checks if a variable is empty (nil, zero value, or empty string).
func Empty(value interface{}) bool {
	if value == nil {
		return true
	}

	// Use reflection to check if the value is the zero value for its type.
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return val.Len() == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	}

	// If the type is not recognized, consider it non-empty.
	return false
}

func Isset(value interface{}) bool {
	return value != nil
}
