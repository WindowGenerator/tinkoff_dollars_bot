//go:generate ../bin/go-enum -f=$GOFILE -a "+:Plus,#:Sharp"

package enums

// Bank x ENUM(
// All,
// Tinkoff,
// Sber,
// VTB
// Alpha
// Raiffaizen
// GasProm
// )
type Bank int32
