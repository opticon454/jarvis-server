# Nodes feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-node-setup

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-node-setup/tree/73b3660c4bd1c2800dbc54dc6b2d5a51abca0ad6/commands), [services](https://github.com/alexberardi/jarvis-node-setup/tree/73b3660c4bd1c2800dbc54dc6b2d5a51abca0ad6/services), [core](https://github.com/alexberardi/jarvis-node-setup/tree/73b3660c4bd1c2800dbc54dc6b2d5a51abca0ad6/core), [setup](https://github.com/alexberardi/jarvis-node-setup/tree/73b3660c4bd1c2800dbc54dc6b2d5a51abca0ad6/setup).

### NODE-01

**Node enrolment, discovery, MQTT connection and command registration**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-01](Get-Orchestration-Features.md#cc-01), [CC-02](Get-Orchestration-Features.md#cc-02), [CONFLIB-02](Get-Foundations-Features.md#conflib-02).

Go evidence or governing decision: [internal/agent/node](../../internal/agent/node), [cmd/jarvis-agent](../../cmd/jarvis-agent).

Go agent performs the node protocol; existing Python nodes remain compatible.

### NODE-02

**Wake-word model, microphone capture, endpointing and spoken playback**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-01](Get-Nodes-Features.md#node-01), [CC-04](Get-Orchestration-Features.md#cc-04).

Go evidence or governing decision: [internal/agent/voice](../../internal/agent/voice), [internal/audio](../../internal/audio), [docs/agent/go-pi-node.md](../../docs/agent/go-pi-node.md).

Implemented; Pi parity and 24-hour soak remain open.

### NODE-03

**Follow-up listening, barge-in, client-tool turns and speech confirmation**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-02](Get-Nodes-Features.md#node-02), [CC-05](Get-Orchestration-Features.md#cc-05).

Go evidence or governing decision: [internal/agent/voice/followup.go](../../internal/agent/voice/followup.go), [internal/agent/voice/bargein.go](../../internal/agent/voice/bargein.go).

Desktop and rig evidence exists; Pi acceptance remains open.

### NODE-04

**ReSpeaker codec/audio setup, LEDs, physical button and shutdown**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-02](Get-Nodes-Features.md#node-02).

Go evidence or governing decision: [internal/agent/hat](../../internal/agent/hat), [internal/agent/pinode](../../internal/agent/pinode), [cmd/jarvis-agent/hat.go](../../cmd/jarvis-agent/hat.go).

HAT hardware proof recorded, broader install/boot verification remains open.

### NODE-05

**Wi-Fi scan, hotspot provisioning, encrypted credential commit and recovery**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-01](Get-Nodes-Features.md#node-01).

Go evidence or governing decision: [internal/agent/provisioning](../../internal/agent/provisioning), [cmd/jarvis-agent/provision.go](../../cmd/jarvis-agent/provision.go).

HTTP and NetworkManager backends built; reset-to-hotspot hardware rehearsal remains open.

### NODE-06

**Node settings, encrypted secrets/K2, snapshots and command-data storage**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-01](Get-Nodes-Features.md#node-01), [CC-29](Get-Orchestration-Features.md#cc-29).

Go evidence or governing decision: [internal/agent/settings](../../internal/agent/settings), [internal/agent/nodesettings](../../internal/agent/nodesettings).

Desktop support exists; Pi K2/snapshot parity remains an explicit gate.

### NODE-07

**Package install/update/revert/uninstall, manifests, Python environment and rollback**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-06](Get-Nodes-Features.md#node-06), [CC-30](Get-Orchestration-Features.md#cc-30).

Go evidence or governing decision: [internal/agent/packages](../../internal/agent/packages), [cmd/jarvis-agent/packages.go](../../cmd/jarvis-agent/packages.go).

Implemented bridge; package-specific acceptance and Pi rehearsal remain.

### NODE-08

**Foreground SDK command execution and runtime backends**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-07](Get-Nodes-Features.md#node-07), [CC-05](Get-Orchestration-Features.md#cc-05).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/sdkhost](../../internal/agent/sdkhost).

Commands remain Python; desktop has fewer injected backends than the Pi profile.

### NODE-09

**General IJarvisAgent background scheduler, polling and agent context**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go), [internal/agent/packages/manager.go](../../internal/agent/packages/manager.go).

Current loader runs commands/device families/managers, not arbitrary package agents; reminders have a dedicated bridge.

### NODE-10

**Timers: set, list, cancel and expiry announcements**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [TTS-01](Get-Services-Features.md#tts-01).

Go evidence or governing decision: [internal/agent/timers](../../internal/agent/timers), [internal/agent/sdkhost/host.go](../../internal/agent/sdkhost/host.go).

Go timer backend with Python command surface; validate restart/expiry semantics on the Pi.

### NODE-11

**Reminders: create/list/delete/snooze/recurrence and due notifications**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [internal/agent/reminders](../../internal/agent/reminders), [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Dedicated reminder polling is implemented; it does not imply general background agents work.

### NODE-12

**Alert queue, whats-up briefing and notification button playback**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [NODE-11](Get-Nodes-Features.md#node-11).

Go evidence or governing decision: [internal/agent/sdkhost/alerts.go](../../internal/agent/sdkhost/alerts.go), [internal/agent/sdkhost/announce.go](../../internal/agent/sdkhost/announce.go).

Queue/announcement path built; briefing completeness depends on background package agents.

### NODE-13

**Calculator, unit conversion, current time, chat, factual answers and jokes**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [LLM-01](Get-Services-Features.md#llm-01).

Go evidence or governing decision: [internal/agent/packages/python.go](../../internal/agent/packages/python.go), [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Loaded from pinned Python built-ins rather than reimplemented in Go.

### NODE-14

**Shopping/to-do lists, exports and act-on-items callbacks**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [NODE-06](Get-Nodes-Features.md#node-06), [CC-31](Get-Orchestration-Features.md#cc-31).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Python built-ins use Go storage/backends; per-command parity remains to verify.

### NODE-15

**Voice control_device, routine and send_link built-ins**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-18](Get-Orchestration-Features.md#cc-18), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/packages/python.go](../../internal/agent/packages/python.go).

Explicitly in PiBuiltins.Skip; server APIs and mobile control do not close this gap.

### NODE-16

**Device protocol/manager discovery, listing, state, pairing and app control**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/devices](../../internal/agent/devices), [internal/agent/sidecar/devices.go](../../internal/agent/sidecar/devices.go), [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Pi bridge built; vendor devices and voice control need separate checks.

### NODE-17

**Bluetooth scan/pair/connect/disconnect/forget and speaker/phone/bridge roles**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-01](Get-Nodes-Features.md#node-01), [NODE-02](Get-Nodes-Features.md#node-02).

Go evidence or governing decision: [internal/agent/bluetooth](../../internal/agent/bluetooth), [cmd/jarvis-agent/bluetooth.go](../../cmd/jarvis-agent/bluetooth.go).

D-Bus implementation and fake tests exist; real Bluetooth/audio routing acceptance remains.

### NODE-18

**Music ducking, player takeover and post-response playback callbacks**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-02](Get-Nodes-Features.md#node-02), [NODE-08](Get-Nodes-Features.md#node-08).

Go evidence or governing decision: [internal/agent/music](../../internal/agent/music), [internal/agent/voice/music.go](../../internal/agent/voice/music.go).

First music parity slice built; verify real players and microphone conditions.

### NODE-19

**Acoustic echo cancellation while music plays**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NODE-18](Get-Nodes-Features.md#node-18).

Go evidence or governing decision: [docs/agent/go-pi-node.md](../../docs/agent/go-pi-node.md).

Open Pi parity slice 17b.

### NODE-20

**Tentative wake and pre-wake speech handling**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NODE-02](Get-Nodes-Features.md#node-02).

Go evidence or governing decision: [docs/agent/go-pi-node.md](../../docs/agent/go-pi-node.md).

Open Pi parity slice 17c; ordinary audio preroll alone is not full parity.

### NODE-21

**Wake auto-calibration and false-wake cooldown**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NODE-20](Get-Nodes-Features.md#node-20).

Go evidence or governing decision: [docs/agent/go-pi-node.md](../../docs/agent/go-pi-node.md).

Open Pi parity slice 17d.

### NODE-22

**Greeting prefetch, maintenance restart and ACT-LED fallback**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NODE-02](Get-Nodes-Features.md#node-02), [CC-03](Get-Orchestration-Features.md#cc-03).

Go evidence or governing decision: [docs/agent/go-pi-node.md](../../docs/agent/go-pi-node.md).

Open Pi parity slice 17e; split into implementation issues when scheduled.

### NODE-23

**Pi installer, system service activation and package system setup**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-04](Get-Nodes-Features.md#node-04), [NODE-05](Get-Nodes-Features.md#node-05).

Go evidence or governing decision: [scripts/install-pi.sh](../../scripts/install-pi.sh), [internal/agent/packages/system.go](../../internal/agent/packages/system.go).

apt/post_install helpers exist; opt-in activation and hardware rehearsal remain.

### NODE-24

**Signed self-update, healthy-start acknowledgement and rollback**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-23](Get-Nodes-Features.md#node-23).

Go evidence or governing decision: [internal/agent/selfupdate](../../internal/agent/selfupdate), [cmd/jarvis-agent/update.go](../../cmd/jarvis-agent/update.go).

Pi on-device crash-loop rollback reproduction remains open.

### NODE-25

**Factory reset, routines sync and context operation replies**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-05](Get-Nodes-Features.md#node-05), [NODE-06](Get-Nodes-Features.md#node-06), [CC-18](Get-Orchestration-Features.md#cc-18).

Go evidence or governing decision: [internal/agent/actions](../../internal/agent/actions), [internal/agent/reset](../../internal/agent/reset), [cmd/jarvis-agent/reset.go](../../cmd/jarvis-agent/reset.go).

Implemented paths; Pi reset/reprovision and context ACL end-to-end gate remain.

### NODE-26

**Ambient-noise measurement, voice settings, volume/mute and LED preview**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-02](Get-Nodes-Features.md#node-02), [NODE-04](Get-Nodes-Features.md#node-04), [NODE-06](Get-Nodes-Features.md#node-06).

Go evidence or governing decision: [cmd/jarvis-agent/nodesettings.go](../../cmd/jarvis-agent/nodesettings.go), [internal/agent/sdkhost/audio.go](../../internal/agent/sdkhost/audio.go).

Built; hardware calibration, volume and persisted settings need Pi verification.

### NODE-27

**Microphone voice-profile enrolment and verification takes**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-02](Get-Nodes-Features.md#node-02), [STT-03](Get-Services-Features.md#stt-03).

Go evidence or governing decision: [cmd/jarvis-agent/voiceprofile.go](../../cmd/jarvis-agent/voiceprofile.go), [internal/agent/voice/take.go](../../internal/agent/voice/take.go).

Collection/upload flow built; real room and microphone calibration remain.

### NODE-28

**Legacy Pi identity/settings transition and supported-node retirement**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NODE-15](Get-Nodes-Features.md#node-15), [NODE-19](Get-Nodes-Features.md#node-19), [NODE-21](Get-Nodes-Features.md#node-21), [NODE-22](Get-Nodes-Features.md#node-22), [NODE-24](Get-Nodes-Features.md#node-24), [NODE-25](Get-Nodes-Features.md#node-25), [NODE-26](Get-Nodes-Features.md#node-26), [NODE-27](Get-Nodes-Features.md#node-27).

Go evidence or governing decision: [docs/agent/go-pi-node.md](../../docs/agent/go-pi-node.md).

Needs explicit enrol-new versus preserve-identity choice, hardware parity, 24-hour soak and rollback evidence.

### NODE-29

**Node adapter training and custom wake-model training tools**

Status: **not started**. Disposition: **cut adapters; retain offline wake tooling**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

LoRA is cut; wake-model training remains offline tooling, not an embedded trainer.

### NODE-30

**Inbox attachments from Python SDK packages**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [NOTIFY-05](Get-Services-Features.md#notify-05).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

AgentInbox.post lacks attachments; SDK returns attachments_unsupported. Server attachment support alone is insufficient.

### NODE-31

**Legacy local pre-routing/fast-path cache and acoustic test/wake-model tooling**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NODE-02](Get-Nodes-Features.md#node-02).

Go evidence or governing decision: [internal/agent/voice](../../internal/agent/voice), [docs/agent/go-pi-node.md](../../docs/agent/go-pi-node.md).

No complete native replacement demonstrated; separate runtime parity from offline calibration/training utilities.

## jarvis-host-agent

Current disposition: **absorb**.

Legacy evidence: [internal](https://github.com/alexberardi/jarvis-host-agent/tree/87118b92689aa2cbf382d644cd3d79866e3c1096/internal), [cmd](https://github.com/alexberardi/jarvis-host-agent/tree/87118b92689aa2cbf382d644cd3d79866e3c1096/cmd), [spikes/pinode](https://github.com/alexberardi/jarvis-host-agent/tree/87118b92689aa2cbf382d644cd3d79866e3c1096/spikes/pinode).

### DESK-01

**Desktop enrolment, credentials, MQTT tools and start-at-login service**

Status: **complete**. Disposition: **absorb**.

Prerequisites: [NODE-01](Get-Nodes-Features.md#node-01).

Go evidence or governing decision: [internal/agent/node](../../internal/agent/node), [internal/agent/deploy](../../internal/agent/deploy), [cmd/jarvis-agent](../../cmd/jarvis-agent).

Existing Go proof of concept moved into this repository.

### DESK-02

**Discover/launch allowed apps, open links and get/set host volume**

Status: **complete**. Disposition: **absorb**.

Prerequisites: [DESK-01](Get-Nodes-Features.md#desk-01).

Go evidence or governing decision: [internal/agent/apps](../../internal/agent/apps), [internal/agent/host](../../internal/agent/host), [internal/agent/desktop](../../internal/agent/desktop).

Platform implementations exist; Windows hardware coverage is limited.

### DESK-03

**List/focus/minimise/close windows with platform accessibility support**

Status: **complete**. Disposition: **absorb**.

Prerequisites: [DESK-02](Get-Nodes-Features.md#desk-02).

Go evidence or governing decision: [internal/agent/host](../../internal/agent/host), [internal/agent/macax](../../internal/agent/macax), [internal/agent/x11](../../internal/agent/x11).

Uses native platform paths and permissions.

### DESK-04

**Python command sidecar proof of concept**

Status: **complete**. Disposition: **absorb**.

Prerequisites: [DESK-01](Get-Nodes-Features.md#desk-01).

Go evidence or governing decision: [internal/agent/sidecar](../../internal/agent/sidecar).

Bridge absorbed; full package feature migration remains under NODE-08/NODE-09.

## jarvis-command-sdk

Current disposition: **retain Python boundary**.

Legacy evidence: [jarvis_command_sdk](https://github.com/alexberardi/jarvis-command-sdk/tree/708193cc3e74d23294cf05625dabdac61336d57e/jarvis_command_sdk).

### SDK-01

**Command metadata, parameters, examples, validation and responses**

Status: **in progress**. Disposition: **retain Python boundary**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Python SDK pinned and consumed by Go agent; no native Go command-authoring SDK.

### SDK-02

**Scoped secrets, persistent records, schemas and summary hooks**

Status: **in progress**. Disposition: **retain Python boundary**.

Prerequisites: [NODE-06](Get-Nodes-Features.md#node-06), [NODE-08](Get-Nodes-Features.md#node-08).

Go evidence or governing decision: [internal/agent/settings](../../internal/agent/settings), [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Storage and record hooks bridged; verify each caller and schema flow.

### SDK-03

**LLM/CC/logging/timer/audio/alert/inbox/signal runtime backends**

Status: **in progress**. Disposition: **retain Python boundary**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08).

Go evidence or governing decision: [internal/agent/sdkhost](../../internal/agent/sdkhost).

Pi backends exist; desktop safe defaults and attachment gap limit parity.

### SDK-04

**Agents, schedules, context providers, proposals and interactive callbacks**

Status: **in progress**. Disposition: **retain Python boundary**.

Prerequisites: [NODE-09](Get-Nodes-Features.md#node-09), [CC-23](Get-Orchestration-Features.md#cc-23), [CC-31](Get-Orchestration-Features.md#cc-31).

Go evidence or governing decision: [internal/agent/actions](../../internal/agent/actions), [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Callbacks/context hooks exist; general agent scheduling is missing.

### SDK-05

**Device protocols/managers, field metadata, colours and geocoding**

Status: **in progress**. Disposition: **retain Python boundary**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16).

Go evidence or governing decision: [internal/agent/sidecar/devices.go](../../internal/agent/sidecar/devices.go).

Protocols still implemented by Python SDK packages.

### SDK-06

**Photo parameters and inbox media attachments**

Status: **in progress**. Disposition: **retain Python boundary**.

Prerequisites: [CC-12](Get-Orchestration-Features.md#cc-12), [NODE-30](Get-Nodes-Features.md#node-30).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Photo argument bridge exists; package inbox attachment upload is incomplete.

### SDK-07

**Manifest/spec generation for Forge and authoring tools**

Status: **not started**. Disposition: **retain Python boundary**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

External Python SDK tooling remains; no Go implementation found.

### SDK-08

**Pre-route/fast-path hooks, process helpers and Bluetooth runtime access**

Status: **not started**. Disposition: **retain Python boundary**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

No general sidecar dispatch for these SDK interfaces found; foreground run support alone is insufficient.
