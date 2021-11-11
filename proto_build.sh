#bash

protoc --proto_path=proto/pmongo --go_out=../../../ objectid.proto

protoc --proto_path=test --proto_path=proto --proto_path=proto/third_party --go_out=paths=source_relative:test codecs_test.proto
