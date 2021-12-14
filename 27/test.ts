import fs from "fs";
import R, { identical } from "ramda";

const file = fs.readFileSync("./27/in.txt", "utf-8");

const input = file.split("\n");

let initialState = input.shift() ?? "";
input.shift();

const rules: { [key: string]: string } = {};
input.forEach((line) => {
  const [pattern, insert] = line.split(" -> ");
  rules[pattern] = insert;
});

const applyRules = () => {
  let newState = "";
  for (let i = 0; i < initialState?.length; i++) {
    const window = initialState[i] + initialState[i + 1];
    if (rules[window]) {
      newState += initialState[i] + rules[window];
    } else {
      newState += initialState[i];
    }
  }
  initialState = newState;
};

console.log(initialState);
for (let i = 0; i < 10; i++) {
  applyRules();
  console.log(initialState);
}

const score = () => {
  const stats = initialState.split("").reduce((acc, char) => {
    if (acc[char]) {
      acc[char] += 1;
    } else {
      acc[char] = 1;
    }
    return acc;
  }, {} as { [key: string]: number });

  const max = Object.entries(stats).reduce(
    (acc,[key,value]) => {
      if (value>acc.count) {
        return {char:key,count:value}
      }
      return acc;
    },
    { char: "", count: 0 }
  )
  const min = Object.entries(stats).reduce(
    (acc,[key,value]) => {
      if (value<acc.count) {
        return {char:key,count:value}
      }
      return acc;
    },
    { char: "", count: Number.MAX_SAFE_INTEGER }
  )
  return max.count - min.count
};

console.log(score());
