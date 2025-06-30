package cmd

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
	"github.com/byvfx/git-anomaly/pkg/game"
	"github.com/byvfx/git-anomaly/pkg/ui"
)

var (
	useTUI bool
)

var rootCmd = &cobra.Command{
	Use:   "scp-git",
	Short: "SCP-████: The Self-Modifying Codebase - Git Learning Game",
	Long: `An interactive CLI game that teaches Git concepts through 
SCP Foundation horror/mystery narrative. You are a Foundation 
researcher tasked with containing anomalous digital entities 
using proper Git procedures.`,
	Run: runGame,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&useTUI, "tui", "t", true, "Use the modern TUI interface (default)")
	rootCmd.Flags().BoolP("classic", "c", false, "Use the classic CLI interface")
}

func runGame(cmd *cobra.Command, args []string) {
	// Check if classic mode is requested
	classic, _ := cmd.Flags().GetBool("classic")
	if classic {
		useTUI = false
	}
	
	if useTUI {
		// Use modern Bubble Tea TUI
		runTUIGame()
	} else {
		// Use classic CLI interface
		runClassicGame()
	}
}

func runTUIGame() {
	model := ui.NewModel()
	
	program := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error running TUI: %v\n", err)
		os.Exit(1)
	}
}

func runClassicGame() {
	// Initialize game components
	engine := game.NewEngine()
	terminal := ui.NewTerminal()
	
	// Display welcome screen
	terminal.ClearScreen()
	terminal.DisplayWelcome()
	
	// Configure readline
	rl, err := setupReadline(engine)
	if err != nil {
		fmt.Printf("Error setting up readline: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := rl.Close(); err != nil {
			fmt.Printf("Error closing readline: %v\n", err)
		}
	}()
	
	// Main game loop
	gameStarted := false
	
	for {
		// Set prompt
		prompt := "[SCP-████] $ "
		if gameStarted && engine.State.IsInitialized {
			prompt = fmt.Sprintf("[SCP-████:%s] $ ", engine.State.CurrentBranch)
		}
		rl.SetPrompt(prompt)
		
		// Read user input
		line, err := rl.Readline()
		if err != nil { // io.EOF or interrupt
			break
		}
		
		input := strings.TrimSpace(line)
		if input == "" {
			continue
		}
		
		// Handle meta commands
		switch input {
		case "quit", "exit", "q":
			fmt.Println("\nExiting containment protocols...")
			fmt.Println("Progress has been saved. The anomaly remains contained.")
			return
			
		case "help", "h", "?":
			terminal.DisplayHelp()
			continue
			
		case "clear", "cls":
			terminal.ClearScreen()
			if gameStarted {
				terminal.DisplayGameStatus(engine.State)
			}
			continue
			
		case "start":
			if !gameStarted {
				gameStarted = true
				err := engine.StartLevel(1)
				if err != nil {
					terminal.DisplayError(fmt.Sprintf("Failed to start level: %v", err))
					continue
				}
				terminal.DisplayLevelIntro(engine.CurrentLevel)
				terminal.DisplayGameStatus(engine.State)
			} else {
				terminal.DisplayError("Game already in progress")
			}
			continue
			
		case "status":
			if gameStarted {
				terminal.DisplayGameStatus(engine.State)
			} else {
				terminal.DisplayError("No game in progress. Type 'start' to begin.")
			}
			continue
			
		case "brief", "briefing", "objective", "objectives":
			if gameStarted && engine.CurrentLevel != nil {
				terminal.DisplayLevelIntro(engine.CurrentLevel)
			} else {
				terminal.DisplayError("No active containment protocol. Type 'start' to begin.")
			}
			continue
		}
		
		// Process game commands
		if !gameStarted {
			terminal.DisplayError("Type 'start' to begin containment protocols")
			continue
		}
		
		// Execute command
		result := engine.ProcessCommand(input)
		terminal.DisplayCommandResult(result)
		
		// Check for critical states
		if engine.State.AnomalyLevel >= 100 {
			terminal.DisplayError("CRITICAL CONTAINMENT BREACH!")
			terminal.DisplayIncidentReport("Anomaly level reached critical threshold")
			fmt.Println("\nGame Over. The anomaly has escaped containment.")
			return
		}
		
		if engine.State.ResearcherSanity <= 0 {
			terminal.DisplayError("RESEARCHER COMPROMISED!")
			fmt.Println("\nYou have been affected by the anomaly's influence.")
			fmt.Println("Report to Medical immediately.")
			return
		}
		
		// Check for level completion
		if engine.IsLevelComplete() {
			fmt.Println()
			terminal.DisplaySuccess("LEVEL COMPLETE!")
			fmt.Printf("Score: %d\n", engine.State.Score)
			
			nextLevel := engine.GetNextLevel()
			if nextLevel > 0 && nextLevel <= 3 {
				rl.SetPrompt(fmt.Sprintf("\nProceed to Level %d? (y/n): ", nextLevel))
				response, err := rl.Readline()
				if err == nil {
					response = strings.ToLower(strings.TrimSpace(response))
					if response == "y" || response == "yes" {
						err := engine.StartLevel(nextLevel)
						if err != nil {
							terminal.DisplayError(fmt.Sprintf("Failed to start level: %v", err))
						} else {
							terminal.ClearScreen()
							terminal.DisplayLevelIntro(engine.CurrentLevel)
							terminal.DisplayGameStatus(engine.State)
						}
					}
				}
			} else {
				fmt.Println("\nCongratulations! You have completed all available levels.")
				fmt.Println("Final Score:", engine.State.Score)
				fmt.Println("\nThank you for your service to the Foundation.")
				return
			}
		}
		
		// Show status if something significant changed
		if result.AnomalyDelta != 0 || result.SanityDelta != 0 {
			fmt.Println()
			terminal.DisplayGameStatus(engine.State)
		}
	}
}

// setupReadline configures readline with tab completion and command history
func setupReadline(engine *game.Engine) (*readline.Instance, error) {
	// Define completer for tab completion
	completer := readline.NewPrefixCompleter(
		// Meta commands
		readline.PcItem("help"),
		readline.PcItem("start"),
		readline.PcItem("status"),
		readline.PcItem("brief"),
		readline.PcItem("briefing"),
		readline.PcItem("objective"),
		readline.PcItem("objectives"),
		readline.PcItem("breathe"),
		readline.PcItem("clear"),
		readline.PcItem("quit"),
		readline.PcItem("exit"),
		
		// Git commands
		readline.PcItem("git",
			readline.PcItem("config",
				readline.PcItem("user.name"),
				readline.PcItem("user.email"),
			),
			readline.PcItem("init"),
			readline.PcItem("add",
				readline.PcItemDynamic(func(line string) []string {
					// Dynamic completion for filenames
					if engine.State == nil || engine.State.WorkingDir == nil {
						return []string{}
					}
					
					var files []string
					for filename := range engine.State.WorkingDir {
						files = append(files, filename)
					}
					// Also add common wildcards
					files = append(files, ".", "*")
					return files
				}),
			),
			readline.PcItem("commit",
				readline.PcItem("-m"),
				readline.PcItem("-a"),
			),
			readline.PcItem("status"),
			readline.PcItem("diff"),
			readline.PcItem("log",
				readline.PcItem("-p"),
			),
			readline.PcItem("show"),
			readline.PcItem("branch"),
			readline.PcItem("checkout",
				readline.PcItemDynamic(func(line string) []string {
					// Dynamic completion for branch names
					if engine.State == nil || engine.State.Branches == nil {
						return []string{}
					}
					
					var branches []string
					for branch := range engine.State.Branches {
						branches = append(branches, branch)
					}
					branches = append(branches, "-b")
					return branches
				}),
			),
			readline.PcItem("switch",
				readline.PcItem("-c"),
				readline.PcItemDynamic(func(line string) []string {
					// Dynamic completion for branch names
					if engine.State == nil || engine.State.Branches == nil {
						return []string{}
					}
					
					var branches []string
					for branch := range engine.State.Branches {
						branches = append(branches, branch)
					}
					return branches
				}),
			),
			readline.PcItem("merge",
				readline.PcItemDynamic(func(line string) []string {
					// Dynamic completion for branch names
					if engine.State == nil || engine.State.Branches == nil {
						return []string{}
					}
					
					var branches []string
					for branch := range engine.State.Branches {
						if branch != engine.State.CurrentBranch {
							branches = append(branches, branch)
						}
					}
					return branches
				}),
			),
		),
	)
	
	// Configure readline
	config := &readline.Config{
		Prompt:              "[SCP-████] $ ",
		HistoryFile:         "/tmp/scp-git-history",
		AutoComplete:        completer,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	}
	
	return readline.NewEx(config)
}

// filterInput allows certain special characters in input
func filterInput(r rune) (rune, bool) {
	switch r {
	// Allow all printable characters
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}