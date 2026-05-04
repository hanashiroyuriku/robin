package main

import (
	"fmt"
	"os"
	"robin/internal/model"
	"robin/internal/player"
	"robin/internal/repository"
	"robin/internal/ui"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	repo := repository.NewSongRepository("./songs")

	for {
		m := ui.InitialModel(repo)
		p := tea.NewProgram(m)

		finalModel, err := p.Run()
		if err != nil {
			fmt.Printf("\nError: %v", err)
			os.Exit(1)
		}

		updateModel := finalModel.(ui.Model)
		if updateModel.Selected == nil {
			fmt.Println("Good Bye!")
			break
		}
		playSong(updateModel.Selected)
	}
}

func playSong(s *model.Song) {
	fmt.Println("\033[H\033[2J")
	fmt.Printf("\n♪ %s - %s ♪\n", s.Title, s.Artist)
	fmt.Println("=========================================")

	for _, line := range s.Lyrics {
		player.PlayLysrics(line.Text, line.CharDelay)
		time.Sleep(time.Duration(line.LinePause * float64(time.Second)))
	}
	fmt.Println("\n-- Song Ended ---")
}