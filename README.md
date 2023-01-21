# Megacron
A simple tool that lets you conveniently define multiple cronjobs in helm

Comes with Prometheus metrics that track for:
- successful cron runs
- failed cron runs
- running time of each cron

Has readiness/liveliness probes.

Waits for crons to finish when shutting down.

## Installation
1. Add the chart to your Chart.yaml:
```yaml
  - name: megacron
    version: 1.0.0
    repository: "https://taggart-comet.github.io/megacron/"
```
2. Add the chart to your values.yaml:
```yaml
megacron:
  envsFrom:
    - secretRef:
        name: secret-name-you-want-to-use-for-all-crons
  imagePullSecrets:
    - name: image-pull-secret
  
  # define cronjobs that will be running in a never-stopping pod
  alwaysRunning:
    enabled: true
    
    # image for the crons to run in (it will be reused in cronJobs, if not redefined there)
    image: "busybox:latest"
    crons:
      - "* * * * * bin/symfony system:example-cron"
      - "* * * * * bin/symfony system:example-cron2"
      - "*/5 * * * * /any-command"
    # you can define values that need to be passed to the cronjob
#    volumes:

  # define cronjobs that will run with a custom schedule
  # using kubernetes CronJobs
  cronJobs:
    - name: "fifteen-minutes-crons"
      schedule: "*/15 * * * *"
      crons:
        - "bin/symfony system:example-cron"
        - "bin/symfony system:example-cron2"
    - name: "nightly-crons"
      schedule: "0 0 * * *"
      crons:
        - "bin/symfony system:example-cron"

  # if you need to collect metrics from your crons using Prometheus pushgateway
  prometheusPush:
    enabled: true
    gatewayUrl: http://prometheus-pushgateway:9091
    job: job-name
```
3. You're good to go!

## How does it work?
For each container there's an initContainer that puts the executable (`runner`) file into the main container.

The `runner` runs all the crons inside, tracking their statuses and logging the output of the crons into container's stdout.

By default initContainer uses an image from `docker.io/taggartcomet/megacron:x.x.x`, it's hardcoded into the HelmChart. But can be redefined using `initImage` field.