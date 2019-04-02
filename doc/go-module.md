# Go Module

## Enable Go Module
    GO111MODULE=on

## Init module
    go mod init [module name]

## Dependency management
    go mod tidy
    go build -mod=vendor

## Edit go.mod
    go mod edit -require="{module_name}@{version_number}"