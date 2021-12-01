

with open('./1/in.txt') as f:
    lines = [line.rstrip() for line in f]

count = 0
previousLine = int(lines[0])
for line in lines[1:]:
  if int(line) > previousLine:
    count +=1
  previousLine=int(line)

  
print (count)