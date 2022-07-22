# overcelebrate
Never ever forget a birthday, because overcelebrate will tell you.

overcelebrate checks for today's birthday based on a `config.yaml` file, in which you can put all birthday coordinates of your loved ones.
It will send out a pushover notification directly to your linked device.

overcelebrate works perfectly as a cronjob on kubernetes. 

`configmap.yaml`
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: overcelebrate-configmap
data:
  config.yaml: |
    bdays:
      - name: Dad
        date: 01-01-1955
    pushover:
      usertoken: "myUserToken"
      apitoken: "myApiToken"
```

`deployment.yaml`
```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: overcelebrate-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: overcelebrate
  template:
    metadata:
      labels:
        app: overcelebrate
    spec:
      containers:
        - name: overcelebrate
          image: overcelebrate
          imagePullPolicy: Never
          resources:
            limits:
              memory: 100m
              cpu: 100m
          volumeMounts:
            - mountPath: /app/config.yaml
              name: config
              subPath: config.yaml
      volumes:
        - name: config
          configMap:
            name: overcelebrate-configmap

```

