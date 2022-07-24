package protosensitive

import (
	"testing"

	examplev1 "go.buf.build/protocolbuffers/go/einride/sensitive/einride/sensitive/example/v1"
	"google.golang.org/protobuf/testing/protocmp"
	"gotest.tools/v3/assert"
)

func TestRedact(t *testing.T) {
	t.Run("no message copy for non-sensitive message", func(t *testing.T) {
		input := &examplev1.ExampleMessage{
			NonSensitiveField: "foo",
			NonStringField:    7,
			Nested: &examplev1.ExampleMessage_Nested{
				NonSensitiveField: "bar",
				NonStringField:    8,
			},
			RepeatedNested: []*examplev1.ExampleMessage_Nested{
				{NonSensitiveField: "baz"},
			},
			MapNested: map[int64]*examplev1.ExampleMessage_Nested{
				6: {
					NonSensitiveField: "apa",
				},
			},
		}
		output := Redact(input)
		assert.Equal(t, input, output)
	})

	t.Run("redacts sensitive fields", func(t *testing.T) {
		input := &examplev1.ExampleMessage{
			SensitiveField:    "top secret",
			NonSensitiveField: "foo",
			NonStringField:    7,
			Nested: &examplev1.ExampleMessage_Nested{
				SensitiveField:    "super secret",
				NonSensitiveField: "bar",
				NonStringField:    8,
			},
			RepeatedNested: []*examplev1.ExampleMessage_Nested{
				{SensitiveField: "mega secret"},
				{NonSensitiveField: "baz"},
			},
			MapNested: map[int64]*examplev1.ExampleMessage_Nested{
				6: {
					NonSensitiveField:     "apa",
					AnotherSensitiveField: "ultra secret",
				},
			},
		}
		expected := &examplev1.ExampleMessage{
			SensitiveField:    "<redacted>",
			NonSensitiveField: "foo",
			NonStringField:    7,
			Nested: &examplev1.ExampleMessage_Nested{
				SensitiveField:    "<redacted>",
				NonSensitiveField: "bar",
				NonStringField:    8,
			},
			RepeatedNested: []*examplev1.ExampleMessage_Nested{
				{SensitiveField: "<redacted>"},
				{NonSensitiveField: "baz"},
			},
			MapNested: map[int64]*examplev1.ExampleMessage_Nested{
				6: {
					NonSensitiveField:     "apa",
					AnotherSensitiveField: "<redacted>",
				},
			},
		}
		actual := Redact(input)
		assert.DeepEqual(t, expected, actual, protocmp.Transform())
	})
}
