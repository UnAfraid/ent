// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package usergroup

import (
	"time"
)

const (
	// Label holds the string label denoting the usergroup type in the database.
	Label = "user_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldJoinedAt holds the string denoting the joined_at field in the database.
	FieldJoinedAt = "joined_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGroupID holds the string denoting the group_id field in the database.
	FieldGroupID = "group_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the usergroup in the database.
	Table = "user_groups"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_groups"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "user_groups"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_id"
)

// Columns holds all SQL columns for usergroup fields.
var Columns = []string{
	FieldID,
	FieldJoinedAt,
	FieldUserID,
	FieldGroupID,
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
	// DefaultJoinedAt holds the default value on creation for the "joined_at" field.
	DefaultJoinedAt func() time.Time
)