# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [secrets/v1/api.proto](#secrets_v1_api-proto)
    - [BackupRequest](#secrets-v1-BackupRequest)
    - [BackupResponse](#secrets-v1-BackupResponse)
    - [CreateSecretRequest](#secrets-v1-CreateSecretRequest)
    - [CreateSecretResponse](#secrets-v1-CreateSecretResponse)
    - [DeleteSecretRequest](#secrets-v1-DeleteSecretRequest)
    - [DeleteSecretResponse](#secrets-v1-DeleteSecretResponse)
    - [GetSecretRequest](#secrets-v1-GetSecretRequest)
    - [GetSecretResponse](#secrets-v1-GetSecretResponse)
    - [ListSecretsRequest](#secrets-v1-ListSecretsRequest)
    - [ListSecretsResponse](#secrets-v1-ListSecretsResponse)
    - [RevertSecretRequest](#secrets-v1-RevertSecretRequest)
    - [RevertSecretResponse](#secrets-v1-RevertSecretResponse)
    - [Secret](#secrets-v1-Secret)
    - [SecretMetadata](#secrets-v1-SecretMetadata)
    - [UpdateSecretRequest](#secrets-v1-UpdateSecretRequest)
    - [UpdateSecretResponse](#secrets-v1-UpdateSecretResponse)

    - [Kv2Service](#secrets-v1-Kv2Service)

- [Scalar Value Types](#scalar-value-types)



<a name="secrets_v1_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## secrets/v1/api.proto



<a name="secrets-v1-BackupRequest"></a>

### BackupRequest
Request message for `secrets.v1.Kv2Service/Backup`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional | The name of the backup. Defaults to `kv2.db`. |






<a name="secrets-v1-BackupResponse"></a>

### BackupResponse
Empty message. Check for an error code to determine success.






<a name="secrets-v1-CreateSecretRequest"></a>

### CreateSecretRequest
Request message for `secrets.v1.Kv2Service/CreateSecret`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The plain text key of the secret. |
| value | [bytes](#bytes) |  | The encoded value of the secret. |






<a name="secrets-v1-CreateSecretResponse"></a>

### CreateSecretResponse
Response message for `secrets.v1.Kv2Service/CreateSecret`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [SecretMetadata](#secrets-v1-SecretMetadata) |  | The metadata of the created secret. |






<a name="secrets-v1-DeleteSecretRequest"></a>

### DeleteSecretRequest
Request message for `secrets.v1.Kv2Service/DeleteSecret`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The plain text key of the secret. |






<a name="secrets-v1-DeleteSecretResponse"></a>

### DeleteSecretResponse
Empty message. Check for an error code to determine success.






<a name="secrets-v1-GetSecretRequest"></a>

### GetSecretRequest
Request message for `secrets.v1.Kv2Service/GetSecret`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The plain text key of the secret. |






<a name="secrets-v1-GetSecretResponse"></a>

### GetSecretResponse
Response message for `secrets.v1.Kv2Service/GetSecret`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets-v1-Secret) |  | The requested secret. |






<a name="secrets-v1-ListSecretsRequest"></a>

### ListSecretsRequest
Empty message. No request message needed to list secrets.






<a name="secrets-v1-ListSecretsResponse"></a>

### ListSecretsResponse
Response message for `secrets.v1.Kv2Service/ListSecrets`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secrets | [SecretMetadata](#secrets-v1-SecretMetadata) | repeated | The secret metadata for all available secrets. |






<a name="secrets-v1-RevertSecretRequest"></a>

### RevertSecretRequest
Request message for `secrets.v1.Kv2Service/RevertSecret`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The plain text key of the secret. |






<a name="secrets-v1-RevertSecretResponse"></a>

### RevertSecretResponse
Request message for `secrets.v1.Kv2Service/RevertSecret`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [uint32](#uint32) |  | The current version of the reverted secret. |






<a name="secrets-v1-Secret"></a>

### Secret
Secret represents a single secret version.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The plain text key of the secret. |
| value | [bytes](#bytes) |  | The encrypted value of the secret. |
| version | [uint32](#uint32) |  | The version of the represented secret. |






<a name="secrets-v1-SecretMetadata"></a>

### SecretMetadata
SecretMetadata represents a secret and all its versions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The plain text key of the secret. |
| version | [uint32](#uint32) | repeated | All versions of the secret. |






<a name="secrets-v1-UpdateSecretRequest"></a>

### UpdateSecretRequest
Request message for `secrets.v1.Kv2Service/UpdateSecret`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The plain text key of the secret. |
| value | [bytes](#bytes) |  | The encoded value of the secret. |






<a name="secrets-v1-UpdateSecretResponse"></a>

### UpdateSecretResponse
Response message for `secrets.v1.Kv2Service/UpdateSecret`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [uint32](#uint32) |  | The version of the updated secret. |












<a name="secrets-v1-Kv2Service"></a>

### Kv2Service
The Kv2 service provides an encrypted key-value store with versioning and backup capabilities.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateSecret | [CreateSecretRequest](#secrets-v1-CreateSecretRequest) | [CreateSecretResponse](#secrets-v1-CreateSecretResponse) |  |
| GetSecret | [GetSecretRequest](#secrets-v1-GetSecretRequest) | [GetSecretResponse](#secrets-v1-GetSecretResponse) |  |
| UpdateSecret | [UpdateSecretRequest](#secrets-v1-UpdateSecretRequest) | [UpdateSecretResponse](#secrets-v1-UpdateSecretResponse) |  |
| DeleteSecret | [DeleteSecretRequest](#secrets-v1-DeleteSecretRequest) | [DeleteSecretResponse](#secrets-v1-DeleteSecretResponse) |  |
| RevertSecret | [RevertSecretRequest](#secrets-v1-RevertSecretRequest) | [RevertSecretResponse](#secrets-v1-RevertSecretResponse) |  |
| ListSecrets | [ListSecretsRequest](#secrets-v1-ListSecretsRequest) | [ListSecretsResponse](#secrets-v1-ListSecretsResponse) |  |
| Backup | [BackupRequest](#secrets-v1-BackupRequest) | [BackupResponse](#secrets-v1-BackupResponse) |  |





## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
