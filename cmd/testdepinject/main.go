package main

import (
	"fmt"

	"cosmossdk.io/depinject"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
)

type App struct {
	ParamsKeeper   paramskeeper.Keeper
	EvidenceKeeper evidencekeeper.Keeper
}

func main() {
	var (
		// app App
		pk paramskeeper.Keeper
		ek evidencekeeper.Keeper
	)
	err := depinject.Inject(
		depinject.Supply(
			GetEvidenceKeeper,
			GetParamsKeeper,
		),
		&ek, &pk,
	)
	fmt.Println(pk, ek, err)
}
func GetParamsKeeper() *paramskeeper.Keeper     { return &paramskeeper.Keeper{} }
func GetEvidenceKeeper() *evidencekeeper.Keeper { return &evidencekeeper.Keeper{} }
