## Crossplane instructions

> Instructions contributed by [Daniel Mangum](https://github.com/hasheddan). For any issues please contact Daniel or Crossplane.

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
