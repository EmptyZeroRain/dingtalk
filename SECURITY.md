# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 1.5.x   | :white_check_mark: |
| < 1.5   | :x:                |

## Security Considerations

### Credentials Protection

- **Never commit** `access_token` or `secret` to version control
- Use environment variables or config files with proper file permissions (`chmod 600`)
- Rotate credentials immediately if exposed

### Config File Permissions

```bash
chmod 600 ~/.dingtalk/config.toml
```

### Domain Validation

When using private DingTalk deployment, ensure:
- Use HTTPS only
- Verify domain ownership
- Avoid untrusted or public domains

### Input Sanitization

The tool passes user input directly to DingTalk API. Ensure:
- Validate message content before sending
- Sanitize mobile numbers in `--atMobiles`
- Be cautious with user-provided URLs in Link/ActionCard messages

## Reporting a Vulnerability

If you discover a security vulnerability, please report it responsibly:

1. **Do NOT** open a public GitHub issue
2. Email to: [your-email@example.com]
3. Include:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact

### Response Timeline

| Stage | Timeline |
|-------|----------|
| Acknowledgment | 48 hours |
| Initial assessment | 1 week |
| Fix or mitigation | 2 weeks |

### What to Expect

- **Accepted**: Fix will be released in next patch version, credit given (unless anonymity requested)
- **Declined**: Explanation provided, may suggest workarounds

## Security Best Practices

1. **Least Privilege**: Grant only necessary permissions to DingTalk robot
2. **Network Security**: Run in trusted network environments
3. **Logging**: Enable debug mode (`-D`) only for troubleshooting, disable in production
4. **Dependencies**: Regularly update dependencies with `go mod tidy`
