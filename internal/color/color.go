package color

import "fmt"

func Custom(s string, color int) string {
	return fmt.Sprintf("\x1b[38;5;%dm", color) + s + "\x1b[0m"
}
func Green(s string) string {
	return Custom(s, 76)
}
func Red(s string) string {
	return Custom(s, 202)
}
func Yellow(s string) string {
	return Custom(s, 11)
}
func Grey(s string) string {
	return Custom(s, 245)
}
