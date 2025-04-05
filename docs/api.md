# 📃 API

The `kv2` server exposes a multi-protocol API via [Connect RPC](https://connectrpc.com/). It is available over port `443` on your Tailnet while running in production, and port `8081` on your local machine when running in development mode. While running in production mode, port `8081` will not respond to requests.

When running in production mode, `kv2` assumes that you will be connected to a Tailnet with MagicDNS and HTTPS support enabled.

You can interact with the `kv2` API via gRPC, REST+Protobuf, or REST+JSON. The recommended way to interact with the API is with gRPC or REST+Protobuf via the `git.huggins.io/kv2/api/secrets/v1/secretsv1connect` client library. The specification is contained in [/proto/secrets/v1/api.proto](../proto/secrets/v1/api.proto) and described in the [Protocol documentation](./protocol.md). This documentation is useful for all consumers, but is particularly useful for gRPC or REST+Protobuf clients.

If you are using gRPC or REST+Protobuf, the recommended way to interact with the API is with the `git.huggins.io/kv2/api/secrets/v1/secretsv1connect` client library. For users who require REST+JSON - such as those using `kv2` with [External Secrets Operator](https://external-secrets.io/latest/provider/webhook/) - you can access the API via `<host>/secrets.v1.Kv2Service/<method_name>`. For example:

- `POST https://kv2.yak-bebop.ts.net/secrets.v1.Kv2Service/GetSecret`

---

# 📃 API

The `kv2` server exposes a REST API for secrets management. It's available over port `80` over your Tailnet (or locally, for development mode). When connected to a Tailnet, port `80` on the container will not respond to requests.

All errors from the API will bear a `Content-Type` header of `text/plain` and a non-200 status code. The response body will contain a human-readable error message.

## `GET /secrets`

List all secrets and the available versions.

### Request

- Content type: none

### Response

- Content type: none
- Status code: 200

```json
[
  {
    "Key": "sample-secret-1",
    "Versions": [
      1
    ]
  },
  {
    "Key": "sample-secret-2",
    "Versions": [
      1,
      2
    ]
  }
]
```

## `POST /secrets/create`

Create a new secret.

### Request

- Content type: `application/json`

```json
{
  "Key": "sample-secret-3",
  "Value": "RXllcyB1cCwgR3VhcmRpYW4u"
}
```

### Response

- Content type: none
- Status code: 201

## `POST /secrets/read`

Read the latest version of a secret.

### Request

- Content type: `application/json`

```json
{
  "Key": "sample-secret-1"
}
```

### Response

- Content type: `application/json`
- Status code: 200

```json
{
  "Key": "sample-secret-1",
  "Value": "UmVtZW1iZXIgUmVhY2g=",
  "Version": 1
}
```

## `POST /secrets/update`

Create a new version of an existing secret.

### Request

- Content type: `application/json`

```json
{
  "Key": "sample-secret-2",
  "Value": "U2t5IEFib3ZlLCBWb2ljZSBXaXRoaW4="
}
```

### Response

- Content type: none
- Status code: 200

## `DELETE /secrets/delete`

Delete a secret and all its versions.

### Request

- Content type: `application/json`

```json
{
  "Key": "sample-secret-3"
}
```

### Response

- Content type: none
- Status code: 200

## `POST /secrets/revert`

Revert a secret to a previous version.

### Request

- Content type: `application/json`

```json
{
  "Key": "sample-secret-2"
}
```

### Response

- Content type: none
- Status code: 200

## `POST /secrets/backup`

Create a backup of the secrets database. The request body is optional, and the name of the backup will default to `kv2.db`.

### Request (optional)

- Content type: `application/json`

```json
{
  "Name": "kv2_new_final.db"
}
```

### Response

- Content type: none
- Status code: 200
