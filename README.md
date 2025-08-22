# 🔐 kv2

`kv2` is an end-to-end encrypted secrets manager for [tailnets][0]. It provides a simple API for secrets management.

## Features

- **Encrypted**: secret values encrypted & decrypted client-side using [age][1] and user-controlled keys.
- **Simple**: deploy a single binary (or container), perfectly suited for hyperscale cloud provider's free tiers.
- **Versioned**: update and rollback secret versions with basic change history support.
- **Secure**: built with the [Tailscale][0] client library for secure-by-default API access.
- **External KMS**: optionally integrates with cloud key management systems for securely retrieving [age][1] keys.
- **Cloud Storage**: optionally leverages cloud storage systems for backup and recovery of the secrets database.

## Thanks

- [@tailscale/setec][9], which largely inspired `kv2`.

<!-- Links -->
[0]: https://tailscale.com/
[1]: https://github.com/FiloSottile/age
[9]: https://github.com/tailscale/setec
