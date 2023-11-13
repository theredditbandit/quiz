# Quiz
A quiz system implemented in go.

# todo
- [x] Add support for setting quiz timers 
- [x] Add option to shuffle questions.
- [x] refactor again
- [x] Migrate to cobra instead of flags for arguments.
- [x] finish adding support for file schema validators (CSV)
- [ ] Add support for generate command to generate boilerplate for inputting questions
- [ ] Add support for JSON question format/schema to support MCQ questions
  - [ ] only one answer correct schema
  - [ ] Multiple correct answers schema
- [ ] do something to pretty print json (implement the schema command)
- [ ] Add support for nicer ui (using bubbles/bubbletea)
- [ ] Add support for per question time using tickers , display using progress bars
- [ ] Add support to skip questions entirely
- [ ] Add support to supply individual question marks , negative marks and different/custom marking schemes that may be declared in conf or question decleration.
- [ ] Add support to generate quiz reports
- [ ] Add support for reading configurations and question schemas from configuration files (ref spf13/viper)
- [ ] Change filenames to snake_case from camelCase or MixedCase
- [x] Unit Tests (optional)

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
