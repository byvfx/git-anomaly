# Git Command Improvements

## 1. Multiple File Support for `git add`

The `git add` command now supports adding multiple files in a single command:

```bash
# Add multiple specific files
git add file1.txt file2.txt file3.txt

# Mix existing and non-existing files
git add safe_file.txt nonexistent.txt research.txt
# Output: Added 2 files to staging area
#         Warning: pathspec 'nonexistent.txt' did not match any files

# Still supports wildcards
git add .
git add *
```

### Features:
- Adds all valid files specified
- Shows warnings for files that don't exist
- Tracks anomaly files separately with appropriate warnings
- Updates sanity/anomaly levels based on file types

## 2. Modern `git switch` Command

Added support for the modern `git switch` command as a replacement for `git checkout`:

```bash
# Switch to existing branch
git switch main

# Create and switch to new branch
git switch -c experiment-1

# Error handling
git switch nonexistent-branch
# Output: fatal: invalid reference: nonexistent-branch
```

### Comparison with checkout:
- `git checkout <branch>` - Still works for switching branches
- `git checkout -b <branch>` - Create and switch (old style)
- `git switch <branch>` - Switch to existing branch (new style)
- `git switch -c <branch>` - Create and switch (new style)

Both commands are available to match real Git behavior, allowing players to learn either the traditional or modern Git workflows.

## 3. Briefing Command

Added convenience commands to re-display current level information:
- `brief` or `briefing` - Shows current level details
- `objective` or `objectives` - Same as brief

This helps players who forget their current objective or need to see the files they're working with.