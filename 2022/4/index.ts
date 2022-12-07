import { readFile } from "fs/promises";
import { resolve } from "path";

function sum(arr: number[]) {
  return arr.reduce((a:number, c:number) => (a + c), 0);
}

async function main() {
  const inputFilename = resolve(__dirname, 'input1.txt');
  const inputFilenameTest = resolve(__dirname, 'inputTest.txt');
  const raw = (await readFile(inputFilename)).toString().trim();
  const lines = raw.split(/\n/).map(l => l.trim());

  const parsedPairs = lines.map(assignmentPlan => {
    const [left, right] = assignmentPlan.split(',');

    const l = left.split(/-/).map(v => parseInt(v, 10));
    const r = right.split(/-/).map(v => parseInt(v, 10));
    return [l, r];
  });

  const containingPairs = parsedPairs.map(([l, r]) => {
    let overlapping = (
         (l[0] <= r[0] && l[1] >= r[1]) // 2-8 && 3-7
      || (l[0] >= r[0] && l[1] <= r[1]) // 3-7 && 2-8
    );
    console.log(l, r, overlapping);
    return Number(overlapping);
  });

  console.log('%d elve-pairs fully contain %d', lines.length, sum(containingPairs));

  const containingPairs2 = parsedPairs.map(([l, r]) => {
    let overlapping = !(
         (l[0] < r[0] && l[1] < r[0])
      || (l[0] > r[1] && l[1] > r[1])
    );
    console.log(l, r, overlapping);
    return Number(overlapping);
  });
  console.log('%d elve-pairs fully contain %d', lines.length, sum(containingPairs2));
  // part 2
}

main();
