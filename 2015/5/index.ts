import { readFile } from "fs/promises";
import { resolve } from "path";

const VOVELS = ['a', 'e', 'i', 'o', 'u'];
const FORBIDDEN_STRINGS_DEFAULT = ['ab', 'cd', 'pq', 'xy'];

function getNumberOfVovels(str: string): number {
  return str.split(/[aeiou]/).length;
}

function hasRepeatedChars(str: string):boolean {
  return /(\w)\1/i.test(str);
}

function containsForbiddenStrings(str: string, forbiddenStrings = FORBIDDEN_STRINGS_DEFAULT): boolean {
  return forbiddenStrings.some(needle => str.includes(needle));
}

function isNiceString(str: string): boolean {
  return !containsForbiddenStrings(str) &&
    (
      hasRepeatedChars(str) &&
      getNumberOfVovels(str) > 3
    );
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const lines = (await readFile(inputFilename)).toString().trim().split(/\n/);

  // Day1
  const niceStrings = lines.filter(isNiceString);
  console.log('%d strings are nice!\n', niceStrings.length);

  // Day2
  const niceStrings2 = lines.filter(line => {
    const overlapping = line.match(/.*(.{2}).*\1.*/g);
    const pairs = line.match(/.*(.).\1.*/g);
    return overlapping && pairs;
  });
  console.log('%d strings are nice!\n', niceStrings2.length);
}

main('input.txt');
// main('inputTest.txt');
