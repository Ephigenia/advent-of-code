import { readFile } from "fs/promises";
import { resolve } from "path";

type DIRECTION_STRING = 'v' | '<' | '>' | '^';
enum DIRECTION {
    UP = '^', DOWN = 'v', LEFT = '<', RIGHT = '>',
}

class Player {
  public pos: [number, number] = [0, 0];

  constructor(public world: World, public name: string) {}

  move(direction: DIRECTION_STRING) {
    switch(direction) {
      case DIRECTION.UP:
        this.pos[1]--;
        break;
      case DIRECTION.DOWN:
        this.pos[1]++;
        break;
      case DIRECTION.LEFT:
        this.pos[0]--;
        break;
      case DIRECTION.RIGHT:
        this.pos[0]++;
        break;
    };
    this.world.visit(this.pos[0], this.pos[1]);
  }

  toString() {
    return `Player ${this.name} ${this.pos[0]}x${this.pos[1]}`;
  }
}

class World {
  public visited: Map<string, number> = new Map();
                   // x1      y1      x2      y2
  public borders: [number, number, number, number] = [0,0,0,0];

  public players = new Set<Player>();

  createPlayer(name: string) {
    const player = new Player(this, name);
    this.players.add(player);
    this.visit(0, 0);
    return player;
  }

  visit(x: number, y: number): void {
    const key = [x, y].join(',');
    const val = this.getVisitsAt(x, y) + 1;
    this.visited.set(key, val);
    this.extendBorders(x, y);
  }

  private extendBorders(x: number, y: number): void {
    // x
    if (this.borders[0] > x) this.borders[0] = x;
    if (this.borders[2] < x) this.borders[2] = x;
    // y
    if (this.borders[1] > y) this.borders[1] = y;
    if (this.borders[3] < y) this.borders[3] = y;
  }

  getVisitsAt(x: number, y: number): number {
    const key = [x,y].join(',');
    return this.visited.get(key) || 0;
  }

  render() {
    let str = '';
    for (let y = this.borders[1]; y <= this.borders[3]; y++) {
      for (let x = this.borders[0]; x <= this.borders[2]; x++) {
        const visits = this.getVisitsAt(x, y);
        let char = '.';
        if (visits > 8) {
          char = '▓';
        } else if (visits > 4) {
          char = '▒';
        } else if (visits > 0) {
          char = '░';
        }
        str += char;
      }
      str += '\n';
    }
    return str;
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = Array.from((await readFile(inputFilename)).toString().trim());

  const world = new World();
  const santa = world.createPlayer('santa');
  const robosanta = world.createPlayer('robo');
  raw.forEach((directionString, i) => {
    if (i % 2) {
      santa.move(directionString as DIRECTION_STRING);
    } else {
      robosanta.move(directionString as DIRECTION_STRING);
    }
  });

  console.log('total houses visited', world.visited.size);
  console.log(world.borders);
  // console.log(world.visited.size);
  console.log(world.render());

  const onePresentVisits = [...world.visited].filter(([key, val]) => val >= 1);
  console.log('number of houses with at least one present: %d', onePresentVisits.length);
}

main('input.txt');
