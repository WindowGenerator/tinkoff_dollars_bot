//go:generate ../bin/go-enum -f=$GOFILE -a "+:Plus,#:Sharp"

package enums

// Currency x ENUM(
// Any
// USD,
// RUB,
// EUR
// )
type Currency int32
