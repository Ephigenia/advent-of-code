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

function calculateVariants(race: Race) {
  const table = [...Array(race.duration).keys()].map((speed) => {
    const travelTime = race.duration - speed;
    const distance = travelTime * speed;
    const success = distance > race.distance;
    if (success) {
      race.successfulVariants++;
    }
    return { speed, travelTime, distance, success };
  });
  return table;
}

function calculateVariants2(race: Race) {
  const table = [...Array(race.duration).keys()].map((speed) => {
    const travelTime = race.duration - speed;
    const distance = travelTime * speed;
    const success = distance > race.distance;
    return success;
  });
  return table;
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

  const variants = races.map(calculateVariants);
  // t.forEach(t => console.table(t));

  console.log(races);
  console.log('multi-variants', races.reduce((acc, race) => acc * race.successfulVariants, 1));
  console.log('');

  // task2
  const [duration, distance] = input.map(line => line.substring(10).replace(/\s+/g, '')).map(v => parseInt(v, 10));
  const newRace = { distance, duration, successfulVariants: 0 };
  const variants2 = calculateVariants2(newRace);
  console.log('Final Race', newRace);
  console.log('multi-variants', variants2.filter(v => v).length);
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
