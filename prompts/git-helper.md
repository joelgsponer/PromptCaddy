---
id: "git-helper-v1"
title: "Git Workflow Assistant"
description: "Generate commit messages, branching strategies, and Git workflow guidance"
version: "1.0.0"
parameters:
  - name: "git_task"
    description: "Type of Git assistance needed (commit-msg, branching, workflow, or troubleshooting)"
    type: "string"
    required: false
  - name: "project_type"
    description: "Project context (personal, team, open-source, or enterprise)"
    type: "string"
    required: false
  - name: "change_type"
    description: "Type of changes (feature, bugfix, refactor, docs, or hotfix)"
    type: "string"
    required: false
---

# Git Workflow Assistant

## Change Description
{{selection}}

## Git Context
- **Task Type**: {{git_task}}
- **Project Type**: {{project_type}}
- **Change Type**: {{change_type}}

## Git Assistance Request

Please help with Git workflow tasks focusing on:

### 1. Commit Message Generation
- **Conventional Commits**: Standard commit message format
- **Descriptive Messages**: Clear, actionable commit descriptions
- **Change Categorization**: Proper change type classification
- **Scope Definition**: Affected component or module identification
- **Breaking Changes**: Breaking change documentation

### 2. Branching Strategy
- **Branch Naming**: Consistent branch naming conventions
- **Workflow Patterns**: Git Flow, GitHub Flow, GitLab Flow
- **Branch Hierarchy**: Feature, develop, main branch relationships
- **Merge Strategies**: Merge commit, squash, rebase decisions
- **Release Branching**: Release and hotfix branch management

### 3. Workflow Optimization
- **Pre-commit Hooks**: Code quality automation
- **CI/CD Integration**: Automated testing and deployment
- **Code Review Process**: Pull request and merge request workflows
- **Conflict Resolution**: Merge conflict handling strategies
- **History Management**: Clean Git history maintenance

### 4. Repository Management
- **Issue Linking**: Connecting commits to issues
- **Release Management**: Tagging and release workflows
- **Documentation**: README and contributing guidelines
- **Security**: Sensitive data and credential management
- **Collaboration**: Team coordination and communication

## Output Format

### Commit Message Recommendations

#### Primary Commit Message
```
feat(auth): implement OAuth2 authentication flow

- Add OAuth2 provider integration with Google and GitHub
- Implement secure token storage and refresh mechanism
- Add user profile synchronization from OAuth providers
- Include comprehensive error handling for auth failures

Resolves: #123
Breaking Change: Updates authentication API endpoints
```

**Message Structure Explanation**:
- **Type**: `feat` (feature addition)
- **Scope**: `auth` (authentication module)
- **Description**: Clear, imperative mood description
- **Body**: Detailed list of changes made
- **Footer**: Issue references and breaking change notes

#### Alternative Message Formats

**Conventional Commits Standard**:
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

**Common Types**:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Branching Strategy Recommendations

#### Branch Naming Convention
```bash
# Feature branches
feature/AUTH-123-oauth-integration
feature/user-profile-management

# Bug fix branches
fix/AUTH-456-token-refresh-error
hotfix/critical-security-patch

# Release branches
release/v2.1.0
release/2024-01-sprint

# Documentation branches
docs/api-documentation-update
docs/contributing-guidelines
```

#### Git Flow Implementation
```bash
# Initialize Git Flow
git flow init

# Start a new feature
git flow feature start AUTH-123-oauth-integration

# Finish a feature (merges to develop)
git flow feature finish AUTH-123-oauth-integration

# Start a release
git flow release start v2.1.0

# Finish a release (merges to main and develop)
git flow release finish v2.1.0

# Create a hotfix
git flow hotfix start critical-security-patch
git flow hotfix finish critical-security-patch
```

#### GitHub Flow (Simplified)
```bash
# Create feature branch from main
git checkout main
git pull origin main
git checkout -b feature/oauth-integration

# Make changes and commit
git add .
git commit -m "feat(auth): implement OAuth2 authentication flow"

# Push branch and create PR
git push origin feature/oauth-integration
# Create Pull Request via GitHub UI

# After review, merge to main
# Delete feature branch
git branch -d feature/oauth-integration
```

### Workflow Commands & Scripts

#### Pre-commit Setup
```bash
# Install pre-commit hooks
npm install --save-dev husky lint-staged

# Package.json configuration
{
  "husky": {
    "hooks": {
      "pre-commit": "lint-staged",
      "commit-msg": "commitlint -E HUSKY_GIT_PARAMS"
    }
  },
  "lint-staged": {
    "*.{js,ts}": ["eslint --fix", "prettier --write"],
    "*.{json,md}": ["prettier --write"]
  }
}
```

#### Useful Git Aliases
```bash
# Add to ~/.gitconfig
[alias]
    # Pretty log display
    lg = log --oneline --graph --decorate --all
    
    # Show branch status
    st = status -sb
    
    # Quick commit with message
    cm = commit -m
    
    # Amend last commit
    amend = commit --amend --no-edit
    
    # Push current branch
    pushup = push -u origin HEAD
    
    # Interactive rebase
    rb = rebase -i
    
    # Show diff of staged changes
    staged = diff --cached
    
    # Undo last commit (keep changes)
    undo = reset --soft HEAD~1
```

### Repository Setup Guidance

#### .gitignore Template
```gitignore
# Dependencies
node_modules/
vendor/

# Build outputs
dist/
build/
*.exe
*.dll

# Environment files
.env
.env.local
.env.*.local

# IDE files
.vscode/
.idea/
*.swp
*.swo

# OS files
.DS_Store
Thumbs.db

# Logs
*.log
logs/

# Database
*.db
*.sqlite

# Temporary files
tmp/
temp/
```

#### Repository Structure
```
project-name/
├── .github/
│   ├── workflows/          # GitHub Actions
│   ├── ISSUE_TEMPLATE/     # Issue templates
│   └── PULL_REQUEST_TEMPLATE.md
├── docs/                   # Documentation
├── src/                    # Source code
├── tests/                  # Test files
├── .gitignore
├── README.md
├── CONTRIBUTING.md
└── LICENSE
```

### Common Git Scenarios & Solutions

#### Merge Conflict Resolution
```bash
# When merge conflicts occur
git status                  # See conflicted files
# Edit files to resolve conflicts
git add .                   # Stage resolved files
git commit                  # Complete the merge

# Using merge tool
git mergetool              # Open configured merge tool
```

#### History Cleanup
```bash
# Interactive rebase to clean up commits
git rebase -i HEAD~3       # Edit last 3 commits

# Squash commits
# In the interactive editor:
# pick commit1
# squash commit2
# squash commit3

# Reword commit messages
# pick -> reword for commits to change
```

#### Undoing Changes
```bash
# Undo last commit (keep changes staged)
git reset --soft HEAD~1

# Undo last commit (unstage changes)
git reset HEAD~1

# Undo last commit (discard changes)
git reset --hard HEAD~1

# Revert a specific commit (safe for shared repos)
git revert <commit-hash>
```

### Advanced Git Techniques

#### Git Hooks Examples
```bash
#!/bin/sh
# pre-commit hook example
# Check for TODO comments in staged files

if git diff --cached --name-only | xargs grep -l "TODO\|FIXME" > /dev/null; then
    echo "❌ Commit contains TODO or FIXME comments"
    echo "Please resolve these before committing:"
    git diff --cached --name-only | xargs grep -n "TODO\|FIXME"
    exit 1
fi

echo "✅ Pre-commit checks passed"
exit 0
```

#### Git Worktree Usage
```bash
# Create worktree for parallel development
git worktree add ../project-feature feature/new-feature
cd ../project-feature

# Work on feature in separate directory
# Both directories share the same repository

# Remove worktree when done
git worktree remove ../project-feature
```

### Team Collaboration Guidelines

#### Code Review Checklist
- [ ] Code follows project style guidelines
- [ ] All tests pass
- [ ] Documentation is updated
- [ ] Breaking changes are documented
- [ ] Security considerations are addressed
- [ ] Performance impact is considered

#### Pull Request Template
```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests pass
- [ ] Manual testing completed

## Screenshots (if applicable)

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] No merge conflicts
```

Please provide comprehensive Git workflow guidance with practical commands, best practices, and team collaboration strategies.