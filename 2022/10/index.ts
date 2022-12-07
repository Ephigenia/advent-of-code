import { readFile } from "fs/promises";
import { resolve } from "path";

class CPU {
  public x = 1;

  public stack: string[] = [];
  public nextCommand = 1;
  public command: string | undefined;

  public cycle = 0;

  add(command: string) {
    this.stack.push(command);
  }

  public values: number[] = [];

  public crt: string[] = [];

  printCrt() {
    const screen: string[] = [];
    while (this.crt.length) {
      screen.push(this.crt.splice(0, 40).join(""));
    }
    return screen.join("\n");
  }

  step() {
    if (!this.command) {
      this.command = this.stack.shift();
    }
    if (!this.command) {
      // stack empty, stop everythng
      return false;
    }

    if (Math.abs(this.x - (this.cycle % 40)) < 2) {
      this.crt.push("▓");
    } else {
      this.crt.push("░");
    }
    this.cycle++;

    if ([20, 60, 100, 140, 180, 220].includes(this.cycle)) {
      this.values.push(this.signalStrength());
    }

    if (this.cycle > this.nextCommand) {
      this.nextCommand += 1;
      if (this.command?.substring(0, 4) === "addx") {
        const value = parseInt(
          this.command.substring(5, this.command.length),
          10
        );
        this.x += value;
        this.nextCommand += 1;
      }
      delete this.command;
    }

    return true;
  }

  signalStrength() {
    return this.cycle * this.x;
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();
  const lines = raw.split(/\n/);

  const cpu = new CPU();
  lines.forEach((instruction: string) => cpu.add(instruction));

  const interval = setInterval(() => {
    if (!cpu.step()) {
      clearInterval(interval);
      console.log(
        "signalStrength",
        cpu.values.reduce((a, c) => a + c, 0)
      );

      console.log(cpu.printCrt());
      process.exit(0);
    }
  }, 0);
}

main("input.txt");
