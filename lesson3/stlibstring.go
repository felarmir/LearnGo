package lesson3

import (
	"fmt"
	"strings"
)

func StringTest() {
	fmt.Println("---( Learning string feataures }---")
	fmt.Println(
		//true
		strings.Contains("String", "ri"),
		// 1
		strings.Count("String", "t"),
		//true
		strings.HasPrefix("String", "St"),
		//true
		strings.HasSuffix("String", "ing"),
		//2
		strings.Index("String", "r"),
		// Denis | Andreev
		strings.Join([]string{"Denis", "Andreev"}, " | "),
		//DDDDDDD
		strings.Repeat("D", 7),
		//AADD
		strings.Replace("DDDD", "D", "A", 2),
		//[a b c d]
		strings.Split("a-b-c-d", "-"),
		//test
		strings.ToLower("TEST"),
		//TEST
		strings.ToUpper("test"),
	)

	arr := []byte("test")
	str := string([]byte{'D', 'e', 'n', 'i', 's'})
	fmt.Printf("String to []byte: %b | []byte to string:%s", arr, str)
	fmt.Println("---( END }---")
}
