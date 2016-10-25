 //params是输入行的数组参数,main方法为默认方法
 function test(num){
     /*code here*/
     var count = 0;
     num = Number(num);


     if(num < 0 && num > 65536) {
         return
     }

     for(var i=1; i*i <= num; i++) {
         if(num % i == 0) {
             count ++;
         }
     }

     console.log(count);
 }


 test('10');
