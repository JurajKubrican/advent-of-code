import fs from "fs";
import R from "ramda";

const file = fs.readFileSync("./17/in.txt", "utf-8");

const input = file.split("\n");

const mapStr = input.map((i) => i.split(""));
const map = mapStr.map(i=>i.map(j=>parseInt(j)))

const sizeX = map.length;
const sizeY = map[0].length;

const lowPoints: boolean[][] = [];
for (let i = 0; i <= sizeX; i++) {
  lowPoints.push(Array(sizeY).fill(false));
}

let score = 0
for (let i = 0; i < sizeX; i++) {
  for (let j = 0; j < sizeY; j++) {
    const val = map[i][j];
    let success = true;
    const top = map[i - 1] && map[i - 1][j]
    if (top!==undefined && top <= val) success = false;

    const bottom = map[i + 1] && map[i + 1][j]
    if (bottom!==undefined && bottom <= val) success = false;

    const left = map[i][j-1]
    if (left!==undefined && left <= val) success = false;

    const right = map[i][j+1]
    if (right!==undefined && right <= val) success = false;
  

    if (success) {
      lowPoints[i][j] = true;
      score = score + 1 + val
    }
  }
}

const print = ()=>{
  lowPoints.forEach(row => console.log(row.map(v=>v?"X":'.').join('')));
  map.forEach(row => console.log(row.join('')));
}

print()
console.log(score)