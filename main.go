package main

import "fmt"

func main()  {
	var aux uint64
	var suma uint64
	var n uint64
	var i uint64
	var s []uint64
	fmt.Scanln(&n)
	for i=1;i<=n;i++{
	  fmt.Scanln(&aux)
	  s = append(s,aux)
	}
	for i=0;i<n;i++{
	  aux = s[i]
	  suma += aux
	}
	fmt.Println(suma)
}