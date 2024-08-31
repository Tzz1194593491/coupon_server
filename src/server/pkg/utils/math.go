package utils

// Equipartition 近似均匀分配
// num - 待分配待数
// cnt - 分母
func Equipartition(num, cnt int) *[]int {
	if cnt == 0 {
		return nil
	}
	partition := make([]int, cnt)
	eachPartition := num / cnt
	remainder := 0
	if num >= cnt {
		for i := 0; i < cnt; i++ {
			partition[i] = eachPartition
		}
		remainder = num % cnt
	} else {
		remainder = num
	}
	for i := 0; i < remainder; i++ {
		partition[i]++
	}
	return &partition
}
