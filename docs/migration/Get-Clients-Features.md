# Clients feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-node-mobile

Current disposition: **retain client; assess server support**.

Legacy evidence: [src/screens](https://github.com/alexberardi/jarvis-node-mobile/tree/41d72ce0657d8add8043ac8e4438405362175bac/src/screens), [src/api](https://github.com/alexberardi/jarvis-node-mobile/tree/41d72ce0657d8add8043ac8e4438405362175bac/src/api), [src/services](https://github.com/alexberardi/jarvis-node-mobile/tree/41d72ce0657d8add8043ac8e4438405362175bac/src/services), [native](https://github.com/alexberardi/jarvis-node-mobile/tree/41d72ce0657d8add8043ac8e4438405362175bac/native).

### MOB-01

**Server discovery, login/refresh, registration, forced password change and secure token storage**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [CONF-02](Get-Foundations-Features.md#conf-02), [AUTH-03](Get-Foundations-Features.md#auth-03).

Go evidence or governing decision: [contract/auth_test.go](../../contract/auth_test.go), [e2e/coverage.json](../../e2e/coverage.json).

Backend supported; native app/keychain stays outside this repository.

### MOB-02

**Voice/text chat, photo input, SSE replies and recent command detail**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [CC-11](Get-Orchestration-Features.md#cc-11), [CC-12](Get-Orchestration-Features.md#cc-12), [CC-13](Get-Orchestration-Features.md#cc-13), [CC-33](Get-Orchestration-Features.md#cc-33).

Go evidence or governing decision: [e2e](../../e2e), [internal/modules/cc/mobile_chat.go](../../internal/modules/cc/mobile_chat.go).

Server support implemented; external Expo screens are not absorbed.

### MOB-03

**QR/Wi-Fi node provisioning, setup progress and adding a computer**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [NODE-05](Get-Nodes-Features.md#node-05), [DESK-01](Get-Nodes-Features.md#desk-01).

Go evidence or governing decision: [cmd/jarvis-agent/provision.go](../../cmd/jarvis-agent/provision.go), [internal/modules/cc/pairing.go](../../internal/modules/cc/pairing.go).

Keep provisioning and pairing acceptance per platform.

### MOB-04

**Nodes, updates, reset, hardware diagnostics and encrypted settings sync**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [CC-28](Get-Orchestration-Features.md#cc-28), [NODE-06](Get-Nodes-Features.md#node-06), [NODE-24](Get-Nodes-Features.md#node-24).

Go evidence or governing decision: [internal/modules/cc](../../internal/modules/cc), [cmd/jarvis-agent](../../cmd/jarvis-agent).

Server surface exists; Pi gates remain open.

### MOB-05

**Households, rooms, Home Assistant discovery/import and device control**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [CC-25](Get-Orchestration-Features.md#cc-25), [CC-26](Get-Orchestration-Features.md#cc-26), [NODE-16](Get-Nodes-Features.md#node-16).

Go evidence or governing decision: [internal/modules/cc/smarthome.go](../../internal/modules/cc/smarthome.go), [e2e](../../e2e).

External mobile app remains the main management surface for several workflows.

### MOB-06

**Camera view and live stream playback**

Status: **not started**. Disposition: **deferred retained capability**.

Prerequisites: [CC-27](Get-Orchestration-Features.md#cc-27).

Go evidence or governing decision: [internal/modules/cc/smarthome.go](../../internal/modules/cc/smarthome.go).

Go streams return unavailable; listing imported cameras is not playback.

### MOB-07

**Inbox, rich cards, callbacks, interactive lists, push and media playback**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [NOTIFY-05](Get-Services-Features.md#notify-05), [CC-23](Get-Orchestration-Features.md#cc-23).

Go evidence or governing decision: [internal/modules/notifications](../../internal/modules/notifications), [internal/modules/cc/callbacks.go](../../internal/modules/cc/callbacks.go), [e2e](../../e2e).

Backend retained; native delivery/deep links require mobile validation.

### MOB-08

**Memories, routines, schedules and signal-automation settings**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [CC-14](Get-Orchestration-Features.md#cc-14), [CC-18](Get-Orchestration-Features.md#cc-18), [CC-19](Get-Orchestration-Features.md#cc-19), [CC-22](Get-Orchestration-Features.md#cc-22).

Go evidence or governing decision: [internal/modules/cc](../../internal/modules/cc), [e2e](../../e2e).

Server surfaces implemented; native UI remains external.

### MOB-09

**Command data browser, record create/edit and package secrets/OAuth**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [CC-31](Get-Orchestration-Features.md#cc-31), [CC-29](Get-Orchestration-Features.md#cc-29), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/modules/cc/commanddata.go](../../internal/modules/cc/commanddata.go).

Package/node schemas still govern edit forms.

### MOB-10

**Pantry browse/detail, node selection, install progress and temporary tests**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [CC-30](Get-Orchestration-Features.md#cc-30), [CC-38](Get-Orchestration-Features.md#cc-38), [NODE-07](Get-Nodes-Features.md#node-07).

Go evidence or governing decision: [internal/modules/cc/packages.go](../../internal/modules/cc/packages.go).

Keep Pantry external; verify package target restrictions on desktop versus Pi.

### MOB-11

**Phonebook, call context and voice-profile management**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [CC-34](Get-Orchestration-Features.md#cc-34), [CC-33](Get-Orchestration-Features.md#cc-33).

Go evidence or governing decision: [internal/modules/cc/phone](../../internal/modules/cc/phone), [internal/modules/cc/voice_profiles.go](../../internal/modules/cc/voice_profiles.go).

Native app remains outside Go repo.

### MOB-12

**Presence/location events and notification routing**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [CC-24](Get-Orchestration-Features.md#cc-24).

Go evidence or governing decision: [internal/modules/cc/signals.go](../../internal/modules/cc/signals.go).

Native location permissions/background tasks are not consolidated.

### MOB-13

**iOS Action Button, Control Centre/Shortcuts, universal links and key import**

Status: **not started**. Disposition: **retain client; assess server support**.

Prerequisites: [MOB-01](Get-Clients-Features.md#mob-01).

Go evidence or governing decision: [docs/dev/mobile-shots.md](../../docs/dev/mobile-shots.md).

Native app/site association components remain external; no equivalent native client is in this repository.

### MOB-14

**Adapter proposals, deployment detail and rollback screens**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Cut backend feature; remove/hide external mobile screens as a client retirement task.

## jarvis-recipes-mobile

Current disposition: **retain client; assess server support**.

Legacy evidence: [src/screens](https://github.com/alexberardi/jarvis-recipes-mobile/tree/564bb013e2e68ac3ea59d76a1de8ef2b7c53e861/src/screens), [src/api](https://github.com/alexberardi/jarvis-recipes-mobile/tree/564bb013e2e68ac3ea59d76a1de8ef2b7c53e861/src/api).

### RMOB-01

**Login, registration, account settings and recipe browsing/editing**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [AUTH-03](Get-Foundations-Features.md#auth-03), [RECIPE-01](Get-Services-Features.md#recipe-01).

Go evidence or governing decision: [internal/modules/recipes](../../internal/modules/recipes), [contract/recipes_test.go](../../contract/recipes_test.go).

Backend migrated; Expo client remains in its own repository.

### RMOB-02

**URL/browser extraction, multi-photo import, progress and mailbox review**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [RECIPE-03](Get-Services-Features.md#recipe-03), [RECIPE-04](Get-Services-Features.md#recipe-04).

Go evidence or governing decision: [internal/modules/recipes](../../internal/modules/recipes), [contract/recipes_jobs_test.go](../../contract/recipes_jobs_test.go).

Native share/browser/photo UI remains external.

### RMOB-03

**Meal-plan date/day setup, quick plan, progress, results and saved plans**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [RECIPE-06](Get-Services-Features.md#recipe-06), [RECIPE-07](Get-Services-Features.md#recipe-07).

Go evidence or governing decision: [internal/modules/recipes](../../internal/modules/recipes).

Backend available; preserve client workflow acceptance.

### RMOB-04

**Shopping list, product picker, groceries and SKU selections**

Status: **in progress**. Disposition: **retain client; assess server support**.

Prerequisites: [RECIPE-09](Get-Services-Features.md#recipe-09).

Go evidence or governing decision: [internal/modules/recipes/grocery.go](../../internal/modules/recipes/grocery.go).

Backend available; no recipe UI has been absorbed into admin.

## jarvis-recipes

Current disposition: **legacy prototype; decide retirement**.

Legacy evidence: [backend/app](https://github.com/alexberardi/jarvis-recipes/tree/63462bd96cca737173bdab404c78402809f5e131/backend/app), [frontend/src](https://github.com/alexberardi/jarvis-recipes/tree/63462bd96cca737173bdab404c78402809f5e131/frontend/src).

### RPROTO-01

**Prototype password login/JWT and current-user API**

Status: **complete**. Disposition: **legacy prototype; decide retirement**.

Prerequisites: [AUTH-01](Get-Foundations-Features.md#auth-01).

Go evidence or governing decision: [internal/modules/auth](../../internal/modules/auth).

Capability superseded by shared auth; independent prototype token format is not preserved.

### RPROTO-02

**Prototype Next.js login page and themed application shell**

Status: **in progress**. Disposition: **legacy prototype; decide retirement**.

Prerequisites: [RPROTO-01](Get-Clients-Features.md#rproto-01).

Go evidence or governing decision: [web/admin/src/pages/LoginPage.tsx](../../web/admin/src/pages/LoginPage.tsx).

Generic sign-in capability replaced; this checkout contains no complete recipe UI to port.

## jarvis-web

Current disposition: **retire standalone client**.

Legacy evidence: [app](https://github.com/alexberardi/jarvis-web/tree/f37c6de912638014b439bc9c4b735da528db653b/app), [components](https://github.com/alexberardi/jarvis-web/tree/f37c6de912638014b439bc9c4b735da528db653b/components), [lib](https://github.com/alexberardi/jarvis-web/tree/f37c6de912638014b439bc9c4b735da528db653b/lib).

### WEB-01

**Browser sign-in, streamed chat and conversation history**

Status: **complete**. Disposition: **retire standalone client**.

Prerequisites: [CC-11](Get-Orchestration-Features.md#cc-11).

Go evidence or governing decision: [web/admin/src/pages/ChatPage.tsx](../../web/admin/src/pages/ChatPage.tsx), [web/admin/src/pages/LoginPage.tsx](../../web/admin/src/pages/LoginPage.tsx).

Standalone Next.js app retired; embedded React app provides browser chat.

### WEB-02

**Inbox list/detail, actions and account settings**

Status: **complete**. Disposition: **retire standalone client**.

Prerequisites: [NOTIFY-03](Get-Services-Features.md#notify-03), [CC-23](Get-Orchestration-Features.md#cc-23).

Go evidence or governing decision: [web/admin/src/pages/InboxPage.tsx](../../web/admin/src/pages/InboxPage.tsx), [web/admin/src/pages/InboxItemPage.tsx](../../web/admin/src/pages/InboxItemPage.tsx), [web/admin/src/pages/AccountPage.tsx](../../web/admin/src/pages/AccountPage.tsx).

Household member pages exist in embedded web app.

### WEB-03

**Package store and node/add-on management**

Status: **in progress**. Disposition: **retire standalone client**.

Prerequisites: [CC-30](Get-Orchestration-Features.md#cc-30), [CC-28](Get-Orchestration-Features.md#cc-28).

Go evidence or governing decision: [web/admin/src/pages/AddonsPage.tsx](../../web/admin/src/pages/AddonsPage.tsx), [web/admin/src/pages/NodesPage.tsx](../../web/admin/src/pages/NodesPage.tsx).

Replacement pages exist; compare package-management journeys before retiring old bookmarks.

### WEB-04

**Browser registration, household management and smart-home device pages**

Status: **not started**. Disposition: **retire standalone client**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04), [CC-25](Get-Orchestration-Features.md#cc-25).

Go evidence or governing decision: [docs/webapp/01-member-web-app-plan.md](../../docs/webapp/01-member-web-app-plan.md).

Standalone README directs these workflows to mobile; do not call browser parity complete.
