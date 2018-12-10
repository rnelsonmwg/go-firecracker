#!/usr/bin/env bash
set -eo pipefail

for f in /go/src/${GIT_PROVIDER}/${GIT_REPO}/scripts/docker/entrypoint.d/*; do
    case "$f" in
        *.sh)     echo "$0: running $f"; . "$f" ;;
        *)        echo "$0: ignoring $f"        ;;
    esac
done