package Optional

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	uuid2 "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"testing"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUUID(t *testing.T) {
	testUUID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue uuid.UUID
		}
		output uuid.UUID
	}{
		{
			name: "NonExistingKey",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": testUUID,
				}, Key: "NonExisting", DefaultValue: uuid.UUID{}},
			output: uuid.UUID{},
		},
		{
			name: "StringCase",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": "11111111-1111-1111-1111-111111111111",
				}, Key: "value", DefaultValue: uuid.UUID{}},
			output: testUUID,
		},
		{
			name: "StringCaseError",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": "not a valid UUID",
				}, Key: "value", DefaultValue: uuid.UUID{}},
			output: uuid.UUID{},
		},
		{
			name: "BinaryCase",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": primitive.Binary{Data: []byte{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17}},
				}, Key: "value", DefaultValue: uuid.UUID{}},
			output: testUUID,
		},
		{
			name: "BinaryCaseError",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": primitive.Binary{Data: []byte{1, 1, 1, 1}},
				}, Key: "value", DefaultValue: uuid.UUID{}},
			output: uuid.UUID{},
		},
		{
			name: "ByteArrayCase",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": []byte{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
				}, Key: "value", DefaultValue: uuid.UUID{}},
			output: testUUID,
		},
		{
			name: "ByteArrayCaseError",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": []byte{1, 1, 1, 1},
				}, Key: "value", DefaultValue: uuid.UUID{}},
			output: uuid.UUID{},
		},
		{
			name: "CaseUUID",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": testUUID,
				}, Key: "value", DefaultValue: uuid.UUID{}},
			output: testUUID,
		},
		{
			name: "CaseUUIDMongo",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": uuid2.UUID{},
				}, Key: "value", DefaultValue: testUUID},
			output: uuid.UUID{},
		},
		{
			name: "UnsupportedType",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"value": []string{},
				}, Key: "value", DefaultValue: uuid.UUID{}},
			output: uuid.UUID{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UUID(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if fmt.Sprint(got) != fmt.Sprint(tt.output) {
				t.Errorf(
					"expected UUID(%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output,
					got,
				)
			}
		})
	}

}