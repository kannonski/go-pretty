package table

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleAutoIndexColumnID() {
	fmt.Printf("AutoIndexColumnID(    0): \"%s\"\n", AutoIndexColumnID(0))
	fmt.Printf("AutoIndexColumnID(    1): \"%s\"\n", AutoIndexColumnID(1))
	fmt.Printf("AutoIndexColumnID(    2): \"%s\"\n", AutoIndexColumnID(2))
	fmt.Printf("AutoIndexColumnID(   25): \"%s\"\n", AutoIndexColumnID(25))
	fmt.Printf("AutoIndexColumnID(   26): \"%s\"\n", AutoIndexColumnID(26))
	fmt.Printf("AutoIndexColumnID(  702): \"%s\"\n", AutoIndexColumnID(702))
	fmt.Printf("AutoIndexColumnID(18278): \"%s\"\n", AutoIndexColumnID(18278))

	// Output: AutoIndexColumnID(    0): "A"
	// AutoIndexColumnID(    1): "B"
	// AutoIndexColumnID(    2): "C"
	// AutoIndexColumnID(   25): "Z"
	// AutoIndexColumnID(   26): "AA"
	// AutoIndexColumnID(  702): "AAA"
	// AutoIndexColumnID(18278): "AAAA"
}

func TestAutoIndexColumnID(t *testing.T) {
	assert.Equal(t, "A", AutoIndexColumnID(0))
	assert.Equal(t, "Z", AutoIndexColumnID(25))
	assert.Equal(t, "AA", AutoIndexColumnID(26))
	assert.Equal(t, "ZZ", AutoIndexColumnID(701))
	assert.Equal(t, "AAA", AutoIndexColumnID(702))
	assert.Equal(t, "ZZZ", AutoIndexColumnID(18277))
	assert.Equal(t, "AAAA", AutoIndexColumnID(18278))
}

func TestComputeBoxStyle(t *testing.T) {
	assertOutput := func(expected, actual BoxStyle) {
		assert.Equal(t, expected, actual)
		if expected != actual {
			fmt.Printf("%#v", actual)
		}
	}

	t.Run("style 1", func(t *testing.T) {
		input := `
┏━━━┳━━━┓
┣━━━╋━━━┫
┃< >┃< >┃ ~
┗━━━┻━━━┛
`
		expectedOutput := BoxStyle{
			BottomLeft:       "┗",
			BottomRight:      "┛",
			BottomSeparator:  "┻",
			EmptySeparator:   " ",
			Left:             "┃",
			LeftSeparator:    "┣",
			MiddleHorizontal: "━",
			MiddleSeparator:  "╋",
			MiddleVertical:   "┃",
			PaddingLeft:      "<",
			PaddingRight:     ">",
			PageSeparator:    "\n",
			Right:            "┃",
			RightSeparator:   "┫",
			TopLeft:          "┏",
			TopRight:         "┓",
			TopSeparator:     "┳",
			UnfinishedRow:    " ~",
		}

		assertOutput(expectedOutput, ComputeBoxStyle(input))
		assertOutput(expectedOutput, ComputeBoxStyle(strings.TrimSpace(input)))
	})

	t.Run("style 2", func(t *testing.T) {
		input := "╔═══╦═══╗\n" + "╠═══╬═══╣\n" + "║   ║   ║ ≈\n" + "╚═══╩═══╝\n"
		expectedOutput := BoxStyle{
			BottomLeft:       "╚",
			BottomRight:      "╝",
			BottomSeparator:  "╩",
			EmptySeparator:   " ",
			Left:             "║",
			LeftSeparator:    "╠",
			MiddleHorizontal: "═",
			MiddleSeparator:  "╬",
			MiddleVertical:   "║",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			PageSeparator:    "\n",
			Right:            "║",
			RightSeparator:   "╣",
			TopLeft:          "╔",
			TopRight:         "╗",
			TopSeparator:     "╦",
			UnfinishedRow:    " ≈",
		}

		assertOutput(expectedOutput, ComputeBoxStyle(input))
		assertOutput(expectedOutput, ComputeBoxStyle(strings.TrimSpace(input)))
	})
}

func TestComputeBoxConnectorStyle(t *testing.T) {
	assertOutput := func(expected, actual BoxConnectorStyle) {
		assert.Equal(t, expected, actual)
		if expected != actual {
			fmt.Printf("%#v", actual)
		}
	}

	t.Run("style 1", func(t *testing.T) {
		input := `
┣━━━╋━━━┫ ~
`
		expectedOutput := BoxConnectorStyle{
			LeftSeparator:    "┣",
			MiddleHorizontal: "━",
			MiddleSeparator:  "╋",
			RightSeparator:   "┫",
			UnfinishedRow:    " ~",
		}

		assertOutput(expectedOutput, ComputeBoxConnectorStyle(input))
		assertOutput(expectedOutput, ComputeBoxConnectorStyle(strings.TrimSpace(input)))
	})

	t.Run("style 2", func(t *testing.T) {
		input := "╠═══╬═══╣ ≈\n"
		expectedOutput := BoxConnectorStyle{
			LeftSeparator:    "╠",
			MiddleHorizontal: "═",
			MiddleSeparator:  "╬",
			RightSeparator:   "╣",
			UnfinishedRow:    " ≈",
		}

		assertOutput(expectedOutput, ComputeBoxConnectorStyle(input))
		assertOutput(expectedOutput, ComputeBoxConnectorStyle(strings.TrimSpace(input)))
	})
}

func Test_isNumber(t *testing.T) {
	assert.True(t, isNumber(int(1)))
	assert.True(t, isNumber(int8(1)))
	assert.True(t, isNumber(int16(1)))
	assert.True(t, isNumber(int32(1)))
	assert.True(t, isNumber(int64(1)))
	assert.True(t, isNumber(uint(1)))
	assert.True(t, isNumber(uint8(1)))
	assert.True(t, isNumber(uint16(1)))
	assert.True(t, isNumber(uint32(1)))
	assert.True(t, isNumber(uint64(1)))
	assert.True(t, isNumber(float32(1)))
	assert.True(t, isNumber(float64(1)))
	assert.False(t, isNumber("1"))
	assert.False(t, isNumber(nil))
}
