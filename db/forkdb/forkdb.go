package forkdb

import (
	"errors"
	"fmt"

	"github.com/coschain/contentos-go/common"
)

const maxSize = 1024

// DB ...
type DB struct {
	//committed common.BlockID
	start  uint64
	offset uint64
	head   common.BlockID

	list     [][]common.BlockID
	branches map[common.BlockID]common.SignedBlock

	// previous BlockID ===> SignedBlock
	detached map[common.BlockID]common.SignedBlock
}

// NewDB ...
func NewDB() *DB {
	return &DB{
		list:     make([][]common.BlockID, maxSize*2+1),
		branches: make(map[common.BlockID]common.SignedBlock),
		detached: make(map[common.BlockID]common.SignedBlock),
	}
}

// Remove removes a block
func (db *DB) Remove(id common.BlockID) {
	num := id.BlockNum()
	if num > db.head.BlockNum()+1 || num < db.start {
		return
	}
	delete(db.branches, id)
	delete(db.detached, id)
	idx := num - db.start + db.offset
	for i := range db.list[idx] {
		if db.list[idx][i] == id {
			l := len(db.list[idx])
			db.list[idx][i], db.list[idx][l-1] = db.list[idx][l-1], db.list[idx][i]
			db.list[idx] = db.list[idx][:l-1]
		}
	}
}

// FetchBlock fetches a block corresponding to id
func (db *DB) FetchBlock(id common.BlockID) (common.SignedBlock, error) {
	b, ok := db.branches[id]
	if ok {
		return b, nil
	}
	return nil, fmt.Errorf("[ForkDB] No block has id of %v", id)
}

// FetchBlockByNum fetches a block corresponding to the block num
func (db *DB) FetchBlockByNum(num uint64) []common.SignedBlock {
	list := db.list[num-db.start+db.offset]
	ret := make([]common.SignedBlock, len(list))
	for i := range list {
		b, _ := db.branches[list[i]]
		ret[i] = b
	}
	return ret
}

// PushBlock adds a block. If any of the forkchain has more than
// maxSize blocks, purge will be triggered.
func (db *DB) PushBlock(b common.SignedBlock) common.SignedBlock {
	id := b.Id()
	num := id.BlockNum()
	if len(db.branches) == 0 {
		db.head = id
		db.start = num
		db.list[0] = append(db.list[0], db.head)
		return b
	}

	if _, ok := db.branches[id]; ok {
		return db.branches[db.head]
	}

	if num > db.head.BlockNum()+1 || num < db.start {
		return db.branches[db.head]
	}
	db.list[num-db.start+db.offset] = append(db.list[num-db.start+db.offset], id)
	db.branches[id] = b
	prev := b.Previous()
	if _, ok := db.branches[prev]; !ok {
		db.detached[prev] = b
	} else {
		db.pushDetached(id)
	}
	db.tryNewHead(id)
	return db.branches[db.head]
}

func (db *DB) pushDetached(id common.BlockID) {
	ok := true
	var b common.SignedBlock
	for ok {
		b, ok = db.detached[id]
		if ok {
			delete(db.detached, id)
			id = b.Id()
			db.tryNewHead(id)
		}
	}
}

func (db *DB) tryNewHead(id common.BlockID) {
	curNum := id.BlockNum()
	if curNum > db.head.BlockNum() {
		db.head = id
		if curNum-db.start >= maxSize {
			db.start++
			db.offset++
		}
		if db.offset >= maxSize {
			db.purge()
		}
	}
}

func (db *DB) purge() {
	var cnt uint64
	for cnt = 0; cnt < db.offset; cnt++ {
		for i := range db.list[cnt] {
			delete(db.branches, db.list[cnt][i])
			delete(db.detached, db.list[cnt][i])
		}
	}

	newList := make([][]common.BlockID, maxSize*2+1)
	copy(newList, db.list[db.offset:])
	db.list = newList
}

// Head returns the head block of the longest chain, returns nil
// if the db is empty
func (db *DB) Head() common.SignedBlock {
	if len(db.branches) == 0 {
		return nil
	}

	return db.branches[db.head]
}

// Pop pops the head block
// NOTE: The only senario Pop should be called is when a fork switch
// occurs, hence the main branch and the fork branch has a common ancestor
// that should NEVER be poped, which also means the main branch cannot be
// poped empty
func (db *DB) Pop() common.SignedBlock {
	ret := db.branches[db.head]
	db.head = ret.Previous()
	if _, ok := db.branches[db.head]; !ok {
		panic("[ForkDB] The main branch was poped empty")
	}
	return ret
}

// FetchBranch finds the nearest ancestor of id1 and id2, then returns
// the 2 branches
func (db *DB) FetchBranch(id1, id2 common.BlockID) ([2][]common.BlockID, error) {
	num1 := id1.BlockNum()
	num2 := id2.BlockNum()
	tid1 := id1
	tid2 := id2
	var ret [2][]common.BlockID
	for num1 > num2 {
		ret[0] = append(ret[0], tid1)
		if b, err := db.getPrevID(tid1); err == nil {
			tid1 = b
			num1 = tid1.BlockNum()
		}
	}
	for num1 < num2 {
		ret[1] = append(ret[1], tid2)
		if b, err := db.getPrevID(tid2); err == nil {
			tid2 = b
			num2 = tid2.BlockNum()
		}
	}

	headNum := db.head.BlockNum()
	for tid1 != tid2 && tid1.BlockNum() <= headNum-maxSize {
		ret[0] = append(ret[0], tid1)
		ret[1] = append(ret[1], tid2)
		tmp, err := db.FetchBlock(tid1)
		if err != nil {
			return ret, err
		}
		tid1 = tmp.Previous()
		tmp, err = db.FetchBlock(tid2)
		if err != nil {
			return ret, err
		}
		tid2 = tmp.Previous()
	}
	if tid1 == tid2 {
		ret[0] = append(ret[0], tid1)
		ret[1] = append(ret[1], tid2)
	} else {
		// This can happen when multiple fork exist and grows simultaneously. To avoid
		// this, call Commit regularly
		errStr := fmt.Sprintf("[ForkDB] cannot find ancestor of %v and %v, unable to switch fork", id1, id2)
		panic(errStr)
	}

	return ret, nil
}

func (db *DB) getPrevID(id common.BlockID) (common.BlockID, error) {
	b, ok := db.branches[id]
	if !ok {
		return common.BlockID{}, fmt.Errorf("[ForkDB] absent key: %v", id)
	}
	return b.Previous(), nil

}

// FetchBlockFromMainBranch returns the num'th block on main branch
func (db *DB) FetchBlockFromMainBranch(num uint64) (common.SignedBlock, error) {
	headNum := db.head.BlockNum()
	if num > headNum || num < db.start {
		return nil, errors.New("[ForkDB] num out of scope")
	}

	var ret common.SignedBlock
	var err error
	cur := db.head
	for headNum >= num {
		ret, err = db.FetchBlock(cur)
		if err != nil {
			return nil, err
		}
		cur = ret.Previous()
		headNum = cur.BlockNum()
	}
	return ret, nil
}

func (db *DB) fetchBlocksSince(id common.BlockID) ([]common.SignedBlock, []common.BlockID, error) {
	length := db.head.BlockNum() - id.BlockNum() + 1
	list := make([]common.SignedBlock, length)
	list1 := make([]common.BlockID, length)
	cur := db.head
	for idx := length - 1; idx >= 0; idx-- {
		b, err := db.FetchBlock(cur)
		if err != nil {
			return nil, nil, err
		}
		list[idx] = b
		list1[idx] = cur
		cur = b.Previous()
	}
	if list1[0] != id {
		errStr := fmt.Sprintf("block %v is not on main branch", id)
		panic(errStr)
	}
	return list, list1, nil
}

// Commit sets the block pointed by id as irreversible. It peals off all
// other branches, sets id as the start block. It should be regularly
// called when a block is commited to save ram.
func (db *DB) Commit(id common.BlockID) {
	_, ids, err := db.fetchBlocksSince(id)
	if err != nil {
		return
	}

	for i := 0; i < len(db.list); i++ {
		for j := 0; j < len(db.list[i]); j++ {
			keep := false
			for k := range ids {
				if db.list[i][j] == ids[k] {
					keep = true
					break
				}
			}
			if !keep {
				delete(db.branches, db.list[i][j])
				delete(db.branches, db.list[i][j])
			}
		}
	}

	db.list = make([][]common.BlockID, maxSize*2+1)
	for i := 0; i < len(ids); i++ {
		db.list[i] = append(db.list[i], ids[i])
	}
	db.start = ids[0].BlockNum()
	db.offset = 0
}