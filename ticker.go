package ticker

type TickerContext struct {
	Share interface{}
	Count uint64

	CountStart uint64
	CountEnd   uint64

	cancel bool
}

func (cxt *TickerContext) Cancel() {
	cxt.cancel = true
}

type TickerHandler func(cxt *TickerContext)
type Ticker struct {
	cxt TickerContext
}

func (t *Ticker) SetShare(share interface{}) {
	t.cxt.Share = share
}

func (t *Ticker) GetNextCount() uint64 {
	return t.cxt.Count
}

func (t *Ticker) SetNextCount(count uint64) {
	t.cxt.Count = count
}

// SetCountRange range like [start, end)
func (t *Ticker) SetCountRange(rangenum ...uint64) {
	switch len(rangenum) {
	case 1:
		t.cxt.CountStart = rangenum[0]
	case 2:
		t.cxt.CountStart = rangenum[0]
		t.cxt.CountEnd = rangenum[1]
	case 0:
		t.cxt.CountStart = 0
		t.cxt.CountEnd = ^uint64(0)
	}
	t.cxt.Count = t.cxt.CountStart
}

func (t *Ticker) Tick(do TickerHandler) (err interface{}) {
	defer func() {

		err = recover()

		if !t.cxt.cancel {
			t.cxt.Count++
			if t.cxt.Count >= t.cxt.CountEnd {
				t.cxt.Count = t.cxt.CountStart
			}
		} else {
			t.cxt.cancel = false
		}

	}()
	if do != nil {
		do(&t.cxt)
	}
	return nil
}

func New() *Ticker {
	t := &Ticker{}
	t.cxt.CountEnd = ^uint64(0)
	return t
}
