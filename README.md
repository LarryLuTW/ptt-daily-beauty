## Development Setup

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

### Setup GCP App Engine

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

### Go into deploy environment

```
docker-compose up -d --build
docker exec -it PTTDB_deploy_en bash
```

### Deploy app to App Engine(about 1 ~ 2 mins)

```bash
gcloud app deploy -q
gcloud app browse # open with browser
```

### Deploy cron job to App Engine

```bash
gcloud app deploy cron.yaml -q
```
