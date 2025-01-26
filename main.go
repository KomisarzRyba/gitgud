package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/komisarzryba/gitgud/config"
)

var logger = log.New(os.Stdout, "[GITGUD] ", 0)

func main() {
	args := os.Args
	if len(args) < 2 {
		logger.Fatalf("No command provided")
	}
	switch args[1] {
	case "branch-name":
		branchName()
	case "commit-msg":
		logger.Fatal("Not implemented")
	}
}

// Check if the branch name matches the required pattern
func branchName() {
	bName := getCurBranchName()
	bNamePattern := getBranchNamePattern()
	regex := regexp.MustCompile(bNamePattern)
	if !regex.MatchString(bName) {
		logger.Printf("Branch name: %s\n", bName)
		logger.Fatal("Branch name does not match the required pattern")
	}
}

// Get the current branch name
func getCurBranchName() string {
	output, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
	if err != nil {
		logger.Fatalf("Could not get current branch name: %v", err)
	}
	return strings.TrimSpace(string(output))
}

// Get the branch name pattern from the command line arguments or .gitgud file
func getBranchNamePattern() string {
	var bNamePattern string
	if len(os.Args) >= 3 {
		bNamePattern = os.Args[2]
	} else {
		repoRoot := getRepoRoot()
		config, err := config.NewConfigFromFile(path.Join(repoRoot, ".gitgud"))
		if err != nil {
			logger.Fatalf("Could not read .gitgud file: %v", err)
		}
		bNamePattern = config.BranchNamePattern
	}
	return bNamePattern
}

func getRepoRoot() string {
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		logger.Fatalf("Could not get repository root: %v", err)
	}
	return strings.TrimSpace(string(output))

}
