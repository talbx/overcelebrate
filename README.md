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

`cronjob.yaml`
```yaml
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: overcelebrate-cronjob
spec:
  schedule: "0 8 * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: overcelebrate
            image: overcelebrate
            resources:
              limits:
                memory: 50m
                cpu: 100m
            imagePullPolicy: Never
            volumeMounts:
              - mountPath: /app/config.yaml
                name: config
                subPath: config.yaml
          restartPolicy: OnFailure
          volumes:
            - name: config
              configMap:
                name: overcelebrate-configmap
```
