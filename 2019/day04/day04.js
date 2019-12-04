let total = 0
for (let i = 353096; i < 843212; i++) {
  if (test(i, testDouble)) {
    total += 1
  }
}
console.log('part1', total)

total = 0
for (let i = 353096; i < 843212; i++) {
  if (test(i, testDoublePart2)) {
    total += 1
  }
}
console.log('part2', total)


function testDouble(s) {
  if (s.match(/(.)\1/)) {
    return true
  }
  return false
}

function testDoublePart2(s) {
  const s2 = s.replace(/(.)\1\1+/g,'')
  if (s2.match(/(.)\1/)) {
    return true
  }
  return false
}


function test(n, testDoubleFunction) {
  const s = n.toString()

  for (let i = 0; i < s.length - 1; i++) {
    if (s[i] > s[i + 1]) {
      return false
    }
  }

  if (!testDoubleFunction(s)) {
    return false
  }

  return true
}
