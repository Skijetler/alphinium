// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRegistrationDate holds the string denoting the registration_date field in the database.
	FieldRegistrationDate = "registration_date"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeAttachments holds the string denoting the attachments edge name in mutations.
	EdgeAttachments = "attachments"
	// Table holds the table name of the user in the database.
	Table = "users"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "user_metadata"
	// MetadataInverseTable is the table name for the UserMetadata entity.
	// It exists in this package in order to avoid circular dependency with the "usermetadata" package.
	MetadataInverseTable = "user_metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "user_id"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "user_id"
	// AttachmentsTable is the table that holds the attachments relation/edge.
	AttachmentsTable = "attachments"
	// AttachmentsInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentsInverseTable = "attachments"
	// AttachmentsColumn is the table column denoting the attachments relation/edge.
	AttachmentsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEmail,
	FieldPassword,
	FieldRegistrationDate,
	FieldDisabled,
}

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
	// DefaultDisabled holds the default value on creation for the "disabled" field.
	DefaultDisabled bool
)