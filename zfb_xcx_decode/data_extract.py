#!/usr/bin/env python3
import struct
import os
import sys
import glob
import tarfile
from io import BytesIO

def extract_custom_archive(filename, output_dir=None):
    """解压自定义格式的压缩文件"""
    if output_dir is None:
        output_dir = os.path.dirname(filename)

    try:
        with open(filename, 'rb') as f:
            data = f.read()

        print(f"开始解压文件: {filename}")
        print(f"文件总大小: {len(data)} 字节")

        # 查找所有ustar标识符
        ustar_positions = []
        pos = 0
        while True:
            pos = data.find(b'ustar', pos)
            if pos == -1:
                break
            ustar_positions.append(pos)
            pos += 1

        print(f'找到 {len(ustar_positions)} 个ustar标识符')

        file_count = 0

        for i, ustar_pos in enumerate(ustar_positions):
            # 从ustar位置向前查找512字节边界
            tar_start = ((ustar_pos - 257) // 512) * 512
            if tar_start < 0:
                continue

            # 解析tar头部
            header = data[tar_start:tar_start+512]
            if len(header) != 512:
                continue

            # 提取文件名
            filename_bytes = header[0:100]
            filename_str = filename_bytes.decode('ascii', errors='ignore').rstrip('\x00')
            if not filename_str:
                continue

            # 提取文件大小
            size_str = header[124:136].decode('ascii', errors='ignore').rstrip('\x00').strip()
            if not size_str.isdigit():
                continue

            file_size = int(size_str, 8)
            if file_size <= 0:
                continue

            # 提取文件数据
            data_start = tar_start + 512
            if data_start + file_size > len(data):
                continue

            file_data = data[data_start:data_start+file_size]
            if len(file_data) != file_size:
                continue

            print(f'  发现文件: {filename_str} ({file_size} 字节)')

            # 创建完整的输出路径
            output_path = os.path.join(output_dir, filename_str)
            output_dir_path = os.path.dirname(output_path)

            # 创建目录结构
            os.makedirs(output_dir_path, exist_ok=True)

            # 保存文件
            try:
                with open(output_path, 'wb') as out:
                    out.write(file_data)
                file_count += 1
                print(f'  ✓ 已保存到: {output_path}')
            except Exception as e:
                print(f'  ✗ 保存文件失败 {output_path}: {e}')
                continue

            # 如果是tar文件，尝试进一步解压
            if filename_str.endswith('.tar'):
                try:
                    tar_dir = os.path.dirname(output_path)
                    with tarfile.open(output_path, 'r') as tar:
                        tar.extractall(path=tar_dir)
                    print(f'  ✓ 成功解压tar文件到: {tar_dir}')
                except Exception as e:
                    print(f'  ✗ 解压tar文件失败: {e}')

        print(f'成功提取 {file_count} 个文件')
        return file_count

    except Exception as e:
        print(f'处理文件时出错: {e}')
        return 0

def find_data_files(start_dir):
    """递归查找所有的0.data文件"""
    data_files = []
    for root, dirs, files in os.walk(start_dir):
        if '0.data' in files:
            data_files.append(os.path.join(root, '0.data'))
    return data_files

def process_directory(data_file):
    """处理单个0.data文件"""
    print(f"\n=== 处理文件: {data_file} ===")
    try:
        dir_path = os.path.dirname(data_file)
        file_count = extract_custom_archive(data_file, dir_path)
        return file_count
    except Exception as e:
        print(f"处理 {data_file} 时出错: {e}")
        return 0

def main():
    """主函数"""
    # 默认处理 data 目录
    start_dir = './data'
    if len(sys.argv) > 1:
        start_dir = sys.argv[1]

    start_dir = os.path.abspath(start_dir)
    if not os.path.exists(start_dir):
        print(f"目录不存在: {start_dir}")
        return

    print(f"开始搜索目录: {start_dir}")
    data_files = find_data_files(start_dir)

    if not data_files:
        print("未找到任何0.data文件")
        return

    print(f"找到 {len(data_files)} 个0.data文件")

    total_files = 0
    for data_file in sorted(data_files):
        file_count = process_directory(data_file)
        total_files += file_count

    print(f"\n=== 总结 ===")
    print(f"总共处理了 {len(data_files)} 个0.data文件")
    print(f"总共提取了 {total_files} 个文件")

if __name__ == '__main__':
    main()
