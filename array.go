package main

import (
	"fmt"
)

func main() {


	PrAr("Hello", "Play","Jorj","12-02-2014")
	PrArt([]string{"Hello", "playground","Jorik","Linked"})
}



func PrArt(Prntxt []string){
    for _, Pr:=range Prntxt {
           fmt.Println(Pr)
    }
}

func PrAr(Prntxt ...string){
    for _, Pr:=range Prntxt {
           fmt.Println(Pr)
    }
}