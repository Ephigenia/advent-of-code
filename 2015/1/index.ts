import { readFile } from "fs/promises";
import { resolve } from "path";
import { EventEmitter } from "stream";

class Elevator extends EventEmitter {
  static EVENT_MOVE = 'move';
  public moveCount = 0;
  public floor = 0;

  up() {
    this.moveCount++;
    this.floor++;
    this.emit(Elevator.EVENT_MOVE, { floor: this.floor });
  }

  down() {
    this.moveCount++;
    this.floor--;
    this.emit(Elevator.EVENT_MOVE, { floor: this.floor });
  }

  move(instruction: string) {
    switch(instruction) {
      case '(':
        this.up();
        break;
      case ')':
        this.down();
        break;
      default:
        throw new Error(`Invalid elevator instruction: ${JSON.stringify(instruction)}.`)
    }
    return this;
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();

  const elevator = new Elevator();
  const instructions = Array.from(raw);
  elevator.on(Elevator.EVENT_MOVE, function(event) {
    if (event.floor === -1) {
      console.log('arrived in the basement after %d moves', elevator.moveCount);
    }
  });
  instructions.forEach(instruction => elevator.move(instruction));

  console.log(
    'Elevator arrived at %d floor after %d instructions',
    elevator.floor,
    elevator.moveCount,
  );
}

main('input.txt');
