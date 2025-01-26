# GitGud Pre-Commit Hooks

GitGud is a lightweight Git pre-commit hook written in Go to enforce branch naming conventions and commit message formats across your repository.

## Features

- **Branch Name Validation:** Ensures branch names follow the specified pattern.
- **Commit Message Validation:** Enforces a defined format for commit messages.

## Setup

Add the following to your `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: https://github.com/komisarzryba/gitgud
    rev: main
    hooks:
      - id: branch-name
      - id: commit-msg
```

After configuring, install the hooks with:

```bash
pre-commit install
pre-commit install -t commit-msg
```

## Configuration

Create a `.gitgud` file in the root of your repository with the following content:

```toml
branch_name_pattern = "^[A-Z][a-z]+/[a-z-]+$"
commit_msg_pattern = "^[A-Z].*"
```

- `branch_name_pattern`: Regex to enforce branch naming (e.g., `Name/branch-name`).
- `commit_msg_pattern`: Regex to enforce commit message format (e.g., capitalized messages).

## Usage

GitGud runs automatically before each commit to ensure:

- Your branch name adheres to the configured pattern.
- Your commit message follows the defined format.

If any violations occur, GitGud will print an error and prevent the commit.
