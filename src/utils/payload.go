package utils

// Payload 结构体
type Payload struct {
}

var (
	// ICSPORT 工业协议端口
	ICSPORT = map[string]int{
		"Modbus": 502, "BACnet": 47808, "DNP3": 20000, "FINS": 9600, "OpcUA": 48400, "OpcDA": 49231,
		"OpcAE": 0, "S7COMM": 102, "ADS/AMS": 48898, "Umas": 502, "ENIP": 44818,
		"Hart/IP": 5094, "S7COMM_PLUS": 102, "IEC104": 2404, "CIP": 0, "GE_SRTP": 18245, "EGD": 7937,
		"H1": 4001, "FF": 1089, "MELSOFT": 5007, "Ovation": 111,
		"CoAP": 5683, "MQTT": 1883, "DLT645": 304, "MELSOFT(1E)": 5551,
	}
	// ICSPayload 工业协议数据
	ICSPayload = map[string][]string{
		// modbus
		"Modbus": {
			"01000000000601010000000a",     // 1 Read Coils
			"070000000006010200000004",     // 2 Read Discrete Inputs
			"01000000000601030000000a",     // 3 Read Holding Registers
			"000000000006010400000000",     // 4 Read Input Registers
			"c24b00000006ff050001ff00",     // 5
			"0001000000060a060005000b",     // 6 Write Single Register
			"000000000003010715",           // 7
			"0000000000060a0800010000",     // 8-1
			"000000000006010900000000",     // 9
			"000000000006010a00000000",     // 10
			"000000000006010b00000000",     // 11
			"000000000006010c00000000",     // 12
			"000000000006010d00000000",     // 13
			"000000000006010e00000000",     // 14
			"000100000008010f000000030100", // 15
			"00000000000701100000000000",   // 16
			"0000000000020a11",             // 17
			"000000000006011200000000",     // 18
			"000000000006011300000000",     // 19
			"000000000006010600000001",     // 20 Read File Record
			"000000000006011500000000",     // 21
			"0000000000080116000000010001", // 22
			"000000000006011700000000",     // 23
			"000000000006011800000000",     // 24
			"000000000005012b0e0100",       // 43-1
			"0001000000060a2b0e030100",     // 43-2
			"000000000006015a00000000",     // 90
		},
		// BACnet
		"BACnet": {
			"002100000046005a01fe028a8006b0d56f1b37583d110000000037583d1137583d1103000000000000000000000001000000000000000000000000000000000000000108040001010000fa00",
			"810a000e0120000d013dff603a09",
		},
		"DNP3": {
			"056408c404000300b4b8c0d7007ace",                                         // Confirm
			"05640bc404000300e42bd7c801ff000607fb",                                   // Read
			"056412c4040003001e7cc1c10232010701ebe45a87ff002801",                     // Write
			"05641ac404000300c2e6d5c6030c012801009f8603016400000030776400000000005b", // Select

		},
		"FINS": {
			"8000020000000000007a010103cccccc0001", // Memory Area Read
			"8000020000000000007a02018012000000ff", // Program Area Read

		},
		"OpcUA": {
			// ACK
			"48454c463900000000000000ffffff7fffffff7f0000000000000000190000006f70632e7463703a2f2f6c6f63616c686f73743a3438343030",
			// CLO
			"434c4f463e000000070000000d00000017000000170000000100c401020000eb030000381e6585c9c7d6011700000000000000ffffffffe8030000000000",
			// MSG
			"4d5347465e000000060000000d0000000a0000000a00000001007702020000ed03000014ec0f86c9c7d6010a00000000000000ffffffffe803000000000000000000000000000000000001000000005402000000ffffffff0000ffffffff",
			// OPN
			"4f504e4684000000000000002f000000687474703a2f2f6f7063666f756e646174696f6e2e6f72672f55412f5365637572697479506f6c696379234e6f6e65ffffffffffffffff01000000010000000100be010000aaab0086c9c7d6010100000000000000ffffffffe80300000000000000000000000000010000000000000080ee3600",
		},
		"OpcDA": {
			"0500008310000000480000005b380000200000000400060010640000cc0f0000d9548c70b1a27bc40500070000000000000000002df4e40cff69e54d92325fd8219fd4cf00000000",
		},
		// S7COMM
		"S7COMM": {
			// Job Read ver
			"0300002b02f080320100000f81001a00000402120a10020002000083000058120a10020001000083000320",
			// Ack_Data(3) PLC Stop
			"0300001402f08032030000430000010000000029",
			// UserData(7) Mode-transition
			"0300001902f080320700000000000800000100101000000000",
		},
		"ADS/AMS": {
			// 4
			"000028000000c0a8016301017882050d7560010110270400050008000000000000000d0000000000000005000200",
			// 2
			"00002c000000050d756001010100c0a801630101a583020004000c0000000000000073000000030000000200000004000000",
		},
		//
		"Umas": {
			"0058000000dc005a01fe0100d40099bf79931faef47b8d83bcfb8b38cd89ac8bbc2df2954c5814f952e40f22bb228b0fc8d2b1782691c898d2dedb5a1e87bbd6e9a093e4f562041d1f592e3a71144566957ea7d4179f955ffbd9d1eb3c0ab7a0c579c22e9bfbbf6547f9cffbddfed67ebffffaf0a71e0ff6b44bdb696b29cbf63a4554c5462c52bad1de586d85ef2c802dcae2348ff3b8c173b16f110fe48d3c9ed73675070cc4398fcf73b340d902378b942d72c34e69cee136d1d6a4e510f38432fe465ea3bd300bdb696b75a52dddb7fda3bddeefbdbdeecb8303623d4bdbf42f",
			"005300000008005a013400013000",
			"002100000046005a01fe028a8006b0d56f1b37583d110000000037583d1137583d1103000000000000000000000001000000000000000000000000000000000000000108040001010000fa00",
			"edc10000002c005acb3801560b6ab52543acec56da0fe83630c33b71e88788773febce092e70dbab0b6e0d5acb3400013000",
		},
		// 无对应数据
		"ENIP": {
			"70002b00b35e5535000000000000000000000000000000000000000001000200a10004005c58ea84b10017005e004b032100500324010001060000000b0000aa7800aa",
			"65000400000000000000000000000000000000000000000001000000",
		},
		"Hart/IP": {
			"010000000002000d0100007530",
			"01010200000c0008",
			"01000300000b001182264e0000d2300008",
			"01010300000b002086264e0000d2300f00d010040700000002000000000000c2",
			"01010300000a003386264e0000d2142200d07769686172746777000000000000000000000000000000000000000000000000db",
		},
		"S7COMM_PLUS": {
			// 1226
			"030000f602f080720100e731000004ca0000000100000120360000011d00040000000000a1000000d3821f0000a3816900151553657276657253657373696f6e5f31433943333846a3822100153b302e302e302e303a30496e74656c2852292038323537344c2047696761626974204e6574776f726b20436f6e6e656374696f6e2e54435049502e31a38228001500a38229001500a3822a0015184445534b544f502d544838374f39305f3831373632333735a3822b000401a3822c001201c9c38fa3822d001500a1000000d3817f0000a38169001515537562736372697074696f6e436f6e7461696e6572a2a20000000072010000",
		},
		"IEC104": {
			"680401001a00",
			"680e00000000460104000a0000000000",
			"680e50000a002e010a000a0001000001680e52000a00030103000a0001000001",
			"680443000000",
		},
		"CIP": {
			"70001e0087c7cbb8000000000000000000000000000000000000000001000200a1000400d77215deb1000a003a000e0320f524013002",
			"70001f0087c7cbb8000000000000000000000000000000000000000000000200a10004000c00fe80b1000b004300cb0000002000000080",
			"70002d0087c7cbb8000000000000000000000000000000000000000000000200a10004000c00fe80b10019004c00cb00000086000000130000aa7b00000001cc000001ccaa",
			"7000ce0187c7cbb8000000000000000000000000000000000000000000000200a10004000c00fe80b100ba01ab00cb00000086000001b40000aa750000010000019a00007b52167d3c8041a205b60cdc7f2c39e9fac762fca435ce4b239f1f22b8b5e3556fddad96b3c70f6b1c5ba618673de00071753bf90806363e43a718becab496b338567694a5b47e135ce393c139aa52cd1b4644262e570a04b7e52b7bfce439b1906deca4ee4eeeeb8e0dce32180cbc4dbce8113b1ebfc83e2ea6ca1e27aa6ad34fc84b8d86f6aef73df2d885d6c1b7968df7bc2bc822a8065c6bf1d82f41d213a3f0050b3e000ef041df16a6ef9d40843b9e910fdffe8f8f999a073e7f236f711e61a8afe862bf9f108d90e48ed70872a8bb55a24f298199b302e0fd8b6569386b5dafbf75ec476549df9855593c800e17bc4e8bacec6dc6f028b0206a98abac45ab4eba45af5118e8332dd2595352402125795e17318718d64ec849324550efea372b52d3c0975f85240ca5779b90700c3097e9f5c68da35ed6f612cb5986030918fce0dca90b72bd7820dc827208ab4d44d42088a5efc9c3838408f3326d155b93669da09fce756c61ede6ffe33678e5dda51632cfee6a9151f0f9eb38537e7b8050b789640079877941c6a8eca9ef3be6fd8965647bfdc36dd17803455a5300aa",
		},
		"GE_SRTP": {
			"0200b9000000000000010000000000000001000000000000000000000000b3c000000000100e00000101074cba010100000000ffff00ffff",
			"0200b8000000000000010000000000000001000000000000000000000000b2c000000000100e000001010408cf07f801000000ffff00ffff0200b9000000000000010000000000000001000000000000000000000000b3c000000000100e00000101040824034700000000ffff00ffff",
			"02003600000000000001000000000000000100000000000000000000000033c000000000100e00000101041602006900000000ffff00ffff",
			"0300dc003000000000000000000000000001000000000000000021581100d694100e0000003a000001013000d08118000101ff0238001c00000064000000000000000000640000000000000000006400000000000000000064000000000000000000640000000000",
		},
		"EGD": {
			"a9000000000000000000000000004e7000000000000000001020820040082184f0050200000000000000000000ba0f00004d0b000081070000b40f000087280000570900005af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000cdcc0442666683426666064266667c422f8dd743b9becb43b0d2d4439f1fd943bc91d8438faed84380b7703f61b2b249000000000000000000000000000000800000008000000080000000000000000000000000000000000000000000000080000000800000008000000000003096450000000000000000ac1c80c3128382c3bdf480c3b89e6a43ae0772430ad7724396211e3fc738084900000000000000000000000000000000000000000000008000000080000000800000000052ad4c480000000000000000f5285d430c025d43b0725c43b85e7f4385ab7c43a4b07d43e51600005c62d74800000000000000000000000000000000000000000000008000000080000000800000000000026b4600000000000000000030fbc60020a1460000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010270000e02e0000e02e0000e02e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			"0400111101004c00696600c8e12e00004006000000000100696600c8000000000000000000000000000000000000000000000000000000000200780500000000464700000000020000000000",
			"a9000000000000000000000000002e7000000000000000001020820040082184f0050200000000000000000000af0f00004d0b0000c4070000b20f000086280000580900005af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff5af1ffff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000cdcc0442666683426666064266667c422f8dd743b9becb43b0d2d4439f1fd943bc91d8438faed84380b7703f61b2b249000000000000000000000000000000800000008000000080000000000000000000000000000000000000000000000080000000800000008000000000003096450000000000000000ac1c80c3128382c3bdf480c3b89e6a43ae0772430ad7724396211e3fc738084900000000000000000000000000000000000000000000008000000080000000800000000052ad4c480000000000000000f5285d430c025d43b0725c43b85e7f4385ab7c43a4b07d43e51600005c62d74800000000000000000000000000000000000000000000008000000080000000800000000000026b4600000000000000000030fbc60020a1460000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
		"H1": {
			"5335100103050308010300000059ff02",
			"533510010305030801030000000aff02",
			"5335100103060f0300ff0700000000000302005900009e01002000060004360000000000",
		},
		"FF": {
			"000000020000000000000000044c002900000102030405060708010203040050607080102304050607008001020304050607080101020304050607080102030405060708010203040102030405060708010203040506070801020304010203040506000201010203000000020000020304050607080102030401",

			"030405060708010203040050607080102304050607008001020304050607080101020304050607080102030405060708010203040102030405060708010203040506070801020304010203040506000201010203000000020000020304050607080102030401",
		},
		"MELSOFT": {
			// 257
			"57000000001111070000ffff030000fe03000014001c080a0800000000000000040101010000000001",
			// 未知
			"57000900001111070000ffff030000fe03000022001c080a08000000000000000404030a0000000001000000009100a902000000000000",
		},
		"Ovation": {
			"0000001400000003fa5dbeef0000000100000000",
		},
		"CoAP": {
			"024531380334430d0a",
			"430105ca7216332b2e77656c6c2d6b6e6f776e046367265",
			"4302ffcd7216332b2e77656c6c2d6b6e6f776e04636f7265",
			"430105ca7216332b2e77656c6c2d6b6e6f776e04636f7265",
		},
		"MQTT": {
			"3017000b53616d706c65546f70696348656c6c6f204d515454e000",
		},
		"DLT645": {
			"fefefefe6851975903050068910e3e373337338d606a676a8c8d7777f916",
		},
		"MELSOFT(1E)": {
			"e0795e04ba6d6c2b59a6a085810000e6080045000034027d40008006e2030a0901010a090131c08615afd0841cb9d49bdde55018f8328b02000000ff0a007400000020530100",
			"00ff0a006900000020530100",
			"e0795e04ba6d6c2b59a6a085810000e6080045000034029140008006e1ef0a0901010a090131c08615afd0841d0dd49bddfa5018f81d80ae000000ff0a007e00000020530100",
		},
	}
	// BLACKPayload 黑名单数据
	BLACKPayload = map[int][]string{
		1210: {"dddd4141414141414141414141414141414141410000",
			"dddd41414141414141414141414141414141414141412e2e412e2e",
			"dddd414141414141414141414141414141410600000041414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141",
			"dddd41414141414141414141414141414141060000004141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141ffffff8f",
			"0001030000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			"10235467bd02000041410a00050041414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141"},
		12397: {"4141414141414141414141410a2e2e5c2e2e5c2e2e5c2e2e5c2e2e5c"},
		12401: {"4141010041410d0000004141414141414141030000002e2e"},
		20034: {
			"4e45544241414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141410000000000004141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141"},
		777: {
			"9090909090909090414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141"},
		44818: {
			"6f00000000000004020000000000b2000800970220c024000000",
			"6f00000000000004020000000000b20008000503200124013003",
		},
		36667: {"buyaocaowosssssssssssssssssssssssss"},
		7168:  {"b6 b6 b6 b6 b6 b6 b6 b6|gffcsada"},
		10086: {"My IP |3A|asda"},
		5060:  {"|28 29 20 7b|"},
		10001: {"lv0njxq80njxq80"},
		10002: {"! LOLNOGTFO|0A|"},
		10003: {"12345|7A 7A 7A 7A 72 71 71 71 71 73 73 73 73 7D 7D 7D 7D|"},
		10004: {"12345|B5 B5 B5 B5 BD BE BE BE BE BC BC BC BC B2 B2 B2 B2|"},
		10005: {"! PING|0A|"},
	}
	PORT = [...]int{80, 21, 110, 23, 25, 53, 67, 69, 443, 161, 162, 513, 1719, 1720, 554}
)

// OutputICS  工业协议数据
func (p *Payload) OutputICS() [][2]interface{} {
	payload := p.List()
	for M, P := range ICSPayload {
		for m, p := range ICSPORT {
			if M == m {
				for pay := range P {
					if p == 0 {
						break
					}
					payload = append(payload, [2]interface{}{p, P[pay]})
				}

			}
		}
	}
	return payload
}

// OutputBLACK  黑名单数据
func (p *Payload) OutputBLACK() [][2]interface{} {
	payload := p.List()
	for port, v := range BLACKPayload {
		for _, pay := range v {
			payload = append(payload, [2]interface{}{port, pay})
		}
	}

	return payload
}

// OutputTCP  TCP数据
func (p *Payload) OutputTCP() [][2]interface{} {
	return p.PORTRandom(p.List())
}

// OutputUDP  UDP数据
func (p *Payload) OutputUDP() [][2]interface{} {
	return p.PORTRandom(p.List())
}

// List  Pay列表
func (p *Payload) List() [][2]interface{} {
	payload := make([][2]interface{}, 0)
	return payload
}

// PORTRandom 随机端口
func (p *Payload) PORTRandom(payload [][2]interface{}) [][2]interface{} {
	for _, port := range PORT {
		payload = append(payload, [2]interface{}{port, RandStr(10)})
	}
	return payload
}

// Execute 小型payload工厂
func (p *Payload) Execute(mode string) [][2]interface{} {
	switch mode {
	case "ICS":
		{
			return p.OutputICS()

		}
	case "BLACK":
		{
			return p.OutputBLACK()

		}
	case "TCP":
		{
			return p.OutputTCP()
		}
	case "UDP":
		{
			return p.OutputUDP()
		}
	}
	return nil
}

var Output Payload
