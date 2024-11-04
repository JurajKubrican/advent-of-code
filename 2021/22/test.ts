import fs from "fs";
import R from "ramda";

const file = fs.readFileSync("./22/in.txt", "utf-8");

const input = file.split("\n");

let octos = input.map((line) => line.split("").map((char) => parseInt(char)));

const addToAll = () => {
  for (let i = 0; i < octos.length; i++) {
    for (let j = 0; j < octos[0].length; j++) {
      octos[i][j] += 1;
    }
  }
};

const maybeFlash = (x: number, y: number) => {
  if(octos[x][y]!==10) return false

  for (let i = x - 1; i <= x + 1; i++) {
    for (let j = y - 1; j <= y + 1; j++) {
      if (
        i < 0 ||
        j < 0 ||
        i >= octos.length ||
        j >= octos[0].length ||
        octos[i][j] === -1 || // has flashed
        octos[i][j] === 10 // already maxed
      ) {
      } else {
        octos[i][j] += 1;
      }
    }
  }
  octos[x][y] = -1;
  return true
};

const cleanUp = () => {
  let light = true
  for (let i = 0; i < octos.length; i++) {
    for (let j = 0; j < octos[0].length; j++) {
      if(octos[i][j]===-1){
        octos[i][j]=0
      }else{
        light=false
      }
    }
  }
  return light
};

console.log(0)
octos.forEach(line=>console.log(line.join('')))

let score = 0
for (let step = 1; step <= 1000; step++) {
  addToAll()

  let flashed = true
  while (flashed ===true){
    flashed=false
    for (let i = 0; i < octos.length; i++) {
      for (let j = 0; j < octos[0].length; j++) {
        const res = maybeFlash(i,j)
        if(res){
          flashed = true
          score ++
        }
      }
    }

  }
  const res = cleanUp()

  console.log(step,score)
  octos.forEach(line=>console.log(line.join('')))

  console.log('');
  if(res){
    console.log('YAAAY',step)

    break
  }
}

console.log(score);