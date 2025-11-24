package main

import (
	"context"
	"fmt"
	"go-cad/internal/parser"
	"log"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Parse(fileName string) *parser.ParseResult {
	fmt.Printf("Opening file: %q\n", fileName)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	res, err := parser.ParseAuto(f)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (a *App) ParseFile(fileContent string) *parser.ParseResult {
	fmt.Println("Parsing content...")

	reader := strings.NewReader(fileContent)

	res, err := parser.ParseAuto(reader)
	if err != nil {
		fmt.Println("Parse error:", err)
		return nil
	}

	return res
}

func (a *App) OpenFileDialog() (string, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Выберите файл схемы (.NET)",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "NET схемы",
				Pattern:     "*.NET;*.net",
			},
			{
				DisplayName: "Все файлы",
				Pattern:     "*",
			},
		},
	})

	if err != nil {
		return "", err
	}

	return path, nil
}

func (a *App) ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
