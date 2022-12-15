import { readFile } from "fs/promises";
import { resolve } from "path";

// use FORCE_COLOR=0 to disable colors
async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();

}

main('inputTest.txt');
// main('input.txt');
