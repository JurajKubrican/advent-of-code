from statistics import median



with open('./5/in.txt') as f:
  lines = [line.rstrip() for line in f]


a = ''
b = ''
for i in range (0,12):
  column =[int(row[i]) for row in lines]
  med = (int(median(column)))
  a += str(med)
  b += str(int(not(med)))

print (a, b, int(a,2)*int(b,2)) 

