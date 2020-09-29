package main

import "fmt"

func main() {
	var mes int64
	var dia int64

	
	fmt.Scan(&dia)
	
	fmt.Scan(&mes)

	switch mes {
	case 1:
		if dia < 21 {
			fmt.Println("capricornio")
		} else {
			fmt.Println("ccuario")
		}
	
	case 2:
		if dia < 21 {
			fmt.Println("acuario")
		} else {
			fmt.Println("piscis")
		}
	
	case 3:
		if dia < 21 {
			fmt.Println("piscis")
		} else {
			fmt.Println("aries")
		}

	case 4:
		if dia < 21 {
			fmt.Println("aries")
		} else {
			fmt.Println("tauro")
		}

	case 5:
		if dia < 22 {
			fmt.Println("tauro")
		} else {
			fmt.Println("geminis")
		}

	case 6:
		if dia < 22 {
			fmt.Println("geminis")
		} else {
			fmt.Println("cancer")
		}
		
	case 7:
		if dia < 23 {
			fmt.Println("cancer")
		} else {
			fmt.Println("leo")
		}
		
	case 8:
		if dia < 23 {
			fmt.Println("leo")
		} else {
			fmt.Println("virgo")
		}
		
	case 9:
		if dia < 24 {
			fmt.Println("virgo")
		} else {
			fmt.Println("libra")
		}

	case 10:
		if dia < 25 {
			fmt.Println("libra")
		} else {
			fmt.Println("escorpio")
		}

	case 11:
		if dia < 23 {
			fmt.Println("escorpio")
		} else {
			fmt.Println("sagitario")
		}

	case 12:
		if dia < 22 {
			fmt.Println("sagitario")
		} else {
			fmt.Println("capricornio")
		}



	}
}