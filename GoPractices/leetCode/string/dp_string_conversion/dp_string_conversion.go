package main

func main() {
	source := "bit"
	target := "dog"
	words := []string{"but", "put", "big", "pot", "pog", "dog", "lot"}
	w := make(map[string]bool)
	char := ""

	for i := 0; i < len(words); i++ {

		w[words[i]] = true
		for j := 0; j < len(words[i]); j++ {
			char = char + string(words[i][i])
		}
	}

	res := 0
	shortedPath(0, source, target, w, char, 0, &res)

}

func shortedPath(k int, source, target string, w map[string]bool, char string, c int, res *int) {

	if source == target {

		if *res > c {
			*res = c
		}
		return
	}
	for i := 0; i < len(char); i++ {
		if check(k, source, w, char[i]) {
			//temp := source[k]
			//source[k] = uint8(v)
			shortedPath(k+1, source, target, w, char, c+1, res)
			//source[k] = temp
		}
	}
}

func check(k int, source string, w map[string]bool, v byte) bool {

	if k < len(source) {
		/*temp := source[k]
		stringzsource[k] = v*/
		return w[source]
	}

	return false
}
