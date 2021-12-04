import fs from 'fs'

const file = fs.readFileSync('./6/in.txt','utf-8')

// console.log(file)

const lines = file.split('\n')

const fnA = (lines:string[], i:number):string=>{
    const sum = lines.reduce((acc,line)=>{
        return acc + Number(line[i])
    },0)
    const correctNum = sum >= (lines.length)/2 ? "1":"0"

    const filtered = lines.filter((line)=>{

        return line[i]===correctNum
    })

    if( i >= lines[0].length-1)
        return correctNum

    return `${correctNum}${fnA(filtered,i+1)}`

}

const A = fnA(lines,0)

const fnB = (lines:string[], i:number):string=>{
    const sum = lines.reduce((acc,line)=>{
        return acc + Number(line[i])
    },0)
    const correctNum = sum < (lines.length)/2 ? "1":"0"

    const filtered = lines.filter((line)=>{
        return line[i]===correctNum
    })
    if(filtered.length===1){
        return filtered[0].slice(i)
    }

    if( i >= lines[0].length-1)
        return correctNum

    return `${correctNum}${fnB(filtered,i+1)}`

}

const B = fnB(lines,0)


console.log(A,B)
console.log(parseInt(A,2),parseInt(B,2),parseInt(A,2)*parseInt(B,2))