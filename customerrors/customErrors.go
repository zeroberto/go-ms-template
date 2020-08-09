package customerrors

import "fmt"

// Report an error message
func Report(cause error, message string) string {
	if cause != nil && len(message) > 0 {
		return fmt.Sprintf("message=%s, erro=%v", message, cause)
	} else if cause != nil {
		return cause.Error()
	}
	return fmt.Sprintf(message)
}
