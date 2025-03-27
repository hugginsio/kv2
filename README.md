# 🔐 kv2

> [!WARNING]
> `kv2` is still v0. Breaking changes are expected.

`kv2` is a versioned, encrypted, key-value secrets manager for the homelab. This module provides the server and client components, as well as the reference CLI implementation. The `kv2` server is:

- **Simple**: available as a single binary or Docker container. The only interface is a REST API.
- **Lightweight**: the simplicity of the service allows it to run almost anywhere, such as free tier VMs on hyperscale providers.
- **Secure**: secrets are encrypted at rest, and the server is designed to integrate with [Tailscale][0] for controlling API access.

While these traits make `kv2` perfect for my homelab, it may not be suitable for production environments.

## 🚀 Quickstart

If you are just looking to move fast and break things, here is the server container running in development mode. No Tailscale, no persistence, and no encryption.

```sh
docker run --rm --name kv2 -p 80:8080 -e KV2_DEV_MODE=true ghcr.io/hugginsio/kv2:latest
```

## 📚 Documentation

Additional documentation can be found in the [docs](docs) directory.

- [API Reference](docs/api.md)
- [Configuration and Deployment](docs/configure-and-deploy.md)

## 🤝🏻 Thanks

- [@tailscale/setec][9], which largely inspired `kv2`.

---

"Tailscale" is a registered trademark of Tailscale Inc. The `kv2` project is not endorsed by, sponsored by, or affiliated with Tailscale Inc.

<!-- Links -->
[0]: https://tailscale.com/
[9]: https://github.com/tailscale/setec
