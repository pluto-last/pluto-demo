package main

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileHeader struct {
	Magic   uint32
	Version uint8
	Flags   uint8
	Size    uint32
}

func tryMultipleDecompress(data []byte) ([]byte, error) {
	// 尝试多种组合解压方式
	decompressors := []struct {
		name string
		fn   func([]byte) ([]byte, error)
	}{
		{"zlib", tryZlib},
		{"gzip", tryGzip},
		{"raw_deflate", tryDeflate},
		{"custom", tryCustom},
	}

	for _, d := range decompressors {
		if out, err := d.fn(data); err == nil {
			fmt.Printf("成功使用 %s 解压\n", d.name)
			return out, nil
		}
	}
	return nil, fmt.Errorf("无法解压")
}

func tryZlib(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

func tryGzip(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

func tryDeflate(data []byte) ([]byte, error) {
	r := flate.NewReader(bytes.NewReader(data))
	if r == nil {
		return nil, fmt.Errorf("创建deflate reader失败")
	}
	defer r.Close()
	return io.ReadAll(r)
}

func tryCustom(data []byte) ([]byte, error) {
	// 检测自定义压缩格式
	if len(data) < 8 {
		return nil, fmt.Errorf("数据太短")
	}

	magic := binary.LittleEndian.Uint32(data[0:4])
	if magic == 0x000e8651 {
		// 跳过头部
		dataStart := 0x400
		if len(data) <= dataStart {
			return nil, fmt.Errorf("数据部分太短")
		}

		compressedData := data[dataStart:]
		// 尝试不同的解压方式
		if out, err := tryZlib(compressedData); err == nil {
			return out, nil
		}
		if out, err := tryGzip(compressedData); err == nil {
			return out, nil
		}
		if out, err := tryDeflate(compressedData); err == nil {
			return out, nil
		}
	}
	return nil, fmt.Errorf("未知格式")
}

func saveOutput(data []byte, name string) error {
	outPath := filepath.Join(".", name)
	return os.WriteFile(outPath, data, 0644)
}

func tryDecompressSegment(data []byte, segmentSize int) ([]byte, error) {
	if len(data) < segmentSize {
		return nil, fmt.Errorf("数据不足")
	}

	segment := data[:segmentSize]
	// 检查数据段是否全为0
	isZero := true
	for _, b := range segment {
		if b != 0 {
			isZero = false
			break
		}
	}
	if isZero {
		return nil, fmt.Errorf("空段")
	}

	// 尝试各种解压方式
	decompressors := []struct {
		name string
		fn   func([]byte) ([]byte, error)
	}{
		{"zlib", tryZlib},
		{"gzip", tryGzip},
		{"raw_deflate", tryDeflate},
	}

	for _, d := range decompressors {
		if out, err := d.fn(segment); err == nil {
			fmt.Printf("使用 %s 成功解压段数据\n", d.name)
			return out, nil
		}
	}

	return nil, fmt.Errorf("无法解压段数据")
}

func main() {
	inPath := "/home/pluto/go/src/elise/cmd/script/demo2/1.data"
	data, err := os.ReadFile(inPath)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}

	fmt.Printf("文件大小: %d (0x%x) 字节\n", len(data), len(data))
	fmt.Printf("文件头: %02x %02x %02x %02x\n", data[0], data[1], data[2], data[3])

	// 1. 首先提取并保存文件头信息
	headerSize := 0x400
	if len(data) > headerSize {
		header := data[:headerSize]
		saveOutput(header, "1.data.header")
		fmt.Printf("\n头部信息(前32字节):\n")
		for i := 0; i < 32; i++ {
			fmt.Printf("%02x ", header[i])
			if (i+1)%16 == 0 {
				fmt.Println()
			}
		}
	}

	// 2. 尝试以不同大小分段处理数据部分
	if len(data) > headerSize {
		dataSection := data[headerSize:]
		segmentSizes := []int{1024, 2048, 4096, 8192}

		fmt.Println("\n尝试分段解压...")
		for offset := 0; offset < len(dataSection); offset += 1024 {
			for _, size := range segmentSizes {
				if offset+size > len(dataSection) {
					continue
				}
				segment := dataSection[offset : offset+size]
				if out, err := tryDecompressSegment(segment, size); err == nil {
					outName := fmt.Sprintf("1.data.segment_%x_%x", offset, size)
					saveOutput(out, outName)
					fmt.Printf("成功解压段: offset=0x%x, size=0x%x\n", offset, size)
				}
			}
		}
	}

	// 3. 尝试反转数据
	fmt.Println("\n尝试反转数据...")
	if len(data) > headerSize {
		reversed := make([]byte, len(data)-headerSize)
		for i := 0; i < len(reversed); i++ {
			reversed[i] = data[len(data)-1-i]
		}
		if out, err := tryMultipleDecompress(reversed); err == nil {
			saveOutput(out, "1.data.reversed")
			fmt.Println("成功解压反转数据")
		}
	}

	// 4. 尝试XOR组合
	fmt.Println("\n尝试XOR组合...")
	xorKeys := []byte{0x51, 0x86, 0x0e}
	if len(data) > headerSize {
		dataSection := data[headerSize:]
		for i, key := range xorKeys {
			for j, key2 := range xorKeys {
				if i != j {
					decoded := make([]byte, len(dataSection))
					for k := range dataSection {
						if k%2 == 0 {
							decoded[k] = dataSection[k] ^ key
						} else {
							decoded[k] = dataSection[k] ^ key2
						}
					}
					if out, err := tryMultipleDecompress(decoded); err == nil {
						outName := fmt.Sprintf("1.data.xor_%02x_%02x", key, key2)
						saveOutput(out, outName)
						fmt.Printf("成功使用XOR组合(0x%02x,0x%02x)解压\n", key, key2)
					}
				}
			}
		}
	}
}
