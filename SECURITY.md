# Security Policy

## Supported Versions

We release patches for security vulnerabilities. Which versions are eligible for receiving such patches depends on the CVSS v3.0 Rating:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability within this project, please send an email to the project maintainers. All security vulnerabilities will be promptly addressed.

Please include the following information in your report:

- Description of the vulnerability
- Steps to reproduce the issue
- Possible impact of the vulnerability
- Any suggested fixes or mitigations

**Please do not report security vulnerabilities through public GitHub issues.**

## Security Best Practices

When using this application:

1. **Environment Variables**: Never commit `.env` files or hardcode sensitive credentials
2. **AWS Credentials**: Use IAM roles when possible, avoid hardcoding AWS keys
3. **Redis Security**: Secure your Redis instance with authentication and network restrictions
4. **HTTPS**: Always use HTTPS in production environments
5. **Input Validation**: Validate all input data before processing
6. **Rate Limiting**: Implement rate limiting to prevent abuse

## Dependencies

We regularly update dependencies to address security vulnerabilities. Please keep your installation up to date.

## Contact

For any security-related questions or concerns, please contact the project maintainers directly.