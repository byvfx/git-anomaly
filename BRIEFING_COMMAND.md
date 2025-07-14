# Briefing Command Added

Added a convenience command to re-display the current level briefing for players who need a reminder of their objectives.

## New Commands:
- `brief` - Re-display current level information
- `briefing` - Same as brief
- `objective` - Same as brief  
- `objectives` - Same as brief

## What it displays:
- Level number and title
- SCP item number and object class
- Full briefing text
- Current objective
- Files detected in the working directory

## Usage:
```bash
# After starting a level
[SCP-████] $ brief

════════ LEVEL 1: Initial Containment ════════

ITEM #: SCP-████
OBJECT CLASS: Safe

BRIEFING:
A simple anomalous file has been discovered. Establish
initial containment.

OBJECTIVE:
Initialize repository and commit the anomalous file

FILES DETECTED:
  • anomaly.txt
```

This command helps players who:
- Forget what files they need to work with
- Need to re-read the objective
- Want to check the level details without restarting