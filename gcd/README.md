### 欧几里得算法

#### 算法证明
```shell
a > b
gcd(a, b) = gcd(b, a%b)

//假设d为a,b最大公约数
//假设r为a%b的余数
//k为整数
d = gcd(a, b)
a = kb + r
r = a - kb
r/d = a/d - bk/d
因为 a/d 和 bk/d都为整数
所以r/d可除尽 所以d|r d整除于r d也是r的约数
d = gcd(a,b,r) 
r < b
d = gcd(b, r)
gcd(b, a%b)

```
