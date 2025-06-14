---
id: "testgen.generate-unit-tests.v1"
title: "Generate Unit Tests"
description: "Generates unit tests for the selected code using the specified framework."
version: "1.0.2"
parameters:
  - name: "framework"
    description: "The testing framework to use (e.g., Jest, Pytest, Go testing)."
    type: "string"
    required: true
---
Your task is to act as an expert Software Developer in Test. Write a comprehensive suite of unit tests for the following code selection, using the **{{framework}}** framework. 

Please ensure your tests:
- Cover all public functions and methods
- Include edge cases and error scenarios
- Follow the framework's best practices and conventions
- Are well-documented with clear test descriptions
- Mock external dependencies where appropriate

The code to test is:

{{selection}}