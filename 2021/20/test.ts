import fs, { open } from "fs";
import R, { length } from "ramda";

const file = fs.readFileSync("./20/in.txt", "utf-8");

const input = file.split("\n");

const OPENING_CHARS = ["(", "[", "{", "<"];
const CLOSING_CHARS = ["(", "[", "{", ">"];
const CHAR_MAP = {
  "(": ")",
  "[": "]",
  "{": "}",
  "<": ">",
  ")": "(",
  "]": "[",
  "}": "{",
  ">": "<",
};

type Char = keyof typeof CHAR_MAP;

const CHAR_SCORE = {
  ")": 1,
  "]": 2,
  "}": 3,
  ">": 4,
};

type ClosingChar = keyof typeof CHAR_SCORE;

let scores:number[]=[];

input.forEach((line, index) => {
  let success = true;
  let stack: string[] = [];
  for (let i = 0; i < line.length; i++) {
    const char = line[i] as Char;
    if (OPENING_CHARS.includes(char)) {
      stack.push(char);
    } else {
      const stackChar = stack.pop();
      if (stackChar !== CHAR_MAP[char]) {
        success = false;
        stack = [];
      }
    }
  }
  if (success === true) {
    stack.reverse()
    const addition = stack
   
    .map((s) => CHAR_MAP[s as Char])
    .reduce((acc, s) => acc * 5 + CHAR_SCORE[s as ClosingChar], 0);
    console.log(
      addition,stack
        .map((s) => CHAR_MAP[s as Char])
        .join("")
    );
  
    scores.push(addition)
      
  }
});

scores.sort((a,b)=>a-b)
const index = Math.floor(scores.length/2)
console.log(scores[index])