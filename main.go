package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor = 4

const (
	todo status = iota
	inProgress
	done
)

var (
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2)
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	// helpStyle = lipgloss.NewStyle().
	//		Foreground(lipgloss.Color("241"))
)

// CUSTOM ITEM

type status int

type task struct {
	status    status
	title     string
	descption string
}

func (t *task) Next() {
	if t.status == done {
		t.status = todo
	} else {
		t.status++
	}
}

// implement the list.Item interface

func (t task) FilterValue() string {
	return t.title
}

func (t task) Title() string {
	return t.title
}

func (t task) Description() string {
	return t.descption
}

// MAIN MODEL
type model struct {
	quitting bool
	focused  status
	lists    []list.Model
	loaded   bool
}

// TODO : call this on tea.WindowSizeMsg
func (m *model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height/2)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	// Init to do
	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		task{todo, "Buy milk", "Milk is good"},
		task{todo, "Sleep", "Sleep is good"},
	})

	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		task{inProgress, "Eat", "Eat some food"},
	})

	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		task{done, "Go to the gym", "Go to the gym"},
	})
}

// convert the model into a bubbletea model
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			columnStyle.Width(msg.Width / divisor)
			focusedStyle.Width(msg.Width / divisor)
			columnStyle.Height(msg.Height - divisor)
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			m.Prev()
		case "right", "l":
			m.Next()
		case "enter":
			return m, m.MoveToNext
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	if m.loaded {
		todoView := m.lists[todo].View()
		inProgressView := m.lists[inProgress].View()
		doneView := m.lists[done].View()
		switch m.focused {
		case inProgress:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(todoView),
				focusedStyle.Render(inProgressView),
				columnStyle.Render(doneView),
			)
		case done:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(todoView),
				columnStyle.Render(inProgressView),
				focusedStyle.Render(doneView),
			)
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				focusedStyle.Render(todoView),
				columnStyle.Render(inProgressView),
				columnStyle.Render(doneView),
			)
		}
	} else {
		return "Loading..."
	}
}

func New() *model {
	return &model{}
}

func (m *model) MoveToNext() tea.Msg {
	selectedItem := m.lists[m.focused].SelectedItem()
	selectedTask := selectedItem.(task)
	m.lists[selectedTask.status].RemoveItem(m.lists[m.focused].Index())
	selectedTask.Next()
	m.lists[selectedTask.status].InsertItem(len(m.lists[selectedTask.status].Items())-1, list.Item(selectedTask))
	return nil
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// TODO : go to next list
func (m *model) Next() {
	if m.focused == done {
		m.focused = todo
	} else {
		m.focused++
	}
}

// TODO : go to previous list
func (m *model) Prev() {
	if m.focused == todo {
		m.focused = done
	} else {
		m.focused--
	}
}
