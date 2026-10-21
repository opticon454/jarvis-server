# Devices feature inventory

[Tracker](Get-Migration-Tracker.md) | [Dependencies](Get-Dependency-Map.md) | [Repository snapshot](Get-Repository-Inventory.md)

Statuses describe consolidation into the current checkout. Read the tracker definitions before treating a feature as release-ready. Prerequisites are required capabilities, not claims that upstream rows are finished.

## jarvis-device-apple

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-apple/tree/c5effce1570f84f8a872814279f704c54d10a554/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-apple/blob/c5effce1570f84f8a872814279f704c54d10a554/jarvis_package.yaml).

### APPLE-01

**Apple TV/HomePod discovery, PIN pairing, playback, power and volume**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. pyatv-compatible hardware and persistent pairing credentials.

## jarvis-device-govee

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-govee/tree/e760cc5be71e44b81ce61306fd692bf5c9a56cb0/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-govee/blob/e760cc5be71e44b81ce61306fd692bf5c9a56cb0/jarvis_package.yaml).

### GOVEE-01

**Discover and control Govee lights, strips, plugs and supported appliances; power and supported lighting properties**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. Govee developer API key and model-specific capabilities.

## jarvis-device-homeconnect

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-homeconnect/tree/f020c5d1237d954cde372f111fd1b649da31cdd8/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-homeconnect/blob/f020c5d1237d954cde372f111fd1b649da31cdd8/jarvis_package.yaml).

### HC-01

**Discover appliances, inspect status/programs, start selected dishwasher programmes and stop**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. Home Connect OAuth and compatible Bosch/Siemens appliances.

## jarvis-device-homekit

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-homekit/tree/cb966b42c346acabb82cf663d03813777954cd72/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-homekit/blob/cb966b42c346acabb82cf663d03813777954cd72/jarvis_package.yaml).

### HK-01

**Local HAP discovery/pairing; lights/switches, locks, thermostats and read-only sensors**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. aiohomekit and per-accessory pairing records.

## jarvis-device-hue

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-hue/tree/2aeb0e098963b7e5599909bd894f776e9b57e4b7/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-hue/blob/2aeb0e098963b7e5599909bd894f776e9b57e4b7/jarvis_package.yaml).

### HUE-01

**Bridge discovery, power, brightness, RGB/XY colour and colour temperature for individual/multiple lights**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. Hue bridge IP and link-button API credential.

## jarvis-device-kasa

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-kasa/tree/d00df9d894b3594d17761a20d41e50bd5106fd98/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-kasa/blob/d00df9d894b3594d17761a20d41e50bd5106fd98/jarvis_package.yaml).

### KASA-01

**LAN discovery/control of plugs, bulbs, switches, dimmers, strips and supported fans**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. python-kasa; Tapo/Kasa model and authentication differences.

## jarvis-device-lifx

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-lifx/tree/ddc2b083fe824245b44ea566d9d7d9337b1c3781/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-lifx/blob/ddc2b083fe824245b44ea566d9d7d9337b1c3781/jarvis_package.yaml).

### LIFX-01

**UDP LAN discovery, power, brightness, colour and colour temperature**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. lifxlan and reachable LAN bulbs/strips.

## jarvis-device-nest

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-nest/tree/0e9665859947c2cc01a6c59a85390486ad980819/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-nest/blob/0e9665859947c2cc01a6c59a85390486ad980819/jarvis_package.yaml).

### NEST-01

**SDM discovery, thermostat temperature/mode and camera/doorbell metadata**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. Google Device Access project and OAuth credentials.

### NEST-02

**Generate authenticated camera stream source and play live camera video**

Status: **not started**. Disposition: **deferred retained capability**.

Prerequisites: [NEST-01](Get-Devices-Features.md#nest-01), [CC-27](Get-Orchestration-Features.md#cc-27).

Go evidence or governing decision: [internal/modules/cc/smarthome.go](../../internal/modules/cc/smarthome.go).

Go camera streaming deliberately deferred; camera metadata is not streaming.

## jarvis-device-resideo

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-resideo/tree/d85f2d889564684a569a7e8403451612a7431fb6/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-resideo/blob/d85f2d889564684a569a7e8403451612a7431fb6/jarvis_package.yaml).

### RESIDEO-01

**Thermostat discovery/status, target temperature, HVAC/fan mode, holds and schedule resume**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. Honeywell Home API OAuth and unit conversion.

## jarvis-device-schlage

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-schlage/tree/e031bd0a9d142d4b1ef30ebc8be371b1f82add5c/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-schlage/blob/e031bd0a9d142d4b1ef30ebc8be371b1f82add5c/jarvis_package.yaml).

### SCHLAGE-01

**Lock discovery, lock/unlock and lock/jam/battery status**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. Schlage/Allegion authentication and supported Encode locks.

## jarvis-device-simplisafe

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-simplisafe/tree/8e19e978de675ed933ba39e91a873811bb12f08e/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-simplisafe/blob/8e19e978de675ed933ba39e91a873811bb12f08e/jarvis_package.yaml).

### SIMPLE-01

**V2/V3 system discovery, home/away arm, disarm, sensor state and lock/unlock**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. SimpliSafe rotating refresh token and system permissions.

## jarvis-device-zwave

Current disposition: **consolidate**.

Legacy evidence: [device_families](https://github.com/alexberardi/jarvis-device-zwave/tree/ea40b194fff28a835c51691c3c99e719a737a9c9/device_families), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-device-zwave/blob/ea40b194fff28a835c51691c3c99e719a737a9c9/jarvis_package.yaml).

### ZWAVE-01

**Z-Wave switch/light/lock/climate/cover discovery, state and control**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-26](Get-Orchestration-Features.md#cc-26).

Go evidence or governing decision: [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py), [internal/agent/devices](../../internal/agent/devices).

Python protocol remains external; app-control bridge built, vendor hardware and voice parity unverified. External Z-Wave JS UI/server and WebSocket connection.

### ZWAVE-02

**Background device cache refresh**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [ZWAVE-01](Get-Devices-Features.md#zwave-01), [NODE-09](Get-Nodes-Features.md#node-09).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

ZWaveAgent requires the general package agent scheduler.

## jarvis-home-assistant-integration

Current disposition: **consolidate**.

Legacy evidence: [ha_shared](https://github.com/alexberardi/jarvis-home-assistant-integration/tree/e403d5a9fc3da2aac4a859edc6cebc10e9eac732/ha_shared), [commands](https://github.com/alexberardi/jarvis-home-assistant-integration/tree/e403d5a9fc3da2aac4a859edc6cebc10e9eac732/commands), [agents](https://github.com/alexberardi/jarvis-home-assistant-integration/tree/e403d5a9fc3da2aac4a859edc6cebc10e9eac732/agents), [device_managers](https://github.com/alexberardi/jarvis-home-assistant-integration/tree/e403d5a9fc3da2aac4a859edc6cebc10e9eac732/device_managers), [jarvis_package.yaml](https://github.com/alexberardi/jarvis-home-assistant-integration/blob/e403d5a9fc3da2aac4a859edc6cebc10e9eac732/jarvis_package.yaml).

### HA-01

**Home Assistant discovery/import and device-manager registry**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [NODE-16](Get-Nodes-Features.md#node-16), [CC-25](Get-Orchestration-Features.md#cc-25).

Go evidence or governing decision: [internal/agent/sidecar/devices.go](../../internal/agent/sidecar/devices.go), [internal/modules/cc/smarthome.go](../../internal/modules/cc/smarthome.go).

Python manager bridge exists; external HA remains authoritative for entities.

### HA-02

**Voice device control, fuzzy entity/room resolution and state queries**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [HA-01](Get-Devices-Features.md#ha-01), [CC-26](Get-Orchestration-Features.md#cc-26), [NODE-15](Get-Nodes-Features.md#node-15).

Go evidence or governing decision: [internal/modules/cc/smarthome_control.go](../../internal/modules/cc/smarthome_control.go), [internal/agent/sidecar/jarvis_sidecar.py](../../internal/agent/sidecar/jarvis_sidecar.py).

Package and generic control paths differ; Pi voice built-in gap remains.

### HA-03

**WebSocket state updates, background cache and proactive device-state signals**

Status: **not started**. Disposition: **consolidate**.

Prerequisites: [HA-01](Get-Devices-Features.md#ha-01), [NODE-09](Get-Nodes-Features.md#node-09), [CC-22](Get-Orchestration-Features.md#cc-22).

Go evidence or governing decision: [internal/agent/packages/manifest.go](../../internal/agent/packages/manifest.go).

Legacy home_assistant background agent is not generally scheduled by Go; display needs a dependable live bridge.

### HA-04

**Home Assistant credentials, reconnect and state refresh across restarts**

Status: **in progress**. Disposition: **consolidate**.

Prerequisites: [HA-01](Get-Devices-Features.md#ha-01), [NODE-06](Get-Nodes-Features.md#node-06).

Go evidence or governing decision: [internal/agent/settings](../../internal/agent/settings), [internal/agent/devices](../../internal/agent/devices).

Stored secrets/manager calls exist; validate persistent live event subscription separately.
