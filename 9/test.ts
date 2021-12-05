import fs from "fs";

const file = fs.readFileSync("./9/in.txt", "utf-8");

const input = file.split("\n");

type Line = {a:{x:number,y:number},b:{x:number,y:number}}

const lines:Line[] = input.map((line) => {
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

const field:number[][] = []
for(let i = 0; i<=size; i++){
  field.push(Array(size+1).fill(0))
}


const drawLine = (line:Line) =>{
  //==
  if(line.a.x === line.b.x){
    const from = Math.min(line.a.y ,line.b.y)
    const to = Math.max(line.a.y ,line.b.y)
    for (let i = from; i<=to; i++){
      field[i][line.a.x]+=1
    }
  }
  // ||
  if(line.a.y === line.b.y){
    const from = Math.min(line.a.x ,line.b.x)
    const to = Math.max(line.a.x ,line.b.x)
    for (let i = from; i<=to; i++){
      field[line.a.y][i]+=1
    }
  }
}

lines.forEach(line=>{
  drawLine(line)
})


const countPoints = ()=>
  field.reduce((acc,row)=>acc+row.reduce((accc,item)=>item>1?accc+1:accc,0),0)

console.log(countPoints())

// const print = ()=>{
//   field.forEach(row => console.log(row.join('')));
// }

// print()
