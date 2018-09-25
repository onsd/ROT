package main

import "fmt"

func Rot(N int,plain string) ( string, error){
	var words string
	var ascii int
	for _,v := range plain{
		num := uint8(v)
		fmt.Println(num)
		if num >= 65 && num <= 77{
			ascii = int(num) + N
		}else if num >= 78 && num <= 90 {
			ascii = int(num) - N
		}else if num >= 97 && num <= 109 {
			ascii = int(num) + N
		}else if num >= 110 && num <= 123{
			ascii = int(num) - N
		}else{
			ascii = int(num)
		}
		words += string(ascii)
	}
	return words,nil
}

func main(){
	//str := "Hello, World."
	var str string
	var N int
	fmt.Println("input the string")
	fmt.Scan(&str)
	fmt.Println("input rotation number")
	fmt.Scan(&N)
	new, _  := Rot(N,str)
	fmt.Println(new)

}
