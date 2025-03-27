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
