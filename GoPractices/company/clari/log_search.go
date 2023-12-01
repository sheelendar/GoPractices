package main

import "fmt"

/*
  - Complete the 'exec_query' function below.
    *
  - The function is expected to return an INTEGER_ARRAY.
  - The function accepts following parameters:
  - 1. STRING_ARRAY log_types
  - 2. STRING_ARRAY log_messages
  - 3. INTEGER_ARRAY timestamps
  - 4. STRING_ARRAY query_log_types
  - 5. INTEGER_ARRAY query_start_timestamps
  - 6. INTEGER_ARRAY query_end_timestamps

STRING_ARRAY log_types = ["DEBUG", "DEBUG", "INFO", "DEBUG", "DEBUG", "ERROR", "ERROR", "ERROR"];
STRING_ARRAY log_messages = ["message A", "message B", "message BD", "message A", "message AB", "message ABC", "message ABCD", "message ABCD"];
INTEGER_ARRAY timestamps = [5, 8, 12, 15, 18, 20, 22, 25];

STRING_ARRAY query_log_types = ["DEBUG", "INFO", "ERROR"];
INTEGER_ARRAY query_start_timestamps = [8, 5, 8];
INTEGER_ARRAY query_end_timestamps = [15, 15, 15];

Expected output: [2, 1, 0]
Explanation:
There are 2 logs of type "DEBUG" between timestamps 8 and 15, 1 log of type "INFO" between timestamps 5 and 15, and 0 logs of type "ERROR" between timestamps 8 and 15.
*/
func main() {
	log_types := []string{"DEBUG", "DEBUG", "INFO", "DEBUG", "DEBUG", "ERROR", "ERROR", "ERROR"}
	log_messages := []string{"message A", "message B", "message BD", "message A", "message AB", "message ABC", "message ABCD", "message ABCD"}
	timestamps := []int32{5, 8, 12, 15, 18, 20, 22, 25}

	query_log_types := []string{"DEBUG", "INFO", "ERROR"}
	query_start_timestamps := []int32{5, 5, 8}
	query_end_timestamps := []int32{15, 15, 15}
	fmt.Println(exec_query(log_types, log_messages, timestamps, query_log_types, query_start_timestamps, query_end_timestamps))
}

func exec_query(log_types []string, log_messages []string, timestamps []int32, query_log_types []string, query_start_timestamps []int32, query_end_timestamps []int32) []int32 {
	dp := make(map[string][]int32)
	n := len(timestamps)
	for i := 0; i < n; i++ {
		dp[log_types[i]] = append(dp[log_types[i]], timestamps[i])
	}
	var res []int32

	for i := 0; i < len(query_start_timestamps); i++ {
		r := totalCount(query_start_timestamps[i], query_end_timestamps[i], dp[query_log_types[i]])
		res = append(res, r)
	}
	return res
}
func totalCount(start, end int32, arr []int32) int32 {
	count := int32(0)

	for i := 0; i < len(arr); i++ {
		if arr[i] >= start && arr[i] <= end {
			count++
		} else if arr[i] > end {
			break
		}
	}
	return count
}

/*
func exec_query1(log_types []string, log_messages []string, timestamps []int32, query_log_types []string, query_start_timestamps []int32, query_end_timestamps []int32) []int32 {
	dp := make(map[int32]int)
	n := len(timestamps)
	for i := 0; i < n; i++ {
		dp[timestamps[i]] = i
	}

	var res []int32
	for i := 0; i < len(query_start_timestamps); i++ {
		index := fetchIndex(dp, query_start_timestamps[i], query_end_timestamps[i])
		if index == -1 {
			res = append(res, 0)
		} else {
			k := getTotal(index, query_end_timestamps[i], timestamps, log_types, query_log_types[i], n)
			res = append(res, k)
		}
	}
	return res
}

func fetchIndex(dp map[int32]int, start, end int32) int {
	for i := start; i <= end; i++ {
		if v, ok := dp[i]; ok {
			return v
		}
	}
	return -1
}
func getTotal(index int, endTime int32, timestamps []int32, log_types []string, logType string, size int) int32 {
	count := int32(0)
	fmt.Println(index, endTime, logType)
	for i := index; i < size; i++ {
		if logType == log_types[i] {
			count++
		}
		if timestamps[i] >= endTime {
			break
		}
	}
	return count
}
*/
