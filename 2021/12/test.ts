import fs from "fs";

const file = fs.readFileSync("./12/in.txt", "utf-8");

const input = file.split(",");
const fish = input.map((s) => parseInt(s));

const fishStates = fish.reduce(
  (acc, fish) => {
    acc[fish]++
    return acc;
  },
  [0, 0, 0, 0, 0, 0, 0, 0, 0]
);

const simulate = () => {
  const newFishCount = fishStates.shift()??0
  fishStates[8] = newFishCount
  fishStates[6] += newFishCount
};

for (let cycle = 1; cycle <= 256; cycle++) {
  simulate();
  console.log(cycle , fishStates.reduce((acc,val)=>acc+val,0));
}
