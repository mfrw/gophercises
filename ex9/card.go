package deck

import "fmt"

//go:generate stringer -type=Suit,Rank
type Suit uint

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // Suitless in theory
)

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}
