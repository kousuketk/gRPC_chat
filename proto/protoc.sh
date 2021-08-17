#!/bin/bash

protoc messenger.proto --go_out=.
protoc messenger.proto --go-grpc_out=.