# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [nexus/v2/api.proto](#nexus_v2_api-proto)
    - [CreateSecretRequest](#nexus-v2-CreateSecretRequest)
    - [CreateSecretResponse](#nexus-v2-CreateSecretResponse)
    - [GetSecretByTitleRequest](#nexus-v2-GetSecretByTitleRequest)
    - [GetSecretByTitleResponse](#nexus-v2-GetSecretByTitleResponse)
    - [GetSecretVersionByTitleRequest](#nexus-v2-GetSecretVersionByTitleRequest)
    - [GetSecretVersionByTitleResponse](#nexus-v2-GetSecretVersionByTitleResponse)
    - [ListSecretInner](#nexus-v2-ListSecretInner)
    - [ListSecretRequest](#nexus-v2-ListSecretRequest)
    - [ListSecretResponse](#nexus-v2-ListSecretResponse)
    - [UpdateSecretRequest](#nexus-v2-UpdateSecretRequest)
    - [UpdateSecretResponse](#nexus-v2-UpdateSecretResponse)

    - [Kv2NexusService](#nexus-v2-Kv2NexusService)

- [relay/v2/api.proto](#relay_v2_api-proto)
- [Scalar Value Types](#scalar-value-types)



<a name="nexus_v2_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nexus/v2/api.proto



<a name="nexus-v2-CreateSecretRequest"></a>

### CreateSecretRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| content | [bytes](#bytes) |  |  |
| public_key | [string](#string) |  |  |






<a name="nexus-v2-CreateSecretResponse"></a>

### CreateSecretResponse







<a name="nexus-v2-GetSecretByTitleRequest"></a>

### GetSecretByTitleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |






<a name="nexus-v2-GetSecretByTitleResponse"></a>

### GetSecretByTitleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| versions | [uint32](#uint32) | repeated |  |
| created_by | [string](#string) |  |  |
| created_at | [string](#string) |  |  |






<a name="nexus-v2-GetSecretVersionByTitleRequest"></a>

### GetSecretVersionByTitleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| version | [uint32](#uint32) | optional |  |






<a name="nexus-v2-GetSecretVersionByTitleResponse"></a>

### GetSecretVersionByTitleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [uint32](#uint32) | optional |  |
| content | [bytes](#bytes) |  |  |
| created_by | [string](#string) |  |  |
| created_at | [string](#string) |  |  |






<a name="nexus-v2-ListSecretInner"></a>

### ListSecretInner



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| versions | [uint32](#uint32) | repeated |  |
| created_by | [string](#string) |  |  |
| created_at | [string](#string) |  |  |






<a name="nexus-v2-ListSecretRequest"></a>

### ListSecretRequest







<a name="nexus-v2-ListSecretResponse"></a>

### ListSecretResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secrets | [ListSecretInner](#nexus-v2-ListSecretInner) | repeated |  |






<a name="nexus-v2-UpdateSecretRequest"></a>

### UpdateSecretRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| content | [bytes](#bytes) |  |  |
| public_key | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="nexus-v2-UpdateSecretResponse"></a>

### UpdateSecretResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [uint32](#uint32) |  |  |












<a name="nexus-v2-Kv2NexusService"></a>

### Kv2NexusService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListSecret | [ListSecretRequest](#nexus-v2-ListSecretRequest) | [ListSecretResponse](#nexus-v2-ListSecretResponse) | List all secrets. |
| CreateSecret | [CreateSecretRequest](#nexus-v2-CreateSecretRequest) | [CreateSecretResponse](#nexus-v2-CreateSecretResponse) | Create a new secret. |
| UpdateSecret | [UpdateSecretRequest](#nexus-v2-UpdateSecretRequest) | [UpdateSecretResponse](#nexus-v2-UpdateSecretResponse) | rpc GetSecretByTitle(GetSecretByTitleRequest) returns (GetSecretByTitleResponse); // Get a secret details by title. rpc GetSecretVersionByTitle(GetSecretVersionByTitleRequest) returns (GetSecretVersionByTitleResponse); // Get a particular secret version by title.

Update an existing secret. |





<a name="relay_v2_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## relay/v2/api.proto












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
