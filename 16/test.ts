import fs from "fs";
import R from "ramda";

const file = fs.readFileSync("./16/in.txt", "utf-8");

const input = file.split("\n");

const ONE: Symbol[] = ["c", "f"];
const FOUR: Symbol[] = ["b", "c", "d", "f"];
const SEVEN: Symbol[] = ["a", "c", "f"];
const EIGHT: Symbol[] = ["a", "b", "c", "d", "e", "f", "g"];

const TWO: Symbol[] = ["a", "c", "d", "e", "g"];
const THREE: Symbol[] = ["a", "c", "d", "f", "g"];
const FIVE: Symbol[] = ["a", "b", "d", "f", "g"];

const ZERO: Symbol[] = ["a", "b", "c", "e", "f", "g"];
const SIX: Symbol[] = ["a", "b", "d", "e", "f", "g"];
const NINE: Symbol[] = ["a", "b", "c", "d", "f", "g"];

type Symbol = "a" | "b" | "c" | "d" | "e" | "f" | "g";
type Display = { [key in Symbol]: Symbol[] };

const ALPHABET: Symbol[] = ["a", "b", "c", "d", "e", "f", "g"];

const readWord = (str?: string) => {
  if (str === undefined) throw new Error("WORD NOT FOUND");
  const arr = str.split("");
  arr.sort();
  return arr as Symbol[];
};

const filterOptions = (display: Display, keys: Symbol[], values: Symbol[]) => {
  ALPHABET.forEach((i) => {
    let line = [...display[i]];
    if (keys.includes(i)) {
      line = line.filter((l) => values.includes(l));
    } else {
      line = line.filter((l) => !values.includes(l));
    }
    display[i] = line;
  });
  return display;
};

const force = (display: Display, path: Symbol[]): Symbol[][] => {
  if (path.length === 7) {
    // console.log("OPTION:", path.join());
    return [path];
  }
  const letters = display[ALPHABET[path.length]];
  let res: Symbol[][] = [];
  letters.forEach((letter) => {
    if (!path.includes(letter)) {
      const r = force(display, [...path, letter]);
      res = [...res, ...r];
    }
  });
  return res;
};

const toNumber = (t: Symbol[]) =>{
  if (R.equals(t, ZERO)) return 0;
  if (R.equals(t, ONE)) return 1;
  if (R.equals(t, TWO)) return 2;
  if (R.equals(t, THREE)) return 3;
  if (R.equals(t, FOUR)) return 4;
  if (R.equals(t, FIVE)) return 5;
  if (R.equals(t, SIX)) return 6;
  if (R.equals(t, SEVEN)) return 7;
  if (R.equals(t, EIGHT)) return 8;
  if (R.equals(t, NINE)) return 9;
}

const decode = (inputs: string[]) => {
  let segmentOptions: Display = {
    a: ["a", "b", "c", "d", "e", "f", "g"],
    b: ["a", "b", "c", "d", "e", "f", "g"],
    c: ["a", "b", "c", "d", "e", "f", "g"],
    d: ["a", "b", "c", "d", "e", "f", "g"],
    e: ["a", "b", "c", "d", "e", "f", "g"],
    f: ["a", "b", "c", "d", "e", "f", "g"],
    g: ["a", "b", "c", "d", "e", "f", "g"],
  };

  const one = readWord(inputs.find((op) => op.length === ONE.length));
  const four = readWord(inputs.find((op) => op.length === FOUR.length));
  const seven = readWord(inputs.find((op) => op.length === SEVEN.length));
  const eight = readWord(inputs.find((op) => op.length === EIGHT.length));

  const A = seven.filter((x) => !one.includes(x));
  segmentOptions = filterOptions(segmentOptions, ["a"], A);

  segmentOptions = filterOptions(segmentOptions, ONE, one);
  const b_d = four.filter((l) => !one.includes(l));
  segmentOptions = filterOptions(segmentOptions, ["b", "d"], b_d);

  segmentOptions = filterOptions(segmentOptions, FOUR, four);
  segmentOptions = filterOptions(segmentOptions, SEVEN, seven);

  const cyphersRaw = force(segmentOptions, []);

  const cyphers = cyphersRaw.map((c) => {
    const res: any = {};
    for (let i = 0; i < c.length; i++) {
      res[c[i]] = ALPHABET[i];
    }
    return res;
  });

  //DEAFGBC
  const c = cyphers.filter((cypher) => {
    const translated = inputs.map((i) => {
      const r = i.split("").map((l) => cypher[l]);
      r.sort();
      return r;
    }) as Symbol[][];
    if (!translated.find((t) => R.equals(t, ZERO))) return false;
    if (!translated.find((t) => R.equals(t, ONE))) return false;
    if (!translated.find((t) => R.equals(t, TWO))) return false;
    if (!translated.find((t) => R.equals(t, THREE))) return false;
    if (!translated.find((t) => R.equals(t, FOUR))) return false;
    if (!translated.find((t) => R.equals(t, FIVE))) return false;
    if (!translated.find((t) => R.equals(t, SIX))) return false;
    if (!translated.find((t) => R.equals(t, SEVEN))) return false;
    if (!translated.find((t) => R.equals(t, EIGHT))) return false;
    if (!translated.find((t) => R.equals(t, NINE))) return false;

    return true;
  });

  return c[0] as { [key in Symbol]:Symbol};
};

const result = input.reduce((acc, line) => {
  const [u, p] = line.split(" | ");
  const unique = u.split(" ");
  const inputs = p.split(" ");

  const cypher = decode(unique);

  // console.log(cypher);
  const translated = inputs.map((i) => {
    const r = i.split("").map((l) => cypher[l as Symbol]);
    r.sort();
    return r;
  }).map(toNumber).join('') 

  console.log(parseInt(translated))


  return acc + parseInt(translated);
}, 0);

console.log(result);
