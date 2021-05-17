package controldisplay

import (
	"fmt"
	"log"
	"strings"
)

type ResultStatusRenderer struct {
	status string
}

func NewResultStatusRenderer(status string) *ResultStatusRenderer {
	return &ResultStatusRenderer{
		status: status,
	}
}

// Render returns the status
func (r ResultStatusRenderer) Render() string {
	log.Println("[TRACE] begin result status render")
	defer log.Println("[TRACE] end result status render")

	// pad status string to fixed width
	statusString := r.paddedStatusString()

	// get the color for our status
	colorFunc, ok := statusColors[r.status]

	if !ok {
		// for unrecognised status, just return nothing - we should be validating elsewhere
		return ""
	}
	// return status follow by colon and trailing space
	return fmt.Sprintf("%-5s%s ", colorFunc(statusString), colorStatusColon(":"))
}

// pad out status toi length of longest status string = "ERROR" - 5 chars
func (r ResultStatusRenderer) paddedStatusString() string {
	return fmt.Sprintf("%-5s", strings.ToUpper(r.status))
}
