import { readFile } from "fs/promises";
import { resolve } from "path";

function getCalibrationValue(digits: Array<Array<string|number>>): Array<number> {
  return digits
    .filter(d => d.length > 0)
    .map(digits => {
      const firstDigit = digits.shift();
      const secondDigit = digits.pop() || firstDigit;
      return parseInt(String(firstDigit + '' + secondDigit), 10);
    });
}

function processInput(input: string) {
  const lines = input.split('\n').map(line => line.trim()).filter(v => v);
  const digits = lines.map(line => line.split('').filter(v => v.match(/\d/)));
  const calibrationValues = getCalibrationValue(digits);
  const value = calibrationValues.reduce((acc, v) => acc + v, 0);
  return value;
}


function processInput2(input: string) {
  const lines = input.split('\n').map(line => line.trim()).filter(v => v);
  const digits = lines.map(extractDigits);
  const calibrationValues = getCalibrationValue(digits);
  const value = calibrationValues.reduce((acc, v) => acc + v, 0);
  return value;
}

function extractDigits(line: string) {
  const map = {
    one: 1,
    two: 2,
    three: 3,
    four: 4,
    five: 5,
    six: 6,
    seven: 7,
    eight: 8,
    nine: 9,
  }
  const digit = line.replace(/one|two|three|four|five|six|seven|eight|nine/, (match: string): string => {
    return String(map[match as keyof typeof map]);
  })
    .split('')
    .filter(v => parseInt(v, 10)
  );

  const map2 = {
    eno: 1,
    owt: 2,
    eerht: 3,
    ruof: 4,
    evif: 5,
    xis: 6,
    neves: 7,
    thgie: 8,
    enin: 9,
  }

  const reversedString = line.split('').reverse().join('');
  const digit2 = reversedString.replace(/eno|owt|eerht|ruof|evif|xis|neves|thgie|enin/, (match: string): string => {
    return String(map2[match as keyof typeof map2]);
  })
  .split('')
    .filter(v => parseInt(v, 10)
  );

  return [digit[0], digit2[0]];
}


async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();

  const value = processInput(rawInput);
  console.log(value);

  const value2 = processInput2(rawInput);
  console.log(value2);
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
