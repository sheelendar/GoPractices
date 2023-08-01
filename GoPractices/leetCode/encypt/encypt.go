package main

import (
	"fmt"
)

/*
Test #4 for question "Decrypt Message"
<ACTUAL::4::> "a|cdefghijklmnopqrstuvwxyz"
Test Failed
<LOG:ESC:>Expected
<string>: a|cdefghijklmnopqrstuvwxyz
to equal
<string>: abcdefghijklmnopqrstuvwxyz

<ACTUAL::2::> "encyclopedi{"
Test Failed
<LOG:ESC:>Expected
<string>: encyclopedi{
to equal
<string>: encyclopedia
*/
func main() {
	fmt.Println(Decrypt("flgxswdliefy"))
	//encyclopedia
}

func Decrypt(word string) string {
	size := len(word)
	if size <= 0 {
		return word
	}
	result := ""
	for i := size - 1; i > 0; i-- {
		ch := int8(word[i]) - int8(word[i-1])
		for ch <= 96 {
			ch = ch + 26
		}

		result = fmt.Sprintf("%s%s", string(byte(ch)), result)

	}
	result = fmt.Sprintf("%s%s", string(word[0]-1), result)
	return result
}

//crime
