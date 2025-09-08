# Contributing to Amazon SES Transaction API

Thank you for your interest in contributing to this project! We welcome contributions from the community.

## How to Contribute

### Reporting Bugs

If you find a bug, please create an issue using the bug report template. Include:

- A clear description of the issue
- Steps to reproduce the problem
- Expected vs actual behavior
- Your environment details (OS, Go version, etc.)

### Suggesting Features

We welcome feature suggestions! Please create an issue using the feature request template and include:

- A clear description of the feature
- The problem it solves
- Possible implementation approaches

### Code Contributions

1. **Fork the repository** and create your branch from `main`
2. **Make your changes** following the coding standards below
3. **Add tests** for any new functionality
4. **Update documentation** if needed
5. **Run the test suite** to ensure everything works
6. **Submit a pull request** using the PR template

## Development Setup

1. Clone your fork:
   ```bash
   git clone https://github.com/your-username/golang-amazon-ses.git
   cd golang-amazon-ses
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up Redis for testing:
   ```bash
   docker run -d -p 6379:6379 redis:alpine
   ```

4. Copy environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your values
   ```

5. Run tests:
   ```bash
   go test ./...
   ```

## Coding Standards

### Go Guidelines

- Follow [Effective Go](https://golang.org/doc/effective_go.html) principles
- Use `gofmt` to format your code
- Run `go vet` to check for common errors
- Write clear, descriptive variable and function names
- Add comments for exported functions and complex logic

### Code Structure

- Keep functions small and focused on a single responsibility
- Use dependency injection where appropriate
- Handle errors explicitly - don't ignore them
- Write tests for new functionality

### Git Commit Messages

- Use clear, descriptive commit messages
- Start with a verb in the imperative mood (e.g., "Add", "Fix", "Update")
- Keep the first line under 50 characters
- Reference issues when relevant (e.g., "Fixes #123")

Example:
```
Add email validation for transaction endpoint

- Validate email format before processing
- Return appropriate error for invalid emails
- Add tests for email validation

Fixes #45
```

## Testing

- Write unit tests for new functionality
- Ensure tests are deterministic and can run in any order
- Use table-driven tests where appropriate
- Mock external dependencies (AWS SES, Redis) in unit tests

## Pull Request Process

1. **Update documentation** if your changes affect the API or configuration
2. **Add tests** for any new functionality
3. **Ensure all tests pass** locally before submitting
4. **Update the README** if you add new features or change behavior
5. **Fill out the PR template** completely
6. **Link related issues** in your PR description

## Code Review

All submissions require code review. We use GitHub's review features for this purpose. Here's what we look for:

- **Correctness**: Does the code do what it's supposed to do?
- **Testing**: Are there adequate tests for the changes?
- **Documentation**: Is the code well-documented?
- **Style**: Does the code follow our coding standards?
- **Performance**: Are there any obvious performance issues?

## Getting Help

If you need help with your contribution:

- Check existing issues and PRs for similar work
- Create a draft PR early to get feedback on your approach
- Ask questions in issue comments

## License

By contributing to this project, you agree that your contributions will be licensed under the MIT License.