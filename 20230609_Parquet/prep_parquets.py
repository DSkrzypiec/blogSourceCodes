import pandas as pd
import pyarrow as pa
import pyarrow.parquet as pq


def df_to_parquet(df: pd.DataFrame, file_name: str) -> None:
    table = pa.Table.from_pandas(df)
    pq.write_table(
        table,
        file_name,
        row_group_size=1000,
        use_dictionary=False,
        compression="NONE",
        write_statistics=False,
        column_encoding="PLAIN",
        store_schema=False
    )

def prep_two_row_df() -> pd.DataFrame:
    return pd.DataFrame({
        "colnum": [42, 66],
        "colstr": ["ds", "sd"]
    })

def prep_and_save_two_row_df() -> None:
    file_name = "two_rows2.parquet"
    df_to_parquet(prep_two_row_df(), file_name)
    print(f"Two rows data frame saved to {file_name}")

def main() -> None:
    prep_and_save_two_row_df()

if __name__ == '__main__':
    main()
