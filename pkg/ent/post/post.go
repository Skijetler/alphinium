// Code generated by ent, DO NOT EDIT.

package post

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldThreadID holds the string denoting the thread_id field in the database.
	FieldThreadID = "thread_id"
	// EdgeThread holds the string denoting the thread edge name in mutations.
	EdgeThread = "thread"
	// EdgeDescribedThread holds the string denoting the described_thread edge name in mutations.
	EdgeDescribedThread = "described_thread"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeAttachments holds the string denoting the attachments edge name in mutations.
	EdgeAttachments = "attachments"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// ThreadTable is the table that holds the thread relation/edge.
	ThreadTable = "posts"
	// ThreadInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	ThreadInverseTable = "threads"
	// ThreadColumn is the table column denoting the thread relation/edge.
	ThreadColumn = "thread_id"
	// DescribedThreadTable is the table that holds the described_thread relation/edge.
	DescribedThreadTable = "threads"
	// DescribedThreadInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	DescribedThreadInverseTable = "threads"
	// DescribedThreadColumn is the table column denoting the described_thread relation/edge.
	DescribedThreadColumn = "description_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "posts"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// AttachmentsTable is the table that holds the attachments relation/edge.
	AttachmentsTable = "attachments"
	// AttachmentsInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentsInverseTable = "attachments"
	// AttachmentsColumn is the table column denoting the attachments relation/edge.
	AttachmentsColumn = "post_id"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldMessage,
	FieldDate,
	FieldUserID,
	FieldThreadID,
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
