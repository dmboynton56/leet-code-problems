# Sherpa 6 GoLang Technical Interview Prep Pack

Interview: Monday, June 29, 2026, 11:30 AM MDT  
Format: Microsoft Teams, expected 60-90 minutes  
Role: GoLang Developer - Secret Clearance Required  
Recruiter thread: Jessica Lucido, Sherpa 6  
Interviewers named in thread: Brian Cook, Daniel Nelson, Colleen Vinomano  

Use the calendar invite for the Teams link. Log in 5 minutes early.

## Source Notes

- Jessica Lucido email, June 25, 2026: prepare to discuss GoLang background and complete a live coding challenge in an IDE of choice.
- Attached resume: Go listed as a language; strongest matching project is Senior Capstone: AI Enabled Community Platform using React, Go, OpenAI/HuggingFace, tests, CI, deployment docs, and runbooks.
- Job posting: [GoLang Developer - Secret Clearance Required](https://recruiting.paylocity.com/Recruiting/Jobs/Details/3971839)
- Company site: [Sherpa 6 home](https://sherpa6.com/) and [Sherpa 6 about/mission](https://sherpa6.com/about/#mission)
- Glassdoor pages reviewed in Chrome:
  - [Sherpa 6 interview questions](https://www.glassdoor.com/Interview/Sherpa-6-Interview-Questions-E2962220.htm)
  - [Sherpa 6 Software Engineer interview questions](https://www.glassdoor.com/Interview/Sherpa-6-Software-Engineer-Interview-Questions-EI_IE2962220.0,8_KO9,26.htm)

## Highest Probability Interview Shape

Glassdoor reports for Sherpa 6 software roles point to this pattern:

1. Recruiter/background screen: experience, resume walkthrough, salary/logistics, clearance/travel fit.
2. Technical interview: coding challenge plus conceptual questions.
3. Hiring manager or panel interview: behavioral questions, team interaction, role fit.

For your specific invite, expect the June 29 interview to combine steps 2 and part of step 3: Go background, live coding, conceptual backend questions, and a few team/mission-fit questions.

## Role Targets From The Posting

The posting is backend-heavy. Prepare to speak to these in practical terms:

- Go REST API development in a microservices architecture.
- Relational databases: PostgreSQL, MySQL, schema design, queries, indexes, transactions.
- Unit tests, design documentation, maintainable and fault-tolerant code.
- Software engineering principles: readable code, decomposition, reuse, debugging, version control.
- Tooling: Git, Jira/Trello-style workflows, Linux/Windows scripting.
- Preferred differentiators: Go concurrency patterns, STIG-compliant environments, client/server certificates, certificate authorities, client certificate authentication, encryption/cryptography basics, tactical radio APIs.

Watch the experience gap honestly. The posting asks for 5+ years software engineering and 3+ years Go. Your resume is strongest on CS fundamentals, backend/API thinking, ML/data systems, testing/CI, and one major Go project. Do not oversell years. Instead, show you can ramp quickly and write correct, tested Go under interview pressure.

## 90-Second Opening Pitch

"I am a software engineer with a CS degree from CU Boulder and a strong backend/data systems foundation. My recent work combines production-style API development, ML systems, and Go. On my senior capstone, I helped build an AI-enabled community platform in React and Go, including a resume parser, chat summarizer, tests, CI pipelines, deployment docs, and runbooks. I have also built data-backed APIs like Sports Edge, where I worked with large datasets, prediction models, BigQuery, and Supabase. For this role, the strongest match is backend service development: REST APIs, testing, documentation, data modeling, and learning mission-specific constraints quickly. I am especially interested in Sherpa 6 because the work supports mission-critical users, and I like building practical systems where reliability and clarity matter."

## Resume Stories To Have Ready

### Go / Backend Story

Use: Senior Capstone, AI Enabled Community Platform  
Emphasize: Go service design, API boundaries, parsing/summarization workflow, tests, CI, runbook docs.  
Likely question: "Tell us about your Go background."  
Answer shape:

- Problem: needed a platform service that supported resume parsing and chat summarization.
- Role: owned or contributed to backend/API logic and quality practices.
- Technical choices: Go for service logic, React frontend, OpenAI/HF integrations, tests and CI.
- Reliability: documented deployment and runbooks, wrote tests around key flows.
- Lesson: clear interfaces and testable units matter when integrating external AI services.

### Data/API Story

Use: Sports Edge API  
Emphasize: API design, data pipeline thinking, BigQuery, Supabase, real-time predictions, model outputs exposed to users.  
Likely question: "How have you worked with databases or APIs?"  
Answer shape:

- Problem: serve sports predictions across several leagues.
- Data: large-scale data in BigQuery; persisted/served results through Supabase.
- Engineering: model output had to be queryable, stable, and understandable.
- Lesson: define schemas and contracts carefully before optimizing.

### Testing / CI Story

Use: Capstone and portfolio GitHub Actions  
Emphasize: unit tests, automated workflows, deployment/runbook docs.  
Likely question: "How do you know your code works?"  
Answer shape:

- Unit-test core behavior, then add integration checks around API/data boundaries.
- Use CI to prevent regressions.
- Document run and deploy procedures so another engineer can reproduce the system.

### Reliability / Teamwork Story

Use: seasonal service and field work plus client freelance work  
Emphasize: reliability, requirements gathering, practical execution, team-based work.  
Likely question: "Tell me about working under ambiguity or with nontechnical stakeholders."  
Answer shape:

- Freelance clients: requirements were not always precise; clarify what success looks like.
- Field work: reliability, showing up, safety, coordination, and finishing under constraints.
- Bridge to Sherpa 6: mission-critical users require clarity, discipline, and responsiveness.

## Likely Behavioral Questions

Practice answers to these out loud:

1. Tell us about yourself and how your background fits a Go developer role.
2. Why Sherpa 6?
3. What interests you about mission-critical or DoD-adjacent software?
4. Describe the Go parts of your senior capstone.
5. Tell us about a time you had to learn a new technology quickly.
6. Tell us about a technical decision you made and the tradeoffs.
7. Tell us about a bug you diagnosed and fixed.
8. How do you approach code quality and testing?
9. Describe a time you worked with unclear requirements.
10. Describe a time you disagreed with a teammate or stakeholder.
11. How do you communicate progress or blockers?
12. Are you comfortable with hybrid work in Natick, MA, travel up to 25%, and clearance/background requirements?
13. What kind of team environment helps you do your best work?
14. What are your biggest gaps for this role, and how would you close them?
15. What questions do you have for us?

Best positioning:

- Mission fit: practical, reliable software for users with real operational stakes.
- Technical fit: Go, REST APIs, data systems, tests, CI, documentation, fast learning.
- Gap handling: "I have less professional Go tenure than the posting describes, but I can demonstrate core Go fluency today and I have already used it in a shipped academic/team project."

## Likely Technical Questions

### Go Fundamentals

- What is the difference between a slice and an array?
- How do maps behave when a key is missing?
- How do you handle errors idiomatically in Go?
- When would you use an interface?
- What does `defer` do, and what are common pitfalls?
- How do pointers work in Go?
- What are zero values and why do they matter?
- What is the difference between `nil` slice and empty slice?
- How do Go modules work?
- How do you write table-driven tests?
- How do you use `context.Context` in API/database calls?
- What is the race detector and when would you use it?

### Concurrency

- Goroutines vs OS threads.
- Channels vs mutexes: when to use each.
- Buffered vs unbuffered channels.
- Worker pool pattern.
- Fan-out/fan-in pattern.
- Cancellation with context.
- Avoiding goroutine leaks.
- Coordinating shutdown.
- Protecting shared maps with `sync.Mutex` or `sync.RWMutex`.
- Handling the first error from concurrent work.

### REST APIs And Microservices

- Design a REST endpoint for creating/updating a resource.
- Explain status codes: 200, 201, 204, 400, 401, 403, 404, 409, 422, 500.
- How would you validate JSON input?
- How do you structure handlers, services, and repositories?
- How do you add request IDs, logs, timeouts, and metrics?
- How do you handle pagination and filtering?
- What makes an endpoint idempotent?
- How would you version an API?
- How would you test a handler in Go?

### SQL / Relational Databases

- Design tables for users, devices, readings/events, and audit logs.
- Write a query for latest status per device.
- Explain indexes and when not to add one.
- Explain transactions and rollback.
- Describe isolation levels at a high level.
- Prevent SQL injection in Go.
- Handle database timeouts with context.
- Compare `database/sql`, `sqlx`, and ORM-style approaches.
- Migrations: how do you evolve schema safely?

### Security / DoD-Relevant Concepts

- TLS vs mTLS.
- What a certificate authority does.
- How client certificate authentication works at a high level.
- Secret handling: environment variables, secret stores, no commits.
- Password hashing vs encryption.
- Symmetric vs asymmetric encryption.
- Input validation and least privilege.
- Audit logging and traceability.
- STIG/CMMC conceptually: secure configuration, access control, patching, logging, documentation.

### Docker / Linux / Tooling

- How to run a container with ports and environment variables.
- Difference between image and container.
- Multi-stage Dockerfile for Go.
- How to inspect logs.
- How to pass configuration to a service.
- Basic Linux commands for troubleshooting.
- Git workflow: branch, commit, PR, resolving conflicts.

Glassdoor specifically reported one software candidate being asked a Docker command. Be ready to say:

```bash
docker run --rm -p 8080:8080 --env-file .env my-service:latest
```

And explain each flag.

## Probable Live-Coding Topics

The invite says live coding in your IDE. Glassdoor mentions coding challenges and conceptual technical questions. Expect a practical backend/Go problem more than a hard LeetCode puzzle.

### 1. Concurrent Worker Pool

Prompt: Given a list of jobs, process them with N workers, collect successful results, stop on context cancellation, and return errors cleanly.  
What they test: goroutines, channels, WaitGroup, context, error handling, no goroutine leaks.

Practice goal:

- Define `Job` and `Result`.
- Start N workers.
- Close result channel when workers finish.
- Respect `ctx.Done()`.
- Add a table-driven test for success and cancellation.

### 2. REST Handler With Validation

Prompt: Build a `POST /devices` or `POST /events` endpoint that accepts JSON, validates required fields, stores it in memory, and returns JSON.  
What they test: `net/http`, JSON decoding, status codes, small design choices, tests.

Practice goal:

- Use structs with JSON tags.
- Reject malformed JSON and missing fields.
- Separate handler from storage.
- Write `httptest` coverage.

### 3. Latest Status Aggregation

Prompt: Given device event records, return the latest status per device.  
What they test: maps, sorting/time comparison, edge cases, clean function signatures.

Practice goal:

- Parse timestamps or compare `time.Time`.
- Handle empty input.
- Handle duplicate timestamps deterministically.

### 4. Rate Limiter

Prompt: Implement a simple per-client request limiter.  
What they test: maps, mutexes, time windows, API/server thinking.

Practice goal:

- Use `map[string][]time.Time` or token bucket.
- Protect state with a mutex.
- Write focused tests using injected clock/time values.

### 5. SQL Query Design

Prompt: Given `devices(id, unit_id)` and `events(device_id, status, created_at)`, return each device's most recent event.  
What they test: SQL fundamentals and ability to reason about relational data.

Practice goal:

- Use a window function or join to max timestamp.
- Explain index on `(device_id, created_at DESC)`.

### 6. Parse And Count

Prompt: Parse log lines or JSON records, count errors by type, and return top N.  
What they test: string parsing, maps, sorting, edge cases.

Practice goal:

- Keep parsing and aggregation separate.
- Return deterministic ordering for ties.

## Go Coding Checklist During The Interview

Before coding:

- Restate the problem in one sentence.
- Ask about input size, error behavior, concurrency requirements, and ordering.
- Name edge cases before typing.

While coding:

- Start with simple types and a pure function.
- Keep handler/business/storage concerns separate if API-related.
- Return errors instead of panicking.
- Use `context.Context` for cancellation/timeouts when external calls or concurrency are involved.
- Add at least one quick test or explain test cases if time is short.

Before finishing:

- Walk through sample input.
- Mention time and space complexity.
- Point out production hardening: logs, metrics, auth, validation, race tests, database transaction.

## 60-Minute Mock Interview Drill

Run this once Saturday and once Sunday. Record yourself or have a friend ask the questions.

### 0:00-0:05 - Setup And Opening

- Open your Go IDE with a blank module.
- Say your 90-second pitch.
- Confirm the role focus: backend Go, REST APIs, relational databases, mission-critical systems.

Score yourself:

- Clear and concise: 1 point.
- Connects resume to role: 1 point.
- Mentions Go without overstating experience: 1 point.

### 0:05-0:15 - Behavioral Screen

Answer these in STAR format:

1. Why Sherpa 6?
2. Tell me about the Go part of your capstone.
3. Tell me about a time you had unclear requirements.
4. What is a gap you have for this role, and how are you addressing it?

Target answer length: 60-90 seconds each.

### 0:15-0:25 - Technical Rapid Fire

Answer in 45-75 seconds each:

1. Explain goroutines, channels, and when you would use a mutex.
2. How would you structure a Go REST API?
3. How do you use context with HTTP handlers and database calls?
4. What indexes would you add for latest event per device?
5. What is mTLS and why might a defense/customer environment use it?
6. What does this Docker command do?

```bash
docker run --rm -p 8080:8080 --env-file .env my-service:latest
```

### 0:25-0:50 - Live Coding

Use this prompt:

Build a Go function that takes a slice of device events and returns the latest event for each device.

```go
type Event struct {
    DeviceID  string
    Status    string
    CreatedAt time.Time
}

func LatestByDevice(events []Event) map[string]Event {
    // implement
}
```

Requirements:

- Empty input returns an empty map.
- If a device has multiple events, keep the newest.
- If timestamps tie, keep the event with lexicographically larger status so output is deterministic.
- Add table-driven tests.

Stretch after 10-15 minutes:

- Return a sorted slice instead of a map.
- Add JSON parsing from `[]byte`.
- Discuss SQL schema and query equivalent.

What good looks like:

- Correct map logic.
- Clear tie handling.
- Tests for empty input, one device, many devices, tie case.
- Explains O(n) time and O(d) space where d is device count.

### 0:50-0:57 - Mission/Team Fit

Answer:

1. How would you handle a cold or quiet panel?
2. How do you work with users who have operational constraints?
3. How would you ramp into a codebase with security/compliance requirements?

Use Sherpa-specific language: mission-critical customers, Warfighter support, user-focused design, reliability, documentation, and maintainability.

### 0:57-1:00 - Questions For Them

Ask 2-3:

1. What Go services would this role own first?
2. What databases and API frameworks are currently in use?
3. How does the team test and review code before it reaches mission users?
4. What does success look like in the first 90 days?
5. How much of the work is greenfield development versus sustainment?
6. Where do STIG, CMMC, certificates, or mTLS show up in the day-to-day work?

## Weekend Study Plan

### Saturday - Go And Live Coding

- 45 minutes: implement `LatestByDevice` with tests.
- 60 minutes: implement worker pool with context cancellation.
- 45 minutes: build a small `net/http` JSON handler with `httptest`.
- 30 minutes: review slices, maps, interfaces, errors, defer, context, goroutines, channels, mutexes.

### Sunday - Role-Specific Polish

- 45 minutes: SQL latest-event query and indexes.
- 30 minutes: Docker basics and a minimal Go Dockerfile.
- 30 minutes: mTLS/certificate/CA conceptual review.
- 60 minutes: full mock interview drill.
- 20 minutes: prepare your IDE, notes, and questions.

## Technical Cheat Sheet

### Go REST Skeleton

```go
func createDevice(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }

        var req CreateDeviceRequest
        dec := json.NewDecoder(r.Body)
        dec.DisallowUnknownFields()
        if err := dec.Decode(&req); err != nil {
            http.Error(w, "invalid json", http.StatusBadRequest)
            return
        }
        if req.ID == "" || req.Name == "" {
            http.Error(w, "missing required fields", http.StatusUnprocessableEntity)
            return
        }

        device, err := store.Create(r.Context(), req)
        if err != nil {
            http.Error(w, "create failed", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        _ = json.NewEncoder(w).Encode(device)
    }
}
```

### Worker Pool Talking Points

- Jobs channel feeds workers.
- Results channel is closed after `WaitGroup` finishes.
- Context cancellation is checked before sending/receiving.
- Shared state is avoided or protected.
- Errors are returned through a channel or cancel the context.

### SQL Latest Event Query

```sql
SELECT device_id, status, created_at
FROM (
    SELECT
        device_id,
        status,
        created_at,
        ROW_NUMBER() OVER (
            PARTITION BY device_id
            ORDER BY created_at DESC
        ) AS rn
    FROM events
) ranked
WHERE rn = 1;
```

Suggested index:

```sql
CREATE INDEX idx_events_device_created_at
ON events (device_id, created_at DESC);
```

## Glassdoor-Derived Cautions

- Some candidates reported disorganized scheduling, late interviewers, or panels that felt cold. Do not mirror that energy. Stay concise, friendly, and structured.
- Several reports mention assessments that did not perfectly match the role. If a prompt feels odd, ask clarifying questions and solve the stated problem cleanly.
- Reports mention resume/background questions and technology checklist questions. Keep your resume walkthrough tight and be ready to map each resume item to Go, REST, SQL, tests, or teamwork.
- Travel and clearance came up in prior software reports. Have a direct answer ready.

## Final Prep Checklist

- [ ] Go IDE opens quickly with working `go test ./...`.
- [ ] Know which folder you will use for the live challenge.
- [ ] Practice the opening pitch twice.
- [ ] Practice `LatestByDevice` and worker pool from memory.
- [ ] Review `context`, goroutines, channels, mutexes, table-driven tests.
- [ ] Review REST status codes, validation, idempotency, pagination.
- [ ] Review SQL indexes, transactions, latest-row queries.
- [ ] Review Docker run command and minimal Go Dockerfile.
- [ ] Prepare 3 questions for the panel.
- [ ] Join Teams 5 minutes early from the calendar invite.
