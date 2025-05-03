# 🔐 kv2

> [!WARNING]
> A significant refactor is currently underway in preparation for 2.0.0, and `main` is very much "under construction". You can view the repository as it appeared at 1.0.0 release [here](https://github.com/hugginsio/kv2/tree/v1.0.0).

`kv2` is an end-to-end encrypted secrets manager for tailnets. This repository contains the server and two client implementations: a proxy agent and the CLI.

## ✨ Features

- **Encrypted**: secrets are end-to-end encrypted with [age][1] and user-controlled keys.
- **Versioned**: track and coordinate secret changes with basic integer-indexed version history.
- **Secure**: built with the [Tailscale][0] client library to provide secure access between components.
- **KMS Support**: optionally integrate with a cloud key management system for securely retrieving [age][1] keys.
- **Backup Support**: optionally leverage cloud storage for automated backup & recovery of the secrets database.

These features make `kv2` the perfect secrets management solution for my homelab environment.

## 🤝🏻 Thanks

- [@tailscale/setec][9], which largely inspired `kv2`.

---

"Tailscale" is a registered trademark of Tailscale Inc. The `kv2` project is not endorsed by, sponsored by, or affiliated with Tailscale Inc.

<!-- Links -->
[0]: https://tailscale.com/
[1]: https://github.com/FiloSottile/age
[9]: https://github.com/tailscale/setec
