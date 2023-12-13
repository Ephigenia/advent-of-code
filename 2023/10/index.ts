import chalk from "chalk";
import { readFile } from "fs/promises";
import { resolve } from "path";


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

  static fromString(input: string): Surface {
    const surface = new Surface(input);
    return surface;
  }

  public findStart(): [number, number] {
    for (let y = 0; y < this.height; y++) {
      for (let x = 0; x < this.width; x++) {
        if (this.get(x, y) == 'S') return [x, y];
      }
    }
    return [0, 0];
  }

  public setColor(x: number, y: number, color: chalk.ChalkFunction): Surface {
    this.shapes[y][x] = color(this.get(x, y));
    return this;
  }

  public traceRoute(start: [number, number], steps: number = 1) {
    const symbol = this.get(start[0], start[1]);
    this.setColor(start[0], start[1], chalk.red.bold);
    if (symbol === 'S') {
      const next = [start[0], start[1] - 1]
      this.setColor(next[0], next[1], chalk.green);
    }
  }

  public toString(): string {
    return this.shapes.map(row => row.join('')).join('\n')
      .replace(/F/g, '┌')
      .replace(/L/g, '└')
      .replace(/7/g, '┐')
      .replace(/J/g, '┘')
      .replace(/\./g, ' ')
      .replace(/-/g, '─')
      .replace(/\|/g, '│');
  }
}

async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();

  const surface = Surface.fromString(rawInput);
  const start = surface.findStart();
  surface.traceRoute(start, 2);

  console.log(surface.toString());
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
