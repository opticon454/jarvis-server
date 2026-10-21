# External feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-osx-api

Current disposition: **external host service; consolidation undecided**.

Legacy evidence: [api](https://github.com/alexberardi/jarvis-osx-api/tree/65d980499f8396b6af7d4eaa39a75ba81bf2d8fb/api), [auth](https://github.com/alexberardi/jarvis-osx-api/tree/65d980499f8396b6af7d4eaa39a75ba81bf2d8fb/auth), [services](https://github.com/alexberardi/jarvis-osx-api/tree/65d980499f8396b6af7d4eaa39a75ba81bf2d8fb/services).

### OSX-01

**Local scoped API-key creation and permission checks**

Status: **not started**. Disposition: **external host service; consolidation undecided**.

Prerequisites: [AUTH-05](Get-Foundations-Features.md#auth-05).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Mac service has its own credentials; desktop enrolment is not an equivalent API port.

### OSX-02

**Unread messages, thread history/style context and mark-read ledger**

Status: **not started**. Disposition: **external host service; consolidation undecided**.

Prerequisites: [OSX-01](Get-External-Features.md#osx-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

No Messages chat.db/typedstream integration found in Go agent.

### OSX-03

**Contact resolution and message sending with sanitisation/idempotency**

Status: **not started**. Disposition: **external host service; consolidation undecided**.

Prerequisites: [OSX-02](Get-External-Features.md#osx-02).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Native desktop control does not replace structured Messages/Contacts APIs.

### OSX-04

**Run allowlisted macOS Shortcuts**

Status: **not started**. Disposition: **external host service; consolidation undecided**.

Prerequisites: [OSX-01](Get-External-Features.md#osx-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

No equivalent scoped Shortcuts service found; retain external dependency for messages packages.

## jarvis-relay

Current disposition: **external cloud boundary; consolidation undecided**.

Legacy evidence: [app](https://github.com/alexberardi/jarvis-relay/tree/f106350428c0d8783a3a9d2d5cebff84cd7408e7/app), [README.md](https://github.com/alexberardi/jarvis-relay/blob/f106350428c0d8783a3a9d2d5cebff84cd7408e7/README.md).

### RELAY-01

**Household token registration and authenticated Expo push fan-out**

Status: **in progress**. Disposition: **external cloud boundary; consolidation undecided**.

Prerequisites: [NOTIFY-02](Get-Services-Features.md#notify-02).

Go evidence or governing decision: [internal/modules/notifications/push.go](../../internal/modules/notifications/push.go), [tools/fakerelay](../../tools/fakerelay).

Go server has a relay client and test fake, not a production cloud relay implementation.

### RELAY-02

**Per-household/token/burst rate limits and abuse escalation**

Status: **not started**. Disposition: **external cloud boundary; consolidation undecided**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [tools/fakerelay](../../tools/fakerelay).

Fake relay is test-only; production abuse controls remain in external relay.

### RELAY-03

**HTTPS OAuth bounce into mobile custom schemes**

Status: **not started**. Disposition: **external cloud boundary; consolidation undecided**.

Prerequisites: [MOB-01](Get-Clients-Features.md#mob-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

External cloud redirect remains; self-hosted CC OAuth handling is not the public bounce endpoint.

## jarvis-loop

Current disposition: **external development tooling; consolidation undecided**.

Legacy evidence: [README.md](https://github.com/alexberardi/jarvis-loop/blob/3ea1682f6578cc54ec5263da664f736953f8e295/README.md), [internal](https://github.com/alexberardi/jarvis-loop/tree/3ea1682f6578cc54ec5263da664f736953f8e295/internal), [cmd](https://github.com/alexberardi/jarvis-loop/tree/3ea1682f6578cc54ec5263da664f736953f8e295/cmd).

### LOOP-01

**Work-item state machine, queue, SQLite store, API/CLI and board**

Status: **not started**. Disposition: **external development tooling; consolidation undecided**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/dev/agent-loop.md](../../docs/dev/agent-loop.md).

Already Go, but separate development control plane; not part of jarvisd.

### LOOP-02

**Capability-based runners, triage/build/review/verify/land stages and merge gates**

Status: **in progress**. Disposition: **external development tooling; consolidation undecided**.

Prerequisites: [QA-01](Get-Operations-Features.md#qa-01).

Go evidence or governing decision: [scripts/sandbox](../../scripts/sandbox), [scripts/e2e](../../scripts/e2e), [docs/dev/agent-loop.md](../../docs/dev/agent-loop.md).

This repo supplies verification harnesses; runner/control plane remains external.

### LOOP-03

**Slack item/plan threads, questions, approvals, digest and orchestrator**

Status: **not started**. Disposition: **external development tooling; consolidation undecided**.

Prerequisites: [LOOP-01](Get-External-Features.md#loop-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

External development workflow, not household Jarvis notifications.

### LOOP-04

**Planning/critique, multi-model execution, resource limits and isolated harnesses**

Status: **not started**. Disposition: **external development tooling; consolidation undecided**.

Prerequisites: [LOOP-01](Get-External-Features.md#loop-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

No consolidation into jarvisd found; document ownership before scheduling a move.
