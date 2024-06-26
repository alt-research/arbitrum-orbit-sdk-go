#!/bin/bash -e
if [[ $DEBUG == "true" ]]; then
  set -x
fi

# https://go.dev/dl
GO_VERSION="1.22"
# https://github.com/golangci/golangci-lint/releases
GOLANGCI_LINT_VERSION="1.59.1"
# https://github.com/ethereum/go-ethereum/releases
GETH_VERSION=1.14.5
ABIGEN_VERSION=$GETH_VERSION

# directory variables
TMPDIR=${TMPDIR:-/tmp}
READLINK=$([[ $OSTYPE == 'darwin'* ]] && echo 'greadlink' || echo 'readlink')
script_path=$($READLINK -f "${BASH_SOURCE:-$0}")
SCRIPTS_ROOT=$(dirname $script_path)
REPO_ROOT=$(dirname $SCRIPTS_ROOT)
BIN=$REPO_ROOT/bin
export GOBIN=$BIN
export PATH=$BIN:$PATH

# executable variables
GOLANGCI_LINT=$BIN/golangci-lint
ABIGEN=$BIN/abigen

info() {
  echo -e "\033[1;32m$@\033[0m"
}

error() {
  echo -e "\033[1;31m$@\033[0m"
}

fatal() {
  error "$@"
  exit 1
}

must_executable() {
  executable=$1
  message=$2
  if ! command -v $1 &>/dev/null; then
    fatal "$message"
  fi
}

# ensure basic tools
ensure_tools() {
  # ensure go
  must_executable go "[deps:go] go not found"
  go_version=$(go version | awk '{print $3}')
  # if VER is not prefixed with GO_VERSION
  if [[ "$go_version" != "go${GO_VERSION}."* ]]; then
    fatal "[deps:go] go version $go_version is not ${GO_VERSION}.*"
  fi
  # ensure curl exist
  must_executable curl "[deps:curl] curl not found"
  # unzip
  must_executable unzip "[deps:unzip] unzip not found"
  # git
  must_executable git "[deps:git] git not found"
}
ensure_tools

# machine variables
# https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
GOHOSTOS=$(go env GOHOSTOS)     # darwin, linux, windows
GOHOSTARCH=$(go env GOHOSTARCH) # amd64, arm64, 386

ensure_bin_dir() {
  mkdir -p $BIN
}

er() {
  # echo and run
  info "[RUN] $@"
  $@
}

# golangci-lint
golangci_lint_version() {
  $GOLANGCI_LINT --version | awk '{print $4}'
}

install_golangci_lint() {
  ensure_bin_dir
  # https://golangci-lint.run/welcome/install/#local-installation
  info "[RUN] curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b $BIN v${GOLANGCI_LINT_VERSION}"
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b $BIN v${GOLANGCI_LINT_VERSION}
  info "[deps] golangci-lint $(golangci_lint_version) installed"
}

ensure_golangci_lint() {
  if [[ ! -x $GOLANGCI_LINT ]]; then
    info "golangci-lint not found, installing..."
    install_golangci_lint
  elif [[ $(golangci_lint_version) != "${GOLANGCI_LINT_VERSION}" ]]; then
    info "golangci-lint version mismatch, found $(golangci_lint_version), installing $GOLANGCI_LINT_VERSION..."
    install_golangci_lint
  else
    info "[deps] golangci-lint $(golangci_lint_version) found"
  fi
  touch $GOLANGCI_LINT
}

# abigen
# abigen version 1.14.5-stable
abigen_version() {
  $ABIGEN --version | awk '{print $3}' | cut -d'-' -f1
}

install_abigen() {
  ensure_bin_dir
  # https://geth.ethereum.org/docs/interface/go-ethereum-abigen/
  er go install github.com/ethereum/go-ethereum/cmd/abigen@v${ABIGEN_VERSION}
  info "[deps] abigen $(abigen_version) installed"
}

ensure_abigen() {
  if [[ ! -x $ABIGEN ]]; then
    info "abigen not found, installing..."
    install_abigen
  elif [[ $(abigen_version) != "${ABIGEN_VERSION}" ]]; then
    info "abigen version mismatch, found $(abigen_version), installing $ABIGEN_VERSION..."
    install_abigen
  else
    info "[deps] abigen $(abigen_version) found"
  fi
}

# main
ensure_golangci_lint
ensure_abigen

info "[deps] all dependencies installed and updated"
