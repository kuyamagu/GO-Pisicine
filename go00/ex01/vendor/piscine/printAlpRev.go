package piscine

import "ft"

func PrintAlpRev() {
	for	str := 'z'; 'a' <= str; str--{
		ft.PrintRune(str)
	}
	ft.PrintRune('\n')
}
