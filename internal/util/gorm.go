package util

import "go/types"

var chainingMethods = map[string]bool{
	"Create":          true,
	"CreateInBatches": true,
	"Save":            true,
	"First":           true,
	"Take":            true,
	"Last":            true,
	"Find":            true,
	"FindInBatches":   true,
	"FirstOrInit":     true,
	"FirstOrCreate":   true,
	"Update":          true,
	"Updates":         true,
	"UpdateColumn":    true,
	"UpdateColumns":   true,
	"Delete":          true,
	"Count":           true,
	"Row":             true,
	"Rows":            true,
	"Scan":            true,
	"Pluck":           true,
	"ScanRows":        true,
	"Connection":      true,
	"Transaction":     true,
	"Begin":           true,
	"Commit":          true,
	"Rollback":        true,
	"SavePoint":       true,
	"RollbackTo":      true,
	"Exec":            true,
}

var finisherMethods = map[string]bool{
	"Model":      true,
	"Clauses":    true,
	"Table":      true,
	"Distinct":   true,
	"Select":     true,
	"Omit":       true,
	"MapColumns": true,
	"Where":      true,
	"Not":        true,
	"Or":         true,
	"Joins":      true,
	"InnerJoins": true,
	"joins":      true,
	"Group":      true,
	"Having":     true,
	"Order":      true,
	"Limit":      true,
	"Offset":     true,
	"Scopes":     true,
	"Preload":    true,
	"Attrs":      true,
	"Assign":     true,
	"Unscoped":   true,
	"Raw":        true,
}

var newSessionMethods = map[string]bool{
	"Session":     true,
	"WithContext": true,
	"Debug":       true,
}

func IsGormChainingMethod(name string) bool {
	return chainingMethods[name]
}

func IsGormFinisherMethod(name string) bool {
	return finisherMethods[name]
}

func IsGormNewSessionMethod(name string) bool {
	return newSessionMethods[name]
}

func IsGormDBType(t types.Type) bool {
	ptr, ok := t.(*types.Pointer)
	if !ok {
		return false
	}

	named, ok := ptr.Elem().(*types.Named)
	if !ok {
		return false
	}

	return named.Obj().Pkg() != nil &&
		named.Obj().Pkg().Path() == "gorm.io/gorm" &&
		named.Obj().Name() == "DB"
}
