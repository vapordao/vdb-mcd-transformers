package test_data

const (
	KovanCatABI        = `[{"constant":true,"inputs":[],"name":"vat","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x36569e77"},{"constant":true,"inputs":[],"name":"vow","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x626cb3c5"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"flips","outputs":[{"name":"ilk","type":"bytes32"},{"name":"urn","type":"bytes32"},{"name":"ink","type":"uint256"},{"name":"tab","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x70d9235a"},{"constant":true,"inputs":[],"name":"nflip","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x76181a51"},{"constant":true,"inputs":[],"name":"live","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x957aa58c"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"wards","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xbf353dbb"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"ilks","outputs":[{"name":"flip","type":"address"},{"name":"chop","type":"uint256"},{"name":"lump","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xd9638d36"},{"constant":true,"inputs":[],"name":"pit","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xf03c7c6e"},{"inputs":[{"name":"vat_","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor","signature":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"ilk","type":"bytes32"},{"indexed":true,"name":"urn","type":"bytes32"},{"indexed":false,"name":"ink","type":"uint256"},{"indexed":false,"name":"art","type":"uint256"},{"indexed":false,"name":"tab","type":"uint256"},{"indexed":false,"name":"flip","type":"uint256"},{"indexed":false,"name":"iArt","type":"uint256"}],"name":"Bite","type":"event","signature":"0x99b5620489b6ef926d4518936cfec15d305452712b88bd59da2d9c10fb0953e8"},{"anonymous":true,"inputs":[{"indexed":true,"name":"sig","type":"bytes4"},{"indexed":true,"name":"guy","type":"address"},{"indexed":true,"name":"foo","type":"bytes32"},{"indexed":true,"name":"bar","type":"bytes32"},{"indexed":false,"name":"wad","type":"uint256"},{"indexed":false,"name":"fax","type":"bytes"}],"name":"LogNote","type":"event","signature":"0x644843f351d3fba4abcd60109eaff9f54bac8fb8ccf0bab941009c21df21cf31"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"rely","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x65fae35e"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"deny","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x9c52a7f1"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"what","type":"bytes32"},{"name":"data","type":"uint256"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x1a0b287e"},{"constant":false,"inputs":[{"name":"what","type":"bytes32"},{"name":"data","type":"address"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0xd4e8be83"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"what","type":"bytes32"},{"name":"flip","type":"address"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0xebecb39d"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"urn","type":"bytes32"}],"name":"bite","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x72f7b593"},{"constant":false,"inputs":[{"name":"n","type":"uint256"},{"name":"wad","type":"uint256"}],"name":"flip","outputs":[{"name":"id","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0xe6f95917"}]`
	KovanFlapperABI    = `[{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"bids","outputs":[{"name":"bid","type":"uint256"},{"name":"lot","type":"uint256"},{"name":"guy","type":"address"},{"name":"tic","type":"uint48"},{"name":"end","type":"uint48"},{"name":"gal","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"ttl","outputs":[{"name":"","type":"uint48"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"gem","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"beg","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"tau","outputs":[{"name":"","type":"uint48"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"kicks","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"dai","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"dai_","type":"address"},{"name":"gem_","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"uint256"},{"indexed":false,"name":"lot","type":"uint256"},{"indexed":false,"name":"bid","type":"uint256"},{"indexed":false,"name":"gal","type":"address"},{"indexed":false,"name":"end","type":"uint48"}],"name":"Kick","type":"event"},{"anonymous":true,"inputs":[{"indexed":true,"name":"sig","type":"bytes4"},{"indexed":true,"name":"guy","type":"address"},{"indexed":true,"name":"foo","type":"bytes32"},{"indexed":true,"name":"bar","type":"bytes32"},{"indexed":false,"name":"wad","type":"uint256"},{"indexed":false,"name":"fax","type":"bytes"}],"name":"LogNote","type":"event"},{"constant":false,"inputs":[{"name":"gal","type":"address"},{"name":"lot","type":"uint256"},{"name":"bid","type":"uint256"}],"name":"kick","outputs":[{"name":"id","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"},{"name":"lot","type":"uint256"},{"name":"bid","type":"uint256"}],"name":"tend","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"deal","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
	KovanFlipperABI    = `[{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"bids","outputs":[{"name":"bid","type":"uint256"},{"name":"lot","type":"uint256"},{"name":"guy","type":"address"},{"name":"tic","type":"uint48"},{"name":"end","type":"uint48"},{"name":"urn","type":"bytes32"},{"name":"gal","type":"address"},{"name":"tab","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x4423c5f1"},{"constant":true,"inputs":[],"name":"ttl","outputs":[{"name":"","type":"uint48"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x4e8b1dd5"},{"constant":true,"inputs":[],"name":"gem","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x7bd2bea7"},{"constant":true,"inputs":[],"name":"beg","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x7d780d82"},{"constant":true,"inputs":[],"name":"tau","outputs":[{"name":"","type":"uint48"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xcfc4af55"},{"constant":true,"inputs":[],"name":"kicks","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xcfdd3302"},{"constant":true,"inputs":[],"name":"dai","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xf4b9fa75"},{"inputs":[{"name":"dai_","type":"address"},{"name":"gem_","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor","signature":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"uint256"},{"indexed":false,"name":"lot","type":"uint256"},{"indexed":false,"name":"bid","type":"uint256"},{"indexed":false,"name":"gal","type":"address"},{"indexed":false,"name":"end","type":"uint48"},{"indexed":true,"name":"urn","type":"bytes32"},{"indexed":false,"name":"tab","type":"uint256"}],"name":"Kick","type":"event","signature":"0xbac86238bdba81d21995024470425ecb370078fa62b7271b90cf28cbd1e3e87e"},{"anonymous":true,"inputs":[{"indexed":true,"name":"sig","type":"bytes4"},{"indexed":true,"name":"guy","type":"address"},{"indexed":true,"name":"foo","type":"bytes32"},{"indexed":true,"name":"bar","type":"bytes32"},{"indexed":false,"name":"wad","type":"uint256"},{"indexed":false,"name":"fax","type":"bytes"}],"name":"LogNote","type":"event","signature":"0x644843f351d3fba4abcd60109eaff9f54bac8fb8ccf0bab941009c21df21cf31"},{"constant":true,"inputs":[],"name":"era","outputs":[{"name":"","type":"uint48"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x143e55e0"},{"constant":false,"inputs":[{"name":"urn","type":"bytes32"},{"name":"gal","type":"address"},{"name":"tab","type":"uint256"},{"name":"lot","type":"uint256"},{"name":"bid","type":"uint256"}],"name":"kick","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0xeae19d9e"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"tick","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0xfc7b6aee"},{"constant":false,"inputs":[{"name":"id","type":"uint256"},{"name":"lot","type":"uint256"},{"name":"bid","type":"uint256"}],"name":"tend","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x4b43ed12"},{"constant":false,"inputs":[{"name":"id","type":"uint256"},{"name":"lot","type":"uint256"},{"name":"bid","type":"uint256"}],"name":"dent","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x5ff3a382"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"deal","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0xc959c42b"}]`
	KovanFlopperABI    = `[{"constant":true,"inputs":[],"name":"era","outputs":[{"name":"","type":"uint48"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"bids","outputs":[{"name":"bid","type":"uint256"},{"name":"lot","type":"uint256"},{"name":"guy","type":"address"},{"name":"tic","type":"uint48"},{"name":"end","type":"uint48"},{"name":"vow","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"ttl","outputs":[{"name":"","type":"uint48"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"},{"name":"lot","type":"uint256"},{"name":"bid","type":"uint256"}],"name":"dent","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"rely","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"gem","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"beg","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"deny","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"gal","type":"address"},{"name":"lot","type":"uint256"},{"name":"bid","type":"uint256"}],"name":"kick","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"wards","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"deal","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"tau","outputs":[{"name":"","type":"uint48"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"kicks","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"dai","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"dai_","type":"address"},{"name":"gem_","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"uint256"},{"indexed":false,"name":"lot","type":"uint256"},{"indexed":false,"name":"bid","type":"uint256"},{"indexed":false,"name":"gal","type":"address"},{"indexed":false,"name":"end","type":"uint48"}],"name":"Kick","type":"event"},{"anonymous":true,"inputs":[{"indexed":true,"name":"sig","type":"bytes4"},{"indexed":true,"name":"guy","type":"address"},{"indexed":true,"name":"foo","type":"bytes32"},{"indexed":true,"name":"bar","type":"bytes32"},{"indexed":false,"name":"wad","type":"uint256"},{"indexed":false,"name":"fax","type":"bytes"}],"name":"LogNote","type":"event"}]`
	KovanJugABI        = `[{"constant":true,"inputs":[],"name":"vat","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"repo","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"vow","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"wards","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"ilks","outputs":[{"name":"tax","type":"uint256"},{"name":"rho","type":"uint48"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"vat_","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":true,"inputs":[{"indexed":true,"name":"sig","type":"bytes4"},{"indexed":true,"name":"guy","type":"address"},{"indexed":true,"name":"foo","type":"bytes32"},{"indexed":true,"name":"bar","type":"bytes32"},{"indexed":false,"name":"wad","type":"uint256"},{"indexed":false,"name":"fax","type":"bytes"}],"name":"LogNote","type":"event"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"rely","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"deny","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"}],"name":"init","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"what","type":"bytes32"},{"name":"data","type":"uint256"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"what","type":"bytes32"},{"name":"data","type":"uint256"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"what","type":"bytes32"},{"name":"data","type":"bytes32"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"}],"name":"drip","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
	KovanMedianizerABI = `[{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"orcl","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"age","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"a","type":"address"}],"name":"lift","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"val","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"min_","type":"uint256"}],"name":"setMin","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"wat","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"read","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"peek","outputs":[{"name":"","type":"bytes32"},{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"val_","type":"uint256[]"},{"name":"age_","type":"uint256[]"},{"name":"v","type":"uint8[]"},{"name":"r","type":"bytes32[]"},{"name":"s","type":"bytes32[]"}],"name":"poke","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"a","type":"address"}],"name":"drop","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"min","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_wat","type":"bytes32"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"who","type":"address"},{"indexed":false,"name":"val","type":"uint256"},{"indexed":false,"name":"age","type":"uint256"}],"name":"LogFeedPrice","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"val","type":"uint256"},{"indexed":false,"name":"age","type":"uint256"}],"name":"LogMedianPrice","type":"event"}]'`
	KovanUpdatedVatABI = `[{"constant":true,"inputs":[],"name":"debt","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"},{"name":"","type":"bytes32"}],"name":"urns","outputs":[{"name":"ink","type":"uint256"},{"name":"art","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"vice","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"can","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"live","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"sin","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"Line","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"wards","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"},{"name":"","type":"bytes32"}],"name":"gem","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"ilks","outputs":[{"name":"Art","type":"uint256"},{"name":"rate","type":"uint256"},{"name":"spot","type":"uint256"},{"name":"line","type":"uint256"},{"name":"dust","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"dai","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":true,"inputs":[{"indexed":true,"name":"hash","type":"bytes4"},{"indexed":true,"name":"arg1","type":"bytes32"},{"indexed":true,"name":"arg2","type":"bytes32"},{"indexed":true,"name":"arg3","type":"bytes32"},{"indexed":false,"name":"data","type":"bytes"}],"name":"Note","type":"event"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"rely","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"deny","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"hope","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"nope","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"}],"name":"init","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"what","type":"bytes32"},{"name":"data","type":"uint256"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"what","type":"bytes32"},{"name":"data","type":"uint256"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"guy","type":"bytes32"},{"name":"rad","type":"int256"}],"name":"slip","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"src","type":"bytes32"},{"name":"dst","type":"bytes32"},{"name":"rad","type":"int256"}],"name":"flux","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"src","type":"bytes32"},{"name":"dst","type":"bytes32"},{"name":"rad","type":"int256"}],"name":"move","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"i","type":"bytes32"},{"name":"u","type":"bytes32"},{"name":"v","type":"bytes32"},{"name":"w","type":"bytes32"},{"name":"dink","type":"int256"},{"name":"dart","type":"int256"}],"name":"frob","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"src","type":"bytes32"},{"name":"dst","type":"bytes32"},{"name":"dink","type":"int256"},{"name":"dart","type":"int256"}],"name":"fork","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"i","type":"bytes32"},{"name":"u","type":"bytes32"},{"name":"v","type":"bytes32"},{"name":"w","type":"bytes32"},{"name":"dink","type":"int256"},{"name":"dart","type":"int256"}],"name":"grab","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"u","type":"bytes32"},{"name":"v","type":"bytes32"},{"name":"rad","type":"int256"}],"name":"heal","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"i","type":"bytes32"},{"name":"u","type":"bytes32"},{"name":"rate","type":"int256"}],"name":"fold","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
	KovanVatABI        = `[{"constant":true,"inputs":[],"name":"debt","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x0dca59c1"},{"constant":true,"inputs":[{"name":"","type":"bytes32"},{"name":"","type":"bytes32"}],"name":"urns","outputs":[{"name":"ink","type":"uint256"},{"name":"art","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x26e27482"},{"constant":true,"inputs":[],"name":"vice","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0x2d61a355"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"sin","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xa60f1d3e"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"wards","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xbf353dbb"},{"constant":true,"inputs":[{"name":"","type":"bytes32"},{"name":"","type":"bytes32"}],"name":"gem","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xc0912683"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"ilks","outputs":[{"name":"take","type":"uint256"},{"name":"rate","type":"uint256"},{"name":"Ink","type":"uint256"},{"name":"Art","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xd9638d36"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"dai","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function","signature":"0xf53e4e69"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor","signature":"constructor"},{"anonymous":true,"inputs":[{"indexed":true,"name":"sig","type":"bytes4"},{"indexed":true,"name":"foo","type":"bytes32"},{"indexed":true,"name":"bar","type":"bytes32"},{"indexed":true,"name":"too","type":"bytes32"},{"indexed":false,"name":"fax","type":"bytes"}],"name":"Note","type":"event","signature":"0x8c2dbbc2b33ffaa77c104b777e574a8a4ff79829dfee8b66f4dc63e3f8067152"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"rely","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x65fae35e"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"deny","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x9c52a7f1"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"}],"name":"init","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x3b663195"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"guy","type":"bytes32"},{"name":"rad","type":"int256"}],"name":"slip","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x42066cbb"},{"constant":false,"inputs":[{"name":"ilk","type":"bytes32"},{"name":"src","type":"bytes32"},{"name":"dst","type":"bytes32"},{"name":"rad","type":"int256"}],"name":"flux","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0xa6e41821"},{"constant":false,"inputs":[{"name":"src","type":"bytes32"},{"name":"dst","type":"bytes32"},{"name":"rad","type":"int256"}],"name":"move","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x78f19470"},{"constant":false,"inputs":[{"name":"i","type":"bytes32"},{"name":"u","type":"bytes32"},{"name":"v","type":"bytes32"},{"name":"w","type":"bytes32"},{"name":"dink","type":"int256"},{"name":"dart","type":"int256"}],"name":"tune","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x5dd6471a"},{"constant":false,"inputs":[{"name":"i","type":"bytes32"},{"name":"u","type":"bytes32"},{"name":"v","type":"bytes32"},{"name":"w","type":"bytes32"},{"name":"dink","type":"int256"},{"name":"dart","type":"int256"}],"name":"grab","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x3690ae4c"},{"constant":false,"inputs":[{"name":"u","type":"bytes32"},{"name":"v","type":"bytes32"},{"name":"rad","type":"int256"}],"name":"heal","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x990a5f63"},{"constant":false,"inputs":[{"name":"i","type":"bytes32"},{"name":"u","type":"bytes32"},{"name":"rate","type":"int256"}],"name":"fold","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0xe6a6a64d"},{"constant":false,"inputs":[{"name":"i","type":"bytes32"},{"name":"u","type":"bytes32"},{"name":"take","type":"int256"}],"name":"toll","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function","signature":"0x09b7a0b5"}]`
	KovanVowABI        = `[{"constant":true,"inputs":[],"name":"Awe","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"Joy","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"flap","outputs":[{"name":"id","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"hump","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"wad","type":"uint256"}],"name":"kiss","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"what","type":"bytes32"},{"name":"data","type":"uint256"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"Ash","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"era","type":"uint48"}],"name":"flog","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"vat","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"Woe","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"wait","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"rely","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"bump","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"tab","type":"uint256"}],"name":"fess","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"row","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint48"}],"name":"sin","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"}],"name":"deny","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"flop","outputs":[{"name":"id","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"wards","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"sump","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"Sin","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"what","type":"bytes32"},{"name":"addr","type":"address"}],"name":"file","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"cow","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"wad","type":"uint256"}],"name":"heal","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":true,"inputs":[{"indexed":true,"name":"sig","type":"bytes4"},{"indexed":true,"name":"guy","type":"address"},{"indexed":true,"name":"foo","type":"bytes32"},{"indexed":true,"name":"bar","type":"bytes32"},{"indexed":false,"name":"wad","type":"uint256"},{"indexed":false,"name":"fax","type":"bytes"}],"name":"LogNote","type":"event"}]`
)
