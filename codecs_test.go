package codecs

import (
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/viamrobotics/mongo-go-driver-protobuf/pmongo"
	"github.com/viamrobotics/mongo-go-driver-protobuf/test"
)

func TestCodecs(t *testing.T) {
	rb := bson.NewRegistryBuilder()
	r := Register(rb).Build()

	tm := time.Now()
	// BSON accuracy is in milliseconds
	tm = time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(),
		(tm.Nanosecond()/1000000)*1000000, tm.Location())

	ts := timestamppb.New(tm)

	objectID := primitive.NewObjectID()
	id := pmongo.NewObjectId(objectID)

	t.Run("primitive object id", func(t *testing.T) {
		resultID, err := id.GetObjectID()
		if err != nil {
			t.Errorf("mongodb.ObjectId.GetPrimitiveObjectID() error = %v", err)
			return
		}

		if !reflect.DeepEqual(objectID, resultID) {
			t.Errorf("failed: primitive object ID=%#v, ID=%#v", objectID, id)
			return
		}
	})

	in := test.Data{
		BoolValue:   &wrapperspb.BoolValue{Value: true},
		BytesValue:  &wrapperspb.BytesValue{Value: make([]byte, 5)},
		DoubleValue: &wrapperspb.DoubleValue{Value: 1.2},
		FloatValue:  &wrapperspb.FloatValue{Value: 1.3},
		Int32Value:  &wrapperspb.Int32Value{Value: -12345},
		Int64Value:  &wrapperspb.Int64Value{Value: -123456789},
		StringValue: &wrapperspb.StringValue{Value: "qwerty"},
		Uint32Value: &wrapperspb.UInt32Value{Value: 12345},
		Uint64Value: &wrapperspb.UInt64Value{Value: 123456789},
		Timestamp:   ts,
		Id:          id,
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		b, err := bson.MarshalWithRegistry(r, &in)
		if err != nil {
			t.Errorf("bson.MarshalWithRegistry error = %v", err)
			return
		}

		var out test.Data

		if err = bson.UnmarshalWithRegistry(r, b, &out); err != nil {
			t.Errorf("bson.UnmarshalWithRegistry error = %v", err)
			return
		}

		if !reflect.DeepEqual(in, out) {
			t.Errorf("failed: in=%#v, out=%#v", in, out)
			return
		}
	})

	t.Run("marshal-jsonpb/unmarshal-jsonpb", func(t *testing.T) {
		b, err := protojson.Marshal(in.ProtoReflect().Interface())
		if err != nil {
			t.Errorf("jsonpb.Marshaler.Marshal error = %v", err)
			return
		}

		var out test.Data
		if err := protojson.Unmarshal(b, out.ProtoReflect().Interface()); err != nil {
			t.Errorf("jsonpb.Unmarshal error = %v", err)
			return
		}

		if !reflect.DeepEqual(in, out) {
			t.Errorf("failed: in=%#v, out=%#v", in, out)
			return
		}
	})
}
