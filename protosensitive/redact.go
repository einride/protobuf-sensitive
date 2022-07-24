package protosensitive

import (
	sensitivev1 "go.buf.build/protocolbuffers/go/einride/sensitive/einride/sensitive/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protorange"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Redact sensitive fields from the input message.
// Currently, only fields of string type are supported.
func Redact(input proto.Message) proto.Message {
	var hasSensitiveFields bool
	_ = protorange.Range(input.ProtoReflect(), func(values protopath.Values) error {
		last := values.Index(-1)
		if _, ok := last.Value.Interface().(string); !ok {
			return nil
		}
		if last.Step.Kind() != protopath.FieldAccessStep {
			return nil
		}
		if !proto.GetExtension(last.Step.FieldDescriptor().Options(), sensitivev1.E_Sensitive).(bool) {
			return nil
		}
		hasSensitiveFields = true
		return protorange.Terminate
	})
	if !hasSensitiveFields {
		return input
	}
	output := proto.Clone(input)
	_ = protorange.Range(output.ProtoReflect(), func(values protopath.Values) error {
		last := values.Index(-1)
		if _, ok := last.Value.Interface().(string); !ok {
			return nil
		}
		if last.Step.Kind() != protopath.FieldAccessStep {
			return nil
		}
		if !proto.GetExtension(last.Step.FieldDescriptor().Options(), sensitivev1.E_Sensitive).(bool) {
			return nil
		}
		values.Index(-2).Value.Message().Set(last.Step.FieldDescriptor(), protoreflect.ValueOfString("<redacted>"))
		return nil
	})
	return output
}
