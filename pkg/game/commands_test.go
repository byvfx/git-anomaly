package game

import (
	"strings"
	"testing"
)

func TestInitCommand(t *testing.T) {
	state := NewGameState()
	cmd := &InitCommand{}

	// Test successful init
	result := cmd.Execute([]string{}, state)
	if !result.Success {
		t.Error("Init command should succeed on uninitialized repo")
	}
	if !state.IsInitialized {
		t.Error("State should be initialized after init command")
	}
	if state.CurrentBranch != "main" {
		t.Errorf("Expected branch 'main', got '%s'", state.CurrentBranch)
	}
	// Successful init command completed

	// Test init on already initialized repo
	result = cmd.Execute([]string{}, state)
	if result.Success {
		t.Error("Init command should fail on initialized repo")
	}
	if result.AnomalyDelta <= 0 {
		t.Error("Failed init should increase anomaly level")
	}
}

func TestAddCommand(t *testing.T) {
	state := NewGameState()
	state.IsInitialized = true
	state.WorkingDir["test.txt"] = FileState{
		Content: "test content",
		Hash:    "abc123",
	}

	cmd := &AddCommand{}

	// Test add without init
	uninitState := NewGameState()
	result := cmd.Execute([]string{"test.txt"}, uninitState)
	if result.Success {
		t.Error("Add should fail without initialized repo")
	}

	// Test successful add
	result = cmd.Execute([]string{"test.txt"}, state)
	if !result.Success {
		t.Error("Add command should succeed for existing file")
	}
	if _, exists := state.StagingArea["test.txt"]; !exists {
		t.Error("File should be in staging area after add")
	}

	// Test add non-existent file
	result = cmd.Execute([]string{"nonexistent.txt"}, state)
	if result.Success {
		t.Error("Add should fail for non-existent file")
	}
}

func TestCommitCommand(t *testing.T) {
	state := NewGameState()
	state.IsInitialized = true
	state.CurrentBranch = "main"
	state.Branches["main"] = []string{}

	// Stage a file
	state.StagingArea["test.txt"] = FileState{
		Content: "test",
		Hash:    "abc",
		Staged:  true,
	}

	cmd := &CommitCommand{}

	// Test commit with message
	result := cmd.Execute([]string{"-m", "Initial commit"}, state)
	if !result.Success {
		t.Error("Commit should succeed with staged files")
	}
	if len(state.Commits) != 1 {
		t.Errorf("Expected 1 commit, got %d", len(state.Commits))
	}
	if len(state.StagingArea) != 0 {
		t.Error("Staging area should be empty after commit")
	}

	// Test commit without staged files
	result = cmd.Execute([]string{"-m", "Empty commit"}, state)
	if result.Success {
		t.Error("Commit should fail with empty staging area")
	}
}

func TestStatusCommand(t *testing.T) {
	state := NewGameState()
	state.IsInitialized = true
	state.CurrentBranch = "main"

	cmd := &StatusCommand{}

	result := cmd.Execute([]string{}, state)
	if !result.Success {
		t.Error("Status command should always succeed on initialized repo")
	}
	if !strings.Contains(result.Message, "On branch main") {
		t.Error("Status should show current branch")
	}
}

func TestBranchCommand(t *testing.T) {
	state := NewGameState()
	state.IsInitialized = true
	state.CurrentBranch = "main"
	state.Branches["main"] = []string{}

	cmd := &BranchCommand{}

	// Test creating new branch
	result := cmd.Execute([]string{"feature"}, state)
	if !result.Success {
		t.Error("Branch creation should succeed")
	}
	if _, exists := state.Branches["feature"]; !exists {
		t.Error("New branch should exist in branches map")
	}

	// Test creating duplicate branch
	result = cmd.Execute([]string{"feature"}, state)
	if result.Success {
		t.Error("Creating duplicate branch should fail")
	}
}

func TestCheckoutCommand(t *testing.T) {
	state := NewGameState()
	state.IsInitialized = true
	state.CurrentBranch = "main"
	state.Branches["main"] = []string{}
	state.Branches["feature"] = []string{}

	cmd := &CheckoutCommand{}

	// Test checkout existing branch
	result := cmd.Execute([]string{"feature"}, state)
	if !result.Success {
		t.Error("Checkout should succeed for existing branch")
	}
	if state.CurrentBranch != "feature" {
		t.Errorf("Expected current branch 'feature', got '%s'", state.CurrentBranch)
	}

	// Test checkout non-existent branch
	result = cmd.Execute([]string{"nonexistent"}, state)
	if result.Success {
		t.Error("Checkout should fail for non-existent branch")
	}

	// Test checkout -b
	result = cmd.Execute([]string{"-b", "new-branch"}, state)
	if !result.Success {
		t.Error("Checkout -b should create and switch to new branch")
	}
	if state.CurrentBranch != "new-branch" {
		t.Error("Should be on new-branch after checkout -b")
	}
}
