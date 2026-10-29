# Packages feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-cmd-audacy

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-audacy/tree/7085c2c76074763d3d62642022f9d8908e5a7fdd/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-audacy/blob/7085c2c76074763d3d62642022f9d8908e5a7fdd/jarvis_package.yaml).

### AUDACY-01

**Tune by station/name/alias, resume, stop, now-playing and station catalogue**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [NODE-18](Get-Nodes-Features.md#node-18).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. mpd plus local PulseAudio/PipeWire output; Pi system-service setup.

## jarvis-cmd-barstool

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-barstool/tree/d3d1103f174484fae6fddf8d75544addf234f117/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-barstool/blob/d3d1103f174484fae6fddf8d75544addf234f117/jarvis_package.yaml).

### BARSTOOL-01

**Latest sitemap headlines, sport/category filtering and result limits**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [SCRAPE-02](Get-Foundations-Features.md#scrape-02).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Barstool sitemap access; no API key.

## jarvis-cmd-bluetooth

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-bluetooth/tree/e4206fe7f8fac8c089dc40366c86321b7bcd6151/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-bluetooth/blob/e4206fe7f8fac8c089dc40366c86321b7bcd6151/jarvis_package.yaml).

### BT-01

**Voice scan, pair, connect, disconnect, forget and status; phone/speaker/bridge roles**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [NODE-17](Get-Nodes-Features.md#node-17).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Go Bluetooth support exists, but package uses its own Python providers; verify both paths.

## jarvis-cmd-calendar

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-calendar/tree/a4cc787828b465b7df5ce1d490b19010c79dea74/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-calendar/blob/a4cc787828b465b7df5ce1d490b19010c79dea74/jarvis_package.yaml).

### CAL-01

**iCloud/Google calendar queries by day/range and availability context**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-08](Get-Orchestration-Features.md#cc-08).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. CalDAV credentials or Google OAuth; date resolution.

### CAL-02

**Create calendar events with idempotency keys and appointment-signal handling**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [CAL-01](Get-Packages-Features.md#cal-01), [CC-23](Get-Orchestration-Features.md#cc-23).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

add_event exists beyond the README feature list; verify retries and confirmation against both providers.

### CAL-03

**Background upcoming-event alerts and appointment signals**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [CAL-01](Get-Packages-Features.md#cal-01), [NODE-09](Get-Nodes-Features.md#node-09), [CC-22](Get-Orchestration-Features.md#cc-22).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

calendar_alerts needs a general package agent scheduler.

## jarvis-cmd-comfyui

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-comfyui/tree/ee69cbc8a088be58068d332c9af816cfea837b17/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-comfyui/blob/ee69cbc8a088be58068d332c9af816cfea837b17/jarvis_package.yaml).

### COMFY-01

**Text-to-image, text-to-video and attached-photo animation with model capability detection**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-12](Get-Orchestration-Features.md#cc-12).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Reachable ComfyUI and suitable checkpoints; completion agent and media upload are separate.

### COMFY-02

**Poll render jobs, enforce timeout, deliver video/image or failure to inbox**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [COMFY-01](Get-Packages-Features.md#comfy-01), [NODE-09](Get-Nodes-Features.md#node-09), [NODE-30](Get-Nodes-Features.md#node-30).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Blocked by general agent scheduling and SDK media attachment upload.

## jarvis-cmd-drive-time

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-drive-time/tree/4f407c88968d2b15b75ea231684d64dc0da0162e/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-drive-time/blob/4f407c88968d2b15b75ea231684d64dc0da0162e/jarvis_package.yaml).

### DRIVE-01

**Traffic-aware drive time and distance with home/default or named origin**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Google Directions/Geocoding API and home address.

### DRIVE-02

**Monitor calendar locations and propose deduplicated leave-by reminders**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [DRIVE-01](Get-Packages-Features.md#drive-01), [CAL-01](Get-Packages-Features.md#cal-01), [NODE-09](Get-Nodes-Features.md#node-09), [CC-23](Get-Orchestration-Features.md#cc-23).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

Background drive_time_alerts is not scheduled by the Go package runtime.

## jarvis-cmd-email

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-email/tree/cc1aa7df41d5ff9733a170fa352b789d0eee3db1/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-email/blob/cc1aa7df41d5ff9733a170fa352b789d0eee3db1/jarvis_package.yaml).

### EMAIL-01

**Gmail/IMAP email list/read/search, send/reply/forward, archive/trash, star/read flags, triage and subscription cleanup**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-23](Get-Orchestration-Features.md#cc-23).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. OAuth or IMAP/SMTP credentials; confirmation/callback and user-scoped storage.

### EMAIL-02

**VIP/urgent alerts and daily email digest**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [EMAIL-01](Get-Packages-Features.md#email-01), [NODE-09](Get-Nodes-Features.md#node-09), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

Background email_alerts agent remains external.

### EMAIL-03

**Smart-reply drafts and appointment scanning into calendar proposals**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [EMAIL-01](Get-Packages-Features.md#email-01), [NODE-09](Get-Nodes-Features.md#node-09), [CAL-02](Get-Packages-Features.md#cal-02), [CC-23](Get-Orchestration-Features.md#cc-23).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

smart_reply and appointment_scan both require agent scheduling.

## jarvis-cmd-entertainment-knowledge

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-entertainment-knowledge/tree/4e994c31367c8bc5bf55309ec5d59ff4b8a7040a/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-entertainment-knowledge/blob/4e994c31367c8bc5bf55309ec5d59ff4b8a7040a/jarvis_package.yaml).

### ENT-01

**Movie/TV synopsis, cast/director/creators, related titles and clickable filmography cards**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-23](Get-Orchestration-Features.md#cc-23).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. TMDB key, inbox and callbacks.

## jarvis-cmd-medication

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-medication/tree/c8ecf650c7659575c133c306d25f4d13e96d9c96/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-medication/blob/c8ecf650c7659575c133c306d25f4d13e96d9c96/jarvis_package.yaml).

### MED-01

**Personal/household medication records, daily status and mark dose taken**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-31](Get-Orchestration-Features.md#cc-31).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Command data forms, scoped ownership and callbacks; no medical recommendations implied.

### MED-02

**Due/overdue dose reminders, grace period, repeat warnings and action buttons**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [MED-01](Get-Packages-Features.md#med-01), [NODE-09](Get-Nodes-Features.md#node-09), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

medication_reminders is a package agent, not the built-in reminder bridge.

## jarvis-cmd-messages

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-messages/tree/b2d0dd03b9a2cd24c230744d6b41639e10278d83/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-messages/blob/b2d0dd03b9a2cd24c230744d6b41639e10278d83/jarvis_package.yaml).

### MSG-01

**Check unread texts, read threads and draft a reply with voice-send or editable inbox confirmation**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [OSX-02](Get-External-Features.md#osx-02).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. External macOS API, user credentials, contacts and confirmation path.

## jarvis-cmd-meteo-weather

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-meteo-weather/tree/0230d31ee8f7e441f8c21091d25fa538dc8f28e0/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-meteo-weather/blob/0230d31ee8f7e441f8c21091d25fa538dc8f28e0/jarvis_package.yaml).

### METEO-01

**Current/hourly/minutely/daily weather with default location, geocoding and date-range forecasts**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-08](Get-Orchestration-Features.md#cc-08).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Open-Meteo network access, location and units.

### METEO-02

**Periodically inject weather context into memory**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [METEO-01](Get-Packages-Features.md#meteo-01), [NODE-09](Get-Nodes-Features.md#node-09), [CC-14](Get-Orchestration-Features.md#cc-14).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

weather_context agent needs background scheduling.

## jarvis-cmd-music-assistant

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-music-assistant/tree/f0f3888d71e7db8a0e00f2d3396c368fe5bb913f/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-music-assistant/blob/f0f3888d71e7db8a0e00f2d3396c368fe5bb913f/jarvis_package.yaml).

### MUSIC-01

**Search/play/queue by artist/album/track, select room/player, pause/resume/stop, next/previous, volume/mute/shuffle/repeat**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [NODE-18](Get-Nodes-Features.md#node-18).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. External Music Assistant, token and shairport-sync AirPlay receiver.

## jarvis-cmd-news

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-news/tree/920c88f399b6ded49c97f571bb3238d213bc8a44/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-news/blob/920c88f399b6ded49c97f571bb3238d213bc8a44/jarvis_package.yaml).

### NEWS-01

**RSS headlines by category/count, custom feeds and rich follow-up actions**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-23](Get-Orchestration-Features.md#cc-23).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. RSS connectivity, inbox/callback support.

### NEWS-02

**Background headline alerts and deduplication**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [NEWS-01](Get-Packages-Features.md#news-01), [NODE-09](Get-Nodes-Features.md#node-09), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

news_alerts exists in source although omitted from the short README.

## jarvis-cmd-open-weather

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-open-weather/tree/3c53e8ebb4ae5db8f06244d515dc49ed92ed0705/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-open-weather/blob/3c53e8ebb4ae5db8f06244d515dc49ed92ed0705/jarvis_package.yaml).

### OW-01

**Current conditions and multi-day forecasts with configurable location/units**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-08](Get-Orchestration-Features.md#cc-08).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. OpenWeather API key, geocoding and date context.

### OW-02

**Periodically inject OpenWeather context into memory**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [OW-01](Get-Packages-Features.md#ow-01), [NODE-09](Get-Nodes-Features.md#node-09), [CC-14](Get-Orchestration-Features.md#cc-14).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

Alternative weather provider; resolve get_weather naming conflicts when installing both.

## jarvis-cmd-pandora

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-pandora/tree/576273cbf411343fb29bd6e2e25f4da27ed59645/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-pandora/blob/576273cbf411343fb29bd6e2e25f4da27ed59645/jarvis_package.yaml).

### PANDORA-01

**Find/create radio stations, play/pause/stop, skip, thumbs-up/down, list stations and now-playing**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [NODE-18](Get-Nodes-Features.md#node-18).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Pandora account and local supported audio player.

## jarvis-cmd-rotten-tomatoes

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-rotten-tomatoes/tree/32da9077859fec035368a838e450e1b5634ac841/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-rotten-tomatoes/blob/32da9077859fec035368a838e450e1b5634ac841/jarvis_package.yaml).

### RT-01

**Movie/TV search, critic/audience ratings, cast and in-cinemas listings**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [SCRAPE-02](Get-Foundations-Features.md#scrape-02).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. External page parsing; changes in the provider can break extraction.

## jarvis-cmd-sports

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-sports/tree/c2e7a1e65402ee294a6f8c1a2e5e950edd70b9d4/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-sports/blob/c2e7a1e65402ee294a6f8c1a2e5e950edd70b9d4/jarvis_package.yaml).

### SPORT-01

**Past/live scores and upcoming schedules, fuzzy team names, professional/college leagues and favourites**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-08](Get-Orchestration-Features.md#cc-08).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. ESPN access and scoped favourites data.

### SPORT-02

**Favourite-team background game alerts**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [SPORT-01](Get-Packages-Features.md#sport-01), [NODE-09](Get-Nodes-Features.md#node-09), [NOTIFY-03](Get-Services-Features.md#notify-03).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

sports_alerts requires the general background runtime.

## jarvis-cmd-spotify

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-spotify/tree/2760b26b0fea2a433766e979dcb27a0ebaf425be/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-spotify/blob/2760b26b0fea2a433766e979dcb27a0ebaf425be/jarvis_package.yaml).

### SPOT-01

**Search/play music, pause/stop/skip/previous, volume/shuffle/repeat and now-playing**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [NODE-18](Get-Nodes-Features.md#node-18).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Spotify OAuth, account entitlement and go-librespot local playback daemon.

### SPOT-02

**Playback daemon keepalive and proactive OAuth refresh**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [SPOT-01](Get-Packages-Features.md#spot-01), [NODE-09](Get-Nodes-Features.md#node-09).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

spotify_keepalive is not loaded as a scheduled agent.

## jarvis-cmd-web-search

Current disposition: **consolidate**.

Legacy evidence: [commands](https://github.com/alexberardi/jarvis-cmd-web-search/tree/cda06e97d29f04f8f35c665395fea3dd0ffc6f44/commands), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-cmd-web-search/blob/cda06e97d29f04f8f35c665395fea3dd0ffc6f44/jarvis_package.yaml).

### SEARCH-01

**Live web search through DuckDuckGo or Bing with configurable region**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [CC-16](Get-Orchestration-Features.md#cc-16).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. CC quick-search is related; it does not prove both package providers work.

## jarvis-command-greeting

Current disposition: **consolidate**.

Legacy evidence: [command.py](https://github.com/alexberardi/jarvis-command-greeting/blob/6339538ef4e8052f4bbfb1e85597ea256fb40e36/command.py), [README.md](https://github.com/alexberardi/jarvis-command-greeting/blob/6339538ef4e8052f4bbfb1e85597ea256fb40e36/README.md), [jarvis_command.yaml](https://github.com/alexberardi/jarvis-command-greeting/blob/6339538ef4e8052f4bbfb1e85597ea256fb40e36/jarvis_command.yaml).

### GREET-01

**Named greetings in ten languages**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Original flat-file package layout must be accepted by installer.

## jarvis-command-rotten-tomatoes

Current disposition: **consolidate**.

Legacy evidence: [command.py](https://github.com/alexberardi/jarvis-command-rotten-tomatoes/blob/421814d4f6de808df7a3ae700750696066467022/command.py), [README.md](https://github.com/alexberardi/jarvis-command-rotten-tomatoes/blob/421814d4f6de808df7a3ae700750696066467022/README.md), [jarvis_command.yaml](https://github.com/alexberardi/jarvis-command-rotten-tomatoes/blob/421814d4f6de808df7a3ae700750696066467022/jarvis_command.yaml).

### RTOLD-01

**Earlier flat-file Rotten Tomatoes search and in-cinemas command**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08), [SCRAPE-02](Get-Foundations-Features.md#scrape-02).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. Separate repository from jarvis-cmd-rotten-tomatoes; do not silently count both as one migration.

### RTOLD-02

**README-advertised streaming and opening-this-week browsing**

Status: **not started**. Disposition: **legacy claim; confirm scope**.

Prerequisites: [RTOLD-01](Get-Packages-Features.md#rtold-01).

Go evidence: no corresponding implementation found in the reviewed target surfaces. See the linked legacy source above.

README advertises these, but command parameter enum exposes only search/in_theaters. Treat as legacy claim requiring scope confirmation, not demonstrated source functionality.

## jarvis-command-tell-dad-joke

Current disposition: **consolidate**.

Legacy evidence: [command.py](https://github.com/alexberardi/jarvis-command-tell-dad-joke/blob/cf0577066fdd376141dae6f1b6e0087e1040f42f/command.py), [README.md](https://github.com/alexberardi/jarvis-command-tell-dad-joke/blob/cf0577066fdd376141dae6f1b6e0087e1040f42f/README.md), [jarvis_command.yaml](https://github.com/alexberardi/jarvis-command-tell-dad-joke/blob/cf0577066fdd376141dae6f1b6e0087e1040f42f/jarvis_command.yaml).

### DAD-01

**Fetch a random dad joke from icanhazdadjoke**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-08](Get-Nodes-Features.md#node-08).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/packages](../../internal/agent/packages).

Python package code remains external; bridge available, package-specific Go-host acceptance not demonstrated. External API and original flat-file package layout.
