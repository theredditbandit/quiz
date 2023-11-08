# Quiz
A quiz system implemented in go.

# TODO
- [x] Add support for setting quiz timers 
- [x] Add option to shuffle questions.
- [x] refactor again
- [ ] Migrate to cobra instead of flags for arguments.
- [ ] Add support for per question time using tickers
- [ ] Add support for displaying individual time limit using profress bars ref (bubbles/bubbletea)
- [ ] Add support to skip questions entirely
- [ ] Add support for different question formats
  - [ ] Add support for JSON question format to support MCQ questions
    - [ ] only one answer correct schema
    - [ ] Multiple correct answers schema
  - [ ] `generate` command in the quiz app to generate boilerplate json for McQ and other type questions
  - [ ] Implement data validators that validate the schema of given problem.csv/json, find go equivalent of pydantic
- [ ] Add support for reading configurations and question schemas from configuration files (ref spf13/viper)
- [x] Unit Tests (optional)
  
# Pictures
![image](https://github.com/theredditbandit/quiz/assets/85390033/91371eda-324c-4c64-bb15-bea25ef3cbf5)
