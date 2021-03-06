package diff

import (
	"reflect"

	"github.com/integration-system/bellows"
	"github.com/integration-system/go-cmp/cmp"
)

type Operation string

const (
	Add    Operation = "ADD"
	Delete Operation = "DELETE"
	Change Operation = "CHANGE"

	ArrayAdd    Operation = "ARRAY_ADD"
	ArrayDelete Operation = "ARRAY_DELETE"
	ArrayChange Operation = "ARRAY_CHANGE"
	ArraySwap   Operation = "ARRAY_SWAP"
)

type DiffDescriptor struct {
	OldValue       interface{} `json:"oldValue,omitempty"`
	NewValue       interface{} `json:"newValue,omitempty"`
	Path           string      `json:"path"`
	Operation      Operation   `json:"operation"`
	AdditionalData interface{} `json:"additionalData,omitempty"`
	OldIndex       *int        `json:"oldIndex,omitempty"`
	NewIndex       *int        `json:"newIndex,omitempty"`
}

type (
	Delta            []*DiffDescriptor
	extensionHandler func(diff *DiffDescriptor)
)

func EvalDiff(left, right map[string]interface{}, opts ...Option) (bool, Delta) {
	c := NewDiffCollector(opts...)
	return cmp.Equal(left, right, c), c.Delta
}

func FlattenDelta(delta Delta) map[string]*DiffDescriptor {
	result := make(map[string]*DiffDescriptor, len(delta)*2)
	extensionDelta(func(diff *DiffDescriptor) {
		result[diff.Path] = diff
	}, delta)
	return result
}

func ExtensionDelta(delta Delta) Delta {
	result := make([]*DiffDescriptor, 0)
	extensionDelta(func(diff *DiffDescriptor) {
		result = append(result, diff)
	}, delta)
	return result
}

func extensionDelta(callback extensionHandler, delta Delta) {
	for _, desc := range delta {
		switch {
		case (desc.Operation == Add || desc.Operation == ArrayAdd) && desc.NewValue != nil:
			rt := reflect.TypeOf(desc.NewValue)
			switch rt.Kind() {
			case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
				flattenValue := bellows.Flatten(desc.NewValue)
				for path, value := range flattenValue {
					newPath := getNewPath(desc.Path, path, rt.Kind())
					callback(&DiffDescriptor{NewValue: value, Path: newPath, Operation: Add})
				}
			default:
				callback(desc)
			}
		case (desc.Operation == Delete || desc.Operation == ArrayDelete) && desc.OldValue != nil:
			rt := reflect.TypeOf(desc.OldValue)
			switch rt.Kind() {
			case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
				flattenValue := bellows.Flatten(desc.OldValue)
				for path, value := range flattenValue {
					newPath := getNewPath(desc.Path, path, rt.Kind())
					callback(&DiffDescriptor{OldValue: value, Path: newPath, Operation: Delete})
				}
			default:
				callback(desc)
			}
		default:
			callback(desc)
		}
	}
}

func getNewPath(base string, path string, kind reflect.Kind) string {
	switch kind {
	case reflect.Slice, reflect.Array:
		return base + path
	default:
		return base + "." + path
	}
}
