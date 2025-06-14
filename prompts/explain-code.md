---
id: "docs.explain-code.v1"
title: "Explain Code"
description: "Provides a detailed explanation of the selected code."
version: "1.0.0"
parameters:
  - name: "audience"
    description: "The target audience for the explanation (e.g., beginner, intermediate, expert)."
    type: "string"
    required: false
---
Please provide a clear and comprehensive explanation of the following code.

{{#if audience}}
Target audience: {{audience}} developers
{{/if}}

Focus on:
- What the code does at a high level
- How it works step by step
- Key concepts or patterns used
- Any important edge cases or considerations
- Potential improvements or alternatives

Code to explain:

{{selection}}