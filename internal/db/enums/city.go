//go:generate ../bin/go-enum -f=$GOFILE -a "+:Plus,#:Sharp"

package enums

// City x ENUM(
// Moscow,
// Yekaterinburg
// )
type City int32
