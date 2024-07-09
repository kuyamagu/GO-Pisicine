package piscine

import "ft"

func PrintAlp() {
	for str := 'a'; str <= 'z'; str++ {
		ft.PrintRune(str)
	}
	ft.PrintRune('\n')
}
