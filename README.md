# Quiz
A quiz system implemented in go.

<details open>
<summary>todo</summary>

# Features yet to be added
- [ ] Add support for setting quiz timers 
- [ ] Add support to read questions from google sheets
- [ ] Add keybind hints (eg keybinds to end test completely , keybinds to skip questions and so on)
- [ ] Add support for per question time using tickers , display using progress bars
- [ ] Add support to generate quiz reports.
- [ ] Add support for client server model with auth so people can authenticate with the server and then get questions from the server and also marks are sent back to the server securely
- [ ] implement auto generated gifs in the README using [VHS](https://github.com/charmbracelet/vhs)
- [ ] Testing
  - [ ] implement integration testing using [VHS](https://github.com/charmbracelet/vhs)
</details>


<details> 
<summary>Done</summary>

# Implemented Features
- [x] Add support to supply individual question marks , negative marks and different/custom marking schemes.
- [x] Add support to skip questions entirely
- [x] Add a UI using charm ecosystem (bubbles ,bubbletea , huh, lipgloss) 
- [x] Add option to shuffle questions.
- [x] refactor again
- [x] Migrate to cobra instead of flags for arguments.
- [x] finish adding support for file schema validators (CSV)
- [x] Add support for JSON question format/schema to support MCQ questions
  - [x] only one answer correct schema
  - [x] Multiple correct answers schema
- [x] Add support for generate command to generate boilerplate for inputting questions
- [x] add example flag to the generate command that generates a valid example.json of 5 questions
- [x] implement the json validator
- [x] remove support for CSV based questions , completely migrate to json based question format
</details>

# Installation

### Make sure you have [go](https://go.dev/) installed

### Clone the repo
    git clone https://github.com/theredditbandit/quiz.git
### Navigate into the cloned directory 
    cd quiz
### Download the dependencies
    go mod install
### Compile the binary
    go build .
### Or install system wide using
    go install 


# Pictures
![image](https://github.com/theredditbandit/quiz/assets/85390033/a8703257-b36d-43f2-87b3-a93800003ca8)

![image](https://github.com/theredditbandit/quiz/assets/85390033/b2d5e7da-aaa7-4196-835e-9e03f9870f8f)

![image](https://github.com/theredditbandit/quiz/assets/85390033/153b5555-f2e3-4c4e-acb7-4fcce081f16f)
