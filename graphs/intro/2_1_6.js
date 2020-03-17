//p062
const condition = true;
const promise = new Promise((resolve, reject) => {
    if (condition) {
        resolve('성공')
    } else {
        reject('실패')
    }
});

promise
  .then((message) => {
    console.log(message); //성공
  })
  .catch((error) => {
    console.error(error);
  });



//p064 예제. 돌리지는 않음 
function findAndSaveUser(Users) {
    Users.findOne({})
      .then((user) => {
          user.name = 'hs';
          return user.save();
      })
      .then((user) => {
          return Users.findOne({gender:'m'});
      })
      .then((user) => {
          //생략
      })
      .catch(err => {
           console.error(err); //콜백에서 매번 따로 처리해야 했던 에러도 마지막 catch에서 한번에 처리가능
      });
}



//065
//프로미스 여러개 한번에 실행하는 방법  
//Promise.all에 넣으면 모두 resolve될 때 까지 기다렸다가 then으로 넘어감. 개중 하나라도 reject되면 catch로 넘어감
//Promise.resolve는 즉시 resolve하는 프로미스를 만드는 방법
const promise1 = Promise.resolve('성공1');
const promise2 = Promise.resolve('성공2');
Promise.all([promise1, promise2])
  .then((result) => {
      console.log(result); //[ '성공1', '성공2' ]
  })
  .catch((error) => {
      console.error(error);
  });