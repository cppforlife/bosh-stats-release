#!/bin/bash

set -e

bosh -d bosh-stats -n deploy example.yml \
  -v director_ip="<ip address>" \
  --var-file director_ssl_ca=<(bosh interpolate <path/to/director/creds.yml> --path /director_ssl/ca) \
  -v director_client=admin \
  -v director_client_secret=`bosh interpolate <path/to/director/creds.yml> --path /admin_password`
