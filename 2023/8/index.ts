import { readFile } from "fs/promises";
import { resolve } from "path";

function parseInput(input: string) {
  return input.split('\n');
}

class Node {
  constructor(public name: string, public left: string, public right: string) {

  }

  getSide(side: string) {
    if (side === 'L') return this.left;
    if (side === 'R') return this.right;
    throw new Error(`Invalid side: ${side}`);
  }
}

class CamelMap {
  public nodes = new Map<string, Node>();
  constructor(public instructions: string[]) {
  }

  public i = 0;
  public getNextInstruction() {
    if (!this.instructions[this.i]) this.i = 0;
    return this.instructions[this.i++];
  }

  public cur?: Node;
  public nextNode() {
    if (!this.cur) {
      this.cur = this.nodes.get('AAA');
    } else {
      const nextInstruction = this.getNextInstruction();
      const nextNodeName = this.cur.getSide(nextInstruction);
      this.cur = this.nodes.get(nextNodeName);
    }
    return this.cur;
  }

  public addNode(name: string, left: string, right:string) {
    this.nodes.set(name, new Node(name, left, right));
  }

  public findNode(name:string) {
    return this.nodes.get(name);
  }
}

async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();
  const input = parseInput(rawInput)

  const instructions = input[0].split('');
  const lines = input.splice(2);

  const map = new CamelMap(instructions);
  lines.forEach(line => {
    map.addNode(
      line.substring(0, 3),
      line.substring(7, 10),
      line.substring(12, 15)
    );
  });

  let node = map.nextNode();
  let count = 0;
  while (node && node.name !== 'ZZZ') {
    count++;
    console.log(node)
    node = map.nextNode();
  }
  console.log('it took %d steps to get to ZZZ', count);
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
