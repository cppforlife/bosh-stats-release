# bosh-stats

Collects different BOSH environment statistics and forwards it to some destination.

Available destinations:

- local file
  - `bosh -d bosh-stats logs --num=100 -f|grep stats.reporter.Logger`
- DataDog

## Metrics

- `releases.count`
- `release` (name, version)
- `stemcells.count`
- `stemcell` (name, version)
- `deployments.count`
- `deployment.instances.count`
- `instances.count`
- `azs.count`
- `disk_types.count`
- `vm_types.count`
- `networks.count`
- `networks.manual.count`
- `networks.dynamic.count`
- `networks.vip.count`
- `compilation.workers`
- `addons.count`
- `director.version`
- `director.uuid`
- `director.auth.type`
- `director.cpi`

## Done
- basic datadog integration

## Todo

- datadog integration v0.2
  - what to do for values that cannot be converted to integer/float?
- average time for CPI actions
  - `cpi.call` (cpi, action, object_type, duration_sec)
- deployment times
  - `deployment.run` (duration_sec, success)
- deployment errors
- net configuration
- is CPI config used
- is runtime config used
- opt-in to metrics
- send all stats to stats.bosh.io (configurable)
  - for now forward to some http endpoint

## Test

```
./src/github.com/cppforlife/bosh-stats/bin/test
```
