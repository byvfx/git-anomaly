package ui

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/byvfx/git-anomaly/pkg/game"
	"github.com/byvfx/git-anomaly/pkg/scp"
)

var (
	// SCP Foundation color palette
	SCPRed      = color.New(color.FgRed, color.Bold)
	SCPOrange   = color.New(color.FgYellow, color.Bold)
	SCPGreen    = color.New(color.FgGreen)
	SCPBlue     = color.New(color.FgCyan)
	SCPGray     = color.New(color.FgHiBlack)
	SCPWhite    = color.New(color.FgWhite, color.Bold)
	
	// Status indicators
	StatusSecure   = SCPGreen
	StatusBreach   = SCPOrange
	StatusCritical = SCPRed
	
	// UI elements
	PromptColor    = color.New(color.FgCyan)
	CommandColor   = color.New(color.FgWhite, color.Bold)
	ErrorColor     = color.New(color.FgRed)
	SuccessColor   = color.New(color.FgGreen)
)

// Terminal represents the game terminal UI
type Terminal struct {
	CurrentLevel *game.Level
}

// NewTerminal creates a new terminal UI
func NewTerminal() *Terminal {
	return &Terminal{}
}

// DisplayWelcome shows the welcome screen
func (t *Terminal) DisplayWelcome() {
	fmt.Println()
	SCPWhite.Println(scp.GetSCPLogo())
	fmt.Println()
	SCPWhite.Println("SCP FOUNDATION SECURE FACILITY")
	SCPOrange.Println("Digital Anomalies Division")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println()
	fmt.Println("Welcome, Junior Researcher.")
	fmt.Println("Your assignment: Contain SCP-████ using Git protocols.")
	fmt.Println()
	SCPGray.Println("Type 'help' for available commands or 'start' to begin.")
	fmt.Println()
}

// DisplayGameStatus shows the current game status
func (t *Terminal) DisplayGameStatus(state *game.GameState) {
	fmt.Println(strings.Repeat("═", 60))
	
	// Containment status with color
	statusColor := getStatusColor(state.ContainmentStatus)
	statusColor.Printf("CONTAINMENT STATUS: %s", state.ContainmentStatus)
	
	// Current stats
	fmt.Printf(" | Branch: %s | Anomaly: %d%% | Sanity: %d%%\n", 
		state.CurrentBranch, state.AnomalyLevel, state.ResearcherSanity)
	
	// Working directory status
	if len(state.StagingArea) > 0 {
		SCPBlue.Println("\nSTAGED FOR CONTAINMENT:")
		for filename := range state.StagingArea {
			fmt.Printf("  📁 %s\n", filename)
		}
	}
	
	fmt.Println(strings.Repeat("═", 60))
}

// DisplayLevelIntro shows the level introduction
func (t *Terminal) DisplayLevelIntro(level *game.Level) {
	t.CurrentLevel = level
	
	fmt.Println()
	SCPWhite.Printf("════════ LEVEL %d: %s ════════\n", level.ID, level.Title)
	fmt.Println()
	
	// SCP document header
	SCPOrange.Printf("ITEM #: %s\n", level.SCPNumber)
	SCPOrange.Printf("OBJECT CLASS: %s\n", level.ObjectClass)
	fmt.Println()
	
	// Description
	fmt.Println("BRIEFING:")
	wrapText(level.Description, 60)
	fmt.Println()
	
	// Objective
	SCPBlue.Println("OBJECTIVE:")
	wrapText(level.Objective, 60)
	fmt.Println()
	
	// Files in working directory
	if len(level.InitialFiles) > 0 {
		SCPGray.Println("FILES DETECTED:")
		for filename := range level.InitialFiles {
			fmt.Printf("  • %s\n", filename)
		}
		fmt.Println()
	}
}

// DisplayIncidentReport shows an incident report
func (t *Terminal) DisplayIncidentReport(errorMsg string) {
	if t.CurrentLevel != nil {
		fmt.Print(scp.GenerateIncidentReport(t.CurrentLevel.ID, errorMsg))
	}
}

// DisplayCommandResult shows the result of a command
func (t *Terminal) DisplayCommandResult(result game.CommandResult) {
	// Display git output
	if result.Message != "" {
		fmt.Println(result.Message)
	}
	
	// Display SCP effect
	if result.SCPEffect != "" {
		fmt.Println()
		if result.Success {
			if strings.Contains(result.SCPEffect, "✅") {
				SuccessColor.Println(result.SCPEffect)
			} else if strings.Contains(result.SCPEffect, "⚠️") {
				SCPOrange.Println(result.SCPEffect)
			} else {
				fmt.Println(result.SCPEffect)
			}
		} else {
			ErrorColor.Println(result.SCPEffect)
		}
	}
	
	// Display stat changes
	if result.AnomalyDelta != 0 || result.SanityDelta != 0 {
		fmt.Println()
		if result.AnomalyDelta > 0 {
			SCPRed.Printf("⚠️  Anomaly Level +%d%%\n", result.AnomalyDelta)
		}
		if result.SanityDelta > 0 {
			SCPGreen.Printf("✓ Researcher Sanity +%d%%\n", result.SanityDelta)
		} else if result.SanityDelta < 0 {
			SCPOrange.Printf("⚠️  Researcher Sanity %d%%\n", result.SanityDelta)
		}
	}
}

// DisplayPrompt shows the command prompt
func (t *Terminal) DisplayPrompt(branch string) {
	if branch != "" {
		PromptColor.Printf("[SCP-████:%s] $ ", branch)
	} else {
		PromptColor.Print("[SCP-████] $ ")
	}
}

// DisplayHelp shows available commands
func (t *Terminal) DisplayHelp() {
	fmt.Println()
	SCPWhite.Println("FOUNDATION COMMAND REFERENCE:")
	fmt.Println(strings.Repeat("─", 60))
	
	commands := []struct {
		cmd  string
		desc string
	}{
		{"help", "Display this help message"},
		{"start", "Begin containment protocols"},
		{"status", "Check containment status"},
		{"git init", "Initialize containment repository"},
		{"git add <file>", "Stage files for containment"},
		{"git commit -m \"<msg>\"", "Secure files in containment"},
		{"git status", "View repository status"},
		{"git branch [name]", "Create or list containment branches"},
		{"git checkout <branch>", "Switch containment branches"},
		{"quit", "Exit containment protocols (progress saved)"},
	}
	
	for _, cmd := range commands {
		fmt.Printf("  %-25s %s\n", cmd.cmd, cmd.desc)
	}
	
	fmt.Println()
	SCPGray.Println("Note: Proper Git syntax required for containment success.")
	fmt.Println()
}

// DisplayError shows an error message
func (t *Terminal) DisplayError(message string) {
	ErrorColor.Printf("ERROR: %s\n", message)
}

// DisplaySuccess shows a success message
func (t *Terminal) DisplaySuccess(message string) {
	SuccessColor.Printf("SUCCESS: %s\n", message)
}

// ClearScreen clears the terminal screen
func (t *Terminal) ClearScreen() {
	// Print enough newlines to clear most terminals
	fmt.Print(strings.Repeat("\n", 50))
}

// Helper functions

func getStatusColor(status string) *color.Color {
	switch status {
	case "SECURE":
		return StatusSecure
	case "BREACH":
		return StatusBreach
	case "CRITICAL":
		return StatusCritical
	default:
		return SCPGray
	}
}

func wrapText(text string, width int) {
	words := strings.Fields(text)
	lineLength := 0
	
	for _, word := range words {
		if lineLength+len(word)+1 > width {
			fmt.Println()
			lineLength = 0
		}
		if lineLength > 0 {
			fmt.Print(" ")
			lineLength++
		}
		fmt.Print(word)
		lineLength += len(word)
	}
	fmt.Println()
}