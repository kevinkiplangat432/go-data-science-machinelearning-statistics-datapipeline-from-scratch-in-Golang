# go-data-science-from-scratch

I'm building this repo to understand data science and machine learning at the
level where I could implement any algorithm from a whiteboard definition. No
sklearn. No gonum shortcuts. Every function — mean, sigmoid, gradient descent,
Gini impurity, k-means — written by hand, with comments explaining the math at
each step.

This is not a course. It's a construction site.

---

## Why Go and Python together

Most people learn ML in Python and stay there. I'm doing both deliberately.

**Go** is the engineering brain. Strict types, explicit interfaces, no magic.
When I implement a gradient descent loop in Go, I can't hide behind a library
call — every matrix operation, every update step has to be reasoned about and
written out. It forces architectural discipline and makes me think about how
these systems would actually run in production.

**Python** is the laboratory. I mirror every Go implementation in Python, add
matplotlib visualisations, and compare outputs. If the two implementations
disagree on a result, I don't move on until I know exactly why.

Go first. Python second. Agreement required.

---

## What this is actually building toward — ARVIS

Everything in this repo feeds a single objective: **ARVIS**, an AI-powered
telecom network monitoring SaaS targeting Kenyan enterprises, with Safaricom
as the primary target.

ARVIS needs:
- Anomaly detection on network telemetry streams
- Risk scoring on AI agent outputs
- PII detection compliant with Kenya's Data Protection Act
- Agent monitoring and hallucination detection for LLM-powered features
- A data engineering layer that can ingest and process telemetry at scale

Every track in this repo maps directly to one of those requirements. This is
not academic learning for its own sake — each algorithm I implement here is a
building block for a system I intend to ship.

---

## The 5 tracks

**Track 1 — Core Statistics and Probability**

The foundation everything else sits on. Mean, variance, covariance,
correlation, Bayes' theorem, conditional probability, A/B testing with proper
z-scores and confidence intervals. If I can't implement these from scratch I
have no business touching anything above them.

Modules: `statistics`, `probability`, `ab_testing`

---

**Track 2 — ML From Scratch**

The prediction and detection engine for ARVIS. Linear and logistic regression,
decision trees, k-nearest neighbours, Naive Bayes, k-means clustering — each
one implemented from first principles, with gradient descent written out step
by step.

Modules: `linear_regression`, `logistic_regression`, `decision_trees`, `knn`,
`naive_bayes`, `kmeans`

---

**Track 3 — Data Engineering**

Telemetry pipelines and data ingestion. A mini DataFrame, CSV reader, feature
scaler, data cleaner, and a stream processor for aggregations. The plumbing
ARVIS needs before any ML can run.

Modules: `data_engineering`, `analytics`

---

**Track 4 — AI Governance**

The compliance and safety layer ARVIS sells to enterprises. PII detection with
Kenyan formats (Safaricom phone numbers, National IDs, KRA PINs), a
configurable risk scoring engine, agent event monitoring with anomaly
detection, and a hallucination detection module with confidence scoring and
fact consistency checking.

Modules: `pii_detection`, `risk_scoring`, `agent_monitoring`,
`hallucination_detection`

---

**Track 5 — Production Skills**

Deploying all of the above as real services. This track lives outside this
repo for now — it covers containerisation, CI/CD, and serving the ARVIS API —
but every module here is written with production deployment in mind.

---

## Architecture

`main.go` is the central orchestrator. Every module implements a single
interface:

```go
type Runner interface {
    Name() string
    Run() error
}
```

Modules are registered in a `Registry` struct. Nothing in `main()` calls
module logic directly — everything goes through the registry.

```go
// Run everything
go run main.go

// Run one module by name
go run main.go linear_regression
```

Python mirrors the same pattern:

```bash
python python/main.py
python python/main.py linear_regression
```

Each module lives in its own package directory with a single `module.go` (Go)
and `module.py` (Python). The Python files add matplotlib visualisations where
relevant — loss curves, decision boundaries, centroid convergence plots.

---

## Where this points

The skills I'm building here point toward ML engineering at the level of
Google DeepMind — not as a fantasy, as a target I'm engineering toward,
deliberately, one algorithm at a time.

ARVIS is the startup vehicle. This repo is where the foundation gets built.
