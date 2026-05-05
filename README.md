# Robin
Robin is a command-line interface (CLI) application developed using the Go programming language. This software is designed to render song lyrics within a terminal environment using a synchronized typewriter effect, configurable via external data structures.

## Technical Description and Inspiration
The project is named Robin, inspired by the character of the same name from the Penacony region in Honkai: Star Rail. Reflecting the character's thematic association with music and performance, this application serves as a technical medium for lyrical display.

From an architectural standpoint, the system implements a strict separation of concerns. It utilizes a Repository pattern for data abstraction and a Model-View-Update (MVU) architecture via the `Bubble Tea ` library to manage the terminal user interface (TUI).

## System Requirements
- Go 1.21 or later.

- A terminal emulator with support for ANSI Escape Sequences.

## Installation
1. Clone the repository:
``` bash
git clone https://github.com/username/robin.git
```
2. Navigate to the project directory:
```bash
cd robin
```
3. Synchronize dependencies:
```bash
go mod tidy
```

## Execution
To initiate the application, execute the following command from the root directory:

### Operational Controls:
1. Arrow Up / J: Move cursor upward
2. Arrow Down / K: Move cursor downward
3. Enter: Confirm selection and initiate the lyric engine
4. Q / Ctrl+C: terminate the application

## Data Configuration
The application dynamically scans the `songs/` directory for JSON files. New song data must adhere to the following schema:
```json
{
  "id": "unique-slug",
  "title": "Song Title",
  "artist": "Artist Name",
  "lyrics": [
    {
      "text": "Lyric line",
      "char_delay": 0.08,
      "line_pause": 0.5
    }
  ]
}
```

### Parameters:

- char_delay: Delay between character rendering (seconds).

- line_pause: Delay after a complete line is rendered (seconds).

## Project Structure
```
robin/
├── internal/
│   ├── model/           # Core data entities
│   ├── player/          # Typewriter rendering engine
│   ├── repository/      # File system data retrieval logic
│   └── ui/              # TUI state management and Lipgloss styling
├── songs/               # JSON lyric data storage
├── main.go              # Application entry point
├── go.mod               # Dependency management
└── LICENSE              # MIT License information
```

## Licensing Information
This project is distributed under the MIT License. A software license is a legal instrument providing guidelines for the use and redistribution of software. The MIT License is a permissive license that permits reuse within proprietary software provided all copies include the original copyright notice and this permission notice.

## Acknowledgements
- Hanashiro Yuriku
- Gemini