// left指待排数组左侧的下标，right指右侧下标
func QuickSort(left int, right int, array *[9]int) {
	l := left
	r := right
	// pivo指中轴
	pivot := array[(left + right)/2]
	// for进行动作有，将左侧序列比中轴值大(从左到右)的和右侧序列比中轴值小(从右到左)的进行交换
	// for呈现的结果是，中轴左侧均比中轴值小，右侧均比中轴值大
	//找遍之后退出循环
	for l<r {
		// 找到每一次需要交换到右边的值的位置
		for array[l]<pivot {
			l++
		}
		// 找到每一次需要交换到左边的值的位置
		for array[r]>pivot {
			r--
		}
		// 表示本次分解任务完成，break
		if l>=r {
			break
		}
		// 交换
		array[l],array[r] = array[r],array[l]
		// 优化
		// if array[l]==pivot {
		// 	r--
		// }
		// if array[r]==pivot {
		// 	r++
		// }
	}
	// 
	if l==r {
		l++
		r--
	}
	if left < r {
		QuickSort(left,r,array)
	}
	if right>l {
		QuickSort(l,right,array)
	}
}

// 测试
//func main() {
//	arr := [9]int{-9,78,0,23,-567,70,123,90,-23}
//	QuickSort(0,len(arr)-1,&arr)
//	fmt.Println(arr)
//}
