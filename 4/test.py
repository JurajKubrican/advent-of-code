

with open('./3/in.txt') as f:
    lines = [line.rstrip() for line in f]

aim = 0
depth = 0
position = 0
for line in lines:
  [direction, magnitude] = line.split(' ')
  if direction=='forward':
    position+= int(magnitude)
    depth += aim* int(magnitude)

  if direction=='down':
    aim+=int(magnitude)

  if direction=='up':
    aim-=int(magnitude)
 

print (depth *position )