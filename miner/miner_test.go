package miner

import (
	"time"
)

type fakeTime struct{}

func (f fakeTime) Now() time.Time {
	return time.Date(2018, 10, 9, 12, 48, 30, 0, time.FixedZone("test", 0)) // timestamp: 1539089310
}

// func TestMineBlockReturnsBlock(t *testing.T) {
// 	Now = fakeTime{}.Now

// 	expectedBlock := t.Block{
// 		Timestamp: 1539089310,
// 	}

// 	minedBlock, _ := MineBlock([]t.Transaction{})

// 	assert.Equal(t, expectedBlock, minedBlock)
// }
