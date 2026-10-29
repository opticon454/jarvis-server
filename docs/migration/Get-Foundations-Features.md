# Foundations feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-data-services

Current disposition: **consolidate**.

Legacy evidence: [README.md](https://github.com/alexberardi/jarvis-data-services/blob/32dcdabfa5fe75ea0b4520ec6da9a173f48786c8/README.md), [docker-compose.yml](https://github.com/alexberardi/jarvis-data-services/blob/32dcdabfa5fe75ea0b4520ec6da9a173f48786c8/docker-compose.yml).

### PLAT-01

**Persistent relational data, schema migrations and transactions**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [internal/platform/db](../../internal/platform/db).

PostgreSQL service replaced by embedded SQLite; review migration transforms before importing data.

### PLAT-02

**Durable background jobs, retries and restart recovery**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/platform/queue](../../internal/platform/queue).

Redis queue replaced by SQLite-backed jobs.

### PLAT-03

**Object and media storage**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/platform/blob](../../internal/platform/blob).

Filesystem backend replaces MinIO; storage capability is retained, not the MinIO administration UI.

### PLAT-04

**Database administration UIs and create/drop/rename shell utilities**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Retired infrastructure: pgAdmin and Redis Commander have no Go equivalents; exclude from required migration scope.

## jarvis-auth

Current disposition: **consolidate**.

Legacy evidence: [jarvis_auth/app/api](https://github.com/alexberardi/jarvis-auth/tree/3fa3c1e95a91929a21b16029097770c99b58f486/jarvis_auth/app/api), [jarvis_auth/app/services](https://github.com/alexberardi/jarvis-auth/tree/3fa3c1e95a91929a21b16029097770c99b58f486/jarvis_auth/app/services), [jarvis_auth/app/core](https://github.com/alexberardi/jarvis-auth/tree/3fa3c1e95a91929a21b16029097770c99b58f486/jarvis_auth/app/core).

### AUTH-01

**Registration, login, access tokens and current-user profile**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/modules/auth/auth.go](../../internal/modules/auth/auth.go), [contract/auth_test.go](../../contract/auth_test.go).

Legacy auth contract covered; registration policy is now controlled by server setup.

### AUTH-02

**Refresh rotation, logout and token revocation**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-01](Get-Foundations-Features.md#auth-01).

Go evidence or governing decision: [internal/modules/auth/tokens.go](../../internal/modules/auth/tokens.go), [contract/auth_test.go](../../contract/auth_test.go).

Token family revocation and refresh handling are implemented.

### AUTH-03

**Password changes, forced reset and account deletion**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-01](Get-Foundations-Features.md#auth-01).

Go evidence or governing decision: [internal/modules/auth/deletion.go](../../internal/modules/auth/deletion.go), [contract/auth_account_test.go](../../contract/auth_account_test.go).

Deletion coordinates module data cleanup; preserve cross-module purge checks.

### AUTH-04

**Households, memberships, roles and invitations**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-01](Get-Foundations-Features.md#auth-01).

Go evidence or governing decision: [internal/modules/auth/households.go](../../internal/modules/auth/households.go), [contract/auth_households_test.go](../../contract/auth_households_test.go), [contract/auth_invites_test.go](../../contract/auth_invites_test.go).

Household boundaries and invite flows retained.

### AUTH-05

**App credentials, key rotation and service grants**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-01](Get-Foundations-Features.md#auth-01).

Go evidence or governing decision: [internal/modules/auth/admin.go](../../internal/modules/auth/admin.go), [internal/modules/auth/apps_inprocess.go](../../internal/modules/auth/apps_inprocess.go), [contract/auth_admin_test.go](../../contract/auth_admin_test.go).

Remote compatibility routes and in-process authority coexist.

### AUTH-06

**Node credentials, household ownership and validation**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04).

Go evidence or governing decision: [internal/modules/auth/nodes.go](../../internal/modules/auth/nodes.go), [internal/modules/auth/nodes_inprocess.go](../../internal/modules/auth/nodes_inprocess.go), [contract/auth_internal_test.go](../../contract/auth_internal_test.go).

Keep old node keys through legacy import.

### AUTH-07

**Superuser administration and cross-household views**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-04](Get-Foundations-Features.md#auth-04), [AUTH-05](Get-Foundations-Features.md#auth-05).

Go evidence or governing decision: [internal/modules/auth/admin.go](../../internal/modules/auth/admin.go), [contract/auth_superuser_test.go](../../contract/auth_superuser_test.go).

User, app and node administration retained.

### AUTH-08

**JWT signing keys, local validation and request guards**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-01](Get-Foundations-Features.md#auth-01).

Go evidence or governing decision: [internal/modules/auth/keys.go](../../internal/modules/auth/keys.go), [internal/platform/authn](../../internal/platform/authn).

RS256 signing and legacy HS256 verification replace distributed shared-secret-only operation.

## jarvis-auth-client

Current disposition: **consolidate**.

Legacy evidence: [jarvis_auth_client](https://github.com/alexberardi/jarvis-auth-client/tree/7e1a6769bb4d8d32759215aab288b0d62a7ba958/jarvis_auth_client).

### AUTHLIB-01

**Superuser JWT and app authentication helpers**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTH-05](Get-Foundations-Features.md#auth-05), [AUTH-08](Get-Foundations-Features.md#auth-08).

Go evidence or governing decision: [internal/platform/authn](../../internal/platform/authn).

Replaced by shared Go principals and Authority, not a Python API port.

### AUTHLIB-02

**Household, node and user context forwarding**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [AUTHLIB-01](Get-Foundations-Features.md#authlib-01).

Go evidence or governing decision: [internal/platform/authn](../../internal/platform/authn), [internal/modules/cc/authz.go](../../internal/modules/cc/authz.go).

Boundary checks remain on HTTP routes; internal calls use typed context.

## jarvis-config-service

Current disposition: **consolidate**.

Legacy evidence: [app/routes](https://github.com/alexberardi/jarvis-config-service/tree/6bbbd3e3a1b03f89fc2c132bdbe6714a38d8e3a3/app/routes), [app/main.py](https://github.com/alexberardi/jarvis-config-service/blob/6bbbd3e3a1b03f89fc2c132bdbe6714a38d8e3a3/app/main.py).

### CONF-01

**Service directory, URL lookup and legacy listener discovery**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/modules/config/config.go](../../internal/modules/config/config.go), [contract/config_test.go](../../contract/config_test.go).

Module listeners self-register; external rows remain editable.

### CONF-02

**mDNS advertisement and public/LAN address selection**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CONF-01](Get-Foundations-Features.md#conf-01).

Go evidence or governing decision: [internal/platform/mdns](../../internal/platform/mdns), [internal/modules/config/public.go](../../internal/modules/config/public.go).

Mobile discovery retains the legacy service identity.

### CONF-03

**Health probes and service registration CRUD**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CONF-01](Get-Foundations-Features.md#conf-01), [AUTH-05](Get-Foundations-Features.md#auth-05).

Go evidence or governing decision: [internal/modules/config/config.go](../../internal/modules/config/config.go), [contract/config_test.go](../../contract/config_test.go).

Core service list and health routes are retained.

### CONF-04

**Aggregated runtime settings gateway**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [SET-01](Get-Foundations-Features.md#set-01), [AUTH-07](Get-Foundations-Features.md#auth-07).

Go evidence or governing decision: [internal/modules/admin/bff.go](../../internal/modules/admin/bff.go).

Admin aggregation replaces legacy HTTP fan-out; endpoint paths are not identical.

### CONF-05

**Legacy registry register/rotate-key/probe orchestration surface**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [CONF-03](Get-Foundations-Features.md#conf-03), [AUTH-05](Get-Foundations-Features.md#auth-05).

Go evidence or governing decision: [internal/modules/admin/connections.go](../../internal/modules/admin/connections.go), [internal/modules/config/config.go](../../internal/modules/config/config.go).

Core registration exists; no one-for-one config service-registration router port found. Decide whether remaining legacy consumers require it.

## jarvis-config-client

Current disposition: **consolidate**.

Legacy evidence: [jarvis_config_client/client.py](https://github.com/alexberardi/jarvis-config-client/blob/40a21f90a5c1f580ac785156134f635916209381/jarvis_config_client/client.py), [jarvis_config_client/discovery.py](https://github.com/alexberardi/jarvis-config-client/blob/40a21f90a5c1f580ac785156134f635916209381/jarvis_config_client/discovery.py).

### CONFLIB-01

**Service resolution, refresh and mandatory-service lookup**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CONF-01](Get-Foundations-Features.md#conf-01).

Go evidence or governing decision: [internal/modules/config/inprocess.go](../../internal/modules/config/inprocess.go), [internal/agent/node](../../internal/agent/node).

Server modules use in-process services; agents resolve server endpoints.

### CONFLIB-02

**Automatic config discovery with configured URL fallback**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CONF-02](Get-Foundations-Features.md#conf-02).

Go evidence or governing decision: [internal/platform/mdns](../../internal/platform/mdns), [internal/agent/firstrun](../../internal/agent/firstrun).

Discovery is available; Python helper signatures are not a migration target.

## jarvis-settings-client

Current disposition: **consolidate**.

Legacy evidence: [jarvis_settings_client/service.py](https://github.com/alexberardi/jarvis-settings-client/blob/402674d1a07891fc6b096ab48f370c371558e4cd/jarvis_settings_client/service.py), [jarvis_settings_client/routes.py](https://github.com/alexberardi/jarvis-settings-client/blob/402674d1a07891fc6b096ab48f370c371558e4cd/jarvis_settings_client/routes.py).

### SET-01

**Typed settings, defaults and user/node/household/system cascade**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PLAT-01](Get-Foundations-Features.md#plat-01).

Go evidence or governing decision: [internal/platform/settings](../../internal/platform/settings), [contract/settings_test.go](../../contract/settings_test.go).

Go cascade retained; unrestricted environment fallback is narrowed to configured secret/URL definitions, per M3.

### SET-02

**Settings CRUD, categories, environment sync and cache invalidation**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [SET-01](Get-Foundations-Features.md#set-01), [AUTH-08](Get-Foundations-Features.md#auth-08).

Go evidence or governing decision: [internal/platform/settings](../../internal/platform/settings), [contract/settings_test.go](../../contract/settings_test.go).

Legacy mounted settings surface retained where required.

## jarvis-settings-server

Current disposition: **retire**.

Legacy evidence: [README.md](https://github.com/alexberardi/jarvis-settings-server/blob/a116e32e0b3b9eb349c0e5e7bef317f0d1cc921a/README.md), [app/routes/settings.py](https://github.com/alexberardi/jarvis-settings-server/blob/a116e32e0b3b9eb349c0e5e7bef317f0d1cc921a/app/routes/settings.py).

### SETOLD-01

**Cross-service settings aggregation and update forwarding**

Status: **complete**. Disposition: **retire**.

Prerequisites: [CONF-04](Get-Foundations-Features.md#conf-04).

Go evidence or governing decision: [internal/modules/admin/bff.go](../../internal/modules/admin/bff.go).

Capability replaced by admin aggregator; separate service deliberately retired.

### SETOLD-02

**Standalone category-filter and service-URL debug endpoints**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: [SETOLD-01](Get-Foundations-Features.md#setold-01).

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Cut service; no standalone compatibility port is required by the existing plan.

## jarvis-logs

Current disposition: **consolidate**.

Legacy evidence: [app/routes/logs.py](https://github.com/alexberardi/jarvis-logs/blob/63229cdfeb8b21664a84cd34e5b0a5ae0aebcc8f/app/routes/logs.py), [app/routes/node_logs.py](https://github.com/alexberardi/jarvis-logs/blob/63229cdfeb8b21664a84cd34e5b0a5ae0aebcc8f/app/routes/node_logs.py).

### LOG-01

**Single/batch log ingestion from services and nodes**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [PLAT-01](Get-Foundations-Features.md#plat-01), [AUTH-05](Get-Foundations-Features.md#auth-05), [AUTH-06](Get-Foundations-Features.md#auth-06).

Go evidence or governing decision: [internal/modules/logs/logs.go](../../internal/modules/logs/logs.go), [contract/logs_test.go](../../contract/logs_test.go).

SQLite sink replaces Loki.

### LOG-02

**Filtered log query, service enumeration and live tail**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LOG-01](Get-Foundations-Features.md#log-01).

Go evidence or governing decision: [internal/modules/logs/logs.go](../../internal/modules/logs/logs.go), [contract/logs_test.go](../../contract/logs_test.go).

HTTP query and streaming preserved.

### LOG-03

**Log retention and scheduled cleanup**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LOG-01](Get-Foundations-Features.md#log-01), [PLAT-02](Get-Foundations-Features.md#plat-02).

Go evidence or governing decision: [internal/modules/logs/logs.go](../../internal/modules/logs/logs.go).

Retention implemented by a scheduled purge.

### LOG-04

**Standalone Loki/Grafana operations**

Status: **not started**. Disposition: **retire / explicit cut**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [docs/PLAN.md](../../docs/PLAN.md).

Replaced infrastructure, not a required Go port; embedded Logs page is the user-facing successor.

## jarvis-log-client

Current disposition: **consolidate**.

Legacy evidence: [jarvis_log_client/client.py](https://github.com/alexberardi/jarvis-log-client/blob/7d6a3e8b62d2df8e1165f3c535481eba87ac5388/jarvis_log_client/client.py), [jarvis_log_client/handlers.py](https://github.com/alexberardi/jarvis-log-client/blob/7d6a3e8b62d2df8e1165f3c535481eba87ac5388/jarvis_log_client/handlers.py).

### LOGLIB-01

**Structured levels, batching, flush and remote node/app log shipping**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [LOG-01](Get-Foundations-Features.md#log-01).

Go evidence or governing decision: [internal/platform/logging](../../internal/platform/logging).

Server logger and batched sink replace library; external Python clients remain supported at the HTTP boundary.

### LOGLIB-02

**Package logging through the node runtime**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [LOG-01](Get-Foundations-Features.md#log-01).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Sidecar maps logs to stderr; check correlation and delivery of package fields before claiming full Python client parity.

## jarvis-web-scraper

Current disposition: **consolidate**.

Legacy evidence: [jarvis_web_scraper](https://github.com/alexberardi/jarvis-web-scraper/tree/648661dfda5eb4a20c28b1fbeb0a340e39b6c4e5/jarvis_web_scraper).

### SCRAPE-01

**Fetch readable page content with bounded redirects and extraction fallback**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: [CONF-01](Get-Foundations-Features.md#conf-01).

Go evidence or governing decision: [internal/modules/cc/servertools](../../internal/modules/cc/servertools), [internal/modules/recipes/extract](../../internal/modules/recipes/extract).

Capability implemented in Go search/research and recipe extraction; not byte-identical trafilatura output.

### SCRAPE-02

**SSRF protection for DNS resolution, addresses and every redirect**

Status: **complete**. Disposition: **consolidate**.

Prerequisites: none within this inventory.

Go evidence or governing decision: [internal/platform/ssrf](../../internal/platform/ssrf).

Shared Go URL boundary used by outbound fetchers.
