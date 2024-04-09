// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1010_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106102925760003560e01c8063935f029e11610160578063d8444715116100d8578063f45e65d81161008c578063f8c68de011610071578063f8c68de01461062e578063fd32aa0f14610636578063ffa1ad741461063e57600080fd5b8063f45e65d814610611578063f68016b71461061a57600080fd5b8063e0e2016d116100bd578063e0e2016d146105ed578063e81b2c6d146105f5578063f2fde38b146105fe57600080fd5b8063d8444715146105dd578063dac6e63a146105e557600080fd5b8063bc49ce5f1161012f578063c71973f611610114578063c71973f614610483578063c9b26f6114610496578063cc731b02146104a957600080fd5b8063bc49ce5f14610473578063c4e8ddfa1461047b57600080fd5b8063935f029e1461043d5780639b7d7f0a14610450578063a711986914610458578063b40a817c1461046057600080fd5b80634add321d1161020e578063550fcdc9116101c257806361d15768116101a757806361d157681461040f578063715018a6146104175780638da5cb5b1461041f57600080fd5b8063550fcdc9146103ff5780635d73369c1461040757600080fd5b80634d9f1559116101f35780634d9f1559146103875780634f16540b1461038f57806354fd4d50146103b657600080fd5b80634add321d146103535780634c1e843d1461037457600080fd5b806318d13918116102655780631fd19ee11161024a5780631fd19ee11461030d5780634397dfef1461031557806348cd4cb11461034b57600080fd5b806318d13918146102f057806319f5cea81461030557600080fd5b806306c9265714610297578063078f29cf146102b25780630a49cb03146102df5780630c18c162146102e7575b600080fd5b61029f610646565b6040519081526020015b60405180910390f35b6102ba610674565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102a9565b6102ba6106ad565b61029f60655481565b6103036102fe3660046120c9565b6106dd565b005b61029f6106f1565b6102ba61071c565b61031d610746565b6040805173ffffffffffffffffffffffffffffffffffffffff909316835260ff9091166020830152016102a9565b61029f61075a565b61035b61078a565b60405167ffffffffffffffff90911681526020016102a9565b61030361038236600461223c565b6107b0565b6102ba610bbb565b61029f7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b6103f26040518060400160405280600681526020017f312e31322e30000000000000000000000000000000000000000000000000000081525081565b6040516102a991906123f1565b6103f2610beb565b61029f610bf5565b61029f610c20565b610303610c4b565b60335473ffffffffffffffffffffffffffffffffffffffff166102ba565b61030361044b366004612404565b610c5f565b6102ba610c75565b6102ba610ca5565b61030361046e366004612426565b610cd5565b61029f610ce6565b6102ba610d11565b610303610491366004612441565b610d41565b6103036104a436600461245d565b610d52565b61056d6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b6040516102a99190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b6103f2610d63565b6102ba610d6d565b61029f610d9d565b61029f60675481565b61030361060c3660046120c9565b610dc8565b61029f60665481565b60685461035b9067ffffffffffffffff1681565b61029f610e7c565b61029f610ea7565b61029f600081565b61067160017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d6124a5565b81565b60006106a86106a460017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad63776124a5565b5490565b905090565b60006106a86106a460017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad6124a5565b6106e5610ff4565b6106ee81611075565b50565b61067160017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a86124a5565b60006106a87f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c085490565b600080610751611132565b90939092509050565b60006106a86106a460017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a06124a5565b6069546000906106a89063ffffffff6a01000000000000000000008204811691166124bc565b600054610100900460ff16158080156107d05750600054600160ff909116105b806107ea5750303b1580156107ea575060005460ff166001145b61087b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108d957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6108e16111af565b6108ea8a610dc8565b6108f38761124e565b6108fd8989611276565b61090686611307565b61092f7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08869055565b61096261095d60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc5986124a5565b849055565b61099661099060017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce95806376124a5565b83519055565b6109cd6109c460017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a86124a5565b60208401519055565b610a046109fb60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad63776124a5565b60408401519055565b610a3b610a3260017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a68718166124a5565b60608401519055565b610a72610a6960017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad6124a5565b60808401519055565b610aa9610aa060017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d6124a5565b60a08401519055565b610ab16113e5565b610abe8260c0015161144d565b610ac784611988565b610acf61078a565b67ffffffffffffffff168667ffffffffffffffff161015610b4c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610872565b8015610baf57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b60006106a86106a460017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a68718166124a5565b60606106a8611dfc565b61067160017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce95806376124a5565b61067160017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a68718166124a5565b610c53610ff4565b610c5d6000611ed6565b565b610c67610ff4565b610c718282611276565b5050565b60006106a86106a460017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d6124a5565b60006106a86106a460017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce95806376124a5565b610cdd610ff4565b6106ee81611307565b61067160017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc5986124a5565b60006106a86106a460017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a86124a5565b610d49610ff4565b6106ee81611988565b610d5a610ff4565b6106ee8161124e565b60606106a8611f4d565b60006106a86106a460017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc5986124a5565b61067160017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a06124a5565b610dd0610ff4565b73ffffffffffffffffffffffffffffffffffffffff8116610e73576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610872565b6106ee81611ed6565b61067160017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad63776124a5565b61067160017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad6124a5565b9055565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b610f58610f2060017f04adb1412b2ddc16fcc0d4538d5c8f07cf9c83abecc6b41f6f69037b708fbcec6124a5565b74ff000000000000000000000000000000000000000060a086901b1673ffffffffffffffffffffffffffffffffffffffff8716179055565b610faf610f8660017f657c3582c29b3176614e3a33ddd1ec48352696a04e92b3c0566d72010fa8863d6124a5565b604051610f979085906020016124e8565b604051602081830303815290604052610ed290612504565b610fee610fdd60017fa48b38a4b44951360fbdcbfaaeae5ed6ae92585412e9841b70ec72ed8cd057646124a5565b604051610f979084906020016124e8565b50505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610c5d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610872565b61109e7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08829055565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161112691906123f1565b60405180910390a35050565b600080806111646106a460017f04adb1412b2ddc16fcc0d4538d5c8f07cf9c83abecc6b41f6f69037b708fbcec6124a5565b73ffffffffffffffffffffffffffffffffffffffff811693509050826111a3575073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee92601292509050565b60a081901c9150509091565b600054610100900460ff16611246576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610872565b610c5d612000565b60678190556040805160208082018490528251808303909101815290820190915260006110f5565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516112fa91906123f1565b60405180910390a3505050565b61130f61078a565b67ffffffffffffffff168167ffffffffffffffff16101561138c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610872565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff831690811790915560408051602080820193909352815180820390930183528101905260026110f5565b6114136106a460017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a06124a5565b600003610c5d57610c5d61144860017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a06124a5565b439055565b73ffffffffffffffffffffffffffffffffffffffff811661147f575073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5b604080518082018252600581527f45746865720000000000000000000000000000000000000000000000000000006020808301919091528251808401909352600383527f45544800000000000000000000000000000000000000000000000000000000009083015260129173ffffffffffffffffffffffffffffffffffffffff841673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee146118cd578373ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611566573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061158a9190612549565b92508260ff166012146115f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f53797374656d436f6e6669673a2062616420646563696d616c730000000000006044820152606401610872565b8373ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b8152600401600060405180830381865afa158015611644573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261168a9190810190612566565b915060208260405160200161169f91906124e8565b604051602081830303815290604052511115611763576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f53797374656d436f6e6669673a206e616d65206f662067617320706179696e6760448201527f20746f6b656e2063616e6e6f742062652067726561746572207468616e20333260648201527f2062797465730000000000000000000000000000000000000000000000000000608482015260a401610872565b8373ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b8152600401600060405180830381865afa1580156117ae573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526117f49190810190612566565b905060208160405160200161180991906124e8565b6040516020818303038152906040525111156118cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f53797374656d436f6e6669673a2073796d626f6c206f6620676173207061796960448201527f6e6720746f6b656e2063616e6e6f742062652067726561746572207468616e2060648201527f3332206279746573000000000000000000000000000000000000000000000000608482015260a401610872565b73ffffffffffffffffffffffffffffffffffffffff841673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14610fee5761190a84848484610ef2565b6119126106ad565b73ffffffffffffffffffffffffffffffffffffffff1663ef3e242d858585856040518563ffffffff1660e01b81526004016119509493929190612631565b600060405180830381600087803b15801561196a57600080fd5b505af115801561197e573d6000803e3d6000fd5b5050505050505050565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff161115611a38576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d6178206261736500000000000000000000006064820152608401610872565b6001816040015160ff1611611acf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e203100000000000000000000000000000000006064820152608401610872565b6068546080820151825167ffffffffffffffff90921691611af0919061267b565b63ffffffff161115611b5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610872565b6000816020015160ff1611611bf5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f74206265203000000000000000000000000000000000006064820152608401610872565b8051602082015163ffffffff82169160ff90911690611c1590829061269a565b611c1f91906126e4565b63ffffffff1614611cb2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d69740000000000000000006064820152608401610872565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b60606000611e08611132565b5090507fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff821601611e8157505060408051808201909152600381527f4554480000000000000000000000000000000000000000000000000000000000602082015290565b611eaf6106a460017fa48b38a4b44951360fbdcbfaaeae5ed6ae92585412e9841b70ec72ed8cd057646124a5565b604051602001611ec191815260200190565b60405160208183030381529060405291505090565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000611f59611132565b5090507fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff821601611fd257505060408051808201909152600581527f4574686572000000000000000000000000000000000000000000000000000000602082015290565b611eaf6106a460017f657c3582c29b3176614e3a33ddd1ec48352696a04e92b3c0566d72010fa8863d6124a5565b600054610100900460ff16612097576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610872565b610c5d33611ed6565b803573ffffffffffffffffffffffffffffffffffffffff811681146120c457600080fd5b919050565b6000602082840312156120db57600080fd5b6120e4826120a0565b9392505050565b803567ffffffffffffffff811681146120c457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561215557612155612103565b60405290565b803563ffffffff811681146120c457600080fd5b60ff811681146106ee57600080fd5b600060c0828403121561219057600080fd5b60405160c0810181811067ffffffffffffffff821117156121b3576121b3612103565b6040529050806121c28361215b565b815260208301356121d28161216f565b602082015260408301356121e58161216f565b60408201526121f66060840161215b565b60608201526122076080840161215b565b608082015260a08301356fffffffffffffffffffffffffffffffff8116811461222f57600080fd5b60a0919091015292915050565b6000806000806000806000806000898b0361028081121561225c57600080fd5b6122658b6120a0565b995060208b0135985060408b0135975060608b0135965061228860808c016120eb565b955061229660a08c016120a0565b94506122a58c60c08d0161217e565b93506122b46101808c016120a0565b925060e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60820112156122e657600080fd5b506122ef612132565b6122fc6101a08c016120a0565b815261230b6101c08c016120a0565b602082015261231d6101e08c016120a0565b604082015261232f6102008c016120a0565b60608201526123416102208c016120a0565b60808201526123536102408c016120a0565b60a08201526123656102608c016120a0565b60c0820152809150509295985092959850929598565b60005b8381101561239657818101518382015260200161237e565b83811115610fee5750506000910152565b600081518084526123bf81602086016020860161237b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006120e460208301846123a7565b6000806040838503121561241757600080fd5b50508035926020909101359150565b60006020828403121561243857600080fd5b6120e4826120eb565b600060c0828403121561245357600080fd5b6120e4838361217e565b60006020828403121561246f57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156124b7576124b7612476565b500390565b600067ffffffffffffffff8083168185168083038211156124df576124df612476565b01949350505050565b600082516124fa81846020870161237b565b9190910192915050565b80516020808301519190811015612543577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8160200360031b1b821691505b50919050565b60006020828403121561255b57600080fd5b81516120e48161216f565b60006020828403121561257857600080fd5b815167ffffffffffffffff8082111561259057600080fd5b818401915084601f8301126125a457600080fd5b8151818111156125b6576125b6612103565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156125fc576125fc612103565b8160405282815287602084870101111561261557600080fd5b61262683602083016020880161237b565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516815260ff8416602082015260806040820152600061266960808301856123a7565b828103606084015261262681856123a7565b600063ffffffff8083168185168083038211156124df576124df612476565b600063ffffffff808416806126d8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600063ffffffff8083168185168183048111821515161561270757612707612476565b0294935050505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
	immutableReferences["SystemConfig"] = false
}
