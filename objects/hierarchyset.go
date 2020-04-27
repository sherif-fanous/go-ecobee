package objects

// The HierarchySet object contains the name, path and additional information
// associated with a management set.
//
// All set paths start with '/' which is equivalent to the "My Sets" set. A set
// path without a leading'/' signifies a special set, such as "Unassigned".
type HierarchySet struct {
	// The name of the hierarchy set.
	SetName *string `json:"setName,omitempty"`

	// The full path from the root to the hierarchy set.
	SetPath *string `json:"setPath,omitempty"`

	// The list of child hierarchy sets.
	Children []HierarchySet `json:"children,omitempty"`

	// The list of hierarchy privileges assigned to this hierarchy set.
	Privileges []HierarchyPrivilege `json:"privileges,omitempty"`

	// The list of thermostats assigned to this hierarchy set..
	Thermostats []string `json:"thermostats,omitempty"`
}
