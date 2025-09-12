---
id: "refactor-code-v1"
title: "Code Refactoring Assistant"
description: "Suggest refactoring opportunities and clean code improvements for better maintainability"
version: "1.0.0"
parameters:
  - name: "language"
    description: "Programming language of the code to refactor"
    type: "string"
    required: false
  - name: "goal"
    description: "Refactoring goal (readability, performance, testability, or maintainability)"
    type: "string"
    required: false
---

# Code Refactoring Suggestions

Analyzing the following code for refactoring opportunities:

```{{language}}
{{selection}}
```

## Refactoring Goal
Primary focus: {{goal}}

## Analysis

Please provide refactoring suggestions focusing on:

1. **Code Structure & Organization**
   - Function/method decomposition
   - Class design improvements
   - Module organization
   - Separation of concerns

2. **Clean Code Principles**
   - Meaningful naming
   - Function size optimization
   - Reducing complexity
   - Eliminating code duplication

3. **Design Patterns Application**
   - Appropriate pattern usage
   - Dependency injection opportunities
   - Strategy pattern applications
   - Factory pattern usage

4. **Code Smells Elimination**
   - Long parameter lists
   - Large classes/methods
   - Feature envy
   - Data clumps
   - Primitive obsession

5. **Maintainability Improvements**
   - Code documentation
   - Error handling enhancement
   - Configuration externalization
   - Testing improvements

## Output Format

For each refactoring suggestion:

**Refactoring Opportunity**: Brief title
**Current Issue**: Explain what's problematic
**Proposed Solution**: Detailed refactoring approach
**Benefits**: Why this improvement matters
**Code Example**: Show before and after code

### Before:
```{{language}}
[Current code snippet]
```

### After:
```{{language}}
[Refactored code snippet]
```

**Effort Level**: (Low/Medium/High)
**Risk Level**: (Low/Medium/High)

Please prioritize suggestions by impact and provide concrete examples showing the transformation from current to improved code.