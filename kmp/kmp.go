package kmp

// 获取匹配字符串出错时回溯位置，当回溯位置为-1时表示需要移动主串
func Get_NextVal(pattern string) []int {
	next:=make([]int, len(pattern))
	i:=1
	next[0]=-1
	j:=0
	len:=len(pattern)-1
	for i<len{
		if j==-1||pattern[i]==pattern[j]{
			i++
			j++
			if pattern[i]!=pattern[j]{
				next[i]=j
			} else {
				next[i]=next[j]
			}
		} else {
			j=next[j]
		}
	}
	return next
}
