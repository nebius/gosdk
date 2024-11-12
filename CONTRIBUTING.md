# Contribution Guide

We appreciate your interest in contributing! Here’s how you can get involved.

## 🐞 Reporting Issues

### Security Vulnerabilities

If you discover a security issue, please report it promptly via the GitHub ["Report a Vulnerability"][new-security] tab.
For more details on our security policies, see [SECURITY.md](SECURITY.md).

### Bugs

Found a bug?
Before opening a new issue, please check the [existing issues][issues] to see if it’s already been reported.
If it’s a new bug, you can [create an issue here][new-issue].

### Feature Requests

Got an idea for a new feature? We’d love to hear it!
Please [submit a feature request][new-issue] and provide as much detail as possible.

## 🛠️ Contributing Code

If you’d like to contribute code, follow these steps:

1. **Open an Issue:** Start by [opening an issue][new-issue] to discuss your proposal or bug fix.
2. **Fork the Repository:** Create a fork and work on your changes in a new branch.
3. **Submit a Pull Request:** Once your changes are ready, submit a Pull Request (PR) for review.

## 💻 Development Setup

To set up your development environment, ensure you have the following tools installed:

- Go
- Make
- Docker

## 🧪 Testing

### Writing Tests

All new code must include unit tests to ensure coverage and stability.

### Running Tests

Run the tests locally with:

```bash
make test
```

## 🔍 Code Quality

### Linting

Ensure your code meets project standards by running the linter:

```bash
make lint
```

## 🐳 Using Docker for Development

You can use the Docker development environment included in the project.
To start a shell session inside the container, use:

```bash
make bash
```

## 📋 Makefile Commands

To see a list of available `make` commands, run:

```bash
make help
```


[issues]: https://github.com/nebius/gosdk/issues
[new-issue]: https://github.com/nebius/gosdk/issues/new/choose
[new-security]: https://github.com/nebius/gosdk/security/advisories/new
