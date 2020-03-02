# Admission Revealer
Kubernetes admission webhook application that logs full AdmissionReview request and always responds with allow. This is to be used as an admission controller development and debugging tool.

The need for this tool arose from struggles debugging an admission webhook with admission review request coming from a builtin controller, in this case, the replicaset controller. The issue came from the assumption that when the replicaset controller created a new pod, the pod's namepsace field is set in its metadata (it does not but is in the namespace field of the admission review). There wasn't an easy way to know this from documentation alone. With this tool, if a webhook isn't functioning as expected, you can copy the entire admission review request from this application's log and make the exact same request to your webhook running locally in your IDE's debugger. 

This approach is slightly more graceful than having to rebuild your webhook to log the requests and can be re-used to debug other webhooks (just scale the deployment to zero while not in use). Update the MutatingWebhookConfiguration manifests to select the type of resource you want admission review request logged for.

---

## Build Image
```make docker-build```

## Deploy
After pushing your image to your registry, update the image field in config/deployment.yaml

Note the config/ directory includes a kcertifier manifest (https://github.com/att-cloudnative-labs/kcertifier). This is used to automate creation of the TLS certificate. If you don't have kcertifier installed and/or prefer to use an alternative, deploy the other manifests manually without using make.

```make deploy```