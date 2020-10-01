package Week_08

import (
	"testing"
)

func TestBitManipulationActual(t *testing.T) {
	var x, y int
	// 1. x的最右边n位清零
	x = 30
	t.Logf("1. 判断奇偶 (%v)", x)
	t.Logf(" %8d, %v", x, x&1)
	t.Logf(" %8d, %v", x+1, (x+1)&1)

	t.Logf("2.除以二(%v)", x)
	t.Logf(" %8d, %v", x, x>>1)
	t.Logf(" %8d, %v", x+1, (x+1)>>1)

	t.Logf("3.清零最低位的1(%v)", x)
	t.Logf(" %8d, %08b", x, x)
	y = x & (x - 1)
	t.Logf(" %8d, %08b", y, y)

	t.Logf("4.得到低位的1(%v)", x)
	t.Logf(" %8d, %08b", x, x)
	y = x & (-x)
	t.Logf(" %8d, %08b", y, y)

	t.Logf("5. 0(%v)", x)
	t.Logf(" %8d, %08b", x, x)
	y = x & (^x)
	t.Logf(" %8d, %08b", y, y)
	y = x &^ x
	t.Logf(" %8d, %08b", y, y)
}

func TestBitManipulation(t *testing.T) {
	var x, y, n int
	x, n = 30, 3
	t.Logf("1. x的最右边n位清零 (%v)(%v)", x, n)
	t.Logf(" %8d, %08b", x, x)
	y = x & (^0 << n)
	t.Logf(" %8d, %08b", y, y)

	t.Logf("2. 获取x的第n位值(%v)(%v)", x, n)
	t.Logf(" %8d, %08b", x, x)
	y = (x >> n) & 1
	t.Logf(" %8d, %08b", y, y)

	t.Logf("3. 获取x的第n位幂值(%v)(%v)", x, n)
	t.Logf(" %8d, %08b", x, x)
	y = x & (1 << n)
	t.Logf(" %8d, %08b", y, y)

	t.Logf("4. 仅将第n位置为1(%v)(%v)", x, n)
	t.Logf(" %8d, %08b", x, x)
	y = x | (1 << 5)
	t.Logf(" %8d, %08b", y, y)

	t.Logf("5. 仅将第n位置为0(%v)(%v)", x, n)
	t.Logf(" %8d, %08b", x, x)
	y = x & (^(1 << n))
	t.Logf(" %8d, %08b", y, y)

	t.Logf("6. 将x的最最高位至第n位(含)清零(%v)(%v)", x, n)
	t.Logf(" %8d, %08b", x, x)
	y = x & ((1 << n) - 1)
	t.Logf(" %8d, %08b", y, y)

}
