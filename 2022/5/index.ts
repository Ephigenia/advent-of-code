import { readFile } from "fs/promises";
import { resolve } from "path";


interface Instruction {
  from: number;
  to: number;
  amount: number;
}


class Crate {
  constructor(public code: string) {}

  toString() {
    return `[${this.code}]`;
  }
}


class CrateStack {
  constructor(public crates: Crate[]) {}

  length() {
    return this.crates.length;
  }

  toString() {
    return this.crates.map(crate => crate.toString()).join(' ');
  }

  pop() {
    return this.crates.shift();
  }

  add(crate: Crate) {
    this.crates = [crate, ...this.crates];
  }

  first() {
    return this.crates[0];
  }
}


class CargoRoom {
  public stacks: CrateStack[] = [];

  constructor(stacks: string[][]) {
    this.stacks = stacks.map((stackItems) => {
      return new CrateStack(stackItems.map(itemCode => new Crate(itemCode)));
    });
  }

  public move(from: number, to: number, amount: number):void {
    for (let i = 0; i < amount; i++) {
      const item = this.stacks[from].pop();
      if (item) {
        this.stacks[to].add(item);
      }
    }
  }

  static parseInstructions(instructions: string): Array<Instruction|null> {
    const lines = instructions.split(/\n/);
    const validLines = lines
      .map((instructionString) => instructionString.match(/move (\d+) from (\d+) to (\d+)/i))
      .filter(v => v);
    return validLines.map((match) => {
      if (!match) return null;
      return {
        amount: parseInt(match[1], 10),
        from: parseInt(match[2], 10) - 1,
        to: parseInt(match[3], 10) - 1,
      }
    });
  }

  static parseFromMap(mapString: string) {
    const lines = mapString.split(/\n/);
    const count = (lines[lines.length - 1].trim().split(/\s+/)).length;

    const stacks:string[][] = [];
    for(let x = 0; x < count ; x++) {
      if (!stacks[x]) stacks[x] = [];
      for (let y = 0; y < lines.length - 1; y++) {
        const pos = (x * 4) + 1;
        if (pos < lines[y].length) {
          const itemCode = lines[y].substring((x * 4) + 1, (x * 4) + 2).trim();
          if (itemCode) {
            stacks[x].push(itemCode);
          }
        }
      }
    }

    return stacks;
  }

  toString() {
    return this.stacks.map((stack, i) => `${i} ${stack.toString()}`).join('\n');
  }
}


class CargoRoom2 extends CargoRoom {
  public move(from: number, to: number, amount: number):void {
    const movingItems:Crate[] = [];
    for (let i = 0; i < amount; i++) {
      const item = this.stacks[from].pop();
      if (item) {
        movingItems.push(item);
      }
    }
    // console.log(movingItems);
    movingItems.reverse().forEach(item => {
      this.stacks[to].add(item);
    })
  }
}

async function main() {
  const inputFilename = resolve(__dirname, 'input1.txt');
  const inputFilenameTest = resolve(__dirname, 'inputTest.txt');
  const raw = (await readFile(inputFilename)).toString();

  const [rawStartMap, rawInstructions] = raw.split(/\n\n/m);
  const stacks = CargoRoom.parseFromMap(rawStartMap);
  const instructions = CargoRoom.parseInstructions(rawInstructions);
  const room = new CargoRoom(stacks);

  console.log(`${room.toString()}`);
  instructions.forEach(instruction => {
    if (!instruction) return;
    room.move(instruction?.from, instruction?.to, instruction?.amount);
  });
  console.log(`------RESULT------`);
  console.log(`${room.toString()}`);
  console.log('Result is: ' + room.stacks.map(stack => stack.first().code).join(''));

  // second task
  console.log();
  console.log('USING CRANE 2');
  const room2 = new CargoRoom2(stacks);
  instructions.forEach((instruction,i) => {
    if (!instruction) return;
    room2.move(instruction?.from, instruction?.to, instruction?.amount);
  });
  console.log(`------RESULT------`);
  console.log(`${room2.toString()}`);
  console.log('Result is: ' + room2.stacks.map(stack => stack.first().code).join(''));
}

main();
