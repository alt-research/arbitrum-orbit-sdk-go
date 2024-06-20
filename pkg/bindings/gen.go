package bindings

//go:generate bash -c "abigen --abi ethereum/RollupCreator.abi --pkg ethereum --type RollupCreator --out ethereum/RollupCreator.go"
//go:generate bash -c "abigen --abi arbone/RollupCreator.abi --pkg arbone --type RollupCreator --out arbone/RollupCreator.go"
//go:generate bash -c "abigen --abi arbsepolia/RollupCreator.abi --pkg arbsepolia --type RollupCreator --out arbsepolia/RollupCreator.go"
