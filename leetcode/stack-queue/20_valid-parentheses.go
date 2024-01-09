package stack_queue

func isValid(s string) bool {

	mmap := map[byte]byte{')': '(', '}': '{', ']': '['}
	arr := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			arr = append(arr, s[i])
		} else if len(arr) > 0 && mmap[s[i]] == arr[len(arr)-1] {
			arr = arr[:len(arr)-1]
			continue
		} else {
			return false
		}
	}
	return len(arr) == 0
}
