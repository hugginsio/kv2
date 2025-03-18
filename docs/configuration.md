# ⚙️ Configuration

The `kv2` server configuration is controlled through environment variables. The following variables are supported:

| Variable          | Description                                                                                 | Default Value |
| ----------------- | ------------------------------------------------------------------------------------------- | ------------- |
| `KV2_DEV_MODE`    | If enabled, the server will use an in-memory database and not attempt a Tailnet connection. | `false`       |
| `KV2_PRIVATE_KEY` | The age private key used to decrypt secrets (`AGE-SECRET-KEY*`).                            | `""`          |
| `KV2_PUBLIC_KEY`  | The age public key used to decrypt secrets (`age1*`).                                       | `""`          |
| `KV2_TS_AUTHKEY`  | The authentication key used to connect to the Tailnet.                                      | `""`          |
