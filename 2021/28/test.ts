import fs from "fs";
import { resourceUsage } from "process";
import R, { identical, merge } from "ramda";

const file = fs.readFileSync("./28/in.txt", "utf-8");

const input = file.split("\n");

let initialState = input.shift() ?? "";
input.shift();

const rules: { [key: string]: string } = {};
input.forEach((line) => {
  const [pattern, insert] = line.split(" -> ");
  rules[pattern] = insert;
});

type Hist =  {[key: string]: number }

const histogram = (str:string)=>str.split("").reduce((acc, char) => {
  if (acc[char]) {
    acc[char] += 1;
  } else {
    acc[char] = 1;
  }
  return acc;
}, {} as Hist);

const mergeHistograms = (h1:Hist,h2:Hist)=>{
  return Object.entries(h1).reduce((acc,[char,value])=>{
    if (acc[char]) {
      acc[char] += value;
    } else {
      acc[char] = value;
    }
    return acc
  },{...h2})
}

const dfsMem:{[key:string]:Hist} = {}

const dfs = (depth:number,str:string):Hist=>{

  depth++
  if(depth>40){
    return histogram(str.slice(0,-1))
  }

  if(rules[str]){
    if(dfsMem[depth+str])return dfsMem[depth+str]
    const h1 = dfs(depth,str[0]+rules[str])
    const h2 = dfs(depth,rules[str]+str[1])
    const res =  mergeHistograms(h1,h2)
    dfsMem[depth+str] = res
    // console.log('add', depth+'_'+str)
    return res
  }

  return histogram(str.slice(0,-1))
}

let stats:Hist = {}
stats = mergeHistograms(stats,histogram(initialState.slice(-1)))
for (let i = 0; i < initialState?.length-1; i++) {
  const window = initialState[i] + initialState[i + 1];
  const h = dfs(0,window) 
  console.log(i)
  stats = mergeHistograms(stats,h)
}



const score = () => {

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




console.log(score(),stats);
