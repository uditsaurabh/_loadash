package array

import "errors"

func ChunkIntArray(arr []int, chunk_size int) (res [][]int, err error) {
	arr_len := len(arr)
	if arr_len == 0 {
		return nil, errors.New("empty array has been provided")
	} else if chunk_size < 0 {
		return nil, errors.New("chunk size is invalid")
	} else if chunk_size > arr_len {
		return nil, errors.New("array size is bigger than the array length")
	} else {
		index := 0
		res = [][]int{}
		if chunk_size > 0 {
			for (index + chunk_size) < arr_len {
				res = append(res, arr[index:index+chunk_size])
				index = index + chunk_size
			}
		}
		res = append(res, arr[index:])
	}
	return res, nil
}
