

with open('./1/in.txt') as f:
    lines = [int(line.rstrip()) for line in f]

count = 0
windowA = []
windowB = lines[0:3]
for line in lines[3:]:
  windowA = windowB
  windowB = windowB[1:]
  windowB.append(line) 

  if sum(windowB) > sum(windowA):
    count +=1

print (count)