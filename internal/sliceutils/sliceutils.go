package sliceutils

// RemoveItem removes the first occurrence of an item from a slice that satisfies the given predicate.
// It returns the modified slice and a boolean indicating whether an item was removed.
func RemoveItem[T any](slice []T, predicate func(T) bool) ([]T, bool) {
	for i, item := range slice {
		if predicate(item) {
			return append(slice[:i], slice[i+1:]...), true
		}
	}
	return slice, false
}

// RemoveAllItems removes all occurrences of items from a slice that satisfy the given predicate.
// It returns the modified slice and the number of items removed.
func RemoveAllItems[T any](slice []T, predicate func(T) bool) ([]T, int) {
	var removed int
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if !predicate(item) {
			result = append(result, item)
		} else {
			removed++
		}
	}
	return result, removed
}

// Example usage:
//
// import "github.com/yourusername/sliceutils"
//
// // Remove a specific asteroid
// am.asteroids, _ = sliceutils.RemoveItem(am.asteroids, func(a *Asteroid) bool {
//     return a == specificAsteroid
// })
//
// // Remove all destroyed asteroids
// am.asteroids, numRemoved = sliceutils.RemoveAllItems(am.asteroids, func(a *Asteroid) bool {
//     return a.IsDestroyed()
// })
