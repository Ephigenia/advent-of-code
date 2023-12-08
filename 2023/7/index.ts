import { readFile } from "fs/promises";
import { resolve } from "path";

function parseInput(input: string) {
  return input.split('\n');
}

// relative strengs decreasing, A strongest, 2 weakest
// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2

const scoreMap = ['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'].reverse();

function getScoreForCard(card: string) {
  return scoreMap.indexOf(card) + 1;
}

class Card {
  public score = 0;
  constructor(public symbol : string, public count: number = 0) {
    this.score = getScoreForCard(symbol);
  }
}

class Hand {
  cards: Card[] = [];
  score: number = 0;
  constructor(cardSymbols: string[], public bid: number = 0) {
    cardSymbols.forEach((symbol) => {
      const card = this.cards.find(card => card.symbol === symbol);
      if (card) card.count++;
      else this.cards.push(new Card(symbol));
    });
    this.cards = this.cards.sort((a, b) => b.count - a.count);

    // find the card with the most occurrences
    const cardScore = this.cards.reduce((a, c) => a+c.score, 0);
    if (this.cards.length === 1) {
      this.score = cardScore * 20;
    } else if (this.cards.length === 2) {
      if (this.cards[0].count === 4) {
        this.score = cardScore * 10;
      } else if (this.cards[0].count === 3 && this.cards[1].count === 2) {
        this.score = cardScore * 5;
      }
    } else if (this.cards.length === 3) {
      this.score = cardScore * 3;
    } else if (this.cards.length === 4) {
      this.score = cardScore * 2;
    } else if (this.cards.length === 5) {
      this.score = cardScore;
    }
  }

  toString() {
    return this.cards.map((card) => card.symbol.repeat(card.count+1)).join('');
  }
}

interface Card {
  symbol: string;
  count: number;
}

function getSameCards(symbol: string[]) {
  const counter: Card[] = [];
  symbol.forEach((symbol) => {
    const card = counter.find(card => card.symbol === symbol);
    if (card) card.count++;
  });
  console.log(counter);;
}

function getScoreHand(cards: string[]) {

  getSameCards(cards)
  const score = cards.reduce((acc, card) => acc + getScoreForCard(card), 0);
  return score;
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
  // console.log(hands);

  hands.forEach((hand, i) => {
    console.log(i, hand.toString(), hand.score);
  });
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
