package model_test

import (
	"testing"

	"github.com/bcokert/terragen/model"
	"github.com/bcokert/terragen/noise"
)

func TestNewNoise(t *testing.T) {
	testCases := map[string]struct {
		NoiseFunction string
		Expected      model.Noise
	}{
		"basic": {
			NoiseFunction: "white:1d",
			Expected: model.Noise{
				NoiseFunction: "white:1d",
			},
		},
	}

	for name, testCase := range testCases {
		result := model.NewNoise(testCase.NoiseFunction)
		if !result.IsEqual(&testCase.Expected) {
			t.Errorf("%s failed. Expected %#v, received %#v", name, testCase.Expected, result)
		}
	}
}

func TestGenerate(t *testing.T) {
	testCases := map[string]struct {
		From          []float64
		To            []float64
		Resolution    int
		NoiseFunction noise.Function
		Expected      model.Noise
	}{
		"1d simple": {
			From:       []float64{-1},
			To:         []float64{3},
			Resolution: 4,
			NoiseFunction: func(t []float64) float64 {
				return t[0] * 2
			},
			Expected: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{-1, -0.75, -0.5, -0.25, 0, 0.25, 0.5, 0.75, 1, 1.25, 1.5, 1.75, 2, 2.25, 2.5, 2.75},
					"value": []float64{-2, -1.5, -1, -0.5, 0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5},
				},
				From:       []float64{-1},
				To:         []float64{3},
				Resolution: 4,
			},
		},
		"1d empty range": {
			From:       []float64{3},
			To:         []float64{3},
			Resolution: 4,
			NoiseFunction: func(t []float64) float64 {
				return t[0] * 2
			},
			Expected: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{},
					"value": []float64{},
				},
				From:       []float64{3},
				To:         []float64{3},
				Resolution: 4,
			},
		},
		"1d high resolution": {
			From:       []float64{0},
			To:         []float64{1},
			Resolution: 25,
			NoiseFunction: func(t []float64) float64 {
				return t[0] * 3
			},
			Expected: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96},
					"value": []float64{0, 0.12, 0.24, 0.36, 0.48, 0.60, 0.72, 0.84, 0.96, 1.08, 1.20, 1.32, 1.44, 1.56, 1.68, 1.80, 1.92, 2.04, 2.16, 2.28, 2.40, 2.52, 2.64, 2.76, 2.88},
				},
				From:       []float64{0},
				To:         []float64{1},
				Resolution: 25,
			},
		},
		"2d simple": {
			From:       []float64{-1, 3},
			To:         []float64{3, 5},
			Resolution: 2,
			NoiseFunction: func(t []float64) float64 {
				return t[0] + 10*t[1]
			},
			Expected: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{-1, -1, -1, -1, -0.5, -0.5, -0.5, -0.5, 0, 0, 0, 0, 0.5, 0.5, 0.5, 0.5, 1, 1, 1, 1, 1.5, 1.5, 1.5, 1.5, 2, 2, 2, 2, 2.5, 2.5, 2.5, 2.5},
					"t2":    []float64{3, 3.5, 4, 4.5, 3, 3.5, 4, 4.5, 3, 3.5, 4, 4.5, 3, 3.5, 4, 4.5, 3, 3.5, 4, 4.5, 3, 3.5, 4, 4.5, 3, 3.5, 4, 4.5, 3, 3.5, 4, 4.5},
					"value": []float64{29, 34, 39, 44, 29.5, 34.5, 39.5, 44.5, 30, 35, 40, 45, 30.5, 35.5, 40.5, 45.5, 31, 36, 41, 46, 31.5, 36.5, 41.5, 46.5, 32, 37, 42, 47, 32.5, 37.5, 42.5, 47.5},
				},
				From:       []float64{-1, 3},
				To:         []float64{3, 5},
				Resolution: 2,
			},
		},
		"2d empty range": {
			From:       []float64{-1, 3},
			To:         []float64{-1, 5},
			Resolution: 2,
			NoiseFunction: func(t []float64) float64 {
				return t[0] + 10*t[1]
			},
			Expected: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{},
					"t2":    []float64{},
					"value": []float64{},
				},
				From:       []float64{-1, 3},
				To:         []float64{-1, 5},
				Resolution: 2,
			},
		},
		"2d high resolution": {
			From:       []float64{0, 0},
			To:         []float64{2, 1},
			Resolution: 25,
			NoiseFunction: func(t []float64) float64 {
				return t[0] + 10*t[1]
			},
			Expected: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.04, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.08, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.12, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.16, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.24, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.28, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.32, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.36, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.4, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.44, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.48, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.52, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.56, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.6, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.64, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.68, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.72, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.76, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.84, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.88, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.92, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 0.96, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.04, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.08, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.12, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.16, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.24, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.28, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.32, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.36, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.4, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.44, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.48, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.52, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.56, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.6, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.64, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.68, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.72, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.76, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.8, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.84, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.88, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.92, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96, 1.96},
					"t2":    []float64{0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96, 0, 0.04, 0.08, 0.12, 0.16, 0.2, 0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48, 0.52, 0.56, 0.6, 0.64, 0.68, 0.72, 0.76, 0.8, 0.84, 0.88, 0.92, 0.96},
					"value": []float64{0, 0.4, 0.8, 1.2, 1.6, 2, 2.4, 2.8, 3.2, 3.6, 4, 4.4, 4.8, 5.2, 5.6, 6, 6.4, 6.8, 7.2, 7.6, 8, 8.4, 8.8, 9.2, 9.6, 0.04, 0.44, 0.84, 1.24, 1.64, 2.04, 2.44, 2.84, 3.24, 3.64, 4.04, 4.44, 4.84, 5.24, 5.64, 6.04, 6.44, 6.84, 7.24, 7.64, 8.04, 8.44, 8.84, 9.24, 9.64, 0.08, 0.48, 0.88, 1.28, 1.68, 2.08, 2.48, 2.88, 3.28, 3.68, 4.08, 4.48, 4.88, 5.28, 5.68, 6.08, 6.48, 6.88, 7.28, 7.68, 8.08, 8.48, 8.88, 9.28, 9.68, 0.12, 0.52, 0.92, 1.32, 1.72, 2.12, 2.52, 2.92, 3.32, 3.72, 4.12, 4.52, 4.92, 5.32, 5.72, 6.12, 6.52, 6.92, 7.32, 7.72, 8.12, 8.52, 8.92, 9.32, 9.72, 0.16, 0.56, 0.96, 1.36, 1.76, 2.16, 2.56, 2.96, 3.36, 3.76, 4.16, 4.56, 4.96, 5.36, 5.76, 6.16, 6.56, 6.96, 7.36, 7.76, 8.16, 8.56, 8.96, 9.36, 9.76, 0.2, 0.6, 1, 1.4, 1.8, 2.2, 2.6, 3, 3.4, 3.8, 4.2, 4.6, 5, 5.4, 5.8, 6.2, 6.6, 7, 7.4, 7.8, 8.2, 8.6, 9, 9.4, 9.8, 0.24, 0.64, 1.04, 1.44, 1.84, 2.24, 2.64, 3.04, 3.44, 3.84, 4.24, 4.64, 5.04, 5.44, 5.84, 6.24, 6.64, 7.04, 7.44, 7.84, 8.24, 8.64, 9.04, 9.44, 9.84, 0.28, 0.68, 1.08, 1.48, 1.88, 2.28, 2.68, 3.08, 3.48, 3.88, 4.28, 4.68, 5.08, 5.48, 5.88, 6.28, 6.68, 7.08, 7.48, 7.88, 8.28, 8.68, 9.08, 9.48, 9.88, 0.32, 0.72, 1.12, 1.52, 1.92, 2.32, 2.72, 3.12, 3.52, 3.92, 4.32, 4.72, 5.12, 5.52, 5.92, 6.32, 6.72, 7.12, 7.52, 7.92, 8.32, 8.72, 9.12, 9.52, 9.92, 0.36, 0.76, 1.16, 1.56, 1.96, 2.36, 2.76, 3.16, 3.56, 3.96, 4.36, 4.76, 5.16, 5.56, 5.96, 6.36, 6.76, 7.16, 7.56, 7.96, 8.36, 8.76, 9.16, 9.56, 9.96, 0.4, 0.8, 1.2, 1.6, 2, 2.4, 2.8, 3.2, 3.6, 4, 4.4, 4.8, 5.2, 5.6, 6, 6.4, 6.8, 7.2, 7.6, 8, 8.4, 8.8, 9.2, 9.6, 10, 0.44, 0.84, 1.24, 1.64, 2.04, 2.44, 2.84, 3.24, 3.64, 4.04, 4.44, 4.84, 5.24, 5.64, 6.04, 6.44, 6.84, 7.24, 7.64, 8.04, 8.44, 8.84, 9.24, 9.64, 10.04, 0.48, 0.88, 1.28, 1.68, 2.08, 2.48, 2.88, 3.28, 3.68, 4.08, 4.48, 4.88, 5.28, 5.68, 6.08, 6.48, 6.88, 7.28, 7.68, 8.08, 8.48, 8.88, 9.28, 9.68, 10.08, 0.52, 0.92, 1.32, 1.72, 2.12, 2.52, 2.92, 3.32, 3.72, 4.12, 4.52, 4.92, 5.32, 5.72, 6.12, 6.52, 6.92, 7.32, 7.72, 8.12, 8.52, 8.92, 9.32, 9.72, 10.12, 0.56, 0.96, 1.36, 1.76, 2.16, 2.56, 2.96, 3.36, 3.76, 4.16, 4.56, 4.96, 5.36, 5.76, 6.16, 6.56, 6.96, 7.36, 7.76, 8.16, 8.56, 8.96, 9.36, 9.76, 10.16, 0.6, 1, 1.4, 1.8, 2.2, 2.6, 3, 3.4, 3.8, 4.2, 4.6, 5, 5.4, 5.8, 6.2, 6.6, 7, 7.4, 7.8, 8.2, 8.6, 9, 9.4, 9.8, 10.2, 0.64, 1.04, 1.44, 1.84, 2.24, 2.64, 3.04, 3.44, 3.84, 4.24, 4.64, 5.04, 5.44, 5.84, 6.24, 6.64, 7.04, 7.44, 7.84, 8.24, 8.64, 9.04, 9.44, 9.84, 10.24, 0.68, 1.08, 1.48, 1.88, 2.28, 2.68, 3.08, 3.48, 3.88, 4.28, 4.68, 5.08, 5.48, 5.88, 6.28, 6.68, 7.08, 7.48, 7.88, 8.28, 8.68, 9.08, 9.48, 9.88, 10.28, 0.72, 1.12, 1.52, 1.92, 2.32, 2.72, 3.12, 3.52, 3.92, 4.32, 4.72, 5.12, 5.52, 5.92, 6.32, 6.72, 7.12, 7.52, 7.92, 8.32, 8.72, 9.12, 9.52, 9.92, 10.32, 0.76, 1.16, 1.56, 1.96, 2.36, 2.76, 3.16, 3.56, 3.96, 4.36, 4.76, 5.16, 5.56, 5.96, 6.36, 6.76, 7.16, 7.56, 7.96, 8.36, 8.76, 9.16, 9.56, 9.96, 10.36, 0.8, 1.2, 1.6, 2, 2.4, 2.8, 3.2, 3.6, 4, 4.4, 4.8, 5.2, 5.6, 6, 6.4, 6.8, 7.2, 7.6, 8, 8.4, 8.8, 9.2, 9.6, 10, 10.4, 0.84, 1.24, 1.64, 2.04, 2.44, 2.84, 3.24, 3.64, 4.04, 4.44, 4.84, 5.24, 5.64, 6.04, 6.44, 6.84, 7.24, 7.64, 8.04, 8.44, 8.84, 9.24, 9.64, 10.04, 10.44, 0.88, 1.28, 1.68, 2.08, 2.48, 2.88, 3.28, 3.68, 4.08, 4.48, 4.88, 5.28, 5.68, 6.08, 6.48, 6.88, 7.28, 7.68, 8.08, 8.48, 8.88, 9.28, 9.68, 10.08, 10.48, 0.92, 1.32, 1.72, 2.12, 2.52, 2.92, 3.32, 3.72, 4.12, 4.52, 4.92, 5.32, 5.72, 6.12, 6.52, 6.92, 7.32, 7.72, 8.12, 8.52, 8.92, 9.32, 9.72, 10.12, 10.52, 0.96, 1.36, 1.76, 2.16, 2.56, 2.96, 3.36, 3.76, 4.16, 4.56, 4.96, 5.36, 5.76, 6.16, 6.56, 6.96, 7.36, 7.76, 8.16, 8.56, 8.96, 9.36, 9.76, 10.16, 10.56, 1, 1.4, 1.8, 2.2, 2.6, 3, 3.4, 3.8, 4.2, 4.6, 5, 5.4, 5.8, 6.2, 6.6, 7, 7.4, 7.8, 8.2, 8.6, 9, 9.4, 9.8, 10.2, 10.6, 1.04, 1.44, 1.84, 2.24, 2.64, 3.04, 3.44, 3.84, 4.24, 4.64, 5.04, 5.44, 5.84, 6.24, 6.64, 7.04, 7.44, 7.84, 8.24, 8.64, 9.04, 9.44, 9.84, 10.24, 10.64, 1.08, 1.48, 1.88, 2.28, 2.68, 3.08, 3.48, 3.88, 4.28, 4.68, 5.08, 5.48, 5.88, 6.28, 6.68, 7.08, 7.48, 7.88, 8.28, 8.68, 9.08, 9.48, 9.88, 10.28, 10.68, 1.12, 1.52, 1.92, 2.32, 2.72, 3.12, 3.52, 3.92, 4.32, 4.72, 5.12, 5.52, 5.92, 6.32, 6.72, 7.12, 7.52, 7.92, 8.32, 8.72, 9.12, 9.52, 9.92, 10.32, 10.72, 1.16, 1.56, 1.96, 2.36, 2.76, 3.16, 3.56, 3.96, 4.36, 4.76, 5.16, 5.56, 5.96, 6.36, 6.76, 7.16, 7.56, 7.96, 8.36, 8.76, 9.16, 9.56, 9.96, 10.36, 10.76, 1.2, 1.6, 2, 2.4, 2.8, 3.2, 3.6, 4, 4.4, 4.8, 5.2, 5.6, 6, 6.4, 6.8, 7.2, 7.6, 8, 8.4, 8.8, 9.2, 9.6, 10, 10.4, 10.8, 1.24, 1.64, 2.04, 2.44, 2.84, 3.24, 3.64, 4.04, 4.44, 4.84, 5.24, 5.64, 6.04, 6.44, 6.84, 7.24, 7.64, 8.04, 8.44, 8.84, 9.24, 9.64, 10.04, 10.44, 10.84, 1.28, 1.68, 2.08, 2.48, 2.88, 3.28, 3.68, 4.08, 4.48, 4.88, 5.28, 5.68, 6.08, 6.48, 6.88, 7.28, 7.68, 8.08, 8.48, 8.88, 9.28, 9.68, 10.08, 10.48, 10.88, 1.32, 1.72, 2.12, 2.52, 2.92, 3.32, 3.72, 4.12, 4.52, 4.92, 5.32, 5.72, 6.12, 6.52, 6.92, 7.32, 7.72, 8.12, 8.52, 8.92, 9.32, 9.72, 10.12, 10.52, 10.92, 1.36, 1.76, 2.16, 2.56, 2.96, 3.36, 3.76, 4.16, 4.56, 4.96, 5.36, 5.76, 6.16, 6.56, 6.96, 7.36, 7.76, 8.16, 8.56, 8.96, 9.36, 9.76, 10.16, 10.56, 10.96, 1.4, 1.8, 2.2, 2.6, 3, 3.4, 3.8, 4.2, 4.6, 5, 5.4, 5.8, 6.2, 6.6, 7, 7.4, 7.8, 8.2, 8.6, 9, 9.4, 9.8, 10.2, 10.6, 11, 1.44, 1.84, 2.24, 2.64, 3.04, 3.44, 3.84, 4.24, 4.64, 5.04, 5.44, 5.84, 6.24, 6.64, 7.04, 7.44, 7.84, 8.24, 8.64, 9.04, 9.44, 9.84, 10.24, 10.64, 11.04, 1.48, 1.88, 2.28, 2.68, 3.08, 3.48, 3.88, 4.28, 4.68, 5.08, 5.48, 5.88, 6.28, 6.68, 7.08, 7.48, 7.88, 8.28, 8.68, 9.08, 9.48, 9.88, 10.28, 10.68, 11.08, 1.52, 1.92, 2.32, 2.72, 3.12, 3.52, 3.92, 4.32, 4.72, 5.12, 5.52, 5.92, 6.32, 6.72, 7.12, 7.52, 7.92, 8.32, 8.72, 9.12, 9.52, 9.92, 10.32, 10.72, 11.12, 1.56, 1.96, 2.36, 2.76, 3.16, 3.56, 3.96, 4.36, 4.76, 5.16, 5.56, 5.96, 6.36, 6.76, 7.16, 7.56, 7.96, 8.36, 8.76, 9.16, 9.56, 9.96, 10.36, 10.76, 11.16, 1.6, 2, 2.4, 2.8, 3.2, 3.6, 4, 4.4, 4.8, 5.2, 5.6, 6, 6.4, 6.8, 7.2, 7.6, 8, 8.4, 8.8, 9.2, 9.6, 10, 10.4, 10.8, 11.2, 1.64, 2.04, 2.44, 2.84, 3.24, 3.64, 4.04, 4.44, 4.84, 5.24, 5.64, 6.04, 6.44, 6.84, 7.24, 7.64, 8.04, 8.44, 8.84, 9.24, 9.64, 10.04, 10.44, 10.84, 11.24, 1.68, 2.08, 2.48, 2.88, 3.28, 3.68, 4.08, 4.48, 4.88, 5.28, 5.68, 6.08, 6.48, 6.88, 7.28, 7.68, 8.08, 8.48, 8.88, 9.28, 9.68, 10.08, 10.48, 10.88, 11.28, 1.72, 2.12, 2.52, 2.92, 3.32, 3.72, 4.12, 4.52, 4.92, 5.32, 5.72, 6.12, 6.52, 6.92, 7.32, 7.72, 8.12, 8.52, 8.92, 9.32, 9.72, 10.12, 10.52, 10.92, 11.32, 1.76, 2.16, 2.56, 2.96, 3.36, 3.76, 4.16, 4.56, 4.96, 5.36, 5.76, 6.16, 6.56, 6.96, 7.36, 7.76, 8.16, 8.56, 8.96, 9.36, 9.76, 10.16, 10.56, 10.96, 11.36, 1.8, 2.2, 2.6, 3, 3.4, 3.8, 4.2, 4.6, 5, 5.4, 5.8, 6.2, 6.6, 7, 7.4, 7.8, 8.2, 8.6, 9, 9.4, 9.8, 10.2, 10.6, 11, 11.4, 1.84, 2.24, 2.64, 3.04, 3.44, 3.84, 4.24, 4.64, 5.04, 5.44, 5.84, 6.24, 6.64, 7.04, 7.44, 7.84, 8.24, 8.64, 9.04, 9.44, 9.84, 10.24, 10.64, 11.04, 11.44, 1.88, 2.28, 2.68, 3.08, 3.48, 3.88, 4.28, 4.68, 5.08, 5.48, 5.88, 6.28, 6.68, 7.08, 7.48, 7.88, 8.28, 8.68, 9.08, 9.48, 9.88, 10.28, 10.68, 11.08, 11.48, 1.92, 2.32, 2.72, 3.12, 3.52, 3.92, 4.32, 4.72, 5.12, 5.52, 5.92, 6.32, 6.72, 7.12, 7.52, 7.92, 8.32, 8.72, 9.12, 9.52, 9.92, 10.32, 10.72, 11.12, 11.52, 1.96, 2.36, 2.76, 3.16, 3.56, 3.96, 4.36, 4.76, 5.16, 5.56, 5.96, 6.36, 6.76, 7.16, 7.56, 7.96, 8.36, 8.76, 9.16, 9.56, 9.96, 10.36, 10.76, 11.16, 11.56},
				},
				From:       []float64{0, 0},
				To:         []float64{2, 1},
				Resolution: 25,
			},
		},
	}

	for name, testCase := range testCases {
		noise := &model.Noise{}
		noise.Generate(testCase.From, testCase.To, testCase.Resolution, testCase.NoiseFunction)

		if !noise.IsEqual(&testCase.Expected) {
			t.Errorf("%s failed. Expected %#v, received %#v", name, testCase.Expected, noise)
		}
	}
}

func TestIsEqual(t *testing.T) {
	testCases := map[string]struct {
		Left     model.Noise
		Right    model.Noise
		Expected bool
	}{
		"equal empty": {
			Left:     model.Noise{},
			Right:    model.Noise{},
			Expected: true,
		},
		"equal full": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.10491637639740707, 0.02522960371433485, 0.14240998620874543, -0.06951808844985238, 0.22028277620639095, -0.0036129201768815805, 0.16743035692290176, -0.08762397550026342, 0.25150755724035145, -0.1495588243161205, 0.21620418268684216, -0.11448528972805566, 0.18006882300120994, -0.1482459144529171, 0.12864632754084612, -0.15321647422470572, 0.040347426422610716, -0.14143833187807403, 0.16947245271105693, -0.09200993524802667},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.10491637639740707, 0.02522960371433485, 0.14240998620874543, -0.06951808844985238, 0.22028277620639095, -0.0036129201768815805, 0.16743035692290176, -0.08762397550026342, 0.25150755724035145, -0.1495588243161205, 0.21620418268684216, -0.11448528972805566, 0.18006882300120994, -0.1482459144529171, 0.12864632754084612, -0.15321647422470572, 0.040347426422610716, -0.14143833187807403, 0.16947245271105693, -0.09200993524802667},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			Expected: true,
		},
		"almost equal full": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.10491637639740707, 0.02522960371433485, 0.14240998620874543, -0.06951808844985238, 0.22028277620639095, -0.0036129201768815805, 0.16743035692590176, -0.08762397550026342, 0.25150755724035145, -0.1495588243161205, 0.21620418268684216, -0.11448528972805566, 0.18006882300120994, -0.1482459144529171, 0.12864632754084612, -0.15321647422470572, 0.040347426422610716, -0.14143833187807403, 0.16947245271105693, -0.09200993524802667},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.10491637639740707, 0.02522960371433485, 0.14240998620874543, -0.06951808844985238, 0.22028277620639095, -0.0036129201768815805, 0.16743035692290176, -0.08762397550026342, 0.25150755724035145, -0.1495588243161205, 0.21620418268684216, -0.11448528972805566, 0.18006882300120994, -0.1482459144529171, 0.12864632754084612, -0.15321647422470572, 0.040347426422610716, -0.14143833187807403, 0.16947245271105693, -0.09200993524802667},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			Expected: false,
		},
		"different from": {
			Left: model.Noise{
				From: []float64{1, 2, 3},
			},
			Right: model.Noise{
				From: []float64{1, 3, 3},
			},
			Expected: false,
		},
		"different to": {
			Left: model.Noise{
				To: []float64{1, 2, 3},
			},
			Right: model.Noise{
				To: []float64{1, 3, 3},
			},
			Expected: false,
		},
		"different dimensions": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"t2":    []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Expected: false,
		},
		"different number of dimensions": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{1, 2},
					"t2":    []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Expected: false,
		},
		"different length of dimensions": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{1, 2, 3},
					"value": []float64{1, 2, 3},
				},
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"t1":    []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Expected: false,
		},
	}

	for name, testCase := range testCases {
		if testCase.Left.IsEqual(&testCase.Right) != testCase.Expected {
			t.Errorf("%s failed. Expected %#v == %#v to be %v", name, testCase.Left, testCase.Right, testCase.Expected)
		}
	}
}
