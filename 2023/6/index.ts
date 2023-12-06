import { readFile } from "fs/promises";
import { resolve } from "path";

function parseInput(input: string) {
  return input.split('\n');
}

interface Race {
  duration: number;
  distance: number;
  successfulVariants: number;
}

async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();
  const input = parseInput(rawInput)

  const [durations, distances] = input.map(line => line.split(/\s+/g).slice(1).map(v => parseInt(v, 10)))
  const races = durations.map((duration, i) => ({
    duration,
    distance: distances[i],
    successfulVariants: 0,
  }));

  races.map((race) => {
    const table = [...Array(race.duration).keys()].map((j) => {
      const waitTime = j;
      const speed = j;
      const travelTime = race.duration - waitTime;
      const distance = travelTime * speed;
      const success = distance > race.distance;
      if (success) {
        race.successfulVariants++;
      }
      return { waitTime, speed, travelTime, distance, success};
    });
    console.table(table);
  });

  console.log(races);
  console.log('multi-variants', races.reduce((acc, race) => acc * race.successfulVariants, 1));
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
