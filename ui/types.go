package ui

import (
	"weather-app/apptype"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	Window fyne.Window
	State  *apptype.State
}
