protoc --proto_path=. --proto_path=../../proto --proto_path=../../proto/third_party --proto_path=../third_party --go_out=paths=source_relative:. -I . data.proto

protoc --proto_path=. --proto_path=../../proto --proto_path=../../proto/third_party --proto_path=../third_party --gotag_out=xxx="bson+\"-\"",output_path=.:. -I . data.proto
