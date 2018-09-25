package main

import (
	"fmt"
	"unicode"
	"strings"
	"bytes"
	"os"
	"flag"
	"bufio"
)

const upperCase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowerCase string = "abcdefghijklmnopqrstuvwxyz"



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
	var (
		intOpt  = flag.Int("n", 13, "help message for i option")
	)
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	target := scanner.Text()
	new, _  := Rot(*intOpt,target)
	fmt.Println(new)

}

