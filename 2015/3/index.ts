import { readFile } from "fs/promises";
import { resolve } from "path";

type DIRECTION_STRING = 'v' | '<' | '>' | '^';
enum DIRECTION {
    UP = '^', DOWN = 'v', LEFT = '<', RIGHT = '>',
}

class Player {
  public visited: Map<string, number> = new Map();
  public pos: [number, number] = [0, 0];
  public moves: DIRECTION_STRING[] = [];
                  // x1      y1      x2      y2
  public borders: [number, number, number, number] = [0,0,0,0];

  constructor(public name: string) {}

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
    this.moves.push(direction);
    this.visit(this.pos[0], this.pos[1]);
  }

  visit(x: number, y: number): void {
    const key = [x,y].join(',');
    this.visited.set(key, this.getVisitsAt(x, y) + 1);
    this.extendBorders();
  }

  private extendBorders(): void {
    // X
    if (this.borders[0] > this.pos[0]) this.borders[0] = this.pos[0];
    if (this.borders[2] < this.pos[0]) this.borders[2] = this.pos[0];
    // Y
    if (this.borders[1] > this.pos[1]) this.borders[1] = this.pos[1];
    if (this.borders[3] < this.pos[1]) this.borders[3] = this.pos[1];
  }

  getVisitsAt(x: number, y: number): number {
    const key = [x,y].join(',');
    return this.visited.get(key) || 0;
  }

  toString() {
    return `${this.name} ${this.pos[0]}x${this.pos[1]}`;
  }
}

class World {

  players = new Set<Player>();

  createPlayer(name: string) {
    const player = new Player(name)
    this.players.add(player);
    return player;
  }

  getBorders() {
    const players = [...this.players];
    return [
      Math.min(...players.map(p => p.borders[0])),
      Math.min(...players.map(p => p.borders[1])),
      Math.max(...players.map(p => p.borders[2])),
      Math.max(...players.map(p => p.borders[3])),
    ];
  }

  render() {
    const borders = this.getBorders();
    let str = '';
    for(let y = borders[1]; y < borders[3]; y++) {
      for(let x = borders[0]; x < borders[2]; x++) {
        // const visits = this.getVisitsAt(x, y);
        let char = '░';
        // if (visits > 5) {
        //   char = '▓';
        // } else if (visits > 0) {
        //   char = '▒';
        // }
        str += char;
      }
      str += '\n';
    }
    return str;
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = Array.from((await readFile(inputFilename)).toString().trim()).slice(0, 200);

  // const raw = Array.from('^v');
  // const raw = Array.from('^>v<<');

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

  console.log(world.players);
  console.log(world.getBorders());

  // console.log(world.visited.size);
  console.log(world.render());

  // const onePresentVisits = [...world.visited].filter(([key, val]) => val >= 1);
  // console.log('number of houses with at least one present: %d', onePresentVisits.length);
}

main('input.txt');
