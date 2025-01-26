package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("No command provided")
	}
	switch args[1] {
	case "branch-name":
		branchName()
	case "commit-msg":
		log.Fatal("Not implemented")
	}
}

// Check if the branch name matches the required pattern
func branchName() {
	// get current branch name
	output, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
	if err != nil {
		log.Fatalf("Could not get current branch name: %v", err)
	}

	bName := strings.TrimSpace(string(output))

	bNamePattern := getBranchNamePattern()

	regex := regexp.MustCompile(bNamePattern)
	if !regex.MatchString(bName) {
		log.Printf("Branch name: %s\n", output)
		log.Fatal("Branch name does not match the required pattern")
	}
}

// Get the branch name pattern from the command line arguments or .gitgud file
func getBranchNamePattern() string {
	var bNamePattern string
	if len(os.Args) >= 3 {
		bNamePattern = os.Args[2]
	} else {
		// read from .gitgud file
		file, err := os.Open(".gitgud")
		if err != nil {
			log.Fatalf("Could not open .gitgud file: %v", err)
		}
		defer file.Close()

		// read file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			keyValue := strings.Split(line, ": ")
			if keyValue[0] != "branchNamePattern" {
				continue
			}
			bNamePattern = keyValue[1]
		}

		if bNamePattern == "" {
			log.Fatalf("No branch name pattern provided in .gitgud file")
		}
	}
	return bNamePattern
}
