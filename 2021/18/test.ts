import fs from "fs";
import R from "ramda";

const file = fs.readFileSync("./18/in.txt", "utf-8");

const input = file.split("\n");

const mapStr = input.map((i) => i.split(""));
const map = mapStr.map((i) => i.map((j) => parseInt(j)));

const sizeX = map.length;
const sizeY = map[0].length;

const groups: number[][] = [];
for (let i = 0; i <= sizeX; i++) {
  groups.push(Array(sizeY).fill(-1));
}

for (let i = 0; i < sizeX; i++) {
  for (let j = 0; j < sizeY; j++) {
    const val = map[i][j];
    if (val < 9) {
      groups[i][j] = 0;
    }
  }
}

const groupTree: { size: number; groups: number[] }[] = [];

let curBasin = 1;
for (let i = 0; i < sizeX; i++) {
  for (let j = 0; j < sizeY; j++) {
    const val = groups[i][j];
    if (val === 0) {
      const top = (groups[i - 1] && groups[i - 1][j]) ?? -1;
      const left = groups[i][j - 1] ?? -1;

      if (top > 0) {
        groups[i][j] = top;
        const index = groupTree.findIndex(
          (g) => g && g.groups.includes(top)
        );
        if(!groupTree[index]){
          console.log('WAT')
        }
        groupTree[index].size++;
        if (left > 0 && left !== top && !groupTree[index].groups.includes(left)) {
          console.log(groupTree);
          const index2 = groupTree.findIndex(
            (g) => g && g.groups.includes(left)
          );
          console.log(index2);
          groupTree[index].size += groupTree[index2].size;
          groupTree[index].groups = [
            ...groupTree[index].groups,
            ...groupTree[index2].groups,
          ];
          delete groupTree[index2];
        }
      } else if (left > 0) {
        groups[i][j] = left;
        const index = groupTree.findIndex(
          (g) => g && g.groups.includes(groups[i][j])
        );
        groupTree[index].size++;
      } else {
        groups[i][j] = curBasin++;
        groupTree.push({ size: 1, groups: [groups[i][j]] });
      }
    }
  }
}

const print = () => {
  groups.forEach((row) => console.log(row.join(" ")));
  map.forEach((row) => console.log(row.join(" ")));
};

groupTree.sort((a,b)=>b.size-a.size)

// print();
console.log(groupTree[0].size*groupTree[1].size*groupTree[2].size,);
