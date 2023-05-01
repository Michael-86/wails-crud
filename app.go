package main

import (
	"context"
	"encoding/json"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

//json.Unmarshal(brownBytes, &people)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (b *App) shutdown(ctx context.Context) {

}

type Kabel struct {
	Fname string `json:"first"`
	Sname string `json:"last"`
}

// Greet returns a greeting for the given name
func (a *App) Print1(people Kabel) string {
	var kablar []Kabel
	// json.Unmarshal(content, &friends)
	err2 := json.Unmarshal([]byte(people.Fname), &kablar)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
	for _, x := range kablar {
		fmt.Printf("%s \n", x.Fname)
	}
	fmt.Printf("%s \n", people.Fname)
	return fmt.Sprintf("Hello %s", people.Fname)
}
