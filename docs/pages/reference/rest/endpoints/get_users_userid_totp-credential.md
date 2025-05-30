---
title: "GET /users/[user_id]/totp-credential"
---

# GET /users/[user_id]/totp-credential

Gets a user's TOTP credential.

```
GET https://your-domain.com/users/USER_ID/totp-credential
```

## Response body

Returns the [user TOTP credential model](/reference/rest/models/user-totp-credential) of the credential if it exists.

## Error codes

- [404] `NOT_FOUND`: The user or credential does not exist.
- [500] `UNKNOWN_ERROR`
