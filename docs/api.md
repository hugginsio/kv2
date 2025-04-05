# 📃 API

The `kv2` server exposes a multi-protocol API via [Connect RPC](https://connectrpc.com/). It is available over port `443` on your Tailnet while running in production, and port `8081` on your local machine when running in development mode. While running in production mode, port `8081` will not respond to requests. Additionally, `kv2` assumes that you will be connected to a Tailnet with [MagicDNS](https://tailscale.com/kb/1081/magicdns) and [HTTPS support](https://tailscale.com/kb/1153/enabling-https) enabled. It will attempt to reserve the `kv2` hostname within your Tailnet, and will be accessible at `https://k2.tailnet-name.ts.net`.

You can interact with the `kv2` API via gRPC, REST+Protobuf, or REST+JSON. The recommended way to interact with the API is with gRPC or REST+Protobuf via the `git.huggins.io/kv2/api/secrets/v1/secretsv1connect` client library. The specification is contained in [/proto/secrets/v1/api.proto](../proto/secrets/v1/api.proto) and described in the [Protocol documentation](./protocol.md). This documentation is useful for all consumers, but is particularly useful for gRPC or REST+Protobuf clients.

If you are using gRPC or REST+Protobuf, the recommended way to interact with the API is with the `git.huggins.io/kv2/api/secrets/v1/secretsv1connect` client library. For users who require REST+JSON - such as those using `kv2` with [External Secrets Operator](https://external-secrets.io/latest/provider/webhook/) - you can access the API via `<host>/secrets.v1.Kv2Service/<method_name>`. For example:

- `POST https://kv2.yak-bebop.ts.net/secrets.v1.Kv2Service/GetSecret`
- `POST https://kv2.yak-bebop.ts.net/secrets.v1.Kv2Service/Backup`
