import pandas as pd

def big_countries(world: pd.DataFrame) -> pd.DataFrame:
    df = world[(world['population'] >= 25000000) | (world['area'] >= 3000000)]
    return df[['name', 'population', 'area']]
