package main

import "fmt"

func main75() {
	nums := []int{2, 0, 1}
	fmt.Println("before sort, ", nums)
	sortColors(nums)
	fmt.Println("after sort, ", nums)
}
func sortColors1(nums []int) {
	ind0, ind1, ind2 := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		fmt.Printf("i:%d,ind0 %d ind1 %d ind2 %d,  nums:%v\n", i, ind0, ind1, ind2, nums)
		if nums[i] == 0 {
			//if ind1 & ind2 is smaller than ind0, switch nums[ind0] & nums[ind2] and swith nums[newind0] and nums[ind1]
			ind0 = i
			if ind0 > ind2 && ind2 != -1 {
				nums[ind0], nums[ind2] = nums[ind2], nums[ind0]
				ind0, ind2 = ind2, ind0 //swap
			}
			//then swap ind0 & ind1
			if ind0 > ind1 && ind1 != -1 {
				nums[ind0], nums[ind1] = nums[ind1], nums[ind0]
				ind0, ind1 = ind1, ind0 //swap
			}
		} else if nums[i] == 1 {
			ind1 = i
			//swap ind1 & ind2
			if ind1 > ind2 && ind2 != -1 {
				nums[ind1], nums[ind2] = nums[ind2], nums[ind1]
				ind1, ind2 = ind2, ind1 //swap
			}
		} else if nums[i] == 2 {
			ind2++
			//do nothing
		}
	}
}
func sortColors(nums []int) {
	j, k := 0, len(nums)-1
	for i := 0; i <= k; i++ {
		fmt.Printf("i:%d, j:%d, k:%d, nums:%v\n", i, j, k, nums)
		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			//move 0's pointer to j+1
			j++
		} else if nums[i] == 2 {
			nums[i], nums[k] = nums[k], nums[i]
			//move i backword to check the value of swapped nums[k] again
			i--
			//move 2's pointer to k-1
			k--
		}
	}
}
