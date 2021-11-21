# Email Sender API
## Send messages to different emails via API


###### emailSenderApi is an API that let you to send an e-mail from/to any mail with subject and cc's.
###### To send e-mails, I used Simple Message Transport Protocol (smtp)

## Supported protocols

- HTTP POST request (via [gin](https://github.com/gin-gonic/gin)) 

## Working on

- gRpc
- DataBase logging storage (PostgreSQL)

## Planned

- Docker

## Usage

###### emailSenderApi requires [Golang](https://golang.org/) v1.6+ to run.

##### Install the dependencies and start the server.

```sh
cd emailSenderAPI/src
go run main.go
```

By default the server will be runned on localhost:8080

#### For testing out I used Postman

You can use API via link localhost:8080/sendmsg

##### You need to send JSON in the next format:

| Field | Type | Meaning |
| ------ | ------ | ------ |
| from | string | From what e-mail the message should be sent |
| to | string | To whom the e-mail should be sent |
| subject | string| Subject of the e-mail |
| message |  string| Message of the e-mail |
| copy | array of strings | To whom the e-mail should be cc'd |

Example:
```sh
{"from":"fromExample@gmail.com",
"to":"toExample@gmail.com",
 "subject":"Readme instruction", 
 "message":"Our message",
 "copy": ["toWhom@gmail.com","shouldIcopy@gmail.com"]}
```

### Thanks a lot for your attention. Would be appreciated if you will leave any comment on what I should have done better.


