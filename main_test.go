package main

import (
    "testing"
)

func TestAddition(t *testing.T) {
    testCases := []struct {
        a, b int
        expected int
    }{
        {2, 3, 5},
        {4, 6, 10},
        {5, 7, 12},
        {9, 12, 21},
        {10, 14, 24},
        {15, 16, 31},
    }
    
    for _, tc := range testCases {
        if result := addition(tc.a, tc.b); result != tc.expected {
            t.Errorf("Expected: %v\nResult: %v\n-------\n", tc.expected, result)
        }
    }
    
}

func TestSubtraction(t *testing.T) {
    testCases := []struct {
        a, b int
        expected int
    }{
        {7, 3, 4},
        {18, 6, 12},
        {17, 5, 12},
        {25, 12, 13},
        {10, 14, -4},
        {15, 16, -1},
    }
    for _, tc := range testCases {
        if result := subtraction(tc.a, tc.b); result != tc.expected {
            t.Errorf("Expected: %v\nResult: %v\n-------\n", tc.expected, result)
        }
    }

}

func TestMultiplication(t *testing.T) {
    testCases := []struct {
        a, b float64
        expected float64
    }{
        {2, 3, 6},
        {4, 6, 24},
        {5, 7, 35},
        {9, 12, 108},
        {10, 14, 140},
        {12, 0, 0},
    }
    for _, tc := range testCases {
        if result := multiplication(tc.a, tc.b); result != tc.expected {
            t.Errorf("Expected: %v\nResult: %v\n-------\n", tc.expected, result)
        }
    }

}

func TestDividen(t *testing.T) {
    testCases := []struct {
        a, b float64
        expected float64
    }{
        {4, 2, 2},
        {10, 10, 1},
        {120, 10, 12},
    }
    for _, tc := range testCases {
        if result := dividen(tc.a, tc.b); result != tc.expected {
            t.Errorf("Expected: %v\nResult: %v\n-------\n", tc.expected, result)
        }
    }

}
