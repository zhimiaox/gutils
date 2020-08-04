package gutils

import (
	"testing"
	"time"
)

func TestNewLockTable(t *testing.T) {
	lock := NewLockTable()
	isLock := lock.IsLock("a", 2*time.Second)
	if isLock {
		t.Fatal("首次上锁失败")
	}
	isLock = lock.IsLock("a", 2*time.Second)
	if !isLock {
		t.Fatal("二次上锁穿透")
	}
	time.Sleep(2 * time.Second)
	isLock = lock.IsLock("a", 2*time.Second)
	if isLock {
		t.Fatal("解锁失败")
	}
}
