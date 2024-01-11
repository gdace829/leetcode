package main

func minLength(s string) int {
	zhan := []byte{}
	for i := 0; i < len(s); i++ {
		zhan = append(zhan, s[i])
		m := len(zhan)
		if m >= 2 && (string(zhan[m-2:]) == "AB" || string(zhan[m-2:]) == "CD") {
			zhan = zhan[:m-2]
		}
	}
	return len(zhan)
}
func main() {

}
