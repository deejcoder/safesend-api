// Code generated by entc, DO NOT EDIT.

package ent

import (
	"SafeSend/pkg/ent/entity"
	"SafeSend/pkg/ent/group"
	"SafeSend/pkg/ent/schema"
	"SafeSend/pkg/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	entityFields := schema.Entity{}.Fields()
	_ = entityFields
	// entityDescDateCreated is the schema descriptor for date_created field.
	entityDescDateCreated := entityFields[2].Descriptor()
	// entity.DefaultDateCreated holds the default value on creation for the date_created field.
	entity.DefaultDateCreated = entityDescDateCreated.Default.(time.Time)
	// entityDescDateModified is the schema descriptor for date_modified field.
	entityDescDateModified := entityFields[3].Descriptor()
	// entity.DefaultDateModified holds the default value on creation for the date_modified field.
	entity.DefaultDateModified = entityDescDateModified.Default.(time.Time)
	// entityDescID is the schema descriptor for id field.
	entityDescID := entityFields[0].Descriptor()
	// entity.DefaultID holds the default value on creation for the id field.
	entity.DefaultID = entityDescID.Default.(func() uuid.UUID)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescMaxParticipants is the schema descriptor for maxParticipants field.
	groupDescMaxParticipants := groupFields[2].Descriptor()
	// group.MaxParticipantsValidator is a validator for the "maxParticipants" field. It is called by the builders before save.
	group.MaxParticipantsValidator = func() func(int32) error {
		validators := groupDescMaxParticipants.Validators
		fns := [...]func(int32) error{
			validators[0].(func(int32) error),
			validators[1].(func(int32) error),
			validators[2].(func(int32) error),
		}
		return func(maxParticipants int32) error {
			for _, fn := range fns {
				if err := fn(maxParticipants); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// groupDescInviteOnly is the schema descriptor for inviteOnly field.
	groupDescInviteOnly := groupFields[3].Descriptor()
	// group.DefaultInviteOnly holds the default value on creation for the inviteOnly field.
	group.DefaultInviteOnly = groupDescInviteOnly.Default.(bool)
	// groupDescDateCreated is the schema descriptor for dateCreated field.
	groupDescDateCreated := groupFields[4].Descriptor()
	// group.DefaultDateCreated holds the default value on creation for the dateCreated field.
	group.DefaultDateCreated = groupDescDateCreated.Default.(time.Time)
	// groupDescDateModified is the schema descriptor for dateModified field.
	groupDescDateModified := groupFields[5].Descriptor()
	// group.DefaultDateModified holds the default value on creation for the dateModified field.
	group.DefaultDateModified = groupDescDateModified.Default.(time.Time)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescDateCreated is the schema descriptor for dateCreated field.
	userDescDateCreated := userFields[4].Descriptor()
	// user.DefaultDateCreated holds the default value on creation for the dateCreated field.
	user.DefaultDateCreated = userDescDateCreated.Default.(time.Time)
	// userDescDateModified is the schema descriptor for dateModified field.
	userDescDateModified := userFields[5].Descriptor()
	// user.DefaultDateModified holds the default value on creation for the dateModified field.
	user.DefaultDateModified = userDescDateModified.Default.(time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}