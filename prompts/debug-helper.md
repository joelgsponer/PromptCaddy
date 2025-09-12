---
id: "debug-helper-v1"
title: "Debug Assistant"
description: "Analyze code and logs to identify bugs and suggest debugging strategies and fixes"
version: "1.0.0"
parameters:
  - name: "language"
    description: "Programming language of the code being debugged"
    type: "string"
    required: false
  - name: "issue_description"
    description: "Description of the bug or unexpected behavior"
    type: "string"
    required: false
  - name: "error_message"
    description: "Any error messages or stack traces encountered"
    type: "string"
    required: false
---

# Debug Analysis

## Problem Description
{{issue_description}}

## Error Information
{{error_message}}

## Code Under Investigation
```{{language}}
{{selection}}
```

## Debug Analysis Request

Please help identify and resolve the issue by analyzing:

1. **Root Cause Analysis**
   - Potential sources of the bug
   - Code flow analysis
   - Variable state examination
   - Logic error identification

2. **Error Pattern Recognition**
   - Common error patterns in this language
   - Framework-specific issues
   - Environment-related problems
   - Configuration issues

3. **Data Flow Analysis**
   - Input validation problems
   - Data transformation errors
   - State management issues
   - Concurrency problems

4. **Debugging Strategy**
   - Key variables to monitor
   - Breakpoint placement suggestions
   - Logging additions needed
   - Test cases to create

## Expected Output

### Issue Identification
- **Primary Issue**: Main cause of the bug
- **Secondary Issues**: Contributing factors
- **Affected Areas**: Parts of code impacted

### Debugging Steps
1. **Immediate Checks**: Quick verification steps
2. **Data Inspection**: Variables and states to examine
3. **Test Scenarios**: Edge cases to verify
4. **Monitoring Points**: Where to add logging/debugging

### Proposed Solutions
- **Quick Fix**: Immediate solution if possible
- **Proper Fix**: Complete resolution approach
- **Prevention**: How to avoid similar issues

### Code Examples
Show corrected code with explanations:

```{{language}}
// Fixed version with comments explaining changes
```

### Testing Recommendations
- Unit tests to add
- Integration test scenarios
- Edge cases to verify

Please provide a systematic approach to identify, understand, and resolve the debugging issue.