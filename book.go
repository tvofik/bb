package main

type Book struct {
	id        string
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

/*
For generating range of numbers
package main

import "fmt"

func main() {
	chapters := []int{}
	for i := range [100]int{} {
		fmt.Printf("9%03d\n", i+1) // prints 9000 - 9100
	}
}


*/
// }

/* FOR REFERENCE
// Folder represents a group of snippets in a directory.
type Folder string

// FilterValue is the searchable value for the folder.
func (f Folder) FilterValue() string {
	return string(f)
}

// folderDelegate represents a folder list item.
type folderDelegate struct{ styles FoldersBaseStyle }

// Height is the number of lines the folder list item takes up.
func (d folderDelegate) Height() int {
	return 1
}

// Spacing is the number of lines to insert between folder items.
func (d folderDelegate) Spacing() int {
	return 0
}

// Update is what is called when the folder selection is updated.
// TODO: Update the filter search for the snippets with the folder name.
func (d folderDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

// Render renders a folder list item.
func (d folderDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	f, ok := item.(Folder)
	if !ok {
		return
	}
	fmt.Fprint(w, "  ")
	if index == m.Index() {
		fmt.Fprint(w, d.styles.Selected.Render("→ "+string(f)))
		return
	}
	fmt.Fprint(w, d.styles.Unselected.Render("  "+string(f)))
}
*/
