# 🏗️ Configuration and Deployment

## ⚙️ Configuration

The `kv2` server configuration is controlled through environment variables. The following variables are supported:

| Variable            | Description                                                                                 |
| ------------------- | ------------------------------------------------------------------------------------------- |
| `KV2_DEV_MODE`      | If enabled, the server will use an in-memory database and not attempt a Tailnet connection. |
| `KV2_PRIVATE_KEY`   | The age private key used to decrypt secrets (`AGE-SECRET-KEY*`).                            |
| `KV2_PUBLIC_KEY`    | The age public key used to encrypt secrets (`age1*`).                                       |
| `KV2_TS_AUTHKEY`    | The authentication key used to connect to the Tailnet.                                      |
| `KV2_CLOUD_STORAGE` | The cloud storage provider and path to use for database backup & restore.                   |

### 🔑 External KMS Support

The `kv2` server can automatically fetch secrets from external key management systems for enhanced security and flexibility. The following configuration variables support external KMS references:

- `KV2_PRIVATE_KEY`
- `KV2_PUBLIC_KEY`
- `KV2_TS_AUTHKEY`

The following key management systems are supported:

- **Google Cloud Secret Manager**: `gsm://projects/<project_id>/secrets/<secret_name>` (requires container ADC access)

When a valid KMS prefix is detected, the server will automatically attempt to retrieve the latest version of the secret.

### ☁️ Cloud Storage Support

The `kv2` server can automatically restore the SQLite database. If configured - and no database is found on server startup - the specified provider will attempt to pull down the latest version of `kv2*.db` from the storage location.

While restore operations are automatic, backups are not. You must manually trigger backups using the `/secrets/backup` endpoint.

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
    volumes:
      - kv2db:/root/.config/kv2/
```

Note that this configuration assumes you have your configuration set up in the `.env` file.
