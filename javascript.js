//Pass by
const a = () => {
  let a = 10;
  a = passBy(a);
  // passBy(a);
  console.log(a);
};
a();

function passBy(increment) {
  return ++increment;
}

// Computed constants are permitted at runtime in js
const runtime = 10;
const b = runtime + 30;
// console.log(b);

function testme() {
  return function (x = 10) {
    return x;
  };
}
man = testme();
console.log(man(1022));
console.log(man(33));
