# Typing-test-app

Typing-test-app is a command-line application written in Go for measuring your typing speed and accuracy. The app generates random text using an online API, tracks your typing performance, and displays detailed results including accuracy, keystrokes per minute (PPM), words per minute (WPM), and error count.

## Features

- Generates random text for typing tests using the Random Word API.
- Tracks start and end time for each test.
- Calculates typing accuracy, keystrokes per minute (PPM), and words per minute (WPM).
- Counts errors between the expected and typed text.
- Displays a summary of your results after each test.
- Modular code structure for easy extension and maintenance.

## Project Structure

```
Typing-test-app/
├── src/
│   ├── main.go                # Entry point and main flow of the app
│   └── models/
│       └── test.go            # Data model and methods for typing test results
│   └── handlers/
│       └── typing.go          # Handlers for test flow and user interaction
│   └── utils/
│       └── helpers.go         # Utility functions (API calls, error counting, etc.)
│   └── config/
│       └── config.go          # Configuration management
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
└── README.md                  # Project documentation
```

## How to Run

1. Open a terminal and navigate to the `src` directory:
   ```
   cd src
   ```

2. Run the application using Go:
   ```
   go run main.go
   ```

This will start the typing test app in your terminal. You will be prompted with random text to type, and after you finish, your results will be displayed.

## Requirements

- Go 1.21 or higher
- Internet connection (for fetching random text from the API)

## Future Improvements

- Support for multiple languages
- Option to use local word lists if the API is unavailable
- Save and review previous test results
- Customizable test length and difficulty
- Improved error handling and user feedback
