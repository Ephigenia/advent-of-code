import { readFile } from "fs/promises";
import { resolve } from "path";
import * as crypto from 'crypto';

function findHash(secret: string, numberOfzeros: number): [string, number] {
  const zeros = '0'.repeat(numberOfzeros);
  let hash = '', i = -1;
  while(hash.substring(0, numberOfzeros) !== zeros) {
    i++;
    hash = crypto.createHash('md5').update(String(secret + i)).digest('hex');
  }
  return [hash, i];
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const secret = (await readFile(inputFilename)).toString().trim();

  // example inputs (secrets)
  // const secret = 'abcdef';
  // const secret = 'pqrstuv';
  console.log('Stop guessing the MD5 hash by pressing CTRL+C');
  let [hash, i] = findHash(secret, 5);

  console.log('found the hash starting with 5 zeros %j after trying for %d times', hash, i);

  let [hash2, i2] = findHash(secret, 6);
  console.log('found the hash starting with 6 zeros %j after trying for %d times', hash2, i2);
}

main('input.txt');
