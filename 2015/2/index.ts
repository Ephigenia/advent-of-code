import { readFile } from "fs/promises";
import { resolve } from "path";

function sum(arr: number[]) {
  return arr.reduce((a:number, c:number) => (a + c), 0);
}

function parseLine(line: string): number[] {
  const [l, w, h] = line.split(/x/).map(v => parseInt(v, 10));
  return [l, w, h];
}

function calculate(line: string): number {
  const [l,w,h] = parseLine(line);
  const product = (2 * l * w) + (2 * w * h) + (2 * h * l);
  return product;
}

function smallestSide(line: string) {
  const sides = parseLine(line);
  const sorts = sides.sort((a, b) => a - b);
  return [sorts[0], sorts[1]];
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();

  const lines = raw.split(/\n/);
  const totals = lines.map((line) => {
    const small = smallestSide(line);
    return calculate(line) + small[0] * small[1];
  });

  console.log('total paper required: %d', sum(totals));

}

main('input.txt');
