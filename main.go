package main

import "fmt"

func main() {
	var mes int64
	var dia int64

	fmt.Println("Ingrese su dia de nacimiento: ")
	fmt.Scan(&dia)
	fmt.Println("Ingrese su mes de nacimiento: ")
	fmt.Scan(&mes)

	switch mes {
	case 1:
		if dia < 21 {
			fmt.Println("Capricornio")
		} else {
			fmt.Println("Acuario")
		}
	
	case 2:
		if dia < 21 {
			fmt.Println("Acuario")
		} else {
			fmt.Println("Piscis")
		}
	
	case 3:
		if dia < 21 {
			fmt.Println("Piscis")
		} else {
			fmt.Println("Aries")
		}

	case 4:
		if dia < 21 {
			fmt.Println("Aries")
		} else {
			fmt.Println("Tauro")
		}

	case 5:
		if dia < 22 {
			fmt.Println("Tauro")
		} else {
			fmt.Println("Geminis")
		}

	case 6:
		if dia < 22 {
			fmt.Println("Geminis")
		} else {
			fmt.Println("Cancer")
		}
		
	case 7:
		if dia < 23 {
			fmt.Println("Cancer")
		} else {
			fmt.Println("Leo")
		}
		
	case 8:
		if dia < 23 {
			fmt.Println("Leo")
		} else {
			fmt.Println("Virgo")
		}
		
	case 9:
		if dia < 24 {
			fmt.Println("Virgo")
		} else {
			fmt.Println("Libra")
		}

	case 10:
		if dia < 25 {
			fmt.Println("Libra")
		} else {
			fmt.Println("Escorpio")
		}

	case 11:
		if dia < 23 {
			fmt.Println("Escorpio")
		} else {
			fmt.Println("Sagitario")
		}

	case 12:
		if dia < 22 {
			fmt.Println("Sagitario")
		} else {
			fmt.Println("Capricornio")
		}



	}
}