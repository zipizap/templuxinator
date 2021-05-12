#!/usr/bin/env bash
# Paulo Aleixo Campos
__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
function shw_info { echo -e '\033[1;34m'"$1"'\033[0m'; }
function error { echo "ERROR in ${1}"; exit 99; }
trap 'error $LINENO' ERR
PS4='████████████████████████${BASH_SOURCE}@${FUNCNAME[0]:-}[${LINENO}]>  '
set -o errexit
set -o pipefail
set -o nounset
set -o xtrace




cd "${__dir}"

# #https://goreleaser.com/quick-start/
# goreleaser init

# Check GITHUB_TOKEN env-var
(set +x; GITHUB_TOKEN="${GITHUB_TOKEN?missing env-var}")

## Create tag
# NOTE: undo tag with:    git tag -d v1.0.0 && git push --delete origin v1.0.0
git tag -a v1.0.0 -m "First release"
git push origin v1.0.0

goreleaser release

