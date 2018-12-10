## [Daily Beauty - 表特日報](https://daily-beauty.xyz)

Daily Beauty 每晚十一點會自動蒐集 PTT 表特版前三名

這些都是 PTT 鄉民篩選過的，想要的話就快點到[網站](https://daily-beauty.xyz)上訂閱吧～～～

![](https://i.imgur.com/RdNBuie.png)

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

### Deploy app to App Engine(about 1 ~ 2 mins)

```bash
make deploy-app
```

### Deploy cron job to App Engine

```bash
make deploy-cron
```
