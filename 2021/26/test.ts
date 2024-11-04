import fs from "fs";
import R from "ramda";

const file = fs.readFileSync("./25/in.txt", "utf-8");

const input = file.split("\n");

type Point = { x: number; y: number };
const size = input.reduce(
  (acc: Point, line: string) => {
    if (line.startsWith("fold") || line === "") return acc;
    const [x, y] = line.split(",");

    return {
      x: Math.max(acc.x, parseInt(x)),
      y: Math.max(acc.y, parseInt(y)),
    };
  },
  { x: 0, y: 0 }
);

const paper: boolean[][] = [];
for (let i = 0; i <= size.y; i++) {
  paper.push(Array(size.x + 1).fill(false));
}

const folds: { dir: "x" | "y"; num: number }[] = [];

input.forEach((line) => {
  if (line === "") return;
  if (line.startsWith("fold")) {
    const ln = line.replace("fold along", "");
    const [dir, num] = ln.split("=");
    folds.push({ dir: dir.trim() as "x" | "y", num: parseInt(num) });
    return;
  }
  const [x, y] = line.split(",");
  paper[parseInt(y)][parseInt(x)] = true;
});

const fold = (direction: "x" | "y", num: number) => {
  if (direction === "y") {
    for (let i = 1; i < size.y; i++) {
      if (paper[num - i] !== undefined && paper[num + i] !== undefined) {
        paper[num + i].forEach((val, index) => {
          paper[num - i][index] = paper[num - i][index] || val;
        });
      }
    }
    for (let i = num; i <= size.y; i++) {
      delete paper[i];
    }
    size.y = num - 1;
  } else {
    for (let i = 1; i <= size.x; i++) {
      for (let j = 0; j <= size.y; j++) {
        if (
          paper[j]?.[num - i] !== undefined &&
          paper[j]?.[num + i] !== undefined
        ) {
          paper[j][num - i] = paper[j][num - i] || paper[j][num + i];
        }
      }
    }
    for (let i = 0; i <= size.y; i++) {
      paper[i].splice(num, size.x - num + 1);
    }
    size.x = num - 1;
  }
};


paper.forEach((line) => console.log(line.map((c) => (c ? "#" : ".")).join("")));

const f = folds[0]
// fold(f.dir, f.num);
//   console.log();

//   paper.forEach((line) =>
//     console.log(line.map((c) => (c ? "#" : ".")).join(""))
//   );

folds.forEach((f) => {
  fold(f.dir, f.num);
  console.log();

  paper.forEach((line) =>
    console.log(line.map((c) => (c ? "#" : ".")).join(""))
  );
});

const score = paper.reduce(
  (acc, line) => acc + line.reduce((accc, c) => accc + (c ? 1 : 0), 0),
  0
);

console.log(score)
