function distance(target) {
  const r = root(target)
  const side = (r-1)/2
  let v = r*r
  let d = side

  while (v > target) {
    v -= 1
    d -= 1
    if (d === -side) {
      d = side
    }
  }
  return Math.abs(d) + side
}

function root(target) {
  let root = 1
  while (root * root < target) {
    root += 2
  }
  return root
}

function check(name, result, expected) {
  if (result != expected) {
    console.log(`fail: ${name}: result: ${result} did not match expected: ${expected}`)
  }
}

check(1, distance(1), 0)
check(12, distance(12), 3)
check(23, distance(23), 2)
check(1024, distance(1024), 31)

console.log('part1', distance(289326))


/// part 2

function g(a,x,y) {
  return a[x][y] || 0;
}

function sum(a, x, y) {
  return g(a, x+1,y-1) +
         g(a, x+1,y) +
         g(a, x+1,y+1) +
         g(a, x,y-1) +
         g(a, x,y) +
         g(a, x,y+1) +
         g(a, x-1,y-1) +
         g(a, x-1,y) +
         g(a, x-1,y+1);
}

function part2(target) {
  const outerRing = root(target); //root(289326)

  let a = []
  for (let i = -outerRing; i < outerRing; i++) {
    a[i] = []
  }

  a[0][0] = 1
  let r = 3
  y = 0
  while (1) {
    const s = (r - 1) / 2
    x = s
    
    // move north
    while (y > -s) {
      const z = sum(a, x, y);
      if (z > target) return z;
      a[x][y] = z
      y--
    }
    // move west
    while (x > -s) {
      const z = sum(a, x, y);
      if (z > target) return z;
      a[x][y] = z
      x--
    }
    // move south
    while (y < s) {
      const z = sum(a, x, y);
      if (z > target) return z;
      a[x][y] = z
      y++
    }
    // move east
    while (x < s) {
      const z = sum(a, x, y);
      if (z > target) return z;
      a[x][y] = z
      x++
    }
    const z = sum(a, x, y);
    if (z > target) return z;
    a[x][y] = z

    r += 2
  }
}

console.log('part2', part2(289326))
