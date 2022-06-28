// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccessTokensColumns holds the columns for the "access_tokens" table.
	AccessTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "token_provider", Type: field.TypeEnum, Enums: []string{"GOOGLE"}},
		{Name: "access_token", Type: field.TypeString},
		{Name: "refresh_token", Type: field.TypeString},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_modified", Type: field.TypeTime},
	}
	// AccessTokensTable holds the schema information for the "access_tokens" table.
	AccessTokensTable = &schema.Table{
		Name:       "access_tokens",
		Columns:    AccessTokensColumns,
		PrimaryKey: []*schema.Column{AccessTokensColumns[0]},
	}
	// EntitiesColumns holds the columns for the "entities" table.
	EntitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "entity_type", Type: field.TypeEnum, Enums: []string{"USER", "GROUP"}},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_modified", Type: field.TypeTime},
		{Name: "date_deleted", Type: field.TypeTime, Nullable: true},
	}
	// EntitiesTable holds the schema information for the "entities" table.
	EntitiesTable = &schema.Table{
		Name:       "entities",
		Columns:    EntitiesColumns,
		PrimaryKey: []*schema.Column{EntitiesColumns[0]},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "group_name", Type: field.TypeString},
		{Name: "max_participants", Type: field.TypeInt32},
		{Name: "invite_only", Type: field.TypeBool, Default: false},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_modified", Type: field.TypeTime},
		{Name: "date_deleted", Type: field.TypeTime, Nullable: true},
		{Name: "entity_groups", Type: field.TypeUUID, Nullable: true},
		{Name: "group_users", Type: field.TypeUUID, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_entities_groups",
				Columns:    []*schema.Column{GroupsColumns[7]},
				RefColumns: []*schema.Column{EntitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "groups_users_users",
				Columns:    []*schema.Column{GroupsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString},
		{Name: "date_accessed", Type: field.TypeTime, Nullable: true},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_modified", Type: field.TypeTime},
		{Name: "deleted_date", Type: field.TypeTime, Nullable: true},
		{Name: "entity_users", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_entities_users",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{EntitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserAccessTokensColumns holds the columns for the "user_access_tokens" table.
	UserAccessTokensColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "access_token_id", Type: field.TypeUUID},
	}
	// UserAccessTokensTable holds the schema information for the "user_access_tokens" table.
	UserAccessTokensTable = &schema.Table{
		Name:       "user_access_tokens",
		Columns:    UserAccessTokensColumns,
		PrimaryKey: []*schema.Column{UserAccessTokensColumns[0], UserAccessTokensColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_access_tokens_user_id",
				Columns:    []*schema.Column{UserAccessTokensColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_access_tokens_access_token_id",
				Columns:    []*schema.Column{UserAccessTokensColumns[1]},
				RefColumns: []*schema.Column{AccessTokensColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccessTokensTable,
		EntitiesTable,
		GroupsTable,
		UsersTable,
		UserAccessTokensTable,
	}
)

func init() {
	GroupsTable.ForeignKeys[0].RefTable = EntitiesTable
	GroupsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = EntitiesTable
	UserAccessTokensTable.ForeignKeys[0].RefTable = UsersTable
	UserAccessTokensTable.ForeignKeys[1].RefTable = AccessTokensTable
}
