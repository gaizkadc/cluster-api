# cluster-api

The Cluster API component provides a method for the Nalej components in the Application cluster to communicate
and trigger operations in the Nalej Management cluster.

## Getting Started

### Prerequisites

To run this component you should have deployed at least the following components:

* [network-manager](https://github.com/nalej/network-manager)
* [conductor](https://github.com/nalej/conductor)
* [device-manager](https://github.com/nalej/device-manager)
* [authx](https://github.com/nalej/authx)
* [nalej-bus](https://github.com/nalej/nalej-bus)
* [cluster-watcher](https://github.com/nalej/cluster-watcher)

### Build and compile

To build and compile this repository use the provided Makefile:

```shell script
make all
```

This operation generates the binaries for this repo, download dependencies,
run existing tests and generate ready-to-deploy Kubernetes files.

### Run tests

Tests are executed using Ginkgo. To run all the available tests:

```shell script
make test
```

### Update dependencies

Dependencies are managed using Godep. For an automatic dependencies download use:

```shell script
make dep
```

To have all dependencies up-to-date run:

```shell script
dep ensure -update -v
```

## Known Issues

## Contributing

Please read [contributing.md](contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

See also the list of [contributors](https://github.com/nalej/grpc-utils/contributors) who participated in this project.

## License
This project is licensed under the Apache 2.0 License - see the [LICENSE-2.0.txt](LICENSE-2.0.txt) file for details.
