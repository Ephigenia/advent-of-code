import chalk from "chalk";
import { readFile } from "fs/promises";
import { resolve } from "path";

// valve
// flow rate (pressure per minute)
// pipe
// tunnels

enum VALVE_STATE {
  OPENED, CLOSED,
}

class Vulcan {
  public current: Valve;
  public releasedPressure = 0;
  public minutes = 1;

  public constructor(public valves: Valve[] = []) {
    this.current = this.getValveByLabel('AA');
  }

  public getOpenValves() {
    return this.valves.filter(valve => valve.isOpen())
  }

  public getFlowRate() {
    return this.getOpenValves().reduce((pressure, valve) => {
      return pressure + valve.flowRate;
    }, 0);
  }

  public move(to: string) {
    if (this.current && this.current.label === to) return;
    const targetValve = this.getValveByLabel(to);
    console.log(`You move to ${to}`);
    this.current = targetValve;
    this.step();
  }

  public getValveByLabel(label: string): Valve {
    const valve = this.valves.find(valve => valve.label === label);
    if (!valve) {
      throw new Error(`Unable to move to non-existent valve with label ${JSON.stringify(label)}`);
    }
    return valve;
  }

  public findNextValve() {
    const valves = this.current.tunnels.map(label => this.getValveByLabel(label));
    const closedValves = valves.filter(valve => valve.isClosed());
    return closedValves.sort((a, b)  => a.flowRate - b.flowRate).pop()
  }

  public open() {
    if (this.current.isOpen()) return;
    console.log(`You open valve ${this.current.label}`);
    this.current.open();
    this.step();
  }

  public step() {
    this.minutes++;
    this.releasedPressure += this.getFlowRate();
  }

  public toString() {
    return [
      'Valve ' + this.getOpenValves().map(valve => valve.label).join(', ') +
        ' are open, releasing ' + chalk.bold.white(this.getFlowRate()) + ' pressure' +
        ` released ${chalk.bold.white(this.releasedPressure)} so far`,
      this.current ? 'current pos: ' + this.current.label : '',
    ].join('\n');
  }
}

class Valve {
  public state = VALVE_STATE.CLOSED;
  constructor(
    public label: string,
    public flowRate: number,
    public tunnels: string[] = [],
  ) {}

  static PARSE_REGEXP = /Valve ([a-z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.+)/i

  static createFromDefinition(definition: string) {
    const matches = definition.match(Valve.PARSE_REGEXP);
    if (!matches) {
      throw new Error(`Unable to parse Valve definition ${JSON.stringify(definition)}`);
    }
    return new Valve(
      matches[1], parseInt(matches[2], 10), matches[3].split(',').map(v => v.trim()),
    );
  }

  open() {
    this.state = VALVE_STATE.OPENED;
  }

  close() {
    this.state = VALVE_STATE.CLOSED;
  }

  isOpen() { return this.state === VALVE_STATE.OPENED; };
  isClosed() { return !this.isOpen(); };

  toString() {
    return `Valve ${this.label} has flow rate=${this.flowRate}, tunnels lead to valves ${this.tunnels.join(',  ')}`;
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();
  const valveDefinitions = raw.split(/\n/);

  const valves = valveDefinitions.map(
  definition => Valve.createFromDefinition(definition));
  valves.forEach((valve) => console.log(valve.toString()));

  const vulcan = new Vulcan(valves);
  let nextValve: Valve|undefined;
  while (nextValve = vulcan.findNextValve()) {
    // console.log(`== Minute ${vulcan.minutes} ==`);
    console.log(vulcan.toString());
    vulcan.move(nextValve.label);
    // console.log(`== Minute ${vulcan.minutes} ==`);
    vulcan.open();
    console.log(vulcan.toString());
  }
  console.log(vulcan.findNextValve());
}

main('inputTest.txt');
// main('input.txt');
