# Operations feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-admin

Current disposition: **consolidate**.

Legacy evidence: [src/pages](https://github.com/alexberardi/jarvis-admin/tree/0d86eb58583f37211686edbdf11f293697992a96/src/pages), [server/src/routes](https://github.com/alexberardi/jarvis-admin/tree/0d86eb58583f37211686edbdf11f293697992a96/server/src/routes), [server/src/services](https://github.com/alexberardi/jarvis-admin/tree/0d86eb58583f37211686edbdf11f293697992a96/server/src/services).

### ADMIN-01

**Dashboard, service health, users, nodes and role-gated navigation**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-07](Get-Foundations-Features.md#auth-07), [CONF-03](Get-Foundations-Features.md#conf-03), [CC-01](Get-Orchestration-Features.md#cc-01).

Go evidence or governing decision: [web/admin/src/pages](../../web/admin/src/pages), [internal/modules/admin/tiers.go](../../internal/modules/admin/tiers.go), [contract/admin_test.go](../../contract/admin_test.go).

React UI embedded in jarvisd; gateway replaces Fastify proxy.

### ADMIN-02

**Runtime settings, quick-set equivalents and service connections**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [CONF-04](Get-Foundations-Features.md#conf-04), [SET-02](Get-Foundations-Features.md#set-02).

Go evidence or governing decision: [web/admin/src/pages/SettingsPage.tsx](../../web/admin/src/pages/SettingsPage.tsx), [web/admin/src/pages/ConnectionsPage.tsx](../../web/admin/src/pages/ConnectionsPage.tsx).

Core settings supported; assess legacy quick-set presets individually.

### ADMIN-03

**Hardware detection, first-run account/privacy wizard and model setup**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-01](Get-Foundations-Features.md#auth-01), [LLM-06](Get-Services-Features.md#llm-06).

Go evidence or governing decision: [web/admin/src/pages/SetupWizard.tsx](../../web/admin/src/pages/SetupWizard.tsx), [internal/modules/admin/setupjobs.go](../../internal/modules/admin/setupjobs.go), [contract/admin_test.go](../../contract/admin_test.go).

Compose wizard replaced with jarvisd setup.

### ADMIN-04

**Model catalogue, downloads, inference labels and prompt selection**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-06](Get-Services-Features.md#llm-06), [CC-07](Get-Orchestration-Features.md#cc-07).

Go evidence or governing decision: [web/admin/src/pages/ModelsPage.tsx](../../web/admin/src/pages/ModelsPage.tsx), [internal/modules/admin](../../internal/modules/admin).

Engine-specific live validation stays under LLM-11 through LLM-13.

### ADMIN-05

**Log viewer and trace list/detail**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LOG-02](Get-Foundations-Features.md#log-02), [CC-13](Get-Orchestration-Features.md#cc-13).

Go evidence or governing decision: [web/admin/src/pages/LogsPage.tsx](../../web/admin/src/pages/LogsPage.tsx), [web/admin/src/pages/TracesPage.tsx](../../web/admin/src/pages/TracesPage.tsx).

Embedded replacements exist.

### ADMIN-06

**Update discovery, signed install/rollback progress and service lifecycle**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [OPS-02](Get-Operations-Features.md#ops-02).

Go evidence or governing decision: [web/admin/src/pages/UpdatePage.tsx](../../web/admin/src/pages/UpdatePage.tsx), [internal/modules/admin/update.go](../../internal/modules/admin/update.go).

Server upgrade capability retained; fresh OS install acceptance remains separate.

### ADMIN-07

**Docker container controls, compose reconciliation, env-file editor and native Python service orchestration**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/admin/00-inventory.md](../../docs/admin/00-inventory.md).

Retired implementation surfaces; jarvisd service/model controls replace their operational purpose.

### ADMIN-08

**Phone/tablet layouts, accessibility and clean-install sign-off**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [ADMIN-01](Get-Operations-Features.md#admin-01), [ADMIN-03](Get-Operations-Features.md#admin-03).

Go evidence or governing decision: [web/admin/src/pages](../../web/admin/src/pages), [docs/STATUS.md](../../docs/STATUS.md).

Existing responsive work is not the requested Display frontend; device sign-off remains open.

## jarvis

Current disposition: **consolidate**.

Legacy evidence: [README.md](https://github.com/alexberardi/jarvis/blob/4a4b1c29ad1a8fcaace10b5c7797c4b02eb0a7ef/README.md), [jarvis](https://github.com/alexberardi/jarvis/blob/4a4b1c29ad1a8fcaace10b5c7797c4b02eb0a7ef/jarvis), [install-e2e](https://github.com/alexberardi/jarvis/tree/4a4b1c29ad1a8fcaace10b5c7797c4b02eb0a7ef/install-e2e), [deploy](https://github.com/alexberardi/jarvis/tree/4a4b1c29ad1a8fcaace10b5c7797c4b02eb0a7ef/deploy).

### OPS-01

**Install/start/stop/status, config, diagnostics and legacy-service coexistence**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CONF-01](Get-Foundations-Features.md#conf-01), [AUTH-01](Get-Foundations-Features.md#auth-01).

Go evidence or governing decision: [cmd/jarvisd](../../cmd/jarvisd), [internal/platform/service](../../internal/platform/service), [scripts/install.sh](../../scripts/install.sh).

Single-binary CLI/service replaces multi-repo Compose orchestration.

### OPS-02

**Release archives, platform installers, signature verification, upgrades and rollback**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [OPS-01](Get-Operations-Features.md#ops-01), [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/update](../../internal/update), [.github/workflows/release.yml](../../.github/workflows/release.yml), [scripts/packaging](../../scripts/packaging).

Implementation exists; release and OS rehearsal status must be checked independently.

### OPS-03

**Legacy users/households/nodes/settings/memories/routines/contacts/inbox import**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-06](Get-Foundations-Features.md#auth-06), [CC-14](Get-Orchestration-Features.md#cc-14), [CC-18](Get-Orchestration-Features.md#cc-18), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [cmd/jarvisd/import_legacy.go](../../cmd/jarvisd/import_legacy.go), [internal/legacyimport](../../internal/legacyimport), [docs/install/legacy-import.md](../../docs/install/legacy-import.md).

Read-only source import with dry-run/apply; recipes use RECIPE-10.

### OPS-04

**Install matrix, fake/live-model behaviour tests and cross-platform rehearsal**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [OPS-02](Get-Operations-Features.md#ops-02), [CC-04](Get-Orchestration-Features.md#cc-04).

Go evidence or governing decision: [e2e](../../e2e), [contract](../../contract), [docs/install/cutover-runbook.md](../../docs/install/cutover-runbook.md).

New harness replaces much of install-e2e; remaining hardware/clean-install gates are documented.

### OPS-05

**Remote GPU satellite deployment mode**

Status: **not started**. Disposition: **deferred retained capability**.

Prerequisites: [LLM-06](Get-Services-Features.md#llm-06), [CONF-01](Get-Foundations-Features.md#conf-01).

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Explicit future work; do not equate a remote REST endpoint with a packaged satellite install.

### OPS-06

**Legacy Compose generators, databases, multi-repo update commands and model conversion/training utilities**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Compose control retired; training/conversion remains offline tooling.

## jarvis-installer

Current disposition: **replace deployment model**.

Legacy evidence: [src/components](https://github.com/alexberardi/jarvis-installer/tree/104cf653fe6cd21b8bd181750432dc830350a722/src/components), [src/lib](https://github.com/alexberardi/jarvis-installer/tree/104cf653fe6cd21b8bd181750432dc830350a722/src/lib), [public/service-registry.json](https://github.com/alexberardi/jarvis-installer/blob/104cf653fe6cd21b8bd181750432dc830350a722/public/service-registry.json).

### INSTALL-01

**Public landing page and service/hardware guided setup journey**

Status: **in progress**. Disposition: **replace deployment model**.

Prerequisites: [OPS-02](Get-Operations-Features.md#ops-02), [ADMIN-03](Get-Operations-Features.md#admin-03).

Go evidence or governing decision: [docs/STATUS.md](../../docs/STATUS.md), [docs/install/00-installers.md](../../docs/install/00-installers.md).

Go wizard exists; public installer landing transition remains open in recorded status.

### INSTALL-02

**Compose/env/database-init generation and downloadable ZIP bundle**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/install/00-installers.md](../../docs/install/00-installers.md).

Deliberately superseded by single-binary installers.

### INSTALL-03

**No-terminal install on Linux/macOS/Windows with first-run browser launch**

Status: **in progress**. Disposition: **replace deployment model**.

Prerequisites: [OPS-02](Get-Operations-Features.md#ops-02), [ADMIN-03](Get-Operations-Features.md#admin-03).

Go evidence or governing decision: [scripts/packaging](../../scripts/packaging), [docs/install/01-no-cli-plan.md](../../docs/install/01-no-cli-plan.md).

Artifacts implemented; full per-OS fresh-user sign-off remains open.

## jarvis-integration-tests

Current disposition: **consolidate**.

Legacy evidence: [tests](https://github.com/alexberardi/jarvis-integration-tests/tree/36e0424347df708073231966b74114157c739a8b/tests), [tools](https://github.com/alexberardi/jarvis-integration-tests/tree/36e0424347df708073231966b74114157c739a8b/tools), [.github/workflows](https://github.com/alexberardi/jarvis-integration-tests/tree/36e0424347df708073231966b74114157c739a8b/.github/workflows).

### QA-01

**Isolated auth/config/CC/MQTT integration harness and fake inference services**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-04](Get-Orchestration-Features.md#cc-04).

Go evidence or governing decision: [contract](../../contract), [e2e](../../e2e), [tools/scriptllm](../../tools/scriptllm), [scripts/sandbox](../../scripts/sandbox).

Go contract and sandbox harness cover replacement stack.

### QA-02

**Real-model routing and multi-turn behaviour corpus**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [QA-01](Get-Operations-Features.md#qa-01), [LLM-01](Get-Services-Features.md#llm-01).

Go evidence or governing decision: [tools/pipeline](../../tools/pipeline), [fixtures/golden](../../fixtures/golden).

New scenarios/fixtures exist; full CASE-by-CASE legacy corpus mapping remains to verify.

### QA-03

**Cross-repo PR dispatch, status reporting and fixture orchestration**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [QA-01](Get-Operations-Features.md#qa-01).

Go evidence or governing decision: [.github/workflows](../../.github/workflows), [docs/dev/e2e.md](../../docs/dev/e2e.md).

Monorepo CI exists; external client/package trigger parity is not established by this inventory.

### QA-04

**Web/mobile/push/HA/voice regression and coverage gates**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [QA-01](Get-Operations-Features.md#qa-01).

Go evidence or governing decision: [e2e/coverage.json](../../e2e/coverage.json), [e2e/web](../../e2e/web), [scripts/e2e](../../scripts/e2e).

Harness implementation present; tests were not run during this documentation review.

## jarvis-developer-toolkit

Current disposition: **authoring tooling; consolidation undecided**.

Legacy evidence: [src/jdt](https://github.com/alexberardi/jarvis-developer-toolkit/tree/4484129d2b3f4161fb9364b74ec415cee9e8a12e/src/jdt).

### JDT-01

**Scaffold commands, agents, protocols, managers, providers and routine bundles**

Status: **not started**. Disposition: **authoring tooling; consolidation undecided**.

Prerequisites: [SDK-01](Get-Nodes-Features.md#sdk-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Python jdt remains external; no equivalent Go-authoring CLI found.

### JDT-02

**Generate/update manifests and inspect secrets/parameters/components**

Status: **not started**. Disposition: **authoring tooling; consolidation undecided**.

Prerequisites: [JDT-01](Get-Operations-Features.md#jdt-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

External introspection tooling; runtime manifest parsing is not an authoring replacement.

### JDT-03

**Manifest validation, static analysis and import/instantiation tests**

Status: **in progress**. Disposition: **authoring tooling; consolidation undecided**.

Prerequisites: [NODE-07](Get-Nodes-Features.md#node-07).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

Runtime performs a subset; complete Pantry-compatible jdt testing remains external.

### JDT-04

**Deploy packages locally, into Docker nodes or over SSH**

Status: **not started**. Disposition: **authoring tooling; consolidation undecided**.

Prerequisites: [JDT-01](Get-Operations-Features.md#jdt-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

No consolidated equivalent of the three developer deployment modes found.

## jarvis-docs

Current disposition: **documentation migration**.

Legacy evidence: [mkdocs.yml](https://github.com/alexberardi/jarvis-docs/blob/531436785a51905125bd1faa048dbb0f7d2b4fb6/mkdocs.yml), [docs](https://github.com/alexberardi/jarvis-docs/tree/531436785a51905125bd1faa048dbb0f7d2b4fb6/docs), [deploy.sh](https://github.com/alexberardi/jarvis-docs/blob/531436785a51905125bd1faa048dbb0f7d2b4fb6/deploy.sh).

### DOCS-01

**Public installation, configuration, upgrade, migration and troubleshooting guides**

Status: **in progress**. Disposition: **documentation migration**.

Prerequisites: [OPS-03](Get-Operations-Features.md#ops-03), [INSTALL-03](Get-Operations-Features.md#install-03), [NODE-28](Get-Nodes-Features.md#node-28).

Go evidence or governing decision: [README.md](../../README.md), [docs/install](../../docs/install), [docs/STATUS.md](../../docs/STATUS.md).

Local Go docs exist; public jarvis-docs still documents much of the Python stack.

### DOCS-02

**Architecture, service API, SDK/package/device and security reference**

Status: **in progress**. Disposition: **documentation migration**.

Prerequisites: [SDK-01](Get-Nodes-Features.md#sdk-01), [CC-32](Get-Orchestration-Features.md#cc-32).

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md), [docs/cc](../../docs/cc), [docs/schema](../../docs/schema).

Source specs exist; public reference needs reconciliation with retained/cut features.

### DOCS-03

**MkDocs site, theme/search/navigation and Cloudflare deployment**

Status: **not started**. Disposition: **documentation migration**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/dev](../../docs/dev).

Site remains in jarvis-docs; hosting source is not absorbed into this repository.

### DOCS-04

**Mobile universal-link association and app-not-installed fallback page**

Status: **not started**. Disposition: **documentation migration**.

Prerequisites: [MOB-13](Get-Clients-Features.md#mob-13).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

External site assets remain required for mobile deep links; no move found.
