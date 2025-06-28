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
	default:
		return nil
	}
}

// Level1 - Initial Containment
var Level1 = Level{
	ID:          1,
	Title:       "Initial Containment",
	SCPNumber:   "SCP-████",
	ObjectClass: "Safe",
	Description: "A simple anomalous file has been discovered. Establish initial containment.",
	Objective:   "Initialize repository and commit the anomalous file",

	InitialFiles: map[string]string{
		"anomaly.txt": "This file writes itself...",
	},

	RequiredCommands: []string{"git init", "git add", "git commit"},

	ContainmentProcs: `All personnel must initialize proper version control before 
handling anomalous digital materials. Standard git protocols 
apply with enhanced monitoring.`,

	IncidentReport: `INCIDENT LOG ████-1
An anomalous text file has been discovered on Foundation servers.
File appears to modify its own contents when unobserved.
Immediate containment via version control required.`,

	ValidateFunc: func(state *GameState) (bool, string) {
		if !state.IsInitialized {
			return false, "Repository not initialized"
		}
		if len(state.Commits) == 0 {
			return false, "No commits found - anomaly not contained"
		}
		return true, "✅ SCP-████ successfully contained"
	},

	ScoreReward: 100,
	UnlocksNext: []int{2},
}

// Level2 - Staging Area Protocols
var Level2 = Level{
	ID:          2,
	Title:       "Staging Area Protocols",
	SCPNumber:   "SCP-████-A",
	ObjectClass: "Euclid",
	Description: "Multiple anomalous files detected. Learn proper staging procedures.",
	Objective:   "Stage and commit only the safe files, avoiding the dangerous anomaly",

	InitialFiles: map[string]string{
		"safe_file.txt":  "Normal file content",
		"anomaly_1.txt":  "I change when you're not looking",
		"anomaly_2.txt":  "I multiply when staged",
		"research.txt":   "Research notes on the anomaly",
		"classified.txt": "[DATA EXPUNGED]",
	},

	RequiredCommands: []string{"git add", "git status", "git commit"},

	ContainmentProcs: `Personnel must carefully stage files to prevent 
cross-contamination. Only verified safe files should be 
committed to the repository. Use 'git status' frequently 
to monitor containment status.`,

	IncidentReport: `INCIDENT LOG ████-2
Multiple anomalous files detected in research directory.
Some files exhibit replication behavior when staged.
Exercise extreme caution during staging procedures.`,

	ValidateFunc: func(state *GameState) (bool, string) {
		if len(state.Commits) < 2 {
			return false, "Insufficient commits for proper containment"
		}
		
		// Check if any anomaly files were committed
		for _, commit := range state.Commits {
			for filename := range commit.Files {
				if filename == "anomaly_1.txt" || filename == "anomaly_2.txt" {
					return false, "⚠️ CONTAINMENT BREACH: Anomalous files committed!"
				}
			}
		}
		
		// Check if safe files were committed
		lastCommit := state.Commits[len(state.Commits)-1]
		if _, hasResearch := lastCommit.Files["research.txt"]; !hasResearch {
			return false, "Research notes not properly secured"
		}
		
		return true, "✅ SCP-████-A contained with proper staging protocols"
	},

	ScoreReward: 150,
	UnlocksNext: []int{3},
}

// Level3 - Branch Containment Procedures
var Level3 = Level{
	ID:          3,
	Title:       "Branch Containment Procedures",
	SCPNumber:   "SCP-████-B",
	ObjectClass: "Euclid",
	Description: "The anomaly has spread. Use branching to isolate the contamination.",
	Objective:   "Create a containment branch and isolate the anomaly",

	InitialFiles: map[string]string{
		"system.txt":     "Core system files",
		"spreading.txt":  "ERROR: FILE CORRUPTION DETECTED",
		"backup.txt":     "Last known good configuration",
		"monitor.log":    "Anomaly detection log",
	},

	RequiredCommands: []string{"git branch", "git checkout", "git add", "git commit"},

	ContainmentProcs: `When contamination is detected, immediately create 
an isolation branch. Move all anomalous materials to the 
containment branch while keeping main branch clean.`,

	IncidentReport: `INCIDENT LOG ████-3
Anomaly has begun spreading through the file system.
Standard containment insufficient. Branch isolation required.
Main branch must remain uncontaminated at all costs.`,

	ValidateFunc: func(state *GameState) (bool, string) {
		// Check if containment branch exists
		if _, exists := state.Branches["containment"]; !exists {
			return false, "No containment branch created"
		}
		
		// Check if main branch is clean
		mainCommits := state.Branches["main"]
		if len(mainCommits) == 0 {
			return false, "Main branch compromised"
		}
		
		// Verify anomaly is isolated in containment branch
		containmentCommits := state.Branches["containment"]
		if len(containmentCommits) == 0 {
			return false, "Anomaly not properly isolated"
		}
		
		return true, "✅ SCP-████-B successfully isolated in containment branch"
	},

	ScoreReward: 200,
	UnlocksNext: []int{4},
}