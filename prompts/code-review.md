---
id: "code-review-v1"
title: "Code Review Assistant"
description: "Analyze code for best practices, potential bugs, security issues, and improvement opportunities"
version: "1.0.0"
parameters:
  - name: "language"
    description: "Programming language of the code being reviewed"
    type: "string"
    required: false
  - name: "focus"
    description: "Specific focus area (security, performance, maintainability, or general)"
    type: "string"
    required: false
---

# Code Review Analysis

Please provide a comprehensive code review for the following code:

```{{language}}
{{selection}}
```

## Review Focus
{{focus}}

## Analysis Areas

Please analyze the code for:

1. **Best Practices & Code Quality**
   - Code organization and structure
   - Naming conventions
   - Code readability and maintainability
   - Adherence to language-specific conventions

2. **Potential Issues**
   - Logic errors or bugs
   - Edge cases not handled
   - Resource management issues
   - Error handling gaps

3. **Security Considerations**
   - Input validation
   - Authentication/authorization
   - Data sanitization
   - Common vulnerability patterns

4. **Performance Opportunities**
   - Algorithm efficiency
   - Memory usage optimization
   - Database query optimization
   - Unnecessary computations

5. **Improvement Suggestions**
   - Code refactoring opportunities
   - Design pattern applications
   - Library/framework best practices
   - Testing considerations

## Output Format

For each identified issue or improvement:
- **Category**: (Best Practice/Bug/Security/Performance/Design)
- **Severity**: (Critical/High/Medium/Low)
- **Description**: Clear explanation of the issue
- **Recommendation**: Specific actionable fix or improvement
- **Code Example**: Show the improved version when applicable

Please prioritize findings by severity and provide concrete, actionable recommendations.