# Email Sender API
## Send messages to different emails via API

## Used technologies:
- ####Golang (gin + gRpc)
- ####PostgreSQL
- ####Docker


##### emailSenderApi is an API that let you to send an e-mail from/to any mail with subject and cc's.
##### To send e-mails, I used Simple Message Transport Protocol (smtp)

## Supported protocols

- HTTP **POST/GET** request (via [gin](https://github.com/gin-gonic/gin))
- gRpc **POST/GET** call

## To-do
- [ ] gRpc GET service 
## Done
- [x] REST support (POST, GET)
- [x] gRpc support (POST)
- [x] DataBase logging storage (PostgreSQL)
- [x] Docker
## Usage

###### emailSenderApi works with docker now! requires [Docker](https://www.docker.com/get-started/).
##### Install docker and then run in shell next command:

###For the first launch:
```sh
docker-compose up --build
```
###For every other:
```sh
docker-compose up
```

## RPC

By default the HTTP server will be available on localhost:8080

#### For testing out I used Postman

You can use **POST** via link localhost:8080/sendMsg

##### You need to send JSON in the next format:

| Field   | Type             | Meaning                                     |
|---------|------------------|---------------------------------------------|
| from    | string           | From what e-mail the message should be sent |
| to      | string           | To whom the e-mail should be sent           |
| subject | string           | Subject of the e-mail                       |
| message | string           | Message of the e-mail                       |
| copy    | array of strings | To whom the e-mail should be cc'd           |

Example:
```sh
{"from":"fromExample@gmail.com",
"to":"toExample@gmail.com",
 "subject":"Readme instruction", 
 "message":"Our message",
 "copy": ["toWhom@gmail.com","shouldIcopy@gmail.com"]}
```

You can use **GET** via link localhost:8080/getMsg/*fromMail*  
"fromMail" is a "from" mail, that is getting used as a key in query to our database.
GET will return every message that was sent by that mail

## GRPC

By default the Grpc server will be available on localhost:8081

GRPC is having json-like fields as well as RPC. Fields are almost the same
![img.png](images/img.png)

| Field   | Type             | Meaning                                     |
|---------|------------------|---------------------------------------------|
| from    | string           | From what e-mail the message should be sent |
| to      | string           | To whom the e-mail should be sent           |
| subject | string           | Subject of the e-mail                       |
| **msg** | string           | Message of the e-mail                       |
| **cc**  | array of strings | To whom the e-mail should be cc'd           |

For testing out GRPC I used [evans](https://github.com/ktr0731/evans).

P.S. GET via gRpc is almost done!
```sh
cd emailSenderApi
evans pkg/API/gRpc/proto/parser.proto -p 8081
----------------------------------------------
                EVANS
                
          ------POST------
call Post
'from (TYPE_STRING) =>' fromMail@mail.com
'to (TYPE_STRING) =>' toMail@mail.com
'subject (TYPE_STRING) =>' subject
'msg (TYPE_STRING) =>' Message
'<repeated> cc (TYPE_STRING) =>' toWhomCopy@mail.com

CTRL-D to interrupt cc input prompt
```

### Thanks a lot for your attention. Would be appreciated if you will leave any comment on what I should have done better.
