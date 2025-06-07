package snowflake

import (
	//"gitlab.inheaden.io/IL/go-service-libs/v3/utils"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const (
	// Epoch is the unix epoch of 5th of September 2021.
	Epoch int64 = 1315164522123

	NodeBits    uint8 = 10
	CounterBits uint8 = 12

	NodeMax    = -1 ^ (-1 << NodeBits)
	CounterMax = -1 ^ (-1 << CounterBits)

	NodeShift = CounterBits
	TimeShift = NodeBits + CounterBits
)

type ID int64

type Generator struct {
	//*utils.WithLogger
	mutex sync.Mutex

	nodeID   int64
	counter  int64
	lastTime int64
	epoch    time.Time
}

// NewGenerator creates a new instance of the snowflake generator.
//
// A snowflake is a pseudo random 63bit ID that is
// * Roughly sorted
// * Most certainly unique across instances
//
// Structure:
//
// | Unix time in ms | Node ID | Counter |
// |----- 41 bit ----|--- 10 --|-- 12 ---|
// </pre>
//
// The unix time makes the IDs roughly sorted, the nodeID ensures uniqueness.
//
// The node ID must be unique across instances.
// The generator is capable of generation 4096 unique IDs within one ms.
func NewGenerator(nodeID int64) (*Generator, error) {
	if nodeID < 0 || nodeID > NodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(NodeMax, 10))
	}

	return &Generator{
		nodeID:  nodeID,
		counter: 0,
		epoch:   time.Unix(Epoch/1000, (Epoch%1000)*1000000), //nolint:gomnd // -
	}, nil
}

// Next returns a new snowflake ID.
func (g *Generator) Next() ID {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	now := time.Since(g.epoch).Nanoseconds() / 1000000
	if now == g.lastTime {
		g.counter++
		g.counter &= CounterMax
	} else {
		g.counter = 0
	}

	result := ID((now << TimeShift) | (g.nodeID << NodeShift) | g.counter)
	g.lastTime = now

	return result
}

// Int64 returns an int64 of the snowflake ID.
func (f ID) Int64() int64 {
	return int64(f)
}

func (f ID) String() string {
	return strconv.FormatInt(f.Int64(), 10)
}
