---
id: "database-query-v1"
title: "Database Query Assistant"
description: "Generate, optimize, and explain SQL queries for various database operations"
version: "1.0.0"
parameters:
  - name: "database_type"
    description: "Database system (mysql, postgresql, sqlite, sqlserver, or oracle)"
    type: "string"
    required: false
  - name: "query_type"
    description: "Type of query needed (select, insert, update, delete, create, or optimize)"
    type: "string"
    required: false
  - name: "table_schema"
    description: "Database table schema or structure information"
    type: "string"
    required: false
---

# Database Query Assistant

## Query Requirements
{{selection}}

## Database Context
- **Database Type**: {{database_type}}
- **Query Type**: {{query_type}}
- **Table Schema**: {{table_schema}}

## Query Generation Request

Please help with database query tasks focusing on:

### 1. Query Generation
- **Requirement Analysis**: Understanding the data needs
- **SQL Construction**: Building appropriate queries
- **Syntax Optimization**: Database-specific syntax
- **Join Optimization**: Efficient table relationships
- **Subquery Alternatives**: Better query structures

### 2. Query Optimization
- **Performance Analysis**: Identifying slow operations
- **Index Recommendations**: Optimal indexing strategies
- **Query Plan Review**: Execution plan analysis
- **Bottleneck Identification**: Performance issues
- **Rewrite Suggestions**: More efficient alternatives

### 3. Database Best Practices
- **Data Integrity**: Constraints and validation
- **Security**: SQL injection prevention
- **Scalability**: Queries that scale well
- **Maintainability**: Readable and maintainable SQL
- **Error Handling**: Robust query patterns

## Output Format

### Generated Query
```sql
-- Main query with comments explaining key parts
SELECT 
    column1,
    column2,
    -- Explanation of complex expressions
    CASE 
        WHEN condition THEN result
        ELSE alternative
    END as calculated_field
FROM table1 t1
JOIN table2 t2 ON t1.id = t2.foreign_id
WHERE condition
ORDER BY column1;
```

### Query Explanation
**Purpose**: What the query accomplishes
**Logic Flow**: Step-by-step execution explanation
**Key Components**: Important parts breakdown
**Assumptions**: Data assumptions made

### Performance Considerations
**Estimated Performance**: Expected query performance
**Index Requirements**: Indexes needed for optimization
**Scalability Notes**: How query performs with data growth
**Alternative Approaches**: Other ways to solve the same problem

### Query Variations
**Basic Version**: Simplified query for getting started
**Enhanced Version**: Feature-rich query with advanced options
**Optimized Version**: Performance-tuned query

### Schema Recommendations
If table schema is provided or can be inferred:
```sql
-- Recommended table structure
CREATE TABLE example_table (
    id INT PRIMARY KEY AUTO_INCREMENT,
    column1 VARCHAR(255) NOT NULL,
    column2 DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_column1 (column1)
);
```

### Testing Queries
**Sample Data**: Example data for testing
**Test Cases**: Different scenarios to verify
**Validation**: How to confirm query correctness

### Database-Specific Features
- **MySQL**: MySQL-specific optimizations and features
- **PostgreSQL**: Postgres-specific capabilities
- **SQLite**: SQLite limitations and workarounds
- **SQL Server**: T-SQL specific features
- **Oracle**: Oracle-specific syntax and features

### Common Patterns
**Pagination**: Efficient data pagination
**Aggregation**: Grouping and summary operations
**Hierarchical Data**: Tree-like data structures
**Time-based Queries**: Date/time operations
**Full-text Search**: Text search capabilities

Please provide complete, tested SQL queries with clear explanations and optimization recommendations.