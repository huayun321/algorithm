package main

import "log"

func gcd(a, b int) int {
	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}

func recursive(a, b int) int {
	if b == 0 {
		return a
	}
	return recursive(b, a%b)
}

func main() {
	log.Println(gcd(8, 6))
	log.Println(recursive(8, 6))
}
