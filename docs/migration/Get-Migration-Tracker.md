# Jarvis feature migration tracker

**Review date:** 2026-10-21, Australia/Perth.

This inventory covers **75 source folders and 315 feature entries**, comparing `/mnt/user/Scripts/GitHub/alexberardi` with this Go checkout. The upstream `alexberardi/jarvis-server` folder is excluded. Pantry is documented but will not be consolidated. The new Display frontend requirement is included.

## Start here

- [Repository inventory and revision snapshot](Get-Repository-Inventory.md): every source folder, including the recipe prototype, developer tools and duplicate command repositories.
- [Dependency maps and migration order](Get-Dependency-Map.md): diagrams, external prerequisites and exact feature-to-feature predecessor IDs.
- [Display frontend](Get-New-Features.md#display-01): Pi external screens, Surface-style devices and Android tablets, with mandatory live Home Assistant status integration.
- [Pantry](Get-Pantry-Features.md): backend, website and isolated runner, explicitly outside consolidation.

## Feature registers

| Register                                       | Features | Complete | In progress | Not started | Excluded |
| ---------------------------------------------- | -------- | -------- | ----------- | ----------- | -------- |
| [Foundations](Get-Foundations-Features.md)     | 33       | 28       | 2           | 3           | 0        |
| [Services](Get-Services-Features.md)           | 44       | 35       | 3           | 6           | 0        |
| [Orchestration](Get-Orchestration-Features.md) | 51       | 39       | 0           | 12          | 0        |
| [Nodes](Get-Nodes-Features.md)                 | 43       | 5        | 26          | 12          | 0        |
| [Clients](Get-Clients-Features.md)             | 24       | 3        | 17          | 4           | 0        |
| [Packages](Get-Packages-Features.md)           | 35       | 0        | 23          | 12          | 0        |
| [Devices](Get-Devices-Features.md)             | 18       | 0        | 15          | 3           | 0        |
| [Operations](Get-Operations-Features.md)       | 29       | 10       | 10          | 9           | 0        |
| [External](Get-External-Features.md)           | 11       | 0        | 2           | 9           | 0        |
| [Pantry](Get-Pantry-Features.md)               | 13       | 0        | 0           | 0           | 13       |
| [New](Get-New-Features.md)                     | 14       | 2        | 6           | 6           | 0        |
| Total                                          | 315      | 122      | 104         | 76          | 13       |

Counts include deliberate cuts, external components and new Go features. They are an inventory distribution, **not a migration completion percentage**. Filter by disposition before turning these into a backlog.

## Status and disposition

Use both fields together. A missing feature can be an intentional cut, an external dependency, or outstanding migration work.

| Status         | Meaning                                                                                                                                                                                                                                            |
| -------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| complete       | The retained capability has an implemented counterpart in this checkout with code/test or recorded contract evidence, and no known implementation gap was identified for the stated slice. This does not certify a fresh runtime or hardware test. |
| in progress    | Some capability is implemented, bridged or supported at the API boundary, but material parity, absorption or explicit acceptance gates remain. A Python sidecar is recorded as partial consolidation.                                              |
| not started    | No functional consolidated counterpart was found for the stated slice, or the implementation is only a plan/stub. This label also appears on explicit cuts; disposition determines whether work is wanted.                                         |
| not applicable | Pantry-only features excluded from consolidation by the user; they remain external and are documented for integration.                                                                                                                             |

The current implementation retains native mobile apps, the cloud relay, macOS host APIs, Python package code and development tools in separate repositories. This inventory records that fact without treating it as user approval to exclude them. Their future absorption or retention needs an explicit scope decision; do not delete or archive those repositories based on this tracker. Only Pantry has the user exclusion given in this task. Existing repository decisions still govern deliberate cuts.

“Complete” is a code migration assessment at the recorded revision, not proof that every legacy response, external provider, hardware combination or client path is equivalent. [The contract policy](../contract/README.md) requires tests to pass against the legacy service before they count as a parity oracle. Existing [STATUS](../STATUS.md) records historical runs; this review did not repeat them. Tests linked from features are evidence sources, not claimed new results.

## Main findings

- Core identity/config/settings/logs, notifications, inference APIs, OCR, recipes and most CC orchestration have Go counterparts. Existing contract and golden suites provide substantial historical coverage.
- Node/package migration is incomplete: commands still load Python SDK code, arbitrary package agents are not scheduled, the Pi skips voice control_device/routine/send_link, and sidecar inbox attachments are unsupported. These shared gaps affect multiple packages.
- Pi audio, provisioning, settings, updates and hardware acceptance remain open even where implementation has landed. Retirement requires the explicit parity/rollback/soak gates, not just a passing build.
- Camera streaming is deferred. fastText, LoRA, old prompt providers, legacy MCP and selected infrastructure/unused surfaces were deliberately cut. Forge test-install was restored after an earlier cut.
- Native mobile functionality supported by Go APIs is not the same as moving the mobile client into this repo. Public documentation and installer transition also remain separate deliverables.
- Display frontend is new work. Existing responsive admin pages are not its completion evidence; live Home Assistant device status and per-device-class acceptance are mandatory.

## Scope boundaries and acceptance

Each ID is a capability slice. Related CRUD variants or command actions share a row where they share an implementation and acceptance gate. Provider alternatives, foreground versus background behaviour, device protocols, client-only functionality and explicit cuts remain distinct. Implementation plans and README claims are labelled rather than silently counted as delivered features.

For a row to move to complete, verify the actual call path, close the remaining condition in its note, and add the relevant test/rehearsal evidence. For hardware/provider integrations, record the platform/device/provider used and the result. For API replacement, record the legacy and Go contract result, including skips. For Python bridges, explicitly decide whether supported execution is the end state or native consolidation is required before closing the consolidation work.

For the new display, acceptance includes live HA changes originating outside Jarvis, authorised actions with resulting state feedback, initial load, stale/offline states, reconnect/resynchronisation, access isolation and operation on all three target device classes. No voice microphone, native Android application, wall-mount hardware purchase or new theme has been added to the requirement.

## Existing material to retain

This tracker supplements the existing architecture and decision documents; it does not replace their implementation detail. Read [PLAN](../PLAN.md), [STATUS](../STATUS.md), [CC specifications](../cc/README.md), [Pi parity and retirement](../agent/go-pi-node.md), [legacy import](../install/legacy-import.md), [cutover runbook](../install/cutover-runbook.md) and [recorded decisions](../decisions/README.md) alongside it. Later decisions and live code take precedence over stale high-level README claims.

## GitHub tracking

[Board, roadmap, issue conventions and heading abbreviations](Get-GitHub-Tracking.md) describe the GitHub tracking setup.

Use the feature IDs as stable external keys; keep them when an item is split or moved. Suggested fields: Feature ID, title, source repository, disposition, migration status, predecessor IDs, target code, acceptance/evidence, owner, issue/PR and validation date. Each feature heading already has a stable Markdown anchor. The dependency ledger is published as blocked-by relations. Pantry entries belong in a reference view, not the consolidation sprint.

Split broad action groups into child issues when scheduling implementation, keeping the parent ID as an epic. Give background package agents their own issues so a working foreground command cannot close them accidentally. Record later discoveries as new IDs instead of renumbering the inventory.

## Review limits

This is a local source review and documentation change. During the source review, no services were started, hardware contacted, source repositories modified or migrations applied. The subsequent GitHub upload publishes this inventory as feature issues in both user-created projects. No new runtime parity, live HA, push, call, provider, mobile or Pi soak result is claimed. Local revisions were not refreshed from remotes. The source inventory makes that snapshot explicit.

## Documentation validation

Validation passed for all 75 source folders, 315 unique feature IDs, 527 dependency edges and 2,313 source/internal links. Checks covered local path existence, Markdown heading anchors, an acyclic graph, status totals, aligned tables and whitespace. Commit-pinned source links were checked against the local source files, not fetched from GitHub. The temporary authoring helper was removed; the delivered tracker is Markdown only.
