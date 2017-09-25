import pandas as pd

# Define column names.
cols = [
        'integercolumn',
        'stringcolumn'
        ]

# Read in the CSV with pandas.
data = pd.read_csv('myfile.csv', names=cols)

# Print out the maximum value in the integer column.
print(data['integercolumn'].max())
