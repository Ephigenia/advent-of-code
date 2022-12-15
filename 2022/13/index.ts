import { readFile } from "fs/promises";
import { resolve } from "path";

type SignalStrength = number | number[];
type PackageValue = SignalStrength[];

class SignalPair {
  left: number[];
  right: number[];
  constructor(left: string, right: string) {
    this.left = this.parsePacketString(left);
    this.right = this.parsePacketString(right);
  }

  parsePacketString(input: string) {
    return JSON.parse(input);
  }
}


function compare(a: PackageValue, b: PackageValue, level = 0): boolean | null {
  console.log(
    '  ' + '  '.repeat(level) + 'Compare %j vs %j', a, b
  );

  for (const [index, left] of a.entries()) {
    const right = b[index];
    // right side ran out of items
    if (right === undefined) return false;
    // compare numbers
    if (typeof left === 'number' && typeof right === 'number') {
      if (left === right) continue;
      return left < right;
    }
    // compare arrays
    const comparison = compare(Array.isArray(left) ? left : [left], Array.isArray(right) ? right : [right], level + 1);
    if (comparison !== null) return comparison;
  }

  // left side runs out of items
  return a.length < b. length ? true : null;
}

// use FORCE_COLOR=0 to disable colors
async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();
  const pairStrings = raw.split(/\n\n/).map(str => str.split(/\n/));

  const pairs = pairStrings.map(([left, right], i) => {
    return new SignalPair(left, right);
  });

  const correctPairs:SignalPair[] = [];
  pairs.forEach((pair, i) => {
    console.log('== Pair %d ==', i + 1);
    const result = compare(pair.left, pair.right);
    if (result) {
      correctPairs.push(pair);
    }
    return pair;
  });

  const correctPairIndices = correctPairs.map(pair => pairs.indexOf(pair) + 1);
  console.log('Found %d correct pairs (%j)', correctPairs.length, correctPairIndices);
  const sumIndices = correctPairIndices.reduce((acc, cur) => acc + cur, 0);
  console.log('indices of those pairs is %d', sumIndices);

  const decoderPackages = [[[2]], [[6]]];
  const sets: PackageValue[] =
    [
      ...pairs.map((pair) => [pair.left, pair.right]).flat(),
      [[2]], [[6]]
      // ...decoderPackages
    ]
    .sort((a, b) => {
      return compare(a, b) ? -1 : 1;
    });
  sets.forEach((set) => {
    console.log(set);
  });

  // const f = sets.indexOf(decoderPackages);
  const stringSets = sets.map(s => JSON.stringify(s));
  console.log(stringSets);
  const decoderPackageIndices = decoderPackages.map((decoderPackage) => {
    const str = (JSON.stringify(decoderPackage));
    return stringSets.indexOf(str) + 1;
  })
  console.log('the divider packages are at %j', decoderPackageIndices);
  console.log('decoder key is %d', decoderPackageIndices.reduce((a, c) => a * c, 1));

}

// main('inputTest.txt');
main('input.txt');
