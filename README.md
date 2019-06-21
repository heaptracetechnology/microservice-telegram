# _Telegram_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.org/heaptracetechnology/microservice-telegram.svg?branch=master)](https://travis-ci.org/heaptracetechnology/microservice-telegram)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-telegram/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-telegram)

An OMG service for Telegram, is a cloud-based instant messaging service.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Get Bot Details
```coffee
>>> telegram getBot
```
##### Send Message TO Group/User
```coffee
>>> telegram send chatId:'chatID' message:'messageText'
```
##### Send Bot Channel Message
```coffee
>>> telegram channelMessage username:'username' message:'messageText'
```
##### Get Chat
```coffee
>>> telegram getChat chatId:'chatID'
```
##### Leave Chat
```coffee
>>> telegram leaveChat chatId:'chatID'
```
##### Send Photo
```coffee
>>> telegram sendPhoto chatId:'chatID' image:'Base64 Data'
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

**Note** : Use "-" as prefix in ChatID for group(chat_id = "-12345678") else for user (chat_id = "12345678")

##### Get Bot Details
```shell
$ omg run getBot -e BOT_TOKEN=<BOT_TOKEN>
```
##### Send Message TO Group/User
```shell
$ omg run send -a chatId=<CHAT_ID> -a message=<MESSAGE> -e BOT_TOKEN=<BOT_TOKEN>
```
##### Send Bot Channel Message
```shell
$ omg run channelMessage -a username=<USERNAME> -a message=<MESSAGE> -e BOT_TOKEN=<BOT_TOKEN>
```
##### Send Bot Channel Message (EXAMPLE)
```shell
$ omg run channelMessage -a username="@firstchannel" -a message="Hello World" -e BOT_TOKEN=<BOT_TOKEN>
```
##### Get Chat
```shell
$ omg run getChat -a chatId=<CHAT_ID> -e BOT_TOKEN=<BOT_TOKEN>
```
##### Leave Chat
```shell
$ omg run leaveChat -a chatId=<CHAT_ID> -e BOT_TOKEN=<BOT_TOKEN>
```
##### Send Photo
```shell
$ omg run sendPhoto -a chatId=<CHAT_ID> -a image=<BASE64_DATA> -e BOT_TOKEN=<BOT_TOKEN>
```
##### Subscribe
```shell
$ omg subscribe bot hears -a channel=<CHANNEL_USERNAME> -e BOT_TOKEN=<BOT_TOKEN>
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/telegram/blob/master/LICENSE).
