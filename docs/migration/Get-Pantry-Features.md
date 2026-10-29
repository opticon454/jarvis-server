# Pantry feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-pantry

Current disposition: **excluded by user**.

Legacy evidence: [app/api](https://github.com/alexberardi/jarvis-pantry/tree/3e015541c8f7758a57fb5792e8c3f59d53b1ed3a/app/api), [app/services](https://github.com/alexberardi/jarvis-pantry/tree/3e015541c8f7758a57fb5792e8c3f59d53b1ed3a/app/services), [config](https://github.com/alexberardi/jarvis-pantry/tree/3e015541c8f7758a57fb5792e8c3f59d53b1ed3a/config), [README.md](https://github.com/alexberardi/jarvis-pantry/blob/3e015541c8f7758a57fb5792e8c3f59d53b1ed3a/README.md).

### PANTRY-01

**Public catalogue, search/category browse, detail, versions and package downloads**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [internal/modules/cc/packages.go](../../internal/modules/cc/packages.go).

Keep external; Go only consumes the store.

### PANTRY-02

**GitHub OAuth, author submissions, ownership and package management**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-01](Get-Pantry-Features.md#pantry-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep external authoring boundary.

### PANTRY-03

**Community ratings/reviews and shared routines catalogue**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-01](Get-Pantry-Features.md#pantry-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep external; routine consumption can use server routine features.

### PANTRY-04

**Validation pipeline: manifest/AST, security review, dependency locks and container tests**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-02](Get-Pantry-Features.md#pantry-02).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep external; runner is an independent isolated execution service.

### PANTRY-05

**Forge generation, conversation/drafts, model choice, BYOK and publishing**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-04](Get-Pantry-Features.md#pantry-04).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep external; source availability does not assert hosted Forge is currently enabled.

### PANTRY-06

**Private/local packages, local LLM settings and upstream catalogue mirror**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-01](Get-Pantry-Features.md#pantry-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep external; configured Pantry base URL must work from server and nodes.

### PANTRY-07

**Admin bulk management, API keys and install/post-install allowlists**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-04](Get-Pantry-Features.md#pantry-04).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep external; node installer validates the permitted operations separately.

## jarvis-pantry-web

Current disposition: **excluded by user**.

Legacy evidence: [src](https://github.com/alexberardi/jarvis-pantry-web/tree/cabd7693ca1426e4f679f349a7a8155c544dfb57/src), [README.md](https://github.com/alexberardi/jarvis-pantry-web/blob/cabd7693ca1426e4f679f349a7a8155c544dfb57/README.md).

### PWEB-01

**Catalogue search/filter/sort and package detail/reviews/security/install pages**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-01](Get-Pantry-Features.md#pantry-01), [PANTRY-03](Get-Pantry-Features.md#pantry-03).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep standalone Pantry frontend.

### PWEB-02

**Repository submission picker and live validation-pipeline progress**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-04](Get-Pantry-Features.md#pantry-04).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep outside consolidation.

### PWEB-03

**Forge chat/editor, model/BYOK selection, AST validation and GitHub publication**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-05](Get-Pantry-Features.md#pantry-05).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep outside consolidation.

### PWEB-04

**Capability-gated local model settings, private-package pages and API proxy**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-06](Get-Pantry-Features.md#pantry-06).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep self-hosted/private Pantry support documented.

## jarvis-pantry-runner

Current disposition: **excluded by user**.

Legacy evidence: [runner.py](https://github.com/alexberardi/jarvis-pantry-runner/blob/3946b6bbc26b635fa7ccdcc87dffe076d129faf6/runner.py), [harness.py](https://github.com/alexberardi/jarvis-pantry-runner/blob/3946b6bbc26b635fa7ccdcc87dffe076d129faf6/harness.py), [.github/workflows](https://github.com/alexberardi/jarvis-pantry-runner/tree/3946b6bbc26b635fa7ccdcc87dffe076d129faf6/.github/workflows), [README.md](https://github.com/alexberardi/jarvis-pantry-runner/blob/3946b6bbc26b635fa7ccdcc87dffe076d129faf6/README.md).

### PRUN-01

**Dispatch isolated package SDK tests from Pantry with per-submission nonce**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PANTRY-04](Get-Pantry-Features.md#pantry-04).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep external workflow runner.

### PRUN-02

**Two-job sandbox/callback separation, signed HMAC results and result validation**

Status: **not applicable**. Disposition: **excluded by user**.

Prerequisites: [PRUN-01](Get-Pantry-Features.md#prun-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

Keep callback key outside untrusted test execution; not a jarvisd job.
