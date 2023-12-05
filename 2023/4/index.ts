import chalk from "chalk";
import { readFile } from "fs/promises";
import { resolve } from "path";

function parseInput(input: string) {
  return input.split('\n');
}

class Card {
  public index;
  public numbers: number[];
  public winningNumbers: number[];
  constructor(private input: string) {
    const NUMBER_SEPERATOR = /\s+/;
    const INDEX_NUMBER_SEPERATOR = ":";
    const WINNING_NUMBERS_SEPERATOR = "|";
    this.index = parseInt(input.substring(5, input.indexOf(INDEX_NUMBER_SEPERATOR)), 10);
    this.numbers = input.substring(input.indexOf(INDEX_NUMBER_SEPERATOR) + 1, input.indexOf(WINNING_NUMBERS_SEPERATOR)).trim().split(NUMBER_SEPERATOR).map(n => parseInt(n, 10));
    this.winningNumbers = input.substring(input.indexOf(WINNING_NUMBERS_SEPERATOR) + 3).trim().split(NUMBER_SEPERATOR).map(n => parseInt(n, 10));
  }

  public toString() {
    return `Card ${this.index}: ${this.numbers.join(' ')} | ${this.winningNumbers.join(' ')}`;
  }

  public isWinningNumber(number: number) {
    return this.winningNumbers.includes(number);
  }

  public getMatchingNumbers() {
    return this.numbers.reduce((acc: number[], cur: number) => {
      return this.isWinningNumber(cur) ? [...acc, cur] : acc;
    }, []);
  }

  public getScore() {
    const matchingNumbers = this.getMatchingNumbers();
    if (matchingNumbers.length === 0) return 0;
    return Math.pow(2, (matchingNumbers.length - 1));
  }
}

async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();
  const input = parseInput(rawInput)

  const cards = input.map(line => new Card(line));

  cards.forEach(card => {
    console.log('%s win: %d points %d', card.toString(), card.getMatchingNumbers().length, card.getScore());
  });
  const sumPoints = cards.reduce((acc, cur) => acc + cur.getScore(), 0);
  console.log('all cards are worth %d points', sumPoints);
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
