# master公式

T(N) = a * T(N/b) + O(N^d)

1) log(b,a) > d -> 复杂度为O(N^log(b,a))
2) log(b,a) = d -> 复杂度为O(N^d * logN)
3) log(b,a) < d -> 复杂度为O(N^d)


## 什么是master公式

估计一系列特殊的递归行为的复杂度

必须满足特殊递归
- 母问题的数据量是N级别的
- 子过程的规模必须是等量的
- 除了子过程的算法用O(N^d)表示

## 例题
[求解 L...R上的最大值](binarysearch/getMax_test.go)

```golang
func getMax(arr []int) {
    return process(arr,0,len(arr)-1)
}

func process(arr []int,L,R int) int {
    if L == R {
        return arr[R]
    }
    M := L + ((R-L) >> 1)

    leftMax := process (arr,L,M)
    rightMax := process(arr,M+1,R)
    return math.Max(leftMax,rightMax)
}
```

这个过程子问题规模调用的是等量的，调用了2次，规模是N/2,其余的两次求中点和计算math.Max 是常数级别复杂度O(1)

代入公式求解复杂度
```
T(N) = a * T(N/b) + O(N^ d)
T(N) = 2 * T(N/2) + O(N^ 0)
即 a =2 ,b =2 ,d=0

代入公式： 
log(2,2) = 1 > 0  -> O(N^log(2,2)) ==> O(N)
```
所以这个算法的复杂度是O(N)