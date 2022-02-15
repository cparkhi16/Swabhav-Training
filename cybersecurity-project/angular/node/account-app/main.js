const User=require('./user')
const ledger=require('./ledger')

ledgernew=new ledger.Ledger();
ledgernew.listenToBalanceChange();

shan=new User("shan");
shan.addAccount(45,1000);
shan.addAccount(23,2000);
shan.addAccount(76,5000);
console.log(shan.getAccount(23).getBalance());
console.log(shan.getAccount(45).getBalance());
console.log(shan.getAccount(76).getBalance());
shan.getAccount(45).withdraw(1000);
shan.getAccount(23).deposit(1000);
shan.getAccount(76).deposit(500);
console.log(shan.getAccount(23).getBalance());
console.log(shan.getAccount(45).getBalance());
console.log(shan.getAccount(76).getBalance());

suma=new User("suma");
suma.addAccount(33,2000);
suma.addAccount(22,3000);
console.log(suma.getAccount(33).getBalance());
console.log(suma.getAccount(22).getBalance());
suma.getAccount(22).withdraw(200);
console.log(suma.getAccount(33).getBalance());
console.log(suma.getAccount(22).getBalance());

