package codecs

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/viamrobotics/mongo-go-driver-protobuf/pmongo"
)

var (
	// Protobuf wrappers types
	boolValueType   = reflect.TypeOf(wrapperspb.BoolValue{})
	bytesValueType  = reflect.TypeOf(wrapperspb.BytesValue{})
	doubleValueType = reflect.TypeOf(wrapperspb.DoubleValue{})
	floatValueType  = reflect.TypeOf(wrapperspb.FloatValue{})
	int32ValueType  = reflect.TypeOf(wrapperspb.Int32Value{})
	int64ValueType  = reflect.TypeOf(wrapperspb.Int64Value{})
	stringValueType = reflect.TypeOf(wrapperspb.StringValue{})
	uint32ValueType = reflect.TypeOf(wrapperspb.UInt32Value{})
	uint64ValueType = reflect.TypeOf(wrapperspb.UInt64Value{})

	// Protobuf Struct type
	structType = reflect.TypeOf(structpb.Struct{})

	// Map type
	mapType = reflect.TypeOf(map[string]interface{}{})

	// Protobuf Timestamp type
	timestampType = reflect.TypeOf(timestamppb.Timestamp{})

	// Time type
	timeType = reflect.TypeOf(time.Time{})

	// Slice type
	sliceType = reflect.TypeOf([]interface{}{})

	// ObjectId type
	objectIDType          = reflect.TypeOf(pmongo.ObjectId{})
	objectIDPrimitiveType = reflect.TypeOf(primitive.ObjectID{})

	// Codecs
	wrapperValueCodecRef = &wrapperValueCodec{}
	structCodecRef       = &structCodec{}
	timestampCodecRef    = &timestampCodec{}
	objectIDCodecRef     = &objectIDCodec{}
)

// wrapperValueCodec is codec for Protobuf type wrappers
type wrapperValueCodec struct {
}

// EncodeValue encodes Protobuf type wrapper value to BSON value
func (e *wrapperValueCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupEncoder(val.Type())
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, val)
}

// DecodeValue decodes BSON value to Protobuf type wrapper value
func (e *wrapperValueCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupDecoder(val.Type())
	if err != nil {
		return err
	}
	return enc.DecodeValue(ectx, vr, val)
}

// structCodec is codec for Protobuf Struct
type structCodec struct {
}

// EncodeValue encodes Protobuf Struct value to BSON value
func (e *structCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(structpb.Struct)
	enc, err := ectx.LookupEncoder(mapType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(v.AsMap()))
}

// DecodeValue decodes BSON value to Timestamp value
func (e *structCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(mapType)
	if err != nil {
		return err
	}
	var m map[string]interface{}
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&m).Elem()); err != nil {
		return err
	}
	st, err := structpb.NewStruct(m)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(*st))
	return nil
}

// timestampCodec is codec for Protobuf Timestamp
type timestampCodec struct {
}

// EncodeValue encodes Protobuf Timestamp value to BSON value
func (e *timestampCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(timestamppb.Timestamp)
	enc, err := ectx.LookupEncoder(timeType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(v.AsTime()))
}

// DecodeValue decodes BSON value to Timestamp value
func (e *timestampCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(timeType)
	if err != nil {
		return err
	}
	var t time.Time
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&t).Elem()); err != nil {
		return err
	}
	ts := timestamppb.New(t)
	val.Set(reflect.ValueOf(*ts))
	return nil
}

// objectIDCodec is codec for Protobuf ObjectId
type objectIDCodec struct {
}

// EncodeValue encodes Protobuf ObjectId value to BSON value
func (e *objectIDCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(pmongo.ObjectId)
	// Create primitive.ObjectId from string
	id, err := primitive.ObjectIDFromHex(v.Value)
	if err != nil {
		return err
	}
	enc, err := ectx.LookupEncoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(id))
}

// DecodeValue decodes BSON value to ObjectId value
func (e *objectIDCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	var id primitive.ObjectID
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&id).Elem()); err != nil {
		return err
	}
	oid := *pmongo.NewObjectId(id)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(oid))
	return nil
}

// Register registers Google protocol buffers types codecs
func Register(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	rb.RegisterTypeMapEntry(bsontype.Array, sliceType)
	return rb.RegisterCodec(boolValueType, wrapperValueCodecRef).
		RegisterCodec(bytesValueType, wrapperValueCodecRef).
		RegisterCodec(doubleValueType, wrapperValueCodecRef).
		RegisterCodec(floatValueType, wrapperValueCodecRef).
		RegisterCodec(int32ValueType, wrapperValueCodecRef).
		RegisterCodec(int64ValueType, wrapperValueCodecRef).
		RegisterCodec(stringValueType, wrapperValueCodecRef).
		RegisterCodec(uint32ValueType, wrapperValueCodecRef).
		RegisterCodec(uint64ValueType, wrapperValueCodecRef).
		RegisterCodec(structType, structCodecRef).
		RegisterCodec(timestampType, timestampCodecRef).
		RegisterCodec(objectIDType, objectIDCodecRef)
}
