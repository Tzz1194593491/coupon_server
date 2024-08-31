package utils

import "github.com/yitter/idgenerator-go/idgen"

func Init() {
	idGen()
}

func idGen() {
	var options = idgen.NewIdGeneratorOptions(1)
	options.SeqBitLength = 10
	idgen.SetIdGenerator(options)
}
