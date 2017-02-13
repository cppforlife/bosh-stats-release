# bosh-stats

Collects different BOSH environment statistics and forwards it to some destination.

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

## Todo

- average time for CPI actions
  - `cpi.call` (cpi, action, object_type, duration_sec)
- deployment times
  - `deployment.run` (duration_sec, success)
- deployment errors
- net configuration
- is CPI config used
- is runtime config used
