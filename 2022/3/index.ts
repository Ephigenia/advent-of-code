import { readFile } from "fs/promises";
import { resolve } from "path";

function sum(arr: number[]) {
  return arr.reduce((a:number, c:number) => (a + c), 0);
}

function intersect<T>(arr: Array<T>, arr2: Array<T>): Array<T> {
  return arr.filter(v => arr2.includes(v));
}

function intersectA<T>(arr: Array<Array<T>>) {
  return arr.reduce((intersection, items) => {
    return intersect(intersection, items);
  }, arr[0]);
}

class Compartment {

  get length() {
    return this.items.length;
  }

  constructor(public items: Item[] = []) {
  };

  hasItem(code: string|Item): boolean {
    if (code instanceof Item) code = code.code;
    return Boolean(this.findItem(code));
  }

  findItem(code: string|Item): Item | undefined {
    if (code instanceof Item) code = code.code;
    return this.items.find(item => item.code === code);
  }
}

class Item {
  priority: number;
  code: string;
  constructor(code: string) {
    this.code = code;
    this.priority = this.getPriority(code);
  }

  getPriority(itemCode: string) {
    // A = 65 Z = 90
    // a = 97 z = 122
    const charCode = itemCode.charCodeAt(0);
    if (charCode >= 65 && charCode <= 90) {
      return charCode - 65 + 27;
    } else if (charCode >= 97 && charCode <= 122) {
      return charCode - 97 + 1;
    } else {
      return 0;
    }
  }
}
class Rucksack {
  left = new Compartment;
  right = new Compartment;

  constructor(str: string) {
    this.fromString(str);
  }

  fromString(str: string) {
    const len = str.length
    this.left = new Compartment(Array.from(str.substring(0, len / 2)).map(code => new Item(code)));
    this.right = new Compartment(Array.from(str.substring(len / 2, len)).map(code => new Item(code)));
  }

  getItems() {
    return [...this.left.items, ...this.right.items];
  }

  hasItem(code: Item|string) {
    return this.left.hasItem(code) || this.right.hasItem(code);
  }

  findSameItem(): Item|undefined {
    let sameItem:Item|undefined = undefined;
    let i = 0;
    do {
      sameItem = this.right.items.find(rightItem => rightItem.code === this.left.items[i].code);
      i++;
    } while (!sameItem && i < this.left.length)
    return sameItem;
  }
}

class Elv {
  constructor(public rucksack: Rucksack) {}
}

async function main() {
  const inputFilename = resolve(__dirname, 'input1.txt');
  const inputFilenameTest = resolve(__dirname, 'inputTest.txt');
  const raw = (await readFile(inputFilename)).toString().trim();
  const lines = raw.split(/\n/).map(l => l.trim());

  const elves = lines.map((rucksackString: string) => {
    const rucksack = new Rucksack(rucksackString);
    return new Elv(rucksack);
  })

  const prios: number[] = elves
    .map((elve, i) => {
      const sameItem = elve.rucksack.findSameItem();
      console.log('elf %d has %d items in his rucksack, %d has prio %d',
        i, elve.rucksack.left.length, sameItem?.code, sameItem?.priority
      );
      return sameItem?.priority || 0;
    });
  console.log('%d elves, prio sum: %d', elves.length, sum(prios));

  // identify badges, same item in 3 subsequent groups
  let group:Elv[] = [];
  const elveGroups = elves.reduce<Elv[][]>((groups, elv, i) => {
    group.push(elv);
    if ((i + 1) % 3 === 0) {
      groups.push(group);
      group = [];
    }
    return groups;
  }, []);

  console.log('found %d groups of elves', elveGroups.length);


  const prios2 = elveGroups.map((group) => {
    const itemGroups = group.map(elv => elv.rucksack.getItems().map(i => i.code));

    const item = new Item(intersectA(itemGroups)[0]);
    return item.priority;
  });

  console.log('badge sum is %d', sum(prios2));

}

main();
