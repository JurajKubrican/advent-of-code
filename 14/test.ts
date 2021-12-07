import fs from "fs";
import R  from "ramda";

const file = fs.readFileSync("./14/in.txt", "utf-8");

const input = file.split(",");
const crabs = input.map((s) => parseInt(s));

const maxCrab = R.apply(Math.max,crabs)
const minCrab = R.apply(Math.min,crabs)


const f = (len:number)=>{

  let fuel = 0
  for(let i = 0; i<=len;i++){
    fuel+=i
  }
  return fuel
}
const fitness = R.memoizeWith(R.identity as any,f)


let minFuel = Number.MAX_SAFE_INTEGER;
for (let i = minCrab; i < maxCrab; i++) {
  let fuel = 0;
  for (let j = 0; j < crabs.length; j++) {
    fuel += fitness(Math.abs(i-crabs[j]));
    if(fuel>mingit addFuel) break
  }

  if(fuel<minFuel){
    minFuel = fuel;
    console.log("min: ", minFuel);
  }

}

console.log("MIN: ", minFuel);