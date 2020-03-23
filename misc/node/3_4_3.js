//p086
/*
setTimeout(콜백 함수, 밀리초): 주어진 밀리초 이후에 콜백 함수를 실행
setInterval(콜백 함수, 밀리초): 주어진 밀리초마다 콜백 함수를 반복 실행
setImmediate(콜백 함수): 콜백 함수를 즉시 실행
위의 타이머 함수들은 모두 아이디를 반환. 이를 사용하여 취소 가능.
clearTimeout(아이디), clearInterval(아이디), clearImmediate(아이디)
*/

const timeout = setTimeout(() => {
    console.log('1.5초 후 실행');
}, 1500);

const interval = setInterval(() => {
    console.log('1초마다 실행');
}, 1000);

const timeout2 = setTimeout(() => {
    console.log('실행되지 않음');
}, 3000);

setTimeout(() => {
    clearTimeout(timeout2);
    clearInterval(interval);
},2500);

const immediate = setImmediate(() => {
    console.log('즉시 실행');
});


const immediate2 = setImmediate(() => {
    console.log('실행되지 않음');
});

clearImmediate(immediate2);

/*
즉시 실행
1초마다 실행
1.5초 후 실행
1초마다 실행
*/

/*
setImmediate(콜백) vs setTimeout(콜백, 0)
setImmediate(콜백) 과 setTimeout(콜백, 0) 에 담긴 콜백 함수는 이벤트 루푸를 거친 뒤 즉시 실행됨
특수한 경우에 setImmediate는 setTimeout(콜백, 0) 보다 먼저 실행 됨.
파일 시스템 접근, 네트워킹 같은 I/O 작업의 콜백 함수 안에서 타이머를 호출하는 경우.
하지만 setImmediate는가 항상 setTimeout(콜백, 0) 보다 먼저 호출되는 것은 아님. 
헷갈리니 setTimeout(콜백, 0) 사용하지 않는 것 권장.
*/
