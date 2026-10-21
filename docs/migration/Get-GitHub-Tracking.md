# GitHub migration tracking

[Feature issue index](Get-Feature-Issues.md): direct links for all 315 feature IDs.

[Migration inventory](Get-Migration-Tracker.md) | [Dependency map](Get-Dependency-Map.md)

The [Jarvis-server Board](https://github.com/users/opticon454/projects/5) and [Jarvis-server roadmap](https://github.com/users/opticon454/projects/3) use the same feature issues in [opticon454/jarvis-server](https://github.com/opticon454/jarvis-server/issues). Updating an issue updates the shared record. Project fields are separate, so keep the Status field aligned in both projects when moving work.

## Identification and labels

Issue titles begin with the stable feature ID, for example `[AUTH-01]` or `[DISPLAY-01]`. Keep these IDs when splitting work, adding implementation PRs or revising scope. Each issue links to the reviewed feature, source evidence and prerequisites.

- `migration`: all 315 inventory entries.
- `area:*`: the feature register, with `area:display` for the five display requirements.
- `status:complete`, `status:in-progress`, `status:not-started`, `status:not-applicable`: the audited migration status.
- `scope:consolidate`, `scope:absorb`, `scope:new`, `scope:deferred`: the intended treatment of the feature.
- `scope:cut`, `scope:retire`: explicit cuts or replacement/retirement work. Read the issue disposition before scheduling work.
- `scope:excluded`: the 13 Pantry reference entries, excluded by the user.
- `scope:review-scope`: retained clients, external boundaries, documentation, deployment changes and other cases needing the detailed disposition in the issue. This label does not authorise exclusion.

## Project workflow

| Audit or disposition         | Issue state         | Project Status            |
| ---------------------------- | ------------------- | ------------------------- |
| complete (122)               | Closed: completed   | Done                      |
| in progress (104)            | Open                | In Progress / In progress |
| not started, actionable (52) | Open                | Todo                      |
| explicit cuts (24)           | Closed: not planned | Not planned               |
| Pantry excluded (13)         | Closed: not planned | Reference                 |

The 24 explicit cuts retain `status:not-started`: they have not been implemented, and existing decisions say not to migrate them. Closing them as not planned avoids counting cuts as delivered functionality. Pantry entries are references, not sprint work. Other external components remain open unless the audited capability is already complete.

Statuses are the source review snapshot of 21 October 2026. Complete does not certify fresh runtime, legacy contract or hardware acceptance. Update the issue label, both project Status fields and the Markdown row when new evidence changes an assessment. Link the implementation PR and validation results before closing an outstanding feature as completed.

Owners, delivery dates, quarters and iterations are unset. Assign these during planning. Use the roadmap's existing Start date and Target date fields once dates are agreed.

## Dependencies and ordering

The inventory has 527 directed prerequisite links. Issue bodies link to prerequisite issues; GitHub blocked-by relationships expose the same graph. These describe required capabilities, including already completed capabilities and references, rather than promising that every open predecessor is scheduled migration work. The [dependency ledger](Get-Dependency-Map.md) records the complete order and external prerequisites.

Start with the lowest outstanding dependency level, resolve scope decisions for retained components, then schedule dependent features. A completed capability may support later work without reopening it. Check disposition before treating a closed not-planned prerequisite as supplied functionality.

## Heading abbreviations

| Prefix   | Source repository or feature family        |
| -------- | ------------------------------------------ |
| ADMIN    | jarvis-admin                               |
| APPLE    | jarvis-device-apple                        |
| AUDACY   | jarvis-cmd-audacy                          |
| AUTH     | jarvis-auth                                |
| AUTHLIB  | jarvis-auth-client                         |
| BARSTOOL | jarvis-cmd-barstool                        |
| BT       | jarvis-cmd-bluetooth                       |
| CAL      | jarvis-cmd-calendar                        |
| CC       | jarvis-command-center                      |
| COMFY    | jarvis-cmd-comfyui                         |
| CONF     | jarvis-config-service                      |
| CONFLIB  | jarvis-config-client                       |
| DAD      | jarvis-command-tell-dad-joke               |
| DESK     | jarvis-host-agent                          |
| DISPLAY  | jarvis-server (current checkout additions) |
| DOCS     | jarvis-docs                                |
| DRIVE    | jarvis-cmd-drive-time                      |
| EMAIL    | jarvis-cmd-email                           |
| ENT      | jarvis-cmd-entertainment-knowledge         |
| GOVEE    | jarvis-device-govee                        |
| GREET    | jarvis-command-greeting                    |
| HA       | jarvis-home-assistant-integration          |
| HC       | jarvis-device-homeconnect                  |
| HERMES   | jarvis-pp-hermes                           |
| HK       | jarvis-device-homekit                      |
| HUE      | jarvis-device-hue                          |
| INSTALL  | jarvis-installer                           |
| JDT      | jarvis-developer-toolkit                   |
| KASA     | jarvis-device-kasa                         |
| LIFX     | jarvis-device-lifx                         |
| LLM      | jarvis-llm-proxy-api                       |
| LOG      | jarvis-logs                                |
| LOGLIB   | jarvis-log-client                          |
| LOOP     | jarvis-loop                                |
| MCP      | jarvis-mcp                                 |
| MED      | jarvis-cmd-medication                      |
| METEO    | jarvis-cmd-meteo-weather                   |
| MISTRAL  | jarvis-pp-mistral                          |
| MOB      | jarvis-node-mobile                         |
| MSG      | jarvis-cmd-messages                        |
| MUSIC    | jarvis-cmd-music-assistant                 |
| NEST     | jarvis-device-nest                         |
| NEW      | jarvis-server (current checkout additions) |
| NEWS     | jarvis-cmd-news                            |
| NODE     | jarvis-node-setup                          |
| NOTIFY   | jarvis-notifications                       |
| OCR      | jarvis-ocr-service                         |
| OPS      | jarvis                                     |
| OSX      | jarvis-osx-api                             |
| OW       | jarvis-cmd-open-weather                    |
| PANDORA  | jarvis-cmd-pandora                         |
| PANTRY   | jarvis-pantry                              |
| PHONE    | jarvis-phone-gateway                       |
| PLAT     | jarvis-data-services                       |
| PRUN     | jarvis-pantry-runner                       |
| PWEB     | jarvis-pantry-web                          |
| QA       | jarvis-integration-tests                   |
| RECIPE   | jarvis-recipes-server                      |
| RELAY    | jarvis-relay                               |
| RESIDEO  | jarvis-device-resideo                      |
| RMOB     | jarvis-recipes-mobile                      |
| RPROTO   | jarvis-recipes                             |
| RT       | jarvis-cmd-rotten-tomatoes                 |
| RTOLD    | jarvis-command-rotten-tomatoes             |
| SCHLAGE  | jarvis-device-schlage                      |
| SCRAPE   | jarvis-web-scraper                         |
| SDK      | jarvis-command-sdk                         |
| SEARCH   | jarvis-cmd-web-search                      |
| SET      | jarvis-settings-client                     |
| SETOLD   | jarvis-settings-server                     |
| SIMPLE   | jarvis-device-simplisafe                   |
| SPORT    | jarvis-cmd-sports                          |
| SPOT     | jarvis-cmd-spotify                         |
| STT      | jarvis-whisper-api                         |
| TTS      | jarvis-tts                                 |
| WEB      | jarvis-web                                 |
| ZWAVE    | jarvis-device-zwave                        |
