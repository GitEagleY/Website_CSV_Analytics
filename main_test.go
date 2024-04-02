package main

import (
	"testing"
)

func BenchmarkVariant(b *testing.B) {
	day1File := "day1.csv"
	day2File := "day2.csv"

	day1Visits := readCSV(day1File)
	day2Visits := readCSV(day2File)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findUsersWithNewVisits(day1Visits)
		findUsersWithNewVisits(day2Visits)
	}
}
