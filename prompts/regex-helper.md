---
id: "regex-helper-v1"
title: "Regular Expression Helper"
description: "Generate, explain, and optimize regular expressions for text processing tasks"
version: "1.0.0"
parameters:
  - name: "regex_task"
    description: "Type of regex task (generate, explain, optimize, or test)"
    type: "string"
    required: false
  - name: "regex_flavor"
    description: "Regex flavor or engine (javascript, python, java, pcre, or posix)"
    type: "string"
    required: false
  - name: "complexity"
    description: "Desired complexity level (simple, moderate, or advanced)"
    type: "string"
    required: false
---

# Regular Expression Helper

## Task Description
{{selection}}

## Regex Context
- **Task Type**: {{regex_task}}
- **Regex Flavor**: {{regex_flavor}}
- **Complexity Level**: {{complexity}}

## Regular Expression Request

Please help with regex tasks focusing on:

### 1. Pattern Generation
- **Text Matching**: Creating patterns for specific text
- **Validation Patterns**: Input validation regex
- **Extraction Patterns**: Data extraction from text
- **Replacement Patterns**: Find and replace operations
- **Complex Patterns**: Multi-part pattern matching

### 2. Pattern Explanation
- **Component Breakdown**: Explaining each part of the regex
- **Logic Flow**: How the pattern matches
- **Edge Cases**: What the pattern handles
- **Limitations**: What the pattern doesn't cover
- **Performance Notes**: Efficiency considerations

### 3. Pattern Optimization
- **Performance Improvement**: Faster matching patterns
- **Readability Enhancement**: Clearer pattern structure
- **Maintainability**: Easier to modify patterns
- **Cross-Platform**: Compatible across regex engines
- **Memory Efficiency**: Lower memory usage patterns

### 4. Testing & Validation
- **Test Cases**: Example matches and non-matches
- **Edge Case Testing**: Boundary condition testing
- **Performance Testing**: Speed and efficiency tests
- **Cross-Platform Testing**: Different regex engine behavior
- **Regression Testing**: Pattern change validation

## Output Format

### Regex Solution

#### Generated Pattern
```regex
# Main regex pattern with comments
(?<protocol>https?)://           # Protocol (http or https)
(?<domain>[\w.-]+)              # Domain name
(?::\d+)?                       # Optional port number
(?<path>/[\w/.-]*)?             # Optional path
(?:\?(?<query>[\w&=%-]*))?      # Optional query string
(?:#(?<fragment>[\w%-]*))?      # Optional fragment
```

**Pattern Explanation**:
- `(?<protocol>https?)`: Named capture group for protocol, matches 'http' or 'https'
- `://`: Literal characters for protocol separator
- `(?<domain>[\w.-]+)`: Domain name with word characters, dots, and hyphens
- `(?::\d+)?`: Non-capturing group for optional port number
- Each component serves a specific matching purpose

#### Usage Examples

**JavaScript Implementation**:
```javascript
const urlRegex = /^(?<protocol>https?):\/\/(?<domain>[\w.-]+)(?::\d+)?(?<path>\/[\w\/.-]*)?(?:\?(?<query>[\w&=%-]*))?(?:#(?<fragment>[\w%-]*))?$/;

// Test the pattern
const testUrls = [
    'https://example.com',
    'http://subdomain.example.com:8080/path?query=value#fragment',
    'https://invalid..domain.com'  // Should not match
];

testUrls.forEach(url => {
    const match = url.match(urlRegex);
    if (match) {
        console.log(`✓ Matched: ${url}`);
        console.log('Groups:', match.groups);
    } else {
        console.log(`✗ No match: ${url}`);
    }
});

// Extract data using the pattern
function parseUrl(url) {
    const match = url.match(urlRegex);
    return match ? match.groups : null;
}
```

**Python Implementation**:
```python
import re

url_pattern = r'^(?P<protocol>https?):\/\/(?P<domain>[\w.-]+)(?::\d+)?(?P<path>\/[\w\/.-]*)?(?:\?(?P<query>[\w&=%-]*))?(?:#(?P<fragment>[\w%-]*))?$'

# Compile for better performance
url_regex = re.compile(url_pattern)

def parse_url(url):
    match = url_regex.match(url)
    return match.groupdict() if match else None

# Test examples
test_urls = [
    'https://example.com',
    'http://subdomain.example.com:8080/path?query=value#fragment'
]

for url in test_urls:
    result = parse_url(url)
    if result:
        print(f"✓ Parsed: {url}")
        for key, value in result.items():
            if value:
                print(f"  {key}: {value}")
    else:
        print(f"✗ Failed to parse: {url}")
```

### Pattern Breakdown & Analysis

#### Character Classes Used
- `\w`: Word characters (letters, digits, underscore)
- `\d`: Digit characters (0-9)
- `\.`: Escaped dot (literal period)
- `[...]`: Character class (custom character set)

#### Quantifiers Explained
- `?`: Zero or one occurrence (optional)
- `+`: One or more occurrences (required)
- `*`: Zero or more occurrences (optional, any amount)
- `{n,m}`: Between n and m occurrences

#### Special Constructs
- `(...)`: Capturing group
- `(?:...)`: Non-capturing group
- `(?<name>...)`: Named capturing group
- `^`: Start of string anchor
- `$`: End of string anchor

### Testing Strategy

#### Test Cases
```javascript
const testCases = [
    // Valid matches
    { input: 'https://example.com', shouldMatch: true, description: 'Basic HTTPS URL' },
    { input: 'http://subdomain.example.com:8080', shouldMatch: true, description: 'HTTP with subdomain and port' },
    { input: 'https://example.com/path/to/resource', shouldMatch: true, description: 'URL with path' },
    
    // Invalid cases
    { input: 'ftp://example.com', shouldMatch: false, description: 'FTP protocol should not match' },
    { input: 'https://', shouldMatch: false, description: 'Missing domain' },
    { input: 'not-a-url', shouldMatch: false, description: 'Plain text should not match' }
];

// Run tests
testCases.forEach(test => {
    const match = urlRegex.test(test.input);
    const result = match === test.shouldMatch ? '✓ PASS' : '✗ FAIL';
    console.log(`${result}: ${test.description} - "${test.input}"`);
});
```

### Performance Considerations

#### Optimization Techniques
**Anchoring**: Using `^` and `$` for full string matching
**Non-capturing Groups**: Using `(?:...)` when capture isn't needed
**Character Classes**: Efficient character matching with `[...]`
**Lazy vs Greedy**: Choosing appropriate quantifier behavior

#### Performance Comparison
```javascript
// Less efficient (backtracking issues)
const slowPattern = /^https?:\/\/.+/;

// More efficient (specific matching)
const fastPattern = /^https?:\/\/[\w.-]+/;

// Benchmark example
function benchmarkRegex(pattern, testString, iterations = 100000) {
    const start = performance.now();
    for (let i = 0; i < iterations; i++) {
        pattern.test(testString);
    }
    const end = performance.now();
    return end - start;
}
```

### Common Regex Patterns Library

#### Email Validation
```regex
^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
```

#### Phone Number (US Format)
```regex
^\+?1?[-.\s]?\(?([0-9]{3})\)?[-.\s]?([0-9]{3})[-.\s]?([0-9]{4})$
```

#### Date Validation (YYYY-MM-DD)
```regex
^\d{4}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12]\d|3[01])$
```

#### Password Strength
```regex
^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$
```

### Troubleshooting Guide

#### Common Issues
**Catastrophic Backtracking**: Patterns that cause exponential time complexity
**Greedy vs Lazy**: Unexpected matching behavior
**Character Escaping**: Special characters not properly escaped
**Case Sensitivity**: Forgetting case-insensitive flags
**Anchor Misuse**: Incorrect use of `^` and `$`

#### Debugging Techniques
**Pattern Testing**: Use online regex testers
**Step-by-Step Building**: Build patterns incrementally
**Verbose Mode**: Use comments and whitespace (where supported)
**Capture Group Analysis**: Examine what each group captures
**Performance Profiling**: Measure pattern performance

### Advanced Techniques

#### Lookahead/Lookbehind
```regex
# Positive lookahead: match 'cat' only if followed by 'dog'
cat(?=dog)

# Negative lookahead: match 'cat' only if NOT followed by 'dog'
cat(?!dog)
```

#### Conditional Patterns
```regex
# Match different patterns based on a condition
(?(condition)yes-pattern|no-pattern)
```

Please provide comprehensive regex solutions with clear explanations, test cases, and performance considerations.