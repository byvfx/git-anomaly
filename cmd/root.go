package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/byvfx/git-anomaly/pkg/game"
	"github.com/byvfx/git-anomaly/pkg/ui"
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
	// Add flags here if needed
}

func runGame(cmd *cobra.Command, args []string) {
	// Initialize game components
	engine := game.NewEngine()
	terminal := ui.NewTerminal()
	scanner := bufio.NewScanner(os.Stdin)
	
	// Display welcome screen
	terminal.ClearScreen()
	terminal.DisplayWelcome()
	
	// Main game loop
	gameStarted := false
	
	for {
		// Display prompt
		if gameStarted && engine.State.IsInitialized {
			terminal.DisplayPrompt(engine.State.CurrentBranch)
		} else {
			terminal.DisplayPrompt("")
		}
		
		// Read user input
		if !scanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(scanner.Text())
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
				fmt.Printf("\nProceed to Level %d? (y/n): ", nextLevel)
				if scanner.Scan() {
					response := strings.ToLower(strings.TrimSpace(scanner.Text()))
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