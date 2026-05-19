---
description: Cross-functional review that consults the right roles
agent: head-of-product
subtask: true
---
Review this request:

$ARGUMENTS

First, classify the request into the relevant domains. Possible domains include:
- product and roadmap
- architecture and implementation
- security and compliance
- UX and design
- field operations
- onboarding and migration
- AI, OCR, and automation
- company strategy

Consult only the roles that are materially relevant to the request.

Consult the CTO for architecture, APIs, data models, sync correctness, infrastructure, reliability, maintainability, or major implementation tradeoffs.

Consult the Security and Compliance Officer for authentication, authorization, permissions, tenant isolation, auditability, compliance, financial data safety, immutable records, or recovery-sensitive decisions.

Consult the Head of Field Operations for dispatcher workflows, technician workflows, field usability, training burden, and operational edge cases under real-world pressure.

Consult the Head of Implementation and Migration for onboarding, imports, OCR cleanup, data migration, duplicate handling, rollout safety, and go-live risk.

Consult the Chief UX Officer for cognitive load, workflow clarity, onboarding friction, interaction simplicity, and calm operational UX.

Consult the Designer for layouts, review screens, information hierarchy, interface patterns, and visual clarity.

Consult the AI Systems Lead for OCR, extraction, confidence scoring, human-in-the-loop review, automation, and AI-related trust or recovery concerns.

Consult the CEO only when the request involves wedge strategy, market positioning, pricing, major roadmap prioritization, what not to build yet, company-level tradeoffs, or expansion beyond the current product wedge.

Do not consult unnecessary roles just for completeness.

If the request is primarily technical architecture, let the CTO's view carry extra weight on architecture constraints.
If the request is primarily security-sensitive, let the Security and Compliance Officer's constraints override convenience.
If the request is primarily onboarding or migration-sensitive, let the Head of Implementation and Migration's view carry extra weight on rollout safety.
If the request is primarily workflow or adoption-sensitive, give extra weight to the Head of Field Operations and Chief UX Officer.
If the request is primarily strategic, give extra weight to the CEO and Head of Product.

Agents should disagree when appropriate. Do not converge prematurely.
If consulted roles disagree, summarize the disagreement clearly before recommending a path.

Default to the smallest useful recommendation that fits the company mission, product philosophy, and bootstrap constraints in AGENTS.md.

Return:
## Recommendation

## Why

## Roles Consulted

## Points of Disagreement

## Key Tradeoffs

## Risks

## Lean Next Step
