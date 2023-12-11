import { readFile } from "fs/promises";
import { resolve } from "path";


class Surface {
  public shapes: string[][] = [[]];
  public width = 0;
  public height = 0;

  public constructor(lines: string) {
    this.shapes = lines.split('\n').map(line => line.split(''));
    this.width = lines[0].length - 1;
    this.height = lines.length - 1;
  }

  public get(x: number, y: number): string {
    return this.shapes[y][x];
  }

  static fromString(input: string): Surface {
    const surface = new Surface(input);
    return surface;
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
  console.log(surface.toString());
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
