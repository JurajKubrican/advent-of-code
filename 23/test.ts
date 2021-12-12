import fs from "fs";
import R from "ramda";

const file = fs.readFileSync("./23/in.txt", "utf-8");

const input = file.split("\n");

type Node = {
  id: string,
  v: string[] 
}

const lower = /[a-z]+/
const upper = /[A-Z]+/

const graph: {[key:string]:Node}  ={}
input.forEach(line=>{
  const [a,b]=line.split('-')
  if(graph[a]){
    graph[a].v.push(b)
  }
  if(graph[b]){
    graph[b].v.push(a)
  }
  if(!graph[a]){
    graph[a] = {id:a, v:[b]}
  }
  if(!graph[b]){
    graph[b] = {id:b, v:[a]}
  }


  console.log(a,b)
})

const results:string[][] = []

const traverse = (path:string[])=>{
  const current = path[path.length-1]
  const node = graph[current]
  node.v.forEach((vertex)=>{
    if(lower.exec(vertex)){
      if(path.includes(vertex)){
        return;
      }
    }
    if(vertex==='end'){
      results.push([...path,vertex])
      return
    }
    traverse([...path,vertex])
  })

}

traverse(['start'])

console.log(graph)
console.log(results,results.length)