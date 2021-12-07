import fs from "fs";

const file = fs.readFileSync("./11/in.txt", "utf-8");

const input = file.split(",");
const fish = input.map((s) => parseInt(s));

const simulate = () => {
  const len = fish.length
  for (let i = 0; i < len; i++) {
    if (fish[i] === 0) {
      fish.push(8);
      fish[i] = 6;
    } else {
      fish[i]--;
    }
  }
};

// console.log("Initial state:", fish.join(","));
for (let cycle = 1; cycle <= 80; cycle++) {
  simulate();
  // console.log(cycle, fish.join(","));
}
console.log('there are ' + fish.length + ' lanternfish')
