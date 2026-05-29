#!/usr/bin/env python3
import json, sys
d = json.load(open(sys.argv[1]))
print(f"Component: {d['component']}")
print(f"Repo: {d.get('repo','')}")
for k in ['crds','webhooks','services','deployments','external_connections','prometheus_metrics','controller_watches','reconcile_sequences']:
    v = d.get(k, []) or []
    if len(v) > 0:
        print(f"  {k}: {len(v)}")
