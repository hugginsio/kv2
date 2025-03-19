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

## 🛠️ Deployment

The `kv2` server is available on the GitHub Container Registry, at `ghcr.io/hugginsio/kv2`. While the `:latest` tag is available, you should really use a version tag (or even better, a hash) to ensure stability and security.

The best way to use `kv2` is most likely going to involve Docker Compose. Here is an example:

```yaml
name: "kv2"
services:
  server:
    image: "ghcr.io/hugginsio/kv2:latest"
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "--max-time", "2", "http://localhost:8080/health"]
    env_file:
      - .env
```

Note that this configuration assumes you have your configuration set up in the `.env` file.

### ⚙️ Configuration

The `kv2` server configuration is controlled through environment variables. The following variables are supported:

| Variable          | Description                                                                                 | Default Value |
| ----------------- | ------------------------------------------------------------------------------------------- | ------------- |
| `KV2_DEV_MODE`    | If enabled, the server will use an in-memory database and not attempt a Tailnet connection. | `false`       |
| `KV2_PRIVATE_KEY` | The age private key used to decrypt secrets (`AGE-SECRET-KEY*`).                            | `""`          |
| `KV2_PUBLIC_KEY`  | The age public key used to decrypt secrets (`age1*`).                                       | `""`          |
| `KV2_TS_AUTHKEY`  | The authentication key used to connect to the Tailnet.                                      | `""`          |

#### 🔑 External KMS Support

The `kv2` server can automatically fetch secrets from external key management systems for enhanced security and flexibility. The following configuration variables support external KMS references:

- `KV2_PRIVATE_KEY`
- `KV2_PUBLIC_KEY`
- `KV2_TS_AUTHKEY`

The following key management systems are supported:

- **Google Cloud Secret Manager**: `gsm://projects/<project_id>/secrets/<secret_name>`

When a valid KMS prefix is detected, the server will automatically attempt to retrieve the latest version of the secret.

## 🤝🏻 Thanks

- [@tailscale/setec][9], which largely inspired `kv2`.

---

"Tailscale" is a registered trademark of Tailscale Inc. The `kv2` project is not endorsed by, sponsored by, or affiliated with Tailscale Inc.

<!-- Links -->
[0]: https://tailscale.com/
[9]: https://github.com/tailscale/setec
