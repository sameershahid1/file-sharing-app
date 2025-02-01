package app

import "fmt"

func (a *App) FileShare(url string, message string) string {
	return fmt.Sprintf("Hello %s, It's show time!", url)
}




