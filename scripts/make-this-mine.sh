#!/bin/bash
# This script can quickly replace keywords in a whole directory, making a cloned project yours.

# Directory to search
directory="../"

# Script filename
scriptFilename=$(basename "$0")

# List of key-value pairs for word replacement
replacementList=("sirjager:orgnisation" "gomicro:repositoryname" "GoMicro:RepositoryName")

# Ignore list - directories and files to skip
ignoreDirs=("ignoredir1" "ignoredir2")
ignoreFiles=("$scriptFilename" "ignorefile1.txt" "ignorefile2.txt")

# Function to recursively replace the words in files
replaceWords() {
    local file="$1"

    # Check if the file is in the ignore list
    for ignoreFile in "${ignoreFiles[@]}"; do
        if [[ $file == *"$ignoreFile"* ]]; then
            return
        fi
    done

    # Replace the words in the file
    for replacementPair in "${replacementList[@]}"; do
        oldWord="${replacementPair%%:*}"
        newWord="${replacementPair#*:}"
        sed -i "s/$oldWord/$newWord/g" "$file"
    done
}

# Function to search and replace in a directory
searchAndReplace() {
    local dir="$1"

    # Iterate over files and directories
    for file in "$dir"/*; do
        if [ -d "$file" ]; then
            local skip=0

            # Check if the directory is in the ignore list
            for ignoreDir in "${ignoreDirs[@]}"; do
                if [[ $file == *"$ignoreDir"* ]]; then
                    skip=1
                    break
                fi
            done

            # Skip ignored directories
            if [ $skip -eq 1 ]; then
                continue
            fi

            # Recursively search in subdirectories
            searchAndReplace "$file"
        else
            # Replace the words in files
            replaceWords "$file"
        fi
    done
}

# Call the searchAndReplace function with the specified directory
searchAndReplace "$directory"

echo "Word replacement completed."
