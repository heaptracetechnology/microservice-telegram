# Telegram as a microservice
An OMG service for Telegram, is a cloud-based instant messaging service.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)


## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Get Bot Details
```sh
$ omg run get_bot -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Send Bot Group Message
```sh
$ omg run send -a chat_id=<CHAT_ID> -a message=<MESSAGE> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Send Bot Group Message (EXAMPLE)
```sh
$ omg run group_message -a chat_id="-349200002" -a message="Hello World" -e ACCESS_TOKEN=<ACCESS_TOKEN>
```

##### Send Bot Channel Message
```sh
$ omg run channel_message -a username=<USERNAME> -a message=<MESSAGE> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Send Bot Channel Message (EXAMPLE)
```sh
$ omg run channel_message -a username="@firstchannel" -a message="Hello World" -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Get Chat
```sh
$ omg run get_chat -a chat_id=<CHAT_ID> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Leave Chat
```sh
$ omg run leave_chat -a chat_id=<CHAT_ID> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-telegram .
```
### RUN
```
docker run -p 3000:3000 microservice-telegram
```
