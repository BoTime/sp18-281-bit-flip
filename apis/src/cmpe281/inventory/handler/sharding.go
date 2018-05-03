package handler

import (
	"github.com/gocql/gocql"
	"log"
)

func SelectShard(storeId gocql.UUID, ctx ShardedDatabaseContext) (*gocql.Session) {
	runeUuid := []rune(storeId.String())
	lastRune := runeUuid[len(runeUuid)-1]

	if bitsSet(lastRune) % 2 == 0 {
		log.Println("Selected Shard 1")
		return ctx.Shard1
	} else {
		log.Println("Selected Shard 2")
		return ctx.Shard2
	}
}

func bitsSet(rune rune) rune {
	// Bit hacks to count number of bits set
	// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
	rune = (rune & 0x55) + ((rune >> 1) & 0x55)
	rune = (rune & 0x33) + ((rune >> 2) & 0x33)
	return (rune + (rune >> 4)) & 0xF
}