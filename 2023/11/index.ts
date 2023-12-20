import { readFile } from "fs/promises";
import { resolve } from "path";

const SPACE = '.';
const GALAXY = '#';

interface Node {
  x: number;
  y: number;
}

function insertRow(universe: string[][], y: number) {
  return universe.splice(y, 0, Array(universe[0].length).fill(SPACE));
}
function insertColumn(universe: string[][], x: number) {
  universe.forEach(row => row.splice(x, 0, SPACE));
}

function generatePairs(arr: any[]) {
  return arr.map( (v, i) => arr.slice(i + 1).map(w => [v, w]) ).flat();
}

function getDistance(a: Node, b: Node) {
  return Math.abs(a.x - b.x) + Math.abs(a.y - b.y);
}

async function main(filename: string) {
  console.time('main');
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();

  // expand the galaxy
  // find rows and columns with no galaxy
  const universe = rawInput.split('\n').map((row) => row.split(''));

  const expandRows = universe.reduce((acc, row, y) => {
    if (row.every(v => v === SPACE)) return [...acc, y];
    return acc;
  }, [] as number[]);

  const expandCols:number[] = [];
  for (let x = 0; x < universe[0].length; x++) {
    const column = [...Array(universe.length).keys()].map(y => universe[y][x]);
    if (column.every(v => v === SPACE)) expandCols.push(x);
  }
  expandRows.forEach((rowIndex, i) => insertRow(universe, rowIndex + i));
  expandCols.forEach((colIndex, i) => insertColumn(universe, colIndex + i));

  // create tree of nodes
  const nodes:Node[] = [];
  for(let y = 0; y < universe.length; y++) {
    for(let x = 0; x < universe[0].length; x++) {
      const node = universe[y][x];
      if (node === GALAXY) {
        nodes.push({ x, y });
      }
    }
  }

  // distance between nodes
  const pairs = generatePairs(nodes);
  const distances = pairs.map(([a, b]) => getDistance(a, b));

  console.log();
  console.log(universe.map(m => m.join('')).join('\n'));

  console.log('distances sum', distances.reduce((acc, v) => acc + v, 0));

  console.timeEnd('main');
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
