# HTTP SMTP Gateway

This is a simple SMTP gateway that sends mails using an HTTP API.

## Usage

The gateway is configured using environment variables:

* `SMTP_HOST` - the SMTP host to use
* `SMTP_USERNAME` - the SMTP username to use
* `SMTP_PASSWORD` - the SMTP password to use
* `AUTH_TOKEN` - the authentication token to use


## Example Request

```http request
POST http://localhost:8080
Content-Type: application/json
token: my-token

{
  "from": "Sender Name <sender@remote.de>",
  "to": "receiver@mail.de",
  "subject": "SUBJECT",
  "body": "HTML BODY"
}
```

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for more info.