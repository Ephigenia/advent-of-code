import { readFile } from "fs/promises";
import { resolve } from "path";

function sum(arr: number[]) {
  return arr.reduce((a:number, c:number) => (a + c), 0);
}

function parseLine(line: string): [number, number, number] {
  const [l, w, h] = line.split(/x/).map(v => parseInt(v, 10));
  return [l, w, h];
}

function calculate(line: string): number {
  const [l,w,h] = parseLine(line);
  const product = (2 * l * w) + (2 * w * h) + (2 * h * l);
  return product;
}

function calculateBow(x: number, y: number, z: number): number {
  return (x + x + y + y) + x * y * z;
}

function smallestSide(line: string) {
  const sides = parseLine(line);
  const sortedSides = sides.sort((a, b) => a - b);
  return sortedSides;
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const lines = (await readFile(inputFilename)).toString().trim().split(/\n/);

  const totals = lines.map((line) => {
    const small = smallestSide(line);
    return calculate(line) + small[0] * small[1];
  });
  console.log('total paper required: %d', sum(totals));

  const bows = lines.map((line) => {
    const s = smallestSide(line);
    return (s[0] + s[0] + s[1] + s[1]) + s[0] * s[1] * s[2];
  });
  console.log('total feet of ribbon required: %d', sum(bows));
}

main('input.txt');
