package main

import(
   "fmt"
//   "unicode/utf8"
)

func main(){ 
	fmt.Println(string(rune(30d7)))
	
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) &&  s[:len(prefix)] == prefix
}

func HasSufix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Substring(s, substr string) bool {
	for i := 0; i< len(s) ;i++ {
		if HasPrefix(s[i:], substr){
			return true
		}
	}
	return false
}
