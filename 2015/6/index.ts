import { readFile } from "fs/promises";
import { resolve } from "path";

type Point = [number, number];
enum COMMAND {
  'ON', 'OFF', 'TOGGLE'
};

class LightGrid {
  public data: Uint8Array;
  constructor(
    private width = 1000,
    private height = 1000,
  ) {
    this.data = Uint8Array.from(Array(width * height).fill(false));
  }

  totalBrightness() {
    return this.data.reduce((a, c) => a + c, 0);
  }

  index(x: number, y: number) {
    if (x >= this.width) x = this.width - 1;
    if (y >= this.height) y = this.height - 1;
    return y * this.height + x;
  }

  get(x: number, y: number): number {
    return this.data[this.index(x, y)];
  }

  set(val: number, x:number, y: number) {
    this.data[this.index(x, y)] = Number(val);
  }

  toggle(x: number, y: number) {
    this.set(this.get(x, y) ? 0 : 1, x, y);
  }

  *indexRect(tl: Point, br: Point) {
    const x1 = Math.min(...[tl[0], br[0]]);
    const x2 = Math.max(...[tl[0], br[0]]);
    const y1 = Math.min(...[tl[1], br[1]]);
    const y2 = Math.max(...[tl[1], br[1]]);
    for (let x = x1; x <= x2; x++) {
      for (let y = y1; y <= y2; y++) {
        yield [x, y] as Point;
      }
    }
  }

  toggleRect(tl: Point, br: Point) {
    for (let point of this.indexRect(tl, br)) {
      this.toggle(point[0], point[1]);
    }
  }

  addBrightness(val: number, x: number, y: number) {
    const brightness = this.get(x, y) + val;
    this.set(brightness > 0 ? brightness : 0, x, y);
  }

  addBrightnessRect(val: number, tl: Point, br: Point) {
    for (let point of this.indexRect(tl, br)) {
      this.addBrightness(val, point[0], point[1]);
    }
  }

  setRect(val: number, tl: Point, br: Point) {
    for (let point of this.indexRect(tl, br)) {
      this.set(val, point[0], point[1]);
    }
  }

  filter(predicate: (value: number, index: number, array: Uint8Array) => any) {
    return this.data.filter(predicate);
  }

  toString() {
    let str = '';
    for(let y = 0; y < this.height; y++) {
      for(let x = 0; x < this.width; x++) {
        let val = this.get(x, y);
        str += val > 10 ? 'X' : String(val);
      }
      str += '\n';
    }
    return str;
  }
}

class Day {
  public grid: LightGrid;

  constructor() {
    this.grid = new LightGrid(1000, 1000);
  }

  parseCommand(line: string): [COMMAND, Point, Point] {
    const matches = line.match(/(turn on|toggle|turn off) (\d+),(\d+) through (\d+),(\d+)/);
    if (!matches) {
      throw new Error(`Unable to parse line ${JSON.stringify(line)}`);
    }
    const tl:Point = [parseInt(matches[2], 10), parseInt(matches[3], 0)];
    const br:Point = [parseInt(matches[4], 10), parseInt(matches[5], 0)];

    let command: COMMAND = COMMAND.OFF;
    switch(matches[1]) {
      case 'turn off':
        command = COMMAND.OFF;
        break;
      case 'turn on':
        command = COMMAND.ON;
        break;
      case 'toggle':
        command = COMMAND.TOGGLE;
        break;
    }
    return [command, tl, br];
  }

  runCommandsDay1(lines: string[]) {
    lines.forEach(line => {
      const command = this.parseCommand(line);
      switch(command[0]) {
        case COMMAND.ON:
          this.grid.setRect(1, command[1], command[2]);
          break;
        case COMMAND.OFF:
          this.grid.setRect(0, command[1], command[2]);
          break;
        case COMMAND.TOGGLE:
          this.grid.toggleRect(command[1], command[2]);
          break;
      }
    });
  }

  runCommandsDay2(lines: string[]) {
    lines.forEach(line => {
      const command = this.parseCommand(line);
      switch(command[0]) {
        case COMMAND.ON:
          this.grid.addBrightnessRect(1, command[1], command[2]);
          break;
        case COMMAND.OFF:
          this.grid.addBrightnessRect(-1, command[1], command[2]);
          break;
        case COMMAND.TOGGLE:
          this.grid.addBrightnessRect(2, command[1], command[2]);
          break;
      }
    });
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const lines = (await readFile(inputFilename)).toString().trim().split(/\n/);

  // const day1 = new Day();
  // day1.runCommandsDay1(lines);
  // console.log('there are %d lights on', day1.grid.filter(v => v > 0).length);

  const day2 = new Day();
  day2.runCommandsDay2(lines);
  console.log('there are %d lights on', day2.grid.filter(v => v > 0).length);
  console.log('with a total brightnewss of %d', day2.grid.totalBrightness());
}

main('input.txt');
// main('inputTest.txt');
