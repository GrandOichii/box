package box

import (
	"errors"
	"fmt"
	"strings"

	"github.com/GrandOichii/colorwrapper"
)

const (
	sepPrefix = "&sep"
	HLine     = "\u2500"
	VLine     = "\u2502"
	RDCorner  = "\u250C"
	LDCorner  = "\u2510"
	RUCorner  = "\u2514"
	LUCorner  = "\u2518"
	RTree     = "\u251C"
	LTree     = "\u2524"
	DTree     = "\u252C"
	UTree     = "\u2534"
	Cross     = "\u253C"
)

func StrWidthSplit(message string, width int) []string {
	result := []string{}
	split := strings.Split(message, " ")
	line := split[0]
	for i, word := range split {
		if i == 0 {
			continue
		}
		if len(line+" "+word) > width {
			result = append(result, line)
			line = word
		} else {
			line += " " + word
		}
	}

	return append(result, line)
}

func Separator(colorPair string) string {
	return sepPrefix + colorPair
}

func Draw(height int, width int, lines []string, borderColorPair string) error {
	if len(lines) > height-2 {
		return fmt.Errorf("boxch - can't fit all lines in box(lines: %v, height: %v)", len(lines), height)
	}

	coloredTop, err := colorwrapper.GetColored(borderColorPair, RDCorner+strings.Repeat(HLine, width-2)+LDCorner)
	if err != nil {
		return err
	}
	coloredBottom, err := colorwrapper.GetColored(borderColorPair, RUCorner+strings.Repeat(HLine, width-2)+LUCorner)
	if err != nil {
		return err
	}
	coloredVLine, err := colorwrapper.GetColored(borderColorPair, VLine)
	if err != nil {
		return err
	}
	for i := 0; i < height; i++ {
		out := ""
		switch i {
		case 0:
			out += coloredTop
		case height - 1:
			out += coloredBottom
		default:
			out += coloredVLine
			mid := ""
			mid = strings.Repeat(" ", width-2)
			out += mid + coloredVLine
			if i < len(lines)+1 {
				line := lines[i-1]
				if strings.HasPrefix(line, sepPrefix) {
					if len(line) < 4 {
						return errors.New("box - no color for separator")
					}
					sepColorPair := line[4:]
					coloredSeparator, err := colorwrapper.GetColored(sepColorPair, RTree+strings.Repeat(HLine, width-2)+LTree)
					if err != nil {
						return err
					}
					out += "\r" + coloredSeparator
				} else {
					out += "\r" + coloredVLine + line
				}
			}
		}
		fmt.Println(out)
	}
	return nil
}
