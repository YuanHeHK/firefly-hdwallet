# hdwallet

1.  编译

   执行`make`命令

2.  配置文件配置

   2.1 在配置文件中```conf/app.conf```中配置助记词、加密文件存放路径配置，当应用本地私钥进行签名时，需配置私钥文件路径和对应的转账地址(目前只支持fil)

   	*  `PlainPath` 参数对应助记词存放路径
   	*  `HiddenPath` 参数对应助记词密后的存放路径
   	*  `LocalPrk` 参数对应私钥文件存放路径
   	*  `FilAddress` 参数对应fil的转账地址(出金地址)

   2.2  网络连接配置

   	*   `EthUrl`  、 `TrxUrl` 、`LotusHost` 分别对应ethereum、tron、filecoin的节点url
   	*   `UsdtLagNumber` 、`FilLagNumber` 、`Trc20UsdtLagNumber` 分别对应对ethereum、filecoin、tron网络监听时滞后最新块高的块数
   	*   `LotusToken`  对应调用lotus节点的write权限token，因为调用一些api时需要write权限

   2.3  fil转账参数配置

   	*  `MaxFee`  对应最高手续费
   	*  `ChangeGasLag`  对应改变手续费的等待块数，在第一次将message推送至消息池中后，会等待该块高数，若没有被打包，会改变gas设置的相关参数
   	*  `GasFeeCap`、`GasPremium`、`GasLimit` 若接收方的地址是第一次接收fil转账，在第一次转账不被打包后，改变gas相关参数时，会按照改参数设定
   	*  `TranRespUrl` 在接收到转账请求后，会立即给前端返回成功接到请求的信息，待转账被确认后(message上链)，将message状态等信息通过该参数对应的回调地址返回至前端

3.  运行

   3.1 第一次运行

   		* 执行`./walletmanager -encode` ， 提示input encode password后输入密码，出现input decode  password时，再次输入密码即可运行
   		* 在提示eth client、tran client、fil client 初始化成功后，先执行`ctrl+z`暂停进程，再执行`bg`即可后台运行程序

   3.2 之后运行

   	* 直接执行`./walletmanager`，出现input decode  password输入密码即可运行

4.  请求发送

   4.1 创建钱包地址

   	* 请求方式：GET

   	* `http://ip+port/newWallet?user_id=1`

   > 在使用助记词的情况下才可创建，user_id不能为0   

   4.2 余额查询

   	* 请求方式：GET
   	* `http://ip+port/balance?address=f1o4bgarfxjkbnnedjlm2wnruzkfkpne6j2mbvb6y&type=fil`

   > type的类型包含fil、erc-usdt、trc-usdt

   4.3 转账请求

   * 请求方式：POST

   * `http://ip+port/transfer`

   * 数据结构

     ```json
     {
     "order_id":"ff99979987887899",
     "user_id":0, //当type为“fil”时，由归集账户转账时user_id必须为0，由本地账户中账户转账，user_id为大于0的任意整
     "to":"t1p6eqsgfbvxm43apn7jp2hgkvdsfrol5hzjqcibq",
     "value":7, //float类型
     "type":"fil" //有"erc-usdt"(之前的usdt)，"fil"，"trc-usdt"(trc20网络)
     }
     ```

   4.4 归集请求

   	* 请求方式：Get
   	* `http://ip+port/collection?type=fil&level=high`

   > 1. 归集目前支持trc-usdt和fil
   > 2. level："medium","high"两种，是fil在转账过程中，估算转账手续费的等级，建议high，归集稳妥

