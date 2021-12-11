import fs, { open } from "fs";
import R from "ramda";

const file = fs.readFileSync("./19/in.txt", "utf-8");

const input = file.split("\n");

const stack:string[] = []


const OPENING_CHARS = ['(','[','{','<']
const CLOSING_CHARS = ['(','[','{','>']
const CHAR_MAP={
  "(":")",
  "[":"]",
  "{":"}",
  "<":">",
  ")":"(",
  "]":"[",
  "}":"{",
  ">":"<",
}

type Char= keyof typeof CHAR_MAP

const CHAR_SCORE={
  ")":3,
  "]":57,
  "}":1197,
  ">":25137,
}

type ClosingChar = keyof typeof CHAR_SCORE

let score = 0
input.forEach((line,index)=>{
  for(let i = 0 ;i<line.length;i++){
    const char = line[i] as Char
    if(OPENING_CHARS.includes(char)){
      stack.push(char)
    }else{
      const stackChar = stack.pop()
      if(stackChar!==CHAR_MAP[char]){
        console.log(index,char,'!=',stackChar)
        score += CHAR_SCORE[char as ClosingChar]
      }
    }

  }
})

console.log(score)