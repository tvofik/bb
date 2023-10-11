/* PART ONE */
package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wrap"
)

/* STYLING */
var (
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2)
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

// ! testing
var p = "1 The book of the generation of Jesus Christ, the son of David, the son of Abraham. 2 Abraham begat Isaac; and Isaac begat Jacob; and Jacob begat Judas and his brethren; 3 And Judas begat Phares and Zara of Thamar; and Phares begat Esrom; and Esrom begat Aram; 4 And Aram begat Aminadab; and Aminadab begat Naasson; and Naasson begat Salmon; 5 And Salmon begat Booz of Rachab; and Booz begat Obed of Ruth; and Obed begat Jesse; 6 And Jesse begat David the king; and David the king begat Solomon of her that had been the wife of Urias; 7 And Solomon begat Roboam; and Roboam begat Abia; and Abia begat Asa; 8 And Asa begat Josaphat; and Josaphat begat Joram; and Joram begat Ozias; 9 And Ozias begat Joatham; and Joatham begat Achaz; and Achaz begat Ezekias; 10 And Ezekias begat Manasses; and Manasses begat Amon; and Amon begat Josias; 11 And Josias begat Jechonias and his brethren, about the time they were carried away to Babylon: 12 And after they were brought to Babylon, Jechonias begat Salathiel; and Salathiel begat Zorobabel; 13 And Zorobabel begat Abiud; and Abiud begat Eliakim; and Eliakim begat Azor; 14 And Azor begat Sadoc; and Sadoc begat Achim; and Achim begat Eliud; 15 And Eliud begat Eleazar; and Eleazar begat Matthan; and Matthan begat Jacob; 16 And Jacob begat Joseph the husband of Mary, of whom was born Jesus, who is called Christ. 17 So all the generations from Abraham to David are fourteen generations; and from David until the carrying away into Babylon are fourteen generations; and from the carrying away into Babylon unto Christ are fourteen generations. 18 Now the birth of Jesus Christ was on this wise: When as his mother Mary was espoused to Joseph, before they came together, she was found with child of the Holy Ghost. 19 Then Joseph her husband, being a just man, and not willing to make her a publick example, was minded to put her away privily. 20 But while he thought on these things, behold, the angel of the LORD appeared unto him in a dream, saying, Joseph, thou son of David, fear not to take unto thee Mary thy wife: for that which is conceived in her is of the Holy Ghost. 21 And she shall bring forth a son, and thou shalt call his name JESUS: for he shall save his people from their sins. 22 Now all this was done, that it might be fulfilled which was spoken of the Lord by the prophet, saying, 23 Behold, a virgin shall be with child, and shall bring forth a son, and they shall call his name Emmanuel, which being interpreted is, God with us. 24 Then Joseph being raised from sleep did as the angel of the Lord had bidden him, and took unto him his wife: 25 And knew her not till she had brought forth her firstborn son: and he called his name JESUS."

type pane int

const (
	bookPane pane = iota
	chapterPane
	passagePane
)

type Book struct {
	title     string
	fullTitle string
	chapters  int
}

// implement the List.Item interface
func (b Book) FilterValue() string {
	return b.title
}
func (b Book) Title() string {
	return b.title
}

func (b Book) Description() string {
	return b.fullTitle
}

type Chapter struct {
	number      int
	description string
}

// implement the List.Item interface
func (c Chapter) FilterValue() string {
	return ""
}
func (c Chapter) Title() string {
	return fmt.Sprintf("%d", c.number)
}

func (c Chapter) Description() string {
	return c.description
}

/* MAIN MODEL */
type Model struct {
	loaded   bool
	focused  pane
	lists    []list.Model
	viewport viewport.Model
	content  string
	err      error
	quitting bool
}

func New() *Model {
	return &Model{}
}
func (m *Model) updateChapters() tea.Cmd {
	return func() tea.Msg {
		selectedItem := m.lists[m.focused].SelectedItem()
		selectedBook := selectedItem.(Book)
		fmt.Println(selectedBook.chapters)
		return selectedBook.chapters
	}
}

// type chapterMsg int

// func (m *Model) updateFolders() tea.Cmd {
// 	return func() tea.Msg {
// 		msg := m.updateFoldersView()
// 		return msg
// 	}
// }

/* Move between the panes*/
func (m *Model) Next() {
	if m.focused == bookPane {
		m.focused = chapterPane
	}
}

func (m *Model) Prev() {
	if m.focused == chapterPane {
		m.focused = bookPane
	}
}

func (m *Model) initLists(width, height int) {

	// Query the DB to get the books
	books, err := GetBooks()
	if err != nil {
		fmt.Println(err)
	}
	bookList := list.New(books, list.NewDefaultDelegate(), width, height)
	chapterList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	passageView := viewport.New(width, height)
	m.lists = []list.Model{bookList, chapterList}

	// Init Books
	m.lists[bookPane].Title = "Books"
	m.lists[bookPane].FilterInput.Prompt = "Find Book: "
	m.lists[bookPane].SetStatusBarItemName("Book", "Books")

	// Init Chapters
	m.lists[chapterPane].Title = "Chapters"
	m.lists[chapterPane].FilterInput.Prompt = "Find Chapter: "
	m.lists[chapterPane].SetStatusBarItemName("Chapter", "Chapters")
	m.lists[chapterPane].SetItems([]list.Item{
		Chapter{number: 1, description: ""},
		Chapter{number: 2, description: ""},
		Chapter{number: 3, description: ""},
		Chapter{number: 4, description: ""},
		Chapter{number: 5, description: ""},
		Chapter{number: 6, description: ""},
		Chapter{number: 7, description: ""},
		Chapter{number: 8, description: ""},
		Chapter{number: 9, description: ""},
		Chapter{number: 10, description: ""},
		Chapter{number: 11, description: ""},
	})

	// Init Passage
	m.viewport = passageView
	passage := wrap.String(p, width-10)
	m.viewport.SetContent(passage)
}

func (m Model) Init() tea.Cmd {
	return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "right":
			m.Next()
		case "left":
			m.Prev()
		case "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
			// case "enter":
			// return m, nil
			// return m, m.updateChapters
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	if m.loaded {
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.lists[bookPane].View(),
			m.lists[chapterPane].View(),
			m.viewport.View(),
		)
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