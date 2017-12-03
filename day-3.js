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
