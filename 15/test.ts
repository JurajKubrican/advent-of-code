import fs from "fs";

const file = fs.readFileSync("./15/in.txt", "utf-8");

const input = file.split("\n");

const result = input.reduce((acc, line) => {
  const [, p] = line.split(" | ");
  const input = p.split(" ");
  return (
    acc +
    input.reduce(
      (acc, val) =>
        val.length === 2 || val.length === 3 || val.length === 4|| val.length === 7
          ? acc + 1
          : acc,
      0
    )
  );
}, 0);

console.log(result);
