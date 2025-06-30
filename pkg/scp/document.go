package scp

import (
	"fmt"
	"strings"
	"time"
)

// SCPDocument represents an SCP Foundation document
type SCPDocument struct {
	Number      string
	ObjectClass string
	Procedures  string
	Description string
	Addendum    []string
}

// Format returns the SCP document formatted with ASCII art header
func (d *SCPDocument) Format() string {
	header := `
███████╗ ██████╗██████╗ 
██╔════╝██╔════╝██╔══██╗
███████╗██║     ██████╔╝
╚════██║██║     ██╔═══╝ 
███████║╚██████╗██║     
╚══════╝ ╚═════╝╚═╝     

SCP FOUNDATION SECURE FACILITY - DIGITAL ANOMALIES DIVISION
═══════════════════════════════════════════════════════════

ITEM #: %s
OBJECT CLASS: %s

SPECIAL CONTAINMENT PROCEDURES: %s

DESCRIPTION: %s
`
	
	doc := fmt.Sprintf(header, d.Number, d.ObjectClass, d.Procedures, d.Description)
	
	if len(d.Addendum) > 0 {
		doc += "\n\nADDENDUM:\n"
		for i, add := range d.Addendum {
			doc += fmt.Sprintf("\n%s-%d: %s", d.Number, i+1, add)
		}
	}
	
	return doc
}

// GenerateIncidentReport creates an incident report for the given level and error
func GenerateIncidentReport(level int, errorMsg string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	
	report := fmt.Sprintf(`
════════════════════════════════════════════════════════════
INCIDENT REPORT - LEVEL %d CONTAINMENT BREACH
════════════════════════════════════════════════════════════

RESEARCHER: Dr. ████████
TIMESTAMP: %s
ANOMALY DETECTED: %s

RECOMMENDED ACTION: Review containment procedures and retry operation.

STATUS: AWAITING RESOLUTION

- Dr. ████████, Senior Containment Specialist
════════════════════════════════════════════════════════════
`, level, timestamp, errorMsg)
	
	return report
}

// FormatContainmentStatus creates a visual representation of containment status
func FormatContainmentStatus(status string, anomalyLevel int, sanity int) string {
	var statusColor, statusIcon string
	
	switch status {
	case "SECURE":
		statusColor = "🟢"
		statusIcon = "✓"
	case "BREACH":
		statusColor = "🟡"
		statusIcon = "⚠"
	case "CRITICAL":
		statusColor = "🔴"
		statusIcon = "✗"
	default:
		statusColor = "⚫"
		statusIcon = "?"
	}
	
	// Create progress bars
	anomalyBar := createProgressBar(anomalyLevel, 100, 20, "█", "░")
	sanityBar := createProgressBar(sanity, 100, 20, "█", "░")
	
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════╗
║ CONTAINMENT STATUS: %s %s %-8s                      ║
╠══════════════════════════════════════════════════════════╣
║ Anomaly Level:  [%s] %3d%%                  ║
║ Researcher Sanity: [%s] %3d%%                  ║
╚══════════════════════════════════════════════════════════╝`,
		statusColor, statusIcon, status,
		anomalyBar, anomalyLevel,
		sanityBar, sanity)
}

// createProgressBar creates a visual progress bar
func createProgressBar(current, max, width int, filled, empty string) string {
	percentage := float64(current) / float64(max)
	filledCount := int(percentage * float64(width))
	emptyCount := width - filledCount
	
	return strings.Repeat(filled, filledCount) + strings.Repeat(empty, emptyCount)
}

// FormatWarning creates an SCP-styled warning message
func FormatWarning(message string) string {
	return fmt.Sprintf(`
⚠️  ═══════════════════════════════════════════════════════
⚠️  WARNING: ANOMALOUS BEHAVIOR DETECTED
⚠️  ═══════════════════════════════════════════════════════
⚠️  %s
⚠️  ═══════════════════════════════════════════════════════
`, message)
}

// FormatSuccess creates an SCP-styled success message
func FormatSuccess(message string) string {
	return fmt.Sprintf(`
✅ ═══════════════════════════════════════════════════════
✅ CONTAINMENT SUCCESSFUL
✅ ═══════════════════════════════════════════════════════
✅ %s
✅ ═══════════════════════════════════════════════════════
`, message)
}

// FormatError creates an SCP-styled error message
func FormatError(message string) string {
	return fmt.Sprintf(`
🔴 ═══════════════════════════════════════════════════════
🔴 CONTAINMENT BREACH DETECTED
🔴 ═══════════════════════════════════════════════════════
🔴 %s
🔴 ═══════════════════════════════════════════════════════
`, message)
}

// GetSCPLogo returns the small SCP Foundation logo
func GetSCPLogo() string {
	return `███████╗ ██████╗██████╗ 
██╔════╝██╔════╝██╔══██╗
███████╗██║     ██████╔╝
╚════██║██║     ██╔═══╝ 
███████║╚██████╗██║     
╚══════╝ ╚═════╝╚═╝     `
}