---
id: "code-migration-v1"
title: "Code Migration Assistant"
description: "Help migrate code between frameworks, languages, or versions with best practices"
version: "1.0.0"
parameters:
  - name: "source_tech"
    description: "Source technology, framework, or language version"
    type: "string"
    required: false
  - name: "target_tech"
    description: "Target technology, framework, or language version"
    type: "string"
    required: false
  - name: "migration_scope"
    description: "Migration scope (full-rewrite, incremental, or compatibility-layer)"
    type: "string"
    required: false
---

# Code Migration Assistant

## Source Code
```{{source_tech}}
{{selection}}
```

## Migration Context
- **From**: {{source_tech}}
- **To**: {{target_tech}}
- **Migration Scope**: {{migration_scope}}

## Migration Analysis Request

Please help with code migration focusing on:

### 1. Migration Strategy Assessment
- **Compatibility Analysis**: Breaking changes identification
- **Migration Complexity**: Effort and risk estimation
- **Incremental Approach**: Step-by-step migration plan
- **Risk Assessment**: Potential migration pitfalls
- **Rollback Strategy**: Safety measures and rollback plans

### 2. Code Transformation
- **Syntax Conversion**: Language/framework syntax changes
- **API Mapping**: Old API to new API equivalents
- **Design Pattern Updates**: Modern pattern adoption
- **Best Practice Alignment**: Target technology conventions
- **Performance Optimization**: Leverage new capabilities

### 3. Dependency Management
- **Library Migration**: Alternative library identification
- **Version Compatibility**: Dependency version alignment
- **Polyfill Requirements**: Backward compatibility needs
- **Package Updates**: Package manager changes
- **Configuration Updates**: Setup and config changes

### 4. Testing Strategy
- **Test Migration**: Test suite updates needed
- **Compatibility Testing**: Cross-version validation
- **Regression Testing**: Functionality preservation
- **Performance Testing**: Performance comparison
- **Integration Testing**: End-to-end validation

### 5. Deployment Considerations
- **Environment Updates**: Infrastructure changes needed
- **CI/CD Pipeline**: Build and deployment updates
- **Database Migration**: Schema and data changes
- **Configuration Management**: Environment settings
- **Monitoring Updates**: Observability adjustments

## Output Format

### Migration Overview
**Migration Type**: Framework/Language/Version upgrade
**Complexity Level**: Simple/Moderate/Complex/Major rewrite
**Estimated Effort**: Time and resource estimation
**Risk Level**: Low/Medium/High migration risk
**Success Probability**: Likelihood of smooth migration

### Pre-Migration Analysis
**Breaking Changes**: Major compatibility issues
**Deprecated Features**: Features removed in target
**New Opportunities**: Benefits of target technology
**Migration Blockers**: Issues that prevent migration
**Prerequisites**: Requirements before starting

### Step-by-Step Migration Plan

#### Phase 1: Preparation
1. **Environment Setup**: Target technology setup
2. **Dependency Analysis**: Library compatibility check
3. **Test Suite Preparation**: Existing test preservation
4. **Backup Strategy**: Code and data backup plan

#### Phase 2: Core Migration
1. **Critical Path**: Most important components first
2. **Incremental Changes**: Small, testable changes
3. **Validation Points**: Check migration success
4. **Rollback Points**: Safe rollback opportunities

#### Phase 3: Optimization
1. **Performance Tuning**: Leverage new capabilities
2. **Code Cleanup**: Remove legacy workarounds
3. **Best Practice Adoption**: Modern patterns
4. **Documentation Updates**: Updated documentation

### Code Transformation Examples

#### Original Code ({{source_tech}})
```{{source_tech}}
// Current implementation with comments explaining the approach
```

**Analysis**: What this code does and why it needs migration

#### Migrated Code ({{target_tech}})
```{{target_tech}}
// Migrated implementation with modern best practices
// Comments explaining migration decisions and improvements
```

**Migration Notes**:
- Key changes made during migration
- Why specific approaches were chosen
- New features leveraged
- Performance improvements gained

#### Alternative Approaches
```{{target_tech}}
// Alternative migration approach if applicable
// Showing different ways to solve the same problem
```

### Migration Checklist

#### Code Changes
- [ ] Syntax updated to target technology
- [ ] APIs replaced with target equivalents
- [ ] Error handling updated
- [ ] Logging and monitoring updated
- [ ] Configuration format updated

#### Testing Requirements
- [ ] Unit tests migrated and updated
- [ ] Integration tests validated
- [ ] Performance benchmarks created
- [ ] Compatibility testing completed
- [ ] User acceptance testing planned

#### Documentation Updates
- [ ] Code comments updated
- [ ] API documentation revised
- [ ] Deployment guides updated
- [ ] Troubleshooting guides revised
- [ ] Migration guide created

### Common Migration Patterns

#### Configuration Migration
```{{target_tech}}
// Configuration format transformation
// Old format -> New format mapping
```

#### API Endpoint Migration
```{{target_tech}}
// API route definition changes
// Middleware updates
// Request/response handling updates
```

#### Database Integration Updates
```{{target_tech}}
// Database connection and query updates
// ORM or query builder changes
// Migration script examples
```

### Risk Mitigation Strategies

#### Technical Risks
**Breaking Changes**: How to handle incompatible changes
**Performance Regression**: Performance monitoring approach
**Data Loss**: Data preservation strategies
**Integration Failures**: Third-party integration updates

#### Operational Risks
**Deployment Issues**: Safe deployment strategies
**Rollback Procedures**: Quick rollback plans
**Team Training**: Knowledge transfer requirements
**Timeline Risks**: Schedule and milestone management

### Validation Strategy

#### Pre-Migration Validation
- Code analysis and compatibility checking
- Dependency conflict resolution
- Test suite execution
- Performance baseline establishment

#### Post-Migration Validation
- Functionality verification testing
- Performance comparison analysis
- Integration testing execution
- User acceptance testing

### Long-term Maintenance

#### Code Maintenance
**Modernization Opportunities**: Further improvements
**Technical Debt Reduction**: Clean up opportunities
**Performance Optimization**: Ongoing improvements
**Security Updates**: Security enhancement opportunities

#### Process Improvements
**Development Workflow**: Updated development practices
**CI/CD Pipeline**: Automated testing and deployment
**Monitoring Strategy**: Enhanced observability
**Documentation Maintenance**: Keeping docs current

Please provide a comprehensive migration plan with concrete code examples, risk assessment, and step-by-step implementation guidance.