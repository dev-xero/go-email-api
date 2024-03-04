# `GO Email API`

An example of sending emails using Google SMTP in Go

## Sending requests

Before you can send emails, you have to create a `.env` file in the root directory with the same keys as the example env file. Your valid email and passkey (password) is required. For the smtp, using Google's requires the passkey and also a port number, most commonly 465.

### URL:

```url
[POST] http://localhost:8080/email/send
```

### Body:

The sender has to be the email account that you provide the passkey for.

```json
{
    "sender": "usr@root.ssh",
    "recipient": "client@mail.com"
}
```

### Response (if successful):

```json
{
    "message": "Email sent successfully.",
    "code": 200,
    "success": true,
    "payload": null
}
```
