import fs from "fs";

const file = fs.readFileSync("./10/in.txt", "utf-8");

const input = file.split("\n");

type Line = { a: { x: number; y: number }; b: { x: number; y: number } };

const lines: Line[] = input.map((line) => {
  const [a, b] = line.split(" -> ");
  const aa = a.split(",");
  const bb = b.split(",");
  return {
    a: { x: parseInt(aa[0]), y: parseInt(aa[1]) },
    b: { x: parseInt(bb[0]), y: parseInt(bb[1]) },
  };
});

const size = lines.reduce(
  (acc, line) => Math.max(line.a.x, line.a.y, line.b.x, line.b.y, acc),
  0
);

const field: number[][] = [];
for (let i = 0; i <= size; i++) {
  field.push(Array(size + 1).fill(0));
}

const normalize = (line: Line) => ({
  x: line.a.x - line.b.x,
  y: line.a.y - line.b.y,
});

const drawLine = (line: Line) => {
  const { a, b } = line;
  const normal = normalize(line);
  // ==
  if (normal.x === 0) {
    console.log(line,'==')
    const from = Math.min(a.y, b.y);
    const to = Math.max(a.y, b.y);
    for (let i = from; i <= to; i++) {
      field[i][a.x] += 1;
    }
    return;
  }
  // ||
  if (normal.y === 0) {
    console.log(line,'||')
    const from = Math.min(a.x, b.x);
    const to = Math.max(a.x, b.x);
    for (let i = from; i <= to; i++) {
      field[a.y][i] += 1;
    }
    return;
  }

  // \
  if (normal.x + normal.y !== 0) {
    console.log(line,'\\')
    const step = normal.x < 0 ? +1 : -1;
    for (
      let point = line.a;
      point.x !== line.b.x;
      point = { x: point.x + step, y: point.y + step }
    ) {
      field[point.y][point.x] += 1;
    }
    field[b.y][b.x] += 1;
    return
  }
  
  // /
  console.log(line,'/')
  const step = normal.x < 0 ? +1 : -1;
    for (
      let point = line.a;
      point.x !== line.b.x;
      point = { x: point.x + step, y: point.y - step }
    ) {
      field[point.y][point.x] += 1;
    }
    field[b.y][b.x] += 1;
};

lines.forEach((line) => {
  drawLine(line);
});

const countPoints = () =>
  field.reduce(
    (acc, row) =>
      acc + row.reduce((accc, item) => (item > 1 ? accc + 1 : accc), 0),
    0
  );

console.log(countPoints());

  // field.forEach(row => console.log(row.join('').replaceAll('0','.')));

