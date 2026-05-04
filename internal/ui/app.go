package ui

import (
	"fmt"
	"robin/internal/model"
	"robin/internal/repository"

	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	viewList sessionState = iota
	viewReady
	viewExit
)

type Model struct {
	state    sessionState
	repo     *repository.SongRepository
	songs    []model.Song
	cursor   int
	Selected *model.Song
}

func InitialModel(repo *repository.SongRepository) Model {
	songs, _ := repo.FetchAllSongs()
	return Model{
		state: viewList,
		repo:  repo,
		songs: songs,
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.songs)-1 {
				m.cursor++
			}
		case "enter":
			switch m.state {
			case viewList:
				m.Selected = &m.songs[m.cursor]
				m.state = viewReady
			case viewReady:
				m.state = viewExit
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := ""
	switch m.state {
	case viewList:
		s += "\nSelect Song:\n\n"
		for i, song := range m.songs {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s - %s\n", cursor, song.Title, song.Artist)
		}
	case viewReady:
		s += fmt.Sprintf("\nReady to play: %s\n\n[ Press Enter to Play ]", m.Selected.Title)
	}
	return s + "\n\n(q: Quit)"
}