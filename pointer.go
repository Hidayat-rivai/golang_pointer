package main

import "fmt"

func main(){
	var angka = 10
	var alamat_angka *int = &angka; 
	fmt.Println("Alamat angka : ",alamat_angka)
}