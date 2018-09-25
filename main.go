package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)


func swap(A,B string)(string,string){
	return B,A
}
func Rot(N int, plain string, flag *bool) (string, error) {
	var (
		upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lowerCase  = "abcdefghijklmnopqrstuvwxyz"
		rotedUpper = upperCase[N:] + upperCase[:N]
		rotedLower = lowerCase[N:] + lowerCase[:N]
	)

	if *flag {
		swap(upperCase,rotedUpper)
		swap(lowerCase,rotedLower)
	}

	var roted bytes.Buffer
	for _, v := range plain {
		if unicode.IsLower(v) {
			index := strings.Index(lowerCase, string(v))
			roted.WriteString(string(rotedLower[index]))
		} else if unicode.IsUpper(v) {
			index := strings.Index(upperCase, string(v))
			roted.WriteString(string(rotedUpper[index]))
		} else {
			roted.WriteString(string(v))
		}

	}

	return roted.String(), nil
}

func main() {
	var (
		intOpt     = flag.Int("n", 13, "help message for i option")
		decodeFlag = flag.Bool("d", false, "help message for d option")
	)
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	target := scanner.Text()
	new, _ := Rot(*intOpt, target, decodeFlag)
	fmt.Println(new)

}
