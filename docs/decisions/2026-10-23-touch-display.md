# Local touch display with centrally inherited configuration

- **Date:** 2026-10-23, user plan and follow-up instructions.

The server owns display defaults and sparse node overrides. Built-ins, defaults and node values merge
in that order, and a content revision covers the result and the protected PIN verifier. Nodes fetch,
cache and acknowledge through their existing authentication. Themes receive a separate versioned
message API and cannot perform privileged operations directly. See [Use the touch display](../agent/Use-TouchDisplay.md).

The user replaced the planned six-digit PIN with **4 to 20 numeric digits**, requested a dedicated
Surface kiosk user with Xfce automatic login, and requested a circular waveform reacting to spoken
replies, subsequently clarified to occupy the full screen during replies. The waveform uses existing voice amplitude callbacks rather than adding audio capture.
The Surface deployment keeps a reversible copy of the original voice installation and runs only
one microphone owner at a time. Cast V2 reception is experimental and separately switchable.
