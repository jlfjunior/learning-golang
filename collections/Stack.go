package collections

type Stack struct {
	Items  []int
	Length int
}

func (c *Stack) Add(value int) {
	c.Items = append(c.Items, value)
	c.Length++
}

func (c *Stack) Pop() int {
	i := c.Length - 1
	value := c.Items[i]
	c.Items = c.Items[:i]
	c.Length--
	return value
}

func (c *Stack) Peek() int {
	return c.Items[c.Length-1]
}
