package game

import (
	"crypto/sha1"
	"fmt"
	"strings"
	"time"
)

// GitCommand interface for all git commands
type GitCommand interface {
	Execute(args []string, state *GameState) CommandResult
	Help() string
	RequiredArgs() int
}

// CommandResult represents the result of executing a command
type CommandResult struct {
	Success      bool
	Message      string
	SCPEffect    string // Special SCP-themed message
	AnomalyDelta int    // Change in anomaly level
	SanityDelta  int    // Change in researcher sanity
}

// CommandRegistry maps command names to their implementations
var CommandRegistry = map[string]GitCommand{
	"init":   &InitCommand{},
	"add":    &AddCommand{},
	"commit": &CommitCommand{},
	"status": &StatusCommand{},
	"branch": &BranchCommand{},
	"checkout": &CheckoutCommand{},
}

// InitCommand implements git init
type InitCommand struct{}

func (c *InitCommand) Execute(args []string, state *GameState) CommandResult {
	if state.IsInitialized {
		return CommandResult{
			Success:      false,
			Message:      "Repository already initialized",
			SCPEffect:    "⚠️  CONTAINMENT BREACH: Attempting to re-initialize secure repository",
			AnomalyDelta: 5,
		}
	}
	
	state.IsInitialized = true
	state.CurrentBranch = "main"
	state.Branches["main"] = []string{}
	
	return CommandResult{
		Success:     true,
		Message:     "Initialized empty Git repository",
		SCPEffect:   "✅ CONTAINMENT ESTABLISHED: Digital anomaly repository secured",
		SanityDelta: 2,
	}
}

func (c *InitCommand) Help() string {
	return "Initialize a new repository for anomaly containment"
}

func (c *InitCommand) RequiredArgs() int {
	return 0
}

// AddCommand implements git add
type AddCommand struct{}

func (c *AddCommand) Execute(args []string, state *GameState) CommandResult {
	if !state.IsInitialized {
		return CommandResult{
			Success:      false,
			Message:      "Repository not initialized",
			SCPEffect:    "🔴 ERROR: Cannot stage files without containment protocols",
			AnomalyDelta: 3,
		}
	}
	
	if len(args) == 0 {
		return CommandResult{
			Success:   false,
			Message:   "Nothing specified, nothing added",
			SCPEffect: "⚠️  WARNING: Specify files for containment",
		}
	}
	
	// Handle "git add ." or "git add *"
	if args[0] == "." || args[0] == "*" {
		count := 0
		for filename, fileState := range state.WorkingDir {
			if !fileState.Staged {
				stagingFile := fileState
				stagingFile.Staged = true
				state.StagingArea[filename] = stagingFile
				count++
			}
		}
		
		if count == 0 {
			return CommandResult{
				Success:   true,
				Message:   "No changes to stage",
				SCPEffect: "✓ All files already contained in staging area",
			}
		}
		
		return CommandResult{
			Success:     true,
			Message:     fmt.Sprintf("Added %d files to staging area", count),
			SCPEffect:   fmt.Sprintf("✅ %d anomalous files staged for containment", count),
			SanityDelta: 1,
		}
	}
	
	// Handle specific file
	filename := args[0]
	if fileState, exists := state.WorkingDir[filename]; exists {
		stagingFile := fileState
		stagingFile.Staged = true
		state.StagingArea[filename] = stagingFile
		
		// Check for anomaly files
		if strings.Contains(filename, "anomaly") {
			return CommandResult{
				Success:      true,
				Message:      fmt.Sprintf("Added '%s' to staging area", filename),
				SCPEffect:    "⚠️  WARNING: Anomalous file staged - exercise caution",
				AnomalyDelta: 2,
			}
		}
		
		return CommandResult{
			Success:     true,
			Message:     fmt.Sprintf("Added '%s' to staging area", filename),
			SCPEffect:   fmt.Sprintf("✅ File '%s' staged for containment", filename),
			SanityDelta: 1,
		}
	}
	
	return CommandResult{
		Success:      false,
		Message:      fmt.Sprintf("pathspec '%s' did not match any files", filename),
		SCPEffect:    "🔴 ERROR: File not found in containment area",
		AnomalyDelta: 1,
	}
}

func (c *AddCommand) Help() string {
	return "Stage files for containment"
}

func (c *AddCommand) RequiredArgs() int {
	return 1
}

// CommitCommand implements git commit
type CommitCommand struct{}

func (c *CommitCommand) Execute(args []string, state *GameState) CommandResult {
	if !state.IsInitialized {
		return CommandResult{
			Success:      false,
			Message:      "Repository not initialized",
			SCPEffect:    "🔴 ERROR: Cannot commit without containment protocols",
			AnomalyDelta: 3,
		}
	}
	
	if len(state.StagingArea) == 0 {
		return CommandResult{
			Success:   false,
			Message:   "nothing to commit, working tree clean",
			SCPEffect: "⚠️  No files staged for containment",
		}
	}
	
	// Parse commit message
	message := "Initial containment"
	if len(args) >= 2 && args[0] == "-m" {
		message = strings.Join(args[1:], " ")
	}
	
	// Create commit
	commitID := generateCommitID()
	commit := Commit{
		ID:        commitID,
		Message:   message,
		Author:    "Dr. ████████",
		Timestamp: time.Now(),
		Files:     make(map[string]string),
		Branch:    state.CurrentBranch,
	}
	
	// Add files to commit
	for filename, fileState := range state.StagingArea {
		commit.Files[filename] = fileState.Hash
	}
	
	// Update game state
	state.Commits = append(state.Commits, commit)
	state.Branches[state.CurrentBranch] = append(state.Branches[state.CurrentBranch], commitID)
	
	// Clear staging area
	fileCount := len(state.StagingArea)
	state.StagingArea = make(map[string]FileState)
	
	return CommandResult{
		Success:     true,
		Message:     fmt.Sprintf("[%s %s] %s\n %d files changed", state.CurrentBranch, commitID[:7], message, fileCount),
		SCPEffect:   fmt.Sprintf("✅ CONTAINMENT SUCCESSFUL: %d anomalies secured with ID %s", fileCount, commitID[:7]),
		SanityDelta: 3,
	}
}

func (c *CommitCommand) Help() string {
	return "Commit staged files to secure containment"
}

func (c *CommitCommand) RequiredArgs() int {
	return 0
}

// StatusCommand implements git status
type StatusCommand struct{}

func (c *StatusCommand) Execute(args []string, state *GameState) CommandResult {
	if !state.IsInitialized {
		return CommandResult{
			Success:      false,
			Message:      "fatal: not a git repository",
			SCPEffect:    "🔴 ERROR: No containment protocols initialized",
			AnomalyDelta: 1,
		}
	}
	
	var status strings.Builder
	status.WriteString(fmt.Sprintf("On branch %s\n", state.CurrentBranch))
	
	if len(state.Commits) == 0 {
		status.WriteString("\nNo commits yet\n")
	}
	
	// Check for staged files
	if len(state.StagingArea) > 0 {
		status.WriteString("\nChanges to be committed:\n")
		status.WriteString("  (use \"git restore --staged <file>...\" to unstage)\n")
		for filename := range state.StagingArea {
			status.WriteString(fmt.Sprintf("\tnew file:   %s\n", filename))
		}
	}
	
	// Check for unstaged files
	unstagedCount := 0
	for _, fileState := range state.WorkingDir {
		if !fileState.Staged {
			unstagedCount++
		}
	}
	
	if unstagedCount > 0 {
		status.WriteString("\nUntracked files:\n")
		status.WriteString("  (use \"git add <file>...\" to include in what will be committed)\n")
		for filename, fileState := range state.WorkingDir {
			if !fileState.Staged {
				status.WriteString(fmt.Sprintf("\t%s\n", filename))
			}
		}
	}
	
	if len(state.StagingArea) == 0 && unstagedCount == 0 {
		status.WriteString("\nnothing to commit, working tree clean\n")
	}
	
	return CommandResult{
		Success:   true,
		Message:   status.String(),
		SCPEffect: "📋 Containment status report generated",
	}
}

func (c *StatusCommand) Help() string {
	return "Show containment status of repository"
}

func (c *StatusCommand) RequiredArgs() int {
	return 0
}

// BranchCommand implements git branch
type BranchCommand struct{}

func (c *BranchCommand) Execute(args []string, state *GameState) CommandResult {
	if !state.IsInitialized {
		return CommandResult{
			Success:      false,
			Message:      "fatal: not a git repository",
			SCPEffect:    "🔴 ERROR: No containment protocols initialized",
			AnomalyDelta: 1,
		}
	}
	
	// List branches if no args
	if len(args) == 0 {
		var branches strings.Builder
		for branch := range state.Branches {
			if branch == state.CurrentBranch {
				branches.WriteString(fmt.Sprintf("* %s\n", branch))
			} else {
				branches.WriteString(fmt.Sprintf("  %s\n", branch))
			}
		}
		
		return CommandResult{
			Success:   true,
			Message:   branches.String(),
			SCPEffect: "📋 Available containment branches listed",
		}
	}
	
	// Create new branch
	branchName := args[0]
	if _, exists := state.Branches[branchName]; exists {
		return CommandResult{
			Success:      false,
			Message:      fmt.Sprintf("fatal: A branch named '%s' already exists", branchName),
			SCPEffect:    "⚠️  WARNING: Duplicate containment branch rejected",
			AnomalyDelta: 1,
		}
	}
	
	// Copy current branch commits to new branch
	state.Branches[branchName] = append([]string{}, state.Branches[state.CurrentBranch]...)
	
	return CommandResult{
		Success:     true,
		Message:     fmt.Sprintf("Created branch '%s'", branchName),
		SCPEffect:   fmt.Sprintf("✅ New containment branch '%s' established", branchName),
		SanityDelta: 1,
	}
}

func (c *BranchCommand) Help() string {
	return "Create or list containment branches"
}

func (c *BranchCommand) RequiredArgs() int {
	return 0
}

// CheckoutCommand implements git checkout
type CheckoutCommand struct{}

func (c *CheckoutCommand) Execute(args []string, state *GameState) CommandResult {
	if !state.IsInitialized {
		return CommandResult{
			Success:      false,
			Message:      "fatal: not a git repository",
			SCPEffect:    "🔴 ERROR: No containment protocols initialized",
			AnomalyDelta: 1,
		}
	}
	
	if len(args) == 0 {
		return CommandResult{
			Success:   false,
			Message:   "error: switch branch requires a branch name",
			SCPEffect: "⚠️  WARNING: Specify target containment branch",
		}
	}
	
	branchName := args[0]
	
	// Check if branch exists
	if _, exists := state.Branches[branchName]; !exists {
		// Handle -b flag for creating and switching
		if len(args) >= 2 && args[0] == "-b" {
			branchName = args[1]
			state.Branches[branchName] = append([]string{}, state.Branches[state.CurrentBranch]...)
			state.CurrentBranch = branchName
			
			return CommandResult{
				Success:     true,
				Message:     fmt.Sprintf("Switched to a new branch '%s'", branchName),
				SCPEffect:   fmt.Sprintf("✅ New containment branch '%s' created and activated", branchName),
				SanityDelta: 2,
			}
		}
		
		return CommandResult{
			Success:      false,
			Message:      fmt.Sprintf("error: pathspec '%s' did not match any known branches", branchName),
			SCPEffect:    "🔴 ERROR: Unknown containment branch",
			AnomalyDelta: 2,
		}
	}
	
	state.CurrentBranch = branchName
	
	return CommandResult{
		Success:     true,
		Message:     fmt.Sprintf("Switched to branch '%s'", branchName),
		SCPEffect:   fmt.Sprintf("✅ Containment branch switched to '%s'", branchName),
		SanityDelta: 1,
	}
}

func (c *CheckoutCommand) Help() string {
	return "Switch between containment branches"
}

func (c *CheckoutCommand) RequiredArgs() int {
	return 1
}

// Helper function to generate commit ID
func generateCommitID() string {
	h := sha1.New()
	h.Write([]byte(time.Now().String()))
	return fmt.Sprintf("%x", h.Sum(nil))
}