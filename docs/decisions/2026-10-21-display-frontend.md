# Display frontend with mandatory live Home Assistant status

- **Date:** 2026-10-21 (user)

Add a Display frontend to the feature migration inventory. It must serve as an interactive
Jarvis dashboard on a Raspberry Pi with an external display, a Surface-style device and an
Android tablet. Home Assistant integration is mandatory for live status reporting of its
devices.

Track the frontend, live device tiles and controls, Home Assistant state integration,
full-screen/touch operation and acceptance on each device class as separate deliverables.
Live state must include changes made outside Jarvis, initial state, reconnect recovery and
clear stale or unavailable status. Device actions must respect household access and show
their result.

The implementation has not started. The existing admin frontend is a reuse candidate;
responsive admin pages alone do not satisfy this requirement. Browser delivery is a planning
proposal, not a decision to build separate native applications. The connector location and
transport remain to be selected during implementation.

This adds scope to the [migration tracker](../migration/Get-Migration-Tracker.md) and its
[Display frontend entries](../migration/Get-New-Features.md#display-01). It does not supersede
the [home-state watch event decision](2026-10-14-home-state-reaches-watches-as-node-signals.md)
or the existing web visual direction.
