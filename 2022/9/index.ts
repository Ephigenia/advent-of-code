import { readFile } from "fs/promises";
import { resolve } from "path";

enum TILE {
  BODY = '▓',
  DEFAULT = '░',
  HEAD = '█',
  START = 'S',
  TAIL = 'X',
  VISITED = '▒',
}

type DIRECTION = 'U' | 'D' | 'L' | 'R';
class Instruction {
    public direction : DIRECTION;
    public amount : number;
    public constructor(str: string) {
        const p = str.split(/\s+/);
        this.direction = p[0] as DIRECTION;
        this.amount = parseInt(p[1]);
    }

    toString() {
        return this.direction + ' ' + this.amount;
    }
}

class Map {

  public constructor() {}

  public size = [1, 1, 1, 1]
  public cur: [number, number] = [0, 0]; // x, y
  public visited: string[];

  public width () {
    return this.size[2] - this.size[0];
  }

  public height () {
    return this.size[3] - this.size[1];
  }

  public isVisited (x: number, y: number) :boolean {
    return Boolean(this.visited.includes(x + 'x' + y));
  }

  public visit(x: number, y: number): void {
    if (!this.isVisited(x, y)) {
      this.visited.push(x + 'x' + y);
    }
  }

  public isCur(x: number, y: number): boolean {
    return this.cur[0] === x && this.cur[1] === y;
  }

  public move(instruction: Instruction): void {
    for (let i = 0; i < instruction.amount; i++) {
      this.step(instruction.direction);
    }
  }

  public step(direction: DIRECTION): void {
    switch(direction) {
      case 'U':
        this.cur[1]--;
        break;
      case 'D':
        this.cur[1]++;
        break;
      case 'L':
        this.cur[0]--;
        break;
      case 'R':
        this.cur[0]++;
        break;
    }
    const [x, y] = this.cur;
    this.visit(x, y);
    if (this.size[0] <= x) this.size[0] = x;
    if (this.size[2] >= x) this.size[2] = x;
    if (this.size[1] <= y) this.size[1] = y;
    if (this.size[2] >= y) this.size[2] = y;
  }

  public toString(): string {
    let str = '';
    console.log(this.size);
    for (let y = this.size[1]; y < this.size[3]; y++) {
      for (let x = this.size[0]; x < this.size[2]; x++) {
        if (this.isCur(x, y)) {
          str += TILE.HEAD;
        } else if (x === 0 && y === 0) {
          str += TILE.START;
        } else if (this.isVisited(x,y)) {
          str += TILE.VISITED;
        } else {
          str += TILE.DEFAULT;
        }
      }
      str += '\n';
    }
    return str;
  }
}

class Rope {}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();
  const instructions = raw.split(/\n/).map((s) => new Instruction(s));

  const map = new Map();

  // instructions.map(i => map.move(i));

  // map.move(new Instruction('R 3'));
  // map.move(new Instruction('R 1'));
  map.move(new Instruction('U 1'));
  // map.move(new Instruction('L 3'));


  console.log(map.visited);
  console.log(map.toString());
  console.log(map.cur);
}

main("inputTest.txt");
