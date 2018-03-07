# sample-augmentor
Sample Data Augmentor for Beeswax

## Objectives

The goal of this repository is:

- To provide a sample data augmentor that receives an HTTP request, calls out to a Redis database to look up a value, and return a response.

## Usage

This package uses the Go toolchain to build the binary, and `make` as a build tool. That is all that is required to build the binary.

```bash
$ git clone https://github.com/pswaminathan/sample-augmentor.git
$ cd sample-augmentor
$ make
$ ./augmentor -port 8080 -samplefile testdata/sample_data.json
```

This creates an in-memory Redis store, loads sample data, and creates an HTTP server. You can then use the [Augmentor Requests Generator](https://github.com/BeeswaxIO/beeswax-api/tree/master/beeswax/tools/augmentor) to send some requests to the server!

In order to build on it and/or work with the vendored packages, you will need [dep](https://github.com/golang/dep) to manage external Go dependencies, and the protobuf compiler and [gogoprotobuf](https://github.com/gogo/protobuf) to compile the protos.
