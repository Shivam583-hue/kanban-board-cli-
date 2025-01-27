package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor = 4

const (
	todo status = iota
	inProgress
	done
)

var models []tea.Model

const (
	maamaa status = iota
	form
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
	status      status
	title       string
	description string
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
	return t.description
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

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "n":
			models[maamaa] = m // save the state of the current model
			newForm := NewForm(m.focused)
			models[form] = newForm // Update the form model in the slice
			return models[form], nil
		case "x":
			selectedItem := m.lists[m.focused].SelectedItem()
			if selectedTask, ok := selectedItem.(task); ok {
				m.lists[selectedTask.status].RemoveItem(m.lists[m.focused].Index())
			}
			return m, nil
		}
	case task:
		task := msg
		return m, m.lists[task.status].InsertItem(len(m.lists[task.status].Items()), task)
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
	// m.lists[selectedTask.status].InsertItem(len(m.lists[selectedTask.status].Items())-1, list.Item(selectedTask))
	m.lists[selectedTask.status].InsertItem(
		len(m.lists[selectedTask.status].Items()), // Correct index to append
		list.Item(selectedTask),
	)
	return nil
}

// Form model
type Form struct {
	focused     status
	title       textinput.Model
	description textarea.Model
}

func NewForm(focused status) Form {
	form := Form{focused: focused}
	form.title = textinput.New()
	form.title.Placeholder = "Enter title"
	form.title.Focus()
	form.description = textarea.New()
	form.description.Placeholder = "Enter description"
	return form
}

func (m Form) Init() tea.Cmd {
	return nil
}

func NewTask(status status, title, description string) task {
	return task{
		status:      status,
		title:       title,
		description: description,
	}
}

func (m Form) CreateTask() tea.Msg {
	return NewTask(m.focused, m.title.Value(), m.description.Value())
}

func (m Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.title.Focused() {
				m.title.Blur()
				m.description.Focus()
				return m, textarea.Blink
			} else {
				// models[maamaa] = m // save the state of the current model
				return models[maamaa], m.CreateTask
			}
		}
	}
	if m.title.Focused() {
		m.title, cmd = m.title.Update(msg)
		return m, cmd
	} else {
		m.description, cmd = m.description.Update(msg)
		return m, cmd
	}
}

func (m Form) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.title.View(), m.description.View())
}

func main() {
	// m := New()
	models = []tea.Model{New(), NewForm(todo)}
	m := models[0]
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
