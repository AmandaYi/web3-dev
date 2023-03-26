package main

func main() {
	/** 三次握手
	--建立连接STR--
	主动发起连接请求端，发送SYN 2000(0),{这里的SYN是标识，2000是位起始位置，可以随便写，0代表数据包，默认第一次建立连接是0}
	被动接受连接请求端，发送ACK 2001(0) 和 SYN 8000(0) {这里的ACK是应答SYN，2001是必须在主动请求端的起始位置加一即可}，
	{SYN同上，8000是被动端定义的起始位置,8000后面的0是数据包}
	主动发起连接请求端，收到信息后，发送ACK 8001用来应答被动端的SYN标识{ACK是应答被动端的SYN，8001是在被动端的起始位置加一}

	至此，Accept函数调用完成，Dial函数调用完成。Accept阻塞的地方已经开始往下执行代码。
	--建立连接END--

	--数据传输STR--
	主动发起连接请求端，发送数据包1000(50),ACK 8001,{这里的1000是起始位置，50是数据包大小，ACK是为了防止建立连接STR的最后一步握手失败，要重新发一下ACK 8001}
	被动接受连接请求端，发送数据包8001(50),ACK 1050,{这里的ACK代表收到了信息，1050代表1050以前的位的数据都收到了，8001(50)是数据包}
	主动发起连接请求端，发送ACK 8051{这里的ACK代表收到了信息，8051代表8051之前的位的数据都收到了}

	至此，Read和Write函数调用完成。
	--数据传输END--


	--关闭连接STR--
	主动关闭请求端，发送FIN 2000(0){这里的FIN是一个指令暂用一个位，2000是位起始位置，0是带的数据，可以没有，在关闭的时候一般不需要}
	被动关闭请求端，发送ACK 2001{这里的ACK是应答的意思，2001代表2001之前的位的所有指令都收到了，因为FIN占用一个位是，，所以才加一}

	上面两步完成后，主动关闭请求端处理半关闭状态，也就是主动关闭已经听到了对方应答了关闭

	被动关闭请求端，发送FIN 8000{这里的FIN是指令，8000随便找的位起始位置}
	主动关闭请求端，发送ACK 8001{这里的ACK应答一下被动关闭请求端即可}

	至此，四次挥手完毕，两边都关闭了TCP连接
	--关闭连接END--


	为什么，建立连接需要三次握手，关闭连接需要四次挥手，因为半关闭的原因。然后详细说一下挥手的半关闭在什么步骤就行了。
	*/

	//TCP最大连接字节，使用MSS标识在建立连接的时候标识即可，
	//也就是第一次SYN的时候，不过即使主动端写了MSS，这里的一般是主动端最大传输包大小，要求不要超过这个值自己还能接受。
	//同时，服务器也有一个MSS，也是用来给主动端说明，自己能最大接受的传输包
}
