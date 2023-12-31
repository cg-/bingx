package bingx

import (
	"github.com/cg-/bingx/perpetualv2"
	"github.com/cg-/bingx/spot"
	"github.com/cg-/bingx/standard"
)

const BASE_URL string = "https://open-api.bingx.com"
const DEMO_BASE_URL string = "https://open-api-vst.bingx.com"

type BingX struct {
	Perpetual *perpetualv2.PerpetualV2
	Spot      *spot.Spot
	Standard  *standard.Standard
	IsDemo    bool
}

func Create(isDemo bool) *BingX {
	return &BingX{
		Perpetual: nil,
		Spot:      nil,
		Standard:  nil,
		IsDemo:    isDemo,
	}
}

func (obj *BingX) AddPerpetualKey(api, secret string) {
	obj.Perpetual = perpetualv2.Create(api, secret)
}

func (obj *BingX) AddSpotKey(api, secret string) {
	obj.Spot = spot.Create(api, secret)
}

func (obj *BingX) AddStandardKey(api, secret string) {
	obj.Standard = standard.Create(api, secret)
}
