import struct
from thrift.protocol import TCompactProtocol
from thrift.transport import TTransport
from genpy.parquet.ttypes import SchemaElement, ColumnMetaData, FileMetaData, ColumnChunk, PageHeader, DataPageHeader

colnum_metadata_bytes="265a1c150419250006191806636f6c6e756d15001604165216522608491c1500150015020000001500152415242c15041500150615061c000000"
colnum_metadata_bytes2="265a1c150419250006191806636f6c6e756d15001604165216522608491c1500150015020000001500"
col_chunk_1="1500152c152c2c15041500150615061c0000000200000004012a000000000000004200000000000000"

def try_decode_bytes_as_column_metadata(bytes) -> None:
    transport = TTransport.TMemoryBuffer(bytes)
    protocol = TCompactProtocol.TCompactProtocol(transport)
    colnum_metadata = SchemaElement()
    colnum_metadata.read(protocol)
    print(colnum_metadata)

def try_decode_bytes_as_page_header(bytes) -> None:
    transport = TTransport.TMemoryBuffer(bytes)
    protocol = TCompactProtocol.TCompactProtocol(transport)
    ph = ColumnChunk()
    ph.read(protocol)
    print(ph)

def read_parquet_footer(file):
    file.seek(-8, 2)  # Seek to 8 bytes from the end of the file
    footer_length = struct.unpack("<i", file.read(4))[0]  # Read the footer length
    file.seek(-(8 + footer_length), 2)  # Seek to the start of the footer
    footer_bytes = file.read(footer_length)  # Read the footer bytes
    return footer_bytes

def decode_parquet_metadata(file_path):
    with open(file_path, "rb") as file:
        footer_bytes = read_parquet_footer(file)
        print(f"Footer bytes: \n{footer_bytes}")
        print(f"Footer bytes len = {len(footer_bytes)}")
        transport = TTransport.TMemoryBuffer(footer_bytes)
        protocol = TCompactProtocol.TCompactProtocol(transport)
        metadata = FileMetaData()
        metadata.read(protocol)
        return metadata

def main() -> None:
    metadata = decode_parquet_metadata("two_rows2.parquet")
    print(metadata)

    print("SchemaElement")
    try_decode_bytes_as_column_metadata(bytes.fromhex(colnum_metadata_bytes))
    #try_decode_bytes_as_page_header(bytes.fromhex(col_chunk_1))


if __name__ == '__main__':
    main()
