

with open('./3/in.txt') as f:
    lines = [line.rstrip() for line in f]

depth = 0
position = 0
for line in lines:
  [direction, magnitude] = line.split(' ')
  if direction=='forward':
    position+= int(magnitude)

  if direction=='down':
    depth+= int(magnitude)

  if direction=='up':
    depth-= int(magnitude)
 

print (depth ,position )