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
}

// CommandRegistry maps command names to their implementations
var CommandRegistry = map[string]GitCommand{
	"config":   &ConfigCommand{},
	"init":     &InitCommand{},
	"add":      &AddCommand{},
	"commit":   &CommitCommand{},
	"status":   &StatusCommand{},
	"diff":     &DiffCommand{},
	"log":      &LogCommand{},
	"show":     &ShowCommand{},
	"branch":   &BranchCommand{},
	"checkout": &CheckoutCommand{},
	"switch":   &SwitchCommand{},
	"merge":    &MergeCommand{},
}

// ConfigCommand implements git config
type ConfigCommand struct{}

func (c *ConfigCommand) Execute(args []string, state *GameState) CommandResult {
	if len(args) < 2 {
		return CommandResult{
			Success:   false,
			Message:   "usage: git config <key> <value>\nCommon configurations:\n  git config user.name \"Your Name\"\n  git config user.email \"email@example.com\"",
			SCPEffect: "⚠️  Researcher identity required for accountability",
		}
	}
	
	key := args[0]
	value := strings.Join(args[1:], " ")
	
	switch key {
	case "user.name":
		state.ConfigName = value
		return CommandResult{
			Success:     true,
			Message:     fmt.Sprintf("Configured user.name: %s", value),
			SCPEffect:   fmt.Sprintf("✅ Researcher identity confirmed: Dr. %s", value),
		}
	case "user.email":
		state.ConfigEmail = value
		return CommandResult{
			Success:     true,
			Message:     fmt.Sprintf("Configured user.email: %s", value),
			SCPEffect:   "✅ Foundation contact protocol established",
		}
	default:
		return CommandResult{
			Success:      false,
			Message:      fmt.Sprintf("error: key does not contain a section: %s", key),
			SCPEffect:    "🔴 Unknown configuration parameter",
			AnomalyDelta: 1,
		}
	}
}

func (c *ConfigCommand) Help() string {
	return "Configure Git settings (user.name, user.email)"
}

func (c *ConfigCommand) RequiredArgs() int {
	return 2
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
			Success:     false,
			Message:     "Nothing specified, nothing added",
			SCPEffect:   "⚠️  WARNING: Specify files for containment",
			AnomalyDelta: 1,
		}
	}
	
	// Process all arguments (supports multiple files)
	var addedFiles []string
	var notFoundFiles []string
	var anomalyFilesAdded int
	totalAnomalyDelta := 0
	
	for _, arg := range args {
		// Handle "git add ." or "git add *"
		if arg == "." || arg == "*" {
			for filename, fileState := range state.WorkingDir {
				if !fileState.Staged {
					stagingFile := fileState
					stagingFile.Staged = true
					state.StagingArea[filename] = stagingFile
					addedFiles = append(addedFiles, filename)
					
					if strings.Contains(filename, "anomaly") {
						anomalyFilesAdded++
						totalAnomalyDelta += 2
					}
				}
			}
			continue
		}
		
		// Handle specific file
		if fileState, exists := state.WorkingDir[arg]; exists {
			if !fileState.Staged {
				stagingFile := fileState
				stagingFile.Staged = true
				state.StagingArea[arg] = stagingFile
				addedFiles = append(addedFiles, arg)
				
				if strings.Contains(arg, "anomaly") {
					anomalyFilesAdded++
					totalAnomalyDelta += 2
				}
			}
		} else {
			notFoundFiles = append(notFoundFiles, arg)
		}
	}
	
	// Build response based on what happened
	if len(notFoundFiles) > 0 && len(addedFiles) == 0 {
		// All files not found
		return CommandResult{
			Success:      false,
			Message:      fmt.Sprintf("pathspec '%s' did not match any files", strings.Join(notFoundFiles, "', '")),
			SCPEffect:    "🔴 ERROR: Files not found in containment area",
			AnomalyDelta: len(notFoundFiles),
		}
	}
	
	if len(addedFiles) == 0 {
		return CommandResult{
			Success:   true,
			Message:   "No changes to stage",
			SCPEffect: "✓ All specified files already contained in staging area",
		}
	}
	
	// Build success message
	var message strings.Builder
	if len(addedFiles) == 1 {
		message.WriteString(fmt.Sprintf("Added '%s' to staging area", addedFiles[0]))
	} else {
		message.WriteString(fmt.Sprintf("Added %d files to staging area", len(addedFiles)))
	}
	
	if len(notFoundFiles) > 0 {
		message.WriteString(fmt.Sprintf("\nWarning: pathspec '%s' did not match any files", strings.Join(notFoundFiles, "', '")))
	}
	
	// Build SCP effect message
	var scpEffect string
	if anomalyFilesAdded > 0 {
		scpEffect = fmt.Sprintf("⚠️  WARNING: %d anomalous files staged - exercise caution", anomalyFilesAdded)
	} else {
		scpEffect = fmt.Sprintf("✅ %d files staged for containment", len(addedFiles))
	}
	
	
	return CommandResult{
		Success:      len(notFoundFiles) == 0,
		Message:      message.String(),
		SCPEffect:    scpEffect,
		AnomalyDelta: totalAnomalyDelta,
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
			Success:     false,
			Message:     "nothing to commit, working tree clean",
			SCPEffect:   "⚠️  No files staged for containment",
			AnomalyDelta: 1,
		}
	}
	
	// Handle -a flag (commit all tracked modified files)
	if len(args) > 0 && args[0] == "-a" {
		// Stage all modified files
		for filename, fileState := range state.WorkingDir {
			// Only stage files that have been previously committed
			wasCommitted := false
			for _, commit := range state.Commits {
				if _, exists := commit.Files[filename]; exists {
					wasCommitted = true
					break
				}
			}
			if wasCommitted && fileState.Modified {
				stagingFile := fileState
				stagingFile.Staged = true
				state.StagingArea[filename] = stagingFile
			}
		}
		// Remove -a from args for message parsing
		args = args[1:]
	}
	
	// Parse commit message
	message := "Initial containment"
	if len(args) >= 2 && args[0] == "-m" {
		message = strings.Join(args[1:], " ")
	} else if len(args) >= 3 && args[0] == "-a" && args[1] == "-m" {
		message = strings.Join(args[2:], " ")
	}
	
	// Create commit
	author := "Dr. ████████"
	if state.ConfigName != "" {
		author = state.ConfigName
	}
	
	commitID := generateCommitID()
	commit := Commit{
		ID:        commitID,
		Message:   message,
		Author:    author,
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

// DiffCommand implements git diff
type DiffCommand struct{}

func (c *DiffCommand) Execute(args []string, state *GameState) CommandResult {
	if !state.IsInitialized {
		return CommandResult{
			Success:      false,
			Message:      "fatal: not a git repository",
			SCPEffect:    "🔴 ERROR: No containment protocols initialized",
			AnomalyDelta: 1,
		}
	}
	
	var diff strings.Builder
	
	// Show unstaged changes
	diff.WriteString("diff --git\n")
	hasChanges := false
	
	for filename, workingFile := range state.WorkingDir {
		// Check if file has been modified since last commit
		var lastCommittedHash string
		if len(state.Commits) > 0 {
			lastCommit := state.Commits[len(state.Commits)-1]
			if hash, exists := lastCommit.Files[filename]; exists {
				lastCommittedHash = hash
			}
		}
		
		// If file is new or modified
		if lastCommittedHash == "" {
			diff.WriteString(fmt.Sprintf("new file: %s\n", filename))
			diff.WriteString(fmt.Sprintf("+++ %s\n", filename))
			diff.WriteString("@@ -0,0 +1 @@\n")
			diff.WriteString(fmt.Sprintf("+%s\n", strings.ReplaceAll(workingFile.Content, "\n", "\n+")))
			hasChanges = true
		} else if workingFile.Hash != lastCommittedHash {
			diff.WriteString(fmt.Sprintf("--- a/%s\n", filename))
			diff.WriteString(fmt.Sprintf("+++ b/%s\n", filename))
			diff.WriteString("@@ -1 +1 @@\n")
			diff.WriteString(fmt.Sprintf("-%s\n", "[previous content]"))
			diff.WriteString(fmt.Sprintf("+%s\n", strings.ReplaceAll(workingFile.Content, "\n", "\n+")))
			hasChanges = true
		}
	}
	
	if !hasChanges {
		return CommandResult{
			Success:   true,
			Message:   "No changes detected",
			SCPEffect: "✓ All files stable - no anomalous activity detected",
		}
	}
	
	return CommandResult{
		Success:   true,
		Message:   diff.String(),
		SCPEffect: "📊 Anomaly analysis complete - changes documented",
	}
}

func (c *DiffCommand) Help() string {
	return "Show changes between working directory and last commit"
}

func (c *DiffCommand) RequiredArgs() int {
	return 0
}

// LogCommand implements git log
type LogCommand struct{}

func (c *LogCommand) Execute(args []string, state *GameState) CommandResult {
	if !state.IsInitialized {
		return CommandResult{
			Success:      false,
			Message:      "fatal: not a git repository",
			SCPEffect:    "🔴 ERROR: No containment protocols initialized",
			AnomalyDelta: 1,
		}
	}
	
	if len(state.Commits) == 0 {
		return CommandResult{
			Success:   true,
			Message:   "No commits yet",
			SCPEffect: "📋 No containment history available",
		}
	}
	
	var log strings.Builder
	
	// Check for -p flag (show patches)
	showPatch := len(args) > 0 && args[0] == "-p"
	
	// Show commits in reverse chronological order
	for i := len(state.Commits) - 1; i >= 0; i-- {
		commit := state.Commits[i]
		log.WriteString(fmt.Sprintf("commit %s\n", commit.ID))
		log.WriteString(fmt.Sprintf("Author: %s\n", commit.Author))
		log.WriteString(fmt.Sprintf("Date:   %s\n", commit.Timestamp.Format("Mon Jan 02 15:04:05 2006")))
		log.WriteString(fmt.Sprintf("\n    %s\n\n", commit.Message))
		
		if showPatch {
			log.WriteString("    Files changed:\n")
			for filename := range commit.Files {
				log.WriteString(fmt.Sprintf("    - %s\n", filename))
			}
			log.WriteString("\n")
		}
	}
	
	return CommandResult{
		Success:   true,
		Message:   log.String(),
		SCPEffect: "📜 Anomaly timeline retrieved from secure archives",
	}
}

func (c *LogCommand) Help() string {
	return "Show commit history"
}

func (c *LogCommand) RequiredArgs() int {
	return 0
}

// ShowCommand implements git show
type ShowCommand struct{}

func (c *ShowCommand) Execute(args []string, state *GameState) CommandResult {
	if !state.IsInitialized {
		return CommandResult{
			Success:      false,
			Message:      "fatal: not a git repository",
			SCPEffect:    "🔴 ERROR: No containment protocols initialized",
			AnomalyDelta: 1,
		}
	}
	
	if len(args) == 0 {
		// Show latest commit by default
		if len(state.Commits) == 0 {
			return CommandResult{
				Success:   false,
				Message:   "No commits yet",
				SCPEffect: "📋 No containment records found",
			}
		}
		
		commit := state.Commits[len(state.Commits)-1]
		return c.showCommit(commit)
	}
	
	// Find commit by ID (partial match)
	commitID := args[0]
	for _, commit := range state.Commits {
		if strings.HasPrefix(commit.ID, commitID) {
			return c.showCommit(commit)
		}
	}
	
	return CommandResult{
		Success:      false,
		Message:      fmt.Sprintf("fatal: bad object %s", commitID),
		SCPEffect:    "🔴 ERROR: Containment record not found",
		AnomalyDelta: 1,
	}
}

func (c *ShowCommand) showCommit(commit Commit) CommandResult {
	var show strings.Builder
	show.WriteString(fmt.Sprintf("commit %s\n", commit.ID))
	show.WriteString(fmt.Sprintf("Author: %s\n", commit.Author))
	show.WriteString(fmt.Sprintf("Date:   %s\n", commit.Timestamp.Format("Mon Jan 02 15:04:05 2006")))
	show.WriteString(fmt.Sprintf("\n    %s\n\n", commit.Message))
	show.WriteString("Files in this commit:\n")
	for filename, hash := range commit.Files {
		show.WriteString(fmt.Sprintf("    %s [%s]\n", filename, hash))
	}
	
	return CommandResult{
		Success:   true,
		Message:   show.String(),
		SCPEffect: "🔍 Detailed anomaly record retrieved",
	}
}

func (c *ShowCommand) Help() string {
	return "Show details of a specific commit"
}

func (c *ShowCommand) RequiredArgs() int {
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
	}
}

func (c *CheckoutCommand) Help() string {
	return "Switch between containment branches"
}

func (c *CheckoutCommand) RequiredArgs() int {
	return 1
}

// SwitchCommand implements git switch
type SwitchCommand struct{}

func (c *SwitchCommand) Execute(args []string, state *GameState) CommandResult {
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
			Message:   "fatal: missing branch name",
			SCPEffect: "⚠️  WARNING: Specify target containment branch",
		}
	}
	
	// Handle -c flag for creating new branch
	if args[0] == "-c" {
		if len(args) < 2 {
			return CommandResult{
				Success:   false,
				Message:   "fatal: missing branch name after '-c'",
				SCPEffect: "⚠️  WARNING: Specify new containment branch name",
			}
		}
		
		branchName := args[1]
		
		// Check if branch already exists
		if _, exists := state.Branches[branchName]; exists {
			return CommandResult{
				Success:      false,
				Message:      fmt.Sprintf("fatal: a branch named '%s' already exists", branchName),
				SCPEffect:    "⚠️  WARNING: Duplicate containment branch rejected",
				AnomalyDelta: 1,
			}
		}
		
		// Create new branch and switch to it
		state.Branches[branchName] = append([]string{}, state.Branches[state.CurrentBranch]...)
		state.CurrentBranch = branchName
		
		return CommandResult{
			Success:     true,
			Message:     fmt.Sprintf("Switched to a new branch '%s'", branchName),
			SCPEffect:   fmt.Sprintf("✅ New containment branch '%s' created and activated", branchName),
		}
	}
	
	// Switch to existing branch
	branchName := args[0]
	
	if _, exists := state.Branches[branchName]; !exists {
		return CommandResult{
			Success:      false,
			Message:      fmt.Sprintf("fatal: invalid reference: %s", branchName),
			SCPEffect:    "🔴 ERROR: Unknown containment branch",
			AnomalyDelta: 2,
		}
	}
	
	state.CurrentBranch = branchName
	
	return CommandResult{
		Success:     true,
		Message:     fmt.Sprintf("Switched to branch '%s'", branchName),
		SCPEffect:   fmt.Sprintf("✅ Containment branch switched to '%s'", branchName),
	}
}

func (c *SwitchCommand) Help() string {
	return "Switch branches or create new branch with -c"
}

func (c *SwitchCommand) RequiredArgs() int {
	return 1
}

// MergeCommand implements git merge
type MergeCommand struct{}

func (c *MergeCommand) Execute(args []string, state *GameState) CommandResult {
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
			Message:   "fatal: No branch name specified",
			SCPEffect: "⚠️  WARNING: Specify source branch for merge",
		}
	}
	
	sourceBranch := args[0]
	
	// Check if source branch exists
	sourceCommits, exists := state.Branches[sourceBranch]
	if !exists {
		return CommandResult{
			Success:      false,
			Message:      fmt.Sprintf("merge: %s - not something we can merge", sourceBranch),
			SCPEffect:    "🔴 ERROR: Unknown containment branch",
			AnomalyDelta: 2,
		}
	}
	
	// Cannot merge into itself
	if sourceBranch == state.CurrentBranch {
		return CommandResult{
			Success:   false,
			Message:   "Already up to date.",
			SCPEffect: "✓ Current branch already contains all changes",
		}
	}
	
	// Simulate merge by adding source branch commits to current branch
	currentCommits := state.Branches[state.CurrentBranch]
	
	// Find commits that are in source but not in current
	mergedCount := 0
	for _, commitID := range sourceCommits {
		found := false
		for _, existingID := range currentCommits {
			if commitID == existingID {
				found = true
				break
			}
		}
		if !found {
			currentCommits = append(currentCommits, commitID)
			mergedCount++
		}
	}
	
	state.Branches[state.CurrentBranch] = currentCommits
	
	if mergedCount == 0 {
		return CommandResult{
			Success:   true,
			Message:   "Already up to date.",
			SCPEffect: "✓ No new containment data to integrate",
		}
	}
	
	// Create merge commit
	author := "Dr. ████████"
	if state.ConfigName != "" {
		author = state.ConfigName
	}
	
	mergeCommit := Commit{
		ID:        generateCommitID(),
		Message:   fmt.Sprintf("Merge branch '%s' into %s", sourceBranch, state.CurrentBranch),
		Author:    author,
		Timestamp: time.Now(),
		Files:     make(map[string]string),
		Branch:    state.CurrentBranch,
	}
	
	// Add all current files to merge commit
	for filename, fileState := range state.WorkingDir {
		mergeCommit.Files[filename] = fileState.Hash
	}
	
	state.Commits = append(state.Commits, mergeCommit)
	state.Branches[state.CurrentBranch] = append(state.Branches[state.CurrentBranch], mergeCommit.ID)
	
	return CommandResult{
		Success:     true,
		Message:     fmt.Sprintf("Merged %d commits from '%s' into %s", mergedCount, sourceBranch, state.CurrentBranch),
		SCPEffect:   fmt.Sprintf("✅ Containment strategies merged. %d protocols integrated.", mergedCount),
	}
}

func (c *MergeCommand) Help() string {
	return "Merge branches together"
}

func (c *MergeCommand) RequiredArgs() int {
	return 1
}

// Helper function to generate commit ID
func generateCommitID() string {
	h := sha1.New()
	h.Write([]byte(time.Now().String()))
	return fmt.Sprintf("%x", h.Sum(nil))
}