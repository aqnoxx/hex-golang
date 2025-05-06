package main

import (
	"errors"
	"fmt"
)

const hexTable string = "0123456789ABCDEF"

var ASCII [256]uint8 = [256]uint8{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

func EncodeToString(b []byte) string {
	result := make([]rune, len(b)*2)

	for i, v := range b {
		result[i*2+1] = rune(hexTable[v&0x0F])
		result[i*2] = rune(hexTable[(v >> 4)])
	}

	return string(result)
}

func DecodeString(s string) ([]byte, error) {

	result := make([]byte, len(s)/2)
	if len(s)%2 != 0 {
		return make([]byte, 0), errors.New("string is not a hex")
	}
	for i := 0; i < len(s); i += 2 {
		result[i/2] = ASCII[[]byte(string(s[i]))[0]]*16 + ASCII[[]byte(string(s[i+1]))[0]]
	}
	return result, nil
}

func main() {
	fmt.Println(EncodeToString([]byte("Hello")))
	fmt.Println(EncodeToString([]byte("123")))
	fmt.Println(EncodeToString([]byte{1, 2, 3}))
	fmt.Println(EncodeToString([]uint8{'a', 'q', 'n', 'o', 'x'}))

	fmt.Println(DecodeString(EncodeToString(([]byte("Hello")))))
	fmt.Println(DecodeString(EncodeToString(([]byte("123")))))
	fmt.Println(DecodeString(EncodeToString(([]byte{1, 2, 3}))))
	fmt.Println(DecodeString(EncodeToString(([]uint8{'a', 'q', 'n', 'o', 'x'}))))
}
