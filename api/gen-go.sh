#!/usr/bin/env bash
thrift -out .. -r --gen go:thrift_import=git.apache.org/thrift.git/lib/go/thrift api.thrift
