import fs from "fs";
import r, { min } from 'ramda'

const file = fs.readFileSync("./13/in.txt", "utf-8");

const input = file.split(",");
const crabs = input.map((s) => parseInt(s));


// const crabStates = crabs.reduce(
//   (acc:{[key:string]:number}, crab) => {
//     acc[String(crab)] = acc[String(crab)] ? acc[String(crab)] +1 : 1
//     return acc;
//   },
//   {}
// )
// console.log(crabStates)



let minFuel = Number.MAX_SAFE_INTEGER;
for (let i = 0; i < crabs.length; i++) {
  let fuel = 0;
  for (let j = 0; j < crabs.length; j++) {
    // THIS IS WRONG! skips cases where none of the crabs is in the optimal spot
    fuel += Math.abs(crabs[i] - crabs[j]);
    if(fuel>minFuel) break
  }
  if(fuel<minFuel){
    minFuel = fuel;
    console.log("min: ", minFuel);
  }

}

console.log("MIN: ", minFuel);