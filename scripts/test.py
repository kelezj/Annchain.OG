from requests import Session

d = {
    "nonce": "1",
    "from": "0x835f6c9617c39c7473db1956124e2bd142b9ea59",
    "value": "0",
    "pubkey": "04a15c49424acef8fb488047ee93faad517b3f9480a24e9610caa58687ffd0c89d2e4c278eb81f33ce166dd598d42ccead02888f17c298f9a2cd6e6a24cd3574c0",
    "signature": "51ec13fd60579f8a463865bc5da8279400fa41130b32c0d4f609afa08eb8598c74582b5b18759346371b3156078fe3b33fb7c0a5c33dfab826e64b1182407e3501",
    # "data": "0x608060405234801561001057600080fd5b506040805190810160405280600e81526020017f68656c6c6f206861636b65723a200000000000000000000000000000000000008152506000908051906020019061005c929190610062565b50610107565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a357805160ff19168380011785556100d1565b828001600101855582156100d1579182015b828111156100d05782518255916020019190600101906100b5565b5b5090506100de91906100e2565b5090565b61010491905b808211156101005760008160009055506001016100e8565b5090565b90565b61066a806101166000396000f3fe60806040526004361061003b576000357c010000000000000000000000000000000000000000000000000000000090048063e978c36f14610040575b600080fd5b34801561004c57600080fd5b506101066004803603602081101561006357600080fd5b810190808035906020019064010000000081111561008057600080fd5b82018360208201111561009257600080fd5b803590602001918460018302840111640100000000831117156100b457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610108565b005b60606101e160008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101a35780601f10610178576101008083540402835291602001916101a3565b820191906000526020600020905b81548152906001019060200180831161018657829003601f168201915b50505050508360206040519081016040528060008152506020604051908101604052806000815250602060405190810160405280600081525061021f565b9050600081805190602001209050600180600083815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b6060808690506060869050606086905060608690506060869050606081518351855187518951010101016040519080825280601f01601f19166020018201604052801561027b5781602001600182028038833980820191505090505b5090506060819050600080905060008090505b88518110156103415788818151811015156102a557fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561030457fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808060010191505061028e565b5060008090505b87518110156103fb57878181518110151561035f57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f01000000000000000000000000000000000000000000000000000000000000000283838060010194508151811015156103be57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610348565b5060008090505b86518110156104b557868181518110151561041957fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561047857fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610402565b5060008090505b855181101561056f5785818151811015156104d357fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561053257fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506104bc565b5060008090505b845181101561062957848181518110151561058d57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f01000000000000000000000000000000000000000000000000000000000000000283838060010194508151811015156105ec57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610576565b5081985050505050505050509594505050505056fea165627a7a72305820202d57d8f1975669a47c7c73b4605fedd13d897c23e6f49d9bb58c22e2fb83ff0029"
    'data': '608060405234801561001057600080fd5b5060408051808201909152600e8082527f68656c6c6f206861636b65723a2000000000000000000000000000000000000060209092019182526100559160009161005b565b506100f6565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009c57805160ff19168380011785556100c9565b828001600101855582156100c9579182015b828111156100c95782518255916020019190600101906100ae565b506100d59291506100d9565b5090565b6100f391905b808211156100d557600081556001016100df565b90565b6104bc806101056000396000f3fe608060405260043610610045577c010000000000000000000000000000000000000000000000000000000060003504638e7d4b1d811461004a578063e978c36f14610088575b600080fd5b34801561005657600080fd5b506100746004803603602081101561006d57600080fd5b503561013d565b604080519115158252519081900360200190f35b34801561009457600080fd5b5061013b600480360360208110156100ab57600080fd5b8101906020810181356401000000008111156100c657600080fd5b8201836020820111156100d857600080fd5b803590602001918460018302840111640100000000831117156100fa57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610152945050505050565b005b60016020526000908152604090205460ff1681565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815260609361021593919290918301828280156101e15780601f106101b6576101008083540402835291602001916101e1565b820191906000526020600020905b8154815290600101906020018083116101c457829003601f168201915b505060408051602081810183526000808352835180830185528181528451928301909452815289955090935090915061023b565b8051602091820120600090815260019182905260409020805460ff191690911790555050565b6060808690506060869050606086905060608690506060869050606081518351855187518951010101016040519080825280601f01601f19166020018201604052801561028f576020820181803883390190505b509050806000805b88518110156102f55788818151811015156102ae57fe5b90602001015160f860020a900460f860020a0283838060010194508151811015156102d557fe5b906020010190600160f860020a031916908160001a905350600101610297565b5060005b875181101561035757878181518110151561031057fe5b90602001015160f860020a900460f860020a02838380600101945081518110151561033757fe5b906020010190600160f860020a031916908160001a9053506001016102f9565b5060005b86518110156103b957868181518110151561037257fe5b90602001015160f860020a900460f860020a02838380600101945081518110151561039957fe5b906020010190600160f860020a031916908160001a90535060010161035b565b5060005b855181101561041b5785818151811015156103d457fe5b90602001015160f860020a900460f860020a0283838060010194508151811015156103fb57fe5b906020010190600160f860020a031916908160001a9053506001016103bd565b5060005b845181101561047d57848181518110151561043657fe5b90602001015160f860020a900460f860020a02838380600101945081518110151561045d57fe5b906020010190600160f860020a031916908160001a90535060010161041f565b50909d9c5050505050505050505050505056fea165627a7a723058200fabcda4da8c9a52ab07862e04db9410eaca1a6caa179bb54d33f75408ce84e30029'
}

# private: ac74d6ee082f9b5bbbde23d31e4dd2cb94a5f3d5210c755501ff2744992d5490
# public: 04a15c49424acef8fb488047ee93faad517b3f9480a24e9610caa58687ffd0c89d2e4c278eb81f33ce166dd598d42ccead02888f17c298f9a2cd6e6a24cd3574c0
# address: 835f6c9617c39c7473db1956124e2bd142b9ea59


d2 = {
    "nonce": "1",
    "from": "835f6c9617c39c7473db1956124e2bd142b9ea59",
    "to": "0xa3fa6b326c91bfa1addd6fa9fd27f266cce43e0d",
    "value": "0",
    "data": "0xe978c36f0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000b3133303030303030303030000000000000000000000000000000000000000000",
    "signature": "f8c494c951506b6fb7f6ea2fdd7c3ea708196343337a8cdca4d87e985210cec750ffea50168c1593ab474c52cbe928ce950a02cd02ede5a72c26ebfaabd83fa400",
    "pubkey": "04a15c49424acef8fb488047ee93faad517b3f9480a24e9610caa58687ffd0c89d2e4c278eb81f33ce166dd598d42ccead02888f17c298f9a2cd6e6a24cd3574c0"
}
d3 = {
    'address':'0xa3fa6b326c91bfa1addd6fa9fd27f266cce43e0d',
    'data': '8e7d4b1d8f0e3ba4f86410dfeea816dfd55b788552cdb7a1a1fdcd746a4a9e38a2d49e4b'
}

if __name__ == '__main__':
    s = Session()
    s.trust_env = False
    resp = s.post('http://127.0.0.1:8000/new_transaction', json=d)
    print(resp.text)

if __name__ == '__main__1':
    s = Session()
    s.trust_env = False
    resp = s.post('http://127.0.0.1:8000/query_contract', json=d3)
    print(resp.text)