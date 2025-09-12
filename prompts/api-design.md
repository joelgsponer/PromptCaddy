---
id: "api-design-v1"
title: "API Design Assistant"
description: "Design RESTful APIs following best practices for scalability, security, and usability"
version: "1.0.0"
parameters:
  - name: "api_type"
    description: "Type of API (rest, graphql, grpc, or websocket)"
    type: "string"
    required: false
  - name: "domain"
    description: "Business domain or use case for the API"
    type: "string"
    required: false
  - name: "framework"
    description: "Framework or platform (express, fastapi, spring, gin, or framework-agnostic)"
    type: "string"
    required: false
---

# API Design Assistant

## API Requirements
{{selection}}

## Design Context
- **API Type**: {{api_type}}
- **Business Domain**: {{domain}}
- **Framework**: {{framework}}

## API Design Request

Please help design a comprehensive API following these principles:

### 1. RESTful Design Principles
- **Resource Identification**: Clear resource naming
- **HTTP Method Usage**: Appropriate verb selection
- **Status Code Standards**: Correct HTTP status codes
- **URI Design**: Clean, intuitive URL structures
- **Stateless Operations**: Stateless request/response

### 2. API Structure & Organization
- **Resource Hierarchy**: Logical resource relationships
- **Endpoint Design**: Consistent naming conventions
- **Version Management**: API versioning strategy
- **Collection vs Resource**: Proper resource handling
- **Nested Resources**: Sub-resource relationships

### 3. Data Design
- **Request/Response Format**: JSON structure design
- **Field Naming**: Consistent naming conventions
- **Data Validation**: Input validation requirements
- **Error Responses**: Standardized error formats
- **Pagination**: Large dataset handling

### 4. Security Design
- **Authentication**: Auth mechanism selection
- **Authorization**: Permission and role management
- **Rate Limiting**: Usage control strategies
- **Input Sanitization**: Security validation
- **HTTPS Requirements**: Transport security

### 5. Performance & Scalability
- **Caching Strategy**: Response caching approach
- **Database Optimization**: Efficient data access
- **Async Operations**: Long-running task handling
- **Load Balancing**: Distribution strategies
- **CDN Integration**: Content delivery optimization

## Output Format

### API Overview
**Purpose**: What the API accomplishes
**Target Users**: Who will use this API
**Key Features**: Main functionality provided
**Technical Approach**: Overall architecture strategy

### Resource Design
```yaml
# OpenAPI 3.0 specification format
openapi: 3.0.0
info:
  title: API Name
  version: 1.0.0
  description: API description

paths:
  /resources:
    get:
      summary: List resources
      parameters:
        - name: limit
          in: query
          schema:
            type: integer
            default: 10
      responses:
        '200':
          description: Successful response
          content:
            application/json:
              schema:
                type: object
                properties:
                  data:
                    type: array
                    items:
                      $ref: '#/components/schemas/Resource'
                  pagination:
                    $ref: '#/components/schemas/Pagination'

components:
  schemas:
    Resource:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: string
          format: uuid
        name:
          type: string
```

### Endpoint Documentation

#### GET /api/v1/resources
**Purpose**: Retrieve list of resources
**Authentication**: Required
**Parameters**:
- `limit` (query, optional): Number of items to return (default: 10, max: 100)
- `offset` (query, optional): Number of items to skip (default: 0)
- `filter` (query, optional): Filter criteria

**Response Example**:
```json
{
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "Resource Name",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "total": 100,
    "limit": 10,
    "offset": 0,
    "has_more": true
  }
}
```

**Error Responses**:
- `400 Bad Request`: Invalid parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Access denied
- `500 Internal Server Error`: Server error

### Implementation Guidelines

#### Framework-Specific Implementation
```javascript
// Express.js example
const express = require('express');
const router = express.Router();

router.get('/resources', async (req, res) => {
  try {
    const { limit = 10, offset = 0, filter } = req.query;
    
    // Validate parameters
    if (limit > 100) {
      return res.status(400).json({
        error: 'Limit cannot exceed 100'
      });
    }
    
    // Fetch data
    const resources = await resourceService.list({
      limit: parseInt(limit),
      offset: parseInt(offset),
      filter
    });
    
    res.json({
      data: resources.items,
      pagination: {
        total: resources.total,
        limit: parseInt(limit),
        offset: parseInt(offset),
        has_more: resources.hasMore
      }
    });
  } catch (error) {
    console.error('Error fetching resources:', error);
    res.status(500).json({
      error: 'Internal server error'
    });
  }
});
```

### Security Considerations
**Authentication Strategy**: JWT/OAuth2/API Keys
**Authorization Model**: RBAC/ABAC implementation
**Rate Limiting**: Request throttling rules
**Input Validation**: Request validation middleware
**Security Headers**: CORS, CSP, security headers

### Testing Strategy
**Unit Tests**: Individual endpoint testing
**Integration Tests**: End-to-end API testing
**Load Testing**: Performance validation
**Security Testing**: Vulnerability assessment
**Contract Testing**: API contract validation

### Documentation & Developer Experience
**Interactive Documentation**: Swagger/OpenAPI UI
**Code Examples**: Multiple language examples
**SDK Generation**: Client library generation
**Error Handling Guide**: Error response guide
**Rate Limiting Guide**: Usage limitation docs

### Monitoring & Analytics
**Logging Strategy**: Request/response logging
**Metrics Collection**: Performance metrics
**Error Tracking**: Error monitoring
**Usage Analytics**: API usage tracking
**Health Checks**: Service health endpoints

Please provide a complete API design with implementation examples, security considerations, and documentation guidelines.