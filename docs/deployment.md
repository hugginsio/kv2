# 🚀 Deployment

The `kv2` server is available on the GitHub Container Registry, at `ghcr.io/hugginsio/kv2`. While the `:latest` tag is available, you should really use a version tag (or even better, a hash) to ensure stability and security.

The best way to use `kv2` is most likely going to involve Docker Compose. Here is an example based off of [deployment/compose.yml](./deployment/compose.yml):

```yaml
# Compose file for kv2
# yaml-language-server: $schema=https://raw.githubusercontent.com/compose-spec/compose-spec/master/schema/compose-spec.json

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

Note that this configuration assumes you have your [configuration](./configuration.md) set up in the `.env` file.
