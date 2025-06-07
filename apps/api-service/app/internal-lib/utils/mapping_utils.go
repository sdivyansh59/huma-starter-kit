package utils

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Todo : uuid must of version 7

// ZeroUUID is the zero UUID, mostly used in tests.
var ZeroUUID = uuid.MustParse("00000000-0000-0000-0000-000000000000")

// ParseUUID will parse a UUID.
func ParseUUID(input string) (uuid.UUID, error) {
	res, err := uuid.Parse(input)
	if err != nil {
		return uuid.UUID{}, errors.New("could not parse uuid")
	}

	return res, err
}

// ParseMultipleUUIDs will parse multiple UUIDs.
func ParseMultipleUUIDs(input []string) ([]uuid.UUID, error) {
	result := make([]uuid.UUID, len(input))

	for _, id := range input {
		parsed, err := ParseUUID(id)
		if err != nil {
			return nil, err
		}

		result = append(result, parsed)
	}

	return result, nil
}

// ToPointer returns the pointer of the given value.
func ToPointer[T any](value T) *T {
	return &value
}

// SafeDereference returns the value of the pointer if it is not nil, otherwise returns the default value.
func SafeDereference[T any](ptr *T, def T) T {
	if ptr == nil {
		return def
	}

	return *ptr
}

func BooleanToString(b bool) string {
	if b {
		return "true"
	}

	return "false"
}

func StringToBoolean(s string) bool {
	return s == "true"
}

// Deprecated: use ToPointer instead
// StringPointer returns the pointer to the given string.
func StringPointer(v string) *string {
	return &v
}

// Deprecated: use ToPointer instead
// Int64Pointer returns the pointer to the given input.
func Int64Pointer(v int64) *int64 {
	return &v
}

// Deprecated: use ToPointer instead
// IntPointer returns the pointer to the given input.
func IntPointer(v int) *int {
	return &v
}

// Deprecated: use ToPointer instead
// FloatPointer returns the pointer to the given input.
func FloatPointer(v float64) *float64 {
	return &v
}

// Deprecated: use ToPointer instead
// TimePointer returns the pointer to a given time.
func TimePointer(v time.Time) *time.Time {
	return &v
}

// BoolPointer returns the pointer to a given bool.
// Deprecated.
func BoolPointer(b bool) *bool {
	return &b
}

// Deprecated: use SafeDereference instead
// SafeDereferenceString will safely dereference the given string pointer.
// It will return an empty string if v is nil.
func SafeDereferenceString(v *string) string {
	if v == nil {
		return ""
	}

	return *v
}

// Deprecated: use SafeDereference instead
// SafeDereferenceInt will safely dereference the given int pointer.
// It will return 0 if v is nil.
func SafeDereferenceInt(v *int) int {
	if v == nil {
		return 0
	}

	return *v
}

// Deprecated: use SafeDereference instead
// SafeDereferenceInt64 will safely dereference the given int pointer.
// It will return 0 if v is nil.
func SafeDereferenceInt64(v *int64) int64 {
	if v == nil {
		return 0
	}

	return *v
}

// Deprecated: use SafeDereference instead
// SafeDereferenceBool will safely dereference the given bool pointer.
// It will return false if v is nil.
func SafeDereferenceBool(v *bool) bool {
	if v == nil {
		return false
	}

	return *v
}
