/*
 * 纸喵软件
 * Copyright (c) 2017~2020 http://zhimiao.org All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2020/3/3 下午4:26
 */
package gutils

import (
	"sync"
	"time"
)

type lockItem struct {
	Key      string
	LifeSpan time.Duration // 生命周期
	CreateOn time.Time     // 创建时间
}

type lockTable struct {
	sync.RWMutex
	CleanerDuraction time.Duration       // 触发定时清理器的时间
	Cleaner          *time.Timer         // 定时清理器
	Items            map[string]lockItem // 子集
}

func NewLockTable() *lockTable {
	return &lockTable{
		Items: make(map[string]lockItem),
	}
}

func (l *lockTable) IsLock(key string, lock_time time.Duration) bool {
	l.Lock()
	if item, ok := l.Items[key]; ok {
		l.Unlock()
		if time.Now().Sub(item.CreateOn) > item.LifeSpan {
			l.cleanerCheck()
			return false
		}
		return true
	}
	l.Items[key] = lockItem{
		Key:      key,
		LifeSpan: lock_time,
		CreateOn: time.Now(),
	}
	cleannerDuraction := l.CleanerDuraction
	l.Unlock()
	if cleannerDuraction == 0 {
		l.cleanerCheck()
	}
	return false
}

func (l *lockTable) cleanerCheck() {
	l.Lock()
	defer l.Unlock()
	if l.Cleaner != nil {
		l.Cleaner.Stop()
	}
	// 遍历当前限制的key, 遇到过期的将其删掉
	// 其余的则从中找到最近一个将要过期的key并且将它还有多少时间过期作为下一次清理任务的定时时间
	now := time.Now()
	smallestDuracton := 0 * time.Second
	for key, item := range l.Items {
		lifeSpan := item.LifeSpan
		createOn := item.CreateOn
		if now.Sub(createOn) >= lifeSpan {
			delete(l.Items, key)
		} else {
			if smallestDuracton == 0 || lifeSpan-now.Sub(createOn) < smallestDuracton {
				smallestDuracton = lifeSpan - now.Sub(createOn)
			}
		}
	}
	l.CleanerDuraction = smallestDuracton
	// 将最近一个将要过期的key距离现在的时间作为启动清理任务的定时时间
	if l.CleanerDuraction > 0 {
		l.Cleaner = time.AfterFunc(l.CleanerDuraction, func() {
			go l.cleanerCheck()
		})
	}
}
