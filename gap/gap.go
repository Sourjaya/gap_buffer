package gap

type Gap struct {
	buffer []rune
	start  int
	end    int
}

func New(s string) *Gap {
	gap := new(Gap)
	gap.buffer = []rune(s)
	gap.start = len(gap.buffer)
	gap.end = len(gap.buffer) - 1
	return gap
}
func (gap *Gap) GetString() string {

	// create a new rune slice and append the preGap and postGap slices to it before returning
	text := append([]rune{}, gap.buffer[:gap.start]...)
	text = append(text, gap.buffer[(gap.end+1):]...)

	return string(text)
}
func (gap *Gap) GapLength() int {
	return gap.end - gap.start + 1
}
func (gap *Gap) BufferLength() int {
	return len(gap.buffer)
}
func (gap *Gap) Insert(s []rune, position int) *Gap {
	if gap.GapLength() == 0 {
		gap = gap.Grow(gap.BufferLength())
	}
	if position != gap.start {
		gap = gap.MoveCursor(position)
	}
	for i := 0; i < len(s); i++ {
		if gap.GapLength() == 0 {
			gap = gap.Grow(position)
		}
		gap.buffer[gap.start] = s[i]
		gap.start += 1
		position += 1
	}
	return gap
}
func (gap *Gap) MoveCursor(position int) *Gap {
	if position < gap.start {
		gap = gap.Left(position)
	} else {
		gap = gap.Right(position)
	}
	return gap
}
func (gap *Gap) SelectDelete(l, r int, isFront bool) *Gap {

	if isFront {
		gap = gap.MoveCursor(r)
		for i := 0; i < r-l; i++ {
			gap = gap.Backspace()
		}
	} else {
		gap = gap.MoveCursor(l)
		for i := 0; i < r-l; i++ {
			gap = gap.Delete()
		}
	}
	return gap
}
func (gap *Gap) Delete() *Gap {
	if gap.end == gap.BufferLength()-1 {
		return gap
	}
	gap.end += 1
	gap.buffer[gap.end] = 0
	return gap
}
func (gap *Gap) Backspace() *Gap {
	if gap.start == 0 {
		return gap
	}
	gap.start -= 1
	gap.buffer[gap.start] = 0
	return gap
}
func (gap *Gap) Grow(position int) *Gap {
	newbuffer := make([]rune, len(gap.buffer)*2)
	copy(newbuffer, gap.buffer[:position])
	copy(newbuffer[gap.end+1+len(gap.buffer):], gap.buffer[gap.end+1:])
	gap.start = position
	gap.end = gap.end + len(gap.buffer)
	gap.buffer = newbuffer
	return gap
}
func (gap *Gap) Left(position int) *Gap {
	for position < gap.start {
		gap.start -= 1
		gap.end -= 1
		gap.buffer[gap.end+1] = gap.buffer[gap.start]
		gap.buffer[gap.start] = 0
	}
	return gap
}
func (gap *Gap) Right(position int) *Gap {
	for position > gap.start {
		gap.start += 1
		gap.end += 1
		gap.buffer[gap.start-1] = gap.buffer[gap.end]
		gap.buffer[gap.end] = 0
	}
	return gap
}
