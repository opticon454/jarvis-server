# New feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-server (current checkout additions)

Current disposition: **new capability, not legacy migration**.

Source: current Go checkout; these are additions, not features imported from the excluded upstream checkout.

### NEW-01

**Desktop wake word, voice pipeline, status icon and on-screen island**

Status: **complete**. Disposition: **new capability, not legacy migration**.

Prerequisites: [DESK-01](Get-Nodes-Features.md#desk-01), [CC-04](Get-Orchestration-Features.md#cc-04).

Go evidence or governing decision: [internal/agent/voice](../../internal/agent/voice), [internal/agent/tray](../../internal/agent/tray), [internal/agent/island](../../internal/agent/island).

New desktop capability; platform/hardware acceptance is not universal.

### NEW-02

**Browser extension/native messaging, site permission grants and browser actions**

Status: **complete**. Disposition: **new capability, not legacy migration**.

Prerequisites: [DESK-01](Get-Nodes-Features.md#desk-01).

Go evidence or governing decision: [internal/agent/browser](../../internal/agent/browser), [internal/agent/browser/extension](../../internal/agent/browser/extension).

Source implementation complete; store distribution/signing remains an operational gate.

### NEW-03

**Read screen and dictate/type/click with platform permission handling**

Status: **in progress**. Disposition: **new capability, not legacy migration**.

Prerequisites: [NEW-01](Get-New-Features.md#new-01), [NEW-02](Get-New-Features.md#new-02).

Go evidence or governing decision: [internal/agent/screen](../../internal/agent/screen), [internal/agent/dictate](../../internal/agent/dictate), [internal/agent/input](../../internal/agent/input).

Built; GNOME/KDE/Windows hardware checks and default-policy review remain open.

### NEW-04

**Device-code pairing, first-run window, sign-out and agent download/update UX**

Status: **in progress**. Disposition: **new capability, not legacy migration**.

Prerequisites: [DESK-01](Get-Nodes-Features.md#desk-01), [AUTH-06](Get-Foundations-Features.md#auth-06).

Go evidence or governing decision: [internal/agent/firstrun](../../internal/agent/firstrun), [internal/modules/cc/pairing.go](../../internal/modules/cc/pairing.go), [internal/modules/admin/agentdownload.go](../../internal/modules/admin/agentdownload.go).

Built; clean-install/store-listing sign-off remains open.

### NEW-05

**Custom rules, approval gates, deterministic web/signal watches and activity**

Status: **in progress**. Disposition: **new capability, not legacy migration**.

Prerequisites: [CC-20](Get-Orchestration-Features.md#cc-20), [CC-22](Get-Orchestration-Features.md#cc-22), [SCRAPE-02](Get-Foundations-Features.md#scrape-02).

Go evidence or governing decision: [internal/modules/cc/policy](../../internal/modules/cc/policy), [internal/modules/cc/watches](../../internal/modules/cc/watches), [docs/autonomy/PLAN.md](../../docs/autonomy/PLAN.md).

P0-P2 recorded complete; admin UI, richer watch sources and mobile screens remain open.

### NEW-06

**Goal wake-loop, budgets, child errands and automatic action review**

Status: **not started**. Disposition: **new capability, not legacy migration**.

Prerequisites: [NEW-05](Get-New-Features.md#new-05).

Go evidence or governing decision: [docs/autonomy/PLAN.md](../../docs/autonomy/PLAN.md).

Planned autonomy phases P5/P6; do not count plans as implementation.

### NEW-07

**Computer task execution, confirmations, checkout safeguards and verification**

Status: **in progress**. Disposition: **new capability, not legacy migration**.

Prerequisites: [NEW-02](Get-New-Features.md#new-02), [NEW-03](Get-New-Features.md#new-03), [CC-20](Get-Orchestration-Features.md#cc-20).

Go evidence or governing decision: [internal/modules/cc/computer](../../internal/modules/cc/computer), [internal/agent/computer](../../internal/agent/computer), [docs/agent/COMPUTER-CONTROL.md](../../docs/agent/COMPUTER-CONTROL.md).

Code exists; project plan still describes active implementation and validation.

### NEW-08

**Record demonstrations, keyframes, teach/edit/replay workflows and spreadsheet steps**

Status: **in progress**. Disposition: **new capability, not legacy migration**.

Prerequisites: [NEW-07](Get-New-Features.md#new-07), [CC-21](Get-Orchestration-Features.md#cc-21).

Go evidence or governing decision: [internal/modules/cc/recordings.go](../../internal/modules/cc/recordings.go), [internal/modules/cc/workflows_learn.go](../../internal/modules/cc/workflows_learn.go), [web/admin/src/pages/WorkflowsPage.tsx](../../web/admin/src/pages/WorkflowsPage.tsx).

Implemented surfaces, active development and platform acceptance remain.

### NEW-09

**Household-editable text prompts and model tool/thinking traits**

Status: **in progress**. Disposition: **new capability, not legacy migration**.

Prerequisites: [CC-07](Get-Orchestration-Features.md#cc-07).

Go evidence or governing decision: [internal/modules/cc/household_prompt.go](../../internal/modules/cc/household_prompt.go), [internal/modules/cc/llmmode.go](../../internal/modules/cc/llmmode.go), [docs/cc/prompt-contract.md](../../docs/cc/prompt-contract.md).

Implementation is ahead of STATUS checklist; re-run prompt/voice/admin acceptance before closing.

### DISPLAY-01

**Display frontend: touch-friendly Jarvis dashboard for Raspberry Pi external screens, Surface-style devices and Android tablets**

Status: **not started**. Disposition: **new capability, not legacy migration**.

Prerequisites: [WEB-01](Get-Clients-Features.md#web-01), [AUTH-04](Get-Foundations-Features.md#auth-04), [DISPLAY-03](Get-New-Features.md#display-03).

Go evidence or governing decision: [web/admin/src/pages](../../web/admin/src/pages).

New user requirement, 2026-10-21. Existing admin pages are a reuse candidate, not a completed display frontend.

### DISPLAY-02

**Device/room tiles, live status, availability and controls with feedback**

Status: **not started**. Disposition: **new capability, not legacy migration**.

Prerequisites: [DISPLAY-01](Get-New-Features.md#display-01), [DISPLAY-03](Get-New-Features.md#display-03), [CC-25](Get-Orchestration-Features.md#cc-25), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/modules/cc/smarthome.go](../../internal/modules/cc/smarthome.go).

Dashboard must show real Home Assistant device state and allow authorised actions; no static/mock-only acceptance.

### DISPLAY-03

**Mandatory Home Assistant live status bridge for the display frontend**

Status: **not started**. Disposition: **new capability, not legacy migration**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04), [CC-25](Get-Orchestration-Features.md#cc-25).

Go evidence or governing decision: [internal/modules/cc/smarthome.go](../../internal/modules/cc/smarthome.go).

Mandatory: initial Home Assistant entity state plus live change events, household-scoped delivery, reconnect/resynchronisation and stale/unavailable indication. Reuse the existing node device.state event design where suitable; connector location/transport remains an implementation decision. A static dashboard or manual refresh is insufficient.

### DISPLAY-04

**Full-screen/kiosk use, responsive touch layout and reconnect recovery on all three device classes**

Status: **not started**. Disposition: **new capability, not legacy migration**.

Prerequisites: [DISPLAY-02](Get-New-Features.md#display-02).

Go evidence or governing decision: [web/admin](../../web/admin).

Browser-based delivery is the initial proposal; no separate native app assumed. Validate actual Pi screen, Surface browser and Android tablet.

### DISPLAY-05

**Display acceptance: live HA changes, control round-trip, access isolation and loss/recovery of connection**

Status: **not started**. Disposition: **new capability, not legacy migration**.

Prerequisites: [DISPLAY-04](Get-New-Features.md#display-04), [QA-04](Get-Operations-Features.md#qa-04).

Go evidence or governing decision: [e2e](../../e2e).

Not complete until external HA changes appear live, failed controls are visible and reconnection restores authoritative state on each target class.
