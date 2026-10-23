# Changelog

## 20261023-2209

- Fixed the PIN-exit test fixture to lock shared engine state while the supervisor runs; retained every exit, rate-limit and voice-continuation assertion (issue 322).

## 20261023-2201

- Added inherited server and per-node speech strand colour/emphasis settings, with stronger 150% default emphasis and bounded colour/intensity validation. Preserved the revision format for older cached configuration.
- Increased the Home screen clock by 50% while retaining the separate 75% idle clock layout.

## 20261023-2147

- Replaced the travelling speech history with fixed radial flare shapes that burst outwards together with speech amplitude, retaining full-screen presentation and reduced-motion support.
- Enlarged the idle clock to a responsive 75% display area for digital, analogue and combined views.

## 20261023-2140

- Changed the live circular speech waveform to fill the display during replies, retaining the trusted exit control above it. Added rendered landscape, portrait and phone coverage for its size and return to the normal view.

## 20261023-2135

- Increased touch control boundary contrast to 3.7:1 and added a rendered accessibility regression check (issue 321).

## 20261023-2127

- Replaced the obsolete Chromium restore-popup switch so restarting the kiosk does not obscure touch controls (issue 320).

## 20261023-2124

- Fixed the local Chromium display profile prompting for a new login keyring in an auto-login kiosk account (issue 319). The separate streaming profile keeps its credential storage.

## 20261023-2116

- Fixed PIN exit and remote display disablement to close display-owned media and receivers while voice continues (issue 318).

## 20261023-2040

- Added the opt-in Linux touch display, inherited server configuration, node acknowledgements, private offline cache, uploaded photos, sandboxed themes and a numeric exit PIN of 4 to 20 digits.
- Added device catalogue controls through existing household and plugin paths, local and Home Assistant media routing, separate Playercast and experimental Cast receivers, and live circular speech visualisation.
- Added display, access-control and browser checks, a starter theme and display operation/theme API documentation. Requires cc migration `00320`; existing nodes retain their behaviour until display mode is enabled.
