import { readFile } from "fs/promises";
import { resolve } from "path";

function parseInput(input: string) {
  return input.split('\n');
}

function getScoreForCard(card: string) {
  const scoreMap = ['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'].reverse();
  return scoreMap.indexOf(card) + 2;
}

class Card {
  public score = 0;
  constructor(public symbol : string, public count: number = 1) {
    this.score = getScoreForCard(symbol);
  }
}

class Hand {
  sortedCards: Card[] = [];
  cards: Card[] = [];
  score: number = 0;
  constructor(cardSymbols: string[], public bid: number = 0) {
    this.cards = cardSymbols.map((symbol) => new Card(symbol));

    cardSymbols.forEach((symbol) => {
      const card = this.sortedCards.find(card => card.symbol === symbol);
      if (card) card.count++;
      else this.sortedCards.push(new Card(symbol));
    });
    this.sortedCards = this.sortedCards.sort((a, b) => b.count - a.count);

    if (this.sortedCards[0].count === 5) {
      this.score = 6;
    } else if (this.sortedCards[0].count === 4) {
      this.score = 5;
    } else if (this.sortedCards[0].count === 3 && this.sortedCards[1].count === 2) {
      this.score = 4;
    } else if (this.sortedCards[0].count === 3 && this.sortedCards[1].count === 1) {
      this.score = 3;
    } else if (this.sortedCards[0].count === 2 && this.sortedCards[1].count === 2) {
      this.score = 2;
    } else if (this.sortedCards[0].count === 2 && this.sortedCards[1].count === 1) {
      this.score = 1;
    } else {
      this.score = 0;
    }
  }
}

interface Card {
  symbol: string;
  count: number;
}

async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();
  const input = parseInput(rawInput)

  const hands = input.map((hand) => {
    const [cardsString, bidString] = hand.split(' ');
    const cards = cardsString.split('');
    const bid = parseInt(bidString, 10);
    return new Hand(cards, bid);
  });

  const sortedHands = hands.sort((a: Hand, b: Hand) => {
    if (a.score > b.score) return 1;
    if (a.score < b.score) return -1;
    if (a.score === b.score) {
      // both ranks are equal, compare hands card by card
      for (let i = 0; i < a.cards.length; i++) {
        if (a.cards[i].score > b.cards[i].score) return 1;
        if (a.cards[i].score < b.cards[i].score) return -1;
        continue;
      }
    }
    return 0;
  });

  const sumBids = sortedHands.reduce((a, hand, index) => a + hand.bid * (index + 1), 0);
  console.log('sum', sumBids);
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
