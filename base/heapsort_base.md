# 堆排序

## 基础概念

### 完全二叉树
一个数组，从0开始依次填满树结构，可以认为是一个完全二叉树，从下标来看，树的结构是这样的
```
[2,4,5,6,2,4,5]
下标完全二叉树：
            0
        |       |
        1       2
      |   |   |   |
      3   4   5   6
```
则有如下公式：   
对任意一个i位置的下标：   
左孩子=  2 * i + 1   
右孩子=  2 * i + 2   
父节点=  i -1 / 2   

### 大根堆 & 小根堆
 一个完全二叉树里，每个子树的最大值就是父节点的值，就是大根堆，反之就是小根堆

#### 如何构建大根堆
每添加一个数，都和自己的父节点进行比较，如果大于父节点，则和父节点进行位置交换，再和父节点的父节点进行比较，直到不比父..节点大了，再停止

这个过程叫做 ```heapinsert```

golang代码实现
```golang
func heapInsert([]int arr,index int) {
    for {
        if arr[index] > arr[(index-1)/2] {
            swap.Swap(arr,index,(index-1) /2 )
            index = (index-1) /2 
        }else{
            break
        }
    }
}

```

#### 如何拿掉大根堆最大值保证剩下的节点依然是大根堆
返回最大值后，把最后一个下标的元素拿到最大值的位置（根节点），然后依次和子孩子进行比较，找到孩子中的最大值，交换位置，再和子孩子进行比较，周而复始

（条件是，原来就是大根堆）这个过程叫做```heapify```

golang代码实现
```golang
func heapify(arr []int ,index ,heapsize int) {
    left := 2 * index + 1 // 左孩子下标
    for{
        if left >= heapsize{
            break
        }

        // 两个孩子中谁最大
        var largest int
        if left+1 < heapsize && arr[left+1] > arr[left] {
            largest = left +1
        }else {
            largest = left 
        }

        // 父孩子和largest谁大
        if arr[largest] <= arr[index] {
            largest = index
        }

        if largest == index {
            break
        }

        swap.Swap(arr,largest,index)
        index = largest
        left = index *2 +1
    }

}

```
