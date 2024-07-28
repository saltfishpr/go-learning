package main

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	for _, booking := range bookings {
		first, last, seats := booking[0], booking[1], booking[2]
		diff[first-1] += seats
		if last < n {
			diff[last] -= seats
		}
	}
	for i := 1; i < n; i++ {
		diff[i] = diff[i] + diff[i-1]
	}
	return diff
}
