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