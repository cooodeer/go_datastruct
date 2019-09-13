// 思路：待排序列看成一个有序序列和一个序列。开始时有序序列只有一个元素。
// 过程是每次从无序序列中去出第一个元素，将它依次与有序序列中元素进行比较，得到它应该在的位置。
// 时间复杂度：O(N^2)

//从小到大排列
func InsertSort(arr []int) {
    length := len(arr)
    if length <= 1 {
        return
    }

    for i := 1; i < length; i++ {
        insertVal := arr[i] // 指无序序列的第一个。先假设为二个数
        insertIndex := i - 1 // 指待插入的位置。先假设为无需中的末尾 
        
        // 找位置
        // 每次取到的数都跟左侧的数做比较，如果左侧的数比取的数大，就将左侧的数右移一位，直至左侧没有数字比取的数大为止
        for insertIndex >= 0 && insertval < arr[insertIndex] {
            arr[insertIndex+1] = arr[insertIndex] // 数据后移一位，知道找到待插入的位置
            insertIndex--
        }

        // 放入找到的位置
        // 将取到的数插入到不小于左侧数的位置
        if insertIndex+1 != i { // 如果找到的位置还是原来的位置，就不赋值了
            arr[insertIndex+1] = insertval
        }
    }
}
