package game

import (
	"testing"
)

func TestNewGameState(t *testing.T) {
	state := NewGameState()
	
	if state == nil {
		t.Fatal("NewGameState returned nil")
	}
	
	if state.IsInitialized {
		t.Error("New game state should not be initialized")
	}
	
	if state.AnomalyLevel != 0 {
		t.Errorf("Expected anomaly level 0, got %d", state.AnomalyLevel)
	}
	
	if state.ResearcherSanity != 100 {
		t.Errorf("Expected researcher sanity 100, got %d", state.ResearcherSanity)
	}
	
	if state.ContainmentStatus != "SECURE" {
		t.Errorf("Expected containment status SECURE, got %s", state.ContainmentStatus)
	}
}

func TestUpdateContainmentStatus(t *testing.T) {
	tests := []struct {
		name         string
		anomalyLevel int
		expected     string
	}{
		{"Secure", 0, "SECURE"},
		{"Secure High", 49, "SECURE"},
		{"Breach", 50, "BREACH"},
		{"Breach High", 79, "BREACH"},
		{"Critical", 80, "CRITICAL"},
		{"Critical Max", 100, "CRITICAL"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := NewGameState()
			state.AnomalyLevel = tt.anomalyLevel
			state.UpdateContainmentStatus()
			
			if state.ContainmentStatus != tt.expected {
				t.Errorf("Expected status %s for anomaly level %d, got %s",
					tt.expected, tt.anomalyLevel, state.ContainmentStatus)
			}
		})
	}
}

func TestIncreaseAnomaly(t *testing.T) {
	state := NewGameState()
	
	state.IncreaseAnomaly(10)
	if state.AnomalyLevel != 10 {
		t.Errorf("Expected anomaly level 10, got %d", state.AnomalyLevel)
	}
	
	// Test max cap
	state.IncreaseAnomaly(200)
	if state.AnomalyLevel != 100 {
		t.Errorf("Anomaly level should be capped at 100, got %d", state.AnomalyLevel)
	}
}

func TestIncreaseSanity(t *testing.T) {
	state := NewGameState()
	state.ResearcherSanity = 50
	
	state.IncreaseSanity(20)
	if state.ResearcherSanity != 70 {
		t.Errorf("Expected sanity 70, got %d", state.ResearcherSanity)
	}
	
	// Test max cap
	state.IncreaseSanity(50)
	if state.ResearcherSanity != 100 {
		t.Errorf("Sanity should be capped at 100, got %d", state.ResearcherSanity)
	}
}