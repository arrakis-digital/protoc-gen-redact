// Code generated by protoc-gen-redact. DO NOT EDIT.
// source: examples/tests/message.proto

package tests

import (
	context "context"
	redact "github.com/arrakis-digital/protoc-gen-redact/redact"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ grpc.Server
	_ context.Context
	_ redact.Redactor
	_ codes.Code
	_ status.Status
	_ redact.FieldRules
	_ empty.Empty
)

// Redact method implementation for TestMessage
func (x *TestMessage) Redact() {
	if x == nil {
		return
	}

	// Redacting field: Float
	x.Float = 0

	// Redacting field: FloatCustom
	x.FloatCustom = 3.2

	// Redacting field: Double
	x.Double = 0

	// Redacting field: DoubleCustom
	x.DoubleCustom = 6.4

	// Redacting field: Int32
	x.Int32 = 0

	// Redacting field: Int32Custom
	x.Int32Custom = 32

	// Redacting field: Int64
	x.Int64 = 0

	// Redacting field: Int64Custom
	x.Int64Custom = 64

	// Redacting field: Uint32
	x.Uint32 = 0

	// Redacting field: Uint32Custom
	x.Uint32Custom = 32

	// Redacting field: Uint64
	x.Uint64 = 0

	// Redacting field: Uint64Custom
	x.Uint64Custom = 64

	// Redacting field: Sint32
	x.Sint32 = 0

	// Redacting field: Sint32Custom
	x.Sint32Custom = 32

	// Redacting field: Sint64
	x.Sint64 = 0

	// Redacting field: Sint64Custom
	x.Sint64Custom = 64

	// Redacting field: Fixed32
	x.Fixed32 = 0

	// Redacting field: Fixed32Custom
	x.Fixed32Custom = 32

	// Redacting field: Fixed64
	x.Fixed64 = 0

	// Redacting field: Fixed64Custom
	x.Fixed64Custom = 64

	// Redacting field: Sfixed32
	x.Sfixed32 = 0

	// Redacting field: Sfixed32Custom
	x.Sfixed32Custom = 32

	// Redacting field: Sfixed64
	x.Sfixed64 = 0

	// Redacting field: Sfixed64Custom
	x.Sfixed64Custom = 64

	// Redacting field: Bool
	x.Bool = false

	// Redacting field: BoolCustom
	x.BoolCustom = true

	// Redacting field: String_
	x.String_ = "REDACTED"

	// Redacting field: StringCustom
	x.StringCustom = `redacted-custom-value`

	// Redacting field: Bytes
	x.Bytes = nil

	// Redacting field: BytesCustom
	x.BytesCustom = []byte(`redacted-custom-value`)

	// Redacting field: Enum
	x.Enum = 0

	// Redacting field: EnumCustom
	x.EnumCustom = 2

	// Redacting field: Message
	redact.Apply(x.Message)

	// Redacting field: MessageNil
	x.MessageNil = nil

	// Redacting field: MessageSkip
	// MessageSkip redaction is skipped

	// Redacting field: MessageEmpty
	x.MessageEmpty = &TestMessage{}

	// Redacting field: Map1
	x.Map1 = nil

	// Redacting field: Map2
	x.Map2 = nil

	// Redacting field: Map1Empty
	x.Map1Empty = map[int64]string{}

	// Redacting field: Map2Empty
	x.Map2Empty = map[string]*empty.Empty{}

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
		x.Map2ItemEmpty[k] = &empty.Empty{}
	}
}

// Redact method implementation for RepeatedM
func (x *RepeatedM) Redact() {
	if x == nil {
		return
	}

	// Redacting field: Floats
	x.Floats = nil

	// Redacting field: FloatCustomEmpties
	x.FloatCustomEmpties = []float32{}

	// Redacting field: FloatCustomNested
	for k := range x.FloatCustomNested {
		x.FloatCustomNested[k] = 0
	}

	// Redacting field: FloatCustoms
	for k := range x.FloatCustoms {
		x.FloatCustoms[k] = 3.2
	}

	// Redacting field: Doubles
	x.Doubles = nil

	// Redacting field: DoubleCustomEmpties
	x.DoubleCustomEmpties = []float64{}

	// Redacting field: DoubleCustomNested
	for k := range x.DoubleCustomNested {
		x.DoubleCustomNested[k] = 0
	}

	// Redacting field: DoubleCustoms
	for k := range x.DoubleCustoms {
		x.DoubleCustoms[k] = 6.4
	}

	// Redacting field: Int32S
	x.Int32S = nil

	// Redacting field: Int32CustomEmpties
	x.Int32CustomEmpties = []int32{}

	// Redacting field: Int32CustomNested
	for k := range x.Int32CustomNested {
		x.Int32CustomNested[k] = 0
	}

	// Redacting field: Int32Customs
	for k := range x.Int32Customs {
		x.Int32Customs[k] = 32
	}

	// Redacting field: Int64S
	x.Int64S = nil

	// Redacting field: Int64CustomEmpties
	x.Int64CustomEmpties = []int64{}

	// Redacting field: Int64CustomNested
	for k := range x.Int64CustomNested {
		x.Int64CustomNested[k] = 0
	}

	// Redacting field: Int64Customs
	for k := range x.Int64Customs {
		x.Int64Customs[k] = 64
	}

	// Redacting field: Uint32S
	x.Uint32S = nil

	// Redacting field: Uint32CustomEmpties
	x.Uint32CustomEmpties = []uint32{}

	// Redacting field: Uint32CustomNested
	for k := range x.Uint32CustomNested {
		x.Uint32CustomNested[k] = 0
	}

	// Redacting field: Uint32Customs
	for k := range x.Uint32Customs {
		x.Uint32Customs[k] = 32
	}

	// Redacting field: Uint64S
	x.Uint64S = nil

	// Redacting field: Uint64CustomEmpties
	x.Uint64CustomEmpties = []uint64{}

	// Redacting field: Uint64CustomNested
	for k := range x.Uint64CustomNested {
		x.Uint64CustomNested[k] = 0
	}

	// Redacting field: Uint64Customs
	for k := range x.Uint64Customs {
		x.Uint64Customs[k] = 64
	}

	// Redacting field: Sint32S
	x.Sint32S = nil

	// Redacting field: Sint32CustomEmpties
	x.Sint32CustomEmpties = []int32{}

	// Redacting field: Sint32CustomNested
	for k := range x.Sint32CustomNested {
		x.Sint32CustomNested[k] = 0
	}

	// Redacting field: Sint32Customs
	for k := range x.Sint32Customs {
		x.Sint32Customs[k] = 32
	}

	// Redacting field: Sint64S
	x.Sint64S = nil

	// Redacting field: Sint64CustomEmpties
	x.Sint64CustomEmpties = []int64{}

	// Redacting field: Sint64CustomNested
	for k := range x.Sint64CustomNested {
		x.Sint64CustomNested[k] = 0
	}

	// Redacting field: Sint64Customs
	for k := range x.Sint64Customs {
		x.Sint64Customs[k] = 64
	}

	// Redacting field: Fixed32S
	x.Fixed32S = nil

	// Redacting field: Fixed32CustomEmpties
	x.Fixed32CustomEmpties = []uint32{}

	// Redacting field: Fixed32CustomNested
	for k := range x.Fixed32CustomNested {
		x.Fixed32CustomNested[k] = 0
	}

	// Redacting field: Fixed32Customs
	for k := range x.Fixed32Customs {
		x.Fixed32Customs[k] = 32
	}

	// Redacting field: Fixed64S
	x.Fixed64S = nil

	// Redacting field: Fixed64CustomEmpties
	x.Fixed64CustomEmpties = []uint64{}

	// Redacting field: Fixed64CustomNested
	for k := range x.Fixed64CustomNested {
		x.Fixed64CustomNested[k] = 0
	}

	// Redacting field: Fixed64Customs
	for k := range x.Fixed64Customs {
		x.Fixed64Customs[k] = 64
	}

	// Redacting field: Sfixed32S
	x.Sfixed32S = nil

	// Redacting field: Sfixed32CustomEmpties
	x.Sfixed32CustomEmpties = []int32{}

	// Redacting field: Sfixed32CustomNested
	for k := range x.Sfixed32CustomNested {
		x.Sfixed32CustomNested[k] = 0
	}

	// Redacting field: Sfixed32Customs
	for k := range x.Sfixed32Customs {
		x.Sfixed32Customs[k] = 32
	}

	// Redacting field: Sfixed64S
	x.Sfixed64S = nil

	// Redacting field: Sfixed64CustomEmpties
	x.Sfixed64CustomEmpties = []int64{}

	// Redacting field: Sfixed64CustomNested
	for k := range x.Sfixed64CustomNested {
		x.Sfixed64CustomNested[k] = 0
	}

	// Redacting field: Sfixed64Customs
	for k := range x.Sfixed64Customs {
		x.Sfixed64Customs[k] = 64
	}

	// Redacting field: Bools
	x.Bools = nil

	// Redacting field: BoolCustomEmpties
	x.BoolCustomEmpties = []bool{}

	// Redacting field: BoolCustomNested
	for k := range x.BoolCustomNested {
		x.BoolCustomNested[k] = false
	}

	// Redacting field: BoolCustoms
	for k := range x.BoolCustoms {
		x.BoolCustoms[k] = true
	}

	// Redacting field: Strings
	x.Strings = nil

	// Redacting field: StringCustomEmpties
	x.StringCustomEmpties = []string{}

	// Redacting field: StringCustomNested
	for k := range x.StringCustomNested {
		x.StringCustomNested[k] = "REDACTED"
	}

	// Redacting field: StringCustoms
	for k := range x.StringCustoms {
		x.StringCustoms[k] = `redacted-custom-value`
	}

	// Redacting field: Bytess
	x.Bytess = nil

	// Redacting field: BytesCustomEmpties
	x.BytesCustomEmpties = [][]byte{}

	// Redacting field: BytesCustomNested
	for k := range x.BytesCustomNested {
		x.BytesCustomNested[k] = nil
	}

	// Redacting field: BytesCustoms
	for k := range x.BytesCustoms {
		x.BytesCustoms[k] = []byte(`redacted-custom-value`)
	}

	// Redacting field: Enums
	x.Enums = nil

	// Redacting field: EnumCustomEmpties
	x.EnumCustomEmpties = []TestEnum{}

	// Redacting field: EnumCustomNested
	for k := range x.EnumCustomNested {
		x.EnumCustomNested[k] = 0
	}

	// Redacting field: EnumCustoms
	for k := range x.EnumCustoms {
		x.EnumCustoms[k] = 2
	}

	// Redacting field: Messages
	x.Messages = nil

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
}
