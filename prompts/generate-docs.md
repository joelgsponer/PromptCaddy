---
id: "generate-docs-v1"
title: "Documentation Generator"
description: "Generate comprehensive documentation for code, APIs, and software components"
version: "1.0.0"
parameters:
  - name: "language"
    description: "Programming language of the code to document"
    type: "string"
    required: false
  - name: "doc_type"
    description: "Type of documentation needed (api, code, readme, or user-guide)"
    type: "string"
    required: false
  - name: "audience"
    description: "Target audience (developer, end-user, or technical-writer)"
    type: "string"
    required: false
---

# Documentation Generation

## Code to Document
```{{language}}
{{selection}}
```

## Documentation Requirements
- **Type**: {{doc_type}}
- **Target Audience**: {{audience}}
- **Language**: {{language}}

## Documentation Request

Please generate comprehensive documentation following these guidelines:

### For API Documentation
1. **Overview**: Purpose and functionality
2. **Endpoints/Methods**: Complete interface description
3. **Parameters**: Input requirements and types
4. **Return Values**: Output format and examples
5. **Error Handling**: Possible errors and responses
6. **Usage Examples**: Practical code examples
7. **Authentication**: Security requirements if applicable

### For Code Documentation
1. **Module/Class Overview**: High-level purpose
2. **Function/Method Documentation**: 
   - Purpose and behavior
   - Parameters with types and descriptions
   - Return values
   - Exceptions/errors
   - Complexity notes
3. **Usage Examples**: How to use the code
4. **Dependencies**: Required imports/libraries
5. **Performance Notes**: Time/space complexity where relevant

### For README Documentation
1. **Project Overview**: What it does and why
2. **Installation**: Setup instructions
3. **Quick Start**: Basic usage examples
4. **Configuration**: Setup options
5. **Examples**: Common use cases
6. **Contributing**: Development guidelines
7. **License**: Legal information

### For User Guide Documentation
1. **Getting Started**: Initial setup
2. **Features**: Functionality overview
3. **Step-by-Step Guides**: Common workflows
4. **Troubleshooting**: Common issues and solutions
5. **FAQ**: Frequently asked questions
6. **Advanced Usage**: Power user features

## Output Format

Please structure the documentation with:
- Clear headings and sections
- Code examples with syntax highlighting
- Table of contents for longer documents
- Cross-references where applicable
- Consistent formatting and style

### Code Examples Format:
```{{language}}
// Example with clear comments
function example() {
    // Implementation details
}
```

### API Examples Format:
```http
GET /api/endpoint
Content-Type: application/json

{
    "example": "request"
}
```

Please ensure the documentation is:
- **Complete**: Covers all functionality
- **Accurate**: Reflects actual behavior
- **Clear**: Easy to understand for the target audience
- **Maintainable**: Structured for easy updates