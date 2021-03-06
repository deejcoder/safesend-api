// Code generated by entc, DO NOT EDIT.

package accesstoken

import (
	"SafeSend/pkg/interfaces"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the accesstoken type in the database.
	Label = "access_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTokenProvider holds the string denoting the token_provider field in the database.
	FieldTokenProvider = "token_provider"
	// FieldAccessToken holds the string denoting the access_token field in the database.
	FieldAccessToken = "access_token"
	// FieldRefreshToken holds the string denoting the refresh_token field in the database.
	FieldRefreshToken = "refresh_token"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// FieldDateCreated holds the string denoting the date_created field in the database.
	FieldDateCreated = "date_created"
	// FieldDateModified holds the string denoting the date_modified field in the database.
	FieldDateModified = "date_modified"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the accesstoken in the database.
	Table = "access_tokens"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_access_tokens"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for accesstoken fields.
var Columns = []string{
	FieldID,
	FieldTokenProvider,
	FieldAccessToken,
	FieldRefreshToken,
	FieldExpiry,
	FieldDateCreated,
	FieldDateModified,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "access_token_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AccessTokenValidator is a validator for the "access_token" field. It is called by the builders before save.
	AccessTokenValidator func(string) error
	// RefreshTokenValidator is a validator for the "refresh_token" field. It is called by the builders before save.
	RefreshTokenValidator func(string) error
	// DefaultDateCreated holds the default value on creation for the "date_created" field.
	DefaultDateCreated time.Time
	// DefaultDateModified holds the default value on creation for the "date_modified" field.
	DefaultDateModified time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// TokenProviderValidator is a validator for the "token_provider" field enum values. It is called by the builders before save.
func TokenProviderValidator(tp interfaces.TokenProvider) error {
	switch tp {
	case "GOOGLE":
		return nil
	default:
		return fmt.Errorf("accesstoken: invalid enum value for token_provider field: %q", tp)
	}
}
