#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

GO_ENUM_VERSION=v0.3.11
HACK_DIR=$(dirname "${BASH_SOURCE[0]}")

cd ${HACK_DIR}/../bin && curl -fsSL "https://github.com/abice/go-enum/releases/download/${GO_ENUM_VERSION}/go-enum_$(uname -s)_$(uname -m)" -o go-enum

