package models

import (
	"github.com/uptrace/bun"
)

type Calc struct {
	bun.BaseModel `bun:"calcs"`

	ID  int32 `json:"id" bun:"id,pk"`
	A   int32 `bun:"a"`
	B   int32 `bun:"b"`
	Res int32 `bun:"res"`
}
