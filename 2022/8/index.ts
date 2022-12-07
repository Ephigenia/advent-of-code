import { readFile } from "fs/promises";
import { resolve } from "path";

class Map {
  public map: number[] = [];

  constructor(public width = 0, public height = 0) {}

  parseFromString(input: string) {
    const lines = input.split(/\n/).map(v => v.trim());
    if (lines.length > 0) {
      this.height = lines.length;
      this.width = lines[0].length;
      this.map = lines.map(line => Array.from(line).map(v => parseInt(v, 10))).flat();
    }
  }

  getSize() {
    return this.width * this.height;
  }

  get(x: number, y: number): number {
    return this.map[y * this.width + x];
  }

  getPos(index: number): [number, number] {
    const y = Math.floor(index / this.width);
    const x = index - y * this.width;
    return [x, y];
  }

  getTrees(x1: number, y1: number, x2: number, y2: number) {
    let trees:number[] = [];
    for(let x = x1; x <= x2; x++) {
      for(let y = y1; y <= y2; y++) {
        trees.push(this.get(x, y));
      }
    }
    return trees;
  }

  isVisibleTree(x: number, y: number): boolean {
    const treeHeight = this.get(x, y);
    const leftVisible = this.getTrees(0, y, x - 1, y)
    const rightVisible = this.getTrees(x + 1, y, this.width - 1, y);
    const topVisible = this.getTrees(x, 0, x, y - 1);
    const bottomVisible = this.getTrees(x, y + 1, x, this.height - 1);
    const isVisible =
      leftVisible.every(h => h < treeHeight) ||
      rightVisible.every(h => h < treeHeight) ||
      topVisible.every(h => h < treeHeight) ||
      bottomVisible.every(h => h < treeHeight);
    return isVisible;
  }

  calculateScenicScore(x: number, y: number): number {
    const treeHeight = this.get(x, y);
    const leftVisible = this.getTrees(0, y, x - 1, y).reverse();
    const rightVisible = this.getTrees(x + 1, y, this.width - 1, y);
    const topVisible = this.getTrees(x, 0, x, y - 1).reverse();
    const bottomVisible = this.getTrees(x, y + 1, x, this.height - 1);
    let topScenic = topVisible.map(v => v < treeHeight).findIndex(v => !v) + 1;
    let bottomScenic = bottomVisible.map(v => v < treeHeight).findIndex(v => !v) + 1;
    let leftScenic = leftVisible.map(v => v < treeHeight).findIndex(v => !v) + 1;
    let rightScenic = rightVisible.map(v => v < treeHeight).findIndex(v => !v) + 1;
    if (topScenic === 0) topScenic = topVisible.length;
    if (bottomScenic === 0) bottomScenic = bottomVisible.length;
    if (leftScenic === 0) leftScenic = leftVisible.length;
    if (rightScenic === 0) rightScenic = rightVisible.length;
    return topScenic * bottomScenic * leftScenic * rightScenic;
  }

  getScenicScores(): number[] {
    let scores:number[] = [];
    for (let y = 0; y < this.height; y++) {
      for (let x = 0; x < this.width; x++) {
        scores.push(this.calculateScenicScore(x, y));
      }
    }
    return scores;
  }

  countVisibleTrees(): number {
    let count = this.width * 2 + this.height * 2 - 4;
    for (let y = 1; y < this.height - 1; y++) {
      for (let x = 1; x < this.width - 1; x++) {
        if (this.isVisibleTree(x, y)) { count++ }
      }
    }
    return count;
  }

  toString() {
    const lines: string[] = [];
    for (let y = 0; y < this.height; y++) {
      lines.push(this.map.slice(y * this.width, (y + 1) * this.width).join(''));
    }
    return lines.join('\n');
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();

  const map = new Map();
  map.parseFromString(raw);

  console.log(map.toString());
  console.log('visible trees', map.countVisibleTrees());

  const scenicScores = map.getScenicScores();
  const max = Math.max(...scenicScores);
  const pos = map.getPos(scenicScores.indexOf(max));
  console.log('Best position has a score of %d (x: %d, y: %d)', max, pos[0], pos[1]);
}

main('input.txt');
// main('inputTest.txt');
