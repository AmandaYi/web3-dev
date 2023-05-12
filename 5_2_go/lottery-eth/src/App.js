
import { useEffect, useState } from "react";
import web3 from "./utils/web3Instance"
import LotterContractInstance from "./eth/LotterInstance";
import Home from "./components/Home";
import { Button } from "semantic-ui-react";


function App() {

  // 管理员是谁
  const [manager, setManager] = useState("");
  // 多少期了
  const [round, setRound] = useState(0);
  // 胜利的人
  const [winner, setWinner] = useState("");
  // 获取彩民人数
  const [playersCount, setPlayersCount] = useState(0);
  // 获取余额
  const [dappBalance, setDappBalance] = useState(0);
  // 获取所有彩民的数据
  const [allPlayers, setAllPlayers] = useState([]);
  const [currentAccount, setCurrentAccount] = useState("");
  async function init() {
    setManager(await LotterContractInstance.methods.manager().call())

    setRound(await LotterContractInstance.methods.round().call())
    setWinner(await LotterContractInstance.methods.winner().call())
    setPlayersCount(await LotterContractInstance.methods.getPlayersCount().call())
    let dappBalanceWei = await LotterContractInstance.methods.getBalance().call()
    setDappBalance(web3.utils.fromWei(dappBalanceWei, 'ether'))
    // allPlayers,
    // setAllPlayers(await LotterContractInstance.methods.getAllPlayers().call())

    let accounts = await web3.eth.getAccounts()
    // console.log(accounts)
    setCurrentAccount(accounts[0])
  }
  setInterval(() => {
    init()
  }, 5000)
  useEffect(() => {
    init()

  }, [manager, round, winner, playersCount, dappBalance, currentAccount])
  function open() {
    // LotterContractInstance.methods.open().transaction()
  }
  async function cancel() {
    let accounts = []
    try {
      // await window.ethereum.enable()

      // Request account access if needed
      // let result  = await window.ethereum.send('eth_requestAccounts');
      // accounts= result.result
      // Accounts now exposed, use them
      // ethereum.send('eth_sendTransaction', { from: accounts[0], /* ... */ })
      let accounts = await web3.eth.getAccounts()
      console.log(accounts)
    } catch (ex) {
      return
    }
    // console.log(manager)
    // console.log(currentAccount)
    // let accounts = await web3.eth.getAccounts()

    LotterContractInstance.methods.cancel().send({
      from: accounts[0],
      gas: currentAccount,
      value: 0
    })
  }
  function play() {
    // LotterContractInstance.methods.play().call()
  }
  return (
    <><Home
      manager={manager}
      round={round}
      winner={winner}
      playersCount={playersCount}
      dappBalance={dappBalance}
      allPlayers={allPlayers}
      currentAccount={currentAccount}
    />
      <div>
        <Button onClick={open} >开奖</Button>
        <Button onClick={cancel}>退款</Button>
        <Button onClick={play}>投注</Button>
      </div>
    </>
  );
}

export default App;
