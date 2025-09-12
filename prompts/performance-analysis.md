---
id: "performance-analysis-v1"
title: "Performance Analysis Assistant"
description: "Analyze code for performance bottlenecks and optimization opportunities"
version: "1.0.0"
parameters:
  - name: "language"
    description: "Programming language of the code to analyze"
    type: "string"
    required: false
  - name: "performance_goal"
    description: "Performance optimization target (speed, memory, throughput, or latency)"
    type: "string"
    required: false
  - name: "context"
    description: "Application context (web-app, mobile-app, batch-processing, or real-time)"
    type: "string"
    required: false
---

# Performance Analysis

## Code Under Analysis
```{{language}}
{{selection}}
```

## Performance Context
- **Language**: {{language}}
- **Optimization Goal**: {{performance_goal}}
- **Application Context**: {{context}}

## Performance Review Request

Please analyze the code for performance optimization opportunities focusing on:

### 1. Algorithm Efficiency
- **Time Complexity**: Big O analysis of algorithms
- **Space Complexity**: Memory usage patterns
- **Algorithm Selection**: More efficient algorithms available
- **Data Structure Optimization**: Better data structure choices
- **Loop Optimization**: Inefficient iteration patterns

### 2. Memory Management
- **Memory Leaks**: Unreleased resources
- **Memory Allocation**: Excessive object creation
- **Garbage Collection**: GC pressure and optimization
- **Memory Pooling**: Object reuse opportunities
- **Buffer Management**: Efficient buffer usage

### 3. I/O Operations
- **Database Queries**: Query optimization opportunities
- **File Operations**: Efficient file handling
- **Network Calls**: API call optimization
- **Caching Strategies**: Data caching opportunities
- **Batch Operations**: Bulk processing improvements

### 4. Concurrency & Parallelism
- **Threading Issues**: Synchronization overhead
- **Parallel Processing**: Parallelization opportunities
- **Async Operations**: Asynchronous improvement areas
- **Resource Contention**: Bottleneck identification
- **Load Balancing**: Work distribution optimization

### 5. Language-Specific Optimizations
- **Compiler Optimizations**: Language-specific improvements
- **Framework Performance**: Framework usage optimization
- **Library Efficiency**: Alternative library suggestions
- **Runtime Optimization**: Runtime-specific improvements
- **Profile-Guided Optimization**: Profiling recommendations

## Output Format

### Performance Assessment
**Overall Performance Rating**: Excellent/Good/Fair/Poor
**Primary Bottlenecks**: Top performance issues identified
**Optimization Potential**: Expected improvement range

### Detailed Analysis

For each performance issue found:

#### Issue Identification
**Performance Issue**: Brief description
**Severity**: Critical/High/Medium/Low
**Category**: Algorithm/Memory/I-O/Concurrency/Language-Specific
**Current Complexity**: Time/Space complexity analysis

#### Impact Analysis
**Performance Impact**: Quantified impact description
**Scalability Issues**: How it affects scaling
**Resource Usage**: CPU/Memory/I-O impact
**User Experience**: End-user impact

#### Evidence
```{{language}}
// Performance bottleneck code section
// Include timing or profiling annotations where helpful
```

#### Optimization Recommendation
**Proposed Solution**: Specific optimization approach
**Expected Improvement**: Performance gain estimate
**Implementation Effort**: Low/Medium/High
**Risk Assessment**: Change complexity and risk

#### Optimized Code
```{{language}}
// Optimized version with performance improvements
// Include comments explaining optimization techniques
```

**Explanation**: Why this optimization works
**Trade-offs**: Any compromises made
**Alternatives**: Other possible optimizations

### Performance Monitoring
**Profiling Recommendations**: Tools and techniques
**Metrics to Track**: Key performance indicators
**Benchmarking**: How to measure improvements
**Continuous Monitoring**: Ongoing performance tracking

### Implementation Strategy
1. **Quick Wins**: Low-effort, high-impact optimizations
2. **Major Optimizations**: Significant performance improvements
3. **Long-term Improvements**: Architectural optimizations
4. **Testing Strategy**: Performance testing approach

Please prioritize optimizations by impact vs. effort ratio and provide concrete performance improvement estimates where possible.