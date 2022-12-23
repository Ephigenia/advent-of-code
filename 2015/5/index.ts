import { readFile } from "fs/promises";
import { resolve } from "path";

const VOVELS = ['a', 'e', 'i', 'o', 'u'];
const FORBIDDEN_STRINGS_DEFAULT = ['ab', 'cd', 'pq', 'or', 'xy'];

// get the distinct number of vovels (max 5)
function getDistinctNumberOfVovels(str: string): number {
  return [...new Set(str)].filter(s => VOVELS.includes(s)).length;
}

// get number of vovels in the string
function getNumberOfVovels(str: string): number {
  return str.length - str.replace(/[aeiou]/g, '').length;
}

// check if any of the characters is repeated 2 or more times
function hasRepeatedChars(str: string):boolean {
  return /(\w)\1{1,}/ig.test(str);
}

function containsForbiddenStrings(str: string, forbiddenStrings = FORBIDDEN_STRINGS_DEFAULT): boolean {
  return forbiddenStrings
    .filter(needle => str.indexOf(needle) > 0).length > 0;
}

function isNiceString(str: string): boolean {
  return !isNaughtyString(str);
}

function isNaughtyString(str: string): boolean  {
  return containsForbiddenStrings(str) ||
    !(
      hasRepeatedChars(str) &&
      getNumberOfVovels(str) >= 3
    );
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const lines = (await readFile(inputFilename)).toString().trim().split(/\n/);

  const niceStrings = lines.filter((str) => {
    const isNice = isNiceString(str);
    console.log(
      str,
      getDistinctNumberOfVovels(str),
      isNice,
    );
    return isNice;
  });

  console.log('%d strings are nice!\n', niceStrings.length);
}

main('input.txt');
main('inputTest.txt');
