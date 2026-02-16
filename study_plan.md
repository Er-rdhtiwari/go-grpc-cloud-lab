# 1-Month Beginner-Friendly Study Plan — IBM Cloud Data Services (Go + Kubernetes + CI/CD + GitOps)

This role is basically: **build Go microservices**, **run them on Kubernetes**, **automate deploys with CI/CD + GitOps**, and **operate them in production** (alerts, on-call, security, reliability).

## Time options (pick one)
- **Minimum (1 hour/day):** read notes + do the mini-lab steps (skip the extras).
- **Ideal (2–3 hours/day):** notes + mini-lab + add 1 small capstone improvement daily.

## Daily DSA track (Beginner → Medium, 30–45 min/day)
Many interviews for this kind of role still ask **DSA (Data Structures & Algorithms)**.
- Do **one DSA topic per day** + **3–5 practice questions** (mix easy/medium).
- Focus on patterns you’ll reuse: **two pointers, sliding window, BFS/DFS, heaps, DP**, etc.
- Track for every problem: **approach**, **time complexity**, **space complexity**, and **common mistakes**.

## How to use this plan (simple)
- Each day has a **ready-to-copy prompt**. Paste it into ChatGPT (or your preferred LLM) to get **beginner-friendly notes + a mini-lab**.
- Save outputs as `notes/day-XX.md` (optional).
- Don’t try to master everything at once. You’ll revisit topics every week through the capstone.

## Before Day 1 (30–60 minutes)
- Make sure you can run: `go version`, `git --version`, `docker version`, `kubectl version --client`, `helm version`
- If you’re missing tools, install them first (Go, Docker, kubectl, Helm). Add `kind` or `minikube` later if needed.

## Key words (plain English)
- **Microservice:** a small service that does one job (ex: Orders).
- **gRPC + protobuf:** a fast way for services to talk; protobuf is the “message format”.
- **Kubernetes (K8s):** runs your containers, restarts them, scales them, and connects them.
- **Helm:** “package manager” for Kubernetes YAML (templating + releases).
- **CI/CD:** automated build/test/release pipelines.
- **GitOps:** Kubernetes deploys are driven by Git as the source of truth.
- **Observability:** logs + metrics + traces so you can debug production issues.
- **On-call:** you respond to alerts and fix incidents.

## Fundamentals checklist (explicit topics + subtopics)
Use this to confirm the plan covers the “basic + must-know” foundations for the role.

- **Go fundamentals**
  - Go modules + dependency basics (Day 1)
  - Repo layout (`cmd/`, `internal/`) (Day 1)
  - Concurrency + `context` + timeouts (Day 2)
  - gRPC/protobuf + backward compatibility rules (Day 3)
  - Errors + logging + graceful shutdown (Day 5)
  - Testing (unit + integration) (Day 4)
  - Reliability patterns (timeouts/retries/idempotency) (Day 6)
- **Git fundamentals (commonly expected)**
  - Branching + PR flow, rebase vs merge (Day 1)
  - Revert/rollback mindset (Day 20)
  - CI + protected branches + required checks (Day 18)
- **Linux + networking fundamentals (important for on-call)**
  - Processes + signals (SIGTERM/SIGKILL) (Day 5 + Day 25)
  - Ports/listeners (`ss`/`netstat`), basic connectivity (`curl`) (Day 25)
  - DNS/service discovery (`nslookup`/`dig`) (Day 11 + Day 25)
  - TLS/HTTPS basics (certificates, termination) (Day 11)
- **Docker fundamentals**
  - Multi-stage build, minimal images, non-root user (Day 8)
- **Kubernetes fundamentals**
  - Core objects: Pod/Deployment/Service/ConfigMap/Secret (Day 9)
  - Workload types: Job/CronJob/DaemonSet/StatefulSet (Day 9 + Day 13)
  - Probes + resources + autoscaling + rollouts (Day 10)
  - Networking: Service types/Ingress/DNS/NetworkPolicy (Day 11)
  - Security: RBAC + ServiceAccounts + pod security basics (Day 12)
- **Helm fundamentals**
  - Chart structure + values + templating + upgrades (Day 15–16)
- **GitOps fundamentals**
  - Desired state + reconciliation + rollback (Day 17)
- **CI/CD fundamentals**
  - Build/test/scan/publish pipeline stages (Day 18)
- **Security/compliance fundamentals**
  - Scanning + SBOM + secrets handling + secure defaults (Day 19)
- **Observability fundamentals**
  - Logs/metrics/traces basics (Days 22–24)
  - Alerting + SLO basics (Day 23)
- **Data fundamentals**
  - Postgres basics + migrations + performance pitfalls (Day 26)
  - Redis caching basics + invalidation pitfalls (Day 27)
  - Queues (RabbitMQ/Kafka), retries + DLQ basics (Day 27)
- **Cloud + IaC fundamentals**
  - VPC + IAM + storage (Day 28)
  - Terraform basics + state + drift (Day 29)
- **DSA fundamentals (Beginner → Medium)**
  - Daily DSA track + weekly mixed practice (All days; see each day’s DSA line)

## Choose one option in each pair (to stay focused)
- **CI:** GitHub Actions (simpler) OR Jenkins (common in enterprises)
- **GitOps:** Argo CD (popular) OR Flux
- **Queue:** RabbitMQ (task queue style) OR Kafka (event stream style)
- **Cloud:** AWS OR Azure OR GCP (pick one and stick to it)
- **IaC:** Terraform (recommended)

## Capstone (you build this slowly across the month)
Build a small **Go “orders” service** (gRPC) that:
- Stores orders in **Postgres**
- Uses **Redis** for caching (basic)
- Publishes an **OrderCreated** message to a queue (RabbitMQ or Kafka)
- Runs in **Kubernetes**, packaged with **Helm**, deployed with **GitOps**
- Has basic **CI** (lint + test + build + scan)
- Has **logs + metrics + traces**, plus 2–5 useful alerts and runbooks

---

## Week 1 — Production Go for Microservices (Days 1–7)

### Day 1 — Go project structure + tooling fundamentals
**Goals**
- Set up a “production-ready” Go workspace: modules, linting, formatting, dependency hygiene
- Learn conventions: `cmd/`, `internal/`, `pkg/`, config, logging
**DSA (30–45 min):** Big-O basics + Arrays

**Ready-to-copy prompt**
```
You are a patient Golang mentor. Assume I am a beginner and explain everything in simple words first, then go deeper. Define any new terms.

Day 1 topic: Production Go project structure + tooling.

Create descriptive notes that include:
1) A clear explanation of Go modules, versioning, and dependency management (replace/replace pitfalls).
2) Recommended repo layout for microservices (cmd/internal/pkg), with pros/cons and when to break into multiple modules.
3) A practical “baseline toolchain” list: gofmt, go test, go vet, staticcheck, golangci-lint, govulncheck, buf (if using protobuf), and why each matters.
4) A checklist for a new microservice repo (Makefile targets, CI tasks, build tags, config, logging).
5) A short exercise: propose a directory layout for a gRPC service called “orders” and show example file names.
6) Git basics for daily work: branches, PR workflow, merge vs rebase (high-level), revert, and good commit messages.

Keep it structured with headings, bullet points, and short code snippets where useful.
Include: Key terms (with 1-line definitions), 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with 5 interview questions and strong sample answers.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Big-O basics + Arrays.
- Explain with simple examples: time vs space, common array operations, when arrays are good/bad.
- Give 5 practice questions (2 easy, 3 medium) with: approach, time/space complexity, and Go-style pseudocode.
```

### Day 2 — Concurrency, context, cancellation, timeouts
**DSA (30–45 min):** Strings + HashMap/Set basics
**Ready-to-copy prompt**
```
You are a patient Golang mentor. Assume I am a beginner. Start with a simple mental model, then go deeper. Define any new terms.

Day 2 topic: Go concurrency + context.

Write descriptive notes covering:
1) Goroutines, channels, select, worker pools, backpressure.
2) context.Context: propagation, cancellation, deadlines, timeouts (especially for DB + RPC calls).
3) Common pitfalls: goroutine leaks, unbounded concurrency, shared memory races.
4) Patterns: errgroup, bounded parallelism, retries with jitter, circuit breaker concept (no need for a library).
5) Debugging: race detector, pprof basics for goroutines.

Mini-lab:
- Provide a small Go example that runs N concurrent tasks with a timeout, cancels on first error, and avoids leaks.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a “production checklist” for concurrency in services.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Strings + HashMap/Set basics.
- Teach: frequency counting, anagrams, duplicates, first unique, and common pitfalls.
- Give 5 practice questions (2 easy, 3 medium) with hints + final solutions in Go.
```

### Day 3 — API design: gRPC + protobuf + error semantics
**DSA (30–45 min):** Two pointers (array/string)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain with simple examples and define any new terms.

Day 3 topic: gRPC + protobuf for microservices.

Create descriptive notes including:
1) When to use gRPC vs REST in internal platforms.
2) Protobuf fundamentals: messages, enums, field numbers, optional fields, backward compatibility rules.
3) API versioning strategies (package names, service names, message evolution).
4) Error handling: gRPC status codes, mapping domain errors, returning details, retries vs non-retries.
5) Security basics: mTLS concept, authn/authz patterns at gateway and service levels.

Mini-lab:
- Design an `OrderService` proto (Create/Get/List) with pagination and idempotency key.
- Show examples of request/response messages and how you would handle errors.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with 10 proto design “dos and don’ts”.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Two pointers.
- Teach patterns: left/right, slow/fast, removing duplicates in-place, palindrome checks.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 4 — Testing in Go: unit, integration, mocks, table-driven tests
**DSA (30–45 min):** Sorting + Binary search basics
**Ready-to-copy prompt**
```
You are a patient Go testing mentor. Assume I am a beginner. Define terms like unit test, integration test, mock, fake, and table-driven test.

Day 4 topic: Testing in Go for production microservices.

Write descriptive notes that cover:
1) Table-driven tests, subtests, golden tests, and coverage expectations.
2) Mocking strategies: interfaces, fakes, testify/mock vs hand-rolled fakes, and tradeoffs.
3) Testing concurrency: avoiding flakes, using timeouts, deterministic patterns.
4) Integration tests with Postgres/Redis (describe approaches with Docker Compose or testcontainers).
5) CI test strategy: fast unit tests vs slower integration tests; caching; parallelization.

Mini-lab:
- Provide a sample service function and show a good test suite structure (unit + integration outline).

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with “How to write bug-preventing tests” checklist.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Sorting + Binary search basics.
- Teach: when to sort, stable vs unstable (high-level), binary search template, off-by-one mistakes.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 5 — Configuration, logging, and error handling conventions
**DSA (30–45 min):** Stack + Queue (and common uses)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain configuration and logging like I’ve never deployed to Kubernetes before.

Day 5 topic: Config + logging + error handling in Go services.

Create descriptive notes including:
1) Config sources: env vars, config files, Kubernetes ConfigMaps/Secrets.
2) 12-factor configuration principles and anti-patterns.
3) Structured logging: fields, correlation IDs, request IDs, log levels, sampling.
4) Error handling: wrapping, sentinel errors vs typed errors, returning vs logging, user-facing vs internal errors.
5) Graceful shutdown: signal handling, draining, timeouts, in-flight requests.

Mini-lab:
- Show a small Go “main” skeleton for a service that loads config, starts a server, handles SIGTERM, and logs structured events.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with an interview-style explanation of graceful shutdown in Kubernetes.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Stack + Queue.
- Teach: valid parentheses, next greater element (intro), BFS queue idea (preview).
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 6 — Service reliability patterns: retries, idempotency, rate limiting
**DSA (30–45 min):** Linked List basics (slow/fast pointers)
**Ready-to-copy prompt**
```
You are a patient distributed-systems mentor. Assume I am a beginner. Use simple examples (like “retrying a payment” or “placing an order twice”).

Day 6 topic: Reliability patterns for distributed systems.

Write descriptive notes covering:
1) Timeouts, retries (exponential backoff + jitter), and when NOT to retry.
2) Idempotency keys and deduplication (especially for Create operations).
3) Rate limiting (token bucket/leaky bucket), load shedding, bulkheads.
4) Circuit breakers and dependency health.
5) Exactly-once vs at-least-once semantics in message queues.

Mini-lab:
- Propose an idempotent CreateOrder flow (DB schema hints + logic outline).

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with 8 “failure modes” and how to mitigate them.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Linked List basics.
- Teach: traverse, reverse, cycle detection (Floyd), merge two lists.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 7 — Weekly review + small capstone checkpoint
**DSA (30–60 min):** Week 1 mixed practice (easy → medium)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Keep the review very practical.

Day 7 topic: Weekly review + checkpoint.

1) Summarize the key takeaways from Days 1–6 in a concise study sheet.
2) Create a short self-assessment quiz (15 questions) with answers.
3) Provide a plan to implement the Week 1 checkpoint in a repo:
   - gRPC proto drafted
   - Go module initialized
   - Basic service skeleton + graceful shutdown
   - Unit test example

Keep it practical and focused on production readiness.

DSA Track:
- Create a review sheet for Week 1 DSA topics (arrays, strings/hashmap, two pointers, binary search, stack/queue, linked list).
- Give 8 practice questions total (4 easy, 4 medium) with brief solutions and complexity.
```

---

## Week 2 — Containers + Kubernetes Core (Days 8–14)

### Day 8 — Docker fundamentals for Go services
**DSA (30–45 min):** Sliding window
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain Docker concepts with a simple “build an app and run it” story.

Day 8 topic: Docker for Go microservices.

Create descriptive notes including:
1) Multi-stage builds, static vs dynamic linking, distroless images, non-root user.
2) Image size, caching layers, build args, reproducible builds.
3) Runtime security: minimal base, read-only filesystem concept, dropping capabilities.
4) Common Dockerfile mistakes for Go (CGO, timezone/certs, caching go mod).

Mini-lab:
- Provide a production-grade Dockerfile for a Go gRPC service and explain each layer.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a checklist to review any Dockerfile in PRs.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Sliding window.
- Teach: fixed vs variable window, “expand/contract” pattern, common mistakes.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 9 — Kubernetes architecture + core objects
**DSA (30–45 min):** Prefix sums + frequency (intro)
**Ready-to-copy prompt**
```
You are a patient Kubernetes mentor. Assume I am a beginner. Use analogies and step-by-step flows.

Day 9 topic: Kubernetes fundamentals.

Write descriptive notes covering:
1) Control plane components (api-server, etcd, scheduler, controller manager) and worker components.
2) Core objects: Pod, Deployment, ReplicaSet, Service, Namespace, ConfigMap, Secret.
3) Common workload types: DaemonSet, Job, CronJob (what they are used for).
4) Labels/selectors, annotations, and common conventions.
5) Rolling updates and safe rollout strategies.

Mini-lab:
- Provide YAML examples for a Deployment + Service for a Go app with sensible defaults.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with “how a request reaches a Pod” explanation.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Prefix sums (and “running sum” thinking).
- Teach: subarray sum queries, range sums, and how prefix sums reduce repeated work.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 10 — Health checks, probes, resources, and autoscaling
**DSA (30–45 min):** Recursion basics
**Ready-to-copy prompt**
```
You are a patient Kubernetes mentor. Assume I am a beginner. Explain probes and resources with simple examples (what happens when CPU/memory is too low).

Day 10 topic: Kubernetes runtime reliability.

Create descriptive notes including:
1) Liveness vs readiness vs startup probes (when to use each).
2) Resource requests/limits, CPU throttling, memory OOM behavior.
3) HPA basics (CPU/memory/custom metrics overview).
4) PodDisruptionBudgets, graceful termination, preStop hooks.

Mini-lab:
- Provide a Deployment spec snippet with probes, resources, and termination settings tuned for a Go service.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a “common outage causes in K8s” list.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Recursion basics.
- Teach: base case, recursion tree idea, stack overflow risk, converting to iteration (conceptual).
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 11 — Networking: Ingress, DNS, service discovery, NetworkPolicies (basics)
**DSA (30–45 min):** Tree basics (traversals)
**Ready-to-copy prompt**
```
You are a patient Kubernetes mentor. Assume I am a beginner. Explain networking from “my laptop → service” step-by-step.

Day 11 topic: Kubernetes networking for microservices.

Write descriptive notes covering:
1) ClusterIP vs NodePort vs LoadBalancer; when each is used.
2) Ingress concepts (controller, rules, TLS termination).
3) DNS/service discovery and common debugging commands (kubectl + nslookup).
4) NetworkPolicy basics: default deny, allow-by-namespace/label patterns.
5) Networking fundamentals (quick but clear): TCP vs HTTP (high-level), TLS/HTTPS basics, what a certificate is, and where TLS is terminated.

Mini-lab:
- Draft a NetworkPolicy that only allows traffic to `orders` from `api-gateway` namespace.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a debugging playbook for “service not reachable”.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Trees basics (DFS traversals).
- Teach: pre/in/post-order, recursion vs iterative stack, common patterns.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 12 — RBAC, ServiceAccounts, Secrets, and least privilege
**DSA (30–45 min):** BFS + DFS on trees/graphs (intro)
**Ready-to-copy prompt**
```
You are a patient Kubernetes security mentor. Assume I am a beginner. Define RBAC, Role, Binding, and least privilege very clearly.

Day 12 topic: Kubernetes security essentials.

Create descriptive notes including:
1) RBAC concepts: Role/ClusterRole, RoleBinding/ClusterRoleBinding.
2) ServiceAccounts and how Pods use them.
3) Secrets vs ConfigMaps; common secret handling mistakes.
4) Pod security basics: non-root, readOnlyRootFilesystem, seccomp, capabilities.

Mini-lab:
- Provide example RBAC YAML for a service that only needs to read ConfigMaps in its namespace.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a K8s security review checklist for PRs.

DSA Track (Beginner → Medium):
- Today’s DSA topic: BFS vs DFS (where to use which).
- Teach: queue vs stack, visited set, cycle prevention, complexity.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 13 — Stateful workloads overview: Postgres/Redis in K8s (conceptual)
**DSA (30–45 min):** Hashing patterns (subarray sum = k, duplicates, grouping)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain “stateful” vs “stateless” with simple examples.

Day 13 topic: Running stateful systems with Kubernetes (overview).

Write descriptive notes covering:
1) StatefulSet vs Deployment; PVCs; storage classes.
2) Backups, restore, upgrades, and failure modes for Postgres/Redis.
3) Why managed services are preferred; when self-managed is acceptable.
4) Connection pooling and service-level best practices (timeouts, retries, migrations).

Mini-lab:
- Outline a safe approach to run Postgres for dev/test in K8s and what NOT to do in prod.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with “things SREs care about” for stateful components.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Hashing patterns.
- Teach: using map for counts/indices, set for membership, avoiding collisions conceptually.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 14 — Weekly review + K8s checkpoint
**DSA (30–60 min):** Week 2 mixed practice (trees + BFS/DFS + sliding window)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Keep the review practical and command-focused.

Day 14 topic: Weekly review + Kubernetes checkpoint.

1) Summarize Days 8–13 into a cheat sheet (YAML + commands).
2) Create a kubectl troubleshooting flowchart (text-based) for common issues:
   - CrashLoopBackOff
   - ImagePullBackOff
   - Readiness failing
   - Service endpoints empty
3) Propose a Week 2 checkpoint for the capstone:
   - Container image builds
   - K8s manifests for service
   - Probes + resources configured

End with 10 interview questions and answers on Kubernetes basics.

DSA Track:
- Summarize Week 2 DSA patterns (sliding window, prefix sums, recursion, trees traversals, BFS/DFS, hashing patterns).
- Give 8 practice questions total (4 easy, 4 medium) with brief solutions + complexity.
```

---

## Week 3 — Helm + GitOps + CI/CD (Days 15–21)

### Day 15 — Helm fundamentals: charts, templates, values, releases
**DSA (30–45 min):** Heap / Priority Queue basics
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain Helm like “Kubernetes YAML templates + versioned releases”.

Day 15 topic: Helm for Kubernetes services.

Create descriptive notes including:
1) Helm chart structure (Chart.yaml, values.yaml, templates/), releases, revisions.
2) Templating basics: values, helpers, functions, conditionals, loops.
3) Best practices: naming, labels, image tags, resources, probes, config separation.
4) Common pitfalls: whitespace, quoting, YAML indentation issues, breaking upgrades.

Mini-lab:
- Provide a minimal but production-minded Helm chart outline for the `orders` service (files + key templates).

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a Helm PR review checklist.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Heap / Priority Queue.
- Teach: min-heap vs max-heap, top-K problems, complexity.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 16 — Advanced Helm: dependency charts, hooks, testing, chart linting
**DSA (30–45 min):** Intervals (merge/insert) patterns
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain “dependencies” and “hooks” with simple examples and warnings.

Day 16 topic: Helm advanced usage.

Write descriptive notes covering:
1) Chart dependencies and when to vendor vs reference.
2) Helm hooks (pre/post install/upgrade) and why they are risky.
3) Testing approaches: helm lint, template rendering tests, chart unit tests (conceptual).
4) Upgrade safety: immutable fields, migrations, rollback strategy.

Mini-lab:
- Show how you’d structure values for dev/stage/prod and avoid config drift.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with “safe upgrade” principles for Helm releases.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Intervals.
- Teach: sorting by start/end, merge intervals, insert interval, edge cases.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 17 — GitOps concepts: Argo CD / Flux, desired state, rollbacks
**DSA (30–45 min):** Greedy basics
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain GitOps with a simple “Git is the deploy source of truth” story.

Day 17 topic: GitOps for Kubernetes.

Create descriptive notes including:
1) GitOps principles: desired state, reconciliation loop, auditability.
2) Argo CD vs Flux: conceptual comparison and typical workflows.
3) Deployment strategies with GitOps: progressive delivery overview (blue/green, canary concepts).
4) Secrets management options (high-level): Sealed Secrets, External Secrets (conceptual).

Mini-lab:
- Outline a Git repo structure for GitOps managing multiple environments (dev/stage/prod).

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a “GitOps incident response” checklist.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Greedy algorithms (intro).
- Teach: “locally best” choice, how to prove/validate, when greedy fails.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 18 — CI/CD pipelines: GitHub Actions + Jenkins patterns
**DSA (30–45 min):** Backtracking basics
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain CI/CD step-by-step from “push code” → “tests” → “build image” → “deploy”.

Day 18 topic: CI/CD for Go microservices (GitHub Actions + Jenkins patterns).

Write descriptive notes covering:
1) Typical pipeline stages: lint, test, build, scan, package, publish, deploy (GitOps).
2) Caching go modules, parallel jobs, artifacts, and how to keep CI fast.
3) Versioning: semantic versioning, build metadata, tagging strategy.
4) “Bots” in repos: dependabot/renovate concept, auto-PR checks, policy gates.
5) Git workflow refresher in CI context: PR checks, protected branches, required status checks, and how to revert safely.

Mini-lab:
- Draft a CI pipeline outline for the capstone that builds an image and updates a GitOps repo (conceptual).

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a CI security checklist.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Backtracking.
- Teach: decision tree, choose/explore/unchoose, pruning, complexity reality.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 19 — Security & compliance for images and services
**DSA (30–45 min):** Dynamic Programming (DP) basics — 1D
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain security and compliance without jargon and with concrete examples.

Day 19 topic: Supply chain security + compliance basics.

Create descriptive notes including:
1) Vulnerability scanning (SCA + image scanning) and what to do with findings.
2) SBOM concept (CycloneDX/SPDX), why organizations require it.
3) Image signing/verification concept (cosign-style flow), policy enforcement concept.
4) Secrets handling: avoiding leaking secrets in logs/builds; rotation.
5) Secure-by-default runtime settings for K8s workloads.

Mini-lab:
- Provide a “secure baseline” checklist for a Go service container + its K8s deployment.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with 10 compliance-focused interview Q&A.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Dynamic Programming (DP) basics — 1D.
- Teach: overlapping subproblems, memoization vs tabulation, building dp arrays.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 20 — Release engineering + rollout/rollback practice
**DSA (30–45 min):** Dynamic Programming — 2D (grid/knapsack intro)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain rollout vs rollback with examples and what to check first.

Day 20 topic: Release engineering in microservices.

Write descriptive notes covering:
1) Deployment safety: feature flags, phased rollouts, metrics-based promotion (conceptual).
2) Rollback strategies: Helm rollback vs GitOps revert, DB migration compatibility.
3) Operational readiness: runbooks, dashboards, SLOs before release.

Mini-lab:
- Define a release checklist and a rollback checklist for the capstone.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a “what can go wrong” section and mitigations.

DSA Track (Beginner → Medium):
- Today’s DSA topic: DP — 2D (grid paths / 0-1 knapsack intro).
- Teach: state definition, transitions, base cases, common mistakes.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 21 — Weekly review + GitOps/CI checkpoint
**DSA (45–60 min):** Week 3 mixed practice (heap/intervals/greedy/backtracking/DP)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Make the checkpoint steps very concrete and ordered.

Day 21 topic: Weekly review + checkpoint (Helm + GitOps + CI/CD).

1) Summarize Days 15–20 into a concise cheat sheet.
2) Provide a step-by-step checkpoint plan:
   - Helm chart exists and renders correctly
   - GitOps repo structure created
   - CI pipeline builds/tests/scans and publishes artifacts
3) Create a 20-question quiz with answers.

End with common interview scenarios and how to respond.

DSA Track:
- Summarize Week 3 DSA topics (heap, intervals, greedy, backtracking, DP 1D/2D).
- Give 8 practice questions total (3 easy, 5 medium) with brief solutions + complexity.
```

---

## Week 4 — Observability + SRE/On-call + Data Systems + Cloud/IaC (Days 22–30)

### Day 22 — Observability foundations: logs, metrics, traces
**DSA (30–45 min):** Graphs basics (adjacency list, traversal)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain observability with simple “how do I debug production?” examples.

Day 22 topic: Observability for Kubernetes microservices.

Create descriptive notes covering:
1) Logs vs metrics vs traces: what each answers and typical tooling stacks (high-level).
2) Structured logging fields (service, version, trace_id, request_id, tenant, latency).
3) Metrics basics: RED/USE/Golden signals; labels vs cardinality pitfalls.
4) Tracing basics: spans, context propagation, sampling, baggage.

Mini-lab:
- Provide an “observability contract” for the capstone service (what must be logged/measured/traced).

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a checklist for adding observability to a new endpoint.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Graphs basics.
- Teach: adjacency list vs matrix, BFS/DFS on graphs, visited, components.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 23 — Prometheus metrics + alerting basics
**DSA (30–45 min):** Topological sort (DAG)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain each metric type with a tiny example.

Day 23 topic: Prometheus-style metrics and alerting.

Write descriptive notes including:
1) Counter vs gauge vs histogram/summary; when to use each.
2) Service-level metrics: request count, error rate, latency buckets, in-flight requests.
3) Alert design: symptom-based vs cause-based alerts, avoiding alert fatigue.
4) SLO/SLI basics and burn-rate concept (high-level).

Mini-lab:
- Propose 5 high-signal alerts for the capstone service and explain why each is actionable.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a “bad alerts” anti-pattern list.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Topological sort (DAG).
- Teach: prerequisites scheduling, Kahn’s algorithm idea, cycle detection.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 24 — OpenTelemetry instrumentation (conceptual) + trace correlation
**DSA (30–45 min):** Union-Find (Disjoint Set Union)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain tracing with a “request journey” story across services.

Day 24 topic: OpenTelemetry tracing for Go services (conceptual + practical tips).

Create descriptive notes covering:
1) How traces propagate across gRPC calls (context propagation).
2) What to instrument: inbound/outbound RPCs, DB queries, queue publish/consume.
3) Correlating logs with traces (trace_id in logs).
4) Performance overhead and sampling strategies.

Mini-lab:
- Provide a minimal example outline showing where tracing hooks would be placed in a Go gRPC server (no need for full runnable code).

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a troubleshooting guide for “missing traces”.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Union-Find (DSU).
- Teach: find/union, path compression, union by rank, where it’s used.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 25 — On-call readiness: incident response, runbooks, postmortems
**DSA (30–45 min):** Shortest path basics (BFS vs Dijkstra concept)
**Ready-to-copy prompt**
```
You are a patient on-call mentor. Assume I am a beginner. Use a calm, step-by-step approach and practical checklists.

Day 25 topic: On-call + incident management for “You build it, you run it”.

Write descriptive notes covering:
1) Incident lifecycle: detect → triage → mitigate → resolve → learn.
2) Severity levels and communication patterns (stakeholders, updates).
3) Runbooks: what good runbooks contain; examples for common service failures.
4) Postmortems: blameless culture, root cause vs contributing factors, action items.
5) Beginner-friendly Linux + networking quick toolkit for incidents:
   - processes/signals (`ps`, `top`, `kill`, why SIGTERM matters)
   - ports (`ss`/`netstat`, “is my service listening?”)
   - DNS (`dig`/`nslookup`)
   - HTTP checks (`curl`), TLS check concept (what failures look like)
   - logs (what to check first in app logs vs Kubernetes events)

Mini-lab:
- Create a runbook template for the capstone service and fill in 2 example incidents:
  a) high error rate due to DB timeouts
  b) increased latency due to CPU throttling

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with an on-call “first 15 minutes” checklist.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Shortest paths basics.
- Teach: unweighted shortest path = BFS, weighted = Dijkstra (high-level), complexity.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 26 — Databases: Postgres fundamentals for service developers
**DSA (30–45 min):** Bit manipulation (common interview tricks)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain Postgres concepts with simple examples and practical advice.

Day 26 topic: Postgres for microservices developers.

Create descriptive notes including:
1) Transactions, isolation levels (practical view), locks, and common pitfalls.
2) Indexing basics, query plans (high-level), and avoiding N+1 queries.
3) Connection pooling, timeouts, retries, and “thundering herd” issues.
4) Schema migrations strategy (backward compatible migrations).

Mini-lab:
- Propose a Postgres schema for Orders with idempotency and explain indices.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a debugging checklist for “DB is slow” incidents.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Bit manipulation.
- Teach: AND/OR/XOR, shifting, checking bits, common patterns (single number).
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 27 — Caching + queues: Redis + RabbitMQ/Kafka fundamentals
**DSA (30–45 min):** Monotonic stack/queue (intro)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain caching and queues with real-world examples and failure cases.

Day 27 topic: Redis + message queues for distributed systems.

Write descriptive notes covering:
1) Redis use cases: cache-aside, write-through, TTLs, cache invalidation, hot keys.
2) Reliability concerns: eviction policies, memory sizing, timeouts.
3) Messaging: RabbitMQ vs Kafka (conceptual), ordering, retries, dead-letter queues, consumer groups.
4) Idempotent consumers and exactly-once illusions.

Mini-lab:
- Design an async “OrderCreated” event flow: producer, schema, consumer behavior, retry policy, DLQ.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with “how to avoid duplications and message loss” checklist.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Monotonic stack/queue (intro).
- Teach: “keep increasing/decreasing”, next greater element, sliding window max idea.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 28 — Cloud basics: VPC, IAM, storage, and least privilege (choose one provider)
**DSA (30–45 min):** Trie basics (prefix search)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain cloud networking and IAM with simple diagrams in words (step-by-step flows).

Day 28 topic: Cloud fundamentals for platform engineers (choose AWS OR Azure OR GCP).

Create descriptive notes including:
1) VPC networking: subnets, routing, NAT, security groups/firewalls (high-level).
2) IAM fundamentals: principals, roles, policies, least privilege, rotation.
3) Storage basics: object storage vs block vs file, encryption at rest/in transit.
4) Typical microservice cloud dependencies and failure modes.

Mini-lab:
- Provide a “secure network + IAM” design checklist for deploying a K8s-based service in the cloud.

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with cloud interview questions focused on VPC and IAM.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Trie (prefix tree).
- Teach: insert/search/startsWith, complexity, memory tradeoffs, when to use.
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 29 — IaC basics: Terraform (preferred) + config management (Ansible concept)
**DSA (30–45 min):** Range queries (prefix sums + difference array) — practical medium
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Explain Terraform with simple “declare desired infrastructure” examples.

Day 29 topic: Infrastructure as Code (Terraform) + dependency management.

Write descriptive notes covering:
1) Terraform basics: providers, resources, modules, state, plan/apply, drift.
2) Safe workflows: remote state, state locking, workspaces vs environments, approvals.
3) Managing service dependencies: networking, IAM, storage, databases.
4) Where Ansible fits vs Terraform (conceptual).

Mini-lab:
- Draft a module-level design (no need for real cloud credentials) for provisioning:
  - network (VPC/subnets)
  - an object store bucket
  - IAM role/policy for a service

Include: Key terms, 3 common beginner mistakes, and a 5-question mini-quiz (with answers).
End with a checklist for reviewing Terraform PRs.

DSA Track (Beginner → Medium):
- Today’s DSA topic: Range queries (prefix sums + difference array idea).
- Teach: when to use prefix sums, how difference arrays help with many range updates (conceptual).
- Give 5 practice questions (2 easy, 3 medium) with Go solutions + complexity.
```

### Day 30 — Final review + capstone readiness + interview prep
**DSA (60–90 min):** Mixed mock interview set (easy → medium)
**Ready-to-copy prompt**
```
You are a patient mentor. Assume I am a beginner. Keep the final review practical and focused on what I should say/do in interviews.

Day 30 topic: Final review + interview preparation.

1) Create a 2-page condensed “IBM Cloud Data Services” study sheet:
   - Go microservices practices
   - Kubernetes essentials
   - Helm + GitOps
   - CI/CD + security/compliance
   - Observability + on-call
   - Postgres/Redis/queues basics
2) Create 25 interview questions with strong answers (mix theory + practical).
3) Provide a capstone completion checklist and a demo script:
   - what to show (metrics, logs, traces, rollout, rollback)
   - what failure drill to simulate and how to explain mitigation

Keep it practical and aligned with “You build it, you run it”.

DSA Track:
- Create a 2-week revision plan (after Day 30) to keep DSA fresh.
- Give a mock interview set of 10 questions (4 easy, 6 medium) across the month’s DSA topics with brief solutions + complexity.
```

---

## Optional: Daily execution checklist (recommended)
- 10 min: review yesterday’s notes
- 60–90 min: deep study + take notes
- 45–90 min: mini-lab / implement capstone increment
- 10 min: write “what I learned + what’s unclear”

## Optional: Tools to install (pick what fits your machine)
- Go latest stable, `protobuf` compiler + `buf`, Docker, `kubectl`, `kind` or `minikube`, Helm, `k9s` (optional)
