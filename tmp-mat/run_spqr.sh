make clean build && \
    ~/go/bin/dlv --listen=:8888 --headless=true --api-version=2 --accept-multiclient \
    exec ./spqr-router -- run --config ./tmp-mat/dev-router.yaml --coordinator-config ./tmp-mat/dev-coordinator.yaml
