// p057
if (true) {
   var x = 3;
}
console.log(x); // 3

if(true) {
   const y =3;
}
console.log(y); //ReferenceError: y is not defined

/*
var은 함수 스코프를 가지므로 if문의 블록과 관계없이 접근할 수 있음.
하지만 const와 let은 블록 스코프를 가지므로 블록 밖에서는 변수에 접근할 수 없음
블록의 범위는 if, while, for, function등의 중괄호
*/
