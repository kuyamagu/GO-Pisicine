package piscine

import "ft"

func PrintDigits(){
	for	num := '0'; num <= '9'; num++{
		ft.PrintRune(num)
	}
	ft.PrintRune('\n')
}
