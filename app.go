package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	notePath    string
	alwaysOnTop bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	notesDir := "notes"

	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error:", err)
		return
	}

	a.notePath = filepath.Join(home, notesDir, "temp.txt")
	a.alwaysOnTop = false
}

func (a *App) SaveNote(content string) error {

	noteDir := filepath.Dir(a.notePath)
	// Create dir
	err := os.MkdirAll(noteDir, 0755)
	if err != nil {
		log.Printf("Error creating note folder: %v", err)
		return err
	}

	err = os.WriteFile(a.notePath, []byte(content), 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (a *App) LoadNote() string {
	data, err := os.ReadFile(a.notePath)
	if err != nil {
		if os.IsNotExist(err) {
			return ""
		}
		log.Println(err)
		return ""
	}

	content := string(data)
	return content
}

func (a *App) ToggleAlwaysOnTop() bool {
	a.alwaysOnTop = !a.alwaysOnTop
	runtime.WindowSetAlwaysOnTop(a.ctx, a.alwaysOnTop)
	return a.alwaysOnTop
}
