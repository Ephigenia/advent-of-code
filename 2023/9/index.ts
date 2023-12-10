import { readFile } from "fs/promises";
import { resolve } from "path";

function parseInput(input: string) {
  return input.split('\n');
}

function findDifference(sequence: number[]) {
  return [...Array(sequence.length - 1).keys()].map(i => {
    return sequence[i + 1] - sequence[i]
  });
}

function sumArr(numArr: number[]) {
  return numArr.reduce((acc, cur) => acc + cur, 0);
}

async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();
  const input = parseInput(rawInput)

  input.forEach((line) => {
    const sequence = line.split(/\s+/).map(v => parseInt(v, 10))
    console.log(sequence)

    let sum = 0;
    let steps:number[][] = [sequence];

    do {
      steps.push(findDifference(steps[steps.length - 1]));
      sum = sumArr(steps[steps.length - 1]);
    } while (sum !== 0);

    console.log(steps);
  })
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
