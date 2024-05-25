package trimmean

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func generateFloats(n int, min, max float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	data := make([]float64, n)
	for i := range data {
		data[i] = min + rand.Float64()*(max-min)
	}
	return data
}

func TestTrimmedMean(t *testing.T) {
	data := generateFloats(100, -100, 100)
	got := TrimmedMean(data, 0.05, 0.05)
	want := 0.0    //expected mean after the trimming process (this will likely change as the R code is getting random floats and thus producing "random" trimmed mean)
	tolerance := 5 //tolerace is important particularly on tests that are run on datasets that can change or when randomly generated during each test run (like this one)
	if math.Abs(got-want) > float64(tolerance) {
		t.Errorf("TrimmedMean() = %v, want %v", got, want)
	}
}
