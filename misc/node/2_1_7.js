//p066
//2_1_6절의 프로미스 코드 다음과 같이 바꿀 수 있다 
async function findAndSaveUser(Users) {
    let user = await Users.findOne({});
    user.name = 'hs';
    user = await user.save();
    user = await Users.findOne({gender: 'm'});
}
//async function 으로 교체한 후, 프로미스 앞에 await를 붙임

//에러 처리하는 부분 추가 
async function findAndSaveUser(Users) {
    try {
        let user = await Users.findOne({});
        user.name = 'hs';
        user = await user.save();
        user = await Users.findOne({gender: 'm'});
    } catch(error){
        console.error(error);
    }
}

//Promise.all대체도 가능. 노드10버전부터 가능한 부분
const promise1 = Promise.resolve('성공1');
const promise2 = Promise.resolve('성공2');
(async () => {
    for await (promise of [promise1, promise2]) {
        console.log(promise);
    }
})();
//성공1
//성공2