import { readFile } from "fs/promises";
import { resolve } from "path";

class Node {
  public parent: Directory|undefined;
  public level: number = 0;

  constructor(public name: string, public bytes: number = 0) {}

  getSize():number {
    return this.bytes;
  }

  toString() {
    let indent = '|   '.repeat(this.level - 1);
    if (this.level > 1) {
      indent += '`---';
    } else {
      return `\`-${indent}${this.name} (${this.getSize()})`;
    }
    return `${indent}${this.name} (${this.getSize()})`;
  }
}

class Directory extends Node {
  public nodes:Node[] = [];

  getSize():number {
    if (!this.nodes.length) return this.bytes;
    return this.nodes.reduce((a, node) => node.getSize() + a, 0);
  }

  addNode(node: Node) {
    node.parent = this;
    node.level = this.level + 1;
    this.nodes.push(node);
    return this;
  }

  getDirectories(): Node[] {
    return this.nodes.filter(node => node instanceof Directory);
  }

  getFiles(): Node[] {
    return this.nodes.filter(node => !(node instanceof Directory));
  }

  hasParent():boolean {
    return Boolean(this.parent);
  }

  hasChildren() {
    return this.nodes.length > 0;
  }

  getByName(name: string): Node|Directory|undefined {
    return this.nodes.find(node => (node.name === name));
  }

  toString() {
    let indent = '';
    if (this.level > 1) {
      indent += '|   '.repeat(this.level - 1 || 0);
      indent += '`---';
    } else if (this.name !== '/') {
      indent += '`-';
    }
    return [
      `${indent}${this.name} (dir) ${this.getSize()}`,
      ...this.nodes.map(node => node.toString())
    ].join('\n');
  }
}

function parseInput(input: string):Directory {
  const lines = input.split(/\n/);
  const root = new Directory('/');

  let currentDir: Directory = root;

  lines.forEach((line) => {
    // check different types of line
    const isComment = line.substring(0, 1) === '#';
    if (isComment) return;
    const isCommand = line.substring(0, 1) === '$';
    const isDir = line.match(/^dir\s(.+)/i);
    const isFile = line.match(/(\d+)+\s(.+)/i);

    if (isCommand) {
      let command = line.substring(2);
      const isCd = command.match(/^cd (.+)/);
      if (isCd) {
        const [,dirname] = isCd;
        if (dirname === '/') return;
        if (dirname === '..' && currentDir.parent) {
          currentDir = currentDir.parent;
        }
        let otherDir = currentDir.getByName(dirname);
        if (otherDir instanceof Directory) {
          currentDir = otherDir;
        }
      }
    } else if (isDir) {
      const [,dirname] = isDir;
      currentDir.addNode(new Directory(dirname));
    } else if (isFile) {
      const [,size,name] = isFile;
      const fileNode = new Node(name, parseInt(size, 10));
      currentDir.addNode(fileNode);
    }
  });

  return root;
}

function walk(node: Node, func: (node:Node) => void) {
  func(node);
  if (node instanceof Directory) {
    node.nodes.forEach(node => walk(node, func));
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();

  const root = parseInput(raw);
  console.log(root.toString());

  console.log('\nsize analysis');

  const sizes:number[] = [];
  walk(root, (node) => {
    const size = node.getSize();
    if (node instanceof Directory && size <= 100000) {
      sizes.push(size);
    }
  });

  const sizeSum = sizes.reduce((a, c) => a + c, 0);
  console.log('sizeSum: ', sizeSum);

  const dirsToBeDeleted:number[] = [];
  const spaceNeeded = root.getSize() - 40000000;
  walk(root, (node) => {
    const size = node.getSize();
    if (node instanceof Directory && size > spaceNeeded) {
      dirsToBeDeleted.push(size);
    }
  });
  console.log(Math.min(...dirsToBeDeleted));
}

main('input.txt');
