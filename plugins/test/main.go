/**
  @author: decision
  @date: 2024/6/27
  @note:
**/

package main

var Instance = TestChain{
	identity: "/storage/ipfs/QmXViwQ1frFwabQHtmpt18SUPhnpcRzhWayt9rnTJ8GTay",
}

type TestChain struct {
	identity string
}

func (tc *TestChain) Identity() (string, error) {
	return tc.identity, nil
}

func (tc *TestChain) Bootstrap() (string, error) {
	return "", nil
}

func (tc *TestChain) Initial(ident string, url string) error {
	return nil
}

func (tc *TestChain) SetIdentity(ident string) error {
	tc.identity = ident
	return nil
}

func (tc *TestChain) Join(url string) error {
	return nil
}

func (tc *TestChain) Setup(address string) error {
	return nil
}
