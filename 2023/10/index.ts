import chalk from "chalk";
import { readFile } from "fs/promises";
import { resolve } from "path";

interface Point { x: number, y: number };

class Surface {
  public shapes: string[][] = [[]];
  public width = 0;
  public height = 0;

  public constructor(lines: string) {
    this.shapes = lines.split('\n').map(line => line.split(''));
    this.width = this.shapes[0].length - 1;
    this.height = this.shapes.length - 1;
  }

  public get(x: number, y: number): string {
    return this.shapes[y][x];
  }

  public set(x: number, y: number, value: string): Surface {
    this.shapes[y][x] = value;
    return this;
  }

  public setColor(x: number, y: number, color: chalk.ChalkFunction): Surface {
    this.shapes[y][x] = color(this.get(x, y));
    return this;
  }

  static fromString(input: string): Surface {
    return new Surface(input);
  }

  public findStart(): [number, number] {
    for (let y = 0; y < this.height; y++) {
      for (let x = 0; x < this.width; x++) {
        if (this.get(x, y) == 'S') return [x, y];
      }
    }
    throw new Error('Unable to find start position');
  }

  public getStartSymbol(x: number, y: number): string {
    const northOk = this.checkNeighbor(x, y - 1, '|7F');
    const southOk = this.checkNeighbor(x, y + 1, '|JL');
    const eastOk = this.checkNeighbor(x + 1, y, '-7J');
    const westOk = this.checkNeighbor(x - 1, y, '-FL');
    if (northOk && southOk) { return '|'; }
    if (northOk && eastOk) { return 'L'; }
    if (northOk && westOk) { return 'J'; }
    if (southOk && eastOk) { return 'F'; }
    if (southOk && westOk) { return '7'; }
    if (westOk  &&  eastOk) { return '-'; }
    throw new Error('ERROR while getting home symbol');
  }

  // check if a symbol at the given position matches one of the chars in the
  // given string
  checkNeighbor(x: number, y: number, chars: string = ''): boolean {
    if (y < 0 || y < 0 || y >= this.width || y >= this.height) { return false; }
    return chars.includes(this.get(x, y));
  }

  public toString(): string {
    return this.shapes.map(line => line.join('')).join('\n')
      .replace(/F/g, '┌').replace(/L/g, '└').replace(/7/g, '┐')
      .replace(/J/g, '┘').replace(/\./g, ' ').replace(/-/g, '─')
      .replace(/\|/g, '│');
  }

  public distanceMap:Int32Array[] = [];

  public search(startX: number, startY: number): number {
    for (let n = 0; n < this.width; n++) {
        const line = new Int32Array(this.height);
        line.fill(-1);
        this.distanceMap.push(line);
    }
    let distance = -1;

    let futureNodes = [ [startX, startY], ];
    while (true) {
      if (!futureNodes.length) {
        return distance;
      }
      const currentNodes = futureNodes;
      distance += 1;
      futureNodes = [];
      for (const node of currentNodes) {
        const [x, y] = node;
        this.distanceMap[x][y] = distance;
        const symbol = this.get(x, y);
        if (symbol == '.') { continue; }
        if (symbol == '-') { addNode(this, x, y, 'W');  addNode(this, x, y, 'E'); continue; }
        if (symbol == '|') { addNode(this, x, y, 'N'); addNode(this, x, y, 'S'); continue; }
        if (symbol == 'L') { addNode(this, x, y, 'N'); addNode(this, x, y, 'E'); continue; }
        if (symbol == 'J') { addNode(this, x, y, 'N'); addNode(this, x, y, 'W'); continue; }
        if (symbol == 'F') { addNode(this, x, y, 'S'); addNode(this, x, y, 'E'); continue; }
        if (symbol == '7') { addNode(this, x, y, 'S'); addNode(this, x, y, 'W'); continue; }
        throw new Error('unknown symbol "' + symbol + '" at ' +  x + ',' + y);
      }
    }

    function addNode(surface: Surface, x: number, y: number, direction: string): void {
      if (direction == 'N') { y -= 1 }
      if (direction == 'S') { y += 1 }
      if (direction == 'W')  { x -= 1 }
      if (direction == 'E')  { x += 1 }
      if (y < 0 || y < 0 || x >= surface.width || y >= surface.height) { return; }
      if (surface.distanceMap[x][y] != -1) { return }
      surface.distanceMap[x][y] = -2 // reserved
      futureNodes.push([x, y])
    }
  } // search
}


async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();

  const surface = Surface.fromString(rawInput);
  const start = surface.findStart();

  const startSymbol = surface.getStartSymbol(start[0], start[1]);
  surface.set(start[0], start[1], startSymbol);

  const distance = surface.search(start[0], start[1]);
  console.log('startSymbol', startSymbol);
  console.log(surface.toString());
  console.log(distance);
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
