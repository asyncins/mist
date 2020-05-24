/*
* 薄雾算法
*
* 1      2                                                     48         56       64
* +------+-----------------------------------------------------+----------+----------+
* retain | increas                                             | salt     | sequence |
* +------+-----------------------------------------------------+----------+----------+
* 0      | 0000000000 0000000000 0000000000 0000000000 0000000 | 00000000 | 00000000 |
* +------+-----------------------------------------------------+------------+--------+
*
* 0. 最高位，占 1 位，保持为 0，使得值永远为正数；
* 1. 自增数，占 47 位，自增数在高位能保证结果值呈递增态势，遂低位可以为所欲为；
* 2. 随机因子，占 8 位，上限数值 255，使结果值不可预测；
* 3. 序列号，占 8 位，上限数值 255，表示同一毫秒生成的第 N 个数；
*
* 编号上限为百万亿级，上限值计算为 140737488355327 即 int64(1 << 47 - 1)，假设每天取值 10 亿，能使用 385+ 年
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const saltBit = uint(8)                       // 随机因子二进制位数
const sequenceBit = uint(8)                   // 序列号二进制位数
const saltMax = int64(1<<saltBit - 1)         // 随机因子数值上限
const sequenceMax = int64(1<<sequenceBit - 1) // 序列号数值上限
const saltShift = sequenceBit                 // 随机因子移位数
const increasShift = saltBit + sequenceBit    // 自增数移位数

/* 生成区间范围内的随机数 */
func RandInt64(min, max, sequence int64) int64 {
	rand.Seed(time.Now().Unix() + sequence) // 时间戳作为随机因子
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

type Mist struct {
	sync.Mutex       // 互斥锁
	increas    int64 // 自增数
	salt       int64 // 随机因子
	sequence   int64 // 序列号
	timestamp  int64 // 时间戳
	saltMax    int64 // 随机数值上限
	saltMin    int64 // 随机数值下限
}

/* 初始化 Mist 结构体*/
func NewMist() *Mist {
	mist := Mist{increas: 1, salt: 0, sequence: 0, saltMax: 255, saltMin: 1}
	return &mist
}

/* 生成唯一编号 */
func (c *Mist) Generate() int64 {
	c.Lock()
	currentTime := time.Now().UnixNano() / 1e6
	// 通过时间戳对比决定序列号的值
	if c.timestamp == currentTime {
		c.sequence = (c.sequence + 1) & sequenceMax
	} else {
		c.sequence = 0
	}
	c.timestamp = currentTime
	c.increas++                                                                     // 自增
	c.salt = RandInt64(c.saltMin, c.saltMax, c.sequence)                            // 获取随机因子数值
	mist := int64((c.increas << increasShift) | (c.salt << saltShift) | c.sequence) // 通过位运算实现自动占位
	c.Unlock()
	return mist
}

func main() {
	// 使用方法
	mist := NewMist()
	for i := 0; i < 10; i++ {
		fmt.Println(mist.Generate())
	}
}
