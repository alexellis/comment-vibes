## comment-vibes

Figure out the vibe of your community through emoji comments.

## Demo

[Tester issue](https://github.com/teamserverless/proposals/issues/1)

## Create a postgres DB

You can use [DigitalOcean](https://digitalocean.com) managed Postgres for example.

## Populate it with the schema

```
psql postgresql://connection-string here
```

Paste in [schema.sql](schema.sql)

## Create your secrets:

```bash
export USERNAME=""
export PASSWORD=""
export HOST=""
export WEBHOOK_SECRET=""

faas-cli secret create username --from-literal $USERNAME
faas-cli secret create password --from-literal $PASSWORD
faas-cli secret create host --from-literal $HOST
faas-cli secret create webhook-secret --from-literal $WEBHOOK_SECRET
```

## Deploy to OpenFaaS

```
# Get the additional template
faas-cli template store pull golang-middleware

# Deploy

faas-cli deploy
```

## Rebuild and deploy

```
sed -i stack.yml s/alexellis2/your-docker-hub/g

faas-cli up
```
