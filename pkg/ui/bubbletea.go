package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/byvfx/git-anomaly/pkg/game"
)

// Messages for animations
type tickMsg time.Time
type bootCompleteMsg struct{}

// Styles - CRT Terminal Theme
var (
	// CRT Color scheme (classic green terminal)
	crtGreen      = lipgloss.Color("#00FF41")    // Bright matrix green
	crtDarkGreen  = lipgloss.Color("#008F11")    // Darker green
	crtAmber      = lipgloss.Color("#FFB000")    // Amber warning
	crtRed        = lipgloss.Color("#FF3030")    // Bright red
	crtBlue       = lipgloss.Color("#00AAFF")    // Cyan blue
	crtGray       = lipgloss.Color("#404040")    // Dark gray
	crtBackground = lipgloss.Color("#0D1117")    // Very dark background
	crtBorder     = lipgloss.Color("#1F2937")    // Border color
	
	// CRT Terminal Styles
	crtScreenStyle = lipgloss.NewStyle().
		Background(crtBackground).
		Foreground(crtGreen).
		Padding(1, 2).
		Border(lipgloss.ThickBorder()).
		BorderForeground(crtBorder)
	
	// Header styles - CRT themed
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(crtGreen).
		Background(crtBackground).
		Align(lipgloss.Center).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(crtGreen).
		Padding(0, 2).
		Margin(0, 0, 1, 0)
	
	crtLogoStyle = lipgloss.NewStyle().
		Foreground(crtDarkGreen).
		Align(lipgloss.Center).
		Margin(1, 0)
	
	// Status bar - CRT green glow
	statusBarStyle = lipgloss.NewStyle().
		Background(crtBackground).
		Foreground(crtGreen).
		Border(lipgloss.NormalBorder()).
		BorderForeground(crtGreen).
		Padding(0, 1).
		Bold(true)
	
	// Containment status styles
	secureStyle = lipgloss.NewStyle().
		Foreground(crtGreen).
		Bold(true)
	
	breachStyle = lipgloss.NewStyle().
		Foreground(crtAmber).
		Bold(true)
	
	criticalStyle = lipgloss.NewStyle().
		Foreground(crtRed).
		Bold(true).
		Blink(true)
	
	// Content styles - CRT themed
	promptStyle = lipgloss.NewStyle().
		Foreground(crtGreen).
		Bold(true)
	
	successStyle = lipgloss.NewStyle().
		Foreground(crtGreen)
	
	errorStyle = lipgloss.NewStyle().
		Foreground(crtRed)
	
	warningStyle = lipgloss.NewStyle().
		Foreground(crtAmber)
	
	// Terminal output window
	terminalStyle = lipgloss.NewStyle().
		Background(crtBackground).
		Foreground(crtGreen).
		Border(lipgloss.NormalBorder()).
		BorderForeground(crtDarkGreen).
		Padding(1).
		Height(12)
	
	// Input line - CRT style
	inputLineStyle = lipgloss.NewStyle().
		Background(crtBackground).
		Foreground(crtGreen).
		Border(lipgloss.NormalBorder()).
		BorderForeground(crtGreen).
		Padding(0, 1)
	
	// Help panel - retro computer manual style
	helpStyle = lipgloss.NewStyle().
		Background(crtBackground).
		Foreground(crtDarkGreen).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(crtGreen).
		Padding(1).
		Margin(1, 0)
	
	// Scanline effect simulation
	scanlineStyle = lipgloss.NewStyle().
		Foreground(crtGray).
		Faint(true)
)

// Model represents the Bubble Tea model for the game
type Model struct {
	engine       *game.Engine
	input        string
	gameStarted  bool
	showHelp     bool
	output       []string
	currentLevel *game.Level
	width        int
	height       int
	
	// CRT Effects
	cursorBlink   bool
	scanlinePos   int
	glitchMode    bool
	bootSequence  bool
	bootStep      int
	frameCount    int
}

// NewModel creates a new Bubble Tea model
func NewModel() Model {
	return Model{
		engine:       game.NewEngine(),
		input:        "",
		gameStarted:  false,
		showHelp:     false,
		output:       []string{},
		width:        80,
		height:       24,
		bootSequence: true,
		bootStep:     0,
		frameCount:   0,
	}
}

// Animation commands
func tickEvery() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func bootComplete() tea.Cmd {
	return tea.Tick(time.Second*3, func(t time.Time) tea.Msg {
		return bootCompleteMsg{}
	})
}

// Init implements tea.Model
func (m Model) Init() tea.Cmd {
	return tea.Batch(tickEvery(), bootComplete())
}

// Update implements tea.Model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		// Handle animations
		m.frameCount++
		m.cursorBlink = m.frameCount%10 < 5
		m.scanlinePos = (m.scanlinePos + 1) % 40
		
		// Check for anomaly level effects
		if m.gameStarted && m.engine.State.AnomalyLevel > 70 {
			m.glitchMode = m.frameCount%20 < 3 // Glitch effect when anomaly high
		} else {
			m.glitchMode = false
		}
		
		return m, tickEvery()
		
	case bootCompleteMsg:
		m.bootSequence = false
		return m, nil
		
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit // Always allow Ctrl+C to exit
			
		case "q":
			if !m.gameStarted {
				return m, tea.Quit
			}
			// In game, 'q' should be typed normally
			m.input += "q"
			return m, nil
			
		case "ctrl+h":
			m.showHelp = !m.showHelp
			return m, nil
			
		case "enter":
			if m.input == "" {
				return m, nil
			}
			
			// Process the command
			m.output = append(m.output, fmt.Sprintf("> %s", m.input))
			result := (&m).processCommand(m.input)
			
			// Handle special quit message
			if result == "quit_game" {
				return m, tea.Quit
			}
			
			// Add result to output
			if result != "" {
				lines := strings.Split(result, "\n")
				m.output = append(m.output, lines...)
			}
			
			// Keep only last 15 lines of output
			if len(m.output) > 15 {
				m.output = m.output[len(m.output)-15:]
			}
			
			m.input = ""
			return m, nil
			
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			return m, nil
			
		default:
			// Add character to input
			if len(msg.String()) == 1 {
				m.input += msg.String()
			}
			return m, nil
		}
	}
	
	return m, nil
}

// View implements tea.Model
func (m Model) View() string {
	if m.bootSequence {
		return m.renderBootSequence()
	}
	
	var sections []string
	
	// CRT Terminal Border (simulate the monitor bezel)
	terminalWidth := min(m.width-4, 100)
	
	// Title bar with CRT glow effect
	title := m.renderCRTTitle()
	sections = append(sections, title)
	
	// System status line
	if m.gameStarted {
		status := m.renderCRTStatus()
		sections = append(sections, status)
	}
	
	// Main terminal window
	terminalContent := m.renderTerminalWindow(terminalWidth)
	sections = append(sections, terminalContent)
	
	// Input line with cursor
	inputLine := m.renderCRTInput()
	sections = append(sections, inputLine)
	
	// Footer with scanline effect
	footer := m.renderCRTFooter()
	sections = append(sections, footer)
	
	// Apply overall CRT screen styling
	screen := lipgloss.JoinVertical(lipgloss.Left, sections...)
	
	// Add glitch effect if anomaly level is high
	if m.glitchMode {
		screen = m.applyGlitchEffect(screen)
	}
	
	return crtScreenStyle.Width(terminalWidth + 4).Render(screen)
}

// Helper methods

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m Model) renderBootSequence() string {
	scpLogo := `
███████╗ ██████╗██████╗ 
██╔════╝██╔════╝██╔══██╗
███████╗██║     ██████╔╝
╚════██║██║     ██╔═══╝ 
███████║╚██████╗██║     
╚══════╝ ╚═════╝╚═╝     

SCP FOUNDATION SECURE FACILITY
Digital Anomalies Division`

	bootLines := []string{
		scpLogo,
		"",
		"SECURE TERMINAL v2.1.7",
		"Copyright (c) 19██ SCP Foundation",
		"",
		"Initializing containment protocols...",
		"Loading anomaly detection systems... OK",
		"Establishing secure connection... OK", 
		"Verifying researcher credentials... OK",
		"",
		"SYSTEM READY",
		"",
		"Press any key to continue...",
	}
	
	// Show progressive boot sequence
	visibleLines := min(m.bootStep*2, len(bootLines))
	if m.frameCount%5 == 0 && visibleLines < len(bootLines) {
		m.bootStep++
	}
	
	display := strings.Join(bootLines[:visibleLines], "\n")
	
	bootStyle := lipgloss.NewStyle().
		Foreground(crtGreen).
		Background(crtBackground).
		Padding(2).
		Border(lipgloss.ThickBorder()).
		BorderForeground(crtGreen).
		Align(lipgloss.Center).
		Width(60)
	
	return bootStyle.Render(display)
}

func (m Model) renderCRTTitle() string {
	title := "SCP-████: THE SELF-MODIFYING CODEBASE"
	
	titleBar := lipgloss.NewStyle().
		Foreground(crtGreen).
		Background(crtBackground).
		Bold(true).
		Align(lipgloss.Center).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(crtGreen).
		Padding(0, 1)
	
	return titleBar.Render(title)
}

func (m Model) renderCRTStatus() string {
	if !m.gameStarted {
		return ""
	}
	
	state := m.engine.State
	
	// Status indicators with CRT styling
	var statusColor lipgloss.Style
	var statusIcon string
	
	switch state.ContainmentStatus {
	case "SECURE":
		statusColor = secureStyle
		statusIcon = "●"
	case "BREACH":
		statusColor = breachStyle  
		statusIcon = "▲"
	case "CRITICAL":
		statusColor = criticalStyle
		statusIcon = "■"
	}
	
	statusText := fmt.Sprintf(
		"STATUS: %s %s │ Branch: %s │ Anomaly: %d%% │ Sanity: %d%% │ Score: %d",
		statusColor.Render(statusIcon),
		statusColor.Render(state.ContainmentStatus),
		state.CurrentBranch,
		state.AnomalyLevel,
		state.ResearcherSanity,
		state.Score,
	)
	
	return statusBarStyle.Render(statusText)
}

func (m Model) renderTerminalWindow(width int) string {
	var content []string
	
	if !m.gameStarted {
		// Welcome screen with CRT ASCII
		content = []string{
			"╔══════════════════════════════════════════════════════════╗",
			"║  FOUNDATION SECURE WORKSTATION - DIGITAL ANOMALIES DIV  ║",
			"╚══════════════════════════════════════════════════════════╝",
			"",
			"Welcome, Junior Researcher.",
			"Your assignment: Contain SCP-████ using Git protocols.",
			"",
			"┌─ AVAILABLE COMMANDS ─┐",
			"│ start - Begin mission │",
			"│ help  - Show commands │", 
			"│ quit  - Exit terminal │",
			"└───────────────────────┘",
			"",
			"Awaiting input...",
		}
	} else {
		// Show level info like classic CLI
		if m.currentLevel != nil {
			content = append(content, 
				fmt.Sprintf("════════ LEVEL %d: %s ════════", m.currentLevel.ID, m.currentLevel.Title),
				"",
				fmt.Sprintf("ITEM #: %s", m.currentLevel.SCPNumber),
				fmt.Sprintf("OBJECT CLASS: %s", m.currentLevel.ObjectClass),
				"",
				"BRIEFING:",
				m.currentLevel.Description,
				"",
				"OBJECTIVE:",
				m.currentLevel.Objective,
				"",
			)
			
			// Show initial files like classic CLI
			if len(m.currentLevel.InitialFiles) > 0 {
				content = append(content, "FILES DETECTED:")
				for filename := range m.currentLevel.InitialFiles {
					content = append(content, fmt.Sprintf("  • %s", filename))
				}
				content = append(content, "")
			}
		}
		
		// Show command output with CRT styling
		if len(m.output) > 0 {
			content = append(content, "┌─ TERMINAL OUTPUT ─┐")
			for _, line := range m.output {
				content = append(content, "│ " + line)
			}
			content = append(content, "└───────────────────┘")
		}
	}
	
	// Pad content to fill terminal window
	for len(content) < 15 {
		content = append(content, "")
	}
	
	// Apply terminal styling
	terminalContent := strings.Join(content, "\n")
	
	return terminalStyle.Width(width).Render(terminalContent)
}

func (m Model) renderCRTInput() string {
	var promptText string
	if m.gameStarted && m.engine.State.IsInitialized {
		promptText = fmt.Sprintf("SCP-████:%s>", m.engine.State.CurrentBranch)
	} else {
		promptText = "SCP-████>"
	}
	
	// Add blinking cursor
	cursor := "█"
	if !m.cursorBlink {
		cursor = " "
	}
	
	inputText := m.input + cursor
	
	prompt := promptStyle.Render(promptText + " ")
	input := lipgloss.NewStyle().Foreground(crtGreen).Render(inputText)
	
	return inputLineStyle.Render(prompt + input)
}

func (m Model) renderCRTFooter() string {
	if m.showHelp {
		return m.renderCRTHelp()
	}
	
	// Simulate scanlines and CRT info
	scanline := strings.Repeat("─", m.scanlinePos) + "█" + strings.Repeat("─", 40-m.scanlinePos)
	footer := scanlineStyle.Render(scanline)
	
	helpHint := lipgloss.NewStyle().
		Foreground(crtDarkGreen).
		Render("Press [Ctrl+H] for help │ [Ctrl+C] to exit")
	
	return lipgloss.JoinVertical(lipgloss.Left, footer, helpHint)
}

func (m Model) renderCRTHelp() string {
	helpContent := `╔═══════════════════ FOUNDATION MANUAL ═══════════════════╗
║                     COMMAND REFERENCE                   ║
╠══════════════════════════════════════════════════════════╣
║ SYSTEM COMMANDS:                                        ║
║   help                    Show this manual              ║
║   start                   Begin containment protocols   ║
║   status                  Check system status           ║
║   brief/briefing          Re-display level briefing    ║
║   breathe                 Take a moment to recover      ║
║   clear                   Clear terminal output         ║
║   next                    Proceed to next level         ║
║   quit                    Exit secure session           ║
║                                                          ║
║ GIT CONTAINMENT PROTOCOLS:                              ║
║   git init                Initialize repository         ║
║   git add <file>          Stage files for containment   ║
║   git commit -m "msg"     Secure staged files           ║
║   git status              View repository status        ║
║   git branch [name]       Manage containment branches   ║
║   git checkout <branch>   Switch active branch          ║
║   git switch <branch>     Switch to existing branch     ║
║   git switch -c <branch>  Create and switch new branch  ║
║                                                          ║
║ CONTROLS:                                               ║
║   [Ctrl+H]               Toggle this help panel        ║
║   [Ctrl+C]               Emergency exit                 ║
║   [Enter]                Execute command                ║
╚══════════════════════════════════════════════════════════╝`
	
	return helpStyle.Render(helpContent)
}

func (m Model) createCRTProgressBar(current, max, width int) string {
	percentage := float64(current) / float64(max)
	filled := int(percentage * float64(width))
	
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return bar
}

func (m Model) applyGlitchEffect(content string) string {
	// Simple glitch effect - replace some characters randomly
	lines := strings.Split(content, "\n")
	glitchChars := []string{"█", "▓", "▒", "░", "▀", "▄", "■", "□"}
	
	for i, line := range lines {
		if m.frameCount%7 == i%7 { // Glitch different lines on different frames
			runes := []rune(line)
			for j := range runes {
				if m.frameCount%13 == j%13 { // Random-ish replacement
					if len(glitchChars) > 0 {
						runes[j] = []rune(glitchChars[j%len(glitchChars)])[0]
					}
				}
			}
			lines[i] = string(runes)
		}
	}
	
	return strings.Join(lines, "\n")
}

func (m *Model) processCommand(input string) string {
	input = strings.TrimSpace(input)
	
	// Handle meta commands
	switch input {
	case "quit", "exit":
		return "quit_game" // Special message to trigger exit
		
	case "help":
		m.showHelp = !m.showHelp
		return ""
		
	case "clear":
		m.output = []string{}
		return ""
		
	case "start":
		if !m.gameStarted {
			m.gameStarted = true
			err := m.engine.StartLevel(1)
			if err != nil {
				return fmt.Sprintf("ERROR: Failed to start level: %v", err)
			}
			m.currentLevel = m.engine.CurrentLevel
			// Clear output to show fresh level info
			m.output = []string{}
			return ""
		}
		return "ERROR: Game already in progress"
		
	case "status":
		if m.gameStarted {
			return m.formatGameStatus()
		}
		return "ERROR: No game in progress. Type 'start' to begin."
	}
	
	// Process game commands
	if !m.gameStarted {
		return "ERROR: Type 'start' to begin containment protocols"
	}
	
	result := m.engine.ProcessCommand(input)
	
	// Check for critical states first
	if m.engine.State.AnomalyLevel >= 100 {
		return "🔴 CRITICAL CONTAINMENT BREACH!\nAnomaly level reached critical threshold.\nGame Over. The anomaly has escaped containment.\nType 'quit' to exit."
	}
	
	if m.engine.State.ResearcherSanity <= 0 {
		return "🔴 RESEARCHER COMPROMISED!\nYou have been affected by the anomaly's influence.\nReport to Medical immediately.\nType 'quit' to exit."
	}
	
	// Check for level completion
	if m.engine.IsLevelComplete() {
		nextLevel := m.engine.GetNextLevel()
		completionMsg := fmt.Sprintf("✅ LEVEL COMPLETE! Score: %d", m.engine.State.Score)
		
		if nextLevel > 0 && nextLevel <= 3 {
			completionMsg += fmt.Sprintf("\n\nProceed to Level %d? Type 'next' to continue or 'stay' to practice more.", nextLevel)
		} else {
			completionMsg += "\n\n🎉 Congratulations! You have completed all available levels."
		}
		
		return completionMsg
	}
	
	// Handle next level progression
	if input == "next" && m.engine.IsLevelComplete() {
		nextLevel := m.engine.GetNextLevel()
		if nextLevel > 0 && nextLevel <= 3 {
			err := m.engine.StartLevel(nextLevel)
			if err != nil {
				return fmt.Sprintf("ERROR: Failed to start level: %v", err)
			}
			m.currentLevel = m.engine.CurrentLevel
			return fmt.Sprintf("✅ LEVEL %d STARTED: %s", nextLevel, m.currentLevel.Title)
		}
	}
	
	return m.formatCommandResult(result)
}

func (m Model) renderStatusBar() string {
	if !m.gameStarted {
		return ""
	}
	
	state := m.engine.State
	
	// Containment status with appropriate color
	var statusStyle lipgloss.Style
	switch state.ContainmentStatus {
	case "SECURE":
		statusStyle = secureStyle
	case "BREACH":
		statusStyle = breachStyle
	case "CRITICAL":
		statusStyle = criticalStyle
	default:
		statusStyle = lipgloss.NewStyle().Foreground(crtGray)
	}
	
	statusText := statusStyle.Render(state.ContainmentStatus)
	
	// Progress bars
	anomalyBar := m.createProgressBar(state.AnomalyLevel, 100, 20, "█", "░")
	sanityBar := m.createProgressBar(state.ResearcherSanity, 100, 20, "█", "░")
	
	statusContent := fmt.Sprintf(
		"Status: %s | Branch: %s | Anomaly: [%s] %d%% | Sanity: [%s] %d%%",
		statusText,
		state.CurrentBranch,
		anomalyBar, state.AnomalyLevel,
		sanityBar, state.ResearcherSanity,
	)
	
	return statusBarStyle.Render(statusContent)
}

func (m Model) renderLevelInfo() string {
	if m.currentLevel == nil {
		return ""
	}
	
	levelStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(crtAmber).
		Padding(1).
		Margin(1, 0)
	
	content := fmt.Sprintf("LEVEL %d: %s\nItem #: %s | Class: %s\n\n%s",
		m.currentLevel.ID,
		m.currentLevel.Title,
		m.currentLevel.SCPNumber,
		m.currentLevel.ObjectClass,
		m.currentLevel.Objective,
	)
	
	return levelStyle.Render(content)
}

func (m Model) renderPrompt() string {
	var promptText string
	if m.gameStarted && m.engine.State.IsInitialized {
		promptText = fmt.Sprintf("[SCP-████:%s] $ ", m.engine.State.CurrentBranch)
	} else {
		promptText = "[SCP-████] $ "
	}
	
	prompt := promptStyle.Render(promptText)
	input := inputLineStyle.Render(m.input + "█") // Add cursor
	
	return lipgloss.JoinHorizontal(lipgloss.Top, prompt, input)
}

func (m Model) renderHelp() string {
	helpContent := `FOUNDATION COMMAND REFERENCE:
─────────────────────────────
help                    Display this help message
start                   Begin containment protocols  
status                  Check containment status
git init                Initialize containment repository
git add <file>          Stage files for containment
git commit -m "msg"     Secure files in containment
git status              View repository status
git branch [name]       Create or list containment branches
git checkout <branch>   Switch containment branches
clear                   Clear output history
quit                    Exit containment protocols

CONTROLS:
─────────
Ctrl+H                  Toggle this help
Ctrl+C                  Emergency exit (main menu only)
Enter                   Execute command`
	
	crtHelpStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(crtBlue).
		Padding(1).
		Foreground(crtGreen).
		Background(crtBackground)
	
	return crtHelpStyle.Render(helpContent)
}

func (m Model) createProgressBar(current, max, width int, filled, empty string) string {
	percentage := float64(current) / float64(max)
	filledCount := int(percentage * float64(width))
	emptyCount := width - filledCount
	
	return strings.Repeat(filled, filledCount) + strings.Repeat(empty, emptyCount)
}

func (m Model) formatGameStatus() string {
	state := m.engine.State
	return fmt.Sprintf(`Current Status:
- Containment: %s
- Anomaly Level: %d%%
- Researcher Sanity: %d%%
- Current Branch: %s
- Score: %d`,
		state.ContainmentStatus,
		state.AnomalyLevel,
		state.ResearcherSanity,
		state.CurrentBranch,
		state.Score,
	)
}

func (m Model) formatCommandResult(result game.CommandResult) string {
	var output []string
	
	if result.Message != "" {
		output = append(output, result.Message)
	}
	
	if result.SCPEffect != "" {
		var style lipgloss.Style
		if result.Success {
			if strings.Contains(result.SCPEffect, "✅") {
				style = successStyle
			} else if strings.Contains(result.SCPEffect, "⚠️") {
				style = warningStyle
			} else {
				style = lipgloss.NewStyle()
			}
		} else {
			style = errorStyle
		}
		output = append(output, style.Render(result.SCPEffect))
	}
	
	// Show stat changes
	if result.AnomalyDelta != 0 || result.SanityDelta != 0 {
		var changes []string
		if result.AnomalyDelta > 0 {
			changes = append(changes, errorStyle.Render(fmt.Sprintf("Anomaly Level +%d%%", result.AnomalyDelta)))
		}
		if result.SanityDelta > 0 {
			changes = append(changes, successStyle.Render(fmt.Sprintf("Researcher Sanity +%d%%", result.SanityDelta)))
		} else if result.SanityDelta < 0 {
			changes = append(changes, warningStyle.Render(fmt.Sprintf("Researcher Sanity %d%%", result.SanityDelta)))
		}
		
		if len(changes) > 0 {
			output = append(output, strings.Join(changes, " | "))
		}
	}
	
	return strings.Join(output, "\n")
}