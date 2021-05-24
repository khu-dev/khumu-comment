// Code generated by entc, DO NOT EDIT.

package likecomment

const (
	// Label holds the string label denoting the likecomment type in the database.
	Label = "like_comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeLikedBy holds the string denoting the likedby edge name in mutations.
	EdgeLikedBy = "likedBy"
	// EdgeAbout holds the string denoting the about edge name in mutations.
	EdgeAbout = "about"
	// KhumuUserFieldID holds the string denoting the ID field of the KhumuUser.
	KhumuUserFieldID = "username"
	// Table holds the table name of the likecomment in the database.
	Table = "comment_likecomment"
	// LikedByTable is the table the holds the likedBy relation/edge.
	LikedByTable = "comment_likecomment"
	// LikedByInverseTable is the table name for the KhumuUser entity.
	// It exists in this package in order to avoid circular dependency with the "khumuuser" package.
	LikedByInverseTable = "user_khumuuser"
	// LikedByColumn is the table column denoting the likedBy relation/edge.
	LikedByColumn = "user_id"
	// AboutTable is the table the holds the about relation/edge.
	AboutTable = "comment_likecomment"
	// AboutInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	AboutInverseTable = "comment_comment"
	// AboutColumn is the table column denoting the about relation/edge.
	AboutColumn = "comment_id"
)

// Columns holds all SQL columns for likecomment fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comment_likecomment"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"comment_id",
	"user_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
