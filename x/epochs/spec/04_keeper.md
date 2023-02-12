<!--
order: 4
-->

# Keepers

## Keeper functions

Epochs keeper module provides utility functions to manage epochs.

```go
// Keeper of epoch nodule maintains collections of epochs and hooks.
type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey
	hooks    types.EpochHooks
}
```

```go
// Keeper is the interface for epoch module keeper
type Keeper interface {
  // GetEpochInfo returns epoch info by identifier
  GetEpochInfo(ctx sdk.Context, identifier string) types.EpochInfo
  // SetEpochInfo set epoch info
  SetEpochInfo(ctx sdk.Context, epoch types.EpochInfo) 
  // DeleteEpochInfo delete epoch info
  DeleteEpochInfo(ctx sdk.Context, identifier string)
  // IterateEpochInfo iterate through epochs
  IterateEpochInfo(ctx sdk.Context, fn func(index int64, epochInfo types.EpochInfo) (stop bool))
  // Get all epoch infos
  AllEpochInfos(ctx sdk.Context) []types.EpochInfo
}
```
