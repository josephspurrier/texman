// Package texman provides a few simple functions to make characters and line
// changes to a text file.
package texman

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Simply represents a simple lanaguage to modify text files.
type Simply struct {
	filename string
	content  [][]rune

	// LineEnding should be set to either "\n" or "\r\n". It is "\n" by default.
	LineEnding string
}

// NewFile returns an object of a file.
func NewFile(filename string) *Simply {
	return &Simply{
		filename:   filename,
		content:    make([][]rune, 0),
		LineEnding: "\n",
	}
}

// Load will read a file into memory
func (s *Simply) Load() error {
	s.content = make([][]rune, 0)

	f, err := os.Open(s.filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		col := 0
		arrCol := make([]rune, len([]rune(line)))
		for _, r := range line {
			arrCol[col] = r
			col++
		}
		s.content = append(s.content, arrCol)
		row++
	}

	return nil
}

// String will return a string representation of the text.
func (s *Simply) String() string {
	out := ""
	for r := 0; r < len(s.content); r++ {
		row := s.content[r]
		for c := 0; c < len(row); c++ {
			out += string(row[c])
		}

		if r < len(s.content)-1 {
			out += s.LineEnding
		}
	}
	return out
}

// Byte will return a byte array representation of the text.
func (s *Simply) Byte() []byte {
	return []byte(s.String())
}

func validate(row int, col int) error {
	if row < 1 {
		return errors.New("row cannot be less than 1")
	}
	if col < 1 {
		return errors.New("column cannot be less than 1")
	}
	return nil
}

// Overwrite will overwrite content at a specific location.
func (s *Simply) Overwrite(row int, col int, content string) error {
	err := validate(row, col)
	if err != nil {
		return err
	}

	// Fix the offset since computers start at 0.
	row--
	col--

	// Pad the row.
	for len(s.content) <= row {
		s.content = append(s.content, make([]rune, 0))
	}

	contentLen := len([]rune(content))

	// Pad the column.
	for len(s.content[row]) < col+contentLen {
		s.content[row] = append(s.content[row], ' ')
	}

	// Overwrite the character.
	for i := 0; i < len(content); i++ {
		s.content[row][col+i] = rune(content[i])
	}

	return nil
}

// Insert will insert content at a specific location.
func (s *Simply) Insert(row int, col int, content string) error {
	err := validate(row, col)
	if err != nil {
		return err
	}

	// Fix the offset since computers start at 0.
	row--
	col--

	// Pad the row.
	for len(s.content) <= row {
		s.content = append(s.content, make([]rune, 0))
	}

	contentLen := len([]rune(content))
	columnLen := len(s.content[row])

	padLen := contentLen

	// If the column is already larger than the column, add that as padding too.
	if col > columnLen {
		padLen += col - columnLen
	}

	// Pad the column.
	for i := 0; i < padLen; i++ {
		s.content[row] = append(s.content[row], ' ')
	}

	// Move the text over the length of the text being inserted.
	copy(s.content[row][col+contentLen:], s.content[row][col:])

	// Replace the old characters (which have already been moved).
	for i := 0; i < contentLen; i++ {
		s.content[row][col+i] = rune(content[i])
	}

	return nil
}

// InsertLine will insert a new line.
func (s *Simply) InsertLine(row int, col int) error {
	return s.Insert(row, col, s.LineEnding)
}

// Delete a character at the specified location.
func (s *Simply) Delete(row int, col int) error {
	err := validate(row, col)
	if err != nil {
		return err
	}

	// Fix the offset since computers start at 0.
	row--
	col--

	if row >= len(s.content) {
		return fmt.Errorf("cannot delete, row does not exist: %v", row)
	}

	if col >= len(s.content[row]) {
		return fmt.Errorf("cannot delete, column does not exist: %v", row)
	}

	// Move the slice left.
	copy(s.content[row][col:], s.content[row][col+1:])

	// Delete the last element.
	s.content[row] = s.content[row][:len(s.content[row])-1]

	return nil
}

// DeleteLine will delete a line.
func (s *Simply) DeleteLine(row int) error {
	err := validate(row, 1)
	if err != nil {
		return err
	}

	// Fix the offset since computers start at 0.
	row--

	if row >= len(s.content) {
		return fmt.Errorf("cannot delete line, row does not exist: %v", row)
	}

	// Move slice down.
	copy(s.content[row:], s.content[row+1:])

	// Delete the last element.
	s.content = s.content[:len(s.content)-1]

	return nil
}
