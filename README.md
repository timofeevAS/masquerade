# About
## masquerade
masquerade - reversible PII anonymization proxy between your harness and the LLM.

Based for dev-ops university course.

## Goal
LLM harness agents can send sensitive data out (Anthropic, Google, OpenAI): file contents, command output, user data.

Goal of masquarade is to hide sensitive data from third parties.

## Proposed schema
```mermaid
flowchart LR
    subgraph LOCAL["Trusted 'local'"]
        direction LR
        H["Harness"]
        P["masquerade<br/>local proxy"]
    end
    subgraph THIRD["Third party"]
        L["External LLM"]
    end

    H -- "(1) real data" --> P
    P -- "(2) anonymized" --> L
    L -- "(3) response" --> P
    P -- "(4) restored" --> H
```

## Other information
WIP


### Authors

- [@mariohuq](https://github.com/mariohuq)
- [@timofeevAS](https://github.com/timofeevAS)