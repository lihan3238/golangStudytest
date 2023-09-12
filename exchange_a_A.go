package main

import "fmt"

func main21() {
	a, b := "a", "b"
	fmt.Scanln(&a, &b)
	fmt.Println("Before:" + a + "," + b)
	fmt.Println("After:" + change(a, b))
}

func change(x, y string) string {
	xB, yB := []byte(x), []byte(y)
	for i := 0; i < len(xB); i++ {
		if xB[i] == 'a' || xB[i] == 'c' || xB[i] == 'e' {
			xB[i] = xB[i] - 32
		}
	}
	for i := 0; i < len(yB); i++ {
		if yB[i] == 'B' || yB[i] == 'D' || yB[i] == 'F' {
			yB[i] = yB[i] + 32
		}
	}
	return string(xB) + "," + string(yB)
}
