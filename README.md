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

##### Send Message
```sh
$ omg run get_bot -e ACCESS_TOKEN=<ACCESS_TOKEN>
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
