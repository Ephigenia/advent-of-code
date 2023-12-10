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

function sumArrAbs(numArr: number[]) {
  return numArr.reduce((acc, cur) => acc + Math.abs(cur), 0);
}

function arrLastItem(arr: number[]): number {
  return arr[arr.length - 1];
}

async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();
  const input = parseInput(rawInput)

  const lastSteps:number[] = [];
  let answer = 0;
  input.forEach((line) => {
    const sequence = line.split(/\s+/).map(v => parseInt(v, 10))

    let sum = 0;
    let steps:number[][] = [sequence];

    do {
      steps.push(findDifference(steps[steps.length - 1]));
      sum = sumArrAbs(steps[steps.length - 1]);
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

    // console.log(steps.map((line, i) => ' '.repeat(i) + line.join(' ')).join('\n'));
    // extract the last value of the first step of each
    const lastStepValue = arrLastItem(steps[0]);
    lastSteps.push(lastStepValue);
  });

  console.log(sumArr(lastSteps));

  // second part
  const firstSteps: number[] = [];
  input.forEach((line) => {
    const sequence = line.split(/\s+/).map(v => parseInt(v, 10))
    let sum = 0;
    let steps:number[][] = [sequence];

    do {
      steps.push(findDifference(steps[steps.length - 1]));
      sum = sumArrAbs(steps[steps.length - 1]);
    } while (sum !== 0);
    for (let i = steps.length - 1; i >= 0; i--) {
      if (i === steps.length - 1) {
        steps[i].unshift(0);
      } else if (i === steps.length - 2) {
        steps[i].unshift(steps[i][0]);
      } else {
        steps[i].unshift(steps[i][0] - steps[i+1][0]);
      }
    }

    // console.log(steps.map((line, i) => ' '.repeat(i) + line.join(' ')).join('\n'));
    firstSteps.push(steps[0][0]);
  })
  console.log(sumArr(firstSteps));
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
