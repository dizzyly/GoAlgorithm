package main

import (
	"fmt"
	"math/rand"
)

/*
	设计一个支持在平均时间复杂度O(1)下，执行以下操作的数据结构。
	注意: 允许出现重复元素。
	insert(val)：向集合中插入元素 val。
	remove(val)：当 val 存在时，从集合中移除一个 val。
	getRandom：从现有集合中随机获取一个元素。每个元素被返回的概率应该与其在集合中的数量呈线性相关。

*/
type RandomizedCollection struct {
	idx map[int]map[int]struct{}
	nums []int
}
func main() {
	obj := Constructor()
	param1 := obj.Insert(2)
	param2 := obj.Remove(2)
	param3 := obj.GetRandom()
	fmt.Println(param1, param2, param3)
}
/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		idx: map[int]map[int]struct{}{},
	}
}
/*
	Inserts a value to the collection.
	Returns true if the collection did not already contain the specified element.
*/
func (r *RandomizedCollection) Insert(val int) bool {
	ids, has := r.idx[val]
	if !has {
		ids = map[int]struct{}{}
		r.idx[val] = ids
	}
	ids[len(r.nums)] = struct{}{}
	r.nums = append(r.nums, val)
	return !has
}
/*
	Removes a value from the collection.
	Returns true if the collection contained the specified element.
*/
func (r RandomizedCollection) Remove(val int) bool {
	ids, has := r.idx[val]
	if !has {
		return false
	}
	var i int
	for id := range ids {
		i = id
		break
	}
	n := len(r.nums)
	r.nums[i] = r.nums[n-1]
	delete(ids, i)
	delete(r.idx[r.nums[i]], n-1)
	if i < n-1 {
		r.idx[r.nums[i]][i] = struct{}{}
	}
	if len(ids) == 0 {
		delete(r.idx, val)
	}
	if len(ids) == 0 {
		delete(r.idx, val)
	}
	r.nums = r.nums[:n-1]
	return true
}
/** Get a random element from the collection. */
func (r RandomizedCollection) GetRandom() int {
	return  r.nums[rand.Intn(len(r.nums))]
}

