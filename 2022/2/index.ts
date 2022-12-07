import { readFile } from "fs/promises";

type ROUND_OUTCOME = 'WIN' | 'LOST' | 'DRAW';
type LEFT_OPTION = 'A' | 'B' | 'C';
type RIGHT_OPTION = 'X' | 'Y' | 'Z';

function sum(arr: number[]) {
    return arr.reduce((a:number, c:number) => (a + c), 0);
}

function parseLineToRound(line: string): [LEFT_OPTION, RIGHT_OPTION]{
    const l = String(line).toUpperCase().trim();
    return [
        l.substring(0, 1) as LEFT_OPTION,
        l.substring(2, 3) as RIGHT_OPTION,
    ];
}

function translateRightOption(right: RIGHT_OPTION): LEFT_OPTION {
    switch(right) {
        case 'X':
            return 'A';
        case 'Y': 
            return 'B';
        case 'Z':
            return 'C';
    }
}

function translateLeftOption(right: LEFT_OPTION): RIGHT_OPTION {
    switch(right) {
        case 'A':
            return 'X';
        case 'B': 
            return 'Y';
        case 'C':
            return 'Z';
    }
}

function getBaseScore(right: RIGHT_OPTION) {
    switch(right) {
        case 'X':
            return 1;
        case 'Y': 
            return 2;
        case 'Z':
            return 3;
    }
}

function getRoundResult(left: LEFT_OPTION, right : RIGHT_OPTION): ROUND_OUTCOME {
    const translatedRight = translateRightOption(right);
    if (left === translatedRight) {
        return 'DRAW';
    }
    // rock + paper
    if (left === 'A' && right === 'Y') return 'WIN';
    // rock + scissors
    if (left === 'A' && right === 'Z') return 'LOST';
    // paper + rock
    if (left === 'B' && right === 'X') return 'LOST';
    // paper + scissors
    if (left === 'B' && right === 'Z') return 'WIN';
    // scissors + rock
    if (left === 'C' && right === 'X') return 'WIN';
    // scissors + paper
    if (left === 'C' && right === 'Y') return 'LOST';
    return 'DRAW';
}

function getRoundScore(left: LEFT_OPTION, right: RIGHT_OPTION): number {
    const outcome = getRoundResult(left, right);
    switch(outcome) {
        case 'WIN':
            return getBaseScore(right) + 6;
        case 'DRAW':
            return getBaseScore(right) + 3;
        case 'LOST':
            return getBaseScore(right) + 0;
    }
}

// A Rock
// B Paper
// C Scissors
// X Rock
// Y Paper
// Z Scissors
async function main() {
    const raw = (await readFile('./input.txt')).toString();
    const rawTest = `A Y
        B X
        C Z`;
    const lines = raw.split(/\n/).filter(v => v);

    // // first part
    // const scores = lines.map<number>((line: string) => {
    //     const [left, right] = parseLineToRound(line);
        
    //     // console.log('ROUND', left, right);
    //     const outcome = getRoundResult(left, right);
    //     const score = getRoundScore(left, right);
    //     console.log('%s %s result %s (%s score)', left, right, outcome, score);
    //     return score;
    // });

    const scores = lines.map<number>((line: string) => {
        const [left, right] = parseLineToRound(line);
        let changedRight = right;
        if (right === 'X') {
            // loose
            if (left === 'A') changedRight = 'Z';
            if (left === 'B') changedRight = 'X';
            if (left === 'C') changedRight = 'Y';
        } else if (right === 'Y') {
            // draw
            changedRight = translateLeftOption(left);
        } else {
            // win
            if (left === 'A') changedRight = 'Y';
            if (left === 'B') changedRight = 'Z';
            if (left === 'C') changedRight = 'X';
        }
        
        // console.log('ROUND', left, right);
        const outcome = getRoundResult(left, changedRight);
        const score = getRoundScore(left, changedRight);
        console.log('%s %s (%s) result %s (%s score)', left, changedRight, right, outcome, score);
        return score;
    });
    
    console.log('%d games played with total score of %d', scores.length, sum(scores));
}

main();