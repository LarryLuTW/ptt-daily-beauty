# dailt-beauty

## PTT Crawler

- Nodejs or Go
- Cloud Function

### development

#### initialize

- `cd ptt_crawler`
- run develop environment in container: `docker-compose up -d --build`
- attach into container: `docker exec -it ptt_crawler_functions_dev_1 bash`
- install dependencies: `npm i`
- start functions-emulator: `functions start`

#### run in emualator

- deploy to emulator: `functions deploy helloGET --trigger-http`
- testing in container: `functions call helloGET --data='{"message":"Hello World"}'`
- testing in host: ``

---

## Mail Sender

- Gmail API
- Cloud Function

## Database

- Datastore

## Subscribe API

- PubSub(optional)
- Cloud Function

## Subscribe Website

- Github Page
