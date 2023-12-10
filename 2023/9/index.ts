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

function arrLastItem(arr: number[]): number {
  return arr[arr.length - 1];
}

async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();
  const input = parseInput(rawInput)

  const lastSteps:number[] = [];
  input.forEach((line) => {
    const sequence = line.split(/\s+/).map(v => parseInt(v, 10))

    let sum = 0;
    let steps:number[][] = [sequence];

    do {
      steps.push(findDifference(steps[steps.length - 1]));
      sum = sumArr(steps[steps.length - 1]);
    } while (sum !== 0);

    // reverse extrapolate each step using the difference from the next step
    for (let i = steps.length - 1; i >= 0; i--) {
      if (i === steps.length - 1) {
        steps[i].push(0);
      } else if (i === steps.length - 2) {
        steps[i].push(steps[i][steps[i].length - 1]);
      } else {
        const incr = arrLastItem(steps[i + 1]);
        const lastItem = arrLastItem(steps[i]);
        steps[i].push(lastItem + incr);
      }
    }

    // extract the last value of the first step of each
    const lastStepValue = arrLastItem(steps[0]);
    lastSteps.push(lastStepValue);
  });

  console.log(sumArr(lastSteps));
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
