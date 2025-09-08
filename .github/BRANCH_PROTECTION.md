# Branch Protection Configuration

This document outlines the recommended branch protection settings for this repository to ensure code quality and security.

## Main Branch Protection Settings

The following settings should be applied to the `main` branch:

### Required Status Checks
- [x] Require status checks to pass before merging
- [x] Require branches to be up to date before merging
- Required status checks:
  - `test` (CI workflow)
  - `security` (Security scanning)

### Restrictions
- [x] Restrict pushes that create files larger than 100 MB
- [x] Require signed commits (recommended for enhanced security)

### Required Pull Request Reviews
- [x] Require pull request reviews before merging
- Required approving reviews: **2**
- [x] Dismiss stale pull request approvals when new commits are pushed
- [x] Require review from code owners (if CODEOWNERS file exists)
- [x] Restrict reviews to users with push access

### Additional Settings
- [x] Require conversation resolution before merging
- [x] Require linear history (no merge commits)
- [x] Include administrators (apply rules to admins as well)
- [x] Allow force pushes: **Disabled**
- [x] Allow deletions: **Disabled**

## GitHub CLI Commands

You can apply these settings using the GitHub CLI:

```bash
# Enable branch protection for main branch
gh api repos/:owner/:repo/branches/main/protection \
  --method PUT \
  --field required_status_checks='{"strict":true,"contexts":["test","security"]}' \
  --field enforce_admins=true \
  --field required_pull_request_reviews='{"required_approving_review_count":2,"dismiss_stale_reviews":true,"require_code_owner_reviews":true,"require_last_push_approval":true}' \
  --field restrictions='null' \
  --field allow_force_pushes=false \
  --field allow_deletions=false
```

## Repository Settings

### General Security Settings
- [x] Enable private vulnerability reporting
- [x] Enable Dependabot alerts
- [x] Enable Dependabot security updates
- [x] Enable Dependabot version updates

### Code Security and Analysis
- [x] Enable CodeQL analysis
- [x] Enable secret scanning
- [x] Enable push protection for secrets

### Branch Naming Convention
- `main` - Production-ready code
- `develop` - Development integration branch
- `feature/*` - Feature development branches
- `hotfix/*` - Critical bug fixes
- `release/*` - Release preparation branches

## Implementation Steps

1. **Repository Admin**: Navigate to Settings > Branches
2. **Add Rule**: Click "Add rule" and enter `main` as branch name pattern
3. **Configure Settings**: Apply all the settings listed above
4. **Save Changes**: Click "Create" to apply the protection rules

## Rationale

These protection rules ensure:

- **Code Quality**: All code is reviewed before merging
- **Security**: Automated security scanning prevents vulnerabilities
- **Stability**: CI tests must pass before changes are accepted
- **Traceability**: Linear history maintains clear change tracking
- **Compliance**: Signed commits provide additional security verification

## Notes

- Repository administrators can temporarily disable these rules for emergency situations
- Consider adding a CODEOWNERS file to automatically request reviews from relevant team members
- Review and update these settings periodically as the project evolves