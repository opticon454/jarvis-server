# Household key held on the phone, with a recovery code

- **Date:** 2026-11-06 (user)

Config for the whole fleet is sealed once with a household key (KH) and stored on jarvisd as ciphertext. The user chose that the phone keeps the key and jarvisd never holds it, with a recovery code the user keeps in a password vault (1Password and similar) to get back in after a lost, stolen or reinstalled phone. Nothing is stored in plaintext or in variable files.

Reason: per-node K2 meant entering every secret once per node, a node added later missed earlier pushes, and losing the phone lost every key. A server-held key would have made new nodes hands-free, but the user wanted jarvisd to stay ciphertext-only.

Consequences, accepted:

- A new node still needs the household key delivered once, through the existing K2 relay, from the phone. It then pulls everything the household already holds.
- Lose the phone and the code and the fleet secrets are gone. The user creates a new key and enters them again.
- One code covers every node and every phone in the household, so a leaked code means rotating the key.
- Go only. Python nodes (jarvis-node-setup) are out of scope.

Not done: encrypting the key and secrets at rest on a node (they stay in 0600 files in a 0700 directory, as before). See [docs/cc/14-fleet-config-and-recovery.md](../cc/14-fleet-config-and-recovery.md) section 8.
