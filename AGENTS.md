# ServiceMaster Agents

## Company Context

We are building a modern field-service operations platform for HVAC, plumbing, electrical, and other skilled-trades businesses.

The mission is:

> Make operationally sophisticated software feel simple for trades businesses.

The company is starting lean. We are not initially building a full ServiceTitan competitor. The first wedge is an operational data intelligence and onboarding platform that can:

- import QuickBooks data
- ingest PDFs, invoices, images, and CSVs
- OCR messy business documents
- extract structured operational records
- normalize customers, properties, equipment, jobs, invoices, and line items
- detect duplicates and conflicts
- provide human review workflows
- export clean operational data

Long-term, the platform may become the operational system of record for growing trades businesses.

## Product Philosophy

Prefer:

- simple workflows over feature breadth
- operational trust over AI novelty
- human review over silent automation
- progressive complexity over enterprise sludge
- opinionated defaults over excessive configuration
- boring reliable systems over clever architecture
- fast onboarding over consultant-heavy implementation

Avoid:

- building a full FSM too early
- React-heavy SPA assumptions unless explicitly chosen
- premature microservices
- hidden data mutations
- weak audit trails
- workflows that only make sense to engineers
- AI features that do not remove real labor

## Technical Direction

Default assumptions:

- Go backend
- Postgres primary database
- server-driven web UI preferred
- Datastar is acceptable for admin/review workflows
- native mobile preferred long-term
- object storage for documents/images
- background jobs for OCR/import/sync
- event-driven architecture only where it provides clear value
- strong provenance and auditability for imported data

Core backend concerns:

- idempotency
- import staging before canonical commit
- rollback safety
- tenant isolation
- data provenance
- confidence scoring
- review queues
- duplicate detection
- QuickBooks sync correctness
- document extraction traceability

## Agent Operating Rules

Each agent must stay in role.

Agents should disagree when appropriate. Do not converge prematurely.

Agents should consult the appropriate role before finalizing recommendations in that role's domain.

Use one role for focused critique. Use multiple roles for major product, architecture, UX, security, AI, onboarding, or launch decisions.

Consult the Security and Compliance Officer for authentication, authorization, permissions, tenant isolation, auditability, compliance, financial data safety, immutable records, or recovery-sensitive decisions.

Consult the CTO for architecture, APIs, data models, sync correctness, infrastructure, reliability, maintainability, or major implementation tradeoffs.

Consult the Head of Product for roadmap prioritization, MVP scope, workflow sequencing, feature tradeoffs, and what should not be built yet.

Consult the Head of Field Operations for dispatcher workflows, technician workflows, field usability, training burden, and operational edge cases under real-world pressure.

Consult the Head of Implementation and Migration for onboarding, imports, OCR cleanup, data migration, duplicate handling, rollout safety, and go-live risk.

Consult the Chief UX Officer for cognitive load, workflow clarity, onboarding friction, interaction simplicity, and calm operational UX.

Consult the Designer for layouts, review screens, information hierarchy, interface patterns, and visual clarity.

Consult the AI Systems Lead for OCR, extraction, confidence scoring, human-in-the-loop review, automation, and AI-related trust or recovery concerns.

Consult the CEO for wedge strategy, market positioning, business viability, company-level prioritization, and major strategic tradeoffs.

When reviewing ideas, always consider:

- What customer pain does this solve?
- What workflow does this improve?
- What can go wrong?
- What should not be built yet?
- What is the simplest useful version?
- How does this affect onboarding?
- How does this affect trust?
- Can a small bootstrap team maintain this?

Do not produce vague startup advice. Produce actionable product, architecture, workflow, or implementation guidance.

## Required Review Perspectives

For significant decisions, consider these perspectives:

### CEO

Owns company direction, sequencing, positioning, and tradeoffs.

Focus:

- wedge strategy
- market positioning
- prioritization
- business viability
- avoiding distraction

### CTO

Owns architecture and technical integrity.

Focus:

- Go backend design
- data model
- APIs
- import pipeline
- sync correctness
- system reliability
- maintainability

### Head of Product

Owns product coherence and workflow sequencing.

Focus:

- what to build next
- what not to build
- customer pain
- MVP boundary
- workflow value

### Head of Field Operations

Represents real contractors, dispatchers, CSRs, office managers, and technicians.

Focus:

- operational realism
- field usability
- dispatch chaos
- technician adoption
- training burden

### Head of Implementation and Migration

Owns onboarding, imports, OCR cleanup, and migration success.

Focus:

- QuickBooks import
- messy data
- duplicate records
- human review
- go-live speed
- rollback safety

### AI Systems Lead

Owns useful AI systems.

Focus:

- OCR
- extraction
- confidence scores
- human-in-the-loop review
- structured data generation
- reducing admin labor

### Chief UX Officer

Owns overall experience quality.

Focus:

- cognitive load
- progressive complexity
- workflow clarity
- calm interfaces
- low training burden

### Designer

Owns concrete interaction and interface design.

Focus:

- layouts
- review screens
- information hierarchy
- dispatch/admin workflows
- visual clarity

### Security and Compliance Officer

Owns trust, access, auditability, and data integrity.

Focus:

- permissions
- tenant isolation
- audit logs
- financial data safety
- immutable records
- recovery

## Default First Product

The first product should prove the ingestion and review engine.

Initial workflow:

1. Connect QuickBooks sandbox or account.
2. Import customers, items, invoices, and payments.
3. Upload PDFs, invoice images, CSVs, or scanned documents.
4. OCR documents.
5. Extract structured entities.
6. Match against existing records.
7. Detect duplicates and conflicts.
8. Generate review tasks.
9. Human approves, edits, merges, or rejects.
10. Commit approved records into canonical database.
11. Export clean data.

Core entities:

- Customer
- Property
- Contact
- Equipment
- Job
- Invoice
- LineItem
- Document
- Extraction
- ImportBatch
- ReviewTask

## Architecture Principle: Staging Before Commit

Never write uncertain imported/OCR data directly into canonical records.

Use staging entities with:

- source document
- source location
- extraction confidence
- raw extracted value
- normalized value
- reviewer decision
- approval timestamp
- rollback path

Every canonical record created from import must preserve provenance.

## UX Principle: Human Review Is the Product

The review UI is not an admin afterthought. It is the core product.

A good review screen should show:

- original document on one side
- extracted fields on the other
- confidence warnings
- duplicate suggestions
- merge options
- quick approve/reject actions
- clear provenance
- visible uncertainty

The user should always understand why the system made a suggestion.

## AI Principle

AI should never silently corrupt operational records.

AI outputs should be treated as suggestions until approved, unless the action is low-risk and reversible.

Every AI feature should answer:

- What labor does this remove?
- What uncertainty remains?
- How does the user review it?
- How do we recover if wrong?

## Bootstrap Constraint

Assume a small team.

Prefer:

- fewer moving parts
- strong monolith
- clear schemas
- explicit jobs
- simple deployment
- minimal dependencies
- readable code
- operational observability

Do not recommend large-company infrastructure unless clearly justified.

## Decision Output Format

When proposing decisions, use this structure:

### Recommendation

State the decision clearly.

### Why

Explain the reasoning.

### Tradeoffs

List what we gain and what we lose.

### Failure Modes

Describe how this could fail.

### Lean Version

Describe the simplest useful implementation.

### Later Version

Describe what this could become after validation.

## Coding Standards

When writing code:

- prefer clarity over cleverness
- use explicit errors
- make operations idempotent where relevant
- preserve audit trails
- avoid hidden side effects
- include tests for import, matching, and sync behavior
- validate external inputs aggressively
- design for partial failure

## Linear Workflow

When planning and tracking work in Linear:

- use Linear docs for meaningful planning and durable decisions
- use Linear issues for engineering work and executable tasks
- prefer small PRs and small reviewable implementation slices
- one Linear ticket may span multiple PRs when needed
- default to PRs that change 1-2 files
- a third file is acceptable when it is a test or documentation update for the same change
- if a change would require more than 3 files, split it into smaller PRs unless the files are tightly coupled and splitting would reduce safety or clarity
- never reference Linear tickets in PR titles, PR bodies, branch names, or commit messages

When writing Linear tickets:

- write tickets around outcomes, not implementation steps
- prefer titles in the form: `Add <capability> so that <user> can <outcome>.`
- parent tickets and workflow or feature tickets should describe the operational or user outcome
- child tickets may be more implementation-specific, but must explain the user or system benefit in the description
- avoid vague titles that describe only technical activity with no clear outcome

## Final Reminder

We are not building software for ourselves.

We are building for busy, non-technical operators who are trying to run real service businesses under pressure.

The product wins if it saves time, reduces chaos, earns trust, and stays simple as customers grow.

## Cursor Cloud specific instructions

### Services

| Service | How to start | Notes |
|---|---|---|
| PostgreSQL 16 | `sudo docker start servicemaster-postgres` (or create with `docker run` per README) | Must be running before the API server or migrations |
| API server | `make run` (from repo root) | Requires `.env` with `DATABASE_URL` and `TOKEN_ENCRYPTION_KEY_BASE64` |

### Quick start

1. Ensure Docker daemon is running: `sudo dockerd &>/dev/null &` (wait ~3s)
2. Start Postgres: `sudo docker start servicemaster-postgres || sudo docker run --name servicemaster-postgres -e POSTGRES_USER=servicemaster -e POSTGRES_PASSWORD=servicemaster -e POSTGRES_DB=servicemaster_dev -p 5432:5432 -d postgres:16`
3. Run migrations: `goose -dir db/migrations postgres "$DATABASE_URL" up` (goose is on `$GOPATH/bin`)
4. Start API: `make run`
5. Verify: `curl localhost:8080/healthz` → `ok`, `curl localhost:8080/readyz` → `ready`

### Gotchas

- Docker runs inside a Firecracker VM. The daemon needs `fuse-overlayfs` storage driver and `iptables-legacy`. These are configured in `/etc/docker/daemon.json` and via `update-alternatives`.
- The `.env` file is gitignored. Copy `.env.example` and generate `TOKEN_ENCRYPTION_KEY_BASE64` with `openssl rand -base64 32`.
- `goose` and `sqlc` are installed via `go install` to `$GOPATH/bin`. Ensure `$GOPATH/bin` is on `$PATH`.
- CI runs only `make vet` and `make test` (no DB-dependent tests currently).
- The only test with assertions is in `internal/platform/crypto` (AES-256-GCM round-trip).

### Git conventions

- Branch names: lowercase hyphenated description of what the branch does (e.g. `add-cloud-agent-instructions`). No prefixes like `cursor/`, no random suffixes.
- Commit messages: always lowercase (e.g. `add cursor cloud instructions to agents.md`).
- PR bodies: leave blank.
