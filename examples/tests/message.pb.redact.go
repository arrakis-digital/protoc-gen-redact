// Code generated by protoc-gen-redact. DO NOT EDIT.
// source: examples/tests/message.proto

package tests

import (
	context "context"
	redact "github.com/arrakis-digital/protoc-gen-redact/v3/redact/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ grpc.Server
	_ context.Context
	_ redact.Redactor
	_ codes.Code
	_ status.Status
	_ emptypb.Empty
	_ redact.FieldRules
)

// Redact method implementation for TestMessage
func (x *TestMessage) Redact() string {
	if x == nil {
		return ""
	}

	// Redacting field: FloatValue
	x.FloatValue = 3.2

	// Redacting field: DoubleValue
	x.DoubleValue = 6.4

	// Redacting field: Int32Value
	x.Int32Value = 32

	// Redacting field: Int64Value
	x.Int64Value = 64

	// Redacting field: Uint32Value
	x.Uint32Value = 32

	// Redacting field: Uint64Value
	x.Uint64Value = 64

	// Redacting field: Sint32Value
	x.Sint32Value = 32

	// Redacting field: Sint64Value
	x.Sint64Value = 64

	// Redacting field: Fixed32Value
	x.Fixed32Value = 32

	// Redacting field: Fixed64Value
	x.Fixed64Value = 64

	// Redacting field: Sfixed32Value
	x.Sfixed32Value = 32

	// Redacting field: Sfixed64Value
	x.Sfixed64Value = 64

	// Redacting field: BoolValue
	x.BoolValue = true

	// Redacting field: StringValue
	x.StringValue = `redacted-value-value`

	// Redacting field: BytesValue
	x.BytesValue = []byte(`redacted-value-value`)

	// Redacting field: EnumValue
	x.EnumValue = 2

	// Redacting field: MessageNil
	x.MessageNil = nil

	// Redacting field: MessageSkip
	// MessageSkip redaction is skipped

	// Redacting field: MessageEmpty
	x.MessageEmpty = &TestMessage{}

	// Redacting field: Map1Empty
	x.Map1Empty = map[int64]string{}

	// Redacting field: Map2Empty
	x.Map2Empty = map[string]*emptypb.Empty{}

	// Redacting field: Map1Nested
	for k := range x.Map1Nested {
		x.Map1Nested[k] = "REDACTED"
	}

	// Redacting field: Map2Nested
	for k := range x.Map2Nested {
		redact.Apply(x.Map2Nested[k])
	}

	// Redacting field: Map1Item
	for k := range x.Map1Item {
		x.Map1Item[k] = `3`
	}

	// Redacting field: Map2ItemNil
	for k := range x.Map2ItemNil {
		x.Map2ItemNil[k] = nil
	}

	// Redacting field: Map2ItemSkip
	// Map2ItemSkip redaction is skipped

	// Redacting field: Map2ItemEmpty
	for k := range x.Map2ItemEmpty {
		x.Map2ItemEmpty[k] = &emptypb.Empty{}
	}
	return x.String()
}

// Redact method implementation for RepeatedM
func (x *RepeatedM) Redact() string {
	if x == nil {
		return ""
	}

	// Redacting field: FloatValueEmpties
	x.FloatValueEmpties = []float32{}

	// Redacting field: FloatValueNested
	for k := range x.FloatValueNested {
		x.FloatValueNested[k] = 0
	}

	// Redacting field: FloatValues
	for k := range x.FloatValues {
		x.FloatValues[k] = 3.2
	}

	// Redacting field: DoubleValueEmpties
	x.DoubleValueEmpties = []float64{}

	// Redacting field: DoubleValueNested
	for k := range x.DoubleValueNested {
		x.DoubleValueNested[k] = 0
	}

	// Redacting field: DoubleValues
	for k := range x.DoubleValues {
		x.DoubleValues[k] = 6.4
	}

	// Redacting field: Int32ValueEmpties
	x.Int32ValueEmpties = []int32{}

	// Redacting field: Int32ValueNested
	for k := range x.Int32ValueNested {
		x.Int32ValueNested[k] = 0
	}

	// Redacting field: Int32Values
	for k := range x.Int32Values {
		x.Int32Values[k] = 32
	}

	// Redacting field: Int64ValueEmpties
	x.Int64ValueEmpties = []int64{}

	// Redacting field: Int64ValueNested
	for k := range x.Int64ValueNested {
		x.Int64ValueNested[k] = 0
	}

	// Redacting field: Int64Values
	for k := range x.Int64Values {
		x.Int64Values[k] = 64
	}

	// Redacting field: Uint32ValueEmpties
	x.Uint32ValueEmpties = []uint32{}

	// Redacting field: Uint32ValueNested
	for k := range x.Uint32ValueNested {
		x.Uint32ValueNested[k] = 0
	}

	// Redacting field: Uint32Values
	for k := range x.Uint32Values {
		x.Uint32Values[k] = 32
	}

	// Redacting field: Uint64ValueEmpties
	x.Uint64ValueEmpties = []uint64{}

	// Redacting field: Uint64ValueNested
	for k := range x.Uint64ValueNested {
		x.Uint64ValueNested[k] = 0
	}

	// Redacting field: Uint64Values
	for k := range x.Uint64Values {
		x.Uint64Values[k] = 64
	}

	// Redacting field: Sint32ValueEmpties
	x.Sint32ValueEmpties = []int32{}

	// Redacting field: Sint32ValueNested
	for k := range x.Sint32ValueNested {
		x.Sint32ValueNested[k] = 0
	}

	// Redacting field: Sint32Values
	for k := range x.Sint32Values {
		x.Sint32Values[k] = 32
	}

	// Redacting field: Sint64ValueEmpties
	x.Sint64ValueEmpties = []int64{}

	// Redacting field: Sint64ValueNested
	for k := range x.Sint64ValueNested {
		x.Sint64ValueNested[k] = 0
	}

	// Redacting field: Sint64Values
	for k := range x.Sint64Values {
		x.Sint64Values[k] = 64
	}

	// Redacting field: Fixed32ValueEmpties
	x.Fixed32ValueEmpties = []uint32{}

	// Redacting field: Fixed32ValueNested
	for k := range x.Fixed32ValueNested {
		x.Fixed32ValueNested[k] = 0
	}

	// Redacting field: Fixed32Values
	for k := range x.Fixed32Values {
		x.Fixed32Values[k] = 32
	}

	// Redacting field: Fixed64ValueEmpties
	x.Fixed64ValueEmpties = []uint64{}

	// Redacting field: Fixed64ValueNested
	for k := range x.Fixed64ValueNested {
		x.Fixed64ValueNested[k] = 0
	}

	// Redacting field: Fixed64Values
	for k := range x.Fixed64Values {
		x.Fixed64Values[k] = 64
	}

	// Redacting field: Sfixed32ValueEmpties
	x.Sfixed32ValueEmpties = []int32{}

	// Redacting field: Sfixed32ValueNested
	for k := range x.Sfixed32ValueNested {
		x.Sfixed32ValueNested[k] = 0
	}

	// Redacting field: Sfixed32Values
	for k := range x.Sfixed32Values {
		x.Sfixed32Values[k] = 32
	}

	// Redacting field: Sfixed64ValueEmpties
	x.Sfixed64ValueEmpties = []int64{}

	// Redacting field: Sfixed64ValueNested
	for k := range x.Sfixed64ValueNested {
		x.Sfixed64ValueNested[k] = 0
	}

	// Redacting field: Sfixed64Values
	for k := range x.Sfixed64Values {
		x.Sfixed64Values[k] = 64
	}

	// Redacting field: BoolValueEmpties
	x.BoolValueEmpties = []bool{}

	// Redacting field: BoolValueNested
	for k := range x.BoolValueNested {
		x.BoolValueNested[k] = false
	}

	// Redacting field: BoolValues
	for k := range x.BoolValues {
		x.BoolValues[k] = true
	}

	// Redacting field: StringValueEmpties
	x.StringValueEmpties = []string{}

	// Redacting field: StringValueNested
	for k := range x.StringValueNested {
		x.StringValueNested[k] = "REDACTED"
	}

	// Redacting field: StringValues
	for k := range x.StringValues {
		x.StringValues[k] = `redacted-value-value`
	}

	// Redacting field: BytesValueEmpties
	x.BytesValueEmpties = [][]byte{}

	// Redacting field: BytesValueNested
	for k := range x.BytesValueNested {
		x.BytesValueNested[k] = nil
	}

	// Redacting field: BytesValues
	for k := range x.BytesValues {
		x.BytesValues[k] = []byte(`redacted-value-value`)
	}

	// Redacting field: EnumValueEmpties
	x.EnumValueEmpties = []TestEnum{}

	// Redacting field: EnumValueNested
	for k := range x.EnumValueNested {
		x.EnumValueNested[k] = 0
	}

	// Redacting field: EnumValues
	for k := range x.EnumValues {
		x.EnumValues[k] = 2
	}

	// Redacting field: MessageNils
	for k := range x.MessageNils {
		x.MessageNils[k] = nil
	}

	// Redacting field: MessageSkips
	// MessageSkips redaction is skipped

	// Redacting field: MessageNested
	for k := range x.MessageNested {
		redact.Apply(x.MessageNested[k])
	}

	// Redacting field: MessageEmpties
	for k := range x.MessageEmpties {
		x.MessageEmpties[k] = &TestMessage{}
	}
	return x.String()
}
