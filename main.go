package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Welcome to passowrd Stength Checker.")

	for {

		fmt.Println("Enter the password (Ctrl + C for exit):")
		var password string
		fmt.Scan(&password)
		if len([]rune(password)) < 1 {

			fmt.Println("Can't use negative or zero password.")
			continue
		} else {
			result := score(password)
			fmt.Println("Your score for\"", password, "\" password is :", result)
		}
	}

}
func score(pass string) string {
	password := []rune(pass)

	if len(password) >= 8 {
		if strings.ContainsAny(pass, "abcdefghijklmnopqrstuvwxyz") {

			if strings.ContainsAny(pass, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {

				if strings.ContainsAny(pass, "0123456789") {

					if strings.ContainsAny(pass, "!@#$%^&*()-_=+[]{}|;:,.<>/") {

						return "strong"
					} else {
						return "medium and no-special-symbols"
					}

				} else {
					if strings.ContainsAny(pass, "!@#$%^&*()-_=+[]{}|;:,.<>/") {

						return "strong"
					} else {
						return "medium and no-special-symbols"
					}
				}

			} else {
				if strings.ContainsAny(pass, "0123456789") {

					if strings.ContainsAny(pass, "!@#$%^&*()-_=+[]{}|;:,.<>/") {

						return "strong"
					} else {
						return "medium and no-special-symbols"
					}

				} else {
					return "medium and no-number"
				}
			}

		} else {
			if strings.ContainsAny(pass, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {

				if strings.ContainsAny(pass, "0123456789") {

					if strings.ContainsAny(pass, "!@#$%^&*()-_=+[]{}|;:,.<>/") {

						return "strong"
					} else {
						return "medium and no-special-symbols"
					}

				} else {
					if strings.ContainsAny(pass, "!@#$%^&*()-_=+[]{}|;:,.<>/") {

						return "strong"
					} else {
						return "medium and no-special-symbols"
					}
				}

			} else {
				if strings.ContainsAny(pass, "0123456789") {

					if strings.ContainsAny(pass, "!@#$%^&*()-_=+[]{}|;:,.<>/") {

						return "strong"
					} else {
						return "medium and no-special-symbols"
					}

				} else {
					return "medium and no-number"
				}
			}
			
		}

	} else {
		return "weak"
	}

}
