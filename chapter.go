package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	// helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

/* Implementing the Default */
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

// /* Implementing the Customized */
// type Chapter string

// // implement the List.Item interface
// func (c Chapter) FilterValue() string {
// 	return string(c)
// }

// type chapterDelegate struct{}

// func (d chapterDelegate) Height() int                               { return 1 }
// func (d chapterDelegate) Spacing() int                              { return 0 }
// func (d chapterDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

// func (d chapterDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
// 	c, ok := item.(Chapter)
// 	if !ok {
// 		return
// 	}
// 	str := fmt.Sprintf("%s", c)

// 	fn := itemStyle.Render
// 	if index == m.Index() {
// 		fn = func(s ...string) string {
// 			return selectedItemStyle.Render("[ " + strings.Join(s, "") + " ]")
// 		}
// 	}
// 	fmt.Fprint(w, fn(str))
// }
