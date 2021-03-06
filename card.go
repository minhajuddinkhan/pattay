package pattay

import "fmt"

const (
	Two   = iota
	Three = iota
	Four  = iota
	Five  = iota
	Six   = iota
	Seven = iota
	Eight = iota
	Nine  = iota
	Ten   = iota
	Jack  = iota
	Queen = iota
	King  = iota
	Ace   = iota
)

const (
	Spade   = "Spades"
	Diamond = "Diamonds"
	Club    = "Clubs"
	Heart   = "Hearts"
)

//Card Card
type Card interface {
	//House returns house of the card
	House() string
	//Number returns number of the card
	Number() int
}

//NewCard NewCard
func NewCard(house string, cardNumber int) Card {
	return &card{house: house, cardNumber: cardNumber}
}

type card struct {
	house      string
	cardNumber int
}

func (c *card) House() string {
	return c.house
}

func (c *card) Number() int {
	return c.cardNumber
}

//GetBiggestCard GetBiggestCard
func GetBiggestCard(cards []Card, house string) Card {

	max := 0
	for i := 0; i < len(cards); i++ {
		if cards[i].House() != house {
			continue
		}
		if cards[i].Number() > cards[max].Number() {
			max = i
		}
	}
	return cards[max]
}

//IsSameCard checks if the card is same
func IsSameCard(c1 Card, c2 Card) bool {
	return c1.House() == c2.House() && c1.Number() == c2.Number()
}

//FindCardInCards returns card and index. returns error if not found
func FindCardInCards(card Card, cards []Card) (Card, int, error) {
	for i, c := range cards {
		if IsSameCard(c, card) {
			return card, i, nil
		}
	}
	return nil, -1, fmt.Errorf("card not found")
}
