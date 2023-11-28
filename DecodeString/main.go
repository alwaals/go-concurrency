package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
    message := "VEZEU0ZVVFVTSk9I"
	result := DecodeSecret(message)
	fmt.Println(result)
}
func DecodeSecret(str string) string{
	data,err:=base64.StdEncoding.DecodeString(str)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
	var secret []rune
	for _, char := range data {
		//fmt.Println(rune(char + 1))
		secret = append(secret, rune(char-1))
	}

	//fmt.Println(secret)
	return string(secret)
}