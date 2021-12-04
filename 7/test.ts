import fs from "fs";

const file = fs.readFileSync("./7/in.txt", "utf-8");

const lines = file.split("\n");

const numbers = lines.shift()?.split(",");

const ex = new RegExp(" +");

const boards: string[][][] = [];
const boardResults: boolean[][][] = [];
while (lines.length > 1) {
  const [, ...board] = lines.splice(0, 6);
  const cleanBoard = board.map((line) => line.trim().split(ex));
  boards.push(cleanBoard);
  boardResults.push([
    [false, false, false, false, false],
    [false, false, false, false, false],
    [false, false, false, false, false],
    [false, false, false, false, false],
    [false, false, false, false, false],
  ]);
}

const computeBoard = (boardNumber: number, number: string) => {
  const board = boards[boardNumber];

  for (let x = 0; x < 5; x++) {
    for (let y = 0; y < 5; y++) {
      if (board[x][y] === number) {
        boardResults[boardNumber][x][y] = true;
      }
    }
  }
};

const isWinningBoard = (boardNumber: number) => {
  const board = boardResults[boardNumber];

  for (let x = 0; x < 5; x++) {
    let win = true;
    for (let y = 0; y < 5 && win; y++) {
      if (board[x][y] === false) {
        win = false;
      }
    }
    if (win) {
      return true;
    }
  }

  for (let y = 0; y < 5; y++) {
    let win = true;
    for (let x = 0; x < 5 && win; x++) {
      if (board[x][y] === false) {
        win = false;
      }
    }
    if (win) {
      return true;
    }
  }
  return false;
};

const calculateScore = (boardNumber: number,number:string) => {
  const board = boards[boardNumber];
  const boardResult = boardResults[boardNumber];

  let sum = 0
  for (let x = 0; x < 5; x++) {
    for (let y = 0; y < 5 ; y++) {
      if(!boardResult[x][y]){
          sum += parseInt(board[x][y])
      }
    }
  }
  return sum * parseInt(number)
};

let isWin = false;
while ((numbers?.length ?? 0 > 0) && !isWin) {
  const number = numbers?.shift();
  if (!number) break;

  for (let i = 0; i < boards.length; i++) {
    computeBoard(i, number);
    isWin = isWinningBoard(i);
    if (isWin) {
      console.log("win: " + i + " num " + number);
      const score = calculateScore(i,number)
      console.log(score)
      break;
    }
  }
}
