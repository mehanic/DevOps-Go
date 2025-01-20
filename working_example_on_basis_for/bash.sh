#!/bin/bash
# Script to find directories with a .git folder

# Set the starting point, default to current directory if none provided
START_DIR="${1:-.}"

echo "Searching for .git directories under: $START_DIR"

# Use find to search for directories containing .git
find "$START_DIR" -type d -name ".git" | while read -r git_dir; do
    # Print the parent directory of each .git folder
    echo "Git repository found in: $(dirname "$git_dir")"
done

