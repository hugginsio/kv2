# 🔐 kv2

> [!WARNING]
> `kv2` is still v0. Breaking changes are expected.

`kv2` is a key-value key vault built for homelab secrets management. This repository contains the server and client components, as well as the reference CLI implementation.

## ✨ Features

- **Simple**: deployed as a single binary or Docker container, with a flexible API for management.
- **Encrypted**: secrets are encrypted at rest using [age][1] and user-controlled keys.
- **Versioned**: up to nine versions of each secret are stored to provide basic change history.
- **Secure**: built with the [Tailscale][0] client library to provide secure access to the API.
- **External KMS**: optionally integrates with cloud key management systems for securely retrieving [age][1] keys.
- **Cloud Storage**: optionally leverage cloud storage system for backup and recovery of the secrets database.

These features makes `kv2` the perfect secrets management solution for my homelab, but it may not be suitable for production environments.

## 🚀 Quickstart

If you are just looking to move fast and break things, here is the server container running in development mode. No Tailscale, no persistence, and no encryption.

```sh
docker run --rm --name kv2 -p 8081:8081 -e KV2_DEV_MODE=true ghcr.io/hugginsio/kv2:latest
```

You can interact with the server using [the API](docs/api.md) or the provided CLI. You can download the CLI executable from the Releases page or install it with Homebrew:

```sh
brew install hugginsio/tap/kv2
```

## 📚 Documentation

Additional documentation can be found in the [docs](docs) directory.

- [API Introduction](docs/api.md)
- [Command Line Interface](docs/command-line-interface.md)
- [Configuration and Deployment](docs/configure-and-deploy.md)
- [Protocol Reference](docs/protocol.md)

## 🤝🏻 Thanks

- [@tailscale/setec][9], which largely inspired `kv2`.

---

"Tailscale" is a registered trademark of Tailscale Inc. The `kv2` project is not endorsed by, sponsored by, or affiliated with Tailscale Inc.

<!-- Links -->
[0]: https://tailscale.com/
[1]: https://github.com/FiloSottile/age
[9]: https://github.com/tailscale/setec
