package test_contracts

var IdentityTestAbi = []byte(`{
  "version": "eosio::abi/1.0",
  "types": [{
      "new_type_name": "account_name",
      "type": "name"
    },{
      "new_type_name": "identity_name",
      "type": "name"
    }
  ],
  "structs": [{
      "name": "getowner",
      "base": "",
      "fields": [
        {"name":"identity", "type":"uint64"}
      ]
    },{
      "name": "getidentity",
      "base": "",
      "fields": [
        {"name":"account", "type":"account_name"}
      ]
    }
  ],
  "actions": [{
      "name": "getowner",
      "type": "getowner"
    },{
      "name": "getidentity",
      "type": "getidentity"
    }
  ],
  "tables": [
  ]
}`)
