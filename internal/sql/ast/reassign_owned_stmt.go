package ast

import ()

type ReassignOwnedStmt struct {
	Roles   *List
	Newrole *RoleSpec
}

func (n *ReassignOwnedStmt) Pos() int {
	return 0
}
