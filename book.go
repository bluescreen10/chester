// Code generated from book/Perfect2023.bin. DO NOT EDIT.

//go:generate go run internal/cmd/book/main.go

package main

type polyglotRandoms struct {
	Pieces      [Color(2)][Piece(6)][Square(64)]uint64
	Castling    [16]uint64
	EnPassant   [8]uint64
	WhiteToMove uint64
}

var polyglotTable = polyglotRandoms{
	Pieces: [2][6][64]uint64{
		{
			{
				0x799e81f05bc93f31,
				0x86536b8cf3428a8c,
				0x97d7374c60087b73,
				0xa246637cff328532,
				0x43fcae60cc0eba0,
				0x920e449535dd359e,
				0x70eb093b15b290cc,
				0x73a1921916591cbd,
				0xb9fd7620e7316243,
				0x5a7e8a57db91b77,
				0xb5889c6e15630a75,
				0x4a750a09ce9573f7,
				0xcf464cec899a2f8a,
				0xf538639ce705b824,
				0x3c79a0ff5580ef7f,
				0xede6c87f8477609d,
				0x2171e64683023a08,
				0x5b9b63eb9ceff80c,
				0x506aacf489889342,
				0x1881afc9a3a701d6,
				0x6503080440750644,
				0xdfd395339cdbf4a7,
				0xef927dbcf00c20f2,
				0x7b32f7d1e03680ec,
				0x9c1633264db49c89,
				0xb3f22c3d0b0b38ed,
				0x390e5fb44d01144b,
				0x5bfea5b4712768e9,
				0x1e1032911fa78984,
				0x9a74acb964e78cb3,
				0x4f80f7a035dafb04,
				0x6304d09a0b3738c4,
				0x87b3e2b2b5c907b1,
				0xa366e5b8c54f48b8,
				0xae4a9346cc3f7cf2,
				0x1920c04d47267bbd,
				0x87bf02c6b49e2ae9,
				0x92237ac237f3859,
				0xff07f64ef8ed14d0,
				0x8de8dca9f03cc54e,
				0x27e6ad7891165c3f,
				0x8535f040b9744ff1,
				0x54b3f4fa5f40d873,
				0x72b12c32127fed2b,
				0xee954d3c7b411f47,
				0x9a85ac909a24eaa1,
				0x70ac4cd9f04f21f5,
				0xf9b89d3e99a075c2,
				0x14acbaf4777d5776,
				0xf145b6beccdea195,
				0xdabf2ac8201752fc,
				0x24c3c94df9c8d3f6,
				0xbb6e2924f03912ea,
				0xce26c0b95c980d9,
				0xa49cd132bfbf7cc4,
				0xe99d662af4243939,
				0x5355f900c2a82dc7,
				0x7fb9f855a997142,
				0x5093417aa8a7ed5e,
				0x7bcbc38da25a7f3c,
				0x19fc8a768cf4b6d4,
				0x637a7780decfc0d9,
				0x8249a47aee0e41f7,
				0x79ad695501e7d1e8,
			},
			{
				0xf678647e3519ac6e,
				0x1b85d488d0f20cc5,
				0xdab9fe6525d89021,
				0xd151d86adb73615,
				0xa865a54edcc0f019,
				0x93c42566aef98ffb,
				0x99e7afeabe000731,
				0x48cbff086ddf285a,
				0x4feabfbbdb619cb,
				0x742e1e651c60ba83,
				0x9a9632e65904ad3c,
				0x881b82a13b51b9e2,
				0x506e6744cd974924,
				0xb0183db56ffc6a79,
				0xed9b915c66ed37e,
				0x5e11e86d5873d484,
				0x722ff175f572c348,
				0x1d1260a51107fe97,
				0x7a249a57ec0c9ba2,
				0x4208fe9e8f7f2d6,
				0x5a110c6058b920a0,
				0xcd9a497658a5698,
				0x56fd23c8f9715a4c,
				0x284c847b9d887aae,
				0xa90b24499fcfafb1,
				0x77a225a07cc2c6bd,
				0x513e5e634c70e331,
				0x4361c0ca3f692f12,
				0xd941aca44b20a45b,
				0x528f7c8602c5807b,
				0x52ab92beb9613989,
				0x9d1dfa2efc557f73,
				0x40e087931a00930d,
				0x8cffa9412eb642c1,
				0x68ca39053261169f,
				0x7a1ee967d27579e2,
				0x9d1d60e5076f5b6f,
				0x3810e399b6f65ba2,
				0x32095b6d4ab5f9b1,
				0x35cab62109dd038a,
				0x87d380bda5bf7859,
				0x16b9f7e06c453a21,
				0x7ba2484c8a0fd54e,
				0xf3a678cad9a2e38c,
				0x39b0bf7dde437ba2,
				0xfcaf55c1bf8a4424,
				0x18fcf680573fa594,
				0x4c0563b89f495ac3,
				0xd2b7adeeded1f73f,
				0xf7a255d83bc373f8,
				0xd7f4f2448c0ceb81,
				0xd95be88cd210ffa7,
				0x336f52f8ff4728e7,
				0xa74049dac312ac71,
				0xa2f61bb6e437fdb5,
				0x4f2a5cb07f6a35b3,
				0xc547f57e42a7444e,
				0x78e37644e7cad29e,
				0xfe9a44e9362f05fa,
				0x8bd35cc38336615,
				0x9315e5eb3a129ace,
				0x94061b871e04df75,
				0xdf1d9f9d784ba010,
				0x3bba57b68871b59d,
			},
			{
				0x3d5774a11d31ab39,
				0x8a1b083821f40cb4,
				0x7b4a38e32537df62,
				0x950113646d1d6e03,
				0x4da8979a0041e8a9,
				0x3bc36e078f7515d7,
				0x5d0a12f27ad310d1,
				0x7f9d1a2e1ebe1327,
				0xe479ee5b9930578c,
				0xe7f28ecd2d49eecd,
				0x56c074a581ea17fe,
				0x5544f7d774b14aef,
				0x7b3f0195fc6f290f,
				0x12153635b2c0cf57,
				0x7f5126dbba5e0ca7,
				0x7a76956c3eafb413,
				0x829626e3892d95d7,
				0x92fae24291f2b3f1,
				0x63e22c147b9c3403,
				0xc678b6d860284a1c,
				0x5873888850659ae7,
				0x981dcd296a8736d,
				0x9f65789a6509a440,
				0x9ff38fed72e9052f,
				0xd2733c4335c6a72f,
				0x7e75d99d94a70f4d,
				0x6ced1983376fa72b,
				0x97fcaacbf030bc24,
				0x7b77497b32503b12,
				0x8547eddfb81ccb94,
				0x79999cdff70902cb,
				0xcffe1939438e9b24,
				0xb7a0b174cff6f36e,
				0xd4dba84729af48ad,
				0x2e18bc1ad9704a68,
				0x2de0966daf2f8b1c,
				0xb9c11d5b1e43a07e,
				0x64972d68dee33360,
				0x94628d38d0c20584,
				0xdbc0d2b6ab90a559,
				0x1dd01aafcd53486a,
				0x1fca8a92fd719f85,
				0xfc7c95d827357afa,
				0x18a6a990c8b35ebd,
				0xcccb7005c6b9c28d,
				0x3bdbb92c43b17f26,
				0xaa70b5b4f89695a2,
				0xe94c39a54a98307f,
				0x11317ba87905e790,
				0x7fbf21ec8a1f45ec,
				0x1725cabfcb045b00,
				0x964e915cd5e2b207,
				0x3e2b8bcbf016d66d,
				0xbe7444e39328a0ac,
				0xf85b2b4fbcde44b7,
				0x49353fea39ba63b1,
				0x23b70edb1955c4bf,
				0xc330de426430f69d,
				0x4715ed43e8a45c0a,
				0xa8d7e4dab780a08d,
				0x572b974f03ce0bb,
				0xb57d2e985e1419c7,
				0xe8d9ecbe2cf3d73f,
				0x2fe4b17170e59750,
			},
			{
				0xebe9ea2adf4321c7,
				0x3219a39ee587a30,
				0x49787fef17af9924,
				0xa1e9300cd8520548,
				0x5b45e522e4b1b4ef,
				0xb49c3b3995091a36,
				0xd4490ad526f14431,
				0x12a8f216af9418c2,
				0xabeeddb2dde06ff1,
				0x58efc10b06a2068d,
				0xc6e57a78fbd986e0,
				0x2eab8ca63ce802d7,
				0x14a195640116f336,
				0x7c0828dd624ec390,
				0xd74bbe77e6116ac7,
				0x804456af10f5fb53,
				0x2472f6207c2d0484,
				0xc2a1e7b5b459aeb5,
				0xab4f6451cc1d45ec,
				0x63767572ae3d6174,
				0xa59e0bd101731a28,
				0x116d0016cb948f09,
				0x2cf9c8ca052f6e9f,
				0xb090a7560a968e3,
				0xeb3593803173e0ce,
				0x9c4cd6257c5a3603,
				0xaf0c317d32adaa8a,
				0x258e5a80c7204c4b,
				0x8b889d624d44885d,
				0xf4d14597e660f855,
				0xd4347f66ec8941c3,
				0xe699ed85b0dfb40d,
				0x1fe2cca76517db90,
				0xd7504dfa8816edbb,
				0xb9571fa04dc089c8,
				0x1ddc0325259b27de,
				0xcf3f4688801eb9aa,
				0xf4f5d05c10cab243,
				0x38b6525c21a42b0e,
				0x36f60e2ba4fa6800,
				0x66c1a2a1a60cd889,
				0x9e17e49642a3e4c1,
				0xedb454e7badc0805,
				0x50b704cab602c329,
				0x4cc317fb9cddd023,
				0x66b4835d9eafea22,
				0x219b97e26ffc81bd,
				0x261e4e4c0a333a9d,
				0x4ed0fe7e9dc91335,
				0xe4dbf0634473f5d2,
				0x1761f93a44d5aefe,
				0x53898e4c3910da55,
				0x734de8181f6ec39a,
				0x2680b122baa28d97,
				0x298af231c85bafab,
				0x7983eed3740847d5,
				0xa09e8c8c35ab96de,
				0xfa7e393983325753,
				0xd6b6d0ecc617c699,
				0xdfea21ea9e7557e3,
				0xb67c1fa481680af8,
				0xca1e3785a9e724e5,
				0x1cfc8bed0d681639,
				0xd18d8549d140caea,
			},
			{
				0xcd04f3ff001a4778,
				0xe3273522064480ca,
				0x9f91508bffcfc14a,
				0x49a7f41061a9e60,
				0xfcb6be43a9f2fe9b,
				0x8de8a1c7797da9b,
				0x8f9887e6078735a1,
				0xb5b4071dbfc73a66,
				0x1f2b1d1f15f6dc9c,
				0xb69e38a8965c6b65,
				0xaa9119ff184cccf4,
				0xf43c732873f24c13,
				0xfb4a3d794a9a80d2,
				0x3550c2321fd6109c,
				0x371f77e76bb8417e,
				0x6bfa9aae5ec05779,
				0x9c1169fa2777b874,
				0x78edefd694af1eed,
				0x6dc93d9526a50e68,
				0xee97f453f06791ed,
				0x32ab0edb696703d3,
				0x3a6853c7e70757a7,
				0x31865ced6120f37d,
				0x67fef95d92607890,
				0x5092ef950a16da0b,
				0x9338e69c052b8e7b,
				0x455a4b4cfe30e3f5,
				0x6b02e63195ad0cf8,
				0x6b17b224bad6bf27,
				0xd1e0ccd25bb9c169,
				0xde0c89a556b9ae70,
				0x50065e535a213cf6,
				0x22af003ab672e811,
				0x52e762596bf68235,
				0x9aeba33ac6ecc6b0,
				0x944f6de09134dfb6,
				0x6c47bec883a7de39,
				0x6ad047c430a12104,
				0xa5b1cfdba0ab4067,
				0x7c45d833aff07862,
				0xc0c0f5a60ef4cdcf,
				0xcaf21ecd4377b28c,
				0x57277707199b8175,
				0x506c11b9d90e8b1d,
				0xd83cc2687a19255f,
				0x4a29c6465a314cd1,
				0xed2df21216235097,
				0xb5635c95ff7296e2,
				0xb0774d261cc609db,
				0x443f64ec5a371195,
				0x4112cf68649a260e,
				0xd813f2fab7f5c5ca,
				0x660d3257380841ee,
				0x59ac2c7873f910a3,
				0xe846963877671a17,
				0x93b633abfa3469f8,
				0x6ffe73e81b637fb3,
				0xddf957bc36d8b9ca,
				0x64d0e29eea8838b3,
				0x8dd9bdfd96b9f63,
				0x87e79e5a57d1d13,
				0xe328e230e3e2b3fb,
				0x1c2559e30f0946be,
				0x720bf5f26f4d2eaa,
			},
			{
				0xf1bcc3d275afe51a,
				0xe728e8c83c334074,
				0x96fbf83a12884624,
				0x81a1549fd6573da5,
				0x5fa7867caf35e149,
				0x56986e2ef3ed091b,
				0x917f1dd5f8886c61,
				0xd20d8c88c8ffe65f,
				0x150f361dab9dec26,
				0x9f6a419d382595f4,
				0x64a53dc924fe7ac9,
				0x142de49fff7a7c3d,
				0xc335248857fa9e7,
				0xa9c32d5eae45305,
				0xe6c42178c4bbb92e,
				0x71f1ce2490d20b07,
				0x65fa4f227a2b6d79,
				0xd5f9e858292504d5,
				0xc2b5a03f71471a6f,
				0x59300222b4561e00,
				0xce2f8642ca0712dc,
				0x7ca9723fbb2e8988,
				0x2785338347f2ba08,
				0xc61bb3a141e50e8c,
				0x5e5637885f29bc2b,
				0x7eba726d8c94094b,
				0xa56a5f0bfe39272,
				0xd79476a84ee20d06,
				0x9e4c1269baa4bf37,
				0x17efee45b0dee640,
				0x1d95b0a5fcf90bc6,
				0x93cbe0b699c2585d,
				0x7dc7785b8efdfc80,
				0x8af38731c02ba980,
				0x1fab64ea29a2ddf7,
				0xe4d9429322cd065a,
				0x9da058c67844f20c,
				0x24c0e332b70019b0,
				0x233003b5a6cfe6ad,
				0xd586bd01c5c217f6,
				0xf05d129681949a4c,
				0x964781ce734b3c84,
				0x9c2ed44081ce5fbd,
				0x522e23f3925e319e,
				0x177e00f9fc32f791,
				0x2bc60a63a6f3b3f2,
				0x222bbfae61725606,
				0x486289ddcc3d6780,
				0xf8549e1a3aa5e00d,
				0x7a69afdcc42261a,
				0xc4c118bfe78feaae,
				0xf9f4892ed96bd438,
				0x1af3dbe25d8f45da,
				0xf5b4b0b0d2deeeb4,
				0x962aceefa82e1c84,
				0x46e3ecaaf453ce9,
				0x55b6344cf97aafae,
				0xb862225b055b6960,
				0xcac09afbddd2cdb4,
				0xdaf8e9829fe96b5f,
				0xb5fdfc5d3132c498,
				0x310cb380db6f7503,
				0xe87fbb46217a360e,
				0x2102ae466ebb1148,
			},
		},
		{
			{
				0x42e240cb63689f2f,
				0x6d2bdcdae2919661,
				0x42880b0236e4d951,
				0x5f0f4a5898171bb6,
				0x39f890f579f92f88,
				0x93c5b5f47356388b,
				0x63dc359d8d231b78,
				0xec16ca8aea98ad76,
				0x3253a729b9ba3dde,
				0x8c74c368081b3075,
				0xb9bc6c87167c33e7,
				0x7ef48f2b83024e20,
				0x11d505d4c351bd7f,
				0x6568fca92c76a243,
				0x4de0b0f40f32a7b8,
				0x96d693460cc37e5d,
				0x18727070f1bd400b,
				0x1fcbacd259bf02e7,
				0xd310a7c2ce9b6555,
				0xbf983fe0fe5d8244,
				0x9f74d14f7454a824,
				0x51ebdc4ab9ba3035,
				0x5c82c505db9ab0fa,
				0xfcf7fe8a3430b241,
				0x4c9f34427501b447,
				0x14a68fd73c910841,
				0xa71b9b83461cbd93,
				0x3488b95b0f1850f,
				0x637b2b34ff93c040,
				0x9d1bc9a3dd90a94,
				0x3575668334a1dd3b,
				0x735e2b97a4c45a23,
				0xaa649c6ebcfd50fc,
				0x8dbd98a352afd40b,
				0x87d2074b81d79217,
				0x19f3c751d3e92ae1,
				0xb4ab30f062b19abf,
				0x7b0500ac42047ac4,
				0xc9452ca81a09d85d,
				0x24aa6c514da27500,
				0x5d1a1ae85b49aa1,
				0x679f848f6e8fc971,
				0x7449bbff801fed0b,
				0x7d11cdb1c3b7adf0,
				0x82c7709e781eb7cc,
				0xf3218f1c9510786c,
				0x331478f3af51bbe6,
				0x4bb38de5e7219443,
				0xd7e765d58755c10,
				0x1a083822ceafe02d,
				0x9605d5f0e25ec3b0,
				0xd021ff5cd13a2ed5,
				0x40bdf15d4a672e32,
				0x11355146fd56395,
				0x5db4832046f3d9e5,
				0x239f8b2d7ff719cc,
				0x9d39247e33776d41,
				0x2af7398005aaa5c7,
				0x44db015024623547,
				0x9c15f73e62a76ae2,
				0x75834465489c0c89,
				0x3290ac3a203001bf,
				0xfbbad1f61042279,
				0xe83a908ff2fb60ca,
			},
			{
				0xa4fc4bd4fc5558ca,
				0xe755178d58fc4e76,
				0x69b97db1a4c03dfe,
				0xf9b5b7c4acc67c96,
				0xfc6a82d64b8655fb,
				0x9c684cb6c4d24417,
				0x8ec97d2917456ed0,
				0x6703df9d2924e97e,
				0xd7288e012aeb8d31,
				0xde336a2a4bc1c44b,
				0xbf692b38d079f23,
				0x2c604a7a177326b3,
				0x4850e73e03eb6064,
				0xcfc447f1e53c8e1b,
				0xb05ca3f564268d99,
				0x9ae182c8bc9474e8,
				0x51039ab7712457c3,
				0xc07a3f80c31fb4b4,
				0xb46ee9c5e64a6e7c,
				0xb3819a42abe61c87,
				0x21a007933a522a20,
				0x2df16f761598aa4f,
				0x763c4a1371b368fd,
				0xf793c46702e086a0,
				0x19afe59ae451497f,
				0x52593803dff1e840,
				0xf4f076e65f2ce6f0,
				0x11379625747d5af3,
				0xbce5d2248682c115,
				0x9da4243de836994f,
				0x66f70b33fe09017,
				0x4dc4de189b671a1c,
				0xc5cc1d89724fa456,
				0x5648f680f11a2741,
				0x2d255069f0b7dab3,
				0x9bc5a38ef729abd4,
				0xef2f054308f6a2bc,
				0xaf2042f5cc5c2858,
				0x480412bab7f5be2a,
				0xaef3af4a563dfe43,
				0xa87832d392efee56,
				0x65942c7b3c7e11ae,
				0xded2d633cad004f6,
				0x21f08570f420e565,
				0xb415938d7da94e3c,
				0x91b859e59ecb6350,
				0x10cff333e0ed804a,
				0x28aed140be0bb7dd,
				0x7eed120d54cf2dd9,
				0x22fe545401165f1c,
				0xc91800e98fb99929,
				0x808bd68e6ac10365,
				0xdec468145b7605f6,
				0x1bede3a3aef53302,
				0x43539603d6c55602,
				0xaa969b5c691ccb7a,
				0x56436c9fe1a1aa8d,
				0xefac4b70633b8f81,
				0xbb215798d45df7af,
				0x45f20042f24f1768,
				0x930f80f4e8eb7462,
				0xff6712ffcfd75ea1,
				0xae623fd67468aa70,
				0xdd2c5bc84bc8d8fc,
			},
			{
				0xdc842b7e2819e230,
				0xba89142e007503b8,
				0xa3bc941d0a5061cb,
				0xe9f6760e32cd8021,
				0x9c7e552bc76492f,
				0x852f54934da55cc9,
				0x8107fccf064fcf56,
				0x98954d51fff6580,
				0xe9f6082b05542e4e,
				0xebfafa33d7254b59,
				0x9255abb50d532280,
				0xb9ab4ce57f2d34f3,
				0x693501d628297551,
				0xc62c58f97dd949bf,
				0xcd454f8f19c5126a,
				0xbbe83f4ecc2bdecb,
				0x947ae053ee56e63c,
				0xc8c93882f9475f5f,
				0x3a9bf55ba91f81ca,
				0xd9a11fbb3d9808e4,
				0xfd22063edc29fca,
				0xb3f256d8aca0b0b9,
				0xb03031a8b4516e84,
				0x35dd37d5871448af,
				0x6f423357e7c6a9f9,
				0x325928ee6e6f8794,
				0xd0e4366228b03343,
				0x565c31f7de89ea27,
				0x30f5611484119414,
				0xd873db391292ed4f,
				0x7bd94e1d8e17debc,
				0xc7d9f16864a76e94,
				0x65d34954daf3cebd,
				0xb4b81b3fa97511e2,
				0xb422061193d6f6a7,
				0x71582401c38434d,
				0x7a13f18bbedc4ff5,
				0xbc4097b116c524d2,
				0x59b97885e2f2ea28,
				0x99170a5dc3115544,
				0xd60f6dcedc314222,
				0x56963b0dca418fc0,
				0x16f50edf91e513af,
				0xef1955914b609f93,
				0x565601c0364e3228,
				0xecb53939887e8175,
				0xbac7a9a18531294b,
				0xb344c470397bba52,
				0x37624ae5a48fa6e9,
				0x957baf61700cff4e,
				0x3a6c27934e31188a,
				0xd49503536abca345,
				0x88e049589c432e0,
				0xf943aee7febf21b8,
				0x6c3b8e3e336139d3,
				0x364f6ffa464ee52e,
				0x7f9b6af1ebf78baf,
				0x58627e1a149bba21,
				0x2cd16e2abd791e33,
				0xd363eff5f0977996,
				0xce2a38c344a6eed,
				0x1a804aadb9cfa741,
				0x907f30421d78c5de,
				0x501f65edb3034d07,
			},
			{
				0x26e6db8ffdf5adfe,
				0x469356c504ec9f9d,
				0xc8763c5b08d1908c,
				0x3f6c6af859d80055,
				0x7f7cc39420a3a545,
				0x9bfb227ebdf4c5ce,
				0x89039d79d6fc5c5c,
				0x8fe88b57305e2ab6,
				0x604d51b25fbf70e2,
				0x73aa8a564fb7ac9e,
				0x1a8c1e992b941148,
				0xaac40a2703d9bea0,
				0x764dbeae7fa4f3a6,
				0x1e99b96e70a9be8b,
				0x2c5e9deb57ef4743,
				0x3a938fee32d29981,
				0x9fc10d0f989993e0,
				0xde68a2355b93cae6,
				0xa44cfe79ae538bbe,
				0x9d1d84fcce371425,
				0x51d2b1ab2ddfb636,
				0x2fd7e4b9e72cd38c,
				0x65ca5b96b7552210,
				0xdd69a0d8ab3b546d,
				0x13328503df48229f,
				0xd6bf7baee43cac40,
				0x4838d65f6ef6748f,
				0x1e152328f3318dea,
				0x8f8419a348f296bf,
				0x72c8834a5957b511,
				0xd7a023a73260b45c,
				0x94ebc8abcfb56dae,
				0xa319ce15b0b4db31,
				0x73973751f12dd5e,
				0x8a8e849eb32781a5,
				0xe1925c71285279f5,
				0x74c04bf1790c0efe,
				0x4dda48153c94938a,
				0x9d266d6a1cc0542c,
				0x7440fb816508c4fe,
				0x1b0cab936e65c744,
				0xb559eb1d04e5e932,
				0xc37b45b3f8d6f2ba,
				0xc3a9dc228caac9e9,
				0xf3b8b6675a6507ff,
				0x9fc477de4ed681da,
				0x67378d8eccef96cb,
				0x6dd856d94d259236,
				0xdbc27ab5447822bf,
				0x9b3cdb65f82ca382,
				0xb67b7896167b4c84,
				0xbfced1b0048eac50,
				0xa9119b60369ffebd,
				0x1fff7ac80904bf45,
				0xac12fb171817eee7,
				0xaf08da9177dda93d,
				0xda3a361b1c5157b1,
				0xdcdd7d20903d0c25,
				0x36833336d068f707,
				0xce68341f79893389,
				0xab9090168dd05f34,
				0x43954b3252dc25e5,
				0xb438c2b67f98e5e9,
				0x10dcd78e3851a492,
			},
			{
				0xa4ec0132764ca04b,
				0x733ea705fae4fa77,
				0xb4d8f77bc3e56167,
				0x9e21f4f903b33fd9,
				0x9d765e419fb69f6d,
				0xd30c088ba61ea5ef,
				0x5d94337fbfaf7f5b,
				0x1a4e4822eb4d7a59,
				0x959f587d507a8359,
				0xb063e962e045f54d,
				0x60e8ed72c0dff5d1,
				0x7b64978555326f9f,
				0xfd080d236da814ba,
				0x8c90fd9b083f4558,
				0x106f72fe81e2c590,
				0x7976033a39f7d952,
				0xef02cdd06ffdb432,
				0xa1082c0466df6c0a,
				0x8215e577001332c8,
				0xd39bb9c3a48db6cf,
				0x2738259634305c14,
				0x61cf4f94c97df93d,
				0x1b6baca2ae4e125b,
				0x758f450c88572e0b,
				0xbf84470805e69b5f,
				0x94c3251f06f90cf3,
				0x3e003e616a6591e9,
				0xb925a6cd0421aff3,
				0x61bdd1307c66e300,
				0xbf8d5108e27e0d48,
				0x240ab57a8b888b20,
				0xfc87614baf287e07,
				0xa2ebee47e2fbfce1,
				0xd9f1f30ccd97fb09,
				0xefed53d75fd64e6b,
				0x2e6d02c36017f67f,
				0xa9aa4d20db084e9b,
				0xb64be8d8b25396c1,
				0x70cb6af7c2d5bcf0,
				0x98f076a4f7a2322e,
				0x106c09b972d2e822,
				0x7fba195410e5ca30,
				0x7884d9bc6cb569d8,
				0x647dfedcd894a29,
				0x63573ff03e224774,
				0x4fc8e9560f91b123,
				0x1db956e450275779,
				0xb8d91274b9e9d4fb,
				0x21e0bd5026c619bf,
				0x3b097adaf088f94e,
				0x8d14dedb30be846e,
				0xf95cffa23af5f6f4,
				0x3871700761b3f743,
				0xca672b91e9e4fa16,
				0x64c8e531bff53b55,
				0x241260ed4ad1e87d,
				0x1f837cc7350524,
				0x1877b51e57a764d5,
				0xa2853b80f17f58ee,
				0x993e1de72d36d310,
				0xb3598080ce64a656,
				0x252f59cf0d9f04bb,
				0xd23c8e176d113600,
				0x1bda0492e7e4586e,
			},
			{
				0xcf05daf5ac8d77b0,
				0x49cad48cebf4a71e,
				0x7a4c10ec2158c4a6,
				0xd9e92aa246bf719e,
				0x13ae978d09fe5557,
				0x730499af921549ff,
				0x4e4b705b92903ba4,
				0xff577222c14f0a3a,
				0x963ef2c96b33be31,
				0x74f85198b05a2e7d,
				0x5a0f544dd2b1fb18,
				0x3727073c2e134b1,
				0xc7f6aa2de59aea61,
				0x352787baa0d7c22f,
				0x9853eab63b5e0b35,
				0xabbdcdd7ed5c0860,
				0xf6f7fd1431714200,
				0x30c05b1ba332f41c,
				0x8d2636b81555a786,
				0x46c9feb55d120902,
				0xccec0a73b49c9921,
				0x4e9d2827355fc492,
				0x19ebb029435dcb0f,
				0x4659d2b743848a2c,
				0x4ae7d6a36eb5dbcb,
				0x2d8d5432157064c8,
				0xd1e649de1e7f268b,
				0x8a328a1cedfe552c,
				0x7a3aec79624c7da,
				0x84547ddc3e203c94,
				0x990a98fd5071d263,
				0x1a4ff12616eefc89,
				0xfb152fe3ff26da89,
				0x3e666e6f69ae2c15,
				0x3b544ebe544c19f9,
				0xe805a1e290cf2456,
				0x24b33c9d7ed25117,
				0xe74733427b72f0c1,
				0xa804d18b7097475,
				0x57e3306d881edb4f,
				0x81536d601170fc20,
				0x91b534f885818a06,
				0xec8177f83f900978,
				0x190e714fada5156e,
				0xb592bf39b0364963,
				0x89c350c893ae7dc1,
				0xac042e70f8b383f2,
				0xb49b52e587a1ee60,
				0x5e90277e7cb39e2d,
				0x2c046f22062dc67d,
				0xb10bb459132d0a26,
				0x3fa9ddfb67e2f199,
				0xe09b88e1914f7af,
				0x10e8b35af3eeab37,
				0x9eedeca8e272b933,
				0xd4c718bc4ae8ae5f,
				0x230e343dfba08d33,
				0x43ed7f5a0fae657d,
				0x3a88a0fbbcb05c63,
				0x21874b8b4d2dbc4f,
				0x1bdea12e35f6a8c9,
				0x53c065c6c8e63528,
				0xe34a1d250e7a8d6b,
				0xd6b04d3b7651dd7e,
			},
		},
	},
	Castling: [16]uint64{
		0x0,
		0x31d71dce64b2c310,
		0xf165b587df898190,
		0xc0b2a849bb3b4280,
		0xa57e6339dd2cf3a0,
		0x94a97ef7b99e30b0,
		0x541bd6be02a57230,
		0x65cccb706617b120,
		0x1ef6e6dbb1961ec9,
		0x2f21fb15d524ddd9,
		0xef93535c6e1f9f59,
		0xde444e920aad5c49,
		0xbb8885e26cbaed69,
		0x8a5f982c08082e79,
		0x4aed3065b3336cf9,
		0x7b3a2dabd781afe9,
	},
	EnPassant: [8]uint64{
		0x70cc73d90bc26e24,
		0xe21a6b35df0c3ad7,
		0x3a93d8b2806962,
		0x1c99ded33cb890a1,
		0xcf3145de0add4289,
		0xd0e4427a5514fb72,
		0x77c621cc9fb3a483,
		0x67a34dac4356550b,
	},
	WhiteToMove: 0xf8d626aaaf278509,
}

type BookEntry struct {
	Move   Move
	Weight uint16
}

var Book map[uint64][]BookEntry = map[uint64][]BookEntry{
	0x5f97c294c68444b: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x85afb233adba72ac: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
	},
	0x8dcf2fb0b12bed0f: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x25758f7649e2690e: []BookEntry{
		{Move: Move(0x314), Weight: 5},
	},
	0x87a0d1d1a75869c: []BookEntry{
		{Move: Move(0xb5c), Weight: 65520},
		{Move: Move(0xe73), Weight: 65520},
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x7d431394f8569306: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0xde06cb340bf1b88e: []BookEntry{
		{Move: Move(0xee0), Weight: 1},
	},
	0x2f1826407a645614: []BookEntry{
		{Move: Move(0x18c), Weight: 1},
	},
	0x38c458dc7efdec55: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x5861a1c09e291097: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0x6f2f33e4df80fc4a: []BookEntry{
		{Move: Move(0x218), Weight: 8},
	},
	0xa722b300b45ec85c: []BookEntry{
		{Move: Move(0x144), Weight: 1},
	},
	0xc652cfc0a78f071b: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xd10ebd9b2df803d1: []BookEntry{
		{Move: Move(0xce3), Weight: 2},
	},
	0xf558aee3075cc030: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0xfb282229b4170f5f: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0x329a05c16f5c4ad8: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0x458962f5a3007740: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
	},
	0x8f9b8a9d1cfbf325: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xa801b8eb5a7f3fbe: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0xc42852bfd714b606: []BookEntry{
		{Move: Move(0xee9), Weight: 1},
	},
	0xf44b6961e533d1c4: []BookEntry{
		{Move: Move(0xce3), Weight: 19},
	},
	0xe3352aa01b5ad99: []BookEntry{
		{Move: Move(0xd2c), Weight: 3},
	},
	0x1718576b79b94142: []BookEntry{
		{Move: Move(0x14e), Weight: 1},
	},
	0x9a170e47a0043f60: []BookEntry{
		{Move: Move(0x14c), Weight: 2},
	},
	0xc73069226410c64b: []BookEntry{
		{Move: Move(0xd22), Weight: 1},
	},
	0x5c17a81176f2d5cd: []BookEntry{
		{Move: Move(0xc4), Weight: 1},
	},
	0x6299e56b3cd87c36: []BookEntry{
		{Move: Move(0x14c), Weight: 2},
		{Move: Move(0x6a3), Weight: 1},
	},
	0x6c06b5a3b1ff82b4: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xa0ce40aea07fe762: []BookEntry{
		{Move: Move(0xcaa), Weight: 65520},
	},
	0xc9d06b3543053dfc: []BookEntry{
		{Move: Move(0xe6a), Weight: 3},
	},
	0xc9f0b90f99b2cbf5: []BookEntry{
		{Move: Move(0x210), Weight: 2},
	},
	0xdc8ed81ccb253e43: []BookEntry{
		{Move: Move(0x195), Weight: 3},
		{Move: Move(0x29a), Weight: 3},
		{Move: Move(0x355), Weight: 2},
	},
	0x1c6a4932a7f6bae4: []BookEntry{
		{Move: Move(0x8106), Weight: 3},
	},
	0x56ac9e729bee3259: []BookEntry{
		{Move: Move(0xc69), Weight: 2},
		{Move: Move(0x89a), Weight: 1},
		{Move: Move(0xf74), Weight: 1},
	},
	0x78cda70e17837d9e: []BookEntry{
		{Move: Move(0xf62), Weight: 35280},
		{Move: Move(0xf59), Weight: 65520},
		{Move: Move(0xce3), Weight: 1},
	},
	0x79aa2581574a000a: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0x7d0428cb7a2ae062: []BookEntry{
		{Move: Move(0x210), Weight: 65520},
		{Move: Move(0x31c), Weight: 35280},
	},
	0xd555b109d6a26885: []BookEntry{
		{Move: Move(0x313), Weight: 3},
	},
	0x2a4fd709af581ee6: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x1dbef5bbf88f176f: []BookEntry{
		{Move: Move(0xef3), Weight: 1},
	},
	0x3b1f96130923b26b: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x678fb13aab18b459: []BookEntry{
		{Move: Move(0x2d3), Weight: 2},
	},
	0xd05e0c7ad7178991: []BookEntry{
		{Move: Move(0x90), Weight: 1},
	},
	0x2015ebb94e0d8989: []BookEntry{
		{Move: Move(0xee9), Weight: 1},
	},
	0xa1de6ea2edbf9d8c: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0xae95390715c896eb: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0xd79c024d25772ae6: []BookEntry{
		{Move: Move(0x66a), Weight: 1},
	},
	0xf582a4faa40dca8: []BookEntry{
		{Move: Move(0x3d7), Weight: 2},
	},
	0x6b84fc16fe488df5: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0x714fac11cc1ca77d: []BookEntry{
		{Move: Move(0xc69), Weight: 65520},
		{Move: Move(0xef2), Weight: 2026},
	},
	0x72d4085a66b4293a: []BookEntry{
		{Move: Move(0xa62), Weight: 1},
	},
	0xbd116218627e8e54: []BookEntry{
		{Move: Move(0x7a4), Weight: 1},
	},
	0xc0b9afe525824e3d: []BookEntry{
		{Move: Move(0x195), Weight: 29},
	},
	0xe16bf32cbc1815a8: []BookEntry{
		{Move: Move(0x161), Weight: 1},
	},
	0xf089e06307927b20: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x122224b49e16deb9: []BookEntry{
		{Move: Move(0x2db), Weight: 10},
	},
	0x27541a45a847c590: []BookEntry{
		{Move: Move(0x195), Weight: 22},
	},
	0x32de23faed635534: []BookEntry{
		{Move: Move(0x31c), Weight: 5},
	},
	0x7c1ac1764c9cf0dd: []BookEntry{
		{Move: Move(0xcb), Weight: 1},
	},
	0x99b0f53f7ae359fc: []BookEntry{
		{Move: Move(0xf7c), Weight: 65520},
		{Move: Move(0xceb), Weight: 16380},
	},
	0xb2fb1c80247b3313: []BookEntry{
		{Move: Move(0xb5c), Weight: 1},
	},
	0xf2f4f3f044f4e37d: []BookEntry{
		{Move: Move(0x4b), Weight: 4},
		{Move: Move(0x153), Weight: 4},
		{Move: Move(0x251), Weight: 2},
	},
	0x2e51289d4ce754d: []BookEntry{
		{Move: Move(0xa6), Weight: 1},
	},
	0x7785b00d93d61bff: []BookEntry{
		{Move: Move(0x251), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0xf4acb009aaf2571c: []BookEntry{
		{Move: Move(0xadd), Weight: 1},
	},
	0x69225559f29429db: []BookEntry{
		{Move: Move(0x49c), Weight: 1},
	},
	0x8ec5015e3eeaa43a: []BookEntry{
		{Move: Move(0x2d3), Weight: 3},
	},
	0xa3be1c10e3ed2ab7: []BookEntry{
		{Move: Move(0x51c), Weight: 1},
	},
	0xf0b07a6a42197f2a: []BookEntry{
		{Move: Move(0x829), Weight: 5},
		{Move: Move(0xf3f), Weight: 3},
	},
	0x738e8e18544ad419: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x93e145fd0df4aa4e: []BookEntry{
		{Move: Move(0xea5), Weight: 65519},
		{Move: Move(0xe9e), Weight: 41694},
		{Move: Move(0xd2c), Weight: 11912},
	},
	0xa893b562428c24ca: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0xd56d72226c3094c4: []BookEntry{
		{Move: Move(0x29a), Weight: 9},
	},
	0x34378abed1256d31: []BookEntry{
		{Move: Move(0xfad), Weight: 65520},
		{Move: Move(0xfb4), Weight: 35280},
	},
	0x3cf6bd6a16d8fefd: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x43576b47f62aa166: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x6b56f397b7559569: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0xb3f683986e2ecf81: []BookEntry{
		{Move: Move(0xdb), Weight: 1},
	},
	0xbaf3d3f0f3b8b101: []BookEntry{
		{Move: Move(0xcaa), Weight: 65520},
	},
	0x2e5fe320d56eef57: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
		{Move: Move(0xee9), Weight: 1},
	},
	0x35fc59b714f328fa: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x64d8f414d879981a: []BookEntry{
		{Move: Move(0x46c), Weight: 43680},
		{Move: Move(0xa6), Weight: 65520},
	},
	0x6e9a8ebc2fcf2132: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
	},
	0x81042e9e0c02704d: []BookEntry{
		{Move: Move(0xaa0), Weight: 65520},
		{Move: Move(0xf3f), Weight: 43680},
	},
	0xc97ab689e00bafd1: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xff811291e1619b90: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0x75b9f39e39d5504: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x26c6dcf15b36e99d: []BookEntry{
		{Move: Move(0x6e4), Weight: 1},
	},
	0x296991f1440d5229: []BookEntry{
		{Move: Move(0x691), Weight: 1},
	},
	0x319c5032553d3a89: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x39964b4e429ec62d: []BookEntry{
		{Move: Move(0xd24), Weight: 2},
	},
	0x3bde4b0aa6685f94: []BookEntry{
		{Move: Move(0x3d7), Weight: 3},
	},
	0x4290746a92b9fec7: []BookEntry{
		{Move: Move(0xf7c), Weight: 2},
		{Move: Move(0xc20), Weight: 1},
		{Move: Move(0xc61), Weight: 1},
	},
	0x61e0ef8ef4d8bdd2: []BookEntry{
		{Move: Move(0xf76), Weight: 65520},
	},
	0xacddc811d9f1454: []BookEntry{
		{Move: Move(0x292), Weight: 65520},
		{Move: Move(0x314), Weight: 16380},
	},
	0xe8714ba50389e69: []BookEntry{
		{Move: Move(0x31c), Weight: 1},
	},
	0x17173b7f2799e981: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0x18c287245d9ebe6a: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0x57335827791e45c8: []BookEntry{
		{Move: Move(0x195), Weight: 2},
	},
	0x681ad387583665d8: []BookEntry{
		{Move: Move(0x396), Weight: 1},
	},
	0xd769c6167135a44f: []BookEntry{
		{Move: Move(0xf76), Weight: 3},
	},
	0xe7f093c36dd8fe04: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
		{Move: Move(0xc69), Weight: 11562},
	},
	0x79b955bd0e58511: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x22423723bd4738db: []BookEntry{
		{Move: Move(0x86a), Weight: 1},
	},
	0x524c11fae006ce14: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0x89b0a28f6a8ff408: []BookEntry{
		{Move: Move(0x2d3), Weight: 2},
	},
	0xe26135efb98ee8d2: []BookEntry{
		{Move: Move(0xce3), Weight: 6},
	},
	0xf7345edb85197d6c: []BookEntry{
		{Move: Move(0x54b), Weight: 1},
		{Move: Move(0x55f), Weight: 1},
	},
	0xb0b0d130d9411ed: []BookEntry{
		{Move: Move(0x91b), Weight: 2},
	},
	0x2c6f76c98241cded: []BookEntry{
		{Move: Move(0x6d5), Weight: 1},
		{Move: Move(0xde7), Weight: 1},
	},
	0x5c11bbd3b028b11c: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x6324713ec3cfd18b: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0xef8df0d9bfd41de4: []BookEntry{
		{Move: Move(0x86a), Weight: 2},
	},
	0x2800906e7bc2229f: []BookEntry{
		{Move: Move(0x55b), Weight: 3},
	},
	0x8a6f4d51ae19e145: []BookEntry{
		{Move: Move(0xf3d), Weight: 4},
	},
	0x99d5c32cc267b59a: []BookEntry{
		{Move: Move(0x6a3), Weight: 2},
	},
	0xe69be1144c41857d: []BookEntry{
		{Move: Move(0x564), Weight: 1},
	},
	0x34a8a5435d660432: []BookEntry{
		{Move: Move(0xeb3), Weight: 6},
	},
	0x442a8fa6b2f6d479: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x651841e5468203a0: []BookEntry{
		{Move: Move(0xd24), Weight: 2},
	},
	0x657c4a078b2c94a9: []BookEntry{
		{Move: Move(0xeac), Weight: 2},
	},
	0x7a935cd7c11c9566: []BookEntry{
		{Move: Move(0xae2), Weight: 1},
	},
	0xc25ec1cad719186a: []BookEntry{
		{Move: Move(0x195), Weight: 4},
	},
	0x20c2cc5e683ef587: []BookEntry{
		{Move: Move(0xe6a), Weight: 42},
	},
	0x41a329f49bd71e2a: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0xaaf1c4dad74870e4: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xb4f7ee5f7d4d1e2a: []BookEntry{
		{Move: Move(0x498), Weight: 1},
	},
	0x4344504d71b41cfa: []BookEntry{
		{Move: Move(0x8b4), Weight: 1},
	},
	0x87537ff001ea0094: []BookEntry{
		{Move: Move(0xa19), Weight: 1},
	},
	0xcd9c9630d38763ff: []BookEntry{
		{Move: Move(0x2d3), Weight: 3},
	},
	0xd50bba61f4a1d476: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0xf18bd168473cbdce: []BookEntry{
		{Move: Move(0x9ad), Weight: 1},
		{Move: Move(0x6d1), Weight: 1},
	},
	0xf86a489ed2b06fd6: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x3405761388e5e63: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x3bc73e49c77c8af0: []BookEntry{
		{Move: Move(0xa9b), Weight: 1},
	},
	0x74bbdd6b6e3bb498: []BookEntry{
		{Move: Move(0xf3f), Weight: 6},
	},
	0xaaa8bb0f5c0d75a8: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x3d2210738ad97a76: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x428a83b998d5c62b: []BookEntry{
		{Move: Move(0xc20), Weight: 2},
	},
	0x51a3fa194bbbf436: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xcd0d415d1f10ec71: []BookEntry{
		{Move: Move(0x725), Weight: 1},
	},
	0xcd8983126755021e: []BookEntry{
		{Move: Move(0x6e4), Weight: 1},
	},
	0xf00759eba731c9f1: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0x4e6c5f148ca8cefd: []BookEntry{
		{Move: Move(0x89b), Weight: 4},
		{Move: Move(0xe9e), Weight: 4},
	},
	0x764d587c4be77df5: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0xbe5981819b7d145f: []BookEntry{
		{Move: Move(0x15a), Weight: 16},
	},
	0xdd417f62f8ace911: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xf4cf32ed68a25e71: []BookEntry{
		{Move: Move(0xf3f), Weight: 3},
		{Move: Move(0xee9), Weight: 2},
		{Move: Move(0xe6a), Weight: 2},
	},
	0x1d9302f9ad52bf56: []BookEntry{
		{Move: Move(0x50), Weight: 2},
	},
	0x1e9561bc7718cd16: []BookEntry{
		{Move: Move(0xb5c), Weight: 1},
	},
	0x767badb06d935267: []BookEntry{
		{Move: Move(0xc61), Weight: 8},
	},
	0xa3f6835437ed6b4e: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0xc2a7a8fb1de70a8c: []BookEntry{
		{Move: Move(0xb7c), Weight: 1},
	},
	0xd37b75d4933c5464: []BookEntry{
		{Move: Move(0xe6a), Weight: 6},
	},
	0x83823483714a6e5: []BookEntry{
		{Move: Move(0x8106), Weight: 65520},
		{Move: Move(0x218), Weight: 65520},
		{Move: Move(0x691), Weight: 14560},
	},
	0x11ace0e5d14b00a2: []BookEntry{
		{Move: Move(0x197), Weight: 3},
		{Move: Move(0x195), Weight: 1},
	},
	0x16b3c7cacff62019: []BookEntry{
		{Move: Move(0xb24), Weight: 1},
	},
	0x691de884dcd1bf9d: []BookEntry{
		{Move: Move(0x314), Weight: 12},
	},
	0x8bce053f7a559e81: []BookEntry{
		{Move: Move(0x55f), Weight: 1},
		{Move: Move(0x251), Weight: 1},
	},
	0x9e373c58859c46dc: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xd2887e5f17be84ca: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x542a5939c7a9a028: []BookEntry{
		{Move: Move(0x723), Weight: 1},
	},
	0x5712b06545f75808: []BookEntry{
		{Move: Move(0xc69), Weight: 1},
	},
	0x87a0d66affa948dd: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0x90ced10912cb6f9a: []BookEntry{
		{Move: Move(0x91c), Weight: 1},
	},
	0x997f48123d7b03bb: []BookEntry{
		{Move: Move(0xea5), Weight: 1},
	},
	0xb08d52ffbaff7891: []BookEntry{
		{Move: Move(0xf38), Weight: 1},
	},
	0x4f13883a8cf8a2bb: []BookEntry{
		{Move: Move(0x688), Weight: 1},
		{Move: Move(0x3d7), Weight: 1},
	},
	0x8d102f0986a1c8ba: []BookEntry{
		{Move: Move(0x2db), Weight: 3},
	},
	0xa4e341f7cde82119: []BookEntry{
		{Move: Move(0xea5), Weight: 4},
	},
	0xd69ba88f487beb19: []BookEntry{
		{Move: Move(0x395), Weight: 1},
	},
	0x2dd81a21e529556c: []BookEntry{
		{Move: Move(0x50), Weight: 1},
	},
	0x98858b4776f513bb: []BookEntry{
		{Move: Move(0xe6a), Weight: 3},
	},
	0xba03c86084159f67: []BookEntry{
		{Move: Move(0x674), Weight: 1},
	},
	0xcbd55c2f80249f3f: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
		{Move: Move(0x8106), Weight: 1},
	},
	0xd7d64ec96840301c: []BookEntry{
		{Move: Move(0xae4), Weight: 1},
	},
	0xf9ec78ac04d96377: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x46b8be54463dee2b: []BookEntry{
		{Move: Move(0x55c), Weight: 1},
	},
	0x77b9a466cc98e826: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xb156163904941982: []BookEntry{
		{Move: Move(0xe68), Weight: 1},
		{Move: Move(0xe73), Weight: 1},
	},
	0x28e0deb1710b6e37: []BookEntry{
		{Move: Move(0x2d3), Weight: 8},
		{Move: Move(0x210), Weight: 5},
	},
	0x422bae93a41be9d8: []BookEntry{
		{Move: Move(0x566), Weight: 1},
	},
	0x58aabcc1a7b420d1: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0x7f5def2bd5c14efd: []BookEntry{
		{Move: Move(0x14c), Weight: 3},
		{Move: Move(0x259), Weight: 2},
	},
	0x9110cea63ff02081: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0xa929006c31bc1681: []BookEntry{
		{Move: Move(0xfad), Weight: 3},
	},
	0xbaa80932e7ad04d1: []BookEntry{
		{Move: Move(0xb67), Weight: 1},
	},
	0xbad5fbc00487ea41: []BookEntry{
		{Move: Move(0x564), Weight: 1},
	},
	0x354f6f003a705401: []BookEntry{
		{Move: Move(0x8b), Weight: 6},
		{Move: Move(0xd1), Weight: 3},
	},
	0x62aed443cd21aae0: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
		{Move: Move(0x39e), Weight: 1},
	},
	0xa00c6726dfb996c8: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
		{Move: Move(0x2d3), Weight: 2},
		{Move: Move(0x418), Weight: 1},
	},
	0xa34f95cdcdf59348: []BookEntry{
		{Move: Move(0xdef), Weight: 4},
	},
	0xc00119dca2196b8c: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
	},
	0xc2dff1677290d092: []BookEntry{
		{Move: Move(0x64b), Weight: 1},
	},
	0xd3207fec0612d89d: []BookEntry{
		{Move: Move(0x52), Weight: 65520},
		{Move: Move(0x195), Weight: 7280},
	},
	0xd7a336451d9eedd0: []BookEntry{
		{Move: Move(0x6e4), Weight: 1},
	},
	0x907370a24bda8e7c: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xefefc9a5794577d5: []BookEntry{
		{Move: Move(0x81a), Weight: 1},
	},
	0xffc3eb254822bf9a: []BookEntry{
		{Move: Move(0x54b), Weight: 1},
	},
	0x17d04ff859438394: []BookEntry{
		{Move: Move(0xeac), Weight: 2},
	},
	0x6f9fa4cb29b7b52f: []BookEntry{
		{Move: Move(0x3d6), Weight: 1},
	},
	0xeee545932d916b81: []BookEntry{
		{Move: Move(0x18c), Weight: 2},
	},
	0x37da52b2b7bcb24b: []BookEntry{
		{Move: Move(0x662), Weight: 1},
	},
	0x7b84a13801be94d5: []BookEntry{
		{Move: Move(0x2d3), Weight: 2},
	},
	0xa0d1b3b08f2fe632: []BookEntry{
		{Move: Move(0xdef), Weight: 2},
	},
	0xaf30e8bbfdf42222: []BookEntry{
		{Move: Move(0x6e3), Weight: 1},
	},
	0xd5c1b11883aea207: []BookEntry{
		{Move: Move(0x74f), Weight: 1},
	},
	0xe9696f2c281ba9e7: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x28bbc9c4f83c34e: []BookEntry{
		{Move: Move(0x316), Weight: 1},
	},
	0x1ede3b5a2455e910: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x59f571e97f044205: []BookEntry{
		{Move: Move(0xf7c), Weight: 1},
	},
	0x83e1ea36bb2ae0ab: []BookEntry{
		{Move: Move(0x51b), Weight: 3},
	},
	0x9a4dff27c6780e05: []BookEntry{
		{Move: Move(0xb77), Weight: 2},
		{Move: Move(0xdae), Weight: 1},
		{Move: Move(0xe9e), Weight: 1},
		{Move: Move(0xeac), Weight: 1},
	},
	0xbbf1a0f9936330ae: []BookEntry{
		{Move: Move(0x9ad), Weight: 1},
	},
	0xbf08caac298eadc4: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
	},
	0xc1770b2f2b347640: []BookEntry{
		{Move: Move(0xb5c), Weight: 1},
	},
	0x3d5cdb69606b60e9: []BookEntry{
		{Move: Move(0x2d3), Weight: 7},
		{Move: Move(0x218), Weight: 2},
	},
	0x6a61256b3a8694d4: []BookEntry{
		{Move: Move(0x396), Weight: 18},
	},
	0x78066286c85efa62: []BookEntry{
		{Move: Move(0x292), Weight: 2},
	},
	0xdf49ceec61539386: []BookEntry{
		{Move: Move(0x859), Weight: 1},
	},
	0xeb7de971915d3029: []BookEntry{
		{Move: Move(0x8106), Weight: 4},
	},
	0xebe8615aae5f9dc1: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0xeefe9f2a26e62e76: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0x2dc127fce22bae7d: []BookEntry{
		{Move: Move(0x652), Weight: 1},
	},
	0x4499fd05fcc5f991: []BookEntry{
		{Move: Move(0xd1), Weight: 65520},
		{Move: Move(0x4b), Weight: 35280},
	},
	0x90207f016bc96ddb: []BookEntry{
		{Move: Move(0xcaa), Weight: 3},
	},
	0xa7463ffa3a1761a9: []BookEntry{
		{Move: Move(0xf62), Weight: 1},
	},
	0xbd47d2c6da10e7ea: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xbf6dd67df11dbdeb: []BookEntry{
		{Move: Move(0xf6b), Weight: 2},
		{Move: Move(0xf59), Weight: 1},
	},
	0xfe9404b96c4e3097: []BookEntry{
		{Move: Move(0x161), Weight: 1},
	},
	0x3d5f6734c3794294: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x45e5409cad6ad43e: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
		{Move: Move(0x86a), Weight: 2},
	},
	0x52ddf4c6fbfc14ed: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x6054327bb4c31c6d: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x10771565b5a96f29: []BookEntry{
		{Move: Move(0x73c), Weight: 3},
		{Move: Move(0xceb), Weight: 1},
	},
	0x74dc1408403c5751: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0x8c12bc6a5d7ca47d: []BookEntry{
		{Move: Move(0x9de), Weight: 1},
	},
	0xbc6ffb0d4dbcc7c9: []BookEntry{
		{Move: Move(0xce3), Weight: 5},
		{Move: Move(0xfad), Weight: 2},
	},
	0xbd40f867316bf54a: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
		{Move: Move(0x143), Weight: 1},
	},
	0xcd4731b72ebc25b7: []BookEntry{
		{Move: Move(0x6e4), Weight: 1},
	},
	0xde5a6d98c2c9546a: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x19e7c4b32784b8b9: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x949192ce4b3df2da: []BookEntry{
		{Move: Move(0x89b), Weight: 2},
	},
	0x2bd04f547d36db74: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0x495b878732817c0c: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0xf70a4a4d2f0c9447: []BookEntry{
		{Move: Move(0x6e3), Weight: 1},
	},
	0xf80add6af5b2dcd0: []BookEntry{
		{Move: Move(0x756), Weight: 1},
	},
	0xa51bb7cab6baf388: []BookEntry{
		{Move: Move(0xef2), Weight: 1},
	},
	0xfa4efdd0ee2aa8e7: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0x52e217d309ce67ed: []BookEntry{
		{Move: Move(0x8db), Weight: 1},
	},
	0x6a3d9abfc90ee7ec: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0xd0852cd18aa5813f: []BookEntry{
		{Move: Move(0x218), Weight: 65520},
		{Move: Move(0x6e2), Weight: 43680},
	},
	0xe3b7b2e34ef352bb: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x15c6e20a4df861ac: []BookEntry{
		{Move: Move(0xc61), Weight: 1},
	},
	0x8ab683db52919e68: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0xdd4398afb4bb9841: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x1dc3916e6386368f: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x3d965402f9b8ce1f: []BookEntry{
		{Move: Move(0x899), Weight: 1},
	},
	0x6284639ab0a605f5: []BookEntry{
		{Move: Move(0xeac), Weight: 2},
	},
	0x6ffda0bcb060308f: []BookEntry{
		{Move: Move(0x9ee), Weight: 1},
	},
	0xdb5bdc884a393d46: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0xff50f26a80faa9d8: []BookEntry{
		{Move: Move(0xe39), Weight: 1},
	},
	0x2a6ec30017e3fb66: []BookEntry{
		{Move: Move(0xd2b), Weight: 1},
	},
	0x40600176c67bafad: []BookEntry{
		{Move: Move(0xe6a), Weight: 17},
	},
	0x692e5d16a508ff62: []BookEntry{
		{Move: Move(0x89), Weight: 10},
	},
	0x9153533974f0e91e: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xa9a53b4655879bcb: []BookEntry{
		{Move: Move(0x564), Weight: 2},
	},
	0x44b0ade0df2efdd1: []BookEntry{
		{Move: Move(0x8ab), Weight: 1},
	},
	0x83660352368e2f4b: []BookEntry{
		{Move: Move(0x8e9), Weight: 16380},
		{Move: Move(0xef3), Weight: 65520},
	},
	0xeb4417dc05acb3b6: []BookEntry{
		{Move: Move(0x89), Weight: 65520},
		{Move: Move(0x210), Weight: 43680},
	},
	0xefed90774bab7787: []BookEntry{
		{Move: Move(0x5de), Weight: 1},
	},
	0x58a9c7bf47ad9e4f: []BookEntry{
		{Move: Move(0xf3f), Weight: 6},
	},
	0x2a70af792e3a89b2: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x10bb5f9522be04f: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0x1c315ca00ce89e0d: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x7fb7a7765488cb7a: []BookEntry{
		{Move: Move(0x153), Weight: 3},
		{Move: Move(0x9d), Weight: 1},
	},
	0xb8d869eafd0fc3d5: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x48b8ee122294c120: []BookEntry{
		{Move: Move(0x832), Weight: 1},
		{Move: Move(0x829), Weight: 1},
	},
	0x591df54aace32895: []BookEntry{
		{Move: Move(0x9dd), Weight: 1},
	},
	0x7a344064de1006e5: []BookEntry{
		{Move: Move(0x3d7), Weight: 2},
		{Move: Move(0x48c), Weight: 1},
	},
	0x8ca93e3cd524e2a5: []BookEntry{
		{Move: Move(0xef3), Weight: 2},
	},
	0x8398719e213d547: []BookEntry{
		{Move: Move(0x49c), Weight: 1},
	},
	0x62752c503b712ba1: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0x78617cfc935a5eb3: []BookEntry{
		{Move: Move(0x8106), Weight: 4},
	},
	0x91867a23179995f0: []BookEntry{
		{Move: Move(0x144), Weight: 1},
	},
	0xdf08f5f5a95178b4: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xe095dc711d67250c: []BookEntry{
		{Move: Move(0xca2), Weight: 65520},
	},
	0xaee59e80e04e3336: []BookEntry{
		{Move: Move(0xf74), Weight: 65520},
		{Move: Move(0xf6b), Weight: 65520},
	},
	0x24aa98400a3ef7fc: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x32c94c9bb070395d: []BookEntry{
		{Move: Move(0x210), Weight: 4},
		{Move: Move(0x52), Weight: 1},
	},
	0xda5a5aebff2d2c1a: []BookEntry{
		{Move: Move(0x54b), Weight: 1},
	},
	0x22af5482d7b4867f: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x243be3cbf8245e46: []BookEntry{
		{Move: Move(0xeac), Weight: 2},
	},
	0x26270fa2218d9064: []BookEntry{
		{Move: Move(0x8db), Weight: 65520},
	},
	0x465d9dfa78f9703d: []BookEntry{
		{Move: Move(0x3d7), Weight: 3},
	},
	0x4b50d1be5fe02328: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x7edc683594968df4: []BookEntry{
		{Move: Move(0xf7c), Weight: 1},
	},
	0xcfbf401533796f68: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0x995091bf57faf895: []BookEntry{
		{Move: Move(0xaa4), Weight: 3},
	},
	0x35b04511761851b3: []BookEntry{
		{Move: Move(0xf7c), Weight: 1},
		{Move: Move(0xdef), Weight: 1},
	},
	0x3ecc074f29ab22ea: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
	},
	0x6b965722689ff80e: []BookEntry{
		{Move: Move(0x6ea), Weight: 1},
	},
	0x8ec0089cc883d616: []BookEntry{
		{Move: Move(0xd5), Weight: 1},
	},
	0xdc7ccbf8a370b31c: []BookEntry{
		{Move: Move(0x41a), Weight: 1},
	},
	0x1c12aa7c739acd6: []BookEntry{
		{Move: Move(0x31c), Weight: 2},
	},
	0x38461b1c244f175: []BookEntry{
		{Move: Move(0x2db), Weight: 65520},
		{Move: Move(0x314), Weight: 38220},
		{Move: Move(0x292), Weight: 5460},
	},
	0x1f5cd1c0a1b9804c: []BookEntry{
		{Move: Move(0x6e4), Weight: 1},
	},
	0x45284ee920e59f57: []BookEntry{
		{Move: Move(0x259), Weight: 2},
	},
	0x48ce5f26c2d235cf: []BookEntry{
		{Move: Move(0x91b), Weight: 1},
	},
	0x4d1402f03b07c41f: []BookEntry{
		{Move: Move(0x210), Weight: 1},
		{Move: Move(0x314), Weight: 65520},
	},
	0x7a1c93c9e334efd9: []BookEntry{
		{Move: Move(0xab9), Weight: 1},
	},
	0x7b2eab40813096dc: []BookEntry{
		{Move: Move(0xd24), Weight: 4},
		{Move: Move(0xca2), Weight: 3},
	},
	0x4090319380cc952a: []BookEntry{
		{Move: Move(0x96e), Weight: 3},
	},
	0x6ad8f4f1ba66d7c0: []BookEntry{
		{Move: Move(0x144), Weight: 1},
	},
	0x9dd67d7b7063ddf4: []BookEntry{
		{Move: Move(0xceb), Weight: 2},
	},
	0xe69bb237e989b4ac: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xb6db635d4b90078b: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x14c47eea7b218ed5: []BookEntry{
		{Move: Move(0xf59), Weight: 65520},
		{Move: Move(0xce3), Weight: 32760},
		{Move: Move(0xe6a), Weight: 10920},
	},
	0x986208c42984ef6: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x5de5d974f1078d9c: []BookEntry{
		{Move: Move(0xf76), Weight: 6},
	},
	0x78f6091060d353e1: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x26dcbc8b3b5e455f: []BookEntry{
		{Move: Move(0x8106), Weight: 5},
		{Move: Move(0x6e2), Weight: 1},
	},
	0x10d870d7ca52f8e: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x659e3d42813ff15a: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
	},
	0xab416883f233239f: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xea7c8f69bf9e29d9: []BookEntry{
		{Move: Move(0xd2c), Weight: 65520},
		{Move: Move(0xee9), Weight: 11562},
	},
	0xf3b84f8433b6695e: []BookEntry{
		{Move: Move(0xe73), Weight: 3},
		{Move: Move(0x7a7), Weight: 1},
	},
	0xf953764b053b640e: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x3461b6cafb7ede89: []BookEntry{
		{Move: Move(0xc6a), Weight: 65520},
	},
	0x5fbac772273a7cd5: []BookEntry{
		{Move: Move(0x251), Weight: 2},
	},
	0xbcf06a108c5e35d9: []BookEntry{
		{Move: Move(0x2d3), Weight: 3},
	},
	0x9c997b83cc779314: []BookEntry{
		{Move: Move(0xf62), Weight: 3},
	},
	0xae2d454607cc9259: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0x3adfb3bc92992a3: []BookEntry{
		{Move: Move(0x8b), Weight: 1},
		{Move: Move(0x3d7), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0x35e3f63463af0d94: []BookEntry{
		{Move: Move(0x14c), Weight: 2},
	},
	0x7acd0315055d8e99: []BookEntry{
		{Move: Move(0xd6d), Weight: 1},
	},
	0x42276d9ef110eaf: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
		{Move: Move(0x89), Weight: 1},
	},
	0xe164f127142bc91: []BookEntry{
		{Move: Move(0xe9e), Weight: 65520},
	},
	0x8ff4fc4777b03b79: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0xc0eb36fbcbbac59d: []BookEntry{
		{Move: Move(0xa9b), Weight: 1},
	},
	0x7ef7bd39a3ba5699: []BookEntry{
		{Move: Move(0x8da), Weight: 1},
	},
	0xaae71d9b25f9ce4d: []BookEntry{
		{Move: Move(0x8106), Weight: 4},
	},
	0xc97ea6ebde44f19a: []BookEntry{
		{Move: Move(0x153), Weight: 2},
		{Move: Move(0x89), Weight: 1},
	},
	0xcfce747c640fcc4b: []BookEntry{
		{Move: Move(0xb23), Weight: 1},
	},
	0xf8649a522dfcff27: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0x88972ca443708f5: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0x3a695dff0023cb58: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0x83db41833f9d1c75: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0x8e5c638c6c0e2368: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xf77c2081e9cf1854: []BookEntry{
		{Move: Move(0xf62), Weight: 65520},
	},
	0x1149fe940a601f5: []BookEntry{
		{Move: Move(0x15a), Weight: 1},
	},
	0x3853e67347e42e21: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x38ec8a1702994133: []BookEntry{
		{Move: Move(0xe39), Weight: 1},
	},
	0x3c519aa1dc24813a: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0x2f13a74c6959292d: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x515cc85312c9aae7: []BookEntry{
		{Move: Move(0x314), Weight: 2},
	},
	0x6997b210a48cff30: []BookEntry{
		{Move: Move(0x49b), Weight: 2},
	},
	0x7f31ca2e4b055e01: []BookEntry{
		{Move: Move(0xda6), Weight: 1},
		{Move: Move(0xfad), Weight: 1},
	},
	0xbec67375f4f5caa7: []BookEntry{
		{Move: Move(0x4db), Weight: 1},
		{Move: Move(0x48c), Weight: 1},
	},
	0xc68b4ff3ada6ebca: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x20411bcbc8ca3ca5: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xb0ddb79046382cd6: []BookEntry{
		{Move: Move(0x18c), Weight: 1},
	},
	0xb76c1a21fd685886: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
	},
	0xbf3397a6eb4e0d88: []BookEntry{
		{Move: Move(0x7e5), Weight: 1},
	},
	0xc292b448666f4632: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0xeb2888df66550df3: []BookEntry{
		{Move: Move(0xce3), Weight: 14},
		{Move: Move(0xca2), Weight: 9},
		{Move: Move(0xdae), Weight: 1},
	},
	0x2943fd55f54613e5: []BookEntry{
		{Move: Move(0x94), Weight: 65520},
		{Move: Move(0x3d7), Weight: 65520},
	},
	0x85bd53e0cb861a2d: []BookEntry{
		{Move: Move(0x6e2), Weight: 65520},
	},
	0xc9c6c2594c215257: []BookEntry{
		{Move: Move(0x9d), Weight: 6},
	},
	0xfe769c7277d4b973: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0x641ea3e0acece50a: []BookEntry{
		{Move: Move(0xea5), Weight: 3},
		{Move: Move(0xeac), Weight: 2},
	},
	0x9871f9a72f0bc53e: []BookEntry{
		{Move: Move(0x6e3), Weight: 1},
	},
	0xb2d40edb7c654174: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0xbd6a96a86396d2af: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0xe16c056913734292: []BookEntry{
		{Move: Move(0xf38), Weight: 1},
	},
	0xed4a8884e65aac2c: []BookEntry{
		{Move: Move(0xe39), Weight: 1},
		{Move: Move(0xf74), Weight: 1},
	},
	0x281732e022e316db: []BookEntry{
		{Move: Move(0x795), Weight: 2},
	},
	0x29590a86ff2a2b09: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x52353e8b7d42fe1d: []BookEntry{
		{Move: Move(0xaea), Weight: 1},
	},
	0x58852f93a367bc52: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0xcddb56ff0eb99a6d: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0xdc1ecfc9a72d9c49: []BookEntry{
		{Move: Move(0x8b), Weight: 1},
	},
	0xfc2056ac309264fd: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x6ce5911928b0cce: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0x28bc8407b9397c1d: []BookEntry{
		{Move: Move(0x259), Weight: 2},
	},
	0x7ac30597d2ae4e56: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
	},
	0x9985e8f7112b2514: []BookEntry{
		{Move: Move(0x144), Weight: 1},
	},
	0xa57d12d9dbbfc345: []BookEntry{
		{Move: Move(0x355), Weight: 1},
	},
	0xcf4092fcb18301ee: []BookEntry{
		{Move: Move(0xf6b), Weight: 3},
	},
	0xd2a23215fd3300c2: []BookEntry{
		{Move: Move(0xf76), Weight: 1},
	},
	0xedd80ced9b3dcd21: []BookEntry{
		{Move: Move(0x35d), Weight: 2},
	},
	0x73a4a15e9365b5f9: []BookEntry{
		{Move: Move(0x795), Weight: 21840},
		{Move: Move(0xa9b), Weight: 65520},
	},
	0xf1b7c7f416d549fe: []BookEntry{
		{Move: Move(0x5cd), Weight: 1},
	},
	0xa911540a2a2f9d65: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0xdf07eb60779140f0: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
	},
	0xec526a2e6ee08bb6: []BookEntry{
		{Move: Move(0xdef), Weight: 2},
		{Move: Move(0xce3), Weight: 1},
	},
	0xb08c54249741a0d: []BookEntry{
		{Move: Move(0xaa4), Weight: 65520},
	},
	0x551eae71a1e8f501: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0x7630236b808e038f: []BookEntry{
		{Move: Move(0x49a), Weight: 1},
	},
	0xe9becd1bc91cd1ff: []BookEntry{
		{Move: Move(0x8d9), Weight: 1},
	},
	0xcb023a8fe8d00b10: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
		{Move: Move(0x3d7), Weight: 1},
		{Move: Move(0x688), Weight: 1},
	},
	0xf04c2aeabcf6135a: []BookEntry{
		{Move: Move(0xf6b), Weight: 5},
	},
	0xf8bb5e6e816fc7a1: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x3d1b21f17e3b0b29: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0xa2969754b978af65: []BookEntry{
		{Move: Move(0x853), Weight: 1},
	},
	0xc095a91913e599c7: []BookEntry{
		{Move: Move(0xef3), Weight: 1},
	},
	0xd2099d7656a16e65: []BookEntry{
		{Move: Move(0xe9e), Weight: 1},
	},
	0xe3d336edb62cc9b3: []BookEntry{
		{Move: Move(0xca), Weight: 2},
	},
	0x1674f1eb0a168f27: []BookEntry{
		{Move: Move(0x210), Weight: 1},
		{Move: Move(0x14c), Weight: 1},
	},
	0x3748e3a13d6d165a: []BookEntry{
		{Move: Move(0xf76), Weight: 7},
		{Move: Move(0xe6a), Weight: 1},
	},
	0x6304fcaa0cd7c3c0: []BookEntry{
		{Move: Move(0xaa4), Weight: 6},
	},
	0x6c51bb403eb34229: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0xb94f3037f550a6fe: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xc24a345ff2145413: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x12ccfe8649ac4cc9: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x1313adedac381273: []BookEntry{
		{Move: Move(0x4b), Weight: 3},
	},
	0x3c0af67234f697d8: []BookEntry{
		{Move: Move(0xfad), Weight: 2},
		{Move: Move(0xd2c), Weight: 1},
	},
	0x8162ff8b2f778ab8: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0xdd651d79e16b7995: []BookEntry{
		{Move: Move(0x14e), Weight: 28080},
		{Move: Move(0x4b), Weight: 65520},
	},
	0xfa184b9eb1fa30f9: []BookEntry{
		{Move: Move(0xd23), Weight: 65520},
	},
	0x67ba9ce28e9fa7: []BookEntry{
		{Move: Move(0xf7c), Weight: 5},
		{Move: Move(0xe6a), Weight: 1},
	},
	0x12cd50c88a23d342: []BookEntry{
		{Move: Move(0x396), Weight: 2},
	},
	0x34b3cd92e49484b4: []BookEntry{
		{Move: Move(0xceb), Weight: 2},
	},
	0x4751cddb9bc1a17d: []BookEntry{
		{Move: Move(0x396), Weight: 53607},
		{Move: Move(0x4b), Weight: 65519},
	},
	0x615fda7cd01e4e86: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0x724dec9a595a77bc: []BookEntry{
		{Move: Move(0x4b), Weight: 3},
	},
	0x9471e2658862c51a: []BookEntry{
		{Move: Move(0x8e9), Weight: 1},
	},
	0xa3f5fc71e0bc06f9: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0x28ef662e70606e1d: []BookEntry{
		{Move: Move(0xf3f), Weight: 3},
	},
	0x9742626ad74e489f: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0x526341290d7773ff: []BookEntry{
		{Move: Move(0x51b), Weight: 3},
	},
	0x92327ce9f7c59150: []BookEntry{
		{Move: Move(0xf62), Weight: 1},
	},
	0xfea6f026a9670b3c: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x26af4abd3d0f0d27: []BookEntry{
		{Move: Move(0x49a), Weight: 1},
	},
	0x56da158d0a5b5035: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0x650ab38766d6f3b2: []BookEntry{
		{Move: Move(0xf3f), Weight: 7},
	},
	0xc901645a04f904c5: []BookEntry{
		{Move: Move(0x51c), Weight: 1},
	},
	0x3e6d5488cf2ddb2e: []BookEntry{
		{Move: Move(0x259), Weight: 1},
	},
	0xa508d2a42206e078: []BookEntry{
		{Move: Move(0x91c), Weight: 1},
	},
	0xabda7de3de3b720b: []BookEntry{
		{Move: Move(0x8106), Weight: 65520},
	},
	0xcdd8b88b2019e9a8: []BookEntry{
		{Move: Move(0x8ed), Weight: 1},
	},
	0x449baf763270dad1: []BookEntry{
		{Move: Move(0x14e), Weight: 1},
	},
	0x712e3d8d8a444c2e: []BookEntry{
		{Move: Move(0x8ec), Weight: 1},
	},
	0xd96de15027d7baa3: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0xea6077dfc89d7f98: []BookEntry{
		{Move: Move(0x6ea), Weight: 1},
		{Move: Move(0x29a), Weight: 1},
	},
	0x269000e641748d1f: []BookEntry{
		{Move: Move(0xad0), Weight: 1},
	},
	0x38ec7b8bc53d98c9: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0x97fd90724ef20b4b: []BookEntry{
		{Move: Move(0x91b), Weight: 5},
	},
	0xbbd5737fdc8af8f9: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0x1aa25dc38f49d250: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0x33b3ba68f8a65fb3: []BookEntry{
		{Move: Move(0x9ee), Weight: 1},
	},
	0x39b3e62fe28c4b6a: []BookEntry{
		{Move: Move(0xb24), Weight: 1},
	},
	0x8024fecf0224cb5c: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x9ab56b0fea452fc5: []BookEntry{
		{Move: Move(0x4b), Weight: 11},
	},
	0xdb216a684541afb7: []BookEntry{
		{Move: Move(0x6d2), Weight: 1},
	},
	0xe2ea2f14e2868734: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
	},
	0xe6a218cf8b6ec8aa: []BookEntry{
		{Move: Move(0x210), Weight: 65520},
	},
	0x7ea9b151ff9560f: []BookEntry{
		{Move: Move(0xeb3), Weight: 3},
		{Move: Move(0xf6b), Weight: 1},
	},
	0xd13f4d3679a8c3d4: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xe1e0c4aad071bc20: []BookEntry{
		{Move: Move(0x899), Weight: 1},
	},
	0xfad3c64a5cfba120: []BookEntry{
		{Move: Move(0xeb3), Weight: 65520},
		{Move: Move(0x92a), Weight: 28080},
	},
	0x596b02649c3fa6: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x8e4c2dc304518d3: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
		{Move: Move(0x89b), Weight: 1},
	},
	0x531dc43c09bfcd7f: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x71219b7990731cb1: []BookEntry{
		{Move: Move(0x6ca), Weight: 1},
	},
	0xb0db872a19c9540d: []BookEntry{
		{Move: Move(0x14e), Weight: 6},
	},
	0x5ab693506f2c1c69: []BookEntry{
		{Move: Move(0x89b), Weight: 3},
		{Move: Move(0xeac), Weight: 1},
	},
	0xe09125b1504e3578: []BookEntry{
		{Move: Move(0xc61), Weight: 1},
	},
	0x18d6706eb670dfca: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
	},
	0x2b4f5fb8f2edd4aa: []BookEntry{
		{Move: Move(0xd2e), Weight: 1},
	},
	0x55cb3e7b3758516f: []BookEntry{
		{Move: Move(0x724), Weight: 1},
	},
	0x64904957c58773e0: []BookEntry{
		{Move: Move(0xc28), Weight: 16380},
		{Move: Move(0xc20), Weight: 65520},
	},
	0x6baae14beea9199b: []BookEntry{
		{Move: Move(0xea5), Weight: 6},
	},
	0x8a9f96657bbbbae2: []BookEntry{
		{Move: Move(0xcfd), Weight: 1},
	},
	0x92d2550d34902fa1: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0x9371581625eb33f4: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0x30d7e7e1ec3f4635: []BookEntry{
		{Move: Move(0xdae), Weight: 1},
	},
	0x99b6f1470dc1dfa5: []BookEntry{
		{Move: Move(0xcaa), Weight: 2},
	},
	0xa52cf8196dedd67d: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0xadde4d0e9d829fd2: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0xb3bf2287e2c15167: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0xd4a4bbd662bb0e7c: []BookEntry{
		{Move: Move(0x91b), Weight: 4},
	},
	0xf0acce9a7b657aeb: []BookEntry{
		{Move: Move(0x8b2), Weight: 65520},
	},
	0x512e65b8eca3b513: []BookEntry{
		{Move: Move(0xb5c), Weight: 1},
		{Move: Move(0xe6a), Weight: 1},
	},
	0x595c86e4b023d726: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x9416c7dea1ac7880: []BookEntry{
		{Move: Move(0x6e3), Weight: 8},
		{Move: Move(0x292), Weight: 1},
	},
	0xabb5df230f89d8c5: []BookEntry{
		{Move: Move(0x18c), Weight: 3},
	},
	0xd6fe15bc3a503748: []BookEntry{
		{Move: Move(0xee0), Weight: 1},
	},
	0xe5cc09ccda6786b2: []BookEntry{
		{Move: Move(0xceb), Weight: 2},
	},
	0xf7cdab8795680f7a: []BookEntry{
		{Move: Move(0xc28), Weight: 1},
	},
	0x72c9914b4603b758: []BookEntry{
		{Move: Move(0xeac), Weight: 2},
		{Move: Move(0xfad), Weight: 1},
		{Move: Move(0xe9e), Weight: 1},
	},
	0x9a4d335e8803c39d: []BookEntry{
		{Move: Move(0xd8), Weight: 1},
		{Move: Move(0x314), Weight: 1},
	},
	0xba031a7d2fe463ea: []BookEntry{
		{Move: Move(0x396), Weight: 65520},
	},
	0xbf4a4561358c0f69: []BookEntry{
		{Move: Move(0x4b), Weight: 2},
		{Move: Move(0x8106), Weight: 1},
	},
	0x330bb70e42e06fa1: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0x64170fa612e6224f: []BookEntry{
		{Move: Move(0xf7c), Weight: 1},
	},
	0x9c48bf34f304138d: []BookEntry{
		{Move: Move(0x8d9), Weight: 1},
		{Move: Move(0x8e9), Weight: 1},
	},
	0xc8ae18693763dc9b: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0x4516d155e3dc8f2a: []BookEntry{
		{Move: Move(0xeeb), Weight: 1},
	},
	0x55bb12cbbd7e7186: []BookEntry{
		{Move: Move(0x3d7), Weight: 9},
	},
	0x8c089b7c48341813: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0xbf29a6086ab02bd6: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
		{Move: Move(0x52), Weight: 25200},
		{Move: Move(0x314), Weight: 10080},
	},
	0x143db01855ea7960: []BookEntry{
		{Move: Move(0x4a3), Weight: 1},
	},
	0x188a8d6f6e51fa68: []BookEntry{
		{Move: Move(0x210), Weight: 65520},
		{Move: Move(0x9d), Weight: 16380},
	},
	0x3b63f0c12325544d: []BookEntry{
		{Move: Move(0xb67), Weight: 65520},
	},
	0x58281111cd4c7f05: []BookEntry{
		{Move: Move(0x314), Weight: 2},
	},
	0x676902f53cbe43a1: []BookEntry{
		{Move: Move(0x31c), Weight: 2},
	},
	0x9f0720cc980827ea: []BookEntry{
		{Move: Move(0xd5), Weight: 2},
	},
	0xe525b1311587325: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x9e42f8939e85e3b3: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x791af3596fc45e0f: []BookEntry{
		{Move: Move(0x94), Weight: 1},
		{Move: Move(0x14c), Weight: 1},
	},
	0x7f521d1d34c31a2b: []BookEntry{
		{Move: Move(0xe9e), Weight: 4},
	},
	0xbd7741389222a904: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0xd0e63bb7170cc987: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xe0f95a88688bb5f6: []BookEntry{
		{Move: Move(0xeb1), Weight: 6},
		{Move: Move(0xef4), Weight: 1},
	},
	0xee2b83f65296353c: []BookEntry{
		{Move: Move(0xfad), Weight: 9},
	},
	0xad48b8339b0282a2: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xb51f5d4190a72f53: []BookEntry{
		{Move: Move(0x89a), Weight: 1},
	},
	0xfa772b8ec0e3c6ab: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xfc30f6b3c0a1a4ad: []BookEntry{
		{Move: Move(0xe6a), Weight: 65520},
		{Move: Move(0x89b), Weight: 14040},
		{Move: Move(0xfad), Weight: 14040},
	},
	0x1d6ee74dfac2d79e: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0xe9358625bc1cf0da: []BookEntry{
		{Move: Move(0xcaa), Weight: 8},
	},
	0x66c2a31d66cab3b: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0x27c3705280cb4361: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x371992e68c1692f5: []BookEntry{
		{Move: Move(0x2d2), Weight: 1},
	},
	0xb41a7e8c9f735244: []BookEntry{
		{Move: Move(0x2d3), Weight: 2},
		{Move: Move(0x86a), Weight: 1},
		{Move: Move(0x724), Weight: 1},
	},
	0x271362abe13c0a4e: []BookEntry{
		{Move: Move(0xefc), Weight: 2},
		{Move: Move(0xca2), Weight: 1},
		{Move: Move(0xea5), Weight: 1},
	},
	0x325b696440075d7d: []BookEntry{
		{Move: Move(0x2d3), Weight: 3},
	},
	0xc7ce339e8682d5cb: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0xfbb2dba2cd66050b: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x1119906ae1976389: []BookEntry{
		{Move: Move(0xa6), Weight: 65520},
		{Move: Move(0x9d), Weight: 1},
	},
	0x1143abc12895c3a1: []BookEntry{
		{Move: Move(0x2db), Weight: 2},
	},
	0x1b75af6e647f6522: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
	},
	0x8c1f88acbf209377: []BookEntry{
		{Move: Move(0x8e4), Weight: 2},
	},
	0xd57359159b18bf6f: []BookEntry{
		{Move: Move(0x2db), Weight: 17},
	},
	0xf2b09157f03b82db: []BookEntry{
		{Move: Move(0xc28), Weight: 2},
		{Move: Move(0xf6b), Weight: 1},
	},
	0x2a20585e7fa5ae99: []BookEntry{
		{Move: Move(0x652), Weight: 1},
	},
	0x48b5a37c9494402d: []BookEntry{
		{Move: Move(0x8106), Weight: 9},
	},
	0x7f34438982042268: []BookEntry{
		{Move: Move(0x96c), Weight: 1},
	},
	0x807949893bc2e696: []BookEntry{
		{Move: Move(0xee0), Weight: 1},
	},
	0xaff3d9389d0a18be: []BookEntry{
		{Move: Move(0xce3), Weight: 65520},
	},
	0xbc705657f9554741: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0xf4fc3281bfbce85a: []BookEntry{
		{Move: Move(0xef2), Weight: 4},
	},
	0x179deb873c0ce49c: []BookEntry{
		{Move: Move(0xb23), Weight: 1},
	},
	0x201e0f35a8ca00a3: []BookEntry{
		{Move: Move(0x8b), Weight: 1},
	},
	0x3455e93c6c4e3076: []BookEntry{
		{Move: Move(0xc20), Weight: 65519},
		{Move: Move(0xca2), Weight: 53607},
	},
	0x40d4272132657865: []BookEntry{
		{Move: Move(0x1), Weight: 1},
	},
	0x8a8c1378d77d22b7: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xe3addba97697978e: []BookEntry{
		{Move: Move(0xc29), Weight: 2},
	},
	0x104ddfd46c2d411b: []BookEntry{
		{Move: Move(0xc20), Weight: 65520},
		{Move: Move(0xf3f), Weight: 35280},
	},
	0x49e809d567e8aa73: []BookEntry{
		{Move: Move(0x51d), Weight: 1},
	},
	0x7794e432507f6367: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0xbb5dc2b36938dfba: []BookEntry{
		{Move: Move(0x14e), Weight: 1},
	},
	0x132d92fb5b8dabdb: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x69e9dd003274d78b: []BookEntry{
		{Move: Move(0x3d7), Weight: 2},
	},
	0x71d71e46f6f403ea: []BookEntry{
		{Move: Move(0x218), Weight: 1},
	},
	0x7fcb5af388625a57: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0x9d5f7aee7e779da1: []BookEntry{
		{Move: Move(0x2db), Weight: 65520},
		{Move: Move(0x314), Weight: 16380},
		{Move: Move(0x29a), Weight: 16380},
		{Move: Move(0x195), Weight: 65520},
	},
	0xbbad2b7422e66aef: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xbffb91f4f1e894d5: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0xd20ff64c86031026: []BookEntry{
		{Move: Move(0x564), Weight: 7},
	},
	0x6e883960a31e749b: []BookEntry{
		{Move: Move(0x564), Weight: 3},
	},
	0x72394567ea4d1c47: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x9f61119f554c7267: []BookEntry{
		{Move: Move(0x6c3), Weight: 1},
	},
	0x3b96017e60fbe955: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
		{Move: Move(0x6a3), Weight: 1},
	},
	0x84ff1cf0c3dd665f: []BookEntry{
		{Move: Move(0xf76), Weight: 1},
	},
	0xad110ef800c0d2ef: []BookEntry{
		{Move: Move(0x409), Weight: 65520},
	},
	0xd945738bf9888023: []BookEntry{
		{Move: Move(0xde7), Weight: 1},
	},
	0xe4c0103b96b6f4d5: []BookEntry{
		{Move: Move(0xf3f), Weight: 11},
	},
	0x5579e215c6d29e5: []BookEntry{
		{Move: Move(0xda6), Weight: 1},
	},
	0x5cdaebee74595e7: []BookEntry{
		{Move: Move(0xae4), Weight: 1},
	},
	0x311ffac4bd326e49: []BookEntry{
		{Move: Move(0x3d6), Weight: 1},
	},
	0x71cb276625ad10b2: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0xeb4d617e2bca3701: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x7ac41b3387eb6f7: []BookEntry{
		{Move: Move(0x89), Weight: 3},
		{Move: Move(0x8106), Weight: 1},
	},
	0xe4c85c115fd6440e: []BookEntry{
		{Move: Move(0x153), Weight: 53607},
		{Move: Move(0x6a3), Weight: 65519},
	},
	0xe6fa8799bbb9625d: []BookEntry{
		{Move: Move(0x8106), Weight: 3},
	},
	0x5e2b665eeeaf6758: []BookEntry{
		{Move: Move(0xf7c), Weight: 2},
	},
	0x81456c592c3d8a1d: []BookEntry{
		{Move: Move(0x9d), Weight: 65520},
	},
	0x5123d4c2d9dcc570: []BookEntry{
		{Move: Move(0x31c), Weight: 1},
	},
	0x655a9f3835215673: []BookEntry{
		{Move: Move(0x89a), Weight: 1},
	},
	0x830608254ab8e51d: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0xb9084a2b8449c1e3: []BookEntry{
		{Move: Move(0x3d7), Weight: 2},
		{Move: Move(0x4b), Weight: 2},
	},
	0xd30d8ccb00150eac: []BookEntry{
		{Move: Move(0x3d7), Weight: 65520},
		{Move: Move(0x6e2), Weight: 7280},
	},
	0x621670670f66f98: []BookEntry{
		{Move: Move(0x15a), Weight: 1},
	},
	0xc949748244b3680: []BookEntry{
		{Move: Move(0xf59), Weight: 4},
	},
	0xd5f8be0c26fe003c: []BookEntry{
		{Move: Move(0x85a), Weight: 1},
	},
	0xd603009da8a2b1a3: []BookEntry{
		{Move: Move(0x4da), Weight: 1},
	},
	0x1262dbdfd1b7e3f5: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0x2ba20ab88afb1444: []BookEntry{
		{Move: Move(0xf6b), Weight: 2},
	},
	0xbf5f0caccbc89a88: []BookEntry{
		{Move: Move(0x953), Weight: 1},
	},
	0xfaaee392938103d: []BookEntry{
		{Move: Move(0xe9e), Weight: 65520},
		{Move: Move(0xa9b), Weight: 21840},
	},
	0x1d28429dfcd1effa: []BookEntry{
		{Move: Move(0x6d1), Weight: 2},
	},
	0x2e1da775779c20b1: []BookEntry{
		{Move: Move(0x2db), Weight: 14},
	},
	0x4186a9e5c2cbcb28: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0x4e27f31e52691f70: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0x51d5967f34611089: []BookEntry{
		{Move: Move(0xceb), Weight: 2},
	},
	0x8c4a1721d6c65218: []BookEntry{
		{Move: Move(0xb5e), Weight: 1},
	},
	0xca1f822370873134: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
	},
	0x1148c1a308b051c8: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
		{Move: Move(0xf7c), Weight: 1},
	},
	0x130bb836930e04fb: []BookEntry{
		{Move: Move(0x314), Weight: 2},
	},
	0x409027d3923aeaae: []BookEntry{
		{Move: Move(0xceb), Weight: 65520},
		{Move: Move(0xb5e), Weight: 1},
	},
	0xa5277879b3826730: []BookEntry{
		{Move: Move(0x51b), Weight: 3},
	},
	0xb33ef34fd8283889: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xcdf073c3e8bccf75: []BookEntry{
		{Move: Move(0x4da), Weight: 1},
	},
	0x1fb5508677577e18: []BookEntry{
		{Move: Move(0x2d3), Weight: 3},
	},
	0x74b2b6867b694cb1: []BookEntry{
		{Move: Move(0xf74), Weight: 65520},
		{Move: Move(0xf6b), Weight: 28080},
	},
	0x92ee5d11b09412d3: []BookEntry{
		{Move: Move(0x8dc), Weight: 1},
	},
	0xc74c59ad2aaa3893: []BookEntry{
		{Move: Move(0x7a7), Weight: 65520},
		{Move: Move(0x795), Weight: 21840},
	},
	0x27ca220e8b8ba3ca: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0x3a80900f66be8c68: []BookEntry{
		{Move: Move(0x153), Weight: 3},
	},
	0x670b17fb255aa6f4: []BookEntry{
		{Move: Move(0x2db), Weight: 6},
		{Move: Move(0x195), Weight: 5},
	},
	0xd11ed0fca9bffa9b: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
		{Move: Move(0xc4), Weight: 1},
	},
	0xfbca8ec227e07a2b: []BookEntry{
		{Move: Move(0xf1c), Weight: 4},
	},
	0xc09e7fc4272ad92: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0x874fae9aaf0e90be: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x95fccf288f18fadc: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xc1960089cdb189a5: []BookEntry{
		{Move: Move(0xb25), Weight: 1},
	},
	0x19d3651dd21a71dc: []BookEntry{
		{Move: Move(0xae2), Weight: 1},
	},
	0x81de519d791b3087: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0x8c4bf20f5170019a: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0xcc7298cef441c622: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
	},
	0xd1095449022e6af8: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0xf71669f28bcce23c: []BookEntry{
		{Move: Move(0xc69), Weight: 1},
	},
	0x1297b266d226456d: []BookEntry{
		{Move: Move(0x2db), Weight: 5},
		{Move: Move(0x195), Weight: 3},
	},
	0x370b21c3003066da: []BookEntry{
		{Move: Move(0x48c), Weight: 1},
	},
	0x37d56dbf27b9f8b5: []BookEntry{
		{Move: Move(0xae4), Weight: 1},
	},
	0x4f837c7daacc3079: []BookEntry{
		{Move: Move(0x195), Weight: 8},
	},
	0xbff6a0307d9dde7e: []BookEntry{
		{Move: Move(0x31c), Weight: 1},
	},
	0xe0ad49d142f169ea: []BookEntry{
		{Move: Move(0xe39), Weight: 1},
	},
	0xf559c04f8d86ae15: []BookEntry{
		{Move: Move(0xfbf), Weight: 13104},
		{Move: Move(0xc20), Weight: 26208},
		{Move: Move(0xeac), Weight: 13104},
		{Move: Move(0xeb3), Weight: 65520},
	},
	0x28f1f2c33b9f923a: []BookEntry{
		{Move: Move(0x51c), Weight: 65520},
	},
	0x58d0600b799f7843: []BookEntry{
		{Move: Move(0x9ee), Weight: 1},
	},
	0x700657c7337940e0: []BookEntry{
		{Move: Move(0x195), Weight: 4},
		{Move: Move(0x396), Weight: 2},
	},
	0x8ac3b9d2c61cedbd: []BookEntry{
		{Move: Move(0x9ef), Weight: 1},
	},
	0x92ae5a13cf4e23d2: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0xfcd0e48e3ddd706e: []BookEntry{
		{Move: Move(0x6a2), Weight: 1},
	},
	0x10ebb3738941ec47: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x1b343593cf02bdbf: []BookEntry{
		{Move: Move(0xee3), Weight: 1},
	},
	0x5b42ebc23ff419cb: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x8d7c86d0b048f56d: []BookEntry{
		{Move: Move(0x14c), Weight: 12},
	},
	0xc49cea41e90fd12: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
		{Move: Move(0x8a9), Weight: 2808},
		{Move: Move(0xc61), Weight: 6552},
		{Move: Move(0xe73), Weight: 9360},
		{Move: Move(0xdef), Weight: 9360},
	},
	0x9127b4c010a91f7c: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0x818b8b6b21d66d00: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0xaba1a5923bb6117a: []BookEntry{
		{Move: Move(0x14c), Weight: 3},
		{Move: Move(0x396), Weight: 2},
		{Move: Move(0x4b), Weight: 2},
		{Move: Move(0x6ea), Weight: 1},
		{Move: Move(0x6d5), Weight: 1},
	},
	0xbbba5c50546b3b07: []BookEntry{
		{Move: Move(0xd6d), Weight: 1},
	},
	0xdbd170d44ee5ef76: []BookEntry{
		{Move: Move(0xf7c), Weight: 1},
	},
	0x18ac24e512453966: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x75f39380da666127: []BookEntry{
		{Move: Move(0xca2), Weight: 4},
	},
	0x99e48752953716c1: []BookEntry{
		{Move: Move(0x660), Weight: 29},
	},
	0xcc342fa75459ce70: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xe6bd7d01a5e9099d: []BookEntry{
		{Move: Move(0xcaa), Weight: 65520},
	},
	0x564b3ae2359d81e: []BookEntry{
		{Move: Move(0xf6b), Weight: 8},
	},
	0x78d0c9f59aa25117: []BookEntry{
		{Move: Move(0xe6a), Weight: 35280},
		{Move: Move(0xdef), Weight: 65520},
	},
	0x9a09f140840f72d8: []BookEntry{
		{Move: Move(0xef4), Weight: 1},
	},
	0x9b579b1088780fab: []BookEntry{
		{Move: Move(0xceb), Weight: 65520},
		{Move: Move(0xce3), Weight: 2069},
		{Move: Move(0xc61), Weight: 1379},
	},
	0x9fb7114fdb7922fa: []BookEntry{
		{Move: Move(0x55b), Weight: 65520},
	},
	0xc973965c36418bc1: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0xcc292b6823c8be87: []BookEntry{
		{Move: Move(0x691), Weight: 1},
		{Move: Move(0x688), Weight: 65520},
		{Move: Move(0x3d7), Weight: 65520},
	},
	0x1cf373bf27fd15f5: []BookEntry{
		{Move: Move(0xa49), Weight: 1},
	},
	0x6bf9a850e6f453cc: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0x86d6f1f4c8bcf27e: []BookEntry{
		{Move: Move(0x8b), Weight: 1},
	},
	0xa33b9535a20ec6f2: []BookEntry{
		{Move: Move(0x89b), Weight: 65520},
	},
	0xce5a9bf97eeeb2e9: []BookEntry{
		{Move: Move(0xd6c), Weight: 1},
	},
	0xf4f7a182e5a78db5: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x106291f6789ae1e6: []BookEntry{
		{Move: Move(0xfad), Weight: 10},
		{Move: Move(0xef2), Weight: 6},
		{Move: Move(0xc28), Weight: 2},
		{Move: Move(0xd2c), Weight: 1},
	},
	0x213b7093dce3f2a5: []BookEntry{
		{Move: Move(0xc61), Weight: 2},
		{Move: Move(0xf6b), Weight: 1},
		{Move: Move(0xdef), Weight: 1},
		{Move: Move(0xc28), Weight: 1},
	},
	0x2da2618d27712507: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0x761c0f2f7c607584: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xf25e77a05530a294: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
		{Move: Move(0x3d7), Weight: 1},
	},
	0x1b7fc8b31a2528a2: []BookEntry{
		{Move: Move(0x4b), Weight: 5},
		{Move: Move(0x3d7), Weight: 2},
	},
	0x272b3a620aeb3ab5: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
	},
	0x3d00a3dd23979226: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x840250babc101dbb: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x883608223c221535: []BookEntry{
		{Move: Move(0x3d7), Weight: 3},
		{Move: Move(0x4b), Weight: 1},
	},
	0x96ed7b3f32063725: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x9c193e7c16e54b8c: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0xc7723bac4a4265bb: []BookEntry{
		{Move: Move(0x161), Weight: 3},
	},
	0x4dfb64b3ba21b449: []BookEntry{
		{Move: Move(0x564), Weight: 5},
	},
	0x803bb5e0de7805dc: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0x90901935d014a8dc: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
		{Move: Move(0xdae), Weight: 65520},
		{Move: Move(0xd24), Weight: 65520},
	},
	0x90a8a48fe0c8c7d6: []BookEntry{
		{Move: Move(0x259), Weight: 65520},
	},
	0xe92a1b07971dcd37: []BookEntry{
		{Move: Move(0xce3), Weight: 65520},
	},
	0x32f2da2bc4f28d7c: []BookEntry{
		{Move: Move(0x29a), Weight: 2},
	},
	0x623c3f4ee9ab2d8a: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
		{Move: Move(0xdef), Weight: 1},
	},
	0x71621434962d3579: []BookEntry{
		{Move: Move(0x89), Weight: 1},
		{Move: Move(0x14c), Weight: 1},
	},
	0x7170698ee023aa6d: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x92527ebf52484529: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
		{Move: Move(0xceb), Weight: 1},
	},
	0x9dec58d5df5d0448: []BookEntry{
		{Move: Move(0xea5), Weight: 65520},
		{Move: Move(0xf7c), Weight: 1},
		{Move: Move(0xd2b), Weight: 1},
	},
	0xb10068f6329b7565: []BookEntry{
		{Move: Move(0xc6a), Weight: 2},
	},
	0xf3d38bb8ac163b79: []BookEntry{
		{Move: Move(0x195), Weight: 9},
	},
	0x9fda525cc0b4c832: []BookEntry{
		{Move: Move(0x314), Weight: 10080},
		{Move: Move(0x396), Weight: 5040},
		{Move: Move(0x2d3), Weight: 65520},
		{Move: Move(0x52), Weight: 20160},
	},
	0xf70d3464c23f3ddb: []BookEntry{
		{Move: Move(0x52), Weight: 9},
		{Move: Move(0x195), Weight: 2},
	},
	0xcc0773e12480fc9: []BookEntry{
		{Move: Move(0x795), Weight: 1},
	},
	0x5865596228a91283: []BookEntry{
		{Move: Move(0xa6), Weight: 4},
		{Move: Move(0x396), Weight: 2},
	},
	0xb582806745ea6d39: []BookEntry{
		{Move: Move(0x211), Weight: 1},
	},
	0xcf181abb87b66fb8: []BookEntry{
		{Move: Move(0x55b), Weight: 65520},
	},
	0xdcfa09d5f8464e7b: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xef72096102ebd9a7: []BookEntry{
		{Move: Move(0xeeb), Weight: 1},
	},
	0xf82e0e5b6ff2266b: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x23fe5fae2b18c2c8: []BookEntry{
		{Move: Move(0x688), Weight: 1},
		{Move: Move(0x259), Weight: 1},
	},
	0x49029a7f69359b30: []BookEntry{
		{Move: Move(0x31c), Weight: 35280},
		{Move: Move(0x2d3), Weight: 65520},
		{Move: Move(0x314), Weight: 1},
	},
	0x9fef7ebd6c6cff5c: []BookEntry{
		{Move: Move(0x8106), Weight: 8},
		{Move: Move(0x688), Weight: 2},
	},
	0xc4a3fcea6447d126: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xdd6c44ebb4ce3e36: []BookEntry{
		{Move: Move(0x396), Weight: 1},
	},
	0x3749b9ee098541a1: []BookEntry{
		{Move: Move(0xdef), Weight: 5},
	},
	0x8edb7da86aaf93f1: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0x9842decfdb1efcb6: []BookEntry{
		{Move: Move(0xf6b), Weight: 2},
	},
	0x282a9a2af7efdd8a: []BookEntry{
		{Move: Move(0x8d2), Weight: 1},
	},
	0x43a61fc3a9014ff9: []BookEntry{
		{Move: Move(0x6a3), Weight: 42},
	},
	0x4626bac3e1c7ed26: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
		{Move: Move(0xf7c), Weight: 1},
	},
	0x5794bd5eeddee2f3: []BookEntry{
		{Move: Move(0x58f), Weight: 2},
	},
	0x8331f736fffa5618: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
		{Move: Move(0x756), Weight: 1},
	},
	0xd58b32e053382bde: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x5ab1cf1cefecab0: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
	},
	0x2850f6fe972ec937: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x5e3460ffb439e7b5: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0x77e9a553f47ca664: []BookEntry{
		{Move: Move(0xf7c), Weight: 1},
	},
	0x8f2cb735da7eec44: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x3c63971ac67c422d: []BookEntry{
		{Move: Move(0x161), Weight: 1},
	},
	0x4329a05d1c98bbf9: []BookEntry{
		{Move: Move(0xc20), Weight: 15883},
		{Move: Move(0xe6a), Weight: 15883},
		{Move: Move(0xf3f), Weight: 65520},
		{Move: Move(0xcaa), Weight: 65520},
		{Move: Move(0xdef), Weight: 15883},
		{Move: Move(0xe73), Weight: 15883},
	},
	0x68e583fbef1a2e37: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x93f818bc24fbdb3d: []BookEntry{
		{Move: Move(0xf7c), Weight: 4},
		{Move: Move(0xceb), Weight: 4},
		{Move: Move(0xe6a), Weight: 2},
	},
	0x9c96f0297d2204d: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
	},
	0x1344c5da514ffedc: []BookEntry{
		{Move: Move(0x652), Weight: 65520},
		{Move: Move(0xe73), Weight: 32760},
		{Move: Move(0x660), Weight: 16380},
		{Move: Move(0xcaa), Weight: 49140},
	},
	0x1cd5ddefc35ce360: []BookEntry{
		{Move: Move(0x3d7), Weight: 11562},
		{Move: Move(0x8106), Weight: 65520},
	},
	0x23d575f7bd002e53: []BookEntry{
		{Move: Move(0x8db), Weight: 28080},
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x3a79510544319108: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x46b4e7c7ba22e320: []BookEntry{
		{Move: Move(0x153), Weight: 2},
	},
	0x490f463127586971: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xbdf888a4064a50dc: []BookEntry{
		{Move: Move(0x723), Weight: 16},
	},
	0x679c7dec0dd62005: []BookEntry{
		{Move: Move(0xc28), Weight: 1},
	},
	0x942989d5b4418d87: []BookEntry{
		{Move: Move(0x91b), Weight: 1},
		{Move: Move(0xe73), Weight: 1},
	},
	0xa660c43d12cc815f: []BookEntry{
		{Move: Move(0xfad), Weight: 53607},
		{Move: Move(0xfb4), Weight: 65519},
	},
	0xc9f4d1d0a6cfd852: []BookEntry{
		{Move: Move(0x3df), Weight: 1},
	},
	0xcff87239fd76544a: []BookEntry{
		{Move: Move(0x314), Weight: 65520},
		{Move: Move(0x4b), Weight: 11562},
	},
	0xf51d135bc875570a: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0xf77a839672d73ab: []BookEntry{
		{Move: Move(0xb23), Weight: 1},
	},
	0x716e460b49675dc1: []BookEntry{
		{Move: Move(0xef4), Weight: 2},
	},
	0x91983dd5d1ce6e3f: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
		{Move: Move(0x4a3), Weight: 2},
	},
	0xc471c4592618c61c: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xf64562aa14157bd3: []BookEntry{
		{Move: Move(0x716), Weight: 1},
		{Move: Move(0xbe7), Weight: 1},
	},
	0x10d7469044f840f: []BookEntry{
		{Move: Move(0xca), Weight: 1},
	},
	0x189ec638af30f369: []BookEntry{
		{Move: Move(0xb5c), Weight: 2},
	},
	0x67873b5aeb784c8f: []BookEntry{
		{Move: Move(0x89b), Weight: 2},
	},
	0x6f208b7bdeebfc60: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0x7436d131e56330d6: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0x7aaa9ee3ccc6940c: []BookEntry{
		{Move: Move(0x31c), Weight: 65520},
		{Move: Move(0x314), Weight: 13104},
		{Move: Move(0x566), Weight: 3276},
	},
	0xc51cbe385dca36fc: []BookEntry{
		{Move: Move(0x91c), Weight: 1},
	},
	0xe03ebd685750c181: []BookEntry{
		{Move: Move(0xf74), Weight: 2},
		{Move: Move(0xce3), Weight: 2},
	},
	0x9a2250f4dfc8f82: []BookEntry{
		{Move: Move(0xeac), Weight: 65520},
		{Move: Move(0xe9e), Weight: 32760},
		{Move: Move(0xdef), Weight: 13104},
		{Move: Move(0xd6d), Weight: 19656},
	},
	0x19fa73e30e9bf3e6: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0x81cd80bc6e340b9a: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x9ccca08cae71dac3: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0xb7ea72da115b1f8a: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xbdf69ce7d591a198: []BookEntry{
		{Move: Move(0x31c), Weight: 2},
	},
	0x242754c989138dd: []BookEntry{
		{Move: Move(0x315), Weight: 1},
	},
	0x2c74513cee07acb4: []BookEntry{
		{Move: Move(0xc69), Weight: 3},
		{Move: Move(0x89b), Weight: 2},
	},
	0x5a539970f42bcd85: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
	},
	0x747d713443db49d4: []BookEntry{
		{Move: Move(0x50), Weight: 1},
	},
	0x7d0091a166e64fd5: []BookEntry{
		{Move: Move(0x829), Weight: 12},
	},
	0x385ff069cbd93cbd: []BookEntry{
		{Move: Move(0xef4), Weight: 1},
	},
	0x86e80b5c1c702884: []BookEntry{
		{Move: Move(0x92a), Weight: 65520},
	},
	0xb7741f6e780adfd5: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0xbc8df3a033d54d92: []BookEntry{
		{Move: Move(0xf6b), Weight: 2},
	},
	0xe6a64da2760c32ef: []BookEntry{
		{Move: Move(0x8dd), Weight: 1},
	},
	0x3113408ba02e0643: []BookEntry{
		{Move: Move(0xb23), Weight: 1},
	},
	0x5741bb8bf5928ca9: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xb090a6423cfcff18: []BookEntry{
		{Move: Move(0x84c), Weight: 5},
	},
	0xd69eca07c6ce50a0: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x106b7c2058877f45: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0xa79707379970e6ab: []BookEntry{
		{Move: Move(0xc20), Weight: 2},
	},
	0x2c6af09375b067ec: []BookEntry{
		{Move: Move(0xe39), Weight: 1},
	},
	0x3b7dc1d7c1850ddf: []BookEntry{
		{Move: Move(0x52), Weight: 17},
	},
	0x652f66adf42ea764: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x80f565f0775604a6: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0xa7bd30f5c9bb40a8: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x7b8869100878cbb7: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xaf5b89f48aeed8f7: []BookEntry{
		{Move: Move(0x90), Weight: 1},
		{Move: Move(0x6d1), Weight: 1},
	},
	0xd2c0046863e7a860: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0xd7e4bab0cd56fa09: []BookEntry{
		{Move: Move(0x251), Weight: 2},
		{Move: Move(0xca), Weight: 1},
	},
	0xee63882911959ac4: []BookEntry{
		{Move: Move(0x859), Weight: 1},
	},
	0x1402c5d7a46ccfec: []BookEntry{
		{Move: Move(0x913), Weight: 1},
	},
	0x2f0951ab1d8e6974: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
		{Move: Move(0x8106), Weight: 1},
	},
	0x71cb14d371ca65cb: []BookEntry{
		{Move: Move(0xb24), Weight: 1},
	},
	0x78945827b73acfe9: []BookEntry{
		{Move: Move(0xdef), Weight: 65520},
		{Move: Move(0xf7c), Weight: 29484},
		{Move: Move(0xc20), Weight: 29484},
		{Move: Move(0xe73), Weight: 22932},
		{Move: Move(0xc61), Weight: 16380},
	},
	0x85535d616f7469fb: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
	},
	0xcf5453bb03d2b43d: []BookEntry{
		{Move: Move(0x91c), Weight: 1},
	},
	0xd484bfeeabbbabf4: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xd9637dd926288f6d: []BookEntry{
		{Move: Move(0x724), Weight: 1},
	},
	0x1642c6063bf6120a: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x27832c01e087cd51: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0xaac8b6e33e845b02: []BookEntry{
		{Move: Move(0x218), Weight: 2},
	},
	0xaad5504af13c5853: []BookEntry{
		{Move: Move(0xfad), Weight: 5},
	},
	0xee4c43f892cf59cd: []BookEntry{
		{Move: Move(0xb6c), Weight: 1},
	},
	0x13c0de9acd5e9320: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
		{Move: Move(0xf59), Weight: 1},
	},
	0x167393519b1b6403: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
	},
	0x16d4ebaba3af2f17: []BookEntry{
		{Move: Move(0x49a), Weight: 1},
	},
	0x82b8c0dc7080ce23: []BookEntry{
		{Move: Move(0x314), Weight: 7},
	},
	0xa89565e593227b85: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
	},
	0x1557dbfdb09fd564: []BookEntry{
		{Move: Move(0xe9e), Weight: 65519},
		{Move: Move(0xd2c), Weight: 53607},
	},
	0x21ef36e703c987fa: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
	},
	0xb6b4eaebb92e133f: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
	},
	0xd6f5c68b1125cb22: []BookEntry{
		{Move: Move(0x9d), Weight: 1},
	},
	0x392ca071b9770c9b: []BookEntry{
		{Move: Move(0x6a2), Weight: 2},
	},
	0x3abea3f63f137a7d: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
	},
	0x4f89b23939190a37: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
		{Move: Move(0xc28), Weight: 1},
	},
	0x586d6b6e35169d14: []BookEntry{
		{Move: Move(0xe9e), Weight: 1},
	},
	0x620aa2427895d63f: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
		{Move: Move(0x153), Weight: 1},
	},
	0x81dc76342486b2f6: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0xfb790393f1ac4109: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0x2904c3f0e5479d79: []BookEntry{
		{Move: Move(0x251), Weight: 65520},
		{Move: Move(0x396), Weight: 1},
		{Move: Move(0xcc), Weight: 1},
	},
	0x663563265e75ec88: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x6e1c69ff916af9e4: []BookEntry{
		{Move: Move(0xd6d), Weight: 1},
	},
	0xd0d3ca6b1ee8754e: []BookEntry{
		{Move: Move(0x8da), Weight: 1},
	},
	0xda2dfe442fc09cc6: []BookEntry{
		{Move: Move(0xf74), Weight: 65520},
		{Move: Move(0xdef), Weight: 28080},
	},
	0x3ce5230e88a0e762: []BookEntry{
		{Move: Move(0xe6a), Weight: 3},
		{Move: Move(0xce3), Weight: 2},
	},
	0x81a20febbec9bbe5: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0x8e2296895488aa75: []BookEntry{
		{Move: Move(0xe73), Weight: 65520},
	},
	0x9d008c6fdf2c8184: []BookEntry{
		{Move: Move(0x691), Weight: 16380},
		{Move: Move(0x68c), Weight: 65520},
	},
	0xb510abdf8b1c6fe0: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0xc5db2592f9ce5fa6: []BookEntry{
		{Move: Move(0x396), Weight: 2},
	},
	0xec9a2b9bd5f09e43: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0x1886f25d4571bb8: []BookEntry{
		{Move: Move(0x6a3), Weight: 65520},
		{Move: Move(0x195), Weight: 1},
	},
	0xf3bd69eecfd0631: []BookEntry{
		{Move: Move(0x469), Weight: 2},
	},
	0x14cca9a753deaf41: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0x60a93cc6c2e4eb81: []BookEntry{
		{Move: Move(0x89a), Weight: 2},
	},
	0x7679ffb102ddcf32: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0x81ae80dc23b9f793: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x92b945b467ea7dbb: []BookEntry{
		{Move: Move(0xe6a), Weight: 35280},
		{Move: Move(0xdae), Weight: 65520},
	},
	0xd07f7b905c6183af: []BookEntry{
		{Move: Move(0x2db), Weight: 3},
	},
	0x29d7c69b5526b954: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0x6567dde9ade7fee3: []BookEntry{
		{Move: Move(0xea5), Weight: 65520},
	},
	0x73e8bd5e4a224676: []BookEntry{
		{Move: Move(0x91c), Weight: 1},
		{Move: Move(0xf74), Weight: 1},
	},
	0xd4276a2caa0bfe4: []BookEntry{
		{Move: Move(0x8106), Weight: 4},
	},
	0x32a2affd7125fcd2: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xe2ec8a4565383204: []BookEntry{
		{Move: Move(0xf6b), Weight: 2},
	},
	0x31a64d32477be19c: []BookEntry{
		{Move: Move(0xefa), Weight: 1},
	},
	0x3609e10ca2253de3: []BookEntry{
		{Move: Move(0xd3), Weight: 1},
	},
	0x4db034e28a48e698: []BookEntry{
		{Move: Move(0x49c), Weight: 1},
	},
	0x5e4737557b8ddacd: []BookEntry{
		{Move: Move(0x9ad), Weight: 1},
	},
	0x79be7ffff2b1777a: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x89b4c9fc27f09e92: []BookEntry{
		{Move: Move(0x564), Weight: 1},
	},
	0x8c7935f098f45fbe: []BookEntry{
		{Move: Move(0xef4), Weight: 1},
	},
	0x9523058a82cb004e: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0x1ff40d27db51854e: []BookEntry{
		{Move: Move(0xcea), Weight: 3},
	},
	0x470b7e72f2aac28d: []BookEntry{
		{Move: Move(0x89), Weight: 1},
		{Move: Move(0x853), Weight: 1},
	},
	0x8003e7460b276263: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x9020762ed17ce0cb: []BookEntry{
		{Move: Move(0x314), Weight: 3},
	},
	0x92633a727346a30a: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
		{Move: Move(0x259), Weight: 1},
	},
	0xb0fe07ca77c9f3f8: []BookEntry{
		{Move: Move(0x2db), Weight: 6},
	},
	0xbaaaba59854de0e9: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0xc3d21ab9d385b305: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x85a1b7724bf1bf66: []BookEntry{
		{Move: Move(0x691), Weight: 1},
	},
	0x98900e6ad7d30a11: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0x9e763a36defa29f0: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0xace300916fc57c4d: []BookEntry{
		{Move: Move(0xce3), Weight: 2},
	},
	0xad468aae628be6da: []BookEntry{
		{Move: Move(0x55f), Weight: 1},
	},
	0xd410ccf3b93565ab: []BookEntry{
		{Move: Move(0x6a3), Weight: 65520},
		{Move: Move(0x52), Weight: 11562},
	},
	0xe03783a7e27638a3: []BookEntry{
		{Move: Move(0xeac), Weight: 3},
		{Move: Move(0xcaa), Weight: 1},
		{Move: Move(0xf3f), Weight: 1},
	},
	0x16e0204f6fa21d6a: []BookEntry{
		{Move: Move(0x4b), Weight: 9},
	},
	0x5dbaa524c7f05530: []BookEntry{
		{Move: Move(0x314), Weight: 10},
		{Move: Move(0x29a), Weight: 7},
		{Move: Move(0x9d), Weight: 1},
	},
	0x6ac6f6928cc335fc: []BookEntry{
		{Move: Move(0x6a3), Weight: 65520},
	},
	0xcc93588e252623f5: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0xe94544cb4742158b: []BookEntry{
		{Move: Move(0x35d), Weight: 1},
	},
	0xecf7d9311902271c: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0x50695176ed297217: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x50e7ee8ba6737ab7: []BookEntry{
		{Move: Move(0x544), Weight: 1},
	},
	0x7bfe411cba6ed46a: []BookEntry{
		{Move: Move(0xb23), Weight: 1},
	},
	0xa6150666b9746874: []BookEntry{
		{Move: Move(0xd2c), Weight: 21840},
		{Move: Move(0xce3), Weight: 65520},
	},
	0x233d82c9b90ec1fc: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
		{Move: Move(0xf6b), Weight: 1},
	},
	0x29ee06108bf01fa6: []BookEntry{
		{Move: Move(0xae2), Weight: 1},
	},
	0x6c95cbd607315126: []BookEntry{
		{Move: Move(0xb5c), Weight: 4},
	},
	0xec794db7027ae5cd: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0x85491914675e0e8: []BookEntry{
		{Move: Move(0x218), Weight: 1},
	},
	0x9d57cb793397a962: []BookEntry{
		{Move: Move(0x2db), Weight: 65520},
	},
	0xa452c58582f0629b: []BookEntry{
		{Move: Move(0xe9e), Weight: 1},
	},
	0xb45dd54ca82f3038: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0xc99f09ee6da8c49a: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0xd95e18f624a60759: []BookEntry{
		{Move: Move(0x723), Weight: 1},
	},
	0xee8b056070961835: []BookEntry{
		{Move: Move(0xf6b), Weight: 2},
	},
	0x42367eaeef7223d4: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0x702363e0722884f4: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0x84d3b417652634b3: []BookEntry{
		{Move: Move(0x195), Weight: 5},
	},
	0x8f58030762724280: []BookEntry{
		{Move: Move(0x144), Weight: 2},
	},
	0x9c65e99bc9de392a: []BookEntry{
		{Move: Move(0x2db), Weight: 14},
	},
	0xa31e101825ea253d: []BookEntry{
		{Move: Move(0x8106), Weight: 3},
	},
	0xe27f917352ecdf82: []BookEntry{
		{Move: Move(0x51c), Weight: 2},
	},
	0x66fc332519e5cde: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x65e712d50e861e94: []BookEntry{
		{Move: Move(0xa62), Weight: 1},
	},
	0x72d91769ca19cc70: []BookEntry{
		{Move: Move(0xf62), Weight: 1},
	},
	0x8d8ee0c8ac9156f5: []BookEntry{
		{Move: Move(0x292), Weight: 65520},
		{Move: Move(0x2db), Weight: 45864},
		{Move: Move(0x8106), Weight: 19656},
	},
	0x987e3aeb0c2e82ec: []BookEntry{
		{Move: Move(0x210), Weight: 1},
		{Move: Move(0x218), Weight: 1},
	},
	0x9a138dd21f33bc2a: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0xb4d59bb0edbe190e: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0xb81dc7aee7058549: []BookEntry{
		{Move: Move(0x195), Weight: 3},
	},
	0x14a3366580725803: []BookEntry{
		{Move: Move(0x5ab), Weight: 1},
	},
	0x225ecaf14cc9b8a9: []BookEntry{
		{Move: Move(0x49b), Weight: 2},
	},
	0x37b9a5a83dbb85d8: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0x5fc9da8df63497b5: []BookEntry{
		{Move: Move(0x89a), Weight: 1},
	},
	0xa61db7f1f4945cb7: []BookEntry{
		{Move: Move(0xea5), Weight: 1},
	},
	0xcc4e8d10c3f9340e: []BookEntry{
		{Move: Move(0xab4), Weight: 1},
	},
	0xdb463c4791f8cba1: []BookEntry{
		{Move: Move(0x9d5), Weight: 1},
	},
	0xde8710f8bd6dff43: []BookEntry{
		{Move: Move(0xeb3), Weight: 1},
	},
	0x4df78606794fac90: []BookEntry{
		{Move: Move(0xeb1), Weight: 2},
		{Move: Move(0xee0), Weight: 1},
	},
	0x55155984c0b5cd89: []BookEntry{
		{Move: Move(0xce3), Weight: 3},
	},
	0x5c49dae4afb967e2: []BookEntry{
		{Move: Move(0xfad), Weight: 65520},
		{Move: Move(0x89b), Weight: 7708},
		{Move: Move(0xe9e), Weight: 3854},
	},
	0xe3f3d911e2d9efe2: []BookEntry{
		{Move: Move(0x55b), Weight: 7},
	},
	0xe815fa09bcea55e5: []BookEntry{
		{Move: Move(0x14e), Weight: 6},
	},
	0xef536648b9c30c0c: []BookEntry{
		{Move: Move(0x31c), Weight: 65520},
		{Move: Move(0x314), Weight: 43680},
	},
	0xa625029326cb3585: []BookEntry{
		{Move: Move(0xdae), Weight: 1},
	},
	0xb1044b9af1b0faa7: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xd78d18a3baadda5a: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0xe604829353f81e2e: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xf5f81fdd55a34033: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x506082a62632a7e3: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x586867969c9b423e: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0x1d04fbb4fd4bf2e8: []BookEntry{
		{Move: Move(0x259), Weight: 12},
	},
	0x3da9542413862894: []BookEntry{
		{Move: Move(0x829), Weight: 2},
	},
	0x983fbd0bfde16407: []BookEntry{
		{Move: Move(0xcaa), Weight: 2},
	},
	0xb17b079c6fdbaec7: []BookEntry{
		{Move: Move(0x144), Weight: 2},
		{Move: Move(0x54b), Weight: 1},
		{Move: Move(0x693), Weight: 1},
	},
	0xb46720d9feab1dd5: []BookEntry{
		{Move: Move(0xa6), Weight: 1},
	},
	0xcc6d7507541d56cc: []BookEntry{
		{Move: Move(0xe6a), Weight: 5},
		{Move: Move(0xf74), Weight: 1},
	},
	0x28315e3d89b6fdc: []BookEntry{
		{Move: Move(0xa62), Weight: 1},
	},
	0x42201274e0dd134c: []BookEntry{
		{Move: Move(0xea8), Weight: 1},
	},
	0x8770ce9d8e264051: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0xa6bcbb08b8269136: []BookEntry{
		{Move: Move(0x691), Weight: 1},
	},
	0xe1331a13d19ba616: []BookEntry{
		{Move: Move(0xe73), Weight: 2},
	},
	0x4512a61b5547cb3: []BookEntry{
		{Move: Move(0x259), Weight: 1},
	},
	0x168069c606b9e24b: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x1716180c2dd6fe84: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0x2790656d163868a1: []BookEntry{
		{Move: Move(0xd24), Weight: 11},
	},
	0x2fa6455e75c90500: []BookEntry{
		{Move: Move(0x859), Weight: 1},
	},
	0x5da00ec3c1f3f543: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x69b0e22d14832049: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x89701914a1c7c555: []BookEntry{
		{Move: Move(0xae2), Weight: 6},
	},
	0x1f0a502ff1111f4e: []BookEntry{
		{Move: Move(0x691), Weight: 1},
	},
	0x6ebb4c09d4e75ed2: []BookEntry{
		{Move: Move(0xfad), Weight: 5},
		{Move: Move(0xce3), Weight: 5},
	},
	0x7eafab1ab37bd637: []BookEntry{
		{Move: Move(0x251), Weight: 1},
		{Move: Move(0x7e5), Weight: 1},
	},
	0xa101f86c9a80a16d: []BookEntry{
		{Move: Move(0x94), Weight: 1},
	},
	0xab7e4cddf737f0eb: []BookEntry{
		{Move: Move(0x471), Weight: 1},
	},
	0xb5b8fe1ecfc4e766: []BookEntry{
		{Move: Move(0x161), Weight: 2},
		{Move: Move(0x153), Weight: 1},
	},
	0xd299218871539bd8: []BookEntry{
		{Move: Move(0xf76), Weight: 2},
	},
	0xe3976825fdb8d5ab: []BookEntry{
		{Move: Move(0x2d3), Weight: 11},
	},
	0x2a4732329694fa83: []BookEntry{
		{Move: Move(0xf76), Weight: 65520},
	},
	0xb8be730e56469a6c: []BookEntry{
		{Move: Move(0x812), Weight: 1},
	},
	0x5061987568759a50: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x5857431bffd6b9fa: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
		{Move: Move(0x8a9), Weight: 1},
	},
	0x127e149848263531: []BookEntry{
		{Move: Move(0x218), Weight: 1},
	},
	0x46a369ae946f507a: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0x9f1e9d1d303939b5: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xd74c2eebce750245: []BookEntry{
		{Move: Move(0x8106), Weight: 65520},
	},
	0xe857ab6c7fd2dea9: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0xed110abfaf5bbb97: []BookEntry{
		{Move: Move(0xc28), Weight: 2},
	},
	0x54c31263e9ad3b4f: []BookEntry{
		{Move: Move(0x314), Weight: 65520},
		{Move: Move(0x292), Weight: 1},
		{Move: Move(0x396), Weight: 5241},
	},
	0x55f87be7dcb8fd07: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0x69f7de8704b6d452: []BookEntry{
		{Move: Move(0x3d7), Weight: 3},
		{Move: Move(0x8106), Weight: 1},
		{Move: Move(0xcc), Weight: 1},
	},
	0xc3f9c0dd28ca2880: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0xcb8e0b6529afac49: []BookEntry{
		{Move: Move(0x15a), Weight: 9},
	},
	0xed7adcbdf2d31a90: []BookEntry{
		{Move: Move(0x29a), Weight: 2},
	},
	0xef32058201ac70bd: []BookEntry{
		{Move: Move(0xd2d), Weight: 1},
	},
	0xfec0d497904bf44a: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0x652b6e022a26f43a: []BookEntry{
		{Move: Move(0xf62), Weight: 3},
	},
	0x714520deb77563f7: []BookEntry{
		{Move: Move(0x29a), Weight: 8},
	},
	0xa5418b488576d129: []BookEntry{
		{Move: Move(0x2db), Weight: 5},
	},
	0xacfd056b6b4d5fda: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
	},
	0x804be23d3563f4a: []BookEntry{
		{Move: Move(0x9ee), Weight: 1},
	},
	0x378e94ad467459ae: []BookEntry{
		{Move: Move(0x89a), Weight: 2},
	},
	0x6bbfac8fc4ae8427: []BookEntry{
		{Move: Move(0xc68), Weight: 1},
	},
	0xa852aace5fb5cf6c: []BookEntry{
		{Move: Move(0x89b), Weight: 65520},
		{Move: Move(0xee9), Weight: 16380},
	},
	0xa9b6fdaff0b3ccd6: []BookEntry{
		{Move: Move(0xce3), Weight: 3},
	},
	0xc84c1824c1c86759: []BookEntry{
		{Move: Move(0x210), Weight: 65520},
	},
	0xceb48bdfa26c4fee: []BookEntry{
		{Move: Move(0xcaa), Weight: 2},
	},
	0x327c5767b175f15a: []BookEntry{
		{Move: Move(0x688), Weight: 4},
		{Move: Move(0x3d7), Weight: 1},
	},
	0x3821bc803e9f887b: []BookEntry{
		{Move: Move(0x7a7), Weight: 2},
	},
	0x8f4b0caa82ded588: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xd35e9cffce9aca24: []BookEntry{
		{Move: Move(0xe3), Weight: 2},
	},
	0x18491d6eb8e4d0cd: []BookEntry{
		{Move: Move(0xef3), Weight: 1},
	},
	0x1a47183f866513c9: []BookEntry{
		{Move: Move(0x8106), Weight: 65520},
	},
	0x41a5cdecc84e0bfb: []BookEntry{
		{Move: Move(0x99f), Weight: 4},
	},
	0x46e0599a50c02bdd: []BookEntry{
		{Move: Move(0x14e), Weight: 65520},
		{Move: Move(0x2db), Weight: 7280},
	},
	0x86ee43a55e90bf53: []BookEntry{
		{Move: Move(0xea5), Weight: 21840},
		{Move: Move(0xca2), Weight: 65520},
	},
	0x898883d58df161b7: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
		{Move: Move(0xca2), Weight: 1},
	},
	0xad0aedf1d3dd3b44: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x290ceffe24d4f0e: []BookEntry{
		{Move: Move(0xc61), Weight: 65519},
		{Move: Move(0xf7c), Weight: 53607},
	},
	0x7124b486804cb0e: []BookEntry{
		{Move: Move(0xeed), Weight: 1},
	},
	0x3f5bb57219b15d94: []BookEntry{
		{Move: Move(0xf3f), Weight: 3},
		{Move: Move(0xc20), Weight: 1},
	},
	0xe0264e76c646bcae: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x4d57308e0e6f7f0b: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x4f7b630adb793873: []BookEntry{
		{Move: Move(0x544), Weight: 1},
	},
	0x5e4b03b8cabb8e4d: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0x7db44eb47b96df45: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
		{Move: Move(0x691), Weight: 1},
		{Move: Move(0x688), Weight: 1},
	},
	0xc2e61cc3ddc8bd71: []BookEntry{
		{Move: Move(0xaf4), Weight: 1},
		{Move: Move(0xd65), Weight: 1},
		{Move: Move(0xeac), Weight: 1},
	},
	0xe0192d21a9636ce2: []BookEntry{
		{Move: Move(0xfad), Weight: 17},
	},
	0xe35ddbde1e6115c3: []BookEntry{
		{Move: Move(0x2d2), Weight: 2},
	},
	0xe5c4e46808e839c8: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xed587e74b88117ef: []BookEntry{
		{Move: Move(0x8106), Weight: 65520},
	},
	0x1958565e58f5af5c: []BookEntry{
		{Move: Move(0x9102), Weight: 1},
	},
	0x5bbeaef112bcf1ae: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0x66e831738d7f03b4: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x864ef7afeb882068: []BookEntry{
		{Move: Move(0x9d), Weight: 1},
	},
	0xd614b836700fe8e4: []BookEntry{
		{Move: Move(0x84c), Weight: 65520},
	},
	0x785f9bd7852fd0c: []BookEntry{
		{Move: Move(0xf7c), Weight: 2},
		{Move: Move(0xd2c), Weight: 1},
		{Move: Move(0xc69), Weight: 1},
		{Move: Move(0xefc), Weight: 1},
	},
	0x13a0cb4b08c808c4: []BookEntry{
		{Move: Move(0x259), Weight: 2},
	},
	0x1b0748fc7e923329: []BookEntry{
		{Move: Move(0x668), Weight: 2},
	},
	0x2bba101c45affeb8: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x429bdace5b4f2d6e: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x631293e5b4179162: []BookEntry{
		{Move: Move(0x195), Weight: 3},
	},
	0xc0080d3344142bf5: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0xcc3c7ff5bdc7fc0c: []BookEntry{
		{Move: Move(0x859), Weight: 3},
	},
	0x649b5cbbdced4af4: []BookEntry{
		{Move: Move(0xca), Weight: 2},
	},
	0xb554593572d63a4e: []BookEntry{
		{Move: Move(0x15a), Weight: 1},
	},
	0xfb58fdd1bf4048ed: []BookEntry{
		{Move: Move(0x14e), Weight: 1},
	},
	0x391c8060ac173ec2: []BookEntry{
		{Move: Move(0x89b), Weight: 2},
	},
	0xdc1cd8ba3abf04e0: []BookEntry{
		{Move: Move(0x6a3), Weight: 65520},
	},
	0x4e64d08b2132984: []BookEntry{
		{Move: Move(0xcaa), Weight: 7},
	},
	0x29956f7e6c572817: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
		{Move: Move(0x251), Weight: 1},
	},
	0x6d26cb69015e20be: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
	},
	0xaa667f5fddcd0957: []BookEntry{
		{Move: Move(0xe9e), Weight: 1},
	},
	0xd86f7979188ddaf5: []BookEntry{
		{Move: Move(0x811), Weight: 1},
	},
	0xff7155aa70f1ab5d: []BookEntry{
		{Move: Move(0x218), Weight: 65520},
		{Move: Move(0x498), Weight: 35280},
	},
	0x378afdc45f80939d: []BookEntry{
		{Move: Move(0xf7c), Weight: 1},
	},
	0xe1d504c64404d9c2: []BookEntry{
		{Move: Move(0x96e), Weight: 3},
		{Move: Move(0x8da), Weight: 1},
	},
	0xe908070fbff64b97: []BookEntry{
		{Move: Move(0x144), Weight: 65520},
		{Move: Move(0x355), Weight: 7280},
	},
	0xf659b3658d303746: []BookEntry{
		{Move: Move(0xd1), Weight: 1},
	},
	0x4d5d82d531cb08c5: []BookEntry{
		{Move: Move(0xab2), Weight: 1},
	},
	0xdee52e24b5c517bc: []BookEntry{
		{Move: Move(0x2db), Weight: 5},
	},
	0xea91997b885bf3ab: []BookEntry{
		{Move: Move(0xcc), Weight: 2},
	},
	0x1c56527337e1a22f: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x8c918d6c8e698b47: []BookEntry{
		{Move: Move(0x210), Weight: 2},
	},
	0x97d24b00103dbd9a: []BookEntry{
		{Move: Move(0x195), Weight: 3},
	},
	0xee8a85a443065287: []BookEntry{
		{Move: Move(0xeed), Weight: 3},
	},
	0x1e648f2a6fb28e7c: []BookEntry{
		{Move: Move(0xc28), Weight: 1},
	},
	0x2b17ac6876aaddd9: []BookEntry{
		{Move: Move(0x8d2), Weight: 1},
	},
	0x87913e200b9fb6af: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
		{Move: Move(0x29a), Weight: 35280},
	},
	0xabe96d1e0acb8867: []BookEntry{
		{Move: Move(0xee9), Weight: 1},
	},
	0xe70730dc976afc65: []BookEntry{
		{Move: Move(0x6ea), Weight: 1},
	},
	0xab0de3b4114c732d: []BookEntry{
		{Move: Move(0x89b), Weight: 3},
	},
	0xb90ccb17e25489e6: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
	},
	0xdea3ccaa70155cbd: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0x26a61d21a538c3e1: []BookEntry{
		{Move: Move(0xcea), Weight: 5},
	},
	0x53027d4751c0f95c: []BookEntry{
		{Move: Move(0x89b), Weight: 3},
		{Move: Move(0xe6a), Weight: 1},
		{Move: Move(0xdef), Weight: 1},
	},
	0x9c584dd4027ab842: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xab94586fc8f9a17d: []BookEntry{
		{Move: Move(0x859), Weight: 1},
	},
	0xbd0beeda0d7973ea: []BookEntry{
		{Move: Move(0xe73), Weight: 7},
	},
	0xc101c80f53178ecc: []BookEntry{
		{Move: Move(0xc20), Weight: 2},
		{Move: Move(0xdef), Weight: 1},
	},
	0x237795e568401291: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x29d384aa3fe16b99: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x359425deeb1dbdde: []BookEntry{
		{Move: Move(0xd2c), Weight: 6},
	},
	0x44698161511a19b5: []BookEntry{
		{Move: Move(0x795), Weight: 1},
	},
	0x549025d2e4ed4d3e: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xca3d86ace3ebb5a3: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
		{Move: Move(0xe68), Weight: 1},
		{Move: Move(0xce3), Weight: 1},
	},
	0xeb63660a17053907: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0xfbc828d4b775c13d: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x23042b38099ba0fd: []BookEntry{
		{Move: Move(0x3d7), Weight: 2},
	},
	0x1ff2ec083124dbbc: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x296745847f29a6f4: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x4e8677d97a6b5167: []BookEntry{
		{Move: Move(0x49b), Weight: 1},
		{Move: Move(0x8106), Weight: 1},
	},
	0x54e62ffff6c0f420: []BookEntry{
		{Move: Move(0x292), Weight: 4},
	},
	0x5ce0ec86d3ae98d8: []BookEntry{
		{Move: Move(0xdae), Weight: 6},
		{Move: Move(0xce3), Weight: 2},
		{Move: Move(0xd2c), Weight: 1},
	},
	0x6e138100ac77f5f5: []BookEntry{
		{Move: Move(0xf74), Weight: 2},
	},
	0x6fb302540693b80e: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xa7e055844e5b6b49: []BookEntry{
		{Move: Move(0xf7c), Weight: 65520},
		{Move: Move(0xdef), Weight: 9240},
		{Move: Move(0xc20), Weight: 9240},
	},
	0xa9ad7dbd2ee5d665: []BookEntry{
		{Move: Move(0x31c), Weight: 7},
	},
	0xb73c09374a2c196c: []BookEntry{
		{Move: Move(0xca2), Weight: 1},
		{Move: Move(0xe6a), Weight: 1},
		{Move: Move(0xf7c), Weight: 1},
	},
	0xdccf7258c0547100: []BookEntry{
		{Move: Move(0xc69), Weight: 1},
	},
	0xf7f9351170540933: []BookEntry{
		{Move: Move(0x31c), Weight: 1},
	},
	0x2d3f3caf5ebfdd5d: []BookEntry{
		{Move: Move(0x89a), Weight: 1},
	},
	0x76c6d1adb23e7aae: []BookEntry{
		{Move: Move(0xc69), Weight: 1},
	},
	0x912c02687ee15536: []BookEntry{
		{Move: Move(0x161), Weight: 1},
	},
	0xaac3beefb4d714f6: []BookEntry{
		{Move: Move(0xd5), Weight: 1},
	},
	0xb3f04ab56c12a039: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0xb8f3a13801493927: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0xf389a5f9bf04667d: []BookEntry{
		{Move: Move(0xe73), Weight: 4},
		{Move: Move(0xcaa), Weight: 2},
	},
	0xa056279e4f9ff7e: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0x15672023e333bea0: []BookEntry{
		{Move: Move(0xab9), Weight: 1},
	},
	0x212f85dba904a105: []BookEntry{
		{Move: Move(0xfad), Weight: 2},
	},
	0x5e35275b7960fd34: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
		{Move: Move(0xc20), Weight: 1},
	},
	0x63088230f562b9e2: []BookEntry{
		{Move: Move(0x6ca), Weight: 1},
	},
	0xa14a3af035fb7f3d: []BookEntry{
		{Move: Move(0x50), Weight: 1},
	},
	0xb8a7b1878ad5fc60: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0xbccdcb0490d4932f: []BookEntry{
		{Move: Move(0x15a), Weight: 1},
		{Move: Move(0x14c), Weight: 1},
	},
	0xe4f4a600674dbb8: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x2867642f953169f6: []BookEntry{
		{Move: Move(0x8b), Weight: 1},
	},
	0x7155d2002fe7cd36: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x82e48c22453ffd36: []BookEntry{
		{Move: Move(0xce1), Weight: 2},
	},
	0x98ceb48647f08530: []BookEntry{
		{Move: Move(0xd2c), Weight: 2},
		{Move: Move(0xef2), Weight: 1},
	},
	0xcba925b8864e6727: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
	},
	0xee12fead096add0b: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0xf3108e32ba7ee322: []BookEntry{
		{Move: Move(0xae2), Weight: 1},
	},
	0x35c20c82567a0545: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x38945133665d33e0: []BookEntry{
		{Move: Move(0x691), Weight: 1},
	},
	0x4cab119588ce9bb5: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x872d3612c75f06df: []BookEntry{
		{Move: Move(0x14c), Weight: 20160},
		{Move: Move(0x292), Weight: 65520},
		{Move: Move(0x161), Weight: 10080},
		{Move: Move(0x6e2), Weight: 5040},
	},
	0x8b99e2127468b00c: []BookEntry{
		{Move: Move(0xef2), Weight: 65520},
		{Move: Move(0xf6b), Weight: 28080},
	},
	0xb93efe2205509f34: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xc7c4db8517f10461: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0xd8e08d47aaa29048: []BookEntry{
		{Move: Move(0xca2), Weight: 65520},
		{Move: Move(0xfad), Weight: 7280},
		{Move: Move(0xe6a), Weight: 1},
	},
	0x11e0939c10de3642: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x5c798a3fbed35a6c: []BookEntry{
		{Move: Move(0x471), Weight: 1},
	},
	0xcd58fd6b301042e6: []BookEntry{
		{Move: Move(0xca), Weight: 1},
	},
	0x1f332598e204d80d: []BookEntry{
		{Move: Move(0xe73), Weight: 3},
		{Move: Move(0xcaa), Weight: 2},
	},
	0x96f0a063d72351a9: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x9f0c44191e0fc82d: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0xab8e63a3c52e7390: []BookEntry{
		{Move: Move(0x6e4), Weight: 6},
	},
	0xda5c645e884b9c6d: []BookEntry{
		{Move: Move(0xf74), Weight: 2},
	},
	0x799867ac37aa914b: []BookEntry{
		{Move: Move(0xb73), Weight: 1},
	},
	0xa0181492ebf3db43: []BookEntry{
		{Move: Move(0x161), Weight: 3},
	},
	0xf5937bba814a62e: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
	},
	0x2cd4075ed7fb53eb: []BookEntry{
		{Move: Move(0x51b), Weight: 2},
	},
	0x552b0f209561e83c: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0x5574cdad16524440: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x564990c3511635ae: []BookEntry{
		{Move: Move(0x315), Weight: 1},
	},
	0x5444dacc84c251ab: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0x86e015a7fb8d4dc6: []BookEntry{
		{Move: Move(0x29a), Weight: 10},
	},
	0xdcde7a2dc06a9ef9: []BookEntry{
		{Move: Move(0xca), Weight: 1},
	},
	0x2333cfdbab192008: []BookEntry{
		{Move: Move(0x89a), Weight: 1},
	},
	0x3799d59699abfdea: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xd3ed9e3c2cd169c8: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xe76e1a845c10f6e5: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0x48226e0e4d8c2fc0: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
	},
	0x1cbc7a034053938f: []BookEntry{
		{Move: Move(0x693), Weight: 1},
	},
	0x2b5638eb16b793cb: []BookEntry{
		{Move: Move(0x31c), Weight: 35280},
		{Move: Move(0x195), Weight: 65520},
	},
	0x368dbfe20ff63c3d: []BookEntry{
		{Move: Move(0x2d3), Weight: 7},
	},
	0x403f6d5b2388ab4e: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xaad6f7e6a94bc16e: []BookEntry{
		{Move: Move(0x3d7), Weight: 2},
		{Move: Move(0x8106), Weight: 1},
	},
	0xca18093c559e579b: []BookEntry{
		{Move: Move(0x195), Weight: 10920},
		{Move: Move(0x31c), Weight: 65520},
		{Move: Move(0x29a), Weight: 10920},
		{Move: Move(0x314), Weight: 10920},
		{Move: Move(0x292), Weight: 10920},
	},
	0x1c6b62dc229020f0: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0x331ca38c93779fb9: []BookEntry{
		{Move: Move(0xcaa), Weight: 9},
	},
	0x5f054cb6745fd8c1: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0x68a1909fdf5cf407: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xc31d06c31efbb6f6: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0xf19734b0030315c6: []BookEntry{
		{Move: Move(0x3eb), Weight: 1},
	},
	0x4e48b85d7ee3bd36: []BookEntry{
		{Move: Move(0x91b), Weight: 2},
	},
	0x52e73ab72679947c: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x6273507cf1b46673: []BookEntry{
		{Move: Move(0x315), Weight: 3},
	},
	0xb2c42ff23ae78ca1: []BookEntry{
		{Move: Move(0x210), Weight: 65520},
		{Move: Move(0x52), Weight: 32760},
		{Move: Move(0x195), Weight: 32760},
	},
	0xbc2bbafeabeac37a: []BookEntry{
		{Move: Move(0xf3f), Weight: 4},
		{Move: Move(0xcaa), Weight: 1},
	},
	0x3976815127e5ace9: []BookEntry{
		{Move: Move(0x722), Weight: 1},
	},
	0x5a7488de9032b32e: []BookEntry{
		{Move: Move(0x9d), Weight: 1},
	},
	0x8c43b624c51d8c60: []BookEntry{
		{Move: Move(0x660), Weight: 2},
	},
	0xe7308b62a0331f38: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0xe68e5e41bddcb96: []BookEntry{
		{Move: Move(0x355), Weight: 1},
		{Move: Move(0x14c), Weight: 1},
	},
	0x12da47dc2b3fd0fc: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0x24e0fb29c79abdac: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
		{Move: Move(0x251), Weight: 1},
	},
	0x606c54b2d0aab1e8: []BookEntry{
		{Move: Move(0x8a9), Weight: 1},
	},
	0xf66540b7142d509b: []BookEntry{
		{Move: Move(0xb73), Weight: 1},
	},
	0x350ee08f679ba33a: []BookEntry{
		{Move: Move(0x754), Weight: 1},
	},
	0x9e362616a7188eae: []BookEntry{
		{Move: Move(0x4da), Weight: 1},
	},
	0x9fefea0668743af0: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0xace818716347bb49: []BookEntry{
		{Move: Move(0xb24), Weight: 1},
	},
	0xbfe1bc346df3c46d: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xc0f7e62ca4b295c: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x7b0f62b8d9f37ea2: []BookEntry{
		{Move: Move(0x252), Weight: 1},
	},
	0x9547930ab93516b7: []BookEntry{
		{Move: Move(0xe73), Weight: 65520},
		{Move: Move(0xf7c), Weight: 32760},
		{Move: Move(0xdef), Weight: 17035},
		{Move: Move(0xe9e), Weight: 17035},
	},
	0xa17290540b629be0: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0xabfadd6618d772a0: []BookEntry{
		{Move: Move(0xeeb), Weight: 1},
	},
	0xc0ca2d7f9ef4c69e: []BookEntry{
		{Move: Move(0x153), Weight: 6},
	},
	0xdfab1596914c469c: []BookEntry{
		{Move: Move(0xe68), Weight: 1},
		{Move: Move(0xf3f), Weight: 1},
	},
	0x9b6b454156f87cb: []BookEntry{
		{Move: Move(0x31a), Weight: 1},
	},
	0x16d3a0cb3591ef02: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x40fee3bf3d1d17be: []BookEntry{
		{Move: Move(0xdad), Weight: 1},
	},
	0x4215f6c940398331: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
		{Move: Move(0xf59), Weight: 1},
	},
	0x79b76a7425b5e0ff: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
		{Move: Move(0xeb1), Weight: 1},
	},
	0x79eeee6fac165e01: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x7c096252dcc75f0a: []BookEntry{
		{Move: Move(0xf6b), Weight: 3},
	},
	0xa1b085fe4b2ec0c2: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0x2b63b51dcfd7a8cc: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
		{Move: Move(0xfad), Weight: 1},
		{Move: Move(0xee9), Weight: 1},
	},
	0x426bff3cf7192ee7: []BookEntry{
		{Move: Move(0x2d2), Weight: 6},
	},
	0xd652e4622d1fa790: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0xf11d18fe3aefb60e: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xf34f84755e74095e: []BookEntry{
		{Move: Move(0xaa2), Weight: 65520},
	},
	0xc927a821fd69c9c: []BookEntry{
		{Move: Move(0xae2), Weight: 2},
	},
	0x3525aa8a2313f838: []BookEntry{
		{Move: Move(0x144), Weight: 1},
	},
	0x59537b2d02cf9f16: []BookEntry{
		{Move: Move(0x218), Weight: 1},
	},
	0x6c83d205e17ae194: []BookEntry{
		{Move: Move(0x2db), Weight: 2},
		{Move: Move(0x86a), Weight: 1},
	},
	0x7ffd1411a8a4c0a8: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0xbcbe96108c324d93: []BookEntry{
		{Move: Move(0xcc), Weight: 1},
	},
	0xe84a7364e3902599: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x2085b08da5172919: []BookEntry{
		{Move: Move(0x14c), Weight: 3},
		{Move: Move(0x52), Weight: 2},
		{Move: Move(0x153), Weight: 1},
	},
	0x83917a240abc78bb: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
	},
	0x942fa0b7423635ec: []BookEntry{
		{Move: Move(0x92a), Weight: 5},
	},
	0xa598a887da2ef8e5: []BookEntry{
		{Move: Move(0xceb), Weight: 65520},
	},
	0x3247739781e7bf25: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
		{Move: Move(0x6e2), Weight: 1},
		{Move: Move(0x18c), Weight: 1},
	},
	0x45c65b144307903f: []BookEntry{
		{Move: Move(0x52), Weight: 9},
	},
	0x87db47e581f551b9: []BookEntry{
		{Move: Move(0xa19), Weight: 1},
	},
	0xd0a21391943ab0df: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x1cdb7ee64e428de8: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
		{Move: Move(0xc20), Weight: 1},
	},
	0x2b47db50d6906e11: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x3b20c2bf164def59: []BookEntry{
		{Move: Move(0xdae), Weight: 1},
	},
	0x8d8b6c03b4d3307f: []BookEntry{
		{Move: Move(0x218), Weight: 2},
	},
	0x98eec903957a6fbc: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x3ce10aee1058e45e: []BookEntry{
		{Move: Move(0x953), Weight: 6},
	},
	0x56868f3fb1cc8f09: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0xf81bc8844d46476f: []BookEntry{
		{Move: Move(0xfad), Weight: 14560},
		{Move: Move(0xce3), Weight: 65520},
	},
	0xfbe83f22d9cef145: []BookEntry{
		{Move: Move(0x6e2), Weight: 65520},
		{Move: Move(0x14e), Weight: 43680},
	},
	0x25035b128f0af0a6: []BookEntry{
		{Move: Move(0xf3f), Weight: 7},
		{Move: Move(0xc69), Weight: 2},
	},
	0x2dce9e18319a1037: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xda601cb4ac88427c: []BookEntry{
		{Move: Move(0xf59), Weight: 1},
	},
	0xe9e9f0117daeafb2: []BookEntry{
		{Move: Move(0xf6b), Weight: 2},
	},
	0xd749b14ddc136cb: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0x3dbd5b87f44b2078: []BookEntry{
		{Move: Move(0x161), Weight: 4},
		{Move: Move(0x14c), Weight: 4},
	},
	0xb493c67df6896c90: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
	},
	0xb90ef1bc502503a9: []BookEntry{
		{Move: Move(0x8da), Weight: 1},
	},
	0xc17db198557d76d4: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x11c221316965553f: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x5bd1a8ecd157bb0c: []BookEntry{
		{Move: Move(0xf3f), Weight: 4},
	},
	0x62251c3e8a145072: []BookEntry{
		{Move: Move(0x14e), Weight: 2},
	},
	0x71f20b8c60706d5e: []BookEntry{
		{Move: Move(0x314), Weight: 5},
	},
	0x7a3ae36d530e686d: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
		{Move: Move(0xc20), Weight: 1},
	},
	0xa9dd452c045db21d: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x1d8175aa2a5c172b: []BookEntry{
		{Move: Move(0x2d3), Weight: 2},
	},
	0x581c7ef5686db41e: []BookEntry{
		{Move: Move(0x8eb), Weight: 1},
		{Move: Move(0x218), Weight: 1},
	},
	0x1550d51ebe73e2d4: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0xb272e138b837c5f8: []BookEntry{
		{Move: Move(0x652), Weight: 2},
	},
	0xc6b7b93635b96c6: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x26868226c23144c9: []BookEntry{
		{Move: Move(0xb63), Weight: 16},
	},
	0x4a42d4107cb14134: []BookEntry{
		{Move: Move(0x91c), Weight: 1},
	},
	0x4dbaa903fcc2b39e: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0x9e8bad6edd7b788: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xa1916e195dde612: []BookEntry{
		{Move: Move(0x8a9), Weight: 1},
	},
	0x5f657e5fc94d4462: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x795ad4a75c75ea8b: []BookEntry{
		{Move: Move(0x195), Weight: 3},
		{Move: Move(0x2db), Weight: 2},
	},
	0x9daf61293d9495bd: []BookEntry{
		{Move: Move(0xee9), Weight: 1},
		{Move: Move(0xf3f), Weight: 1},
	},
	0xc37b2985d6bce6a7: []BookEntry{
		{Move: Move(0xd65), Weight: 9635},
		{Move: Move(0xef3), Weight: 65519},
		{Move: Move(0xf74), Weight: 28905},
	},
	0xc9c8ad2d1e403b57: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0xd18eab5ef23f7ec5: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0x6417a8505a53dcc4: []BookEntry{
		{Move: Move(0xcaa), Weight: 2},
	},
	0x8e8a02d257ab5239: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x48b719dc95a341eb: []BookEntry{
		{Move: Move(0xceb), Weight: 19},
	},
	0x4af0f97d1ae1be58: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0x56258dcaebdd1227: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x87c7e066ea9fb988: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0x97260a80977dd622: []BookEntry{
		{Move: Move(0xda6), Weight: 1},
	},
	0xb77f973fec92ed57: []BookEntry{
		{Move: Move(0xcaa), Weight: 3},
	},
	0xd33f7739a5c9d388: []BookEntry{
		{Move: Move(0xce3), Weight: 65520},
		{Move: Move(0xf7c), Weight: 28080},
	},
	0x2c1f34ad86d4e075: []BookEntry{
		{Move: Move(0x6a3), Weight: 65520},
	},
	0x43a7ab86f201809a: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x92c93ed391ada19f: []BookEntry{
		{Move: Move(0x94), Weight: 17869},
		{Move: Move(0x3d7), Weight: 65519},
		{Move: Move(0x688), Weight: 35738},
	},
	0xbc76a27430c370e4: []BookEntry{
		{Move: Move(0x161), Weight: 65519},
		{Move: Move(0x2db), Weight: 53607},
	},
	0xe39eef5afe5f79a7: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0xb3ea34175d62686: []BookEntry{
		{Move: Move(0xa99), Weight: 1},
	},
	0x99dd36a42426b80e: []BookEntry{
		{Move: Move(0x218), Weight: 65520},
		{Move: Move(0x55f), Weight: 11562},
	},
	0x9c2f4cf1b8ab0639: []BookEntry{
		{Move: Move(0x51b), Weight: 65520},
	},
	0xbae07986b5375fce: []BookEntry{
		{Move: Move(0x14c), Weight: 58240},
		{Move: Move(0x3d7), Weight: 65520},
		{Move: Move(0x6e2), Weight: 21840},
	},
	0x421d4ecb41cda9c6: []BookEntry{
		{Move: Move(0xd2c), Weight: 9},
	},
	0xcb0d8210e9bb0b3a: []BookEntry{
		{Move: Move(0xe73), Weight: 2},
		{Move: Move(0xdef), Weight: 2},
	},
	0x6dd680b6a56ff29e: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0x8ad89ec061ca6c0d: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x9bf367f4b3164b49: []BookEntry{
		{Move: Move(0xca), Weight: 3},
	},
	0xa465a8c559d41098: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0xdef66fbb2f97ef0e: []BookEntry{
		{Move: Move(0x144), Weight: 1},
	},
	0x24877bc32edfeba5: []BookEntry{
		{Move: Move(0xb25), Weight: 1},
	},
	0x754012f374c7bdb6: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x8b1bcc78feddc676: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
	},
	0xfa4b692dc1e2c038: []BookEntry{
		{Move: Move(0x292), Weight: 2},
	},
	0xf90ee33671201ad9: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xfb04bbcd96baabf8: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0x7a7ee1102432674: []BookEntry{
		{Move: Move(0x195), Weight: 17},
		{Move: Move(0x292), Weight: 3},
	},
	0xb91d9c2d74f3cef: []BookEntry{
		{Move: Move(0x859), Weight: 2},
	},
	0x4046c4d92e66dae0: []BookEntry{
		{Move: Move(0x51b), Weight: 4},
	},
	0x4ffa40888519aa1b: []BookEntry{
		{Move: Move(0xea5), Weight: 3},
	},
	0x7d4951706df4b6d4: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xfde9cc451691cc10: []BookEntry{
		{Move: Move(0x161), Weight: 6},
	},
	0x44eb756f26f8c15: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x6c3d80e1294d4df4: []BookEntry{
		{Move: Move(0x94), Weight: 1},
		{Move: Move(0x3d7), Weight: 1},
	},
	0x95adb5e856557c88: []BookEntry{
		{Move: Move(0xb24), Weight: 1},
		{Move: Move(0xef2), Weight: 1},
	},
	0xdfd232f3152fdfc8: []BookEntry{
		{Move: Move(0x195), Weight: 1},
		{Move: Move(0x52), Weight: 1},
	},
	0xee66520fbdb434b0: []BookEntry{
		{Move: Move(0xef2), Weight: 2},
		{Move: Move(0xe3a), Weight: 1},
	},
	0xfb73e240dc52b440: []BookEntry{
		{Move: Move(0x252), Weight: 1},
	},
	0x1852fa65f350f464: []BookEntry{
		{Move: Move(0xca2), Weight: 1},
	},
	0xb1d10a39532680c8: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
		{Move: Move(0xe9e), Weight: 1},
	},
	0x844931a6ef4b9a0: []BookEntry{
		{Move: Move(0xfad), Weight: 65519},
		{Move: Move(0xe6a), Weight: 1},
	},
	0x4da06e8b2f0f38b0: []BookEntry{
		{Move: Move(0x89b), Weight: 8},
		{Move: Move(0xce3), Weight: 6},
	},
	0x88ce4b19abe8990b: []BookEntry{
		{Move: Move(0xb5c), Weight: 65520},
		{Move: Move(0xce3), Weight: 35280},
	},
	0xb70fcdbd9cf874c1: []BookEntry{
		{Move: Move(0x90), Weight: 1},
	},
	0x1864c1b3203b6d3b: []BookEntry{
		{Move: Move(0xe6a), Weight: 7},
		{Move: Move(0xf59), Weight: 1},
	},
	0x46479342c6f67823: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0x512fd77bd691ccdc: []BookEntry{
		{Move: Move(0xf6b), Weight: 2},
		{Move: Move(0x89b), Weight: 1},
		{Move: Move(0xf74), Weight: 1},
		{Move: Move(0xee9), Weight: 1},
		{Move: Move(0x9d5), Weight: 1},
	},
	0x64ac850d0e8768a3: []BookEntry{
		{Move: Move(0xf1c), Weight: 1},
	},
	0x8af3b83828e061d7: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0x301b3c500681d1f8: []BookEntry{
		{Move: Move(0xd6d), Weight: 65520},
		{Move: Move(0xd65), Weight: 3448},
	},
	0x6cea994cc5d6656c: []BookEntry{
		{Move: Move(0x4b), Weight: 65520},
		{Move: Move(0x49a), Weight: 35280},
	},
	0x7d55cdefd62168ee: []BookEntry{
		{Move: Move(0xd22), Weight: 1},
		{Move: Move(0xf3f), Weight: 1},
	},
	0x88474efe47e2c0c0: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0x8be88f868c45d5cf: []BookEntry{
		{Move: Move(0xf6b), Weight: 5},
		{Move: Move(0x91c), Weight: 2},
	},
	0x8d376bc94975c941: []BookEntry{
		{Move: Move(0x9ee), Weight: 1},
	},
	0xb3194d06be5b369a: []BookEntry{
		{Move: Move(0x49c), Weight: 1},
		{Move: Move(0x8106), Weight: 1},
	},
	0xd159c6f6501e8a9a: []BookEntry{
		{Move: Move(0xceb), Weight: 65520},
	},
	0x31606dfc03d2d72d: []BookEntry{
		{Move: Move(0x9d), Weight: 16},
	},
	0x36107858fb405950: []BookEntry{
		{Move: Move(0x2d3), Weight: 3},
	},
	0x8302628f474b6d26: []BookEntry{
		{Move: Move(0xd23), Weight: 65520},
	},
	0x99ab27ec2b49af53: []BookEntry{
		{Move: Move(0x15a), Weight: 1},
		{Move: Move(0x153), Weight: 1},
	},
	0xa4c382c494791a36: []BookEntry{
		{Move: Move(0xd6d), Weight: 1},
	},
	0xab8657a88608a7ed: []BookEntry{
		{Move: Move(0x29a), Weight: 3},
	},
	0xfc6561a8b43b8f7d: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x5e667c9f6949739b: []BookEntry{
		{Move: Move(0xcaa), Weight: 2},
	},
	0x7c193bd565b147d7: []BookEntry{
		{Move: Move(0x4b), Weight: 3},
	},
	0xb185d4b526cf23b2: []BookEntry{
		{Move: Move(0xae4), Weight: 1},
	},
	0x33668a084a54816: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x22faf9f2e535387a: []BookEntry{
		{Move: Move(0x6d8), Weight: 1},
	},
	0x649240fd41ee17ba: []BookEntry{
		{Move: Move(0xe6a), Weight: 4},
		{Move: Move(0xef2), Weight: 1},
		{Move: Move(0xf7c), Weight: 1},
	},
	0x7537cb727b34b8c6: []BookEntry{
		{Move: Move(0xf76), Weight: 4},
		{Move: Move(0xa9b), Weight: 1},
	},
	0x87eb7323cabb3176: []BookEntry{
		{Move: Move(0x61d), Weight: 1},
	},
	0x6649ba69b8c9ff8: []BookEntry{
		{Move: Move(0xca2), Weight: 65520},
		{Move: Move(0xfad), Weight: 7280},
	},
	0x7acbb41e0565aeb: []BookEntry{
		{Move: Move(0x49b), Weight: 1},
	},
	0x1929f7102b7629a2: []BookEntry{
		{Move: Move(0xce3), Weight: 65520},
	},
	0x50197b0107ccf149: []BookEntry{
		{Move: Move(0x91c), Weight: 65520},
		{Move: Move(0xe6a), Weight: 7280},
	},
	0x635f3330d1ef3db9: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xb0cceae0297eb493: []BookEntry{
		{Move: Move(0x845), Weight: 1},
	},
	0x76559684f673d0a7: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
		{Move: Move(0x8da), Weight: 1},
	},
	0x7c64e8d81f8a93c: []BookEntry{
		{Move: Move(0xe68), Weight: 1},
		{Move: Move(0xf74), Weight: 1},
		{Move: Move(0xeac), Weight: 1},
	},
	0x35edead8eeb03510: []BookEntry{
		{Move: Move(0x292), Weight: 2},
	},
	0x5a82726c7c233faf: []BookEntry{
		{Move: Move(0xb23), Weight: 1},
	},
	0x9abb0903e9f3e6c7: []BookEntry{
		{Move: Move(0x51b), Weight: 2},
	},
	0xa0d23168348ac7d9: []BookEntry{
		{Move: Move(0x9d), Weight: 2},
		{Move: Move(0x6e2), Weight: 1},
	},
	0xcb47c24a5c5b0de7: []BookEntry{
		{Move: Move(0xf38), Weight: 1},
	},
	0xf5b96ab8d235da65: []BookEntry{
		{Move: Move(0xceb), Weight: 1},
	},
	0xf9890a370cedb46f: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x268e0f6ef835612: []BookEntry{
		{Move: Move(0xc28), Weight: 1},
		{Move: Move(0xef3), Weight: 1},
		{Move: Move(0xda6), Weight: 1},
	},
	0xbe0f7eaa8b8a4cd: []BookEntry{
		{Move: Move(0x99e), Weight: 1},
	},
	0x37361bad0ed92adf: []BookEntry{
		{Move: Move(0x55b), Weight: 2},
	},
	0x99c989fd0216fe36: []BookEntry{
		{Move: Move(0xaa4), Weight: 1},
	},
	0xa8e7c91161416feb: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
		{Move: Move(0x14c), Weight: 11562},
	},
	0xc7d0f3b5998fcb3d: []BookEntry{
		{Move: Move(0xdae), Weight: 1},
		{Move: Move(0xeac), Weight: 1},
	},
	0x787ea30d0825faae: []BookEntry{
		{Move: Move(0x691), Weight: 3},
		{Move: Move(0x6a1), Weight: 1},
	},
	0xd00c496c419b2540: []BookEntry{
		{Move: Move(0x50), Weight: 1},
	},
	0x451dcb836adcc2f: []BookEntry{
		{Move: Move(0x51c), Weight: 1},
	},
	0x2f65988dd5723f58: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x3ff5fc9d8f1882b9: []BookEntry{
		{Move: Move(0x8b), Weight: 1},
	},
	0xc9eae20b017128a4: []BookEntry{
		{Move: Move(0xf76), Weight: 1},
	},
	0xdf3d80e2c8dc7525: []BookEntry{
		{Move: Move(0xcaa), Weight: 4},
		{Move: Move(0xe6a), Weight: 1},
	},
	0xf426005101ff2a92: []BookEntry{
		{Move: Move(0xe9e), Weight: 1},
		{Move: Move(0xf7c), Weight: 65520},
		{Move: Move(0xc20), Weight: 46800},
		{Move: Move(0xe73), Weight: 37440},
		{Move: Move(0xdef), Weight: 37440},
		{Move: Move(0x8a9), Weight: 9360},
	},
	0x6e73513c24c9e47: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0xb29582ddabd719e9: []BookEntry{
		{Move: Move(0xe6a), Weight: 5},
	},
	0xc7e4aa7114aaf05c: []BookEntry{
		{Move: Move(0x48c), Weight: 1},
	},
	0xec0f9b2019323c64: []BookEntry{
		{Move: Move(0x8e9), Weight: 4},
		{Move: Move(0xeac), Weight: 2},
		{Move: Move(0x8d2), Weight: 2},
	},
	0xfc1b09641ddeb986: []BookEntry{
		{Move: Move(0xf62), Weight: 3},
		{Move: Move(0xc20), Weight: 1},
		{Move: Move(0xd65), Weight: 1},
		{Move: Move(0xf74), Weight: 1},
	},
	0x39767dbe2ea9ae50: []BookEntry{
		{Move: Move(0xb1e), Weight: 1},
	},
	0x5c9324b5687c638e: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
		{Move: Move(0xf6b), Weight: 1},
	},
	0x80ddb75fb6ce67bb: []BookEntry{
		{Move: Move(0xb73), Weight: 1},
	},
	0xb6831a52f9e546c6: []BookEntry{
		{Move: Move(0xe6a), Weight: 65520},
	},
	0x268b00784b4c13b4: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0x28ec6d7651ded6bd: []BookEntry{
		{Move: Move(0x52), Weight: 6},
	},
	0x66b7fc10216d657a: []BookEntry{
		{Move: Move(0x218), Weight: 1},
		{Move: Move(0x2d3), Weight: 1},
	},
	0x8b3879078d8c3d03: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0xbb5f52faa383e932: []BookEntry{
		{Move: Move(0xf62), Weight: 1},
	},
	0xd9be1592d75a50f9: []BookEntry{
		{Move: Move(0xc28), Weight: 10},
	},
	0xdc207bff167f32d3: []BookEntry{
		{Move: Move(0x315), Weight: 1},
	},
	0x7dfba425a1fc6891: []BookEntry{
		{Move: Move(0xdef), Weight: 65520},
	},
	0xc056de0cdeca9b19: []BookEntry{
		{Move: Move(0x29a), Weight: 65520},
	},
	0xf0a27165cad6e47f: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
	},
	0xfb0095e193629d47: []BookEntry{
		{Move: Move(0xc61), Weight: 65520},
		{Move: Move(0xdef), Weight: 65520},
		{Move: Move(0xf3f), Weight: 1},
	},
	0xa95fa0cc74d4125: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x184d02bf8385cc74: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
		{Move: Move(0xee9), Weight: 1},
	},
	0x398760651060d466: []BookEntry{
		{Move: Move(0x2da), Weight: 1},
	},
	0xb41464df682eb2a0: []BookEntry{
		{Move: Move(0x91c), Weight: 11562},
		{Move: Move(0xe6a), Weight: 65520},
		{Move: Move(0xe73), Weight: 1},
	},
	0xbe333b6834871480: []BookEntry{
		{Move: Move(0x52), Weight: 65520},
	},
	0xdc36a4f02cf79900: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0xf8104159eff1ab96: []BookEntry{
		{Move: Move(0x688), Weight: 1},
	},
	0x23cb58e6836fce54: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
		{Move: Move(0xeb1), Weight: 1},
	},
	0x3658454605ecd054: []BookEntry{
		{Move: Move(0x89), Weight: 2},
	},
	0x53bbe52c4d69cb27: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x7d4fb6e585097082: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x36c8950e03f838c9: []BookEntry{
		{Move: Move(0xb63), Weight: 2},
	},
	0x4375d8e95f644936: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
		{Move: Move(0x766), Weight: 1},
	},
	0x8ffbb0d6ba2d8010: []BookEntry{
		{Move: Move(0xae3), Weight: 1},
		{Move: Move(0xf7c), Weight: 1},
	},
	0xbbd6e62422222dc8: []BookEntry{
		{Move: Move(0xb23), Weight: 1},
	},
	0x19d8470dd6d8f4d: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xc932d5cab5e6780f: []BookEntry{
		{Move: Move(0x6e5), Weight: 1},
	},
	0x167d305816050a8b: []BookEntry{
		{Move: Move(0xf7c), Weight: 65520},
		{Move: Move(0xe73), Weight: 26208},
		{Move: Move(0xc20), Weight: 26208},
		{Move: Move(0x899), Weight: 13104},
	},
	0x29d31a89f50e8726: []BookEntry{
		{Move: Move(0x355), Weight: 3},
	},
	0x2bcd5715dfd2c562: []BookEntry{
		{Move: Move(0xd2c), Weight: 10},
	},
	0x3d2ee72e5c8f6a2d: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0xcb7272d072071528: []BookEntry{
		{Move: Move(0xee9), Weight: 2},
		{Move: Move(0xf59), Weight: 1},
	},
	0x2714615a2e1f4fdc: []BookEntry{
		{Move: Move(0xcaa), Weight: 65520},
		{Move: Move(0xceb), Weight: 9360},
		{Move: Move(0xc20), Weight: 9360},
		{Move: Move(0xdef), Weight: 9360},
		{Move: Move(0xce3), Weight: 1},
	},
	0x7cba2eb005bd8406: []BookEntry{
		{Move: Move(0xd2c), Weight: 65520},
		{Move: Move(0xe6a), Weight: 39312},
		{Move: Move(0xe9e), Weight: 13104},
		{Move: Move(0xdae), Weight: 13104},
	},
	0xa9e7787a18ac7aa3: []BookEntry{
		{Move: Move(0x8106), Weight: 4},
	},
	0x3a469e6b28127b47: []BookEntry{
		{Move: Move(0x251), Weight: 6},
	},
	0x3d927a9082bee7bc: []BookEntry{
		{Move: Move(0x31c), Weight: 1},
	},
	0x516d5d7406978d04: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x7cee1d1a447d7e46: []BookEntry{
		{Move: Move(0x252), Weight: 1},
	},
	0xa12e30a185526bfd: []BookEntry{
		{Move: Move(0xa99), Weight: 1},
		{Move: Move(0x8d9), Weight: 65520},
	},
	0xcab0cb43c7df72b: []BookEntry{
		{Move: Move(0x66b), Weight: 1},
	},
	0xb0229a3f949d494d: []BookEntry{
		{Move: Move(0xdb), Weight: 5},
	},
	0x69597d94a0041361: []BookEntry{
		{Move: Move(0x91b), Weight: 2},
		{Move: Move(0xeb3), Weight: 1},
	},
	0x7b2cf8abc5c4b411: []BookEntry{
		{Move: Move(0x195), Weight: 42},
	},
	0xbe1ff05927f1410: []BookEntry{
		{Move: Move(0xf74), Weight: 4},
	},
	0x9328e412a7225c2c: []BookEntry{
		{Move: Move(0x91b), Weight: 1},
	},
	0xed6bcb747ed398c5: []BookEntry{
		{Move: Move(0x8e9), Weight: 2},
	},
	0x3858148d1504f5ed: []BookEntry{
		{Move: Move(0x314), Weight: 3},
	},
	0x5dad3e31882b95a9: []BookEntry{
		{Move: Move(0x15a), Weight: 8},
		{Move: Move(0x259), Weight: 2},
		{Move: Move(0x2d3), Weight: 2},
	},
	0x999dbae0ba7754e6: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
	},
	0xd0378562bc5e25a7: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
		{Move: Move(0xdef), Weight: 1},
	},
	0xe87d1fd26b906484: []BookEntry{
		{Move: Move(0x51b), Weight: 3},
	},
	0xf038a83090106a39: []BookEntry{
		{Move: Move(0x31c), Weight: 1},
	},
	0x71b439757eefac4: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x37fa6da596ffd415: []BookEntry{
		{Move: Move(0x8106), Weight: 3},
		{Move: Move(0x251), Weight: 2},
	},
	0x3d14712103bf0bee: []BookEntry{
		{Move: Move(0x51c), Weight: 1},
	},
	0x4577b8b352ddd3c1: []BookEntry{
		{Move: Move(0x14e), Weight: 1},
	},
	0x6a84818e3fdf85a7: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x6b24c6c1467a5085: []BookEntry{
		{Move: Move(0xae4), Weight: 1},
	},
	0xa45526d823e33f3c: []BookEntry{
		{Move: Move(0x566), Weight: 1},
	},
	0xbbf719d404992d74: []BookEntry{
		{Move: Move(0x195), Weight: 17},
	},
	0x45aa8f3fa6ec533: []BookEntry{
		{Move: Move(0x6d1), Weight: 1},
	},
	0xc8162c4989019aab: []BookEntry{
		{Move: Move(0x195), Weight: 8},
	},
	0xe0377d831132d56c: []BookEntry{
		{Move: Move(0x544), Weight: 1},
	},
	0x310f3f8d8e30be26: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x5e94651ae00442ce: []BookEntry{
		{Move: Move(0xaa0), Weight: 35280},
		{Move: Move(0xdef), Weight: 65520},
	},
	0x9734f1704832688a: []BookEntry{
		{Move: Move(0xfad), Weight: 8},
	},
	0xa61a9ff7b418c4f8: []BookEntry{
		{Move: Move(0x754), Weight: 1},
	},
	0xbce451b3442e6dc1: []BookEntry{
		{Move: Move(0x29a), Weight: 2},
		{Move: Move(0x251), Weight: 1},
	},
	0xe86501c6340f92e3: []BookEntry{
		{Move: Move(0x688), Weight: 1},
	},
	0x22b4a107b6be6492: []BookEntry{
		{Move: Move(0xd1), Weight: 1},
	},
	0x4176447cebe1e823: []BookEntry{
		{Move: Move(0xcaa), Weight: 4},
	},
	0x71b1bcbc99da396e: []BookEntry{
		{Move: Move(0x314), Weight: 3},
		{Move: Move(0x52), Weight: 2},
	},
	0xb0fbaa7cd799ac04: []BookEntry{
		{Move: Move(0x195), Weight: 65520},
	},
	0xbd866aa42a506c9a: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0xe722bb1faff3b601: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0xce3691555d99f697: []BookEntry{
		{Move: Move(0xd1), Weight: 2},
	},
	0xe73a253b8a1a5a3: []BookEntry{
		{Move: Move(0x18c), Weight: 1},
	},
	0xb5b0ebbc908b9da9: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xe9ada06dd8965860: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0xf3b92e1610d76021: []BookEntry{
		{Move: Move(0xae2), Weight: 65520},
	},
	0xdd369ab2ea11b3cd: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x3ef5e9285ba4cc2: []BookEntry{
		{Move: Move(0x94), Weight: 3},
	},
	0x317e46cbf4fafc86: []BookEntry{
		{Move: Move(0x2db), Weight: 19},
	},
	0x356823b60e2e1806: []BookEntry{
		{Move: Move(0xaa2), Weight: 1},
	},
	0x4290c7eb332ab307: []BookEntry{
		{Move: Move(0x691), Weight: 5},
	},
	0x70d2d00f6647499b: []BookEntry{
		{Move: Move(0x14c), Weight: 5},
	},
	0x9f05585f830c9563: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
		{Move: Move(0x259), Weight: 1},
	},
	0xc11788ada45c0eb1: []BookEntry{
		{Move: Move(0xc69), Weight: 4},
	},
	0xd1976c0221e20ca2: []BookEntry{
		{Move: Move(0x292), Weight: 65520},
		{Move: Move(0x251), Weight: 35280},
	},
	0xb1fb5640840ff9: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x1f6a1ab5c4f32fcf: []BookEntry{
		{Move: Move(0xf74), Weight: 12},
	},
	0x3ce966319024bd91: []BookEntry{
		{Move: Move(0x218), Weight: 65520},
	},
	0x67c4124bd2e4b4cf: []BookEntry{
		{Move: Move(0x94), Weight: 1},
	},
	0x67d042f54855b86a: []BookEntry{
		{Move: Move(0xd3), Weight: 3},
	},
	0x86933069b8c862bb: []BookEntry{
		{Move: Move(0xce3), Weight: 42},
	},
	0xbfb7b0f645fec6b8: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x497986937fdc04fb: []BookEntry{
		{Move: Move(0x91c), Weight: 3},
	},
	0x5f0d6a71f3152c17: []BookEntry{
		{Move: Move(0xc61), Weight: 1},
	},
	0x8304d6e7a949f4b6: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
		{Move: Move(0xe73), Weight: 1},
	},
	0x8b2399e9a2df9661: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0xb3b10f728924caac: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x161772a00d8df359: []BookEntry{
		{Move: Move(0xd3), Weight: 1},
		{Move: Move(0x49b), Weight: 65520},
	},
	0x2b3bc2ed78af6750: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0x8ddebe5d1e13c1a6: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x91fe8c5eb58422ce: []BookEntry{
		{Move: Move(0xca2), Weight: 1},
	},
	0xac004cca449aae45: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0xbafe2512df61bb39: []BookEntry{
		{Move: Move(0xf74), Weight: 3},
		{Move: Move(0xc28), Weight: 2},
		{Move: Move(0xc69), Weight: 2},
		{Move: Move(0xf6b), Weight: 2},
	},
	0xe6685e5187aa2d0c: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0xe789c585c18c0560: []BookEntry{
		{Move: Move(0xb1e), Weight: 1},
	},
	0x1558f996ed748334: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0xa6a90e5475b4d804: []BookEntry{
		{Move: Move(0xf76), Weight: 19},
	},
	0xb00a34eeb8f7bd65: []BookEntry{
		{Move: Move(0xc28), Weight: 1},
		{Move: Move(0xcaa), Weight: 1},
	},
	0x4cdd1fb6438c9a12: []BookEntry{
		{Move: Move(0x89), Weight: 1},
		{Move: Move(0x661), Weight: 1},
	},
	0x38fca71019edd0f4: []BookEntry{
		{Move: Move(0xef4), Weight: 2},
		{Move: Move(0xe9e), Weight: 1},
	},
	0x3fd5555383e615b3: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x64a3ae844e34aaf0: []BookEntry{
		{Move: Move(0x94), Weight: 65519},
		{Move: Move(0x14c), Weight: 5697},
	},
	0xc0c7e80f36d960bf: []BookEntry{
		{Move: Move(0x8106), Weight: 4},
	},
	0x69ada09749b65892: []BookEntry{
		{Move: Move(0x55b), Weight: 2},
	},
	0x8422ef6582f2bbd2: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x8a51a000aa638992: []BookEntry{
		{Move: Move(0xb25), Weight: 1},
	},
	0xa640f5c6b61e52a2: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xfa4e22fb5a2a5135: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x2622382f832562ea: []BookEntry{
		{Move: Move(0x45a), Weight: 1},
		{Move: Move(0x314), Weight: 1},
	},
	0x3832e876ba05a333: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x5822a6647dfdbd7b: []BookEntry{
		{Move: Move(0xe7), Weight: 1},
	},
	0x768933b188163242: []BookEntry{
		{Move: Move(0xaf2), Weight: 1},
	},
	0xc69cc2108100279b: []BookEntry{
		{Move: Move(0xce5), Weight: 2},
		{Move: Move(0xef2), Weight: 2},
		{Move: Move(0xee9), Weight: 1},
	},
	0x29be28bcee1607e7: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x46f22bb243af7e1a: []BookEntry{
		{Move: Move(0x995), Weight: 1},
	},
	0x5069120c1ad91a6d: []BookEntry{
		{Move: Move(0x89), Weight: 7},
		{Move: Move(0x6a3), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0x854afc715e267cb4: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
		{Move: Move(0xb73), Weight: 1},
	},
	0x93be3e9156758b22: []BookEntry{
		{Move: Move(0x52), Weight: 1},
		{Move: Move(0x259), Weight: 1},
	},
	0xc1a56fa73ce2e2d1: []BookEntry{
		{Move: Move(0x8106), Weight: 3},
	},
	0x24c1a29cc124eda: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x99e851751fca1fd: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x1afceb8e7efa9771: []BookEntry{
		{Move: Move(0x8d2), Weight: 7},
		{Move: Move(0xe6a), Weight: 1},
	},
	0x3ff2f1aa61239319: []BookEntry{
		{Move: Move(0x688), Weight: 1},
	},
	0x4be4c06744078347: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0x7446cd7480dde39f: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x8de7432d7a577955: []BookEntry{
		{Move: Move(0xdae), Weight: 2},
	},
	0xbffe4a9570e802c0: []BookEntry{
		{Move: Move(0xd5), Weight: 3},
	},
	0x228e5ea5fb0a13ae: []BookEntry{
		{Move: Move(0xdae), Weight: 65520},
		{Move: Move(0xeac), Weight: 65520},
		{Move: Move(0xe68), Weight: 65520},
		{Move: Move(0xf7c), Weight: 65520},
	},
	0x33840620182806a9: []BookEntry{
		{Move: Move(0xca2), Weight: 3},
	},
	0x39cdadf0275fe683: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x73018e98fb2089b3: []BookEntry{
		{Move: Move(0xd23), Weight: 65520},
	},
	0x75145e750328578b: []BookEntry{
		{Move: Move(0x54b), Weight: 8},
	},
	0xdf55adb27214dc34: []BookEntry{
		{Move: Move(0x757), Weight: 1},
	},
	0x1260813c1ba8d3d: []BookEntry{
		{Move: Move(0x314), Weight: 2},
	},
	0x4ae5bbb723b3645: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xb9b1a0b16b02df71: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xef300c7965922162: []BookEntry{
		{Move: Move(0x4b), Weight: 9},
	},
	0xf99a4d4409b71542: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0x4b6381081861851d: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x50aeb0cc5f1a7f77: []BookEntry{
		{Move: Move(0x663), Weight: 1},
	},
	0xc1f25cb0e47f9ae2: []BookEntry{
		{Move: Move(0xc61), Weight: 1},
	},
	0x3bf2436b7637104b: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
	},
	0xa69c4030e8dc722a: []BookEntry{
		{Move: Move(0xf76), Weight: 1},
	},
	0xc8378ceba70dfd1e: []BookEntry{
		{Move: Move(0x52), Weight: 7},
		{Move: Move(0x2d3), Weight: 2},
	},
	0xcfe7bce4ecbc862e: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x59cc85901d75669: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x3aff09b23462353e: []BookEntry{
		{Move: Move(0xd2c), Weight: 3},
	},
	0x4acbda5a3ee3508f: []BookEntry{
		{Move: Move(0xe9e), Weight: 2},
	},
	0x933f384b634c232d: []BookEntry{
		{Move: Move(0x6a3), Weight: 3},
	},
	0x22b4e052f18fab1d: []BookEntry{
		{Move: Move(0xfad), Weight: 33},
		{Move: Move(0xe6a), Weight: 22},
		{Move: Move(0xdae), Weight: 1},
	},
	0x29a86ea94f97603b: []BookEntry{
		{Move: Move(0xf76), Weight: 1},
	},
	0x43ed81ed5622df19: []BookEntry{
		{Move: Move(0xef2), Weight: 65520},
		{Move: Move(0xf6b), Weight: 23400},
		{Move: Move(0xe73), Weight: 4680},
	},
	0xcc21827e1eb84aeb: []BookEntry{
		{Move: Move(0xb63), Weight: 3},
	},
	0xf79b36f680f86d16: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x7f5277a1ede7e8d: []BookEntry{
		{Move: Move(0x14e), Weight: 4},
	},
	0xa04ff02cf0a027a6: []BookEntry{
		{Move: Move(0xce2), Weight: 5},
	},
	0xe8ea548090e8f1b0: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x1c02e7931ed8ab02: []BookEntry{
		{Move: Move(0x218), Weight: 1},
	},
	0x5eaada61a64b3e89: []BookEntry{
		{Move: Move(0x3df), Weight: 4},
		{Move: Move(0x4b), Weight: 1},
		{Move: Move(0x14c), Weight: 1},
	},
	0x2d9e89f573f00cfe: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x6697b6f0f6faa98c: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xad8fce31a01695: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x40f70817b15292d8: []BookEntry{
		{Move: Move(0x314), Weight: 65520},
	},
	0xa756c29108af3d5b: []BookEntry{
		{Move: Move(0xd1), Weight: 1},
	},
	0xac875601b38bbc3: []BookEntry{
		{Move: Move(0x44a), Weight: 1},
	},
	0x2ef2e9891d3c7083: []BookEntry{
		{Move: Move(0xcaa), Weight: 2},
	},
	0xd8e83cd0e742a48b: []BookEntry{
		{Move: Move(0xce3), Weight: 65520},
		{Move: Move(0xe6a), Weight: 30240},
	},
	0x758c536f64d4ade1: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0x9b93e4e1d9b4f3e5: []BookEntry{
		{Move: Move(0xfad), Weight: 65520},
	},
	0xf37bb19bb8d44c7a: []BookEntry{
		{Move: Move(0x55b), Weight: 4},
	},
	0xfd66c9d9562d0363: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0xfc4621328f6790: []BookEntry{
		{Move: Move(0xb24), Weight: 1},
	},
	0x3cc863e808c9db35: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x7751fa22d9981d6f: []BookEntry{
		{Move: Move(0xdae), Weight: 2},
		{Move: Move(0xc28), Weight: 2},
	},
	0x98b9676ccf42357a: []BookEntry{
		{Move: Move(0x153), Weight: 6},
	},
	0xf435c7f18061f6bf: []BookEntry{
		{Move: Move(0x766), Weight: 1},
	},
	0x5908f38a5d6c2e53: []BookEntry{
		{Move: Move(0x54b), Weight: 1},
	},
	0x87a10f57046abba5: []BookEntry{
		{Move: Move(0xf62), Weight: 1},
	},
	0xe2f65ff891026e23: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0xfb00b8465975f0aa: []BookEntry{
		{Move: Move(0xaa0), Weight: 1},
	},
	0xfea37baea55c6f3d: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0x7bd618ae257ddd5: []BookEntry{
		{Move: Move(0x722), Weight: 1},
	},
	0xa68280807fe81ce: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x2bff25be1310c925: []BookEntry{
		{Move: Move(0x86a), Weight: 1},
	},
	0xb85da2714b660e26: []BookEntry{
		{Move: Move(0x498), Weight: 2},
	},
	0xe5f771ba2f8496ee: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x4adfabd7fbc45bda: []BookEntry{
		{Move: Move(0x92e), Weight: 6},
	},
	0x56f6270f5c893e0d: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x252c1d8b9acb78fa: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x1254a9c5c8ce8e0d: []BookEntry{
		{Move: Move(0xceb), Weight: 3},
	},
	0xe06314985a5c4765: []BookEntry{
		{Move: Move(0xc28), Weight: 1},
	},
	0x1bd2f7dab6b8ab5: []BookEntry{
		{Move: Move(0x927), Weight: 1},
	},
	0x1803b518fbd89e8f: []BookEntry{
		{Move: Move(0x195), Weight: 2},
	},
	0x2bff014d139d88f6: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
		{Move: Move(0x795), Weight: 1},
	},
	0xa7b5bab80bcf76bc: []BookEntry{
		{Move: Move(0xeed), Weight: 2},
	},
	0xb7e73ef3899c80a9: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xe15688b08a044755: []BookEntry{
		{Move: Move(0xd5), Weight: 1},
	},
	0x3464eaf8fcf2cb01: []BookEntry{
		{Move: Move(0xc28), Weight: 2},
	},
	0x7c50067df37e1b9c: []BookEntry{
		{Move: Move(0x49b), Weight: 1},
	},
	0xaca114247eba3196: []BookEntry{
		{Move: Move(0xf76), Weight: 1},
	},
	0xe4cebfb868eddd5b: []BookEntry{
		{Move: Move(0x8106), Weight: 3},
	},
	0x94aea6091c794078: []BookEntry{
		{Move: Move(0x195), Weight: 3},
	},
	0xd7970105350e1e28: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0xf08a906f22298e1: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x41f83c098ab9cac3: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x5497147d3b9415d6: []BookEntry{
		{Move: Move(0x161), Weight: 2},
	},
	0xcaf142667b9accfc: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0xdbb4d0319f7a3365: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0xe1282c6aa8f2461e: []BookEntry{
		{Move: Move(0x2dc), Weight: 1},
	},
	0xe2ed194e77ac02a9: []BookEntry{
		{Move: Move(0x2db), Weight: 2},
		{Move: Move(0x8106), Weight: 1},
	},
	0x698da2a0265229c: []BookEntry{
		{Move: Move(0x14c), Weight: 3},
	},
	0x18c9c218a67ed390: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0x1d97c18b27b22007: []BookEntry{
		{Move: Move(0x252), Weight: 2},
	},
	0x24f56351cd7e1348: []BookEntry{
		{Move: Move(0xef3), Weight: 1},
	},
	0x277d88be219ff543: []BookEntry{
		{Move: Move(0xef4), Weight: 4},
		{Move: Move(0xf3f), Weight: 1},
	},
	0xad2ac0dc5dab4b45: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
		{Move: Move(0xce2), Weight: 1},
	},
	0xe98052020a558fe1: []BookEntry{
		{Move: Move(0x89), Weight: 1},
		{Move: Move(0x90), Weight: 1},
	},
	0xb3844d49d2be0d0: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x543c88ad355dab25: []BookEntry{
		{Move: Move(0x8106), Weight: 65520},
	},
	0x5a130203501d87de: []BookEntry{
		{Move: Move(0x15a), Weight: 1},
	},
	0x5f7f2012af6208d4: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
		{Move: Move(0x6e2), Weight: 1},
	},
	0xc1aad7383d89e2fb: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
		{Move: Move(0xe73), Weight: 1},
	},
	0xe97f22769fac82a2: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0x260fc6e7a52f3775: []BookEntry{
		{Move: Move(0x713), Weight: 3},
	},
	0x58776855a7e6eecf: []BookEntry{
		{Move: Move(0xd3), Weight: 3},
	},
	0x743fea1104be7acb: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0x8b68686fb4868c1a: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
		{Move: Move(0xef4), Weight: 1},
	},
	0xb3419d7cab23211e: []BookEntry{
		{Move: Move(0x292), Weight: 5},
		{Move: Move(0x55f), Weight: 2},
	},
	0xc50cb8b1b9ab5d45: []BookEntry{
		{Move: Move(0x259), Weight: 1},
	},
	0xcdd6cd2d4e4045e: []BookEntry{
		{Move: Move(0xf74), Weight: 5},
		{Move: Move(0xd6d), Weight: 3},
	},
	0x1ec9ea3de349fdd3: []BookEntry{
		{Move: Move(0x66b), Weight: 1},
	},
	0x4112df7bd300d6e7: []BookEntry{
		{Move: Move(0x15a), Weight: 8},
		{Move: Move(0x195), Weight: 8},
	},
	0xa49a95123a83d7b5: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
		{Move: Move(0xea5), Weight: 1},
	},
	0xacea65f12c86f88: []BookEntry{
		{Move: Move(0x54b), Weight: 6},
	},
	0x13a43b0df61e50da: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0x1cb8a5f6d4f5263f: []BookEntry{
		{Move: Move(0xb63), Weight: 42},
	},
	0x935f6e629e0fc568: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x43b1d0aecaa8ce10: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0xbf6814a4a8f9f19b: []BookEntry{
		{Move: Move(0xae2), Weight: 1},
	},
	0xe782f8e137a5f793: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x473ae6dfce8a3ba: []BookEntry{
		{Move: Move(0xcaa), Weight: 2},
	},
	0x1d6bf345eb0eae1a: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x51b4826398e9d333: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xbbcb748a38b75e4d: []BookEntry{
		{Move: Move(0xd2c), Weight: 7},
	},
	0xccdffb03ff8f9483: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x39ff7c6b9e37ee8b: []BookEntry{
		{Move: Move(0x52), Weight: 65520},
	},
	0xb448fd989f186eb3: []BookEntry{
		{Move: Move(0x94), Weight: 1},
	},
	0xfed1f107e22f1ee0: []BookEntry{
		{Move: Move(0x724), Weight: 1},
	},
	0xa84922ee9429ade: []BookEntry{
		{Move: Move(0xd2d), Weight: 1},
	},
	0xd87103a6610271f: []BookEntry{
		{Move: Move(0xca), Weight: 1},
	},
	0x1690f8cf93e9f527: []BookEntry{
		{Move: Move(0x564), Weight: 1},
	},
	0x9fa8bd4c359d6f20: []BookEntry{
		{Move: Move(0xd1), Weight: 3},
	},
	0xabcd463c7d3b23f1: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0xd6b140a4fc712efc: []BookEntry{
		{Move: Move(0x8a9), Weight: 3},
	},
	0xe236dc8dafc5d42b: []BookEntry{
		{Move: Move(0x9ee), Weight: 1},
	},
	0x143d3e0edbc4a8bc: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x164b3c91d4ba493d: []BookEntry{
		{Move: Move(0x153), Weight: 8},
		{Move: Move(0xd1), Weight: 2},
	},
	0x34ec218438949b6d: []BookEntry{
		{Move: Move(0xda6), Weight: 2},
	},
	0x37e81b1804b32028: []BookEntry{
		{Move: Move(0x756), Weight: 6},
	},
	0x6091c40904ca46bb: []BookEntry{
		{Move: Move(0xcaa), Weight: 65520},
		{Move: Move(0xe73), Weight: 65520},
	},
	0x80790d677ad41bab: []BookEntry{
		{Move: Move(0xef3), Weight: 1},
		{Move: Move(0x8da), Weight: 1},
		{Move: Move(0xf74), Weight: 1},
		{Move: Move(0xc28), Weight: 1},
	},
	0xe2ab9aea075704e8: []BookEntry{
		{Move: Move(0xc69), Weight: 10},
	},
	0xd6accac9a404a7e: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0x183de96daf43e744: []BookEntry{
		{Move: Move(0xfad), Weight: 32760},
		{Move: Move(0xe6a), Weight: 49140},
		{Move: Move(0xd2c), Weight: 65520},
		{Move: Move(0x89b), Weight: 16380},
	},
	0x626c8ae3742eeb2f: []BookEntry{
		{Move: Move(0x259), Weight: 3},
	},
	0x6ed71dcaa7462de8: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0xfff8dbdeb5be4541: []BookEntry{
		{Move: Move(0x251), Weight: 2},
	},
	0x68ce7201fe999de: []BookEntry{
		{Move: Move(0xd2c), Weight: 9},
		{Move: Move(0xef2), Weight: 1},
	},
	0x4df2457031dc9041: []BookEntry{
		{Move: Move(0xc61), Weight: 1},
	},
	0x51f5e04d35d4f163: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xb8f4112ea93e0ad6: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x2ae68e3e7a03b5f: []BookEntry{
		{Move: Move(0x55f), Weight: 5},
	},
	0x17c679fef69a3761: []BookEntry{
		{Move: Move(0x161), Weight: 1},
	},
	0x34f838cab7461708: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
	},
	0x4f15a5adada948d0: []BookEntry{
		{Move: Move(0x3c6), Weight: 1},
	},
	0x56d695fa086fd775: []BookEntry{
		{Move: Move(0xf7c), Weight: 1},
	},
	0x5a95801dbb6966bc: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0x61d81dc2182f3f5d: []BookEntry{
		{Move: Move(0x91b), Weight: 3},
	},
	0xf72b67b227936a81: []BookEntry{
		{Move: Move(0x691), Weight: 2},
	},
	0x475aa8c7e48579aa: []BookEntry{
		{Move: Move(0x6ea), Weight: 1},
		{Move: Move(0xa6), Weight: 1},
	},
	0x79aaac85d69a4b42: []BookEntry{
		{Move: Move(0x925), Weight: 1},
	},
	0x99249479a1e19c5e: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x9fe39e1e6381380d: []BookEntry{
		{Move: Move(0xec3), Weight: 2},
	},
	0xdd007750f36afec5: []BookEntry{
		{Move: Move(0x195), Weight: 41694},
		{Move: Move(0x314), Weight: 65519},
		{Move: Move(0x292), Weight: 11912},
	},
	0xe2ea9f51b01c42da: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x51ce99c2fe6f405a: []BookEntry{
		{Move: Move(0x2db), Weight: 65520},
		{Move: Move(0x2d3), Weight: 11562},
	},
	0x5e133d3801514648: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0x92ab09766a8e678a: []BookEntry{
		{Move: Move(0x6e2), Weight: 2},
	},
	0x941fcc1672a00ad3: []BookEntry{
		{Move: Move(0x8e9), Weight: 65520},
	},
	0x94716541b05f5f63: []BookEntry{
		{Move: Move(0x6e3), Weight: 1},
	},
	0xa3ae59c88f83e0d5: []BookEntry{
		{Move: Move(0x6e4), Weight: 1},
	},
	0xa9f86397c65e014a: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
		{Move: Move(0x90), Weight: 1},
	},
	0xb1f3605e9e4fe9e8: []BookEntry{
		{Move: Move(0xeb3), Weight: 2},
	},
	0x136c634da531aa53: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x31b161becdfac2bd: []BookEntry{
		{Move: Move(0x50), Weight: 2},
	},
	0x68e640b4e709b12d: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0x7ad274fbeedf08c4: []BookEntry{
		{Move: Move(0xe73), Weight: 50960},
		{Move: Move(0xca2), Weight: 65520},
		{Move: Move(0xe6a), Weight: 29120},
	},
	0xad8b26fa5d14d7f5: []BookEntry{
		{Move: Move(0xf3f), Weight: 3},
		{Move: Move(0xe6a), Weight: 1},
		{Move: Move(0x70b), Weight: 1},
	},
	0xc5e0225ccc4bca86: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0xd4aad69cef2d3928: []BookEntry{
		{Move: Move(0xce9), Weight: 18720},
		{Move: Move(0xcaa), Weight: 65520},
		{Move: Move(0xf7c), Weight: 9360},
	},
	0xf160c563622c28d6: []BookEntry{
		{Move: Move(0x8106), Weight: 65520},
	},
	0x4082900f1eebbf07: []BookEntry{
		{Move: Move(0x652), Weight: 5},
	},
	0x463b96181691fc9c: []BookEntry{
		{Move: Move(0xd24), Weight: 65520},
		{Move: Move(0xfad), Weight: 17035},
		{Move: Move(0xce3), Weight: 39312},
		{Move: Move(0xca2), Weight: 9172},
	},
	0x69596bc8f1bd07d1: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
		{Move: Move(0x52), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0x8f7cc452f7af6a54: []BookEntry{
		{Move: Move(0x6c3), Weight: 1},
	},
	0xb966915ccd0da452: []BookEntry{
		{Move: Move(0x936), Weight: 1},
	},
	0xd84dc5c57c1e6952: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x338807bccd84b278: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x76a18fe5099eb6e1: []BookEntry{
		{Move: Move(0x8106), Weight: 3},
	},
	0x854eda82375adce5: []BookEntry{
		{Move: Move(0xc69), Weight: 1},
	},
	0x4062c048d9622d6b: []BookEntry{
		{Move: Move(0xd19), Weight: 1},
	},
	0xce3439e8d69583e9: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0xd3fe5dce21e70ce5: []BookEntry{
		{Move: Move(0x2dc), Weight: 3},
	},
	0x10eff588533cfa22: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x1c4501258e9ca464: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x4b1376a17217ee1d: []BookEntry{
		{Move: Move(0x210), Weight: 65520},
		{Move: Move(0x195), Weight: 43680},
	},
	0xa9372747c0084bc5: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xfd914606a91e7b27: []BookEntry{
		{Move: Move(0x6ea), Weight: 2},
	},
	0x18c794a8b162fa6f: []BookEntry{
		{Move: Move(0x195), Weight: 5},
	},
	0x95078406c795ff3d: []BookEntry{
		{Move: Move(0xc61), Weight: 1},
	},
	0xa8d861e0b13c3783: []BookEntry{
		{Move: Move(0xf76), Weight: 2},
	},
	0xb70c6acdfb843105: []BookEntry{
		{Move: Move(0xdef), Weight: 3},
		{Move: Move(0xce3), Weight: 1},
	},
	0xe2ed249e5995556a: []BookEntry{
		{Move: Move(0x9ee), Weight: 1},
	},
	0x16e636381ce8d6e9: []BookEntry{
		{Move: Move(0x52), Weight: 7},
	},
	0x2e0275f7df671e17: []BookEntry{
		{Move: Move(0x91c), Weight: 1},
	},
	0x52fc0dcc18443f07: []BookEntry{
		{Move: Move(0xc69), Weight: 1},
	},
	0x98c46f5aa7744acb: []BookEntry{
		{Move: Move(0x688), Weight: 1},
	},
	0xa06f3410eeff3740: []BookEntry{
		{Move: Move(0xd2c), Weight: 4},
		{Move: Move(0xe9e), Weight: 1},
	},
	0xda48997503d0: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x43fa35bbb0b57ff5: []BookEntry{
		{Move: Move(0x832), Weight: 1},
	},
	0x57105cffa5de5f15: []BookEntry{
		{Move: Move(0x829), Weight: 4},
	},
	0x644d4afe02564aeb: []BookEntry{
		{Move: Move(0xfad), Weight: 65519},
		{Move: Move(0xcaa), Weight: 2397},
		{Move: Move(0xe6a), Weight: 11985},
	},
	0xfa6866be133e26e0: []BookEntry{
		{Move: Move(0x316), Weight: 1},
	},
	0x217c9925598d53b0: []BookEntry{
		{Move: Move(0x18c), Weight: 5},
	},
	0x91fd9d423ea6606e: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0xb07a7c7ca85bf4de: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0xe8c30b0249c44ff0: []BookEntry{
		{Move: Move(0xc61), Weight: 65520},
		{Move: Move(0xe73), Weight: 65520},
	},
	0xb3c6b529ebe4331: []BookEntry{
		{Move: Move(0x52), Weight: 65520},
	},
	0x3372beae34f36620: []BookEntry{
		{Move: Move(0x2d3), Weight: 2},
	},
	0x43b75c2448df0991: []BookEntry{
		{Move: Move(0x766), Weight: 1},
	},
	0xc6b14e1bd38ddc37: []BookEntry{
		{Move: Move(0xce3), Weight: 7},
	},
	0xd8e2cb360dea06b9: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xe9e9857ba1cbb4f8: []BookEntry{
		{Move: Move(0xeb1), Weight: 3},
	},
	0xece8523fbce10b03: []BookEntry{
		{Move: Move(0xd1), Weight: 1},
	},
	0x143b92313c312e93: []BookEntry{
		{Move: Move(0x161), Weight: 1},
		{Move: Move(0x6ea), Weight: 1},
	},
	0x43d6b740a99a5421: []BookEntry{
		{Move: Move(0x48c), Weight: 1},
	},
	0x911f7693d6608d6a: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0xf13b64bf5cc619c5: []BookEntry{
		{Move: Move(0x195), Weight: 5},
	},
	0xfac0e1aa157e48e5: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0x1abbe9d726e89f85: []BookEntry{
		{Move: Move(0x653), Weight: 1},
		{Move: Move(0x668), Weight: 65520},
	},
	0xca7264f534b3f547: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xcc94dbfc2efdfb48: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0xd8e8e1eec422180a: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0xda9f68d56907cc4c: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
	},
	0x30cb9adf6f7a7852: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x576360497e75d991: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xae2c0a697f9cf35f: []BookEntry{
		{Move: Move(0xe6a), Weight: 3},
	},
	0xbfedf6fe55381c6f: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0xf91b2be15903fa6f: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x683ea74b070cd5: []BookEntry{
		{Move: Move(0x259), Weight: 1},
	},
	0x25c05d79bd88390: []BookEntry{
		{Move: Move(0xf7c), Weight: 65520},
	},
	0x13393868ac2036cf: []BookEntry{
		{Move: Move(0x92a), Weight: 1},
	},
	0x21bf172442e98ec2: []BookEntry{
		{Move: Move(0x6e2), Weight: 65520},
	},
	0x66934f3ee39d46ca: []BookEntry{
		{Move: Move(0xceb), Weight: 9},
	},
	0x9508799b635fa163: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
	},
	0xab30fd8fa233a149: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xec104f97b500572a: []BookEntry{
		{Move: Move(0xfad), Weight: 3},
	},
	0x2d8eedaeab65aded: []BookEntry{
		{Move: Move(0x218), Weight: 1},
	},
	0x569f9910e9398daf: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xca6f05b8e3808954: []BookEntry{
		{Move: Move(0xf3f), Weight: 4},
	},
	0x1ddd137af68b89c9: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
		{Move: Move(0x314), Weight: 1},
	},
	0x57d6737a3514d762: []BookEntry{
		{Move: Move(0xad0), Weight: 1},
	},
	0xc1ea8f5a9e6b846a: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0x6349ea8eb429571e: []BookEntry{
		{Move: Move(0xf59), Weight: 1},
	},
	0x68c2aabe2fb117a9: []BookEntry{
		{Move: Move(0x51b), Weight: 1},
	},
	0x93ad77d9e7ce1631: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0x95dd03fe7583f5cf: []BookEntry{
		{Move: Move(0xd24), Weight: 3},
	},
	0xe4c5ce7c4cdf1959: []BookEntry{
		{Move: Move(0xd5), Weight: 3},
	},
	0x777a86581c6d145d: []BookEntry{
		{Move: Move(0xeed), Weight: 1},
		{Move: Move(0xef4), Weight: 1},
	},
	0x150a2372f9e693ea: []BookEntry{
		{Move: Move(0xaa0), Weight: 1},
	},
	0x51da0eb69e0348c5: []BookEntry{
		{Move: Move(0x6e4), Weight: 3},
	},
	0x569a1d1a0c93b50d: []BookEntry{
		{Move: Move(0x144), Weight: 2},
		{Move: Move(0x2d3), Weight: 1},
	},
	0xabe05a0245180384: []BookEntry{
		{Move: Move(0xd65), Weight: 1},
		{Move: Move(0xeac), Weight: 1},
		{Move: Move(0xeed), Weight: 1},
	},
	0xcdd35b108a043c9d: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0xda8e6dce9eaf6a4a: []BookEntry{
		{Move: Move(0x3df), Weight: 2},
	},
	0x8e78072a0ff68e32: []BookEntry{
		{Move: Move(0xae2), Weight: 1},
	},
	0x906785997599061b: []BookEntry{
		{Move: Move(0x218), Weight: 65520},
		{Move: Move(0x210), Weight: 29436},
	},
	0x92c79dda1cb3cf17: []BookEntry{
		{Move: Move(0xdef), Weight: 3},
		{Move: Move(0xf7c), Weight: 2},
		{Move: Move(0xe73), Weight: 2},
		{Move: Move(0xc61), Weight: 1},
		{Move: Move(0xc20), Weight: 1},
	},
	0x9f81c6f8c36c5f26: []BookEntry{
		{Move: Move(0x252), Weight: 4},
		{Move: Move(0x2d2), Weight: 4},
	},
	0x365ffe84562afc7e: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0xaa0d0ff55f1b402f: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xd928f6dfa4f42fe7: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x2ae9f67c1d93d54e: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x561c3e589ff6552c: []BookEntry{
		{Move: Move(0x161), Weight: 1},
	},
	0xf8fa2cdb06a78340: []BookEntry{
		{Move: Move(0x2d5), Weight: 1},
	},
	0x62ec4f9e5411ca11: []BookEntry{
		{Move: Move(0x8dc), Weight: 3},
	},
	0xeb8f18aac8f38868: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0xc08cf2a3a6e3498: []BookEntry{
		{Move: Move(0xfb4), Weight: 1},
	},
	0xb2342458e69e5d7c: []BookEntry{
		{Move: Move(0xee0), Weight: 1},
	},
	0x22c971fb283a038: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x39d459ba6d323392: []BookEntry{
		{Move: Move(0xf7c), Weight: 1},
	},
	0x5d29a0391419dfcc: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x5db4f32662eda7a5: []BookEntry{
		{Move: Move(0xe73), Weight: 11},
		{Move: Move(0xca2), Weight: 7},
	},
	0x698f29a8eb882f40: []BookEntry{
		{Move: Move(0x161), Weight: 65520},
	},
	0x7d553ff675f88fbf: []BookEntry{
		{Move: Move(0xaa4), Weight: 1},
	},
	0x1ae71d5904e6f28a: []BookEntry{
		{Move: Move(0xce3), Weight: 2},
	},
	0x51f18b0833614a01: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0xb1470b8083a93240: []BookEntry{
		{Move: Move(0xd2c), Weight: 3},
	},
	0x69281490e154d43a: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0xc2ba2e278b0db9c7: []BookEntry{
		{Move: Move(0x210), Weight: 65520},
		{Move: Move(0x2d3), Weight: 7280},
	},
	0xcfbe0a954afb4721: []BookEntry{
		{Move: Move(0x15a), Weight: 1},
	},
	0xd462b557d024f0eb: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0xef4305d2ec429650: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0x13ea16005fe8d54: []BookEntry{
		{Move: Move(0xdae), Weight: 1},
	},
	0x190fcb8f550a770c: []BookEntry{
		{Move: Move(0x91b), Weight: 14},
	},
	0x271aaf451eb51d56: []BookEntry{
		{Move: Move(0x396), Weight: 14},
		{Move: Move(0x314), Weight: 1},
	},
	0x44ed323e295ef431: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x65ae67db5c8699e8: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x6c3c28696d47d29f: []BookEntry{
		{Move: Move(0x622), Weight: 1},
	},
	0x6d1549688e26c176: []BookEntry{
		{Move: Move(0x15a), Weight: 2},
	},
	0x889541213ec615f2: []BookEntry{
		{Move: Move(0x195), Weight: 6},
	},
	0x35f358c5cce8b47f: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
		{Move: Move(0x8106), Weight: 8190},
		{Move: Move(0x218), Weight: 8190},
	},
	0x7d9b42b94a659936: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0x9e685da7bdbc3a39: []BookEntry{
		{Move: Move(0x4b), Weight: 5},
	},
	0xa951e4af5a4ec559: []BookEntry{
		{Move: Move(0xceb), Weight: 4},
	},
	0xbe19ae1c206d6938: []BookEntry{
		{Move: Move(0x8b), Weight: 1},
	},
	0xc2758e737bc6dc5e: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0xf6c25d92f49abf08: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0x1453ca40e62bb5ad: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x21fdd22ade8d34d5: []BookEntry{
		{Move: Move(0xca2), Weight: 1},
	},
	0x3a5be15ad3c1639c: []BookEntry{
		{Move: Move(0x49b), Weight: 1},
	},
	0x5dde121e812506a5: []BookEntry{
		{Move: Move(0xc61), Weight: 2},
	},
	0x86c80dc277be7fcf: []BookEntry{
		{Move: Move(0xeb1), Weight: 2},
	},
	0xa3c3f8d90fff5435: []BookEntry{
		{Move: Move(0xc69), Weight: 1},
	},
	0xb5e2a8d6e0d5e682: []BookEntry{
		{Move: Move(0xef3), Weight: 1},
	},
	0xb9cb8071e8297c10: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x29f436babdb3ed37: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x5375c1a61eafff00: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0x7a1d1fe6073cb8a5: []BookEntry{
		{Move: Move(0x7ac), Weight: 1},
	},
	0xa691c9860f2e59d5: []BookEntry{
		{Move: Move(0xca2), Weight: 1},
	},
	0xa7794faa35da9383: []BookEntry{
		{Move: Move(0x52), Weight: 4},
		{Move: Move(0x314), Weight: 2},
		{Move: Move(0x6ea), Weight: 1},
	},
	0xe3f3f35b4affc4df: []BookEntry{
		{Move: Move(0xceb), Weight: 1},
		{Move: Move(0xf74), Weight: 1},
		{Move: Move(0xce3), Weight: 1},
	},
	0xf21a6d7e626fb0b0: []BookEntry{
		{Move: Move(0x51b), Weight: 4},
	},
	0x29a11925b6cc0ae7: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0x1730369b8a817668: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
	},
	0x86653a6d93027086: []BookEntry{
		{Move: Move(0xe9e), Weight: 1},
	},
	0x99629d6da00c8482: []BookEntry{
		{Move: Move(0x55b), Weight: 10},
	},
	0xbd4fd3445cc56942: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x3f4a46e741dfa503: []BookEntry{
		{Move: Move(0x314), Weight: 3},
		{Move: Move(0x396), Weight: 1},
	},
	0xa506f1abcc0e593f: []BookEntry{
		{Move: Move(0x6a3), Weight: 65520},
	},
	0xa77dd7230d5726ea: []BookEntry{
		{Move: Move(0xde7), Weight: 3},
		{Move: Move(0xfad), Weight: 3},
	},
	0xa86142e8bb6dc40b: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0x95fa639eb4005378: []BookEntry{
		{Move: Move(0xca2), Weight: 1},
		{Move: Move(0xf3f), Weight: 1},
		{Move: Move(0xe73), Weight: 1},
	},
	0xfcaecef8e36dc2f7: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0x84dac15038e2d1eb: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
	},
	0xb6f05b48621d9b21: []BookEntry{
		{Move: Move(0xd5), Weight: 1},
	},
	0xe768c92907230645: []BookEntry{
		{Move: Move(0x55b), Weight: 2},
	},
	0x13592d293859007d: []BookEntry{
		{Move: Move(0x195), Weight: 19},
	},
	0x34d3e3bb02740bbf: []BookEntry{
		{Move: Move(0x292), Weight: 19},
		{Move: Move(0x29a), Weight: 1},
	},
	0x23bccc997422a0f4: []BookEntry{
		{Move: Move(0x913), Weight: 1},
	},
	0x2f84f03c4f929111: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x3d875f5e412c15f0: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0x8b7c0f62add0992d: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0xaf53f4dd841f860d: []BookEntry{
		{Move: Move(0x6e2), Weight: 2},
	},
	0xc20f63283dce1688: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0xe26c45486a16909c: []BookEntry{
		{Move: Move(0x210), Weight: 1},
		{Move: Move(0x89), Weight: 1},
	},
	0xe845f4f9a2f5eb42: []BookEntry{
		{Move: Move(0xef2), Weight: 1},
	},
	0x15dee38a2516b183: []BookEntry{
		{Move: Move(0xd2c), Weight: 65520},
	},
	0x7cf9971884332f44: []BookEntry{
		{Move: Move(0xeb3), Weight: 1},
		{Move: Move(0xf6b), Weight: 1},
	},
	0xca7d7c9c897102b4: []BookEntry{
		{Move: Move(0xfad), Weight: 65520},
	},
	0xe3aaf1cf6da1659d: []BookEntry{
		{Move: Move(0x314), Weight: 35280},
		{Move: Move(0x195), Weight: 65520},
	},
	0x41fac7ad97c147b9: []BookEntry{
		{Move: Move(0x795), Weight: 65520},
		{Move: Move(0x7a7), Weight: 21840},
	},
	0xf372983dd8e7550: []BookEntry{
		{Move: Move(0x72d), Weight: 7},
	},
	0x38fa4493c8cf093a: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0x4a64501df6f8595c: []BookEntry{
		{Move: Move(0xcc), Weight: 1},
	},
	0x4cd7d4c90dd627c2: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x247e90dadca40950: []BookEntry{
		{Move: Move(0xca), Weight: 2},
		{Move: Move(0x4b), Weight: 1},
	},
	0x4be53c41339895eb: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x7d6e48dbdc31a835: []BookEntry{
		{Move: Move(0x481), Weight: 1},
	},
	0xe228d944796bdefb: []BookEntry{
		{Move: Move(0x997), Weight: 1},
	},
	0xf35dabd873d40d70: []BookEntry{
		{Move: Move(0x94), Weight: 1},
	},
	0x420f9c67a4d1577: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x48ce30cc5a06f4fc: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x5f930003a5e8a61f: []BookEntry{
		{Move: Move(0x756), Weight: 1},
	},
	0x66ae581464a66b5f: []BookEntry{
		{Move: Move(0x8ed), Weight: 1},
	},
	0x9eca0b7de77fe68f: []BookEntry{
		{Move: Move(0x8b), Weight: 1},
	},
	0xebd16e1d483310f4: []BookEntry{
		{Move: Move(0x2db), Weight: 10},
	},
	0x27c673c440bea48b: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x11579afdc1eddda9: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x18596b41a5071374: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x630716c8153188c8: []BookEntry{
		{Move: Move(0x92a), Weight: 1},
	},
	0x87260d7f79c8bf1e: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x8edffe3dd59810bb: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0xb1684397aac32cc0: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xc6b9ff8c9e6de8f4: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
		{Move: Move(0xe6a), Weight: 1},
	},
	0x9a780aa21a5028b: []BookEntry{
		{Move: Move(0x67d), Weight: 1},
	},
	0x250b92c24699ecc3: []BookEntry{
		{Move: Move(0x6a3), Weight: 2},
	},
	0x5e62b4ba99b122e9: []BookEntry{
		{Move: Move(0x723), Weight: 1},
	},
	0x69d62f3b7e0c510f: []BookEntry{
		{Move: Move(0x6a3), Weight: 2},
	},
	0x796be06b136da7cd: []BookEntry{
		{Move: Move(0xdef), Weight: 4},
	},
	0xdab4dd232dc62747: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0xdd3402cf55425288: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0x16f831cdf8d5705d: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
	},
	0x2c8f4d599a963341: []BookEntry{
		{Move: Move(0x498), Weight: 1},
	},
	0x5186e5497a300820: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
		{Move: Move(0xd1), Weight: 1},
	},
	0xb112b1e78132d2ec: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
		{Move: Move(0x8106), Weight: 1},
	},
	0xd7ad25c39cc02453: []BookEntry{
		{Move: Move(0xb5c), Weight: 1},
	},
	0xe31d95cdacb378b1: []BookEntry{
		{Move: Move(0xe6a), Weight: 65520},
	},
	0xca2767bcbcf4af5: []BookEntry{
		{Move: Move(0x8b), Weight: 1},
	},
	0x5b1da96d96a561b1: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x70920eda86578186: []BookEntry{
		{Move: Move(0x688), Weight: 2},
		{Move: Move(0x8106), Weight: 2},
	},
	0x99f783d54b562c83: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x146d4888073671ef: []BookEntry{
		{Move: Move(0x91c), Weight: 8},
	},
	0x3699ac18ba28f2a2: []BookEntry{
		{Move: Move(0x292), Weight: 4},
	},
	0x5b8f0c302911599a: []BookEntry{
		{Move: Move(0xe6a), Weight: 3},
		{Move: Move(0xef4), Weight: 1},
		{Move: Move(0x91b), Weight: 1},
	},
	0xaba14136f7f505f6: []BookEntry{
		{Move: Move(0x556), Weight: 1},
	},
	0xbdbfa931f89e699d: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0xe15cca64f6a6b254: []BookEntry{
		{Move: Move(0xa58), Weight: 65520},
		{Move: Move(0xf38), Weight: 65520},
		{Move: Move(0xf74), Weight: 65520},
	},
	0x483c043ac217a732: []BookEntry{
		{Move: Move(0xce2), Weight: 1},
	},
	0x6e5ac0f496e69046: []BookEntry{
		{Move: Move(0x259), Weight: 1},
	},
	0xa77d7c9d438f663d: []BookEntry{
		{Move: Move(0x795), Weight: 1},
	},
	0xc15419e9afc7a10d: []BookEntry{
		{Move: Move(0x85a), Weight: 2},
	},
	0x3c4b4c6daf3506a8: []BookEntry{
		{Move: Move(0x8106), Weight: 33264},
		{Move: Move(0x94), Weight: 65520},
	},
	0x40ad304438b08969: []BookEntry{
		{Move: Move(0xa62), Weight: 2},
	},
	0x549c089dcde20874: []BookEntry{
		{Move: Move(0x55b), Weight: 3},
	},
	0xd21c3f3d8e398d1d: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0xfac7dbd62c6e3dac: []BookEntry{
		{Move: Move(0xc6a), Weight: 2},
	},
	0x9d1789ce6c349039: []BookEntry{
		{Move: Move(0xceb), Weight: 65520},
		{Move: Move(0xdae), Weight: 65520},
		{Move: Move(0xce3), Weight: 65520},
	},
	0xabd30e5cbf153469: []BookEntry{
		{Move: Move(0xc28), Weight: 4},
	},
	0x252c24fcbb48226a: []BookEntry{
		{Move: Move(0x8ab), Weight: 2},
	},
	0x2a49f39949ae7035: []BookEntry{
		{Move: Move(0xdef), Weight: 65520},
		{Move: Move(0xf3f), Weight: 16380},
	},
	0x83d576304109069f: []BookEntry{
		{Move: Move(0xce3), Weight: 17},
		{Move: Move(0xd24), Weight: 1},
	},
	0x99b2049c48f6d2ab: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
	},
	0xbdd8c21738a00496: []BookEntry{
		{Move: Move(0xcaa), Weight: 4},
	},
	0xd89390d6ed2b87c9: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0xf553a8d3106483de: []BookEntry{
		{Move: Move(0x103), Weight: 2},
	},
	0x2c69d892aad95ce6: []BookEntry{
		{Move: Move(0xfad), Weight: 8},
	},
	0x8760ece966df7fd5: []BookEntry{
		{Move: Move(0xdae), Weight: 3},
	},
	0xf6f1ebdfadd0a3c3: []BookEntry{
		{Move: Move(0x4a1), Weight: 1},
	},
	0x22b405c4f5182a55: []BookEntry{
		{Move: Move(0x35d), Weight: 1},
	},
	0x638a6951fa71f7a1: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0x6b29f0fb1cbabe46: []BookEntry{
		{Move: Move(0x8e9), Weight: 3},
		{Move: Move(0x8f4), Weight: 2},
	},
	0x850c78a78d5f9a19: []BookEntry{
		{Move: Move(0x14e), Weight: 2},
	},
	0xba2e5b24be7341b6: []BookEntry{
		{Move: Move(0xae3), Weight: 1},
	},
	0x8e8993f8d29201f: []BookEntry{
		{Move: Move(0x396), Weight: 3},
	},
	0x183558fae2a3d387: []BookEntry{
		{Move: Move(0xce3), Weight: 30},
		{Move: Move(0xdae), Weight: 20},
		{Move: Move(0xca2), Weight: 9},
	},
	0x393e742ab5f7dfe9: []BookEntry{
		{Move: Move(0x4ca), Weight: 1},
	},
	0xbf79a84f3c2c10cf: []BookEntry{
		{Move: Move(0xd3), Weight: 6},
	},
	0xdd7afb3777c6e6b1: []BookEntry{
		{Move: Move(0xeb3), Weight: 1},
	},
	0x3cbc259d1d24992d: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x503a3e3add0480ea: []BookEntry{
		{Move: Move(0x8dc), Weight: 1},
	},
	0x9dc3316ed9141a08: []BookEntry{
		{Move: Move(0x8106), Weight: 6},
	},
	0xab2dcfdba75868fc: []BookEntry{
		{Move: Move(0x4b), Weight: 3},
	},
	0xd775ee0044cc79fe: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0xd78e9152e02a8430: []BookEntry{
		{Move: Move(0xf74), Weight: 6},
	},
	0x11f39d813e9dbda3: []BookEntry{
		{Move: Move(0x3d7), Weight: 2},
		{Move: Move(0x94), Weight: 1},
	},
	0x2201f587fe2164e0: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0x3c195fa8e72a0aae: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0x87990398782ab867: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x9860a1d20a8655e0: []BookEntry{
		{Move: Move(0xeb3), Weight: 1},
	},
	0xd99dd453f7306f56: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
		{Move: Move(0x50), Weight: 1},
	},
	0xf3a51442a2ff7859: []BookEntry{
		{Move: Move(0xea5), Weight: 65520},
	},
	0x1efdc27e1a4ec322: []BookEntry{
		{Move: Move(0x51d), Weight: 1},
	},
	0x64dc3b9ab0d28634: []BookEntry{
		{Move: Move(0xc6a), Weight: 2},
	},
	0x320b6dbaad2b3fe: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0x18ee9778a4f204c6: []BookEntry{
		{Move: Move(0x14e), Weight: 1},
	},
	0x25a1bc1cc8d7809f: []BookEntry{
		{Move: Move(0xb24), Weight: 1},
	},
	0x4e2fe6bc0d2665bf: []BookEntry{
		{Move: Move(0xe6a), Weight: 3},
	},
	0x7295c1b6d591dd2f: []BookEntry{
		{Move: Move(0x8a9), Weight: 1},
	},
	0x848f617c1dc0b54a: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x9518ee6e56e3abb6: []BookEntry{
		{Move: Move(0xf7c), Weight: 4},
	},
	0xb52ef7aadf87162b: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0x67521168b0ded8bb: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0x6b919cc41f6ae26c: []BookEntry{
		{Move: Move(0xd65), Weight: 1},
		{Move: Move(0xf3f), Weight: 1},
		{Move: Move(0xc28), Weight: 1},
	},
	0x732357f7a4c0612: []BookEntry{
		{Move: Move(0x396), Weight: 1},
	},
	0x6f36d0f5546618b1: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0x830eb9b20758d1de: []BookEntry{
		{Move: Move(0x195), Weight: 65519},
		{Move: Move(0x2db), Weight: 53607},
	},
	0xb87a994f94c05dc0: []BookEntry{
		{Move: Move(0x4b), Weight: 5},
	},
	0xdc7f0ad5a665f739: []BookEntry{
		{Move: Move(0xca2), Weight: 65520},
		{Move: Move(0xf76), Weight: 43680},
	},
	0xe657841fc95e7099: []BookEntry{
		{Move: Move(0x89b), Weight: 3},
	},
	0xf9d00ca49969ca20: []BookEntry{
		{Move: Move(0x2db), Weight: 65520},
		{Move: Move(0x251), Weight: 23400},
		{Move: Move(0x292), Weight: 4680},
	},
	0x6307e59d9ade6368: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0xa2dcaf75395e465e: []BookEntry{
		{Move: Move(0x953), Weight: 1},
	},
	0xa697b0027119dda9: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x2eda985b1f08a7de: []BookEntry{
		{Move: Move(0x3df), Weight: 3},
	},
	0x304e806629817f6f: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
	},
	0x31a4fd19dcf390f1: []BookEntry{
		{Move: Move(0xa32), Weight: 1},
	},
	0x4ca125838500cdb8: []BookEntry{
		{Move: Move(0x564), Weight: 1},
	},
	0xa2a1e43de0c38ccc: []BookEntry{
		{Move: Move(0xc61), Weight: 1},
	},
	0x1a83398fccdf7729: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x1ac14344abf65c20: []BookEntry{
		{Move: Move(0x795), Weight: 3},
	},
	0x8cb307d294892e68: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0x98b7316e6a5fc7ef: []BookEntry{
		{Move: Move(0xdae), Weight: 53607},
		{Move: Move(0xd2c), Weight: 65519},
	},
	0x2411c1fff69893a4: []BookEntry{
		{Move: Move(0xc28), Weight: 2},
	},
	0xd2d63f2da32ff410: []BookEntry{
		{Move: Move(0x49c), Weight: 4},
	},
	0xf643dadaf65345ce: []BookEntry{
		{Move: Move(0xef2), Weight: 2},
	},
	0x1c86e3d035dad294: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0x1f9374044fba81d5: []BookEntry{
		{Move: Move(0x31a), Weight: 65520},
	},
	0x3f1443c418881f39: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0x41ae97c7b9f29a70: []BookEntry{
		{Move: Move(0x9df), Weight: 2},
		{Move: Move(0xf6b), Weight: 1},
	},
	0x8c535dd3aee45e7e: []BookEntry{
		{Move: Move(0xd65), Weight: 8},
	},
	0xa26befa14dfa5a12: []BookEntry{
		{Move: Move(0xc28), Weight: 5},
	},
	0xa9bcf09f16550990: []BookEntry{
		{Move: Move(0x218), Weight: 1},
	},
	0xb1adb360d339e7c0: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0x8624c07465e5e41a: []BookEntry{
		{Move: Move(0x218), Weight: 65520},
		{Move: Move(0x691), Weight: 65520},
	},
	0x9b1c43e2ad768db2: []BookEntry{
		{Move: Move(0xb5e), Weight: 1},
	},
	0xc93c61f8996a6e37: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0xdb55d4fcaadc775e: []BookEntry{
		{Move: Move(0x210), Weight: 65520},
		{Move: Move(0x52), Weight: 1},
	},
	0xe06b60ca762dfae: []BookEntry{
		{Move: Move(0xfb4), Weight: 1},
	},
	0x118cd876319c1e7b: []BookEntry{
		{Move: Move(0x210), Weight: 2},
	},
	0x24b4c7f5448340ec: []BookEntry{
		{Move: Move(0x259), Weight: 1},
		{Move: Move(0x51c), Weight: 1},
		{Move: Move(0x210), Weight: 1},
	},
	0x27e3efa1be503b3f: []BookEntry{
		{Move: Move(0xf59), Weight: 1},
	},
	0x6599a1efe54c61a7: []BookEntry{
		{Move: Move(0xca2), Weight: 1},
	},
	0x6780bd452cc80cfc: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0x7a1528f73b6d8505: []BookEntry{
		{Move: Move(0xb63), Weight: 2},
	},
	0x991dfcf7ae9d941c: []BookEntry{
		{Move: Move(0xeb3), Weight: 5},
	},
	0x235cdb0d68ea04dc: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x5e96d510a9a3727e: []BookEntry{
		{Move: Move(0xb63), Weight: 65520},
	},
	0x8eefeb75171cca51: []BookEntry{
		{Move: Move(0xb63), Weight: 5},
	},
	0xa84f8f7f8274fbcd: []BookEntry{
		{Move: Move(0x49a), Weight: 1},
	},
	0xbb466ea803f76c79: []BookEntry{
		{Move: Move(0xcaa), Weight: 2},
	},
	0xb2790472b437dea0: []BookEntry{
		{Move: Move(0xfad), Weight: 65520},
		{Move: Move(0xde7), Weight: 7708},
		{Move: Move(0xcaa), Weight: 3854},
	},
	0xb844c53ac05b240b: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0xc95e54397292bb84: []BookEntry{
		{Move: Move(0xe6a), Weight: 4},
		{Move: Move(0xce3), Weight: 1},
	},
	0x505c95f1e56a5c3e: []BookEntry{
		{Move: Move(0x161), Weight: 5},
	},
	0xa9dce7a7896ed6ce: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
		{Move: Move(0x314), Weight: 1},
	},
	0xe885db57abd08ab9: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0xeac0c2f894f40b87: []BookEntry{
		{Move: Move(0xd1), Weight: 1},
	},
	0xf6d220cf97caf250: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0x339a1752cf7a6bff: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xa95610520aedc209: []BookEntry{
		{Move: Move(0xca2), Weight: 5},
		{Move: Move(0xdef), Weight: 4},
		{Move: Move(0xe73), Weight: 1},
		{Move: Move(0xefc), Weight: 1},
	},
	0xaf0c78718ccbfda0: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
	},
	0xb98f8008e554c418: []BookEntry{
		{Move: Move(0xd65), Weight: 65519},
		{Move: Move(0xeed), Weight: 20690},
	},
	0xab85da604d68cce3: []BookEntry{
		{Move: Move(0x31c), Weight: 1},
	},
	0xd2c8263db45ce25b: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0xef6c09dfddb1d023: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
		{Move: Move(0x8106), Weight: 1},
	},
	0x429bd3b520b60b3f: []BookEntry{
		{Move: Move(0x314), Weight: 2},
	},
	0x4f4dad2b78333dd5: []BookEntry{
		{Move: Move(0xa9b), Weight: 2},
	},
	0x723cb59062632602: []BookEntry{
		{Move: Move(0x8106), Weight: 14},
		{Move: Move(0x218), Weight: 9},
		{Move: Move(0x691), Weight: 9},
	},
	0x8301b594404f73be: []BookEntry{
		{Move: Move(0xfad), Weight: 6},
	},
	0xa78bb069ba8f9d8f: []BookEntry{
		{Move: Move(0xd8), Weight: 1},
	},
	0x1f59853e4e6cf3c0: []BookEntry{
		{Move: Move(0xd2d), Weight: 2},
	},
	0x4bbc48730f565aba: []BookEntry{
		{Move: Move(0xd1), Weight: 2},
		{Move: Move(0x31c), Weight: 2},
	},
	0x4db1e8f97ed165dc: []BookEntry{
		{Move: Move(0x153), Weight: 65520},
		{Move: Move(0x251), Weight: 32760},
		{Move: Move(0x564), Weight: 32760},
	},
	0x604f3f7af659ba29: []BookEntry{
		{Move: Move(0x86a), Weight: 4},
	},
	0xaaa48a910ae07f23: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0xb21351130ba09819: []BookEntry{
		{Move: Move(0xe73), Weight: 4},
		{Move: Move(0xf6b), Weight: 1},
		{Move: Move(0x96e), Weight: 1},
		{Move: Move(0xf59), Weight: 1},
		{Move: Move(0x8da), Weight: 1},
	},
	0xbb7bec8dea70b21c: []BookEntry{
		{Move: Move(0xb23), Weight: 2},
	},
	0xde9443b04de8727f: []BookEntry{
		{Move: Move(0x6e2), Weight: 2},
	},
	0x239aa8ef5dcf0d00: []BookEntry{
		{Move: Move(0x8da), Weight: 1},
	},
	0x358ea87aee782648: []BookEntry{
		{Move: Move(0x8e9), Weight: 1},
	},
	0x3edf319b7d774c78: []BookEntry{
		{Move: Move(0x4a4), Weight: 1},
	},
	0x578b79b22261789c: []BookEntry{
		{Move: Move(0x913), Weight: 1},
	},
	0x64c239bf565ae721: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x87669a18199ed752: []BookEntry{
		{Move: Move(0x161), Weight: 1},
	},
	0xa27873ec2bd96aa5: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0xaa7250561aaf50e8: []BookEntry{
		{Move: Move(0xeb1), Weight: 1},
	},
	0x209ede6caa86f2: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x567e57d109ee9e44: []BookEntry{
		{Move: Move(0x252), Weight: 7},
	},
	0xe613da09d133fa3f: []BookEntry{
		{Move: Move(0xa6), Weight: 11},
		{Move: Move(0x9d), Weight: 7},
		{Move: Move(0x4b), Weight: 1},
	},
	0x371436ea3e73e228: []BookEntry{
		{Move: Move(0x259), Weight: 4},
	},
	0x4baae1b28ca0d2e4: []BookEntry{
		{Move: Move(0x45a), Weight: 1},
	},
	0x7ffd0523148e314e: []BookEntry{
		{Move: Move(0x481), Weight: 1},
		{Move: Move(0x2d3), Weight: 65520},
	},
	0xb8a62bb07e86ad02: []BookEntry{
		{Move: Move(0x8db), Weight: 5},
	},
	0xbb86a4c24916da46: []BookEntry{
		{Move: Move(0x52), Weight: 2},
	},
	0xd8c5fdd0fc01387f: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0xe16be6294ccae2a5: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0xf4da85b96960abbc: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0x25688ef616583c60: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x42bd18dc743a0025: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x5492cb9e7422885a: []BookEntry{
		{Move: Move(0x210), Weight: 65520},
		{Move: Move(0x218), Weight: 65520},
	},
	0xb663577df5f241bb: []BookEntry{
		{Move: Move(0xcaa), Weight: 29},
		{Move: Move(0xf3f), Weight: 8},
		{Move: Move(0xceb), Weight: 5},
	},
	0xff2e4da42ce3fa92: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0x25a4d3d8ada4601c: []BookEntry{
		{Move: Move(0x14e), Weight: 4},
	},
	0x4c501d8620d409b7: []BookEntry{
		{Move: Move(0x52), Weight: 8},
	},
	0xa5afd0c7f356260c: []BookEntry{
		{Move: Move(0xf6b), Weight: 2},
	},
	0xbd605731ea9a9886: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x36d8bb934b7ba7: []BookEntry{
		{Move: Move(0xea5), Weight: 1},
	},
	0x10337d5926e51238: []BookEntry{
		{Move: Move(0x91b), Weight: 1},
	},
	0x237816d0c9e173f0: []BookEntry{
		{Move: Move(0xca2), Weight: 2},
	},
	0x346b0b567f083b3f: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x387ec3b5e17c8e8c: []BookEntry{
		{Move: Move(0xf76), Weight: 6},
	},
	0x442e9a6eaf32b88c: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x6fb8b81580d4b735: []BookEntry{
		{Move: Move(0xe6a), Weight: 35280},
		{Move: Move(0xce3), Weight: 65520},
	},
	0x915ab05d5efddbb6: []BookEntry{
		{Move: Move(0x7ac), Weight: 1},
	},
	0xc1b7a26abff9aff6: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x8fdd2ec5f0dfc251: []BookEntry{
		{Move: Move(0xc28), Weight: 65520},
	},
	0x9119a3394a68480b: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0xa17be3eb6a4cdd82: []BookEntry{
		{Move: Move(0x8106), Weight: 4},
		{Move: Move(0x2db), Weight: 2},
	},
	0xbe7d760738e7f0e5: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
		{Move: Move(0x8106), Weight: 1},
	},
	0xcb067d00107c80db: []BookEntry{
		{Move: Move(0x251), Weight: 5},
		{Move: Move(0x2db), Weight: 5},
	},
	0xd3801b6871adc60c: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0xe151f8b18f5f4670: []BookEntry{
		{Move: Move(0x31c), Weight: 65520},
		{Move: Move(0x314), Weight: 16380},
	},
	0xfc47842a785ce02d: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x5f9c97d001dde5b8: []BookEntry{
		{Move: Move(0x8b), Weight: 65520},
	},
	0xc4ec99d8fc4d1225: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0xd923f8f0336d29c4: []BookEntry{
		{Move: Move(0x161), Weight: 19},
		{Move: Move(0x2db), Weight: 3},
	},
	0x4c4595724944c61d: []BookEntry{
		{Move: Move(0x48c), Weight: 1},
	},
	0x7c2416fb2671e25c: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0x8ecab9c13f81a410: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0x2116b700e8400dc1: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
		{Move: Move(0x153), Weight: 1},
	},
	0x35d84ca7f1375c3a: []BookEntry{
		{Move: Move(0xb63), Weight: 65520},
	},
	0x47d27b78aa89bb95: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x6d6e2d96fac71f78: []BookEntry{
		{Move: Move(0x6e2), Weight: 1},
	},
	0x9cb86a6d8d4a2ea0: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0xb795953d0d3fa5ba: []BookEntry{
		{Move: Move(0x6ea), Weight: 1},
	},
	0x15e4e7f06399397c: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x2cd203225d185792: []BookEntry{
		{Move: Move(0xdef), Weight: 2},
		{Move: Move(0xf7c), Weight: 1},
		{Move: Move(0x8a9), Weight: 1},
	},
	0x43278c9993832518: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
	},
	0x8a1ce68e2a1b7084: []BookEntry{
		{Move: Move(0xe73), Weight: 1},
	},
	0xeb640ad87c03bccd: []BookEntry{
		{Move: Move(0xba5), Weight: 1},
	},
	0x1291aba34e3c6012: []BookEntry{
		{Move: Move(0x259), Weight: 1},
	},
	0x19d9fcbaf70ff87f: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x34e19cfb0f9e6d2e: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x9ed8fd8db9a45d4c: []BookEntry{
		{Move: Move(0xf6b), Weight: 65520},
		{Move: Move(0xc69), Weight: 8736},
		{Move: Move(0xf74), Weight: 13104},
	},
	0xa08d277f6e263551: []BookEntry{
		{Move: Move(0xd5), Weight: 2},
	},
	0xd383846276341f5c: []BookEntry{
		{Move: Move(0x89), Weight: 65520},
	},
	0x34d38e62d9b9cb4e: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0x36e3a1da7a6cef43: []BookEntry{
		{Move: Move(0xab4), Weight: 1},
	},
	0x64a8267580fbd4bf: []BookEntry{
		{Move: Move(0x252), Weight: 1},
	},
	0x65e122941c18ab29: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x8415831f70558477: []BookEntry{
		{Move: Move(0x86a), Weight: 1},
	},
	0x8824072becafdb79: []BookEntry{
		{Move: Move(0x18f), Weight: 1},
	},
	0x94cf94a9c18ccef2: []BookEntry{
		{Move: Move(0x14c), Weight: 65520},
	},
	0xdcbc22c0c809b86c: []BookEntry{
		{Move: Move(0x52), Weight: 2},
	},
	0x6212f055d3c9a7cd: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0x7691b9e84cb6c3c3: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0xdbfea0d031d66fa6: []BookEntry{
		{Move: Move(0x86a), Weight: 2},
	},
	0xe7bdc5643145acb4: []BookEntry{
		{Move: Move(0xeac), Weight: 8},
	},
	0xac65eb80ea89b2f: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0xd04ac04ed275f1c: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
	},
	0x3902884db892e050: []BookEntry{
		{Move: Move(0x15a), Weight: 65520},
	},
	0xae610c4683a849dc: []BookEntry{
		{Move: Move(0x91c), Weight: 65520},
	},
	0xd0aaf92e9618e03c: []BookEntry{
		{Move: Move(0xf62), Weight: 2},
	},
	0xd58df98987438fbc: []BookEntry{
		{Move: Move(0x8dc), Weight: 2},
	},
	0x11c6bc50682e1f0d: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0x28e09df3343587e0: []BookEntry{
		{Move: Move(0x91c), Weight: 1},
	},
	0x3774a361dde63617: []BookEntry{
		{Move: Move(0x52), Weight: 3},
	},
	0x43ce039f28a435cc: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0xc421de9c9fbb9e60: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
		{Move: Move(0x48c), Weight: 1},
	},
	0xd3725eecb0e6d8b4: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0xfb3f874522d5a7a4: []BookEntry{
		{Move: Move(0x195), Weight: 2},
	},
	0x1a7fbad2d6b0043a: []BookEntry{
		{Move: Move(0x564), Weight: 1},
	},
	0x40e808efd5d41069: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x778b5d0fe3ac5b00: []BookEntry{
		{Move: Move(0x94), Weight: 1},
	},
	0xa4b171b3149d3d05: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xd4d1f1a57613a433: []BookEntry{
		{Move: Move(0x3d7), Weight: 2},
	},
	0x30f01e6cddcd8da: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0x76cd323b218c9f55: []BookEntry{
		{Move: Move(0xae4), Weight: 1},
	},
	0xccfe1bdcfca7c5f0: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0xcdcea23704ba7516: []BookEntry{
		{Move: Move(0x86a), Weight: 1},
	},
	0xeb159e897a63ed92: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x7a40a07a03f5920f: []BookEntry{
		{Move: Move(0xd24), Weight: 1},
	},
	0x823c9b50fd114196: []BookEntry{
		{Move: Move(0x31c), Weight: 65520},
		{Move: Move(0x29a), Weight: 43680},
		{Move: Move(0x292), Weight: 1},
		{Move: Move(0x314), Weight: 1},
	},
	0xae3cb923d4784906: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0x202ba95c19a29f0f: []BookEntry{
		{Move: Move(0xb5c), Weight: 5},
	},
	0x46e551346f8a2e9e: []BookEntry{
		{Move: Move(0xeac), Weight: 2},
	},
	0x4ebee1eb1389a342: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x8acfe2f02b08e478: []BookEntry{
		{Move: Move(0x99f), Weight: 1},
	},
	0x6cb643555e688e0b: []BookEntry{
		{Move: Move(0x18c), Weight: 65520},
	},
	0x74f2b7ceb03eeb1f: []BookEntry{
		{Move: Move(0x355), Weight: 1},
	},
	0xadeb8007aad9790e: []BookEntry{
		{Move: Move(0x9d), Weight: 2},
	},
	0xc3adbe732c95a860: []BookEntry{
		{Move: Move(0x55b), Weight: 1},
	},
	0xd6a52ba22c274313: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xf0d296a95f236c3a: []BookEntry{
		{Move: Move(0x8e9), Weight: 10},
	},
	0xfa60b1f968b311c3: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x1e09ea2f1b873d2c: []BookEntry{
		{Move: Move(0x31c), Weight: 1},
		{Move: Move(0x210), Weight: 1},
	},
	0x491006cb8f7638cc: []BookEntry{
		{Move: Move(0x4dd), Weight: 2},
	},
	0x686f001ee61959e4: []BookEntry{
		{Move: Move(0x2d3), Weight: 2},
	},
	0x6e0dfabeddda6f85: []BookEntry{
		{Move: Move(0x251), Weight: 8},
		{Move: Move(0x52), Weight: 1},
		{Move: Move(0x6a3), Weight: 1},
	},
	0xb8beb8ffde7a3920: []BookEntry{
		{Move: Move(0x2d3), Weight: 16380},
		{Move: Move(0x3d7), Weight: 65520},
	},
	0x25936aa438068fc2: []BookEntry{
		{Move: Move(0x51d), Weight: 1},
	},
	0x3d52801243a63dfe: []BookEntry{
		{Move: Move(0xe3a), Weight: 1},
		{Move: Move(0xf6b), Weight: 1},
		{Move: Move(0xef2), Weight: 1},
	},
	0x49998b0e302ba184: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x7e68ed1978f1f2f8: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0xc54dde8695b6adf6: []BookEntry{
		{Move: Move(0x218), Weight: 1},
	},
	0xdf66a63b320d0d37: []BookEntry{
		{Move: Move(0x31c), Weight: 1},
		{Move: Move(0x314), Weight: 1},
	},
	0xb4fbe5d3d852b04: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x24cf5eec5531dae3: []BookEntry{
		{Move: Move(0xadd), Weight: 1},
	},
	0x9a33a1bda927a7fb: []BookEntry{
		{Move: Move(0xb5c), Weight: 2},
		{Move: Move(0xe3a), Weight: 1},
	},
	0xc0952a2c36971b29: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0xf522fe1c8d8b8a0f: []BookEntry{
		{Move: Move(0xef4), Weight: 1},
	},
	0xab961e2c0ceff227: []BookEntry{
		{Move: Move(0xcea), Weight: 1},
	},
	0x50f4dd8b75e88638: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x5862e4d3225e8420: []BookEntry{
		{Move: Move(0xd24), Weight: 65520},
		{Move: Move(0xce3), Weight: 65520},
		{Move: Move(0xfad), Weight: 16380},
		{Move: Move(0xe6a), Weight: 16380},
	},
	0x77f1bf5e3dd7951: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x160a5893b400feca: []BookEntry{
		{Move: Move(0x913), Weight: 2},
	},
	0x284eced5b9682851: []BookEntry{
		{Move: Move(0xef3), Weight: 1},
	},
	0x6b509d2553e461d9: []BookEntry{
		{Move: Move(0x14e), Weight: 1},
	},
	0x76a9691c37494042: []BookEntry{
		{Move: Move(0xf59), Weight: 1},
	},
	0x4a1b790ac434244: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
		{Move: Move(0x35d), Weight: 1},
	},
	0x514316f4cbef4e66: []BookEntry{
		{Move: Move(0x6e4), Weight: 1},
	},
	0x8aba49d914518561: []BookEntry{
		{Move: Move(0x3d7), Weight: 2},
		{Move: Move(0x14c), Weight: 1},
	},
	0xae2b774eabe45281: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x268d3b99e61735d7: []BookEntry{
		{Move: Move(0x18c), Weight: 2},
	},
	0x493db4d659fd3551: []BookEntry{
		{Move: Move(0xce3), Weight: 2},
		{Move: Move(0xf3f), Weight: 1},
	},
	0x83658a9caa3c31fc: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
	},
	0xc6cb7919bf93dc5c: []BookEntry{
		{Move: Move(0xb63), Weight: 2},
	},
	0xe3da1a1d41674ead: []BookEntry{
		{Move: Move(0x4b), Weight: 2},
	},
	0xf309fde4ccbb2e7d: []BookEntry{
		{Move: Move(0xceb), Weight: 16380},
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x9f5d575ed4019d0e: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xb3e9ed04af74b8ef: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xfa4f04cf7a0a403e: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0x21440dd4152ab2cf: []BookEntry{
		{Move: Move(0xd3), Weight: 1},
	},
	0x70b80c1ebb4dd25d: []BookEntry{
		{Move: Move(0xd6d), Weight: 7280},
		{Move: Move(0xf62), Weight: 1},
		{Move: Move(0xe9e), Weight: 65520},
	},
	0x21ef49cc25357c34: []BookEntry{
		{Move: Move(0x91a), Weight: 3},
	},
	0x35af20718f1446b0: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0x6f6083ecd8231958: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x8575eed5ef914657: []BookEntry{
		{Move: Move(0xa6), Weight: 1},
	},
	0x8bb922ed54dc54e0: []BookEntry{
		{Move: Move(0x8db), Weight: 2},
	},
	0xa774f2abaf440610: []BookEntry{
		{Move: Move(0xee3), Weight: 1},
	},
	0x4da4b44f8100f95: []BookEntry{
		{Move: Move(0x29a), Weight: 10},
		{Move: Move(0x251), Weight: 1},
	},
	0x11d08f5f5ae679ac: []BookEntry{
		{Move: Move(0x8e9), Weight: 1},
	},
	0xd48d61d9bcb81b29: []BookEntry{
		{Move: Move(0xe9e), Weight: 1},
	},
	0xefbe506124d4cf05: []BookEntry{
		{Move: Move(0x39e), Weight: 1},
	},
	0x16c2a48572af530f: []BookEntry{
		{Move: Move(0xe3a), Weight: 3},
	},
	0x1aec48017b36216f: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0x5fffb7733927f8f7: []BookEntry{
		{Move: Move(0x9ad), Weight: 3},
	},
	0x62a0a6dbbfaf3870: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
	},
	0x97a12052d0e17670: []BookEntry{
		{Move: Move(0x7d6), Weight: 2},
	},
	0xe0414aa5d3d951c6: []BookEntry{
		{Move: Move(0xf62), Weight: 1},
	},
	0x4aa6f556a54a105e: []BookEntry{
		{Move: Move(0x195), Weight: 2},
	},
	0x423e3759d3c6ddb8: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0x72a956f52c98ac3c: []BookEntry{
		{Move: Move(0x691), Weight: 2},
	},
	0x883934ea89698ff1: []BookEntry{
		{Move: Move(0x35d), Weight: 1},
	},
	0xe838ae5b9e69b1b2: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x1d25fff6cf076f36: []BookEntry{
		{Move: Move(0x84c), Weight: 6},
	},
	0x297128683ee08ce9: []BookEntry{
		{Move: Move(0xc28), Weight: 3},
		{Move: Move(0xef2), Weight: 1},
	},
	0x3cc7ae0de7bc6e19: []BookEntry{
		{Move: Move(0xd3), Weight: 1},
	},
	0x5befd5a6f08ee608: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0x10fd4254dfedaf8b: []BookEntry{
		{Move: Move(0xceb), Weight: 43680},
		{Move: Move(0xf3f), Weight: 65520},
	},
	0xf64dd621d0625040: []BookEntry{
		{Move: Move(0x8db), Weight: 1},
	},
	0x4a2a7db8cbb256ec: []BookEntry{
		{Move: Move(0x8106), Weight: 5},
	},
	0x6bd674c3e3aa643a: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0xa6783103fd7bdb71: []BookEntry{
		{Move: Move(0x55b), Weight: 2},
	},
	0xa8c53ca58090a0b5: []BookEntry{
		{Move: Move(0x716), Weight: 1},
	},
	0x766e0107bb3b5a79: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0xdbefcaf48b921a5e: []BookEntry{
		{Move: Move(0x953), Weight: 3},
	},
	0xf3ba5484cdd8361f: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xff58c8d60dad788c: []BookEntry{
		{Move: Move(0x94), Weight: 3},
	},
	0x35a51d6a4353d380: []BookEntry{
		{Move: Move(0x96c), Weight: 1},
	},
	0x46428f7b575f4fe6: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
		{Move: Move(0xef3), Weight: 1},
		{Move: Move(0xab4), Weight: 1},
		{Move: Move(0xc28), Weight: 1},
	},
	0x57a6ff9e4312b089: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x9336b945c62e5c68: []BookEntry{
		{Move: Move(0xceb), Weight: 1},
	},
	0xc4830018ed3677fb: []BookEntry{
		{Move: Move(0x8a9), Weight: 1},
	},
	0xe88de696d25c4667: []BookEntry{
		{Move: Move(0x210), Weight: 1},
	},
	0xe97bb628062d587f: []BookEntry{
		{Move: Move(0x195), Weight: 5},
	},
	0xc11a67137da74e5: []BookEntry{
		{Move: Move(0xa9b), Weight: 1},
	},
	0x683184c2b6221383: []BookEntry{
		{Move: Move(0xef2), Weight: 1},
		{Move: Move(0xf6b), Weight: 1},
	},
	0x6ecf5687b1072223: []BookEntry{
		{Move: Move(0xeed), Weight: 3},
	},
	0x8e4e245025e9ec78: []BookEntry{
		{Move: Move(0xc28), Weight: 1},
	},
	0xa9ff7eefc9565129: []BookEntry{
		{Move: Move(0xd5), Weight: 1},
	},
	0xb6b0c66ae9c6364c: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xda0afe2641635934: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0x3392af4c170c6902: []BookEntry{
		{Move: Move(0x161), Weight: 3},
	},
	0x4f7a9abf84aa2d7e: []BookEntry{
		{Move: Move(0xca2), Weight: 1},
	},
	0x5f3ee1990a17ad83: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
	},
	0xa08b1c0126a1b84f: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0x1cbd18324a4ab0a8: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0x37e7363d8591f29e: []BookEntry{
		{Move: Move(0x49b), Weight: 1},
	},
	0x4a3cb856134f5e98: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
		{Move: Move(0x4b), Weight: 1},
	},
	0x6e38a4e010de9a0b: []BookEntry{
		{Move: Move(0x2d3), Weight: 65519},
		{Move: Move(0x8106), Weight: 40157},
	},
	0x811fa00ca33ff93b: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0x8b0e76c6c4c10c6d: []BookEntry{
		{Move: Move(0xb63), Weight: 65520},
	},
	0xa1d5e60c1b926f91: []BookEntry{
		{Move: Move(0xe9e), Weight: 2},
		{Move: Move(0xe73), Weight: 1},
	},
	0xd1f151406ae8a397: []BookEntry{
		{Move: Move(0x6a3), Weight: 5},
	},
	0x212a112686ccc9da: []BookEntry{
		{Move: Move(0x89), Weight: 1},
	},
	0x63024f7e68c7d95f: []BookEntry{
		{Move: Move(0xf62), Weight: 1},
	},
	0x8aeb317664e12c2f: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
		{Move: Move(0xef2), Weight: 1},
	},
	0xa730b46e12064fad: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0xc1c440399cb3a28c: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
		{Move: Move(0xceb), Weight: 1},
	},
	0xcd4f7f6de29e3ec5: []BookEntry{
		{Move: Move(0x29a), Weight: 1},
		{Move: Move(0x6ea), Weight: 1},
	},
	0xd4830b853bb58601: []BookEntry{
		{Move: Move(0xe39), Weight: 1},
		{Move: Move(0xeac), Weight: 1},
	},
	0xfb1e9666a5a09e9e: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
		{Move: Move(0xe68), Weight: 1},
	},
	0x2261affbb0ef4651: []BookEntry{
		{Move: Move(0xdae), Weight: 1},
		{Move: Move(0xef4), Weight: 1},
	},
	0x6e1e07142474a455: []BookEntry{
		{Move: Move(0x14c), Weight: 2},
	},
	0x7bb009a651418951: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x8a470482d88334ff: []BookEntry{
		{Move: Move(0x314), Weight: 65520},
		{Move: Move(0x292), Weight: 11562},
	},
	0xe434d52226759e3e: []BookEntry{
		{Move: Move(0x316), Weight: 6},
	},
	0x1cf0dc1189f1c4e1: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x64fa7731b5725eaf: []BookEntry{
		{Move: Move(0xc20), Weight: 1},
	},
	0xc359059bc7a58679: []BookEntry{
		{Move: Move(0x195), Weight: 5},
		{Move: Move(0x314), Weight: 1},
	},
	0x49e3fd0a4cb255ff: []BookEntry{
		{Move: Move(0xd2c), Weight: 3},
	},
	0xceaefc1e653a8f1d: []BookEntry{
		{Move: Move(0x652), Weight: 8},
		{Move: Move(0xf3f), Weight: 6},
	},
	0xdac3e56f9bd5e92d: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x20f0d6b04c90ffbf: []BookEntry{
		{Move: Move(0xef3), Weight: 3},
	},
	0x3f75684e71170b62: []BookEntry{
		{Move: Move(0x8d2), Weight: 65520},
		{Move: Move(0x8d9), Weight: 35280},
		{Move: Move(0xa99), Weight: 1},
	},
	0x706fdd6b849f0661: []BookEntry{
		{Move: Move(0xf74), Weight: 3},
		{Move: Move(0xeeb), Weight: 1},
	},
	0x8d283b31f9fd483b: []BookEntry{
		{Move: Move(0xd5), Weight: 2},
	},
	0xb241549328aa1560: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xbcb86377f0077e8: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0xc6816a1ad5cdd1c6: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xec335f803cd05d47: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xa1ff9057ff6fe082: []BookEntry{
		{Move: Move(0x94), Weight: 1},
	},
	0xcd0eb4ba78ca707b: []BookEntry{
		{Move: Move(0xe6a), Weight: 4},
	},
	0xd1d7cc69d02ac116: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0xf8a38ac0b44f296e: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x70aed077ad2f9392: []BookEntry{
		{Move: Move(0xd8), Weight: 1},
	},
	0x974806a002c2909f: []BookEntry{
		{Move: Move(0x89b), Weight: 10},
	},
	0xa83dbf4d01ec7aef: []BookEntry{
		{Move: Move(0xf62), Weight: 9},
	},
	0xdfba6450cb268853: []BookEntry{
		{Move: Move(0x9ad), Weight: 65520},
		{Move: Move(0x99f), Weight: 43680},
	},
	0x36a21bc5e1af7e33: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
		{Move: Move(0x29a), Weight: 1},
	},
	0xa0d939de7fe86acd: []BookEntry{
		{Move: Move(0xd8), Weight: 1},
	},
	0xaa8d97bf911d9a77: []BookEntry{
		{Move: Move(0x4db), Weight: 2},
	},
	0x174c1637a3537029: []BookEntry{
		{Move: Move(0x566), Weight: 4},
	},
	0x37c3829b9ae5f8bb: []BookEntry{
		{Move: Move(0xe9e), Weight: 1},
	},
	0x3a7a1765357cbebb: []BookEntry{
		{Move: Move(0xf38), Weight: 1},
	},
	0x3ed0c1c28f65350a: []BookEntry{
		{Move: Move(0x49b), Weight: 14},
	},
	0x8da0e8ed4d0b1b29: []BookEntry{
		{Move: Move(0xceb), Weight: 2},
	},
	0xb1edfb6798e73ee4: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0xb7e12be75cb8f077: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
		{Move: Move(0x6e2), Weight: 1},
	},
	0xc38a7b8ecc1465fb: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0xa3c359584521a: []BookEntry{
		{Move: Move(0xb25), Weight: 1},
	},
	0x359a73dc4e004f4b: []BookEntry{
		{Move: Move(0xa6), Weight: 65520},
		{Move: Move(0x9d), Weight: 65520},
		{Move: Move(0x314), Weight: 65520},
	},
	0x5f09e211257e88a0: []BookEntry{
		{Move: Move(0xf3f), Weight: 65520},
		{Move: Move(0xcaa), Weight: 16380},
	},
	0x87d02d1dfd824315: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xa11f4cadad02beaa: []BookEntry{
		{Move: Move(0x210), Weight: 2},
	},
	0xf3f3c0192e5643ae: []BookEntry{
		{Move: Move(0xb5c), Weight: 1},
	},
	0x104676bf8ad1922d: []BookEntry{
		{Move: Move(0xce3), Weight: 65520},
		{Move: Move(0xca2), Weight: 1},
		{Move: Move(0xe6a), Weight: 7280},
		{Move: Move(0xfad), Weight: 1},
	},
	0x5ac1bcc138709bfc: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x632029cca8bd4536: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xc8a4326ca5b8f21d: []BookEntry{
		{Move: Move(0x52), Weight: 1},
	},
	0x2e67c1a08d983d0b: []BookEntry{
		{Move: Move(0x2db), Weight: 3},
	},
	0x43f4441aa16b9ffe: []BookEntry{
		{Move: Move(0xa9b), Weight: 3},
		{Move: Move(0xf74), Weight: 1},
	},
	0x9a41ae165a86452e: []BookEntry{
		{Move: Move(0xf3f), Weight: 6},
	},
	0xb522655e6175c793: []BookEntry{
		{Move: Move(0x85a), Weight: 1},
	},
	0xd0396140a1cfa361: []BookEntry{
		{Move: Move(0xee9), Weight: 1},
	},
	0xe46556b89784dbda: []BookEntry{
		{Move: Move(0x314), Weight: 4},
		{Move: Move(0x6a3), Weight: 2},
		{Move: Move(0xd1), Weight: 1},
	},
	0xa19f2aa20254ec6: []BookEntry{
		{Move: Move(0xd1), Weight: 1},
	},
	0x45e6efa252b7f95d: []BookEntry{
		{Move: Move(0x292), Weight: 1},
	},
	0x820df88ac269bc92: []BookEntry{
		{Move: Move(0x292), Weight: 4},
	},
	0xa00dade256266c2c: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0xb960d466a822d9fe: []BookEntry{
		{Move: Move(0x2db), Weight: 2},
		{Move: Move(0x251), Weight: 1},
		{Move: Move(0x2d3), Weight: 1},
	},
	0xd2612d8b118ca3cd: []BookEntry{
		{Move: Move(0x9d), Weight: 1},
		{Move: Move(0x691), Weight: 65520},
	},
	0xfa0474c2d51f9ad4: []BookEntry{
		{Move: Move(0x688), Weight: 1},
	},
	0x1fc7836547bab49e: []BookEntry{
		{Move: Move(0xe9e), Weight: 3},
		{Move: Move(0xdef), Weight: 2},
	},
	0x6033de01ddc7747d: []BookEntry{
		{Move: Move(0xce3), Weight: 65520},
	},
	0xa6f1592f531871da: []BookEntry{
		{Move: Move(0xc6a), Weight: 1},
	},
	0xb5dfa0ffbc013fef: []BookEntry{
		{Move: Move(0xa6), Weight: 4},
		{Move: Move(0x6e4), Weight: 1},
	},
	0xdcd37724853accae: []BookEntry{
		{Move: Move(0x89a), Weight: 7},
		{Move: Move(0xf74), Weight: 2},
	},
	0x21fae0452dc767d9: []BookEntry{
		{Move: Move(0x218), Weight: 65520},
	},
	0x6713c33eed3660f3: []BookEntry{
		{Move: Move(0x31c), Weight: 2},
	},
	0x84625476ead84dfb: []BookEntry{
		{Move: Move(0x2d3), Weight: 4},
		{Move: Move(0x8106), Weight: 1},
	},
	0xc3f3119056a3b78a: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xee34e406d6aa517a: []BookEntry{
		{Move: Move(0x15a), Weight: 22},
	},
	0xf18d1a8a7575b473: []BookEntry{
		{Move: Move(0xaa0), Weight: 1},
		{Move: Move(0xfad), Weight: 1},
	},
	0xf7b0d689ab79015f: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x305b30d5d82d7a05: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x9e18404fd36ed90f: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x6d7798ed17d2b334: []BookEntry{
		{Move: Move(0xe9e), Weight: 3},
	},
	0x9309c4475002a702: []BookEntry{
		{Move: Move(0xb5e), Weight: 2},
	},
	0xb667b8a422a78b3c: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
		{Move: Move(0x210), Weight: 65520},
	},
	0xf0ae069c165c48e1: []BookEntry{
		{Move: Move(0x8106), Weight: 65520},
		{Move: Move(0x94), Weight: 65520},
	},
	0x1b72625b4fba9729: []BookEntry{
		{Move: Move(0x89b), Weight: 1},
	},
	0x1c4fd4e070b91336: []BookEntry{
		{Move: Move(0x396), Weight: 1},
	},
	0x36e1fee9099b4e7f: []BookEntry{
		{Move: Move(0x396), Weight: 1},
	},
	0x9f13fb5504f37e72: []BookEntry{
		{Move: Move(0x564), Weight: 1},
	},
	0xb6ab446626afd63d: []BookEntry{
		{Move: Move(0x35d), Weight: 1},
	},
	0xe4ba131adc3334e0: []BookEntry{
		{Move: Move(0x3d7), Weight: 65519},
		{Move: Move(0x218), Weight: 65519},
		{Move: Move(0x8106), Weight: 65519},
	},
	0x3449c3626fbc8588: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
		{Move: Move(0x8da), Weight: 1},
	},
	0x572018ddfdb7685f: []BookEntry{
		{Move: Move(0xcaa), Weight: 65519},
		{Move: Move(0xc20), Weight: 23681},
	},
	0x22478b99a53987f6: []BookEntry{
		{Move: Move(0x8106), Weight: 3},
	},
	0xaefa283257a98b70: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
		{Move: Move(0x218), Weight: 1},
	},
	0xfced1919326fe84b: []BookEntry{
		{Move: Move(0xe3a), Weight: 1},
	},
	0x65ac7d3cf17e633: []BookEntry{
		{Move: Move(0x195), Weight: 18},
	},
	0xb3569c7083466a6: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x14b193f6fb541c52: []BookEntry{
		{Move: Move(0xb6e), Weight: 1},
	},
	0x2a9354cb8a332f30: []BookEntry{
		{Move: Move(0xf3f), Weight: 4},
	},
	0x34d817b0cf41ac7d: []BookEntry{
		{Move: Move(0x6d3), Weight: 1},
	},
	0x6d0ef39e7bb04327: []BookEntry{
		{Move: Move(0xae2), Weight: 1},
	},
	0x340e4b154d0f3487: []BookEntry{
		{Move: Move(0xce3), Weight: 5},
	},
	0x5bb74163394fc557: []BookEntry{
		{Move: Move(0xd1), Weight: 2},
	},
	0x7c1da35c5d3cf2be: []BookEntry{
		{Move: Move(0xfad), Weight: 7},
	},
	0xa40e12abe10fce93: []BookEntry{
		{Move: Move(0x18c), Weight: 6},
	},
	0xb6e25b280bc70995: []BookEntry{
		{Move: Move(0x314), Weight: 1},
	},
	0xcb229a49e237f310: []BookEntry{
		{Move: Move(0x2db), Weight: 2},
	},
	0xfab110b103580fe4: []BookEntry{
		{Move: Move(0xe6a), Weight: 2},
	},
	0xff2acfabe43445b2: []BookEntry{
		{Move: Move(0x6d1), Weight: 3},
	},
	0x4f16f938d7a8fce: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x23545b0fc9f3c05d: []BookEntry{
		{Move: Move(0x691), Weight: 26208},
		{Move: Move(0x3d7), Weight: 65520},
		{Move: Move(0x688), Weight: 65520},
		{Move: Move(0x94), Weight: 65520},
	},
	0x7c55b10e19964cae: []BookEntry{
		{Move: Move(0x218), Weight: 1},
	},
	0xae01f60c810cb7e8: []BookEntry{
		{Move: Move(0xdef), Weight: 1},
	},
	0x56d23041fb26590: []BookEntry{
		{Move: Move(0xf6b), Weight: 1},
	},
	0x1b7e13263fc0ab38: []BookEntry{
		{Move: Move(0x15a), Weight: 65520},
		{Move: Move(0x195), Weight: 21840},
	},
	0x95cb9a0112871ea2: []BookEntry{
		{Move: Move(0xf76), Weight: 5},
	},
	0xb53138a2d76aed9e: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0xbe5ffef5873da583: []BookEntry{
		{Move: Move(0x90), Weight: 6},
	},
	0xe2214ef03fd6629e: []BookEntry{
		{Move: Move(0x4a3), Weight: 1},
	},
	0xe5861cc0cfe59059: []BookEntry{
		{Move: Move(0xf3f), Weight: 3},
	},
	0x2c282e121f7f7f57: []BookEntry{
		{Move: Move(0x95e), Weight: 65520},
	},
	0x5b762026f3f92bb6: []BookEntry{
		{Move: Move(0x6e2), Weight: 6},
		{Move: Move(0x153), Weight: 1},
	},
	0x73f7d41106030257: []BookEntry{
		{Move: Move(0x252), Weight: 65520},
	},
	0xa8077f1263864d50: []BookEntry{
		{Move: Move(0x6e5), Weight: 1},
	},
	0xef2b1d14b66d12e5: []BookEntry{
		{Move: Move(0xca2), Weight: 1},
	},
	0x17bf3198d41cb03: []BookEntry{
		{Move: Move(0x144), Weight: 1},
	},
	0xc3705bef422e78d: []BookEntry{
		{Move: Move(0xcaa), Weight: 1},
	},
	0x91f63dc9f864160d: []BookEntry{
		{Move: Move(0xe6a), Weight: 18},
		{Move: Move(0xce3), Weight: 14},
	},
	0xdc0be57475e5defa: []BookEntry{
		{Move: Move(0x99f), Weight: 1},
	},
	0x412ece4665428684: []BookEntry{
		{Move: Move(0x691), Weight: 7280},
		{Move: Move(0x2d3), Weight: 65520},
	},
	0xab778561eb486e97: []BookEntry{
		{Move: Move(0xd2c), Weight: 58240},
		{Move: Move(0xca2), Weight: 65520},
		{Move: Move(0xea5), Weight: 1},
		{Move: Move(0xdae), Weight: 21840},
	},
	0xeff384b0e6e1794a: []BookEntry{
		{Move: Move(0x89), Weight: 2},
	},
	0x237f247225a02652: []BookEntry{
		{Move: Move(0x2db), Weight: 1},
	},
	0x3074c7e3a04ed55d: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0x5d28e68af30eed41: []BookEntry{
		{Move: Move(0x51b), Weight: 3},
	},
	0xbe47ad32e8752f05: []BookEntry{
		{Move: Move(0x6e4), Weight: 1},
	},
	0x721363d36ef43d05: []BookEntry{
		{Move: Move(0x4b), Weight: 1},
	},
	0xc23de3ad2136dce9: []BookEntry{
		{Move: Move(0x715), Weight: 1},
	},
	0xc5a619f7f4bd961f: []BookEntry{
		{Move: Move(0x4b), Weight: 2},
	},
	0xe54b5f507cf5b248: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0xf20a93453141626: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0x4cde026e277b37fe: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0x8a71b6e9735769bb: []BookEntry{
		{Move: Move(0x543), Weight: 1},
	},
	0xdd470fbeba11311b: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0xac828cfe68b2f4d: []BookEntry{
		{Move: Move(0xc28), Weight: 2},
	},
	0x20936171f90e60cf: []BookEntry{
		{Move: Move(0x2d3), Weight: 65520},
	},
	0x5757cc46cd0f37ca: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
		{Move: Move(0xa9b), Weight: 1},
		{Move: Move(0xeeb), Weight: 1},
	},
	0x2dbc8cedfb46a637: []BookEntry{
		{Move: Move(0x8106), Weight: 12},
	},
	0x80da42077a5a9093: []BookEntry{
		{Move: Move(0xfad), Weight: 1},
	},
	0x821cf8151e95a552: []BookEntry{
		{Move: Move(0x7a7), Weight: 1},
	},
	0xb3bc856395b91368: []BookEntry{
		{Move: Move(0xaa3), Weight: 1},
	},
	0x1804f35ebf3428c2: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0x3f0c0809a40666e9: []BookEntry{
		{Move: Move(0x15a), Weight: 1},
	},
	0xaa239ac15a0b51ea: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0xd1b4c2c61ef19fe7: []BookEntry{
		{Move: Move(0xef2), Weight: 1},
		{Move: Move(0xee9), Weight: 1},
	},
	0xe4bac8497ebcaa08: []BookEntry{
		{Move: Move(0x51c), Weight: 1},
	},
	0xfbf0a23e206c1d7e: []BookEntry{
		{Move: Move(0x2d3), Weight: 1},
	},
	0x203bb97c5f6ac6db: []BookEntry{
		{Move: Move(0x8106), Weight: 4},
		{Move: Move(0x292), Weight: 1},
	},
	0xa15f2444aea24e72: []BookEntry{
		{Move: Move(0x195), Weight: 2},
	},
	0xb7491acce599b527: []BookEntry{
		{Move: Move(0xdef), Weight: 2},
		{Move: Move(0xe6a), Weight: 1},
		{Move: Move(0xf3f), Weight: 1},
	},
	0xf3309ddd8e18b61: []BookEntry{
		{Move: Move(0xa6), Weight: 1},
	},
	0x42a4118a1a569bed: []BookEntry{
		{Move: Move(0xce3), Weight: 1},
	},
	0xc6247ae1151155d3: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
		{Move: Move(0x688), Weight: 1},
	},
	0xf1cdae34f7f1107c: []BookEntry{
		{Move: Move(0x195), Weight: 1},
	},
	0xf65c642cab87e077: []BookEntry{
		{Move: Move(0xfad), Weight: 14},
		{Move: Move(0xd24), Weight: 3},
	},
	0x6636a8002a4042ce: []BookEntry{
		{Move: Move(0x652), Weight: 1},
	},
	0xc110c2530d12c550: []BookEntry{
		{Move: Move(0xd2c), Weight: 2},
	},
	0xc9dedbdef3454a99: []BookEntry{
		{Move: Move(0x723), Weight: 1},
	},
	0x57bb6ad299b8f4ce: []BookEntry{
		{Move: Move(0xef2), Weight: 65520},
	},
	0x660b3e269ffc96f0: []BookEntry{
		{Move: Move(0xf74), Weight: 1},
	},
	0x94b57e88b07fb73a: []BookEntry{
		{Move: Move(0x14c), Weight: 43680},
		{Move: Move(0x210), Weight: 65520},
	},
	0x41133489a3f4746f: []BookEntry{
		{Move: Move(0xd8), Weight: 1},
	},
	0xb4518a2f8a881fd7: []BookEntry{
		{Move: Move(0x9d), Weight: 8},
		{Move: Move(0x396), Weight: 6},
	},
	0xdd8392c589701b36: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xe2fdaba149171757: []BookEntry{
		{Move: Move(0xdae), Weight: 1},
	},
	0x2be863329e830176: []BookEntry{
		{Move: Move(0xe6a), Weight: 3},
		{Move: Move(0xf76), Weight: 1},
	},
	0x319ef876a26d5cd3: []BookEntry{
		{Move: Move(0x91c), Weight: 1},
	},
	0x4c121ee63c8af4bd: []BookEntry{
		{Move: Move(0x153), Weight: 1},
	},
	0x6fb0ac5293dbea4b: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
	},
	0xf181c8c1df8647f6: []BookEntry{
		{Move: Move(0x18c), Weight: 1},
	},
	0x289cb3eefa8dc56e: []BookEntry{
		{Move: Move(0xe73), Weight: 5},
	},
	0x2c7ce25cc6b45982: []BookEntry{
		{Move: Move(0xb63), Weight: 1},
	},
	0xedb43da29b14a51: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
		{Move: Move(0x995), Weight: 2},
	},
	0x10dbde55f1ce293c: []BookEntry{
		{Move: Move(0xd2c), Weight: 2},
	},
	0x11adaae9b8bd6f3c: []BookEntry{
		{Move: Move(0x2d2), Weight: 5},
	},
	0x5313a9692bad23ea: []BookEntry{
		{Move: Move(0x853), Weight: 1},
		{Move: Move(0x858), Weight: 1},
	},
	0x6795103701bed4ec: []BookEntry{
		{Move: Move(0x252), Weight: 1},
	},
	0x5031c1baeacb24e: []BookEntry{
		{Move: Move(0xe6a), Weight: 1},
	},
	0x127e9dd963293621: []BookEntry{
		{Move: Move(0x8106), Weight: 1},
	},
	0x9212a4898fa64b88: []BookEntry{
		{Move: Move(0xcaa), Weight: 4},
	},
	0x96b1d3cbc3bca267: []BookEntry{
		{Move: Move(0x8106), Weight: 2},
	},
	0x976bdcd90c4df3e3: []BookEntry{
		{Move: Move(0xd2c), Weight: 6},
	},
	0xc37a9235fad06306: []BookEntry{
		{Move: Move(0x8d9), Weight: 3},
	},
	0xe51eee3563a8ef71: []BookEntry{
		{Move: Move(0x3d7), Weight: 1},
	},
	0xfa172f42962b96eb: []BookEntry{
		{Move: Move(0xf3f), Weight: 5},
	},
	0x12d736ef0c32c0e1: []BookEntry{
		{Move: Move(0x251), Weight: 1},
	},
	0x5f63cc3494b3d6d1: []BookEntry{
		{Move: Move(0xf3f), Weight: 2},
	},
	0xcfeb4cf54b7debb4: []BookEntry{
		{Move: Move(0xe39), Weight: 1},
	},
	0xdb580e1c51ddd8c8: []BookEntry{
		{Move: Move(0xca), Weight: 1},
	},
	0xffe8e8072c3cb07c: []BookEntry{
		{Move: Move(0x50), Weight: 1},
	},
	0x438af5468dc12cad: []BookEntry{
		{Move: Move(0x195), Weight: 8},
	},
	0x4b68ca199ef73d66: []BookEntry{
		{Move: Move(0x31c), Weight: 2},
	},
	0x74cef56100c2b2b2: []BookEntry{
		{Move: Move(0x6a3), Weight: 1},
	},
	0x7bd88439357a4386: []BookEntry{
		{Move: Move(0xd2c), Weight: 1},
	},
	0xbfa799045b27adae: []BookEntry{
		{Move: Move(0x314), Weight: 3},
	},
	0xd5f0a154319b17f3: []BookEntry{
		{Move: Move(0xf3f), Weight: 1},
	},
	0xdc079d717092429f: []BookEntry{
		{Move: Move(0x82a), Weight: 65520},
	},
	0x6b07291c42c9f6c6: []BookEntry{
		{Move: Move(0x14c), Weight: 1},
	},
	0x84f48e3bb957ca83: []BookEntry{
		{Move: Move(0xacf), Weight: 1},
	},
	0x6ef43b63b7cc19: []BookEntry{
		{Move: Move(0x691), Weight: 65520},
	},
	0x63da5409b69c581d: []BookEntry{
		{Move: Move(0xfad), Weight: 7},
	},
	0xc5ec2eb13ae8f88b: []BookEntry{
		{Move: Move(0xca), Weight: 1},
	},
	0x684e883e5eeb66f: []BookEntry{
		{Move: Move(0xeac), Weight: 1},
	},
	0xb86612276ae438b: []BookEntry{
		{Move: Move(0x355), Weight: 1},
	},
	0x115a48d0a373275a: []BookEntry{
		{Move: Move(0xf3f), Weight: 8},
	},
	0x2d3888dac361814a: []BookEntry{
		{Move: Move(0x8106), Weight: 65520},
		{Move: Move(0x2db), Weight: 16380},
	},
}
