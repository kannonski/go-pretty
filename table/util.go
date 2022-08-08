package table

import (
	"reflect"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

// AutoIndexColumnID returns a unique Column ID/Name for the given Column Number.
// The functionality is similar to what you get in an Excel spreadsheet w.r.t.
// the Column ID/Name.
func AutoIndexColumnID(colIdx int) string {
	charIdx := colIdx % 26
	out := string(rune(65 + charIdx))
	colIdx = colIdx / 26
	if colIdx > 0 {
		return AutoIndexColumnID(colIdx-1) + out
	}
	return out
}

// ComputeBoxStyle extracts the characters needed to construct a box. For ex.,
// with the following input:
//  ┏━━━┳━━━┓
//  ┣━━━╋━━━┫
//  ┃< >┃< >┃ ~
//  ┗━━━┻━━━┛
// the output would be:
//   BoxStyle{
//       BottomLeft:       "┗",
//       BottomRight:      "┛",
//       BottomSeparator:  "┻",
//       EmptySeparator:   " ",
//       Left:             "┃",
//       LeftSeparator:    "┣",
//       MiddleHorizontal: "━",
//       MiddleSeparator:  "╋",
//       MiddleVertical:   "┃",
//       PaddingLeft:      "<",
//       PaddingRight:     ">",
//       PageSeparator:    "\n",
//       Right:            "┃",
//       RightSeparator:   "┫",
//       TopLeft:          "┏",
//       TopRight:         "┓",
//       TopSeparator:     "┳",
//       UnfinishedRow:    " ~",
//   }
func ComputeBoxStyle(boxDrawing string) BoxStyle {
	charMatrix := [4][10]string{}
	maxLines := len(charMatrix)
	maxLineLen := len(charMatrix[0])

	// extract boxDrawing into charMatrix
	boxDrawing = strings.TrimSpace(boxDrawing)
	boxDrawingLines := strings.Split(boxDrawing, "\n")
	for lineNum, line := range boxDrawingLines {
		if lineNum <= maxLines { // ignore all lines beyond max allowed lines
			line = strings.TrimSpace(line)

			// extract individual "characters"
			charNum := 0
			for _, char := range line {
				charMatrix[lineNum][charNum] += string(char)
				// don't count beyond max allowed strings (last string contains all trailing chars)
				if charNum < maxLineLen-1 {
					charNum++
				}
			}
		}
	}

	return BoxStyle{
		BottomLeft:       charMatrix[3][0],
		BottomRight:      charMatrix[3][8],
		BottomSeparator:  charMatrix[3][4],
		EmptySeparator:   text.RepeatAndTrim(" ", text.RuneWidthWithoutEscSequences(charMatrix[1][4])),
		Left:             charMatrix[2][0],
		LeftSeparator:    charMatrix[1][0],
		MiddleHorizontal: charMatrix[0][2],
		MiddleSeparator:  charMatrix[1][4],
		MiddleVertical:   charMatrix[2][4],
		PaddingLeft:      charMatrix[2][1],
		PaddingRight:     charMatrix[2][3],
		PageSeparator:    "\n",
		Right:            charMatrix[2][8],
		RightSeparator:   charMatrix[1][8],
		TopLeft:          charMatrix[0][0],
		TopRight:         charMatrix[0][8],
		TopSeparator:     charMatrix[0][4],
		UnfinishedRow:    charMatrix[2][9],
	}
}

// WidthEnforcer is a function that helps enforce a width condition on a string.
type WidthEnforcer func(col string, maxLen int) string

// widthEnforcerNone returns the input string as is without any modifications.
func widthEnforcerNone(col string, maxLen int) string {
	return col
}

// isNumber returns true if the argument is a numeric type; false otherwise.
func isNumber(x interface{}) bool {
	if x == nil {
		return false
	}

	switch reflect.TypeOf(x).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	}
	return false
}
