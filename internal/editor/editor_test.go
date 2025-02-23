package editor

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_updateResourceEditor(t *testing.T) {
	SkipEditor = true

	resource := &struct {
		ID   string
		Name string
	}{
		"uuid",
		"name",
	}
	updateRequest := &struct {
		ID   string
		Name string
	}{
		"uuid",
		"",
	}

	_, err := updateResourceEditor(resource, updateRequest, &Config{})
	assert.Nil(t, err)
}

func Test_updateResourceEditor_pointers(t *testing.T) {
	SkipEditor = true

	type UpdateRequest struct {
		ID   string
		Name *string
	}
	resource := &struct {
		ID   string
		Name string
	}{
		"uuid",
		"name",
	}

	updateRequest := &UpdateRequest{
		"uuid",
		nil,
	}

	editedUpdateRequestI, err := updateResourceEditor(resource, updateRequest, &Config{})
	assert.Nil(t, err)
	editedUpdateRequest := editedUpdateRequestI.(*UpdateRequest)

	assert.NotNil(t, editedUpdateRequest.Name)
	assert.Equal(t, resource.Name, *editedUpdateRequest.Name)
}

func Test_updateResourceEditor_map(t *testing.T) {
	SkipEditor = true

	type UpdateRequest struct {
		ID  string             `json:"id"`
		Env *map[string]string `json:"env"`
	}
	resource := &struct {
		ID  string            `json:"id"`
		Env map[string]string `json:"env"`
	}{
		"uuid",
		map[string]string{
			"foo": "bar",
		},
	}

	updateRequest := &UpdateRequest{
		"uuid",
		nil,
	}

	editedUpdateRequestI, err := updateResourceEditor(resource, updateRequest, &Config{
		editedResource: `
id: uuid
env: {}
`,
	})
	assert.Nil(t, err)
	editedUpdateRequest := editedUpdateRequestI.(*UpdateRequest)
	assert.NotNil(t, editedUpdateRequest.Env)
	assert.True(t, len(*editedUpdateRequest.Env) == 0)
}
