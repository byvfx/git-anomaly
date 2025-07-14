# SCP-████ Containment Walkthrough Guide

**Classification**: Foundation Personnel Only  
**Clearance Required**: Level 2  
**Document Purpose**: Complete containment protocol walkthrough

---

## ⚠️ Important Notes

- This guide follows the **official Git tutorial structure** for optimal learning
- Each level builds on previous knowledge - complete them in order
- Use **Tab completion** and **Up/Down arrows** for efficient command entry
- Type `brief` anytime to review current objectives
- Monitor your Anomaly Level and Researcher Sanity at all times

---

## Getting Started

1. **Start the Game**:
   ```bash
   ./bin/scp-git --classic
   # or
   go run main.go --classic
   ```

2. **Begin Containment**:
   ```
   start
   ```

---

# Level 1: Initial Containment Setup
**Object Class**: Safe  
**Objective**: Configure git, initialize repository, and perform initial commit

## Situation Briefing
A new anomalous codebase has been discovered showing autonomous behavior. You must establish initial containment protocols using version control.

**Files Detected**:
- `README.txt` - Project documentation
- `anomaly.txt` - Self-modifying file
- `containment.log` - Incident logs

## Step-by-Step Walkthrough

### Step 1: Configure Researcher Identity
**Why**: Git requires user identification for accountability and tracking.

```bash
git config user.name "Dr. [Your Name]"
git config user.email "researcher@site-19.scp"
```

**Example**:
```bash
git config user.name "Dr. Smith"
git config user.email "smith@site-19.scp"
```

### Step 2: Initialize Repository
**Why**: Creates the containment environment (.git directory).

```bash
git init
```

**Expected Output**: `Initialized empty Git repository`

### Step 3: Stage All Files
**Why**: Prepares all discovered files for initial containment.

```bash
git add .
```

**Note**: The `.` means "all files in current directory"

### Step 4: Perform Initial Commit
**Why**: Secures all files in version control history.

```bash
git commit -m "Initial containment of SCP-████"
```

## ✅ Success Criteria
- Git configuration completed
- Repository initialized
- All 3 files committed
- Foundation protocols established

**Level Complete!** Anomaly contained, researcher sanity maintained.

---

# Level 2: Monitoring Changes
**Object Class**: Euclid  
**Objective**: Track autonomous file modifications using Git monitoring tools

## Situation Briefing
The entity has begun modifying files autonomously. You must track and document all changes for pattern analysis.

**Files Modified**: 4 files now show evidence of entity interference

## Step-by-Step Walkthrough

### Step 1: Check File Status
**Why**: Identifies which files the entity has modified.

```bash
git status
```

**Expected Output**: Shows modified files in red (unstaged changes)

### Step 2: Examine Specific Changes
**Why**: Understand exactly what the entity changed.

```bash
git diff
```

**Advanced**: Check specific files:
```bash
git diff README.txt
git diff anomaly.txt
```

**What to Look For**:
- Lines starting with `-` (removed content)
- Lines starting with `+` (added content)
- Entity behavior patterns

### Step 3: Stage Individual Files (Method A)
**Why**: Selective staging for precise control.

```bash
git add README.txt
git add anomaly.txt
git add containment.log
git add research.txt
```

### Step 3 Alternative: Commit All Changes (Method B)
**Why**: Faster when you want to commit all modifications.

```bash
git commit -a -m "Document autonomous modifications - Day 2"
```

**Note**: The `-a` flag automatically stages and commits all tracked files.

### Step 4: Document Changes
**Why**: Maintains research log for pattern analysis.

```bash
git commit -m "Document entity modifications - pattern analysis required"
```

## 🔍 Pro Tips
- Use `git status` frequently to monitor entity activity
- Use `git diff` before committing to understand changes
- The `-a` flag only works on files already tracked by Git

## ✅ Success Criteria
- File modifications detected and analyzed
- All changes documented in commits
- Pattern analysis data collected

**Level Complete!** Entity behavior documented for research analysis.

---

# Level 3: Historical Analysis
**Object Class**: Euclid  
**Objective**: Investigate entity behavior through commit history forensics

## Situation Briefing
The O5 Council has authorized deep historical analysis. Previous researchers may have left clues in the commit history about the entity's true nature.

## Step-by-Step Walkthrough

### Step 1: View Commit History
**Why**: See the timeline of all containment activities.

```bash
git log
```

**Output Shows**:
- Commit hash (unique identifier)
- Author and date
- Commit messages
- Chronological timeline

**Navigation**:
- Press `Space` to scroll down
- Press `q` to quit log view

### Step 2: Detailed History Analysis
**Why**: See exactly what changed in each commit.

```bash
git log -p
```

**What This Shows**:
- Complete commit information
- Full diff of every change
- Line-by-line modifications
- Entity behavior patterns over time

### Step 3: Examine Specific Incidents
**Why**: Focus on particular commits of interest.

```bash
git show HEAD
```

**Advanced Forensics**:
```bash
# Show the most recent commit
git show HEAD

# Show the previous commit
git show HEAD^

# Show a specific commit (use actual hash)
git show a1b2c3d
```

**Partial Commit IDs**: You can use just the first 7 characters of any commit hash.

### Step 4: Pattern Recognition
Look for:
- Recurring modification patterns
- Time-based behavior changes
- Escalation in entity capabilities
- Research notes from previous containment attempts

## 🔬 Advanced Analysis Commands

### Search Commit History
```bash
# Find commits mentioning specific terms
git log --grep="anomaly"
git log --grep="breach"
```

### Show Commit Statistics
```bash
git log --stat
```

## ✅ Success Criteria
- Historical timeline established
- Entity evolution patterns documented
- Previous researcher data recovered
- Forensic analysis complete

**Level Complete!** Critical intelligence gathered on entity behavior.

---

# Level 4: Parallel Containment Strategies
**Object Class**: Keter  
**Objective**: Test multiple containment approaches using branch management

## Situation Briefing
The entity has evolved beyond previous containment protocols. Multiple parallel strategies must be tested simultaneously to prevent total containment failure.

**Critical Files**:
- `core.sys` - System critical, handle with care
- `anomaly.exe` - Active threat, do not execute
- Strategy files for different approaches

## Step-by-Step Walkthrough

### Step 1: Assess Current Situation
**Why**: Understand the main branch state before branching.

```bash
git status
git branch
```

**Expected**: Should show you're on `main` branch.

### Step 2: Create Isolation Strategy Branch
**Why**: Test containment through isolation.

```bash
git switch -c strategy-isolation
```

**Alternative syntax**:
```bash
git branch strategy-isolation
git switch strategy-isolation
```

### Step 3: Implement Isolation Approach
**Why**: Test if isolation contains the entity.

```bash
# Modify files for isolation strategy
git add strategy_a.txt
git commit -m "Implement isolation containment protocol"
```

### Step 4: Create Neutralization Strategy Branch
**Why**: Test active neutralization approach.

```bash
git switch main
git switch -c strategy-neutralization
```

### Step 5: Implement Neutralization Approach
**Why**: Test if neutralization stops entity.

```bash
# Modify files for neutralization strategy
git add strategy_b.txt
git commit -m "Implement neutralization containment protocol"
```

### Step 6: Create Monitoring Strategy Branch
**Why**: Test enhanced monitoring approach.

```bash
git switch main
git switch -c strategy-monitoring
```

### Step 7: Test All Strategies
Work in each branch:
```bash
# Switch between branches to test
git switch strategy-isolation
git switch strategy-neutralization
git switch strategy-monitoring
git switch main
```

### Step 8: Merge Successful Strategy
**Why**: Implement the most effective containment method.

```bash
# Return to main and merge the best strategy
git switch main
git merge strategy-isolation
```

**If successful merge**:
```bash
git commit -m "Optimal containment strategy implemented"
```

### Step 9: Clean Up Failed Experiments
**Why**: Remove unsuccessful approaches.

```bash
# Delete failed experiment branches
git branch -d strategy-neutralization
```

## 🧪 Advanced Branch Management

### View All Branches
```bash
git branch -a
```

### Switch Between Strategies Quickly
```bash
git switch strategy-isolation
git switch strategy-neutralization
git switch main
```

### Compare Strategies
```bash
git diff main strategy-isolation
```

## ⚠️ Critical Warnings
- Always return to `main` before creating new branches
- Test each strategy thoroughly before merging
- Keep at least 3 parallel experiments running
- Document results in commit messages

## ✅ Success Criteria
- At least 3 parallel strategy branches created
- Different containment approaches tested
- Successful strategy merged to main
- Failed experiments properly documented
- Optimal containment protocol established

**Level Complete!** Multi-pronged containment strategy successful. Entity contained.

---

# Advanced Tips & Troubleshooting

## Tab Completion Features
- Type `git ` + **Tab** → See all git commands
- Type `git add ` + **Tab** → See available files
- Type `git switch ` + **Tab** → See available branches
- Type `git config ` + **Tab** → See config options

## Command History
- **Up Arrow** → Previous command
- **Down Arrow** → Next command
- History persists between sessions

## Useful Commands
- `brief` → Review current level objectives
- `status` → Check game status and stats
- `help` → Full command reference
- `clear` → Clear screen
- `breathe` → Recover sanity (+5%)

## Common Issues

### "Repository not initialized"
**Solution**: Run `git init` first

### "Nothing to commit"
**Solution**: Use `git add .` to stage files first

### "Not all files contained"
**Solution**: Make sure to use `git add .` to stage ALL files

### "No commits found"
**Solution**: Use `git commit -m "message"` after staging files

### "Git config not set"
**Solution**: Set both name and email:
```bash
git config user.name "Your Name"
git config user.email "your.email@site.scp"
```

### Anomaly Level Too High
**Solution**: 
- Take breaks with `breathe` command
- Follow protocols exactly
- Avoid random commands

### Branch Issues
**Solution**:
- Use `git branch` to see current branch
- Use `git switch main` to return to main
- Create branches with `git switch -c branch-name`

---

## Research Notes

### Entity Behavior Patterns
1. **Level 1**: Passive observation phase
2. **Level 2**: Active file modification begins
3. **Level 3**: Historical analysis reveals escalation
4. **Level 4**: Advanced defensive capabilities emerge

### Best Practices
- Always configure git before starting
- Use descriptive commit messages
- Check `git status` frequently
- Use `git diff` to understand changes
- Keep branches organized and clean

### Educational Progression
This walkthrough follows the official Git tutorial structure:
1. **Setup** → Configuration and initialization
2. **Changes** → Tracking modifications
3. **History** → Understanding the timeline
4. **Branching** → Parallel development workflows

---

**Classification**: Foundation Internal Use Only  
**Prepare for Contingency Plan Omega if containment fails**

*Secure. Contain. Protect.*