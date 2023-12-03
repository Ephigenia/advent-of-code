import { readFile } from "fs/promises";
import { resolve } from "path";

// games
//    n
//    draws[] = [red, green, blue]

// red green blue
// game number
// number of stones in each color per game

interface Draw { r: number, g: number, b: number;};
interface Game { id: number, draws: Array<Draw>; };

function parseDraw(draw: string): Draw {
  const c = { r: 0, g: 0, b: 0 };
  draw.split(/,\s*/).map(numberAndColor => {
    const [count, color] = numberAndColor.split(/\s+/);
    const key = color.substring(0, 1) as keyof typeof c;
    c[key] = parseInt(count);
    return c;
  });
  return c;
}

function parseDraws(drawString: string): Draw[] {
  return drawString.split(/;\s*/).map(parseDraw);
}

function findGameMaxRGB(draws: Array<Draw>) {
  const r = draws.map(d => d.r);
  const g = draws.map(d => d.g);
  const b = draws.map(d => d.b);
  return {
    r: Math.max(...r),
    g: Math.max(...g),
    b: Math.max(...b),
  };
}

async function main(filename: string) {
  const inputFilename = resolve(filename);
  const rawInput = (await readFile(inputFilename)).toString().trim();

  const games: Game[] = rawInput
    .split('\n')
    .map(v => v.trim().substring(v.indexOf(':') + 2))
    .map((str, i) => ({
      id: i + 1,
      draws: parseDraws(str)
    }));

  const bagRGB = { r: 12, g: 13, b: 14 };
  const possibleGames = games.filter((game, i) => {
    const maxRGB = findGameMaxRGB(game.draws);

    const isPossible = {
      r: maxRGB.r <= bagRGB.r,
      g: maxRGB.g <= bagRGB.g,
      b: maxRGB.b <= bagRGB.b
    }
    const isPossibleC = isPossible.r && isPossible.g && isPossible.b;
    console.log('Game %d is %s possible', game.id, isPossibleC ? '' : 'not');

    return isPossibleC;
  });

  const sum = possibleGames.reduce((acc, game) => game.id + acc, 0);

  console.log('number of possible games', possibleGames.length);
  console.log('sum of ids', sum);
}

const INPUT_FILENAME = process.argv.pop() || 'input.txt';
main(INPUT_FILENAME);
