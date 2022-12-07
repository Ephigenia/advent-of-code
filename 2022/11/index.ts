import { readFile } from "fs/promises";
import { resolve } from "path";

class Monkey {
  constructor(
    public itemWorryLevels: number[],
    public operation: (old: number) => number,
    public divisor: number,
    public trueTarget: number,
    public falseTarget: number,
  ) {}

  public testCount = 0;

  worryLevelTotal(): number {
    return this.itemWorryLevels.reduce((a, c) => a + c, 0);
  }

  addItem(worryLevel: number) {
    this.itemWorryLevels.push(worryLevel);
  }

  toString(): string {
    return this.itemWorryLevels.join(', ');
  }

  test(worryLevel: number): boolean {
    this.testCount++;
    return worryLevel % this.divisor === 0;
  }

  targetMonkey(worryLevel: number) {
    if (this.test(worryLevel)) return this.trueTarget;
    return this.falseTarget;
  }

  act(mod: number) {
    this.itemWorryLevels = this.itemWorryLevels
      .map((worryLevel) => {
        return this.operation(worryLevel);
      })
      .map(worryLevel => worryLevel % mod);
  }

  static createOperation(
    operationString: string
  ): (old: number) => number {
    const [operator, operand] = operationString.split(/\s/);
    const getOperand = (old: number | string): number => {
      if (operand === 'old') return old as number;
      return parseInt(operand, 10);
    }
    if (operator === '*') {
      return (old: number) => old * getOperand(old);
    } else {
      return (old: number) => old + getOperand(old);
    }
  }

  static parseFromInput(input: string) {
    const lines = input.split(/\n/);
    const worryLevels = lines[1].substring(18).split(',').map(v => parseInt(v, 10));
    const operation = Monkey.createOperation(lines[2].substring(23));
    const divisor = parseInt(lines[3].substring(21), 10);
    const trueTargetMonkey = parseInt(lines[4].substring(29), 10);
    const falseTargetMonkey = parseInt(lines[5].substring(30), 10);
    return new Monkey(worryLevels, operation, divisor, trueTargetMonkey, falseTargetMonkey);
  }
}

class MonkeyGroup {
  public monkeys: Monkey[] = [];
  addMonkey(monkey: Monkey) {
    this.monkeys.push(monkey);
  }

  getMonkey(index: number):Monkey {
    return this.monkeys[index];
  }

  toString(): string {
    return this.monkeys.map((monkey, i) => `Monkey ${i}: ${monkey.toString()}`).join('\n');
  }

  playRound(): void {
    const mod = this.monkeys.reduce((a, m) => a * m.divisor, 1);
    this.monkeys.forEach((monkey, i) => {
      monkey.act(mod);
      let worryLevel: number|undefined;
      while (worryLevel = monkey.itemWorryLevels.shift()) {
        // console.log('monkey %d worry level %d', i, worryLevel);
        const targetMonkeyIndex = monkey.targetMonkey(worryLevel);
        this.getMonkey(targetMonkeyIndex).addItem(worryLevel);
      }
    });
  }

  mostActiveMonkeys() {
    return this.monkeys.sort((a, b) => b.testCount - a.testCount);
  }

  monkeyBusiness(num: number): number {
    const mostActive = this.mostActiveMonkeys().splice(0, num);
    return mostActive.reduce((a, c) => a * c.testCount, 1);
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString();

  const group = new MonkeyGroup();

  const rawMonkeys = raw.split(/\n\n/);
  rawMonkeys.forEach(rawMonkeyInput => {
    const monkey = Monkey.parseFromInput(rawMonkeyInput);
    group.addMonkey(monkey);
  })

  console.log(group.toString() + '\n\n');
  for (let i = 1; i <= 10000; i++) {
    group.playRound();
    if (i % 1000 === 0 || i === 20 || i === 1) {
      console.log('== After round %d ==', i);
      console.log(group.toString());
      group.monkeys.forEach((monkey, i) => {
        console.log('Monkey %d inspected items %d times', i, monkey.testCount);
      });
    }
  }

  console.log(group.monkeyBusiness(2));
}

main('input.txt');
// main("inputTest1.txt");
