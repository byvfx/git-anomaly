package game

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// Engine represents the main game engine
type Engine struct {
	State        *GameState
	CurrentLevel *Level
	LevelNum     int
}

// NewEngine creates a new game engine
func NewEngine() *Engine {
	return &Engine{
		State:    NewGameState(),
		LevelNum: 1,
	}
}

// StartLevel begins a specific level
func (e *Engine) StartLevel(levelNum int) error {
	level := GetLevel(levelNum)
	if level == nil {
		return fmt.Errorf("level %d not found", levelNum)
	}
	
	e.CurrentLevel = level
	e.LevelNum = levelNum
	e.State.CurrentLevel = levelNum
	
	// Initialize working directory with level files
	e.State.WorkingDir = make(map[string]FileState)
	for filename, content := range level.InitialFiles {
		e.State.WorkingDir[filename] = FileState{
			Content:  content,
			Modified: false,
			Staged:   false,
			Hash:     hashContent(content),
		}
	}
	
	return nil
}

// ProcessCommand parses and executes a user command
func (e *Engine) ProcessCommand(input string) CommandResult {
	// Parse the command
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return CommandResult{
			Success: false,
			Message: "No command entered",
		}
	}
	
	// Handle git commands
	if parts[0] == "git" && len(parts) > 1 {
		gitCmd := parts[1]
		args := parts[2:]
		
		if cmd, exists := CommandRegistry[gitCmd]; exists {
			result := cmd.Execute(args, e.State)
			
			// Update game state based on result
			e.State.IncreaseAnomaly(result.AnomalyDelta)
			e.State.IncreaseSanity(result.SanityDelta)
			
			// Check for level completion
			if e.CurrentLevel != nil {
				if completed, msg := e.CurrentLevel.ValidateFunc(e.State); completed {
					result.Success = true
					result.Message += "\n\n" + msg
					result.SCPEffect = "🎉 LEVEL COMPLETE! " + msg
					e.State.Score += e.CurrentLevel.ScoreReward
					e.State.CompletedLevels = append(e.State.CompletedLevels, e.LevelNum)
				}
			}
			
			return result
		}
		
		return CommandResult{
			Success:      false,
			Message:      fmt.Sprintf("git: '%s' is not a git command", gitCmd),
			SCPEffect:    "🔴 ERROR: Unknown containment protocol",
			AnomalyDelta: 1,
		}
	}
	
	// Handle non-git commands
	switch parts[0] {
	case "help":
		return CommandResult{
			Success: true,
			Message: "", // UI will display help
		}
	case "status":
		return CommandResult{
			Success: true,
			Message: "", // UI will display status
		}
	default:
		return CommandResult{
			Success:      false,
			Message:      fmt.Sprintf("Command not found: %s", parts[0]),
			SCPEffect:    "⚠️  Invalid Foundation protocol",
			AnomalyDelta: 1,
		}
	}
}


// IsLevelComplete checks if the current level is complete
func (e *Engine) IsLevelComplete() bool {
	if e.CurrentLevel == nil {
		return false
	}
	
	completed, _ := e.CurrentLevel.ValidateFunc(e.State)
	return completed
}

// GetNextLevel returns the next level number if available
func (e *Engine) GetNextLevel() int {
	if e.CurrentLevel != nil && len(e.CurrentLevel.UnlocksNext) > 0 {
		return e.CurrentLevel.UnlocksNext[0]
	}
	return 0
}

// Helper function to generate file hash
func hashContent(content string) string {
	h := sha1.New()
	h.Write([]byte(content))
	return fmt.Sprintf("%x", h.Sum(nil))[:8]
}