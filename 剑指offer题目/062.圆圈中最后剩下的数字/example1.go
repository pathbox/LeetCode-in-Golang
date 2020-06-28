package Offer062

func lastRemaining(n int, m int) int {
	var f int
	for i := 2; i <= n; i++ {
		f = (f + m) % i
	}
	return f
}

/*
本题的难点在于递推公式的推导或者说是理解。

f(n,m) 表示序号为 0,1,...,n-1 的圈要一直淘汰第 m 个人最终剩下来的序号，这里序号和对应的值是一致的。
f(n-1,m) 表示序号为 0,1,...,n-2 的圈要一直淘汰第 m 个人最终剩下来的序号，这里序号和对应的值是一致的。
当我们从 f(n,m) 中第一次淘汰第 m 个人（序号为 (m-1)%n ）时，该序列的长度就变成了 n-1，再淘汰一个第 m 个人时，情况就变成了 f'(n-1,m), 但是此时 f'(n-1,m) 是以 m%n 为序号开始的，而 f(n-1,m) 是以 0 为序号开始的。要想将 f(n-1,m) 对应的索引转换成 f(n,m) 对应的索引/值，则 f(n,m) = (m%n +f(n-1,m) ) % n = (m+f(n-1,m)) % n

func lastRemaining(n int, m int) int {
  if n == 1{return 0}
  return (m + lastRemaining(n-1,m)) % n
}
*/
