package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	// go func() {
	// 	events := hook.Start()
	// 	defer hook.End()
	//
	// 	for event := range events {
	// 		if event.Kind == hook.KeyDown {
	// 			key := a.HandleModifier(event.Keychar)
	// 			runtime.EventsEmit(ctx, "keyPressed", key)
	// 		}
	// 	}
	// }()
	runtime.WindowSetPosition(ctx, 0, 0)
	a.ctx = ctx
}

func (a *App) HandleModifier(char rune) string {
	keyChar := ""
	fmt.Println(char)
	switch char {
	case 8:
		keyChar = "Backspace"
	case 13:
		keyChar = "Enter"
	case 'A':
		keyChar = "Alt"
	case 'C':
		keyChar = "Ctrl"
	case 'S':
		keyChar = "Shift"
	default:
		keyChar = fmt.Sprintf("%c", char)
	}
	return keyChar
}
