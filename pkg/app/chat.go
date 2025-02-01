package app

import "fmt"

func (a *App) Message(url string, message string) string {
	return fmt.Sprintf("Hello %s, It's show time!", url)
}



