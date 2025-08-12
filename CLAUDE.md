# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

gh-q is a GitHub CLI extension that provides ghq-like functionality for managing GitHub repositories locally. It's written in Go and integrates natively with the GitHub CLI (gh).

## Development Commands

### Build and Install
```bash
# Build the project
make build
# or
go build .

# Install as GitHub CLI extension
make install
# or
gh extension install .

# Uninstall the extension
make uninstall
# or
gh extension remove q
```

### Running the Extension
```bash
# After installation, use via gh
gh q get OWNER/REPO
gh q list
gh q create OWNER/REPO
gh q remove OWNER/REPO
```

### Development Workflow
```bash
# Build and test locally
go build .
./gh-q [command]

# Install for testing with gh CLI
make install
gh q [command]
```

### Linting
```bash
# Run Go linter (used in CI)
golangci-lint run
```

## Architecture Overview

### Command Structure
The project uses Cobra for CLI command management. Each command is defined in `cmd/`:
- `cmd/root.go`: Root command setup and global flags
- `cmd/get.go`: Clone repositories from GitHub
- `cmd/list.go`: List managed repositories with filtering
- `cmd/create.go`: Create new local repositories
- `cmd/remove.go`: Remove local repositories

### Core Components

1. **CLI Client** (`internal/cli/client.go`):
   - Central client that handles GitHub API interactions
   - Manages authentication and configuration
   - Coordinates between commands and Git operations

2. **Git Operations** (`internal/git/`):
   - `client.go`: Git operations wrapper using go-git
   - `auth.go`: SSH and HTTPS authentication handling

3. **Directory Management** (`internal/cli/directory.go`):
   - Determines repository root directory (GHQ_ROOT → git config ghq.root → ~/ghq)
   - Handles repository path resolution

### Key Implementation Details

- **Authentication**: Supports both HTTPS (via gh auth token) and SSH (via GHQ_SSH_KEY_PATH)
- **Repository Paths**: Follows ghq convention: `root/host/owner/repo`
- **Error Handling**: Uses fmt.Errorf with wrapped errors for context
- **Interactive Selection**: Uses go-fzf for fuzzy finding in list/remove commands

### Configuration

The tool respects the following configuration sources:
1. Environment variables:
   - `GHQ_ROOT`: Repository root directory
   - `GHQ_SSH_KEY_PATH`: Path to SSH key for Git operations
   - `GHQ_USER`: Default username for repository operations
2. Git config:
   - `ghq.root`: Repository root directory
   - `ghq.user`: Default username
3. GitHub CLI authentication token

### Testing

Currently, the project does not have test files. When adding tests:
- Create `*_test.go` files alongside implementation files
- Use standard Go testing package
- Run tests with `go test ./...`

### Release Process

Releases are fully automated using Release Please:
1. When changes are pushed to `main`, Release Please creates/updates a PR with version bumps and CHANGELOG updates
2. Merging the Release Please PR triggers the actual release
3. GitHub Actions automatically builds and publishes precompiled binaries via `gh-extension-precompile`

The release workflow is defined in `.github/workflows/release-please.yml`

### Code Style Guidelines

- Follow standard Go conventions
- Use go-git for Git operations (already a dependency)
- Use cobra for command structure (existing pattern)
- Maintain compatibility with GitHub CLI extension interface
- Keep commands focused and Unix-philosophy compliant