/* PART ONE */
package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sessionState int

const (
	translationColumn sessionState = iota
	bookColumn
	chapterColumn
	passageColumn
)

/* MAIN MODEL */
type Model struct {
	loaded bool
	// spinner      spinner.Model //TODO: Add spinner while we fetch the data
	focused sessionState
	// The List for translation, book & chapters Translations to display
	columns []list.Model
	// The viewport for the book and chapter
	passage  viewport.Model
	content  string
	err      error
	quitting bool
}

func New() *Model {
	// return &Model{focused: bookColumn}
	return &Model{focused: bookColumn, loaded: false}
}

type Passage struct {
	book    string  // Name of the book
	chapter string  // Chapter Number
	verses  []Verse // Verse Number
}

type Verse struct {
	verse string // Verse Number
	text  string // Text of the verse
}

type GetPassageMsg struct {
	Err     error
	Passage []Verse
}

/* Used to get the list of number of chapters in a book */
type GetBookChaptersMsg struct {
	chapterList []list.Item
}

// Tells the application to get the passage
func (m *Model) GetPassage(bookId, chapter, translation string) tea.Cmd {
	return func() tea.Msg {
		msg, err := m.GetChapter(bookId, chapter, translation)
		if err != nil {
			return GetPassageMsg{Err: err}
		}
		return GetPassageMsg{Passage: msg}
	}
}

/* Gets the list of chapters in a book */
func (m Model) GetBookChapters() tea.Msg {
	var chapterList []list.Item
	selectedItem := m.columns[bookColumn].SelectedItem()
	selectedBook := selectedItem.(Book)
	chapters := selectedBook.chapters
	for i := 0; i < chapters; i++ {
		num := fmt.Sprintf("%v", i+1)
		chapterList = append(chapterList, Chapter(num))
	}
	return GetBookChaptersMsg{chapterList: chapterList}
}

/* Move to the next column */
func (m *Model) Next() {
	if m.focused == bookColumn {
		m.focused = chapterColumn
	} else {
		m.focused = bookColumn
	}
}

/* Move to the previous column */
func (m *Model) Prev() {
	if m.focused == chapterColumn {
		m.focused = bookColumn
	}
}

/* Get the selected book from field */
func (m Model) GetSelectedBookId() string {
	selectedItem := m.columns[bookColumn].SelectedItem()
	selectedBook := selectedItem.(Book)
	return selectedBook.id
}

/* Get the selected chapter item from field */
func (m Model) GetSelectedChapterId() string {
	selectedItem := m.columns[chapterColumn].SelectedItem()
	selectedChapter := selectedItem.(Chapter)
	return selectedChapter.FilterValue() //returning the filtervalue which is the id
}

/* Get the selected translation & gets the table name */
func (m Model) GetSelectedTranslation() string {
	selectedItem := m.columns[translationColumn].SelectedItem()
	selectedTranslation := selectedItem.(Translation)

	translation := selectedTranslation.abbreviation
	tableName, err := m.GetCurrentTranslationTable(translation)
	if err != nil {
		return fmt.Sprint(err)
	}
	return tableName
}

func (m *Model) initModel(width, height int) {
	// Query the DB to get the books
	books, err := GetBooks()
	if err != nil {
		fmt.Println(err)
	}
	// Query the DB to get the translations
	translations, err := GetTranslations()
	if err != nil {
		fmt.Println(err)
	}

	translationsList := list.New(translations, list.NewDefaultDelegate(), width, height)
	bookList := list.New(books, list.NewDefaultDelegate(), width, height)
	chapterList := list.New([]list.Item{}, chapterDelegate{}, width, height)

	m.columns = []list.Model{translationsList, bookList, chapterList}

	// Init Translation
	m.columns[translationColumn].Title = "Translations"
	m.columns[translationColumn].FilterInput.Prompt = "Find Translation: "
	m.columns[translationColumn].SetStatusBarItemName("Translation", "Translations")

	// Init Books
	m.columns[bookColumn].Title = "Books"
	m.columns[bookColumn].FilterInput.Prompt = "Find Book: "
	m.columns[bookColumn].SetStatusBarItemName("Book", "Books")

	// Init Chapters
	m.columns[chapterColumn].Title = "Chapters"
	m.columns[chapterColumn].FilterInput.Prompt = "Find Chapter: "
	m.columns[chapterColumn].SetStatusBarItemName("Chapter", "Chapters")
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initModel(msg.Width, msg.Height)
			m.loaded = true
			return m, m.GetBookChapters
		}

	case GetBookChaptersMsg:
		m.columns[chapterColumn].SetItems(msg.chapterList)

	case tea.KeyMsg:
		switch msg.String() {
		case "down", "up":
			if m.focused == bookColumn {
				var cmd tea.Cmd
				var cmds []tea.Cmd
				m.columns[bookColumn], cmd = m.columns[bookColumn].Update(msg)
				cmds = append(cmds, m.GetBookChapters, cmd)
				return m, tea.Batch(cmds...)
			}
		case "right":
			m.Next()
		case "left":
			m.Prev()
		}
	}

	var cmd tea.Cmd
	m.columns[m.focused], cmd = m.columns[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return "Jesus ❤️  You!" // TODO: Not working
	}

	if m.loaded {
		translationView := m.columns[translationColumn].View()
		bookView := m.columns[bookColumn].View()
		chapterView := m.columns[chapterColumn].View()
		switch m.focused {
		case translationColumn:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				focusedStyle.Render(translationView),
				columnStyle.Render(bookView),
				columnStyle.Render(chapterView),
			)
		case chapterColumn:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(translationView),
				columnStyle.Render(bookView),
				focusedStyle.Render(chapterView),
			)
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(translationView),
				focusedStyle.Render(bookView),
				columnStyle.Render(chapterView),
			)
		}
	} else {
		return "Loading..."
	}
}

func main() {
	m := New()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
