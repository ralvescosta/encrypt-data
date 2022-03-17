package io

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadInput(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		mock(nil, nil, nil, nil, nil)

		result, err := ReadInput()

		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return err if some error occur getting the input file path", func(t *testing.T) {
		mock(errors.New(""), nil, nil, nil, nil)

		_, err := ReadInput()

		assert.Error(t, err)
	})

	t.Run("should return err if some error occur whiling reading the input file", func(t *testing.T) {
		mock(nil, errors.New(""), nil, nil, nil)

		_, err := ReadInput()

		assert.Error(t, err)
	})

	t.Run("should return err if some error occur whiling parsing input file", func(t *testing.T) {
		mock(nil, nil, nil, errors.New(""), nil)

		_, err := ReadInput()

		assert.Error(t, err)
	})
}

func Test_WriteOutput(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		mock(nil, nil, nil, nil, nil)

		err := WriteOutput([]byte(""))

		assert.NoError(t, err)
	})

	t.Run("should return err if some error occur getting the output file path", func(t *testing.T) {
		mock(errors.New(""), nil, nil, nil, nil)

		err := WriteOutput([]byte(""))

		assert.Error(t, err)
	})

	t.Run("should return err if some error occur whiling parsing output content", func(t *testing.T) {
		mock(nil, nil, nil, nil, errors.New(""))

		err := WriteOutput([]byte(""))

		assert.Error(t, err)
	})

	t.Run("should return err if some error occur writing the output file", func(t *testing.T) {
		mock(nil, nil, errors.New(""), nil, nil)

		err := WriteOutput([]byte(""))

		assert.Error(t, err)
	})
}

func mock(osGetwdErr, readFileErr, writeFileErr, jsonUnmarshalErr, jsonMarshalIdentErr error) {
	osGetwd = func() (string, error) {
		return "dir", osGetwdErr
	}
	readFile = func(filename string) ([]byte, error) {
		return []byte(nil), readFileErr
	}
	writeFile = func(filename string, data []byte, perm fs.FileMode) error {
		return writeFileErr
	}
	jsonUnmarshal = func(data []byte, v any) error {
		v = &InputData{}
		return jsonUnmarshalErr
	}
	jsonMarshalIndent = func(v any, prefix string, indent string) ([]byte, error) {
		return []byte(nil), jsonMarshalIdentErr
	}
}
