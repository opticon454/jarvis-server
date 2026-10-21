# Services feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-notifications

Current disposition: **consolidate**.

Legacy evidence: [app/api](https://github.com/alexberardi/jarvis-notifications/tree/7214ab1c176720b88f0508020031f5d2cfb1cc33/app/api), [app/services](https://github.com/alexberardi/jarvis-notifications/tree/7214ab1c176720b88f0508020031f5d2cfb1cc33/app/services).

### NOTIFY-01

**Register, list and revoke push tokens; account data cleanup**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04), [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/modules/notifications/tokens.go](../../internal/modules/notifications/tokens.go), [contract/notifications_test.go](../../contract/notifications_test.go).

Authenticated device token lifecycle retained.

### NOTIFY-02

**Single/batch targeted push, household relay registration and failure handling**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [NOTIFY-01](Get-Services-Features.md#notify-01).

Go evidence or governing decision: [internal/modules/notifications/push.go](../../internal/modules/notifications/push.go), [contract/notifications_relay_test.go](../../contract/notifications_relay_test.go).

Relay remains an external dependency; contract uses a fake relay.

### NOTIFY-03

**Inbox create/update/list/detail/read/delete, unread counts and bulk operations**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04), [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/modules/notifications/inbox.go](../../internal/modules/notifications/inbox.go), [contract/notifications_test.go](../../contract/notifications_test.go).

Household/user visibility enforced.

### NOTIFY-04

**Admin delivery statistics and inbox retention cleanup**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [NOTIFY-03](Get-Services-Features.md#notify-03), [PLAT-02](Get-Foundations-Features.md#plat-02).

Go evidence or governing decision: [internal/modules/notifications/notifications.go](../../internal/modules/notifications/notifications.go).

Statistics and cleanup retained.

### NOTIFY-05

**Inbox image/video attachments, downloads and lifecycle cleanup**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [NOTIFY-03](Get-Services-Features.md#notify-03), [PLAT-03](Get-Foundations-Features.md#plat-03).

Go evidence or governing decision: [internal/modules/notifications/attachments.go](../../internal/modules/notifications/attachments.go), [contract/notifications_attachments_test.go](../../contract/notifications_attachments_test.go).

Server feature exists; package-side upload support is tracked separately.

## jarvis-llm-proxy-api

Current disposition: **consolidate**.

Legacy evidence: [api](https://github.com/alexberardi/jarvis-llm-proxy-api/tree/12c0235c8f40756c35c138a5c63c1ab543b24451/api), [backends](https://github.com/alexberardi/jarvis-llm-proxy-api/tree/12c0235c8f40756c35c138a5c63c1ab543b24451/backends), [services](https://github.com/alexberardi/jarvis-llm-proxy-api/tree/12c0235c8f40756c35c138a5c63c1ab543b24451/services).

### LLM-01

**Non-streaming chat and request validation**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-05](Get-Foundations-Features.md#auth-05), [SET-01](Get-Foundations-Features.md#set-01).

Go evidence or governing decision: [internal/modules/llm/http.go](../../internal/modules/llm/http.go), [contract/llm_test.go](../../contract/llm_test.go).

Legacy chat response shape retained.

### LLM-02

**Streaming chat, cancellation and OpenAI-compatible stream option**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-01](Get-Services-Features.md#llm-01).

Go evidence or governing decision: [internal/modules/llm/openai_stream.go](../../internal/modules/llm/openai_stream.go), [internal/modules/llm/http.go](../../internal/modules/llm/http.go), [contract/llm_test.go](../../contract/llm_test.go).

Legacy and standard streaming formats are distinct supported paths.

### LLM-03

**Live/background model labels and admission control**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-01](Get-Services-Features.md#llm-01).

Go evidence or governing decision: [internal/modules/llm/service.go](../../internal/modules/llm/service.go), [internal/modules/llm/models](../../internal/modules/llm/models).

Labels route to local or remote engines.

### LLM-04

**Remote OpenAI-compatible and Anthropic backends**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-03](Get-Services-Features.md#llm-03).

Go evidence or governing decision: [internal/modules/llm/client.go](../../internal/modules/llm/client.go), [internal/modules/llm/anthropic.go](../../internal/modules/llm/anthropic.go).

Remote provider credentials and message conversion retained.

### LLM-05

**Local GGUF inference and supervised llama-server lifecycle**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-03](Get-Services-Features.md#llm-03).

Go evidence or governing decision: [internal/modules/llm/engine](../../internal/modules/llm/engine), [internal/platform/engines](../../internal/platform/engines).

Out-of-process engine replaces in-process llama-cpp bindings.

### LLM-06

**Model catalogue, download/remove, hardware fit and label assignment**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-05](Get-Services-Features.md#llm-05), [PLAT-02](Get-Foundations-Features.md#plat-02).

Go evidence or governing decision: [internal/modules/llm/models](../../internal/modules/llm/models).

Model manager replaces ad hoc swap and install scripts.

### LLM-07

**Text embeddings and model enumeration**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-01](Get-Services-Features.md#llm-01).

Go evidence or governing decision: [internal/modules/llm/http.go](../../internal/modules/llm/http.go), [contract/llm_test.go](../../contract/llm_test.go).

Embedding API retained; compatible embedding engine must be configured.

### LLM-08

**Queued inference jobs, results and completion callbacks**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-01](Get-Services-Features.md#llm-01), [PLAT-02](Get-Foundations-Features.md#plat-02).

Go evidence or governing decision: [internal/modules/llm/jobs.go](../../internal/modules/llm/jobs.go), [contract/llm_test.go](../../contract/llm_test.go).

Internal completions replace service-to-service callbacks where appropriate.

### LLM-09

**JSON/schema handling, repair and message folding**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-01](Get-Services-Features.md#llm-01).

Go evidence or governing decision: [internal/modules/llm/jsonmode](../../internal/modules/llm/jsonmode), [internal/modules/llm/fold.go](../../internal/modules/llm/fold.go).

Structured-output processing retained with model-specific constraints.

### LLM-10

**Image/vision chat payloads and capability checks**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-01](Get-Services-Features.md#llm-01).

Go evidence or governing decision: [internal/modules/llm/types.go](../../internal/modules/llm/types.go), [internal/modules/cc/chat_images.go](../../internal/modules/cc/chat_images.go).

A compatible vision model is still required.

### LLM-11

**MLX inference on Apple Silicon**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [LLM-06](Get-Services-Features.md#llm-06).

Go evidence or governing decision: [internal/modules/llm/engine](../../internal/modules/llm/engine), [docs/STATUS.md](../../docs/STATUS.md).

Engine and folder-model support built; documented live Mac validation remains open.

### LLM-12

**vLLM inference on Linux/NVIDIA**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [LLM-06](Get-Services-Features.md#llm-06).

Go evidence or governing decision: [internal/modules/llm/engine](../../internal/modules/llm/engine), [docs/STATUS.md](../../docs/STATUS.md).

Supervised engine built; live GPU parity/performance checks remain open.

### LLM-13

**Transformers inference, quantisation and accelerator selection**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [LLM-06](Get-Services-Features.md#llm-06).

Go evidence or governing decision: [internal/modules/llm/engine](../../internal/modules/llm/engine), [docs/STATUS.md](../../docs/STATUS.md).

Supervised serve engine built; platform package sets and live CPU/MPS checks remain open.

### LLM-14

**Date-key catalogue and date interpretation support**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LLM-01](Get-Services-Features.md#llm-01).

Go evidence or governing decision: [internal/modules/llm/dates](../../internal/modules/llm/dates), [internal/modules/llm/http.go](../../internal/modules/llm/http.go).

Date behaviour also depends on CC time-zone context.

### LLM-15

**LoRA storage/cache, training jobs and adapter deployment**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Explicit cut; retain offline research tooling outside jarvisd if still needed.

### LLM-16

**Experimental pipeline routes and unused model swap/cache endpoints**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Explicit cuts; do not interpret their absence as an accidental omission.

## jarvis-whisper-api

Current disposition: **consolidate**.

Legacy evidence: [app/main.py](https://github.com/alexberardi/jarvis-whisper-api/blob/8254e8775dd969cd20d10d1ca2008e22921eb57c/app/main.py), [app/api/voice_profiles.py](https://github.com/alexberardi/jarvis-whisper-api/blob/8254e8775dd969cd20d10d1ca2008e22921eb57c/app/api/voice_profiles.py), [app](https://github.com/alexberardi/jarvis-whisper-api/tree/8254e8775dd969cd20d10d1ca2008e22921eb57c/app).

### STT-01

**Audio transcription through whisper.cpp**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-05](Get-Foundations-Features.md#auth-05), [LLM-06](Get-Services-Features.md#llm-06).

Go evidence or governing decision: [internal/modules/stt/transcribe.go](../../internal/modules/stt/transcribe.go), [contract/whisper_test.go](../../contract/whisper_test.go).

whisper-server supervised alongside the model manager.

### STT-02

**Voice-profile enrolment, sample list/delete and profile purge**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [STT-01](Get-Services-Features.md#stt-01), [AUTH-04](Get-Foundations-Features.md#auth-04).

Go evidence or governing decision: [internal/modules/stt/speaker.go](../../internal/modules/stt/speaker.go), [contract/whisper_test.go](../../contract/whisper_test.go).

Household member profiles use the new speaker model.

### STT-03

**Speaker verification, matching and quality thresholds**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [STT-02](Get-Services-Features.md#stt-02).

Go evidence or governing decision: [internal/modules/stt/voicematch.go](../../internal/modules/stt/voicematch.go), [internal/modules/stt/dsp.go](../../internal/modules/stt/dsp.go).

ERes2Net/sherpa replaces ECAPA and resemblyzer; legacy voiceprints require re-enrolment.

### STT-04

**Legacy ECAPA/resemblyzer/librosa implementations**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Deliberately replaced algorithms; not a required code port.

## jarvis-tts

Current disposition: **consolidate**.

Legacy evidence: [app/main.py](https://github.com/alexberardi/jarvis-tts/blob/70a96f3fa5323d25acc660c891b5fdefcac6b9d8/app/main.py), [app/services/provider_manager.py](https://github.com/alexberardi/jarvis-tts/blob/70a96f3fa5323d25acc660c891b5fdefcac6b9d8/app/services/provider_manager.py), [app](https://github.com/alexberardi/jarvis-tts/tree/70a96f3fa5323d25acc660c891b5fdefcac6b9d8/app).

### TTS-01

**Kokoro text-to-speech, voice/speed settings and WAV output**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-05](Get-Foundations-Features.md#auth-05), [SET-01](Get-Foundations-Features.md#set-01).

Go evidence or governing decision: [internal/modules/tts/speak.go](../../internal/modules/tts/speak.go), [internal/modules/tts/voices.go](../../internal/modules/tts/voices.go), [contract/tts_test.go](../../contract/tts_test.go).

In-process sherpa synthesis replaces Python provider.

### TTS-02

**Low-latency raw PCM streaming and sentence splitting**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [TTS-01](Get-Services-Features.md#tts-01).

Go evidence or governing decision: [internal/modules/tts/split.go](../../internal/modules/tts/split.go), [internal/modules/tts/speak.go](../../internal/modules/tts/speak.go), [contract/tts_test.go](../../contract/tts_test.go).

Buffered and streaming callers retained.

### TTS-03

**Piper provider and hot-swap fallback**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Explicit cut; Kokoro is the selected engine.

### TTS-04

**Deprecated generate-wake-response endpoint**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: [CC-03](Get-Orchestration-Features.md#cc-03).

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Use CC wake-response capability; legacy TTS shim deliberately omitted.

## jarvis-ocr-service

Current disposition: **consolidate**.

Legacy evidence: [app/main.py](https://github.com/alexberardi/jarvis-ocr-service/blob/2f5b3c0c5e512718204f2a82239cba92ecca9515/app/main.py), [app/provider_manager.py](https://github.com/alexberardi/jarvis-ocr-service/blob/2f5b3c0c5e512718204f2a82239cba92ecca9515/app/provider_manager.py), [app/providers](https://github.com/alexberardi/jarvis-ocr-service/tree/2f5b3c0c5e512718204f2a82239cba92ecca9515/app/providers).

### OCR-01

**Submit single/batch OCR jobs, provider listing and queue/status polling**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PLAT-02](Get-Foundations-Features.md#plat-02), [AUTH-05](Get-Foundations-Features.md#auth-05).

Go evidence or governing decision: [internal/modules/ocr/ocr.go](../../internal/modules/ocr/ocr.go), [internal/modules/ocr/jobs.go](../../internal/modules/ocr/jobs.go), [contract/ocr_test.go](../../contract/ocr_test.go).

Durable Go jobs replace RQ.

### OCR-02

**Image resolution, vision-LLM OCR, Tesseract and Apple Vision**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [OCR-01](Get-Services-Features.md#ocr-01), [LLM-10](Get-Services-Features.md#llm-10), [SCRAPE-02](Get-Foundations-Features.md#scrape-02).

Go evidence or governing decision: [internal/modules/ocr/recognize.go](../../internal/modules/ocr/recognize.go), [internal/modules/ocr/vision.go](../../internal/modules/ocr/vision.go), [internal/modules/ocr/vision_darwin.go](../../internal/modules/ocr/vision_darwin.go).

Provider availability depends on the host; Apple Vision is macOS-only.

### OCR-03

**Provider fallback, result combination, validation and callbacks**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [OCR-02](Get-Services-Features.md#ocr-02).

Go evidence or governing decision: [internal/modules/ocr/pipeline.go](../../internal/modules/ocr/pipeline.go), [contract/ocr_test.go](../../contract/ocr_test.go).

Retained OCR pipeline; recipe hand-off stays queue-based.

### OCR-04

**EasyOCR, PaddleOCR, RapidOCR and dead async-validation modules**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Explicit cuts; assess accuracy against the retained providers before adding alternatives.

## jarvis-recipes-server

Current disposition: **consolidate**.

Legacy evidence: [jarvis_recipes/app/api](https://github.com/alexberardi/jarvis-recipes-server/tree/2f07b15a773c5b52575fedbfb16ba00e49df0d28/jarvis_recipes/app/api), [jarvis_recipes/app/services](https://github.com/alexberardi/jarvis-recipes-server/tree/2f07b15a773c5b52575fedbfb16ba00e49df0d28/jarvis_recipes/app/services).

### RECIPE-01

**Recipe CRUD, search, ingredients, tags and household sharing**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04), [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/modules/recipes/routes_recipes.go](../../internal/modules/recipes/routes_recipes.go), [contract/recipes_test.go](../../contract/recipes_test.go).

Core family recipe data plane retained.

### RECIPE-02

**Ingredient/unit stock reference data and staples**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-01](Get-Services-Features.md#recipe-01).

Go evidence or governing decision: [internal/modules/recipes/stock.go](../../internal/modules/recipes/stock.go), [internal/modules/recipes/shopping_routes.go](../../internal/modules/recipes/shopping_routes.go), [contract/recipes_test.go](../../contract/recipes_test.go).

Food pantry stock is distinct from the Pantry package store.

### RECIPE-03

**URL/HTML import using schema.org, heuristics and LLM fallback**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-01](Get-Services-Features.md#recipe-01), [LLM-01](Get-Services-Features.md#llm-01), [SCRAPE-02](Get-Foundations-Features.md#scrape-02).

Go evidence or governing decision: [internal/modules/recipes/extract](../../internal/modules/recipes/extract), [internal/modules/recipes/ingest.go](../../internal/modules/recipes/ingest.go), [contract/recipes_jobs_test.go](../../contract/recipes_jobs_test.go).

Fetch restrictions and async job polling retained.

### RECIPE-04

**Image OCR import, extraction staging and mailbox commit**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-01](Get-Services-Features.md#recipe-01), [OCR-03](Get-Services-Features.md#ocr-03), [PLAT-03](Get-Foundations-Features.md#plat-03).

Go evidence or governing decision: [internal/modules/recipes/photo_import.go](../../internal/modules/recipes/photo_import.go), [internal/modules/recipes/jobs_routes.go](../../internal/modules/recipes/jobs_routes.go), [contract/recipes_jobs_test.go](../../contract/recipes_jobs_test.go).

OCR queue hand-off and user-scoped staging retained.

### RECIPE-05

**Image storage, media delivery and image chat/refinement**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-04](Get-Services-Features.md#recipe-04), [LLM-10](Get-Services-Features.md#llm-10).

Go evidence or governing decision: [internal/modules/recipes/media.go](../../internal/modules/recipes/media.go), [internal/modules/recipes/photo_chat.go](../../internal/modules/recipes/photo_chat.go).

Validate source recipe media during import.

### RECIPE-06

**AI meal-plan generation with per-day preferences**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-01](Get-Services-Features.md#recipe-01), [LLM-08](Get-Services-Features.md#llm-08).

Go evidence or governing decision: [internal/modules/recipes/mealplan.go](../../internal/modules/recipes/mealplan.go), [contract/recipes_jobs_test.go](../../contract/recipes_jobs_test.go).

Real-model quality remains a separate acceptance check.

### RECIPE-07

**Random meal plans, reroll, save, retrieve, move and delete planned meals**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-01](Get-Services-Features.md#recipe-01).

Go evidence or governing decision: [internal/modules/recipes/random.go](../../internal/modules/recipes/random.go), [internal/modules/recipes/planner.go](../../internal/modules/recipes/planner.go), [contract/recipes_test.go](../../contract/recipes_test.go).

Deterministic planner operations retained.

### RECIPE-08

**Shopping aggregation, quantity/unit normalisation and staples inclusion**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-02](Get-Services-Features.md#recipe-02), [RECIPE-07](Get-Services-Features.md#recipe-07).

Go evidence or governing decision: [internal/modules/recipes/shopping](../../internal/modules/recipes/shopping), [internal/modules/recipes/quantity](../../internal/modules/recipes/quantity).

Shopping output derived from selected recipes and pantry state.

### RECIPE-09

**Grocery catalogue, SKU mappings and queued matching**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-08](Get-Services-Features.md#recipe-08), [LLM-08](Get-Services-Features.md#llm-08).

Go evidence or governing decision: [internal/modules/recipes/grocery.go](../../internal/modules/recipes/grocery.go), [internal/modules/recipes/grocery_match.go](../../internal/modules/recipes/grocery_match.go), [contract/recipes_jobs_test.go](../../contract/recipes_jobs_test.go).

Grocery matching API retained.

### RECIPE-10

**Legacy recipe/media import with user remapping**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-01](Get-Services-Features.md#recipe-01), [PLAT-03](Get-Foundations-Features.md#plat-03).

Go evidence or governing decision: [internal/modules/recipes/legacy_import.go](../../internal/modules/recipes/legacy_import.go), [docs/install/legacy-import.md](../../docs/install/legacy-import.md).

Separate from import-legacy; rehearse owner mapping and missing media.

### RECIPE-11

**Static core-recipe lookup compatibility**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [RECIPE-01](Get-Services-Features.md#recipe-01).

Go evidence or governing decision: [internal/modules/recipes/routes_recipes.go](../../internal/modules/recipes/routes_recipes.go).

Matches the legacy always-404 core lookup behaviour; not a populated recipe catalogue.
