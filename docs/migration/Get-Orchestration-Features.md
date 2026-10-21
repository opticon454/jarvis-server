# Orchestration feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-command-center

Current disposition: **consolidate**.

Legacy evidence: [app/api](https://github.com/alexberardi/jarvis-command-center/tree/bcdb4185315d897614d37aff2d2f94dec2653475/app/api), [app/core](https://github.com/alexberardi/jarvis-command-center/tree/bcdb4185315d897614d37aff2d2f94dec2653475/app/core), [app/services](https://github.com/alexberardi/jarvis-command-center/tree/bcdb4185315d897614d37aff2d2f94dec2653475/app/services).

### CC-01

**Node registration, heartbeat, command catalogue and liveness**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-06](Get-Foundations-Features.md#auth-06), [CONF-01](Get-Foundations-Features.md#conf-01).

Go evidence or governing decision: [internal/modules/cc/nodes.go](../../internal/modules/cc/nodes.go), [internal/modules/cc/commands.go](../../internal/modules/cc/commands.go), [contract/cc_nodes_test.go](../../contract/cc_nodes_test.go).

Core node wire contract retained.

### CC-02

**Authenticated MQTT broker, request correlation and task results**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-01](Get-Orchestration-Features.md#cc-01).

Go evidence or governing decision: [internal/platform/mqtt](../../internal/platform/mqtt), [internal/modules/cc/bus.go](../../internal/modules/cc/bus.go), [contract/cc_mqtt_test.go](../../contract/cc_mqtt_test.go).

Embedded broker replaces Mosquitto; node-side feature execution is tracked independently.

### CC-03

**Wake greeting generation and playback requests**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-01](Get-Services-Features.md#llm-01), [TTS-01](Get-Services-Features.md#tts-01), [CC-02](Get-Orchestration-Features.md#cc-02).

Go evidence or governing decision: [internal/modules/cc/warmup.go](../../internal/modules/cc/warmup.go), [internal/modules/cc/voice.go](../../internal/modules/cc/voice.go).

Server wake-response support retained; Pi prefetch remains separate.

### CC-04

**Blocking/streaming voice commands, continuation and conversation state**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [STT-01](Get-Services-Features.md#stt-01), [TTS-02](Get-Services-Features.md#tts-02), [LLM-02](Get-Services-Features.md#llm-02), [CC-02](Get-Orchestration-Features.md#cc-02).

Go evidence or governing decision: [internal/modules/cc/voice_routes.go](../../internal/modules/cc/voice_routes.go), [internal/modules/cc/continue.go](../../internal/modules/cc/continue.go), [contract/cc_voice_test.go](../../contract/cc_voice_test.go).

Contract proves shapes; full voice quality requires model and microphone checks.

### CC-05

**Server/client tool registry, selection, execution and result reinjection**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-04](Get-Orchestration-Features.md#cc-04).

Go evidence or governing decision: [internal/modules/cc/engine.go](../../internal/modules/cc/engine.go), [internal/modules/cc/nodetools.go](../../internal/modules/cc/nodetools.go), [internal/modules/cc/servertools](../../internal/modules/cc/servertools).

Node client-tool execution is not itself a port of every package.

### CC-06

**Malformed tool-call parsing, argument validation and clarification/refinement**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-05](Get-Orchestration-Features.md#cc-05).

Go evidence or governing decision: [internal/modules/cc/parse](../../internal/modules/cc/parse), [internal/modules/cc/refine.go](../../internal/modules/cc/refine.go), [fixtures/golden](../../fixtures/golden).

Parser and validation behaviour have golden/unit evidence.

### CC-07

**Qwen 3.x and test OpenAI prompt providers; persona/context rendering**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-03](Get-Services-Features.md#llm-03), [CC-05](Get-Orchestration-Features.md#cc-05).

Go evidence or governing decision: [internal/modules/cc/prompts](../../internal/modules/cc/prompts), [fixtures/golden](../../fixtures/golden).

Retained providers are tested against frozen prompt fixtures.

### CC-08

**Relative dates, ISO guards, household time zones and exact current-time context**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-07](Get-Orchestration-Features.md#cc-07), [LLM-14](Get-Services-Features.md#llm-14).

Go evidence or governing decision: [internal/modules/cc/dates](../../internal/modules/cc/dates), [internal/modules/cc/timezone.go](../../internal/modules/cc/timezone.go), [fixtures/golden](../../fixtures/golden).

Date handling has deliberate fixes; preserve DST and time-zone cases.

### CC-09

**Speaker identity, membership validation and conversational stickiness**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [STT-03](Get-Services-Features.md#stt-03), [AUTH-04](Get-Foundations-Features.md#auth-04), [CC-04](Get-Orchestration-Features.md#cc-04).

Go evidence or governing decision: [internal/modules/cc/identity.go](../../internal/modules/cc/identity.go), [internal/modules/cc/voice_profiles.go](../../internal/modules/cc/voice_profiles.go).

Speaker continuity is conversation-scoped. Legacy per-node last-speaker/stickiness is deliberately removed; do not carry identity between conversations.

### CC-10

**Follow-up intent gate, ambient context and transcript filtering**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-04](Get-Orchestration-Features.md#cc-04).

Go evidence or governing decision: [internal/modules/cc/followup_gate.go](../../internal/modules/cc/followup_gate.go), [internal/modules/cc/ambient_context.go](../../internal/modules/cc/ambient_context.go), [internal/modules/cc/textfilter](../../internal/modules/cc/textfilter).

Room-directed and non-directed speech handled; Pi audio heuristics remain incomplete.

### CC-11

**Text/mobile chat, SSE, client-tool continue and conversation history**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-05](Get-Orchestration-Features.md#cc-05), [AUTH-04](Get-Foundations-Features.md#auth-04).

Go evidence or governing decision: [internal/modules/cc/mobile_chat.go](../../internal/modules/cc/mobile_chat.go), [internal/modules/cc/textadapt.go](../../internal/modules/cc/textadapt.go).

Server surface retained for external mobile and embedded web clients.

### CC-12

**Chat photos, tool image arguments and conversation compaction**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-11](Get-Orchestration-Features.md#cc-11), [LLM-10](Get-Services-Features.md#llm-10), [PLAT-03](Get-Foundations-Features.md#plat-03).

Go evidence or governing decision: [internal/modules/cc/chat_images.go](../../internal/modules/cc/chat_images.go), [internal/modules/cc/tool_images.go](../../internal/modules/cc/tool_images.go), [internal/modules/cc/compaction.go](../../internal/modules/cc/compaction.go).

Image payloads and bounded conversational context implemented.

### CC-13

**Traces, recent transcripts and per-turn diagnostics**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-04](Get-Orchestration-Features.md#cc-04), [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/modules/cc/traces.go](../../internal/modules/cc/traces.go), [internal/modules/cc/transcripts.go](../../internal/modules/cc/transcripts.go), [contract/cc_voice_test.go](../../contract/cc_voice_test.go).

Admin trace and mobile recent-command surfaces retained.

### CC-14

**Memory remember/recall/forget, embedding search and mobile CRUD**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04), [LLM-07](Get-Services-Features.md#llm-07), [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/modules/cc/memory_tools.go](../../internal/modules/cc/memory_tools.go), [internal/modules/cc/memory_store.go](../../internal/modules/cc/memory_store.go), [internal/modules/cc/memory_routes.go](../../internal/modules/cc/memory_routes.go).

Household and user scoping retained; import must re-embed old vectors.

### CC-15

**Background memory extraction, characterisation and context injection**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-14](Get-Orchestration-Features.md#cc-14), [LLM-08](Get-Services-Features.md#llm-08).

Go evidence or governing decision: [internal/modules/cc/memory_jobs.go](../../internal/modules/cc/memory_jobs.go), [internal/modules/cc/memory_agent.go](../../internal/modules/cc/memory_agent.go), [internal/modules/cc/memory_prompt.go](../../internal/modules/cc/memory_prompt.go).

Background jobs replace remote callbacks.

### CC-16

**Quick web search and answer-question server tools**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-05](Get-Orchestration-Features.md#cc-05), [SCRAPE-01](Get-Foundations-Features.md#scrape-01), [SCRAPE-02](Get-Foundations-Features.md#scrape-02).

Go evidence or governing decision: [internal/modules/cc/servertools](../../internal/modules/cc/servertools).

Privacy gate controls external web access.

### CC-17

**Deep research jobs with summarisation and inbox delivery**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-16](Get-Orchestration-Features.md#cc-16), [LLM-08](Get-Services-Features.md#llm-08), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [internal/modules/cc/servertools](../../internal/modules/cc/servertools).

Queued research completes through notifications.

### CC-18

**Routines CRUD, sync and multi-step execution**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-05](Get-Orchestration-Features.md#cc-05), [CC-02](Get-Orchestration-Features.md#cc-02), [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/modules/cc/routines.go](../../internal/modules/cc/routines.go), [internal/modules/cc/routines_run.go](../../internal/modules/cc/routines_run.go).

Server routines are implemented; Pi routine voice command is a separate gap.

### CC-19

**Once/interval/cron schedules, time zones and scheduled errands**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-18](Get-Orchestration-Features.md#cc-18), [PLAT-02](Get-Foundations-Features.md#plat-02), [CC-08](Get-Orchestration-Features.md#cc-08).

Go evidence or governing decision: [internal/modules/cc/schedules.go](../../internal/modules/cc/schedules.go), [internal/platform/scheduler](../../internal/platform/scheduler).

Durable scheduling replaces separate worker processes.

### CC-20

**Errand planning, step dependencies, execution, progress and cancellation**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-05](Get-Orchestration-Features.md#cc-05), [LLM-08](Get-Services-Features.md#llm-08), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [internal/modules/cc/errands](../../internal/modules/cc/errands), [internal/modules/cc/errands.go](../../internal/modules/cc/errands.go).

Use tool availability and confirmation policies at execution time.

### CC-21

**Workflow definitions and execution, references to earlier step values**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-20](Get-Orchestration-Features.md#cc-20).

Go evidence or governing decision: [internal/modules/cc/workflows.go](../../internal/modules/cc/workflows.go), [internal/modules/cc/stepvalues](../../internal/modules/cc/stepvalues).

Workflow runtime retained; taught desktop workflows are additional features.

### CC-22

**Signal publication, catalogue, subscriptions and automation execution**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-02](Get-Orchestration-Features.md#cc-02), [PLAT-02](Get-Foundations-Features.md#plat-02).

Go evidence or governing decision: [internal/modules/cc/signals.go](../../internal/modules/cc/signals.go), [internal/modules/cc/signals_reactions.go](../../internal/modules/cc/signals_reactions.go).

Package agents must actually publish signals for these automations to trigger.

### CC-23

**Directed proposals, suppressions and confirmation callbacks**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-22](Get-Orchestration-Features.md#cc-22), [LLM-08](Get-Services-Features.md#llm-08), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [internal/modules/cc/proposals.go](../../internal/modules/cc/proposals.go), [internal/modules/cc/callbacks.go](../../internal/modules/cc/callbacks.go).

Card dispatcher and user-authored automations retained; proactive situation matching is deliberately cut (CC-39).

### CC-24

**Attention arbitration, presence, time windows and notification gating**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-23](Get-Orchestration-Features.md#cc-23), [NOTIFY-02](Get-Services-Features.md#notify-02).

Go evidence or governing decision: [internal/modules/cc/attention.go](../../internal/modules/cc/attention.go), [internal/modules/cc/timewindow](../../internal/modules/cc/timewindow), [internal/modules/cc/notify_gated.go](../../internal/modules/cc/notify_gated.go).

Controls when and where proactive output is delivered.

### CC-25

**Rooms, device registry, discovery/import and device CRUD**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04), [CC-02](Get-Orchestration-Features.md#cc-02).

Go evidence or governing decision: [internal/modules/cc/smarthome.go](../../internal/modules/cc/smarthome.go), [internal/modules/cc/smarthome_devices.go](../../internal/modules/cc/smarthome_devices.go), [internal/modules/cc/smarthome_rooms.go](../../internal/modules/cc/smarthome_rooms.go).

Server registry is separate from node protocol execution.

### CC-26

**Smart-home control/status, provider OAuth and node credential hand-off**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-25](Get-Orchestration-Features.md#cc-25).

Go evidence or governing decision: [internal/modules/cc/smarthome_control.go](../../internal/modules/cc/smarthome_control.go), [internal/modules/cc/smarthome_oauth.go](../../internal/modules/cc/smarthome_oauth.go).

Compatible node-side family or manager still required.

### CC-27

**Camera credentials, sessions and streamed camera access**

Status: **not started**. Disposition: **deferred retained capability**.

Prerequisites: [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/modules/cc/smarthome.go](../../internal/modules/cc/smarthome.go), [docs/cc/07-smart-home.md](../../docs/cc/07-smart-home.md).

D29 explicitly defers camera streaming: listing exists, stream start/proxy are unavailable stubs. No go2rtc/HLS runtime is ported.

### CC-28

**Node provisioning proxy, update tasks, cancellation and reset/sign-out**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-02](Get-Orchestration-Features.md#cc-02).

Go evidence or governing decision: [internal/modules/cc/provisioning.go](../../internal/modules/cc/provisioning.go), [internal/modules/cc/updates.go](../../internal/modules/cc/updates.go), [internal/modules/cc/reset.go](../../internal/modules/cc/reset.go), [contract/cc_node_routes_test.go](../../contract/cc_node_routes_test.go).

Server routes implemented; hardware rollback/reset tests are separate.

### CC-29

**Node settings requests, encrypted K2 forwarding and snapshot coordination**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-02](Get-Orchestration-Features.md#cc-02), [AUTH-04](Get-Foundations-Features.md#auth-04).

Go evidence or governing decision: [internal/modules/cc/settingsreq.go](../../internal/modules/cc/settingsreq.go), [contract/cc_mqtt_rows_test.go](../../contract/cc_mqtt_rows_test.go).

Desktop and Pi node implementations have different validation maturity.

### CC-30

**Pantry browse/install/update/revert/uninstall coordination and privacy switch**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-02](Get-Orchestration-Features.md#cc-02), [AUTH-04](Get-Foundations-Features.md#auth-04).

Go evidence or governing decision: [internal/modules/cc/packages.go](../../internal/modules/cc/packages.go), [contract/cc_mqtt_rows_test.go](../../contract/cc_mqtt_rows_test.go).

External Pantry remains; package runtime compatibility is tracked under NODE and PKG.

### CC-31

**Command data schemas, records, CRUD and mobile summaries**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-02](Get-Orchestration-Features.md#cc-02), [AUTH-04](Get-Foundations-Features.md#auth-04).

Go evidence or governing decision: [internal/modules/cc/commanddata.go](../../internal/modules/cc/commanddata.go), [internal/modules/cc/nodeplugin.go](../../internal/modules/cc/nodeplugin.go).

Server forwards to the owning node with user context.

### CC-32

**Public plugin API: inbox, push, LLM, callbacks, signals and context requests**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-02](Get-Orchestration-Features.md#cc-02), [NOTIFY-03](Get-Services-Features.md#notify-03), [LLM-01](Get-Services-Features.md#llm-01), [CC-22](Get-Orchestration-Features.md#cc-22).

Go evidence or governing decision: [internal/modules/cc/nodeplugin.go](../../internal/modules/cc/nodeplugin.go), [contract/cc_node_routes_test.go](../../contract/cc_node_routes_test.go).

Keep these endpoints stable for packages and SDK callers.

### CC-33

**Mobile audio, TTS/STT media proxy and voice-profile collection**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [STT-03](Get-Services-Features.md#stt-03), [TTS-01](Get-Services-Features.md#tts-01), [CC-02](Get-Orchestration-Features.md#cc-02).

Go evidence or governing decision: [internal/modules/cc/media.go](../../internal/modules/cc/media.go), [internal/modules/cc/mobile_audio.go](../../internal/modules/cc/mobile_audio.go), [contract/cc_media_test.go](../../contract/cc_media_test.go).

Server route coverage does not replace phone/Pi microphone tests.

### CC-34

**Phone contacts, call context, sessions and dial scheduling**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04), [CC-20](Get-Orchestration-Features.md#cc-20), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [internal/modules/cc/phone](../../internal/modules/cc/phone), [internal/modules/cc/phone_wire.go](../../internal/modules/cc/phone_wire.go).

Phone media processing is tracked under PHONE.

### CC-35

**LoRA proposals/training/deployment/rollback and node adapter APIs**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Explicit cut; remove or hide dependent mobile screens before retiring old services.

### CC-36

**Install arbitrary Python prompt providers and unused legacy endpoints**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md), [internal/modules/cc/testinstall.go](../../internal/modules/cc/testinstall.go).

Cut provider mechanism; compatibility stubs are not implementations.

### CC-37

**fastText tool-router training and routing optimisation**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: [CC-05](Get-Orchestration-Features.md#cc-05).

Go evidence or governing decision: [docs/cc/02-tool-loop.md](../../docs/cc/02-tool-loop.md).

D9 explicitly cuts fastText, /tool-router/train, router hints and old fast/tool-stream paths. This supersedes the earlier PLAN reference to porting fastText.

### CC-38

**Forge temporary test installation coordination**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-30](Get-Orchestration-Features.md#cc-30).

Go evidence or governing decision: [internal/modules/cc/testinstall.go](../../internal/modules/cc/testinstall.go), [internal/agent/packages/testinstall.go](../../internal/agent/packages/testinstall.go).

Test-install routes, share-code verification and request lifecycle are ported; the 2026-10-08 decision reverses the earlier cut. Package/Pi runtime acceptance remains under NODE-07.

### CC-39

**Proactive situation matching, anti-nag state and signal precision evaluation harness**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/cc/10-signals-attention-proposals.md](../../docs/cc/10-signals-attention-proposals.md).

D17 explicitly cuts this plane; directed proposals and user-authored signal automations remain.

### CC-40

**Emotion/affect analysis and legacy ambient-grounding pass**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/cc/01-voice-pipeline.md](../../docs/cc/01-voice-pipeline.md).

D38/D9 cuts: affect accepted as null/no-op; newer ambient signal context is not the old affect analyser.

### CC-41

**Unused attention consent/tier/feedback tables and inspection-only endpoints**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/cc/00-platform.md](../../docs/cc/00-platform.md), [docs/cc/10-signals-attention-proposals.md](../../docs/cc/10-signals-attention-proposals.md).

Never-written/dead surfaces deliberately omitted; journal, budgets and quiet hours remain.

### CC-42

**Phone number lookup, do-not-call checks, approval cards and call cancellation/reaping**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-34](Get-Orchestration-Features.md#cc-34).

Go evidence or governing decision: [internal/modules/cc/phone/plan.go](../../internal/modules/cc/phone/plan.go), [internal/modules/cc/phone/reaper.go](../../internal/modules/cc/phone/reaper.go), [internal/modules/cc/phone/plan_test.go](../../internal/modules/cc/phone/plan_test.go).

Includes contact matching, pending draft lifecycle and in-call duration guard.

### CC-43

**Call transcripts/audio, escalation and goal-result reporting**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PHONE-03](Get-Orchestration-Features.md#phone-03), [PLAT-03](Get-Foundations-Features.md#plat-03).

Go evidence or governing decision: [internal/modules/cc/phone](../../internal/modules/cc/phone), [internal/modules/cc/phone/call_test.go](../../internal/modules/cc/phone/call_test.go).

Session persistence and goal result handling ported; call recording retention was not added by the port.

## jarvis-phone-gateway

Current disposition: **consolidate**.

Legacy evidence: [main.py](https://github.com/alexberardi/jarvis-phone-gateway/blob/56eedd40fffd9c5effdaad5821efff34b3ecb292/main.py), [audio](https://github.com/alexberardi/jarvis-phone-gateway/tree/56eedd40fffd9c5effdaad5821efff34b3ecb292/audio), [telephony](https://github.com/alexberardi/jarvis-phone-gateway/tree/56eedd40fffd9c5effdaad5821efff34b3ecb292/telephony), [llm](https://github.com/alexberardi/jarvis-phone-gateway/tree/56eedd40fffd9c5effdaad5821efff34b3ecb292/llm), [services](https://github.com/alexberardi/jarvis-phone-gateway/tree/56eedd40fffd9c5effdaad5821efff34b3ecb292/services).

### PHONE-01

**Twilio webhook authentication, outbound dial and media WebSocket sessions**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CC-34](Get-Orchestration-Features.md#cc-34).

Go evidence or governing decision: [internal/modules/cc/phone/live](../../internal/modules/cc/phone/live), [internal/modules/cc/phone_wire.go](../../internal/modules/cc/phone_wire.go).

Absorbed in-process; external Twilio credentials/connectivity still required.

### PHONE-02

**Mu-law/PCM codec, VAD endpointing and duplex media loop**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PHONE-01](Get-Orchestration-Features.md#phone-01), [STT-01](Get-Services-Features.md#stt-01), [TTS-02](Get-Services-Features.md#tts-02).

Go evidence or governing decision: [internal/modules/cc/phone/live](../../internal/modules/cc/phone/live).

Codec and media runtime implemented; real call quality was not exercised by this review.

### PHONE-03

**Streaming LLM call turns, restricted tools and call completion reporting**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PHONE-02](Get-Orchestration-Features.md#phone-02), [LLM-02](Get-Services-Features.md#llm-02).

Go evidence or governing decision: [internal/modules/cc/phone/live](../../internal/modules/cc/phone/live).

Legacy README described foundation-only status; Go contains the live loop.

## jarvis-pp-hermes

Current disposition: **retire**.

Legacy evidence: [README.md](https://github.com/alexberardi/jarvis-pp-hermes/blob/0b688b4bc50dcabbb1a1fcf3f7f5c924e05b486d/README.md), [prompt_providers](https://github.com/alexberardi/jarvis-pp-hermes/tree/0b688b4bc50dcabbb1a1fcf3f7f5c924e05b486d/prompt_providers).

### HERMES-01

**Hermes untrained, LoRA-trained, compressed and MLX prompt variants**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Explicitly excluded by prompt-provider keep list; future text templates are not a completed port.

## jarvis-pp-mistral

Current disposition: **retire**.

Legacy evidence: [README.md](https://github.com/alexberardi/jarvis-pp-mistral/blob/57b39c791192c074c97f6e89943d6ad9353b0021/README.md), [prompt_providers](https://github.com/alexberardi/jarvis-pp-mistral/tree/57b39c791192c074c97f6e89943d6ad9353b0021/prompt_providers).

### MISTRAL-01

**Mistral 7B and Mixtral prompts; XML/native tool output parsing**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Old providers deliberately cut; generic model-engine support does not restore their prompt contracts.

## jarvis-mcp

Current disposition: **retire**.

Legacy evidence: [jarvis_mcp/server.py](https://github.com/alexberardi/jarvis-mcp/blob/2c4b7fa824dbd7031ce018eee7d0c641d659e4a6/jarvis_mcp/server.py), [jarvis_mcp/tools](https://github.com/alexberardi/jarvis-mcp/tree/2c4b7fa824dbd7031ce018eee7d0c641d659e4a6/jarvis_mcp/tools).

### MCP-01

**MCP SSE/HTTP server, tool discovery and configurable tool groups**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Separate MCP service deliberately dropped.

### MCP-02

**MCP logs, health/debug, Docker/database and test tooling**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: [MCP-01](Get-Orchestration-Features.md#mcp-01).

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Underlying admin/log capabilities exist, but there is no equivalent MCP transport.

### MCP-03

**MCP command execution, maths, conversions and date/time tools**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: [MCP-01](Get-Orchestration-Features.md#mcp-01).

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Count separately from native voice tools; not migrated to MCP in this repository.
