## Development Setup

### GCP App Engine

put your app engine credential here

```js
// deploy-en/credential.json
{
    "type": "service_account",
    "project_id": "YOUR_PROJECT_ID",
    "private_key_id": "YOUR_PRIVATE_KEY_ID",
    "private_key": "YOUR_PRIVATE_KEY",
    // ...
}
```

### Mongodb

```go
// db.config
package db

const (
    user = "YOUR_MONGO_USER"
    pass = "YOUR_MONGO_PASSWORD"
    host = "YOUR_MONGO_HOST"
    port = 27017
)
```

### AWS SES

```go
// mail/config.go
package mail

const (
    user = "YOUR_SES_USER"
    pwd  = "YOUR_SES_PASSWORD"
    host = "YOUR_SES_HOST"
    port = 587
)
```



## Deployment

#### First, go into deploy environment

```
docker-compose up -d --build
docker exec -it mail_sender_deploy_en bash
```

### app

- deploy to app engine(about 1 ~ 2 mins)

```bash
gcloud app deploy -q
```

- open with browser

```bash
gcloud app browse
```

wait for deploy finish and goto [https://daily-beauty-209105.appspot.com/test](https://daily-beauty-209105.appspot.com/test)

- get logs

```
gcloud app logs tail -s default
```

### cron job

```bash
gcloud app deploy cron.yaml -q
```

goto [https://console.cloud.google.com/appengine/taskqueues/cron?project=daily-beauty-209105](https://console.cloud.google.com/appengine/taskqueues/cron?project=daily-beauty-209105)
