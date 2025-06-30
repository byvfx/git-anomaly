package game

// Level represents a game level with SCP theming
type Level struct {
	ID          int
	Title       string
	SCPNumber   string // e.g., "SCP-████"
	Description string
	Objective   string

	// SCP Documentation
	ObjectClass      string // Safe, Euclid, Keter
	ContainmentProcs string
	IncidentReport   string

	// Game mechanics
	InitialFiles     map[string]string
	RequiredCommands []string
	ValidateFunc     func(*GameState) (bool, string)

	// Rewards
	ScoreReward int
	UnlocksNext []int
}

// GetLevel returns the level definition for the given level number
func GetLevel(levelNum int) *Level {
	switch levelNum {
	case 1:
		return &Level1
	case 2:
		return &Level2
	case 3:
		return &Level3
	case 4:
		return &Level4
	default:
		return nil
	}
}

// Level1 - Initial Containment Setup
var Level1 = Level{
	ID:          1,
	Title:       "Initial Containment Setup",
	SCPNumber:   "SCP-████",
	ObjectClass: "Safe",
	Description: "A new anomalous codebase has been discovered. Establish initial containment protocols by setting up version control.",
	Objective:   "Configure git, initialize repository, and perform initial commit of all files",

	InitialFiles: map[string]string{
		"README.txt":      "Project SCP-████: Anomalous Codebase Documentation",
		"anomaly.txt":     "This file writes itself...",
		"containment.log": "Day 1: Initial discovery. File behavior appears autonomous.",
	},

	RequiredCommands: []string{"git config", "git init", "git add .", "git commit"},

	ContainmentProcs: `CONTAINMENT PROTOCOL SCP-████-CP1:
1. Configure researcher identity for accountability
2. Initialize secure version control repository  
3. Stage ALL discovered files for tracking
4. Perform initial containment commit

Note: Use 'git add .' to stage all files at once`,

	IncidentReport: `INCIDENT LOG ████-1
Digital anomaly discovered on Foundation server Gamma-7.
Multiple files showing autonomous behavior patterns.
Dr. ████████ assigned as lead researcher.
IMMEDIATE ACTION: Establish version control for all files.`,

	ValidateFunc: func(state *GameState) (bool, string) {
		// Check git config was used
		if state.ConfigName == "" || state.ConfigEmail == "" {
			return false, "Researcher identity not configured (use git config)"
		}
		if !state.IsInitialized {
			return false, "Repository not initialized"
		}
		if len(state.Commits) == 0 {
			return false, "No commits found - initial containment incomplete"
		}
		// Check that all files were committed
		commit := state.Commits[len(state.Commits)-1]
		if len(commit.Files) < 3 {
			return false, "Not all files contained - use 'git add .' to stage all files"
		}
		return true, "✅ Initial containment established. All files secured."
	},

	ScoreReward: 100,
	UnlocksNext: []int{2},
}

// Level2 - Monitoring Changes
var Level2 = Level{
	ID:          2,
	Title:       "Monitoring Changes",
	SCPNumber:   "SCP-████-A",
	ObjectClass: "Euclid",
	Description: "The entity has begun modifying files. Track and document all changes carefully.",
	Objective:   "Learn to monitor file changes using git status and git diff, then commit the modifications",

	InitialFiles: map[string]string{
		"README.txt":      "Project SCP-████: Anomalous Codebase Documentation\n\n[REDACTED]",
		"anomaly.txt":     "This file writes itself...\nLine 2: Added by the entity",
		"containment.log": "Day 1: Initial discovery. File behavior appears autonomous.\nDay 2: Files showing signs of self-modification.",
		"research.txt":    "Research notes on SCP-████ behavior patterns.",
	},

	RequiredCommands: []string{"git add", "git status", "git diff", "git commit -a"},

	ContainmentProcs: `CONTAINMENT PROTOCOL SCP-████-CP2:
1. Use 'git status' to identify modified files
2. Use 'git diff' to examine specific changes
3. Stage individual files OR use 'git commit -a' to commit all changes
4. Document all modifications in commit messages

IMPORTANT: The entity learns from our actions. Monitor all changes.`,

	IncidentReport: `INCIDENT LOG ████-2
08:30 - Routine file check reveals autonomous modifications
08:45 - Multiple files show timestamp changes without user input
09:00 - Content analysis reveals structured patterns in modifications
ACTION: Document all changes for pattern analysis`,

	ValidateFunc: func(state *GameState) (bool, string) {
		// This level starts with files already committed from Level 1
		// So we need at least 2 commits (initial + modifications)
		if len(state.Commits) < 2 {
			return false, "Modifications not yet committed"
		}
		
		// Check that the player used git diff or status (tracked by command history)
		// For now, we'll just check that they made a commit
		lastCommit := state.Commits[len(state.Commits)-1]
		
		if len(lastCommit.Files) < 3 {
			return false, "Not all modified files were committed"
		}
		
		return true, "✅ All modifications documented. Pattern analysis complete."
	},

	ScoreReward: 150,
	UnlocksNext: []int{3},
}

// Level3 - Historical Analysis  
var Level3 = Level{
	ID:          3,
	Title:       "Historical Analysis",
	SCPNumber:   "SCP-████-B",
	ObjectClass: "Euclid",
	Description: "Investigate the entity's past behavior through commit history analysis.",
	Objective:   "Use git log, git log -p, and git show to understand the anomaly's evolution",

	InitialFiles: map[string]string{
		"anomaly.txt":     "ERROR ERROR ERROR ERROR\nThe pattern is changing...",
		"research.log":    "Day 3: Entity shows learning behavior\nDay 4: Patterns detected in modifications",
		"timeline.txt":    "Tracking anomaly evolution over time",
		"analysis.txt":    "Pattern analysis results pending...",
	},

	RequiredCommands: []string{"git log", "git log -p", "git show"},

	ContainmentProcs: `CONTAINMENT PROTOCOL SCP-████-CP3:
1. Use 'git log' to view commit history
2. Use 'git log -p' to see detailed changes in each commit
3. Use 'git show <commit>' to examine specific commits
4. Document patterns in entity behavior

CRITICAL: Understanding its history may reveal weaknesses.`,

	IncidentReport: `INCIDENT LOG ████-3
10:00 - Historical analysis authorized by O5 Council
10:30 - Previous researchers' notes recovered from commits
11:00 - Pattern emerging in entity's modifications
ACTION: Deep forensic analysis of all commits`,

	ValidateFunc: func(state *GameState) (bool, string) {
		// For historical analysis level, we just need to ensure commits exist
		if len(state.Commits) < 3 {
			return false, "Insufficient historical data for analysis"
		}
		
		// This level is complete once enough history exists
		// In a real implementation, we'd track if git log commands were used
		return true, "✅ Historical analysis complete. Entity patterns documented."
	},

	ScoreReward: 200,
	UnlocksNext: []int{4},
}

// Level4 - Parallel Containment Strategies
var Level4 = Level{
	ID:          4,
	Title:       "Parallel Containment Strategies",
	SCPNumber:   "SCP-████-C",
	ObjectClass: "Keter",
	Description: "The entity has evolved. Test multiple containment approaches simultaneously using branches.",
	Objective:   "Create branches for different strategies, test approaches, and merge successful containment",

	InitialFiles: map[string]string{
		"core.sys":        "CRITICAL: System core - handle with extreme care",
		"anomaly.exe":     "ACTIVE THREAT - DO NOT EXECUTE",
		"strategy_a.txt":  "Containment Strategy A: Isolation Protocol",
		"strategy_b.txt":  "Containment Strategy B: Neutralization Protocol",
		"monitor.log":     "Real-time anomaly behavior tracking",
	},

	RequiredCommands: []string{"git branch", "git switch", "git merge"},

	ContainmentProcs: `CONTAINMENT PROTOCOL SCP-████-CP4:
1. Create separate branches for each containment strategy
2. Use 'git switch -c <branch>' to create and switch
3. Test different approaches in isolation
4. Merge successful strategies back to main
5. Delete failed experiment branches

CRITICAL: The entity adapts. Multiple strategies increase success probability.`,

	IncidentReport: `INCIDENT LOG ████-4
12:00 - Entity shows rapid evolution, previous containment failing
12:30 - O5 authorizes parallel containment experiments
13:00 - Multiple research teams assigned different approaches
ACTION: Implement branching strategy immediately`,

	ValidateFunc: func(state *GameState) (bool, string) {
		// Check for multiple branches
		if len(state.Branches) < 3 {
			return false, "Insufficient parallel experiments (need at least 3 branches)"
		}
		
		// Check for merge (main branch should have commits from other branches)
		mainCommits := state.Branches["main"]
		if len(mainCommits) < 4 {
			return false, "No successful strategies merged to main branch"
		}
		
		return true, "✅ Optimal containment strategy identified and implemented."
	},

	ScoreReward: 250,
	UnlocksNext: []int{5},
}