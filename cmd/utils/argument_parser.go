package utils

import "fmt"

const (
	SingleQuote = '\''
	DoubleQuote = '"'
	EmptySpace  = ' '
	Backslash   = '\\'
	Dollar      = '$'
	Newline     = '\n'
)

func isSpecialEscape(c byte) bool {
	return c == Newline || c == Dollar || c == DoubleQuote || c == Backslash
}

// Backslash enclosed in double quotes "\" preservers
// the special handling if followed by \, $, " or newline

func ParseArguments(userInput string) (params []string) {
	var current string

	for i := 0; i < len(userInput); i++ {
		char := rune(userInput[i])

		switch char {
		case SingleQuote:
			for i++; i < len(userInput) && userInput[i] != '\''; i++ {
				current += string(userInput[i])
			}
		case DoubleQuote:
			for i++; i < len(userInput) && userInput[i] != '"'; i++ {
				if userInput[i] == Backslash {
					if len(userInput) > i+1 && isSpecialEscape(userInput[i]) {
						i++
					} else {
						fmt.Println("Error: malformed input")
					}
				}

				current += string(userInput[i])
			}
		case EmptySpace:
			if current != "" {
				params = append(params, current)
			}

			current = ""
		case Backslash:
			// In this case we should escape the next character
			if i+1 < len(userInput) {
				i++
				current += string(userInput[i])
			}
		default:
			current += string(userInput[i])
		}
	}

	if current != "" {
		params = append(params, current)
	}

	return params
}
