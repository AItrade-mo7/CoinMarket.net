var capital = 1; // 本金
var target = 10000000; // 目标金钱
// var rate = 0.03; // 每次的利润 3%
var rate = 1; // 每次的利润 100%

var num = 0; // 交易次数
for (let i = 0; i < 100000; i++) {
  num++;
  capital = parseInt(capital * rate + capital); // 本金 = 本金 * 利率 + 本金   每次取整
  console.log('交易第', num, '次', ', 本金结余:', capital, '元');
  if (capital > target) {
    // 当本金  >  目标金钱则退出循环
    break;
  }
}
