# CronJob Controller

This is just a sample controller for K8s CronJob Resource. This controller, built with Kubebuilder handles CronJob resource. This controller was build along the [Kubebuilder Book](https://book.kubebuilder.io/introduction.html)

# Use

```
$ make install
$ make run
$ kubectl apply -f config/samples/batch_v1_cronjob.yaml
```

Note: you need to generate certificates for serving the webhooks.