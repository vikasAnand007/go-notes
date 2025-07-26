package main

import (
	"fmt"
	"slices"
)

// slice - dynamic arrry which do not require length in adwance.
// most used construct
// useful array methods

// cap([]) gives current grnated capasity to the slice. It is dynamic and increase as new elements are added in slice.
// Capasity becomes double in size if more elements are added then current capasity

func main() {
	// uninitialised slice in nil
	var nums1 []int
	fmt.Println("nums1", nums1, nums1 == nil, len(nums1), cap(nums1)) // [] true 0

	// intialised slice
	var nums2 = make([]int, 0, 5) // (arg1, initial length, initial capasity)
	fmt.Println("nums2", nums2, nums2 == nil, len(nums2), cap(nums2))
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 5)
	fmt.Println("nums2", nums2, nums2 == nil, len(nums2), cap(nums2))
	nums2 = append(nums2, 6)
	fmt.Println("nums2", nums2, nums2 == nil, len(nums2), cap(nums2))

	// intialised slice - shorthand
	nums3 := []int{}
	// in this case length and capasity will be 0 and increase as we add elements
	fmt.Println("nums3", nums3, nums3 == nil, len(nums3), cap(nums3))
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 1)
	fmt.Println("nums3", nums3, nums3 == nil, len(nums3), cap(nums3))

	// Change value at any index
	strs1 := make([]string, 2)
	// strs1[3] = "a" // this will through error, because length of slice is 2
	strs1[1] = "a"
	fmt.Println("strs1", strs1, strs1 == nil, len(strs1), cap(strs1))

	// copy a slice
	str2 := []string{}
	str2 = append(str2, "vikas")
	str2 = append(str2, "anand")
	var str3 = make([]string, len(str2))
	copy(str3, str2)
	fmt.Println("str3", str3, str3 == nil, len(str3), cap(str3))

	// slice operator - [startingIndex : elementCount]
	str4 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	fmt.Println("str4", str4[0:3])
	fmt.Println("str4", str4[:3])
	fmt.Println("str4", str4[2:6])
	fmt.Println("str4", str4[2:])
	fmt.Println("str4", str4[:])

	// slice comparision
	nums4 := []int{1, 2, 3}
	nums5 := []int{1, 2, 3}
	nums6 := []int{1, 2, 3, 4}
	fmt.Println(slices.Equal(nums4, nums5), slices.Equal(nums4, nums6))

}
