#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# generate-internal-groups generates everything for a project with internal types, e.g. an
# user-provided API server based on k8s.io/apiserver.

if [ "$#" -lt 4 ] || [ "${1}" == "--help" ]; then
  cat <<EOF
Usage: $(basename "$0") <generators> <output-package> <internal-apis-package> <extensiona-apis-package> <groups-versions> ...

  <generators>        the generators comma separated to run (deepcopy,defaulter,conversion,client,lister,informer,openapi) or "all".
  <output-package>    the output package name (e.g. github.com/example/project/pkg/generated).
  <int-apis-package>  the internal types dir (e.g. github.com/example/project/pkg/apis).
  <ext-apis-package>  the external types dir (e.g. github.com/example/project/pkg/apis or githubcom/example/apis).
  <groups-versions>   the groups and their versions in the format "groupA:v1,v2 groupB:v1 groupC:v2", relative
                      to <api-package>.
  ...                 arbitrary flags passed to all generator binaries.

Examples:
  $(basename "$0") all                           github.com/example/project/pkg/client github.com/example/project/pkg/apis github.com/example/project/pkg/apis "foo:v1 bar:v1alpha1,v1beta1"
  $(basename "$0") deepcopy,defaulter,conversion github.com/example/project/pkg/client github.com/example/project/pkg/apis github.com/example/project/apis     "foo:v1 bar:v1alpha1,v1beta1"
EOF
  exit 0
fi

GENS="$1"
OUTPUT_PKG="$2"
EXT_APIS_PKG="$3"
GROUPS_WITH_VERSIONS="$4"
shift 4

echo "GENS is $GENS"
echo "OUTPUT_PKG is $OUTPUT_PKG"
echo "EXT_APIS_PKG is $EXT_APIS_PKG"
echo "GROUPS_WITH_VERSIONS is $GROUPS_WITH_VERSIONS"

# go install ./"$(dirname "${0}")"/cmd/{defaulter-gen,conversion-gen,client-gen,lister-gen,informer-gen,deepcopy-gen,openapi-gen}

function codegen::join() { local IFS="$1"; shift; echo "$*"; }

# enumerate group versions
ALL_FQ_APIS=() # e.g. k8s.io/kubernetes/pkg/apis/apps k8s.io/api/apps/v1
EXT_FQ_APIS=() # e.g. k8s.io/api/apps/v1
for GVs in ${GROUPS_WITH_VERSIONS}; do
  IFS=: read -r G Vs <<<"${GVs}"

  # enumerate versions
  for V in ${Vs//,/ }; do
    ALL_FQ_APIS+=("${EXT_APIS_PKG}/${G}/${V}")
    EXT_FQ_APIS+=("${EXT_APIS_PKG}/${G}/${V}")
  done
done

if [ "${GENS}" = "all" ] || grep -qw "deepcopy" <<<"${GENS}"; then
  echo "Generating deepcopy funcs for $(codegen::join , ${ALL_FQ_APIS[@]})"
  #echo $(codegen::join , "${ALL_FQ_APIS[@]}")
  "${GOPATH}/bin/deepcopy-gen" --input-dirs "$(codegen::join , "${ALL_FQ_APIS[@]}")" -O zz_generated.deepcopy --bounding-dirs "${EXT_APIS_PKG}" "$@"
fi

if [ "${GENS}" = "all" ] || grep -qw "defaulter" <<<"${GENS}"; then
  echo "Generating defaulters for $(codegen::join , ${EXT_FQ_APIS[@]})"
  "${GOPATH}/bin/defaulter-gen"  --input-dirs "$(codegen::join , "${EXT_FQ_APIS[@]}")" -O zz_generated.defaults "$@"
fi

# TODO figure out why this doesn't work
#if [ "${GENS}" = "all" ] || grep -qw "conversion" <<<"${GENS}"; then
#  echo "Generating conversions for $(codegen::join , ${ALL_FQ_APIS[@]})"
#  "${GOPATH}/bin/conversion-gen" --input-dirs "$(codegen::join , "${ALL_FQ_APIS[@]}")" -O zz_generated.conversion "$@"
#fi

if [ "${GENS}" = "all" ] || grep -qw "client" <<<"${GENS}"; then
  echo "Generating clientset for ${GROUPS_WITH_VERSIONS} at ${OUTPUT_PKG}/${CLIENTSET_PKG_NAME:-clientset}"
  "${GOPATH}/bin/client-gen" --clientset-name "${CLIENTSET_NAME_VERSIONED:-versioned}" --input-base "" --input "$(codegen::join , "${EXT_FQ_APIS[@]}")" --output-package "${OUTPUT_PKG}/${CLIENTSET_PKG_NAME:-clientset}" "$@"
fi

if [ "${GENS}" = "all" ] || grep -qw "lister" <<<"${GENS}"; then
  echo "Generating listers for ${GROUPS_WITH_VERSIONS} at ${OUTPUT_PKG}/listers"
  "${GOPATH}/bin/lister-gen" --input-dirs "$(codegen::join , "${ALL_FQ_APIS[@]}")" --output-package "${OUTPUT_PKG}/listers" "$@"
fi

if [ "${GENS}" = "all" ] || grep -qw "informer" <<<"${GENS}"; then
  echo "Generating informers for ${GROUPS_WITH_VERSIONS} at ${OUTPUT_PKG}/informers"
  "${GOPATH}/bin/informer-gen" \
           --input-dirs "$(codegen::join , "${ALL_FQ_APIS[@]}")" \
           --versioned-clientset-package "${OUTPUT_PKG}/${CLIENTSET_PKG_NAME:-clientset}/${CLIENTSET_NAME_VERSIONED:-versioned}" \
           --listers-package "${OUTPUT_PKG}/listers" \
           --output-package "${OUTPUT_PKG}/informers" \
           "$@"
fi

if [ "${GENS}" = "all" ] || grep -qw "openapi" <<<"${GENS}"; then
  echo "Generating OpenAPI definitions for ${GROUPS_WITH_VERSIONS} at ${OUTPUT_PKG}/openapi"
  declare -a OPENAPI_EXTRA_PACKAGES
  "${GOPATH}/bin/openapi-gen" \
           --input-dirs "$(codegen::join , "${EXT_FQ_APIS[@]}" "${OPENAPI_EXTRA_PACKAGES[@]+"${OPENAPI_EXTRA_PACKAGES[@]}"}")" \
           --input-dirs "k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version,go.aporeto.io/gaia" \
           --output-package "${OUTPUT_PKG}/openapi" \
           -O zz_generated.openapi \
           "$@"
fi
