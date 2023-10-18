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

// !  testing
var p = `
1 Blessed are the undefiled in the way, who walk in the law of the LORD. 2 Blessed are they that keep his testimonies, and that seek him with the whole heart. 3 They also do no iniquity: they walk in his ways. 4 Thou hast commanded us to keep thy precepts diligently. 5 O that my ways were directed to keep thy statutes! 6 Then shall I not be ashamed, when I have respect unto all thy commandments. 7 I will praise thee with uprightness of heart, when I shall have learned thy righteous judgments. 8 I will keep thy statutes: O forsake me not utterly. 9 Wherewithal shall a young man cleanse his way? by taking heed thereto according to thy word. 10 With my whole heart have I sought thee: O let me not wander from thy commandments. 11 Thy word have I hid in mine heart, that I might not sin against thee. 12 Blessed art thou, O LORD: teach me thy statutes. 13 With my lips have I declared all the judgments of thy mouth. 14 I have rejoiced in the way of thy testimonies, as much as in all riches. 15 I will meditate in thy precepts, and have respect unto thy ways. 16 I will delight myself in thy statutes: I will not forget thy word. 17 Deal bountifully with thy servant, that I may live, and keep thy word. 18 Open thou mine eyes, that I may behold wondrous things out of thy law. 19 I am a stranger in the earth: hide not thy commandments from me. 20 My soul breaketh for the longing that it hath unto thy judgments at all times. 21 Thou hast rebuked the proud that are cursed, which do err from thy commandments. 22 Remove from me reproach and contempt; for I have kept thy testimonies. 23 Princes also did sit and speak against me: but thy servant did meditate in thy statutes. 24 Thy testimonies also are my delight and my counsellors. 25 My soul cleaveth unto the dust: quicken thou me according to thy word. 26 I have declared my ways, and thou heardest me: teach me thy statutes. 27 Make me to understand the way of thy precepts: so shall I talk of thy wondrous works. 28 My soul melteth for heaviness: strengthen thou me according unto thy word. 29 Remove from me the way of lying: and grant me thy law graciously. 30 I have chosen the way of truth: thy judgments have I laid before me. 31 I have stuck unto thy testimonies: O LORD, put me not to shame. 32 I will run the way of thy commandments, when thou shalt enlarge my heart. 33 Teach me, O LORD, the way of thy statutes; and I shall keep it unto the end. 34 Give me understanding, and I shall keep thy law; yea, I shall observe it with my whole heart. 35 Make me to go in the path of thy commandments; for therein do I delight. 36 Incline my heart unto thy testimonies, and not to covetousness. 37 Turn away mine eyes from beholding vanity; and quicken thou me in thy way. 38 Stablish thy word unto thy servant, who is devoted to thy fear. 39 Turn away my reproach which I fear: for thy judgments are good. 40 Behold, I have longed after thy precepts: quicken me in thy righteousness. 41 Let thy mercies come also unto me, O LORD, even thy salvation, according to thy word. 42 So shall I have wherewith to answer him that reproacheth me: for I trust in thy word. 43 And take not the word of truth utterly out of my mouth; for I have hoped in thy judgments. 44 So shall I keep thy law continually for ever and ever. 45 And I will walk at liberty: for I seek thy precepts. 46 I will speak of thy testimonies also before kings, and will not be ashamed. 47 And I will delight myself in thy commandments, which I have loved. 48 My hands also will I lift up unto thy commandments, which I have loved; and I will meditate in thy statutes. 49 Remember the word unto thy servant, upon which thou hast caused me to hope. 50 This is my comfort in my affliction: for thy word hath quickened me. 51 The proud have had me greatly in derision: yet have I not declined from thy law. 52 I remembered thy judgments of old, O LORD; and have comforted myself. 53 Horror hath taken hold upon me because of the wicked that forsake thy law. 54 Thy statutes have been my songs in the house of my pilgrimage. 55 I have remembered thy name, O LORD, in the night, and have kept thy law. 56 This I had, because I kept thy precepts. 57 Thou art my portion, O LORD: I have said that I would keep thy words. 58 I intreated thy favour with my whole heart: be merciful unto me according to thy word. 59 I thought on my ways, and turned my feet unto thy testimonies. 60 I made haste, and delayed not to keep thy commandments. 61 The bands of the wicked have robbed me: but I have not forgotten thy law. 62 At midnight I will rise to give thanks unto thee because of thy righteous judgments. 63 I am a companion of all them that fear thee, and of them that keep thy precepts. 64 The earth, O LORD, is full of thy mercy: teach me thy statutes. 65 Thou hast dealt well with thy servant, O LORD, according unto thy word. 66 Teach me good judgment and knowledge: for I have believed thy commandments. 67 Before I was afflicted I went astray: but now have I kept thy word. 68 Thou art good, and doest good; teach me thy statutes. 69 The proud have forged a lie against me: but I will keep thy precepts with my whole heart. 70 Their heart is as fat as grease; but I delight in thy law. 71 It is good for me that I have been afflicted; that I might learn thy statutes. 72 The law of thy mouth is better unto me than thousands of gold and silver. 73 Thy hands have made me and fashioned me: give me understanding, that I may learn thy commandments. 74 They that fear thee will be glad when they see me; because I have hoped in thy word. 75 I know, O LORD, that thy judgments are right, and that thou in faithfulness hast afflicted me. 76 Let, I pray thee, thy merciful kindness be for my comfort, according to thy word unto thy servant. 77 Let thy tender mercies come unto me, that I may live: for thy law is my delight. 78 Let the proud be ashamed; for they dealt perversely with me without a cause: but I will meditate in thy precepts. 79 Let those that fear thee turn unto me, and those that have known thy testimonies. 80 Let my heart be sound in thy statutes; that I be not ashamed. 81 My soul fainteth for thy salvation: but I hope in thy word. 82 Mine eyes fail for thy word, saying, When wilt thou comfort me? 83 For I am become like a bottle in the smoke; yet do I not forget thy statutes. 84 How many are the days of thy servant? when wilt thou execute judgment on them that persecute me? 85 The proud have digged pits for me, which are not after thy law. 86 All thy commandments are faithful: they persecute me wrongfully; help thou me. 87 They had almost consumed me upon earth; but I forsook not thy precepts. 88 Quicken me after thy lovingkindness; so shall I keep the testimony of thy mouth. 89 For ever, O LORD, thy word is settled in heaven. 90 Thy faithfulness is unto all generations: thou hast established the earth, and it abideth. 91 They continue this day according to thine ordinances: for all are thy servants. 92 Unless thy law had been my delights, I should then have perished in mine affliction. 93 I will never forget thy precepts: for with them thou hast quickened me. 94 I am thine, save me: for I have sought thy precepts. 95 The wicked have waited for me to destroy me: but I will consider thy testimonies. 96 I have seen an end of all perfection: but thy commandment is exceeding broad. 97 O how I love thy law! it is my meditation all the day. 98 Thou through thy commandments hast made me wiser than mine enemies: for they are ever with me. 99 I have more understanding than all my teachers: for thy testimonies are my meditation. 100 I understand more than the ancients, because I keep thy precepts. 101 I have refrained my feet from every evil way, that I might keep thy word. 102 I have not departed from thy judgments: for thou hast taught me. 103 How sweet are thy words unto my taste! yea, sweeter than honey to my mouth! 104 Through thy precepts I get understanding: therefore I hate every false way. 105 Thy word is a lamp unto my feet, and a light unto my path. 106 I have sworn, and I will perform it, that I will keep thy righteous judgments. 107 I am afflicted very much: quicken me, O LORD, according unto thy word. 108 Accept, I beseech thee, the freewill offerings of my mouth, O LORD, and teach me thy judgments. 109 My soul is continually in my hand: yet do I not forget thy law. 110 The wicked have laid a snare for me: yet I erred not from thy precepts. 111 Thy testimonies have I taken as an heritage for ever: for they are the rejoicing of my heart. 112 I have inclined mine heart to perform thy statutes alway, even unto the end. 113 I hate vain thoughts: but thy law do I love. 114 Thou art my hiding place and my shield: I hope in thy word. 115 Depart from me, ye evildoers: for I will keep the commandments of my God. 116 Uphold me according unto thy word, that I may live: and let me not be ashamed of my hope. 117 Hold thou me up, and I shall be safe: and I will have respect unto thy statutes continually. 118 Thou hast trodden down all them that err from thy statutes: for their deceit is falsehood. 119 Thou puttest away all the wicked of the earth like dross: therefore I love thy testimonies. 120 My flesh trembleth for fear of thee; and I am afraid of thy judgments. 121 I have done judgment and justice: leave me not to mine oppressors. 122 Be surety for thy servant for good: let not the proud oppress me. 123 Mine eyes fail for thy salvation, and for the word of thy righteousness. 124 Deal with thy servant according unto thy mercy, and teach me thy statutes. 125 I am thy servant; give me understanding, that I may know thy testimonies. 126 It is time for thee, LORD, to work: for they have made void thy law. 127 Therefore I love thy commandments above gold; yea, above fine gold. 128 Therefore I esteem all thy precepts concerning all things to be right; and I hate every false way. 129 Thy testimonies are wonderful: therefore doth my soul keep them. 130 The entrance of thy words giveth light; it giveth understanding unto the simple. 131 I opened my mouth, and panted: for I longed for thy commandments. 132 Look thou upon me, and be merciful unto me, as thou usest to do unto those that love thy name. 133 Order my steps in thy word: and let not any iniquity have dominion over me. 134 Deliver me from the oppression of man: so will I keep thy precepts. 135 Make thy face to shine upon thy servant; and teach me thy statutes. 136 Rivers of waters run down mine eyes, because they keep not thy law. 137 Righteous art thou, O LORD, and upright are thy judgments. 138 Thy testimonies that thou hast commanded are righteous and very faithful. 139 My zeal hath consumed me, because mine enemies have forgotten thy words. 140 Thy word is very pure: therefore thy servant loveth it. 141 I am small and despised: yet do not I forget thy precepts. 142 Thy righteousness is an everlasting righteousness, and thy law is the truth. 143 Trouble and anguish have taken hold on me: yet thy commandments are my delights. 144 The righteousness of thy testimonies is everlasting: give me understanding, and I shall live. 145 I cried with my whole heart; hear me, O LORD: I will keep thy statutes. 146 I cried unto thee; save me, and I shall keep thy testimonies. 147 I prevented the dawning of the morning, and cried: I hoped in thy word. 148 Mine eyes prevent the night watches, that I might meditate in thy word. 149 Hear my voice according unto thy lovingkindness: O LORD, quicken me according to thy judgment. 150 They draw nigh that follow after mischief: they are far from thy law. 151 Thou art near, O LORD; and all thy commandments are truth. 152 Concerning thy testimonies, I have known of old that thou hast founded them for ever. 153 Consider mine affliction, and deliver me: for I do not forget thy law. 154 Plead my cause, and deliver me: quicken me according to thy word. 155 Salvation is far from the wicked: for they seek not thy statutes. 156 Great are thy tender mercies, O LORD: quicken me according to thy judgments. 157 Many are my persecutors and mine enemies; yet do I not decline from thy testimonies. 158 I beheld the transgressors, and was grieved; because they kept not thy word. 159 Consider how I love thy precepts: quicken me, O LORD, according to thy lovingkindness. 160 Thy word is true from the beginning: and every one of thy righteous judgments endureth for ever. 161 Princes have persecuted me without a cause: but my heart standeth in awe of thy word. 162 I rejoice at thy word, as one that findeth great spoil. 163 I hate and abhor lying: but thy law do I love. 164 Seven times a day do I praise thee because of thy righteous judgments. 165 Great peace have they which love thy law: and nothing shall offend them. 166 LORD, I have hoped for thy salvation, and done thy commandments. 167 My soul hath kept thy testimonies; and I love them exceedingly. 168 I have kept thy precepts and thy testimonies: for all my ways are before thee. 169 Let my cry come near before thee, O LORD: give me understanding according to thy word. 170 Let my supplication come before thee: deliver me according to thy word. 171 My lips shall utter praise, when thou hast taught me thy statutes. 172 My tongue shall speak of thy word: for all thy commandments are righteousness. 173 Let thine hand help me; for I have chosen thy precepts. 174 I have longed for thy salvation, O LORD; and thy law is my delight. 175 Let my soul live, and it shall praise thee; and let thy judgments help me. 176 I have gone astray like a lost sheep; seek thy servant; for I do not forget thy commandments.
`

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

	// The column focused
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

	// passageView := viewport.New(50, 20)
	// passageView.Style = lipgloss.NewStyle().
	// 	BorderStyle(lipgloss.RoundedBorder()).
	// 	BorderForeground(lipgloss.Color("62")).
	// 	PaddingRight(2)

	// renderer, err := glamour.NewTermRenderer(
	// 	glamour.WithAutoStyle(),
	// 	glamour.WithWordWrap(width),
	// )

	// str, err := renderer.Render(p)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Init Passage
	// m.passage = passageView
	// m.passage.SetContent(str)

	// Init Translation
	m.columns[translationColumn].Title = "Translations"
	m.columns[translationColumn].FilterInput.Prompt = "Find Translation: "
	m.columns[translationColumn].SetStatusBarItemName("Translation", "Translations")
	m.columns[translationColumn].SetShowHelp(false)

	// Init Books
	m.columns[bookColumn].Title = "Books"
	m.columns[bookColumn].FilterInput.Prompt = "Find Book: "
	m.columns[bookColumn].SetStatusBarItemName("Book", "Books")

	// Init Chapters
	m.columns[chapterColumn].Title = "Chapters"
	m.columns[chapterColumn].FilterInput.Prompt = "Find Chapter: "
	m.columns[chapterColumn].SetStatusBarItemName("Chapter", "Chapters")
	m.columns[chapterColumn].SetShowHelp(false)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
		case "t":
			// to change to translation
			if m.focused == translationColumn {
				m.focused = bookColumn
			} else {
				m.focused = translationColumn
			}
		case "esc", "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
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
		// passageView := m.passage.View()
		switch m.focused {
		case translationColumn:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				translationView,
				bookView,
				chapterView,
			)
		case chapterColumn:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				// translationView,
				bookView,
				chapterView,
			)
		case passageColumn:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				// translationView,
				bookView,
				chapterView,
			)
		default: //BookColumn is the default
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				// translationView,
				bookView,
				chapterView,
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
