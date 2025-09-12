---
id: "error-handling-v1"
title: "Error Handling Assistant"
description: "Improve error handling and logging strategies for robust application development"
version: "1.0.0"
parameters:
  - name: "language"
    description: "Programming language of the code"
    type: "string"
    required: false
  - name: "error_type"
    description: "Type of error handling needed (exceptions, async-errors, validation, or logging)"
    type: "string"
    required: false
  - name: "application_type"
    description: "Application context (web-app, cli-tool, library, or microservice)"
    type: "string"
    required: false
---

# Error Handling Enhancement

## Current Code
```{{language}}
{{selection}}
```

## Error Handling Context
- **Language**: {{language}}
- **Error Type Focus**: {{error_type}}
- **Application Type**: {{application_type}}

## Error Handling Analysis Request

Please analyze and improve error handling focusing on:

### 1. Error Detection & Classification
- **Error Types**: Identifying different error categories
- **Error Sources**: System, user, network, business logic errors
- **Error Severity**: Critical, warning, info classification
- **Recovery Potential**: Recoverable vs non-recoverable errors
- **Error Propagation**: How errors flow through the system

### 2. Exception Handling Patterns
- **Try-Catch Blocks**: Proper exception handling
- **Error Wrapping**: Context preservation
- **Custom Exceptions**: Domain-specific error types
- **Error Recovery**: Graceful degradation strategies
- **Resource Cleanup**: Proper resource management

### 3. Validation & Input Handling
- **Input Validation**: Data validation strategies
- **Boundary Checking**: Range and limit validation
- **Type Safety**: Data type validation
- **Business Rule Validation**: Domain constraint checking
- **Early Validation**: Fail-fast principles

### 4. Logging & Monitoring
- **Log Levels**: Appropriate logging levels
- **Structured Logging**: Machine-readable logs
- **Context Information**: Useful debugging info
- **Performance Impact**: Efficient logging
- **Log Security**: Sensitive data protection

### 5. User Experience
- **Error Messages**: User-friendly error communication
- **Error Recovery**: User recovery options
- **Fallback Behavior**: Graceful degradation
- **Progress Indication**: Long-running operation feedback
- **Error Prevention**: Preventing common user errors

## Output Format

### Current Error Handling Assessment
**Error Handling Quality**: Excellent/Good/Fair/Poor
**Main Issues Identified**: Top problems found
**Risk Level**: High/Medium/Low risk assessment
**Improvement Priority**: Critical/Important/Nice-to-have

### Error Handling Improvements

#### Issue Analysis
**Error Handling Gap**: Description of the problem
**Current Behavior**: What happens now
**Risk Assessment**: Potential impact of the issue
**Frequency**: How often this error might occur

**Problematic Code**:
```{{language}}
// Current error handling approach
```

#### Recommended Solution
**Improvement Strategy**: Overall approach
**Error Handling Pattern**: Specific pattern to use
**Recovery Strategy**: How to handle the error gracefully

**Improved Code**:
```{{language}}
// Enhanced error handling implementation
// with detailed comments explaining the approach
```

**Key Improvements**:
- Specific enhancement explanations
- Why this approach is better
- What scenarios it handles

### Comprehensive Error Handling Strategy

#### Error Classification System
```{{language}}
// Custom error types for different scenarios
class ValidationError extends Error {
    constructor(field, value, message) {
        super(message);
        this.name = 'ValidationError';
        this.field = field;
        this.value = value;
        this.timestamp = new Date().toISOString();
    }
}

class BusinessLogicError extends Error {
    constructor(operation, reason, context) {
        super(`${operation} failed: ${reason}`);
        this.name = 'BusinessLogicError';
        this.operation = operation;
        this.reason = reason;
        this.context = context;
    }
}
```

#### Centralized Error Handling
```{{language}}
// Global error handler implementation
function handleError(error, context = {}) {
    const errorInfo = {
        timestamp: new Date().toISOString(),
        type: error.constructor.name,
        message: error.message,
        stack: error.stack,
        context: context,
        severity: determineSeverity(error)
    };
    
    // Log based on severity
    if (errorInfo.severity === 'critical') {
        logger.error('Critical error occurred', errorInfo);
        notifyAdministrators(errorInfo);
    } else if (errorInfo.severity === 'warning') {
        logger.warn('Warning condition detected', errorInfo);
    }
    
    // Return appropriate response
    return formatErrorResponse(errorInfo);
}
```

### Logging Strategy

#### Structured Logging Implementation
```{{language}}
// Structured logging with context
const logger = {
    error: (message, context = {}) => {
        console.error(JSON.stringify({
            level: 'ERROR',
            timestamp: new Date().toISOString(),
            message: message,
            ...context,
            traceId: getTraceId(),
            userId: getCurrentUserId()
        }));
    },
    
    warn: (message, context = {}) => {
        console.warn(JSON.stringify({
            level: 'WARN',
            timestamp: new Date().toISOString(),
            message: message,
            ...context
        }));
    }
};
```

### Validation Framework
```{{language}}
// Comprehensive input validation
function validateInput(data, schema) {
    const errors = [];
    
    for (const [field, rules] of Object.entries(schema)) {
        const value = data[field];
        
        // Required field check
        if (rules.required && (value === undefined || value === null)) {
            errors.push(new ValidationError(field, value, `${field} is required`));
            continue;
        }
        
        // Type validation
        if (value !== undefined && rules.type && typeof value !== rules.type) {
            errors.push(new ValidationError(field, value, `${field} must be of type ${rules.type}`));
            continue;
        }
        
        // Custom validation
        if (rules.validate && !rules.validate(value)) {
            errors.push(new ValidationError(field, value, rules.message || `${field} validation failed`));
        }
    }
    
    if (errors.length > 0) {
        throw new ValidationError('input', data, 'Validation failed', errors);
    }
}
```

### Error Recovery Patterns
```{{language}}
// Retry mechanism with exponential backoff
async function withRetry(operation, maxAttempts = 3, delay = 1000) {
    for (let attempt = 1; attempt <= maxAttempts; attempt++) {
        try {
            return await operation();
        } catch (error) {
            if (attempt === maxAttempts || !isRetryableError(error)) {
                throw error;
            }
            
            logger.warn(`Operation failed, retrying (${attempt}/${maxAttempts})`, {
                error: error.message,
                attempt: attempt,
                nextDelay: delay * Math.pow(2, attempt - 1)
            });
            
            await sleep(delay * Math.pow(2, attempt - 1));
        }
    }
}
```

### Testing Error Scenarios
```{{language}}
// Unit tests for error handling
describe('Error Handling', () => {
    test('should handle validation errors appropriately', () => {
        const invalidInput = { email: 'invalid-email' };
        
        expect(() => validateInput(invalidInput, emailSchema))
            .toThrow(ValidationError);
    });
    
    test('should retry on transient failures', async () => {
        const mockOperation = jest.fn()
            .mockRejectedValueOnce(new NetworkError('Timeout'))
            .mockResolvedValueOnce('Success');
        
        const result = await withRetry(mockOperation);
        
        expect(result).toBe('Success');
        expect(mockOperation).toHaveBeenCalledTimes(2);
    });
});
```

### Monitoring & Alerting
**Error Rate Tracking**: Monitor error frequency and trends
**Error Type Analysis**: Track common error patterns
**Performance Impact**: Monitor error handling overhead
**User Impact Assessment**: Track user-facing errors
**Recovery Success Rate**: Monitor error recovery effectiveness

Please provide comprehensive error handling improvements with concrete code examples and implementation strategies.