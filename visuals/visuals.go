package visuals

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type DataModel struct {
	listedVideos []string
	cursorLoc    int
}

type BaseModel struct {
	name string
}

// The main Tea lifecycle consists of 3 functions, which must be implemented on every model
//    Instantiate a model of any sort, and then define an Init(), Update(), and View() function for them

// Init() can optionally return an initial command to run
func (model BaseModel) Init() tea.Cmd {
	return nil
}

// Update() is called when messages are received. Inspect message, send back updated model
//	Can also return a Command (tea.Cmd), which is a function that performs I/O and returns a message

func (model BaseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return model, tea.Quit
		}
	}
	return model, nil
}

// View() returns a string based on the data in the model. This string is rendered to terminal
func (model BaseModel) View() string {
	return model.name
}

func TestProgram() {
	prog := tea.NewProgram(BaseModel{"Hi mom!"})
	if _, err := prog.Run(); err != nil {
		log.Fatal(err)
	}

}

// NOTE: Example Command (tea.Cmd). Defined as "something that performs I/O, then returns a message"...
//	... where a "message" (tea.Msg) is an interface that you pack with data, which then triggers Update()

// NOTE: I'm pretty sure tea.Msg can just be literally anything since it's defined as an empty interface, so...
//	... we should define our own message types, and then switch on type, I guess
//	We can just have something like ToggleWatchedMsg, or QueryAPIMsg

// type tickMsg time.Time
//
// func tick() tea.Msg {
// 	time.Sleep(time.Second)
// 	return tickMsg{}
// }
