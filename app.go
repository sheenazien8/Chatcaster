package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v2"
)

type App struct {
	ctx    context.Context
	config Config
}

type BubbleChatConfig struct {
	ShowKey struct {
		Mode string `yaml:"mode"` // Options: 'focus', 'global'
	} `yaml:"show_key"`
	DisplayLimit int `yaml:"display_limit"` // Maximum number of chat bubbles to display
	Style        struct {
		BackgroundColor string `yaml:"background_color"` // Hex code or color name
		FontColor       string `yaml:"font_color"`
		FontSize        string `yaml:"font_size"`
	} `yaml:"style"`
	Content struct {
		ShowUserIcon bool   `yaml:"show_user_icon"` // Enable or disable user icon in chat bubbles
		UserIconPath string `yaml:"user_icon_path"` // Path to the user icon image, leave it empty will use default image
	} `yaml:"content"`
}

type Config struct {
	BubbleChat BubbleChatConfig `yaml:"bubble_chat"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	runtime.WindowSetPosition(ctx, 0, 0)
	a.ctx = ctx
	a.generateConfig()
}

func (a *App) generateConfig() {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "chatcaster", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			fmt.Println("Error creating config directory:", err)
			return
		}
		defaultConfigPath := filepath.Join(".config", "config.yaml")
		data, err := os.ReadFile(defaultConfigPath)
		if err != nil {
			fmt.Println("Error reading default config file:", err)
			return
		}
		err = os.WriteFile(configPath, data, 0644)
		if err != nil {
			fmt.Println("Error writing default config file:", err)
			return
		}
		fmt.Println("Default config file created at", configPath)
	} else {
		fmt.Println("Config file found at", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	err = yaml.Unmarshal(data, &a.config)
	if err != nil {
		fmt.Println("Error unmarshalling config file:", err)
		return
	}

	runtime.EventsEmit(a.ctx, "configLoaded", a.config)

	if a.config.BubbleChat.ShowKey.Mode == "global" {
		go func() {
			events := hook.Start()
			defer hook.End()

			for event := range events {
				if event.Kind == hook.KeyDown {
					key := a.HandleModifier(event.Keychar)
					runtime.EventsEmit(a.ctx, "keyPressed", key)
				}
			}
		}()
	}
}

func (a *App) HandleModifier(char rune) string {
	keyChar := ""
	fmt.Println(char)
	switch char {
	case 8:
		keyChar = "Backspace"
	case 9:
		keyChar = "Tab"
	case 27:
		keyChar = "Esc"
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

func (a *App) Config() Config {
	return a.config
}
