import chalk from 'chalk';
import { readFile } from "fs/promises";
import { resolve } from "path";

const MAP_START = 'S';
const MAP_TARGET = 'E';
enum DIRECTION {
  UP, DOWN, LEFT, RIGHT
}

class ElevationMap {
  private map: string[] = [];
  private width = 0;
  private height = 0;
  public player: [number, number] = [0,0];
  private target: [number, number] = [0,0];
  public moveCount = 0;
  public visited: number[] = [];

  public constructor(map: string) {
    const lines = map.split(/\n/);
    if (lines.length) {
      this.height = lines.length;
      this.width = lines[0].length;
      this.map = Array.from(lines.join(''));
    }
    this.player = this.findFirstChar(MAP_START);
    this.set(this.player[0], this.player[1], 'a');
    this.target = this.findFirstChar(MAP_TARGET);
  }

  private findFirstChar(char: string): [number, number] {
    const index = this.map.indexOf(char);
    if (index < -1) {
      throw new Error(`Unable to find char ${JSON.stringify(char)} anywhere in the map`);
    }
    return this.getCoords(index);
  }

  public getCoords(index: number): [number, number] {
    const y = Math.floor(index / this.width);
    const x = index - y * this.width;
    return [x, y];
  }

  public getIndex(x: number, y: number): number {
    return y * this.width + x;
  }

  public get(x: number, y: number): string {
    return this.map[this.getIndex(x, y)];
  }

  public set(x: number, y: number, char: string):void {
    this.map[this.getIndex(x, y)] = char;
  }

  public isFinished(): boolean {
    return this.get(this.player[0], this.player[1]) === MAP_TARGET;
  }

  public isVisited(x: number, y: number):boolean {
    return this.visited.indexOf(this.getIndex(x, y)) > -1;
  }

  /**
   * @returns normalized elevation, 0-25
   */
  public elevation(char: string): number {
    const charCode = char.charCodeAt(0);
    if (charCode < 97) return 0;
    if (charCode > 122) return 25;
    return 122 - charCode;
  }

  public colorForElevation(elevation: number): chalk.Chalk {
    const ranges = [
      chalk.bgRgb(255,255,255).black,
      chalk.bgRgb(255,255,255).black,
      chalk.bgRgb(100,100,100).black,
      chalk.bgRgb(150,150,150).black,
      chalk.bgRgb(180,180,180).black,
      chalk.bgRgb(128,0,0).black,
      chalk.bgRgb(139,69,19).black,
      chalk.bgRgb(210,105,30).black,
      chalk.bgRgb(218,165,32).black,
      chalk.bgRgb(244,164,96).black,
      chalk.bgRgb(107,142,35).black,
      chalk.bgRgb(50,205,50).black,
      chalk.bgRgb(143,188,143).black,
      chalk.bgRgb(60,179,113).black,
    ];
    let color = ranges.find((color, index) => {
      if (elevation / 25 * (ranges.length - 1) < index) return color;
    });
    if (!color) color = ranges[ranges.length - 1];
    return color;
  }

  public toString(): string {
    let str = '';
    for (let y = 0; y < this.height; y++) {
      for (let x = 0; x < this.width; x++) {
        let color: chalk.Chalk = chalk.gray;
        let char = this.get(x, y);
        if (this.player[0] === x && this.player[1] === y) {
          color = chalk.yellow.bold;
          char = '𓀠';
        } else if (char === MAP_TARGET) {
          color = chalk.red.bold;
        } else if (this.isVisited(x, y)) {
          color = chalk.green;
        } else {
          color = this.colorForElevation(this.elevation(char));
        }
        str += color(char);
      }
      str += '\n';
    }
    return str;
  }

  compareElevations(a: string, b: string): boolean {
    if (a === MAP_START) return true;
    if (b === MAP_START) return false;
    if (b === MAP_TARGET) return true;
    return a.charCodeAt(0) + 1 >= b.charCodeAt(0);
  }

  movePlayer(direction: DIRECTION): void {
    this.moveCount++;
    switch(direction) {
      case DIRECTION.UP:
        if (this.player[1] > 0) this.player[1]--;
        break;
      case DIRECTION.DOWN:
        if (this.player[1] < this.height) this.player[1]++;
        break;
      case DIRECTION.LEFT:
        if (this.player[0] > 0) this.player[0]--;
        break;
      case DIRECTION.RIGHT:
        if (this.player[0] < this.width) this.player[0]++;
        break;
    }
    const index = this.getIndex(this.player[0], this.player[1]);
    this.visited.push(index);
  }

  guessDirection() {
    const dx = this.target[0] - this.player[0]; // positive values target is right
    const dy = this.target[1] - this.player[1]; // positive values target is down
    // try to shorten the longest distance first
    if (Math.abs(dx) > Math.abs(dy)) {
      if (dx < 0) return DIRECTION.LEFT;
      return DIRECTION.RIGHT;
    } else {
      if (dy < 0) return DIRECTION.UP;
      return DIRECTION.DOWN;
    }
  }

  findNextStep(): DIRECTION.UP | DIRECTION.DOWN | DIRECTION.LEFT | DIRECTION.RIGHT {
    const cur = this.get(this.player[0], this.player[1]);
    const upIndex = this.getIndex(this.player[0], this.player[1] - 1);
    const up = this.map[upIndex];
    const downIndex = this.getIndex(this.player[0], this.player[1] + 1);
    const down = this.map[downIndex];
    const rightIndex = this.getIndex(this.player[0] + 1, this.player[1]);
    const right = this.map[rightIndex];
    const leftIndex = this.getIndex(this.player[0] - 1, this.player[1]);
    const left = this.map[leftIndex];

    const bestDirection = this.guessDirection();

    // try different directions "one square up, down, left, or right"
    if (up && this.compareElevations(cur, up) && !this.visited.includes(upIndex)) {
      return DIRECTION.UP;
    } else if (down && this.compareElevations(cur, down) && !this.visited.includes(downIndex)) {
      return DIRECTION.DOWN;
    } else if (left && this.compareElevations(cur, left) && !this.visited.includes(leftIndex)) {
      return DIRECTION.LEFT;
    } else if (right && this.compareElevations(cur, right) && !this.visited.includes(rightIndex)) {
      return DIRECTION.RIGHT;
    }

    return DIRECTION.RIGHT;
  }
}

// use FORCE_COLOR=0 to disable colors
async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();

  const map = new ElevationMap(raw);

  for(let i = 0; i < 4; i++) {
    console.log(map.toString());
    console.log('best direction', map.guessDirection());
    const direction = map.findNextStep();
    map.movePlayer(direction);
    // console.log('#%d moving ', i, direction, map.player);
    if (map.isFinished()) {
      console.log('Found in %d steps', map.moveCount);
      continue;
    }
  }
}

// main('input.txt');
main('input.txt');
