---
id: "security-review-v1"
title: "Security Review Assistant"
description: "Identify potential security vulnerabilities and provide recommendations for secure coding practices"
version: "1.0.0"
parameters:
  - name: "language"
    description: "Programming language of the code to review"
    type: "string"
    required: false
  - name: "application_type"
    description: "Type of application (web, mobile, desktop, api, or embedded)"
    type: "string"
    required: false
  - name: "security_focus"
    description: "Specific security concern (authentication, authorization, injection, or general)"
    type: "string"
    required: false
---

# Security Review Analysis

## Code Under Review
```{{language}}
{{selection}}
```

## Context
- **Application Type**: {{application_type}}
- **Language**: {{language}}  
- **Security Focus**: {{security_focus}}

## Security Analysis Request

Please conduct a comprehensive security review focusing on:

### 1. Input Validation & Sanitization
- **SQL Injection**: Check for unsafe database queries
- **XSS Prevention**: Cross-site scripting vulnerabilities
- **Command Injection**: Shell command vulnerabilities
- **Path Traversal**: File system access issues
- **Input Validation**: Proper data validation and sanitization

### 2. Authentication & Authorization
- **Authentication Flaws**: Weak login mechanisms
- **Session Management**: Session security issues
- **Authorization Bypass**: Access control problems
- **Privilege Escalation**: Permission elevation risks
- **Token Security**: JWT and API key handling

### 3. Data Protection
- **Sensitive Data Exposure**: Information leakage
- **Encryption**: Data at rest and in transit
- **Password Storage**: Secure password handling
- **Personal Information**: PII protection
- **Cryptographic Issues**: Weak crypto implementations

### 4. Configuration & Deployment
- **Security Misconfiguration**: Unsafe default settings
- **Error Handling**: Information disclosure in errors
- **Logging Security**: Sensitive data in logs
- **Dependencies**: Vulnerable third-party libraries
- **Environment Variables**: Secure configuration management

### 5. Business Logic Security
- **Race Conditions**: Concurrent access issues
- **Business Rule Bypass**: Logic flaw exploitation
- **Rate Limiting**: DoS protection
- **Data Integrity**: Unauthorized data modification
- **Workflow Security**: Process manipulation

## Output Format

For each security finding:

### Vulnerability Assessment
**Title**: Brief vulnerability description
**Severity**: Critical/High/Medium/Low/Info
**Category**: OWASP category (if applicable)
**CWE ID**: Common Weakness Enumeration ID (if applicable)

### Technical Details
**Description**: Detailed explanation of the vulnerability
**Attack Vector**: How an attacker could exploit this
**Impact**: Potential consequences of exploitation
**Likelihood**: Probability of exploitation

### Evidence
```{{language}}
// Vulnerable code section with line numbers
```

### Remediation
**Recommended Fix**: Specific solution approach
**Code Example**: Secure implementation
```{{language}}
// Corrected secure code
```

**Additional Measures**: Complementary security controls
**Testing**: How to verify the fix

### Prevention Guidelines
- Security best practices for this vulnerability type
- Tools and techniques for ongoing detection
- Code review checklist items

## Security Recommendations Summary

Provide a prioritized list of:
1. **Critical Issues**: Immediate attention required
2. **Security Improvements**: Important enhancements  
3. **Best Practices**: General security guidelines
4. **Monitoring**: Security monitoring recommendations

Please ensure all findings include concrete, actionable remediation steps with code examples.