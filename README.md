# dailt-beauty

## PTT Crawler

- Nodejs or Go
- Cloud Function

### development

#### initialize

- go into ptt_crawler

```bash
cd ptt_crawler
```

- run develop environment in container

```bash
docker-compose up -d --build
```

- attach into container

```bash
docker exec -it ptt_crawler_functions_dev_1 bash
```

- install dependencies

```bash
npm i
```

- start functions-emulator

```bash
functions start
```

#### run in emualator

- deploy to emulator

```bash
functions deploy helloGET --trigger-http
```

- testing in container

```bash
functions call helloGET --data='{"message":"Hello World"}'
```

- testing from host
```
curl http://localhost:8010/daily-beauty-209105/us-central1/helloGET
```

#### reference

- [Cloud Functions Documentation  |  Google Cloud](https://cloud.google.com/functions/docs/emulator)

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
