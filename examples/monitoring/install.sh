# https://github.com/grafana/helm-charts/tree/main/charts/loki-stack
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

# Default config
# helm upgrade --install loki grafana/loki-stack

# Deploy with override
# helm upgrade --install loki grafana/loki-stack --set fluent-bit.enabled=true,promtail.enabled=false

# Deploy with the file override
helm upgrade --install -f grafana-values.yaml loki grafana/loki-stack
