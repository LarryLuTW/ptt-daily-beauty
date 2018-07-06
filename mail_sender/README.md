## Development

#### initialize

- go into project

```bash
cd mail_sender
```

- run develop environment in container

```bash
docker-compose up -d --build
```

- attach into container

```bash
docker exec -it mail_sender_app_engine_py_dev_1 bash
```

- install new dependencies

```bash
pip install flask
pip freeze > requirements.txt
```

## Deployment

- go into container

