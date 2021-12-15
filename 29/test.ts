import fs from "fs";
import R from "ramda";
import Heap from "heap";

const file = fs.readFileSync("./29/in.txt", "utf-8");

const input = file.split("\n");

const field = input.map((line) => line.split("").map((c) => parseInt(c)));

const sizeX = field.length;
const sizeY = field[0].length;
const fitness: number[][] = [];
for (let i = 0; i < sizeX; i++) {
  fitness.push(Array(sizeY).fill(-1));
}

type Point = { x: number; y: number };
fitness[0][0] = 0;

type HeapPoint = Point & { fit: number };
var heap = new Heap<HeapPoint>((a: HeapPoint, b: HeapPoint) => {
  return a.fit - b.fit;
});
heap.push({ x: 0, y: 0, fit: 0 });

const guard = (a?: number) => (a && a > -1 ? a : Number.MAX_SAFE_INTEGER);

while (fitness[sizeX - 1][sizeY - 1] === -1) {
  const {x,y} = heap.pop();
  const fit = fitness[x][y]
  if(fitness[x+1]?.[y]!==undefined && (fitness[x+1][y]===-1||fitness[x+1][y] > fit+ field[x+1][y])){
    fitness[x+1][y] = fit+ field[x+1][y]
    heap.push({x:x+1,y:y,fit: fitness[x+1][y]})
  }
  if(fitness[x-1]?.[y]!==undefined && (fitness[x-1][y]===-1||fitness[x-1][y] > fit+ field[x-1][y])){
    fitness[x-1][y] = fit+ field[x-1][y]
    heap.push({x:x-1,y:y,fit: fitness[x-1][y]})
  }
  if(fitness[x]?.[y+1]!==undefined && (fitness[x][y+1]===-1||fitness[x][y+1] > fit+ field[x][y+1])){
    fitness[x][y+1] = fit+ field[x][y+1]
    heap.push({x:x,y:y+1,fit: fitness[x][y+1]})
  }
  if(fitness[x]?.[y-1]!==undefined && (fitness[x][y-1]===-1||fitness[x][y-1] > fit+ field[x][y-1])){
    fitness[x][y-1] = fit+ field[x][y-1]
    heap.push({x:x,y:y-1,fit: fitness[x][y-1]})
  }
  
  
  // fitness.forEach((line) => console.log(line.join(" ")));
  console.log(fit);
}


fitness.forEach((line) => console.log(line.join(" ")));
console.log(fitness[sizeX - 1][sizeY - 1])