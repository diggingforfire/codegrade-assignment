// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"codegrade.com/take-home/ent/post"
	"codegrade.com/take-home/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[1].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescContent is the schema descriptor for content field.
	postDescContent := postFields[2].Descriptor()
	// post.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	post.ContentValidator = postDescContent.Validators[0].(func(string) error)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[3].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)
	// postDescID is the schema descriptor for id field.
	postDescID := postFields[0].Descriptor()
	// post.DefaultID holds the default value on creation for the id field.
	post.DefaultID = postDescID.Default.(func() uuid.UUID)
}
