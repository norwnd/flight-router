# Flight calculator

This is a simple micro-service to calculate flights. This service is built with [Encore.dev](https://encore.dev/) - 
see more about it (generated interactive API docs, how to run it, etc.) in the sections below.

## Demo API

Demo API docs and Admin dashboard are available at [Encore staging environment](https://app.encore.dev/flight-router-c9fi/envs/staging/api/flight) (requires Admin/Member access) to preview and experiment with.

Endpoints however are publicly available:

```bash
curl 'https://staging-flight-router-c9fi.encr.app/calculate' -d '{"flights":[["IND","EWR"],["SFO","ATL"],["GSO","IND"],["ATL","GSO"]]}'
```

## Running locally

```bash
encore run --port=8080
```

## Open the developer dashboard

While `encore run` is running, open [http://localhost:9400/](http://localhost:9400/) to access Encore's [local developer dashboard](https://encore.dev/docs/observability/dev-dash).

Here you can see API docs, make requests in the API explorer, and view traces of the responses.

## Using the API

To see that your app is running, you can try this Curl command:

```bash
curl 'http://127.0.0.1:8080/calculate' -d '{"flights":[["IND","EWR"],["SFO","ATL"],["GSO","IND"],["ATL","GSO"]]}'
```

## Deployment

Deploy your application to a staging environment in Encore's free development cloud:

```bash
git add -A .
git commit -m 'Commit message'
git push encore
```

Then head over to the [Cloud Dashboard](https://app.encore.dev) to monitor your deployment and find your production URL.

From there you can also connect your own AWS or GCP account to use for deployment.

## Testing

```bash
encore test ./...
```