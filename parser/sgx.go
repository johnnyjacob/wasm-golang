package main

import (
	"errors"
	"fmt"

	"github.com/saferwall/elf"
)

type SGXEnclaveInfo struct {
	MREnclave []uint8
	MRSigner  []uint8
}

const SGX_MAGIC_NUM uint64 = 0x86A80294635D0E4C

// Convenience function to create elf parser from file or from a buffer.
func loadBuffer(data []byte) (*elf.Parser, error) {
	return elf.NewBytes(data)
}

func loadFile(path string) (*elf.Parser, error) {
	return elf.New(path)
}

func getSGXMetadata(buffer []byte, offset uint64) (SGXEnclaveInfo, error) {
	fmt.Println()
	// TODO : Read the magic number . Check on byte ordering.
	for i := 0; i < 8; i++ {
		cOff := offset + uint64(i)
		fmt.Printf(" %02X", buffer[cOff])
	}
	fmt.Println()

	return SGXEnclaveInfo{}, errors.New("Unable to parse SGX Metadata")
}

func getSGXMetaOffset(p *elf.Parser) (uint64, error) {
	err := p.ParseIdent()
	if err != nil {
		panic(err)
	}

	elfClass := p.F.Ident.Class

	err = p.ParseELFHeader(elfClass)
	if err != nil {
		panic(err)
	}
	err = p.ParseELFSectionHeaders(elfClass)
	if err != nil {
		panic(err)
	}
	err = p.ParseELFSections(elfClass)
	if err != nil {
		panic(err)
	}

	for _, section := range p.F.ELFBin64.Sections64 {
		if section.SectionName == ".note.sgxmeta" {
			fmt.Println(section)
			// SGX META OFFSET
			fmt.Println(section.ELF64SectionHeader.Off)
			fmt.Println(section.Size)
			fmt.Println(section.ELF64SectionHeader.Size)
			// FIXME : Dynamically calculate '25' . This is size of 'note' & 'name'
			return section.ELF64SectionHeader.Off + 25, nil
		}
	}
	return 0, errors.New("SGX meta section not found")
}

func GetSGXMetadata(buffer []byte) (SGXEnclaveInfo, error) {
	p, err := loadBuffer(buffer)
	if err != nil {
		return SGXEnclaveInfo{}, errors.New("Unable to initialize ELF parser")
	}

	offset, err := getSGXMetaOffset(p)
	if err != nil {
		return SGXEnclaveInfo{}, errors.New("Unable to get SGX metadata offset")
	}

	return getSGXMetadata(buffer, offset)

}
