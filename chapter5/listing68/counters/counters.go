package counters

// alia
type alertCounter int

// create as factory
func New(value int) alertCounter {
	return alertCounter(value)
}
