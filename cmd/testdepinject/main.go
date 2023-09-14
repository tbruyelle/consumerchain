package main

import (
	"fmt"

	"cosmossdk.io/depinject"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
)

func GetX() int { return 42 }

func main() {
	/*
		var x int
		err := depinject.Inject(
			depinject.Supply(
				GetX,
			),
			&x,
		)
		fmt.Println(x, err)
	*/
	var (
		pk paramskeeper.Keeper
		ek evidencekeeper.Keeper
	)
	err := depinject.Inject(
		depinject.Provide(
			GetEvidenceKeeper,
			GetParamsKeeper,
		),
		&ek, &pk,
	)
	_ = pk
	fmt.Println("ERR", err)
}

func GetParamsKeeper() paramskeeper.Keeper {
	return paramskeeper.Keeper{}
}

func GetEvidenceKeeper() evidencekeeper.Keeper {
	return evidencekeeper.Keeper{}
}
