package main

import (
	"fmt"
	"unicode"
	"strings"
	"bytes"
	"bufio"
	"os"
	"flag"
)

const upperCase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowerCase string = "abcdefghijklmnopqrstuvwxyz"

var (
    intOpt  = flag.Int("n", 13, "help message for i option")
    strOpt  = flag.String("s", "default", "help message for s option")
)
func Rot(N int,plain string) ( string, error){
	 rotedUpper := upperCase[N:] + upperCase[:N]
	 rotedLower := lowerCase[N:] + lowerCase[:N]
	 var roted bytes.Buffer
	 for _,v := range plain {
		 if unicode.IsLower(v) {
			 index := strings.Index(lowerCase, string(v))
			 roted.WriteString(string(rotedLower[index]))
		 }else if unicode.IsUpper(v){
		 	index := strings.Index(upperCase,string(v))
		 	roted.WriteString(string(rotedUpper[index]))
		 }else{
		 	roted.WriteString(string(v))
		 }

	 }

	 return roted.String() , nil
}


func main(){

	var N int
	fmt.Println("input the string")
	//fmt.Scan(&str)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("input rotation number")
	fmt.Scan(&N)
	new, _  := Rot(N,scanner.Text())
	fmt.Println(new)

}

