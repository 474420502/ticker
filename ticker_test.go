package ticker

import (
	"fmt"
	"testing"
)

func TestTicker(t *testing.T) {
	ticker := New()
	ticker.SetCountRange(0, 15)

	var count = 0
	for i := 0; i < 17; i++ {
		ticker.Tick(func(cxt *TickerContext) {
			if cxt.Count >= 15 {
				t.Error("tick error. is not limited. ")
			}
			count++
		})
	}

	if ticker.GetNextCount() != 2 {
		t.Error("ticker count should be 2", ticker.GetNextCount())
	}

	if count != 17 {
		t.Error("my count should be 17", count)
	}

}

func TestTickerCancel(t *testing.T) {
	ticker := New()
	ticker.SetCountRange(0, 15)

	var count int64
	ticker.SetShare(&count)
	for i := 0; i < 17; i++ {
		ticker.Tick(func(cxt *TickerContext) {
			if cxt.Count >= 15 {
				t.Error("tick error. is not limited. ")
			}
			mcount := cxt.Share.(*int64)
			*mcount++

			cxt.Cancel()
		})
	}

	if ticker.GetNextCount() != 0 {
		t.Error("ticker count should be 0", ticker.GetNextCount())
	}

	if count != 17 {
		t.Error("my count should be 17", count)
	}

}

func TestTickerRange(t *testing.T) {
	ticker := New()
	ticker.SetCountRange(5)
	for i := 0; i < 6; i++ {
		ticker.Tick(nil)
		if i == 0 && ticker.GetNextCount() == 6 {
			break
		} else {
			t.Error("Range is error. check it", ticker.GetNextCount())
		}
	}

	ticker.SetCountRange(2, 7)
	ticker.SetNextCount(7)
	ticker.Tick(nil)
	if ticker.GetNextCount() != 2 {
		t.Error("Next Count should be 2. 7 tick() 7 >= limit. so 7 -> 2")
	}
}

func TestTickerPanic(t *testing.T) {
	ticker := New()
	ticker.SetCountRange()
	err := ticker.Tick(func(cxt *TickerContext) {
		panic("error test")
	})

	if err == nil {
		t.Error("should be error")
	}

	if err.(string) != "error test" {
		t.Error("check it")
	}

	err = ticker.Tick(func(cxt *TickerContext) {
		panic(fmt.Errorf("error test"))
	})

	if err.(error).Error() != "error test" {
		t.Error("check it")
	}
}
