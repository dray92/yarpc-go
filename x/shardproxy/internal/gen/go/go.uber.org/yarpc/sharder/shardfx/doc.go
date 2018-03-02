// Code generated by thriftrw-plugin-yarpc
// @generated

// Package shardfx provides better integration for Fx for services
// implementing or calling Shard.
//
// Clients
//
// If you are making requests to Shard, use the Client function to inject a
// Shard client into your container.
//
// 	fx.Provide(shardfx.Client("..."))
//
// Servers
//
// If you are implementing Shard, provide a shardserver.Interface into
// the container and use the Server function.
//
// Given,
//
// 	func NewShardHandler() shardserver.Interface
//
// You can do the following to have the procedures of Shard made available
// to an Fx application.
//
// 	fx.Provide(
// 		NewShardHandler,
// 		shardfx.Server(),
// 	)
package shardfx