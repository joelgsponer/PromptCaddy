---
id: "github-process-issue-v1"
title: "Process issue"
description: "Provides a detailed explanation of how to process an issue. Including planning and branch management."
version: "1.0.0"
parameters:
  - name: "ISSUE_NUMBER"
    description: "The issue number to process"
    type: "integer"
    required: true

---
- Get issue number $ISSUE_NUMBER from github.
- Checkout `dev` and pull latest changes.
- Analyze the current code base.
- Think ultrahard how to solve the issue.
- Perform a web search for relevant documentation if nessesary.
- Formulate a plan.
- Add the plan to the github issue.
- Add labels to the github issue
- Change the issues status to "In Progress" in the repos project page.
- Create and checkout a new branch based on `dev`.
- Reference the issue in the commit message.
- Implment the plan.
- Create a PR (pointing at branch `dev` and close the branch.
- Change the issue status to "In Review" in the repos project page.
- Close the issue.
- Checkout `dev` again

