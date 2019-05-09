package evm

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/Arachnid/evmdis"
	"github.com/eager7/one_chain/common/utils"
	"strings"
)

const swarmHashLength = 43

var swarmHashProgramTrailer = [...]byte{0x00, 0x29}
var swarmHashHeader = [...]byte{0xa1, 0x65}

func CheckERC20InterfaceByDisassemble(codeHex string) bool {
	if str, err := ParseCode([]byte(utils.HexFormat(codeHex))); err != nil {
		fmt.Println("parse code error:", err)
		return false
	} else {
		if strings.Contains(str, KnowMethods["transfer(address,uint256)"]) {
			return true
		}
	}
	return false
}

func ParseCode(data []byte) (string, error) {
	data = bytes.TrimSpace(data)
	byteCode := make([]byte, hex.DecodedLen(len(data)))
	if _, err := hex.Decode(byteCode, data); err != nil {
		return "", errors.New("Could not decode hex string:" + err.Error())
	}
	disassembly, err := Disassemble(byteCode, true, true)
	if err != nil {
		return "", errors.New("Unable to disassemble:" + err.Error())
	}
	return disassembly, nil
}

func Disassemble(byteCode []byte, withSwarmHash bool, ctorMode bool) (disassembly string, err error) {
	// detect swarm hash and remove it from byteCode, see http://solidity.readthedocs.io/en/latest/miscellaneous.html?highlight=swarm#encoding-of-the-metadata-hash-in-the-bytecode
	byteCodeLength := uint64(len(byteCode))
	if byteCode[byteCodeLength-1] == swarmHashProgramTrailer[1] &&
		byteCode[byteCodeLength-2] == swarmHashProgramTrailer[0] &&
		byteCode[byteCodeLength-43] == swarmHashHeader[0] &&
		byteCode[byteCodeLength-42] == swarmHashHeader[1] && withSwarmHash {
		byteCodeLength -= swarmHashLength // remove swarm part
	}

	program := evmdis.NewProgram(byteCode[:byteCodeLength])
	if err := AnalyzeProgram(program); err != nil {
		return "", err
	}

	if ctorMode {
		var codeEntryPoint = FindNextCodeEntryPoint(program)
		if codeEntryPoint == 0 {
			return disassembly, fmt.Errorf("no code entrypoint found in ctor")
		} else if codeEntryPoint >= byteCodeLength {
			return disassembly, fmt.Errorf("code entrypoint outside of currently available code")
		}

		ctor := evmdis.NewProgram(byteCode[:codeEntryPoint])
		code := evmdis.NewProgram(byteCode[codeEntryPoint:byteCodeLength])

		if err := AnalyzeProgram(ctor); err != nil {
			return "", err
		}
		disassembly += fmt.Sprintln("# Constructor part -------------------------")
		disassembly += PrintAnalysisResult(ctor)

		if err := AnalyzeProgram(code); err != nil {
			return "", err
		}
		disassembly += fmt.Sprintln("# Code part -------------------------")
		disassembly += PrintAnalysisResult(code)

	} else {
		disassembly += PrintAnalysisResult(program)
	}

	return disassembly, nil
}

func FindNextCodeEntryPoint(program *evmdis.Program) uint64 {
	var lastPos uint64 = 0
	for _, block := range program.Blocks {
		for _, instruction := range block.Instructions {
			if instruction.Op == evmdis.CODECOPY {
				var expression evmdis.Expression

				instruction.Annotations.Get(&expression)

				arg := expression.(*evmdis.InstructionExpression).Arguments[1].Eval()

				if arg != nil {
					lastPos = arg.Uint64()
				}
			}
		}
	}
	return lastPos
}

func PrintAnalysisResult(program *evmdis.Program) (disassembly string) {
	for _, block := range program.Blocks {
		offset := block.Offset

		// Print out the jump label for the block, if there is one
		var label *evmdis.JumpLabel
		block.Annotations.Get(&label)
		if label != nil {
			disassembly += fmt.Sprintf("%v\n", label)
		}

		// Print out the stack prestate for this block
		var reaching evmdis.ReachingDefinition
		block.Annotations.Get(&reaching)

		blockDisassembly := fmt.Sprintf("# Stack: %v\n", reaching)
		blockRealInstructions := 0

		for _, instruction := range block.Instructions {
			var expression evmdis.Expression
			instruction.Annotations.Get(&expression)

			if expression != nil {
				if instruction.Op.StackWrites() == 1 && !instruction.Op.IsDup() {
					blockDisassembly += fmt.Sprintf("0x%X\tPUSH(%v)\n", offset, expression)
				} else {
					blockDisassembly += fmt.Sprintf("0x%X\t%v\n", offset, expression)
				}

				blockRealInstructions++
			}
			offset += instruction.Op.OperandSize() + 1
		}

		blockDisassembly += fmt.Sprintf("\n")

		// avoid printing empty stack frames with no instructions in the block
		if len(reaching) > 0 || blockRealInstructions > 0 {
			disassembly += blockDisassembly
		}
	}

	return disassembly
}

func AnalyzeProgram(program *evmdis.Program) (err error) {
	if err := evmdis.PerformReachingAnalysis(program); err != nil {
		return errors.New("Error performing reaching analysis:" + err.Error())
	}
	evmdis.PerformReachesAnalysis(program)
	evmdis.CreateLabels(program)
	if err := evmdis.BuildExpressions(program); err != nil {
		return errors.New("Error building expressions:" + err.Error())
	}

	return nil
}
