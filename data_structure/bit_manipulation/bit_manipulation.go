package bitmanipulation

import (
	"fmt"
	"math"
)

func missingNumber(nums []int) int {
	n := len(nums)

	var a int
	for i := 0; i <= n; i++ {
		a ^= i
	}

	var b int
	for _, num := range nums {
		b ^= num
	}

	return a ^ b
}

func singleNumber(nums []int) int {
	bits := make([]int, 32)

	// 统计每个bit位上，所有的数都加一起，总共有多少个1
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			bits[i] += (num >> i) & 1
		}
	}

	var ans int
	for i := 0; i < 32; i++ {
		if bits[i]%3 != 0 {
			ans |= (1 << i)
		}
	}

	return ans
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n == (n&-n)
}

// https://www.bilibili.com/video/BV1ch4y1Q7vd
func reverseBits(num uint32) uint32 {
	n := num
	n = (n&0xaaaaaaaa)>>1 | (n&0x55555555)<<1
	n = (n&0xcccccccc)>>2 | (n&0x33333333)<<2
	n = (n&0xf0f0f0f0)>>4 | (n&0x0f0f0f0f)<<4
	n = (n&0xff00ff00)>>8 | (n&0x00ff00ff)<<8
	n = n>>16 | n<<16

	return n
}

type BitSet struct {
	set []int32 // int array as underlying data structure
}

// NewBitSet ...
// n 表示最多能存放的数字个数
// 能表示的范围：0 ~ n-1
// 0 ~ 31 放在 set[0] 中
// 32 ~ 63 放在 set[1] 中，依次类推
func NewBitSet(n int) *BitSet {
	return &BitSet{
		// a/b 如果想向上取整，可以写成 (a+b-1)/b
		// 前提是 a 和 b 都是非负数
		set: make([]int32, (n+31)/32),
	}
}

func (bs *BitSet) Add(num int) {
	bs.set[num/32] |= 1 << (num % 32)
}

func (bs *BitSet) Remove(num int) {
	bs.set[num/32] &= ^(1 << (num % 32))
}

func (bs *BitSet) Flip(num int) {
	bs.set[num/32] ^= (1 << (num % 32))
}

// Contains ...
func (bs *BitSet) Contains(num int) bool {
	return (bs.set[num/32]>>(num%32))&1 == 1
}

// add add implemented using bit op
func add(a, b int32) int32 {
	ans := a
	// 循环直到计算的相加进位信息为0
	for b != 0 {
		ans = a ^ b      // 无进位相加结果
		b = (a & b) << 1 // a 和 b 相加时的进位信息
		a = ans
	}

	return ans
}

// neg 求相反数
// -x = ~x + 1
func neg(x int32) int32 {
	return add(^x, 1)
}

// minus ...
// a - b = a + (-b)
func minus(a, b int32) int32 {
	return add(a, neg(b))
}

func multiply(a, b int32) int32 {
	var ans int32
	// Go 中没有无符号右移运算符。。
	// 只有一个右移运算符，根据数的类型来决定具体的右移操作
	// 当为无符号类型时，进行的是无符号右移
	// 所以为了对一个有符号的数进行无符号右移时，需要将它转成无符号数。。。
	ub := uint32(b)
	fmt.Printf("b: %v, ub: %v\n", b, ub)
	for ub != 0 {
		if ub&1 == 1 {
			ans = add(ans, a)
		}
		ub >>= 1
		a <<= 1
	}

	return ans
}

// div return the result of a / b
// 必须保证 a 和 b 都不是整数最小值
// 因为计算中涉及取相反数（neg），而有符号数的
// 最小值的相反数是无法计算的（因为没有跟它对应的正数）
// 算法原理：找出 a = b * 2^x + b * 2^y + b * 2^z ...
// 中的 x, y, z
func div(a, b int32) int32 {
	// 若 a 或 b 是负数，先把它们转成正数
	x := a
	y := b
	if a < 0 {
		x = neg(a)
	}
	if b < 0 {
		y = neg(b)
	}

	var ans int32
	// 最高位不用考虑，因为是正数，一定是 0
	for i := int32(30); i >= 0; i = minus(i, 1) {
		// 目的是要判断 x >= y << i
		// 但为了避免左移溢出，改成了如下等价形式
		// 因为： x >= y * 2^i <==> x * 2^-i >= y * 2^i * 2^-i
		// <==> x * 2^-i >= y
		if (x >> i) >= y {
			ans |= 1 << i
			// 那这里为啥又不怕溢出了呢？
			// 因为上面的 if 条件判断保证了
			// x 是大于等于 y<<i 的，而 x 既然没有溢出
			// 那么此时 y<<i 一定也不会溢出
			x = minus(x, y<<i)
		}
	}

	// 确定计算结果的符号
	// TODO: 更简单的写法？ (a < 0) ^ (b < 0) ? neg(ans): ans
	if (a < 0 && b < 0) || (a > 0 && b > 0) {
		return ans
	} else {
		return neg(ans)
	}
}

// divide 兼容整数最小值的除法
func divide(a, b int32) int32 {
	min := int32(math.MinInt32)
	if a == min && b == min {
		return 1
	}

	if a != min && b != min {
		return div(a, b)
	}

	// b is minInt32
	if b == min {
		return 0
	}

	// a is minInt32 and b is -1
	// returns maxInt32, 题目要求
	if b == neg(1) {
		return math.MaxInt32
	}

	// a is minInt32 and b is either minInt32 or -1
	if b > 0 {
		a = add(a, b)
	} else {
		a = minus(a, b)
	}

	ans := div(a, b)

	if b > 0 {
		return minus(ans, 1)
	} else {
		return add(ans, 1)
	}
}
