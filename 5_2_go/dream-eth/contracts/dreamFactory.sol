// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0 <0.9.0;
// import "./supportDreamMap.sol";
// import "./dream.sol";
 
contract SupportDreamMap {
    // 这个合约用来保存，每个在网络中的，在这个平台中，都参与了什么Dream众筹项目
    mapping(address => address[]) supportMapWithDreamAddress;

    function getValue(
        address personAddress
    ) public view returns (address[] memory) {
        return supportMapWithDreamAddress[personAddress];
    }

    function setValue(address personAddress, address dreamAddress) public {
        supportMapWithDreamAddress[personAddress].push(dreamAddress);
    }
}

contract DreamContract {
    // 项目发起人
    address creator;
    // 众筹的项目名称
    string dreamName = unicode"飞机";
    // 众筹项目目标筹集金额
    uint targetAmount = 100 ether;
    // 每次众筹时，限制的金额数值
    uint limitSupportAmount = 10 ether;
    // 众筹截止日期，到此时间如果需要的金额不足，那么众筹失败，单位是秒
    uint endTime = block.timestamp + 3600;
    SupportDreamMap supportDreamMap;

    // 初始化
    constructor(
        string memory _dreamName,
        uint256 _targetAmount,
        uint256 _limitSupportAmount,
        uint256 _sumTime,
        address _creator,
        SupportDreamMap _supportDreamMap
    ) {
        // creator = msg.sender; // 这行代码不正确，因为只要谁new当前合约，那么谁的地址就是msg.sender,所以工厂函数的地址就被当成的msg.sender

        // 正确的是接受工厂传过来的发起人的地址
        creator = _creator;

        dreamName = _dreamName;
        targetAmount = _targetAmount;
        limitSupportAmount = _limitSupportAmount;
        endTime = block.timestamp + _sumTime;

        supportDreamMap = _supportDreamMap;
    }

    address[] supportList;
    mapping(address => bool) supportMap; //用于保存谁投递了

    function supportAdd() public payable isCreator {
        // 每个人只能参与一次
        require(supportMap[msg.sender] == false);
        // 每次只能参与制定的金额
        require(limitSupportAmount == msg.value);

        // 添加到众筹人数组中去
        supportList.push(msg.sender);
        // 标记当前账户为众筹人
        supportMap[msg.sender] = true;

        supportDreamMap.setValue(msg.sender, address(this));
    }

    // 退钱函数
    function backMoney() public payable isCreator {
        for (uint256 i = 0; i < supportList.length; i++) {
            payable(supportList[i]).transfer(limitSupportAmount);

            supportMap[supportList[i]] = false;
            delete supportList;
        }
    }

    // 查看合约当前的余额
    function getBalance() public view returns (uint) {
        return address(this).balance;
    }

    // 返回所有的投资参与人
    function getSupportList() public view returns (address[] memory) {
        return supportList;
    }

    // 花费请求
    struct ShopPayItem {
        // 买什么
        string shopName;
        // 需要多少钱
        uint shopPrice;
        // 像谁购买
        address shopAddress;
        // 多少人赞成了，超过半数就同意
        uint agreePersonCount;
        // 赞成人的标记
        // mapping(address => bool) agreePersonMap;
        // 申请的当前状态
        ApplyStatus payStatus;
    }
    // 赞成人的标记放这里吧
    // key是名字，值是赞成者的列表
    mapping(string => address[]) public agreePersonMapList;
    enum ApplyStatus {
        // 投票中，已批准，已完成
        Collecting,
        Approved,
        Completed
    }
    // 计划要买的东西列表
    ShopPayItem[] public shopPayList;

    function addShopPay(
        string memory _shopName,
        uint _shopPrice,
        address _shopAddress
    ) public isCreator {
        // 判断商品是否已经存在过了
        require(agreePersonMapList[_shopName].length == 0);
        ShopPayItem memory item = ShopPayItem({
            shopName: _shopName,
            shopPrice: _shopPrice,
            shopAddress: _shopAddress,
            agreePersonCount: 0,
            payStatus: ApplyStatus.Collecting
        });
        shopPayList.push(item);
    }

    // 对要买的东西进行审核，需不需要买，需要经过参与人的投票同意
    function agreeShopPayItem(string memory _shopName) public {
        // 检查是不是参与投资的人
        require(supportMap[msg.sender], unicode"您不是投资人");
        // 找到对应的需要付款的shopPayItem，前端或者后端传过来通过索引也好，传过来shopPayName名字也好都行，名字在创建时候做了唯一限制
        // 记录索引位置
        int currentIndex = -1;

        for (uint256 i = 0; i < shopPayList.length; i++) {
            if (
                (keccak256(abi.encodePacked(_shopName)) ==
                    keccak256(abi.encodePacked(shopPayList[i].shopName)))
            ) {
                currentIndex = int(i);
            }
        }

        // ==
        // 查看是否是进行中的投票
        require(
            shopPayList[uint(currentIndex)].payStatus == ApplyStatus.Collecting,
            unicode"当前项目已关闭"
        );
        // ==

        // 代表知道了
        require(currentIndex > -1, unicode"不要传递不存在的商品名称");
        // 确保当前的投票人没有投过票,检查是否投过票了，如果投过了就不需要了
        address[] memory currentPersonList = agreePersonMapList[_shopName];
        for (uint i = 0; i < currentPersonList.length; i++) {
            if (msg.sender == currentPersonList[i]) {
                revert(unicode"您投过票了");
            }
        }

        // 支持票+1
        shopPayList[uint(currentIndex)].agreePersonCount++;
        // 把自己标记为已经投票了
        agreePersonMapList[_shopName].push(msg.sender);
    }

    // 使用对应商品的钱
    function useShopPayItem(string memory _shopName) public isCreator {
        // 查找有没有这个项目
        int currentIndex = -1;

        for (uint256 i = 0; i < shopPayList.length; i++) {
            if (
                (keccak256(abi.encodePacked(_shopName)) ==
                    keccak256(abi.encodePacked(shopPayList[i].shopName)))
            ) {
                currentIndex = int(i);
            }
        }
        // 查看是否是进行中的,只要不是进行中，就停止逻辑
        if (
            shopPayList[uint(currentIndex)].payStatus != ApplyStatus.Collecting
        ) {
            revert();
        }

        // 查看票数是否过半
        require(
            shopPayList[uint(currentIndex)].agreePersonCount * 2 >
                shopPayList.length,
            unicode"票数不过半"
        );
        // 查看当前合约资金是否充足，已经大于这个商品需要的费用了
        require(
            address(this).balance > shopPayList[uint(currentIndex)].shopPrice
        );
        // 转账给商品的所属者
        payable(shopPayList[uint(currentIndex)].shopAddress).transfer(
            shopPayList[uint(currentIndex)].shopPrice
        );
        // 标记状态
        shopPayList[uint(currentIndex)].payStatus = ApplyStatus.Completed;
    }

    // 权限问题
    modifier isCreator() {
        require(msg.sender == creator, unicode"无权限操作");
        _;
    }

    // 返回投资人的数量
    function getSupportListCount() public view returns (uint) {
        return supportList.length; // 只有storage的数组才有长度
    }

    // 返回众筹的剩余时间
    function getDreamEndTime() public payable returns (uint) {
        return (endTime - block.timestamp);
    }

    // 返回购买的东西的数量
    function getShopPayListCount() public view returns (uint) {
        return shopPayList.length;
    }

    // 返回某一个买的东西的具体内容
    function getshopPayItemByShopName(
        string memory _shopName
    ) public view returns (ShopPayItem memory, address[] memory) {
        int currentIndex = -1;

        for (uint256 i = 0; i < shopPayList.length; i++) {
            if (
                (keccak256(abi.encodePacked(_shopName)) ==
                    keccak256(abi.encodePacked(shopPayList[i].shopName)))
            ) {
                currentIndex = int(i);
            }
        }

        // 代表知道了
        require(currentIndex > -1, unicode"不要传递不存在的商品名称");

        return (shopPayList[uint(currentIndex)], agreePersonMapList[_shopName]);
    }
}
 
contract DreamFactory {
    // 平台提供者
    address factoryCreator;
    // 使用当前平台产生的众筹项目
    address[] productDreamList;
    // 使用某个账户发起的所有众筹项目的地址
    mapping(address => address[]) productDreamMap;

// 用来维护一个参与Dream的人，以及人都参与了什么项目投资的map
    SupportDreamMap supportDreamMap;
    constructor() {
        factoryCreator = msg.sender;
    }

    // 通过平台创建一个梦想计划
    //     string memory _dreamName,
    //     uint256 _targetAmount,
    //     uint256 _limitSupportAmount,
    //     uint256 _sumTime
    function createDream(
        string memory _dreamName,
        uint _targetAmount,
        uint256 _limitSupportAmount,
        uint256 _sumTime
    ) public returns(DreamContract) {
        DreamContract dream = new DreamContract(
            _dreamName,
            _targetAmount,
            _limitSupportAmount,
            _sumTime,
            msg.sender,
            supportDreamMap
        );
      
        productDreamList.push(address(dream));

        productDreamMap[msg.sender].push(address(dream));
        return dream;
    }

    // 返回当前平台的所有众筹项目地址
    function getProductDreamList() public view returns (address[] memory) {
        return productDreamList;
    }

    // 返回某个账户发起的所有众筹项目地址
    function getCreateDreamListByPersonAddress() public view returns (address[] memory) {
        return productDreamMap[msg.sender];
    }

    // 返回某个账户所投资的全部众筹项目地址
    function getSupportDreamListByPersonAddress() public payable returns(address[] memory  ) {
        return supportDreamMap.getValue(msg.sender);
    }
}
