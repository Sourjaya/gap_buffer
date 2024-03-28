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
func (gap *Gap) Gap_Length() int {
	return gap.end - gap.start + 1
}
func (gap *Gap) Buffer_Length() int {
	return len(gap.buffer)
}
func (gap *Gap) Insert(s []rune, position int) {
	if gap.Gap_Length() == 0 {
		gap.Grow(gap.Buffer_Length())
	}
	if position != gap.start {
		gap.Move_cursor(position)
	}
	for i := 0; i < len(s); i++ {
		if gap.Gap_Length() == 0 {
			gap.Grow(position)
		}
		gap.buffer[gap.start] = s[i]
		gap.start += 1
		position += 1
	}
}
func (gap *Gap) Move_cursor(position int) {
	if position < gap.start {
		gap.Left(position)
	} else {
		gap.Right(position)
	}
}
func (gap *Gap) Select_Delete(l, r int) {
	if l != gap.start {
		gap.Move_cursor(l)
	}
	// if r < l {
	// 	return
	// }
	// size := l - r + 1
	// if l < gap.start && r < gap.start {

	// } else if l > gap.end && r > gap.end {
	// 	fmt.Println()
	// } else {

	// }
	// gap.start -= size
}
func (gap *Gap) Delete() {
	if gap.end == gap.Buffer_Length()-1 {
		return
	}
	gap.end += 1
	gap.buffer[gap.end] = 0
}
func (gap *Gap) Backspace() {
	if gap.start == 0 {
		return
	}
	gap.start -= 1
	gap.buffer[gap.start] = 0
}
func (gap *Gap) Grow(position int) {
	newbuffer := make([]rune, len(gap.buffer)*2)
	copy(newbuffer, gap.buffer[:position])
	copy(newbuffer[gap.end+1+len(gap.buffer):], gap.buffer[gap.end+1:])
	gap.start = position
	gap.end = gap.end + len(gap.buffer)
	gap.buffer = newbuffer
}
func (gap *Gap) Left(position int) {
	for position < gap.start {
		gap.start -= 1
		gap.end -= 1
		gap.buffer[gap.end+1] = gap.buffer[gap.start]
		gap.buffer[gap.start] = 0
	}
}
func (gap *Gap) Right(position int) {
	for position > gap.start {
		gap.start += 1
		gap.end += 1
		gap.buffer[gap.start-1] = gap.buffer[gap.end]
		gap.buffer[gap.end] = 0
	}
}
