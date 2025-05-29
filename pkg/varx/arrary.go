package varx

import (
	"fmt"
	"strings"

	json "github.com/json-iterator/go"
	"github.com/spf13/cast"
)

func TrimArray[T comparable](arr []T, filter func(v T) (bool, T)) []T {
	var newArr []T
	var zeroValue T

	for _, v := range arr {
		if filter != nil {
			if ok, transformed := filter(v); ok && transformed != zeroValue {
				newArr = append(newArr, transformed)
			}
		} else if v != zeroValue {
			newArr = append(newArr, v)
		}
	}

	return newArr
}

// ArrDiff returns the difference between two slices.
func ArrDiff[T comparable](need, existed []T) []T {
	existedMap := make(map[T]struct{}, len(existed))
	for _, e := range existed {
		existedMap[e] = struct{}{}
	}

	var diff []T
	for _, n := range need {
		if _, found := existedMap[n]; !found {
			diff = append(diff, n)
		}
	}

	return diff
}

// ArrSome returns true if any element in the slice satisfies the predicate.
func ArrSome[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if predicate(v) {
			return true
		}
	}
	return false
}

// ArrFind returns the first element in the slice that satisfies the predicate.
func ArrFind[T any](arr []T, predicate func(T) bool) T {
	var zero T
	for _, v := range arr {
		if predicate(v) {
			return v
		}
	}
	return zero
}

func RemoveDuplicate[T comparable](arr []T) []T {
	var newRtn []T
	for _, t := range arr {
		if !ContainEqual(newRtn, t) {
			newRtn = append(newRtn, t)
		}
	}

	return newRtn
}

// ContainEqual check if the aim is in the arr
func ContainEqual[T comparable](arr []T, aim T) bool {
	for _, t := range arr {
		if t == aim {
			return true
		}
	}

	return false
}

// ContainLike check if the aim contains any element in the arr
func ContainLike(arr []string, aim string) bool {
	for _, t := range arr {
		if strings.Contains(aim, t) {
			return true
		}
	}

	return false
}

func ContainPrefix(arr []string, longStr string) bool {
	for _, t := range arr {
		if strings.HasPrefix(longStr, t) {
			return true
		}
	}

	return false
}

func Remove[T comparable](arr []T, aim T) ([]T, bool) {
	for i, t := range arr {
		if t == aim {
			arr = append(arr[:i], arr[i+1:]...)
			return arr, true
		}
	}

	return arr, false
}

func RemoveFromMap[T any](mp map[string]T, keys []string) map[string]T {
	for _, key := range keys {
		delete(mp, key)
	}

	return mp
}
func Arr2InfArr[T any](arr []T) []any {
	var newRtn []any
	for _, t := range arr {
		newRtn = append(newRtn, t)
	}

	return newRtn
}

func Arr2StrArr[T any](arr []T) []string {
	var newRtn []string
	for _, t := range arr {
		newRtn = append(newRtn, cast.ToString(t))
	}

	return newRtn
}

func Arr2IntArr[T any](arr []T) []int {
	var newRtn []int
	for _, t := range arr {
		newRtn = append(newRtn, cast.ToInt(t))
	}
	return newRtn
}

func ArrContainsIn(subStrArr []string, aim string) bool {
	for _, s := range subStrArr {
		if strings.Contains(aim, s) {
			return true
		}
	}

	return false
}

func ToWindowConfig(m map[string]any) string {
	marshalString, err := json.MarshalToString(m)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(`window.APP_CONFIG = %s`, marshalString)
}
