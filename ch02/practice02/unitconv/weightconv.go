package unitconv

import (
	"fmt"
)

type Kgm float64
type Lb float64

func (k Kgm) String() string { return fmt.Sprintf("%gkg", k) }
func (l Lb) String() string  { return fmt.Sprintf("%glb", l) }
