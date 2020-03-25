## comment-vibes

This is a serverless CRUD sample project with Postgres and GitHub emoji integration. You can use it to tally up the emojis used in various comments across your GitHub projects.

## Kubernetes / Serverless example project

This sample uses the OpenFaaS [golang-middleware](https://github.com/openfaas-incubator/golang-http-template/) template and three endpoints:

* view - to render a HTML template (server-side)
* import-comment - to receive and validate the webhook, then update postgres
* comments - to render JSON from a postgres function

You can see the results in the demo below. Feel free to try it out, if the endpoint is still up.

The code is adapted from the [Serverless Single Page App - aka Open Source leaderboard](https://github.com/alexellis/leaderboard-app/)

## Demo

Example output view the `/view` function:

![Example screenshot](docs/example.png)

Example input from a GitHub issue and webhook:

![Example comments](docs/comments.png)

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

## Deploy to OpenFaaS (Intel)

```
# Get the additional template
faas-cli template store pull golang-middleware

# Deploy

faas-cli deploy
```

## Rebuild and deploy (Intel and ARM)

```
sed -i stack.yml s/alexellis2/your-docker-hub/g
export DOCKER_BUILDKIT=1

faas-cli up --tag=sha
```

## Setup GitHub webhooks

Go to Settings for your repo and click "Webhooks"

Add a webhook

![webhook](docs/webhook.png)

Pick only the issue comments event, there is no issue for reactions at this time.

![tick](docs/tick.png)

Now have someone send a comment to one of the issues in your repo with an emoji i.e. 👍

If you need to create a HTTP tunnel to your local computer, try [inlets and inletsctl](https://docs.inlets.dev/)

View the result on the `/view` function.

## Demo - with [Crossplane](https://crossplane.io/)

1. Install Crossplane

```
arkade install crossplane --helm3
```

2. Install GCP Provider

```
kubectl apply -f gcp.yaml
```

3. Set up Credentials

This guide assumes you have set up a service account with Admin permissions and have included the credentials in a JSON file named `service-account.json` in the `crossplane/` directory.

```
cd crossplane

./credentials.sh <your-gcp-project-name>
```

4. Create CloudSQL Class

```
kubectl apply -f crossplane/cloudsql.yaml
```

5. Create Postgres Claim

```
kubectl apply -f crossplane/postgres.yaml
```

6. Install and Configure OpenFaaS

```
arkade install openfaas

kubectl port-forward -n openfaas svc/gateway 8080:8080 &

PASSWORD=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)                                                                                    
echo -n $PASSWORD | faas-cli login --username admin --password-stdin
```

7. Set up Secrets

```
export USERNAME=$(kubectl -n openfaas-fn get secret pgconn -o=jsonpath={.data.username} | base64 --decode -)
export PASSWORD=$(kubectl -n openfaas-fn get secret pgconn -o=jsonpath={.data.password} | base64 --decode -)
export HOST=$(kubectl -n openfaas-fn get secret pgconn -o=jsonpath={.data.endpoint} | base64 --decode -)
export WEBHOOK_SECRET="mysupersecretvalue"

faas-cli secret create username --from-literal=$USERNAME
faas-cli secret create password --from-literal=$PASSWORD
faas-cli secret create host --from-literal=$HOST
faas-cli secret create webhook-secret --from-literal=$WEBHOOK_SECRET
```

8. Configure Database

```
PGPASSWORD=$(echo $PASSWORD) psql -h $HOST -p 5432 -U $USERNAME -c 'CREATE DATABASE defaultdb;'
PGPASSWORD=$(echo $PASSWORD) psql -h $HOST -p 5432 -U $USERNAME defaultdb < schema.sql
```

9. Set up an Exit Node with [Inlets](https://github.com/inlets/inlets) on GCP

```
curl -sLSf https://inletsctl.inlets.dev | sudo sh
curl -sLS https://get.inlets.dev | sudo sh

export PROJECTID=<your-gcp-project-name>

inletsctl create -p gce --project-id=$PROJECTID -f=crossplane/service-account.json
```

Make sure to execute the commands returned to start the inlets client.

9. Deploy Functions

```
faas-cli template store pull golang-middleware

sed -i 's/postgres_port: 25060/postgres_port: 5432/g' stack.yml

faas-cli deploy
```

10. Set up Github Webhook

See directions above.

11. Comment with emoji!
