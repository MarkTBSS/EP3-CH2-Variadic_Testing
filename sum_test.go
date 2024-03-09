package sum

import "testing"

// test nil
func TestSum(t *testing.T) {
	t.Run("Should return 0 when got nothing", func(t *testing.T) {
		// Arrange
		want := 0
		// Act
		got := sum([]int{}...)
		// Assert
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("Should return 5 when got 2, 3", func(t *testing.T) {
		// Arrange
		want := 5
		// Act
		got := sum(2, 3)
		// Assert
		if got != want {
			t.Error("Expected", 5, "Got", got)
		}
	})
	t.Run("Sum All", func(t *testing.T) {
		// Arrange
		want := 7
		// Act
		got := sum([]int{2, 3, 3, -1}...)
		// Assert
		if got != want {
			t.Error("Expected", 8, "Got", got)
		}
	})
}
