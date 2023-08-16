package table

import (
	"github.com/silence-lml/dxf/format"
	"github.com/silence-lml/dxf/handle"
)

// SymbolTable is interface for AcDbSymbolTableRecord.
type SymbolTable interface {
	IsSymbolTable() bool
	Format(format.Formatter)
	Handle() int
	SetHandle(*int)
	SetOwner(handle.Handler)
	Name() string
}
