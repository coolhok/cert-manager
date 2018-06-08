#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

## This script will generate a reference documentation site into ./docs/generated/reference/output
## It requires a number of tools be installed:
##
## * openapi-gen
## * gen-apidocs
## * docker
##

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..

REFERENCE_PATH="docs/generated/reference"
REFERENCE_ROOT=$(cd "${SCRIPT_ROOT}/${REFERENCE_PATH}" 2> /dev/null && pwd -P)

## cleanup removes files that are leftover from running various tools and not required
## for the actual output
cleanup() {
    pushd "${REFERENCE_ROOT}"
    echo "+++ Cleaning up temporary docsgen files"
    # Clean up old temporary files
    find "./output" \
        -type f \
        -not -name bootstrap.min.css \
        -not -name font-awesome.min.css \
        -not -name highlight.js \
        -not -name stylesheet.css \
        -not -name index.html \
        -not -name scroll.js \
        -not -name tabvisibility.js \
        -not -name default.css \
        -not -name navData.js \
        -not -name jquery.min.js \
        -not -name jquery.scrollTo.min.js \
        -not -name fontawesome-webfont.ttf \
        -not -name fontawesome-webfont.woff \
        -not -name fontawesome-webfont.woff2 \
        -exec rm -Rf {} \; || true
    find "./output" \
        -type d \
        -empty \
        -print0 | xargs -0 rmdir || true
    rm -Rf "openapi-spec" "openapi" "includes" "manifest.json"

    popd
}

trap cleanup EXIT

cleanup
echo "+++ Removing old output"
rm -Rf "${REFERENCE_ROOT}/output"

echo "+++ Creating temporary directories"
# Create all required directories
mkdir -p "${REFERENCE_ROOT}/openapi-spec"

echo "+++ Generating openapi_generated.go into 'github.com/jetstack/cert-manager/${REFERENCE_PATH}/openapi'"
# Generate Golang types for OpenAPI spec
openapi-gen \
        --input-dirs github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version \
        --output-package "github.com/jetstack/cert-manager/${REFERENCE_PATH}/openapi"

echo "+++ Running './${REFERENCE_PATH}/main.go'"
# Generate swagger.json
go run "./${REFERENCE_PATH}/main.go"

echo "+++ Running gen-apidocs"
# Generate Markdown docs
gen-apidocs \
    --copyright "<a href=\"https://jetstack.io\">Copyright 2018 Jetstack Ltd.</a>" \
    --title "Cert-manager API Reference" \
    --config-dir ./docs/generated/reference/

echo "+++ Running brodocs"
# Run brodocs
docker run \
    -v "${REFERENCE_ROOT}/includes:/source" \
    -v "${REFERENCE_ROOT}/output:/build" \
    -v "${REFERENCE_ROOT}:/manifest" \
    "gcr.io/kubebuilder/brodocs@sha256:81a93f8d3dde22288a2ac7ff287d8157ce60c3b4b29a6d66bdef58b4954d55c3"
