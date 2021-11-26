# Go SQL Driver for Hazelcast Sample

This project demonstrates how to use the Hazelcast [database/sql](https://pkg.go.dev/database/sql) driver.

## Build

1. Run `go mod tidy` to download the latest Hazelcast Go client with SQL support
2. Run `make sqlrun` which builds the `sqlrun` binary. 

## Usage

Make sure you use [Hazelcast Platform 5.0](https://hazelcast.com/get-started/) and Jet is enabled in your Hazelcast configuration (it is by default):

	<jet enabled="true"/>

### sqlrun

`sqlrun` can execute SQL commands from a file. You can optionally specify connection arguments:

    $ sqlrun -f myfile.sql
    $ sqlrun -c "10.20.30.40:5701;ClusterName=prod" -f myfile.sql
