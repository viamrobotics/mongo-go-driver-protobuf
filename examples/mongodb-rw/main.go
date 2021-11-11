package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	codecs "github.com/viamrobotics/mongo-go-driver-protobuf"
)

func main() {
	log.Printf("connecting to MongoDB...")

	// Register custom codecs for protobuf Timestamp and wrapper types
	reg := codecs.Register(bson.NewRegistryBuilder()).Build()

	// Create MongoDB client with registered custom codecs for protobuf Timestamp and wrapper types
	// NOTE: "mongodb+srv" protocol means connect to Altas cloud MongoDB server
	//       use just "mongodb" if you connect to on-premise MongoDB server
	client, err := mongo.NewClient(options.Client().
		ApplyURI("mongodb://localhost:27017/experiments").
		SetRegistry(reg),
	)

	if err != nil {
		log.Fatalf("failed to create new MongoDB client: %#v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect client
	if err = client.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to MongoDB: %#v", err)
	}

	log.Printf("connected successfully")

	// Get collection from database
	coll := client.Database("experiments").Collection("proto")

	// Create a protobuf Struct value from golang map[string]interface{}
	ms := `{
		"name": "Test",
		"number": 12345,
		"nestedMap": {
			"foo": "bar"
		},
		"nestedArray": [
			"a", "b", "c", "d"
		]
	}`
	var mm map[string]interface{}
	json.Unmarshal([]byte(ms), &mm)

	sp, err := structpb.NewStruct(mm)
	if err != nil {
		log.Fatalf("failed to create structpb.NewStruct error: %v", err)
		return
	}

	// Create protobuf Timestamp value from golang Time
	ts := timestamppb.Now()

	// Fill in data structure
	in := Data{
		BoolValue:      true,
		BoolProtoValue: &wrapperspb.BoolValue{Value: true},

		BytesValue:      []byte{1, 2, 3, 4, 5},
		BytesProtoValue: &wrapperspb.BytesValue{Value: []byte{1, 2, 3, 4, 5}},

		DoubleValue:      123.45678,
		DoubleProtoValue: &wrapperspb.DoubleValue{Value: 123.45678},

		FloatValue:      123.45,
		FloatProtoValue: &wrapperspb.FloatValue{Value: 123.45},

		Int32Value:      -12345,
		Int32ProtoValue: &wrapperspb.Int32Value{Value: -12345},

		Int64Value:      -123456789000,
		Int64ProtoValue: &wrapperspb.Int64Value{Value: -123456789000},

		StringValue:      "qwerty",
		StringProtoValue: &wrapperspb.StringValue{Value: "qwerty"},

		Uint32Value:      12345,
		Uint32ProtoValue: &wrapperspb.UInt32Value{Value: 12345},

		Uint64Value:      123456789000,
		Uint64ProtoValue: &wrapperspb.UInt64Value{Value: 123456789000},

		Struct: sp,

		Timestamp: ts,
	}

	log.Printf("insert data into collection <experiments.proto>...")

	// Insert data into the collection
	res, err := coll.InsertOne(ctx, &in)
	if err != nil {
		log.Fatalf("insert data into collection <experiments.proto>: %#v", err)
	}
	id := res.InsertedID
	log.Printf("inserted new item with id=%v successfully", id)

	// Create filter and output structure to read data from collection
	var out Data
	filter := bson.D{{Key: "_id", Value: id}}

	// Read data from collection
	err = coll.FindOne(ctx, filter).Decode(&out)
	if err != nil {
		log.Fatalf("failed to read data (id=%v) from collection <experiments.proto>: %#v", id, err)
	}

	m := protojson.MarshalOptions{Indent: "  "}
	b, err := m.Marshal(out.ProtoReflect().Interface())
	if err != nil {
		log.Fatalf("jsonpb.Marshaler.Marshal error = %v", err)
	}

	log.Printf("read successfully:\n%s", string(b))
}
