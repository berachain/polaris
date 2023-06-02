#!/bin/bash
# this script generates the sim_tests variable containing
# the slice of testSpec for the given source

source_file="$1"
dest="$2"
sim_name="$3"

# Check if the Go file argument is provided
if [ -z "$source_file" ] ||  [ -z "$sim_name" ] ||  [ -z "$dest" ]; then
  echo "Usage: ./generate_tests.sh <source> <dest> <sim_name> "
  exit 1
fi

# Extract function names using grep and assign them to a bash array
function_names=($(grep -E "^func" "$source_file" | cut -d '(' -f 1 | cut -d ' ' -f 2))

# Dump function names
echo "function names:"
for name in "${function_names[@]}"; do
  echo "$name"
done

# Generate the Go code to append the function names to the existing variable
go_code="[]testSpec{"
for name in "${function_names[@]}"; do
  capitalized_name=$(awk 'BEGIN {print toupper(substr("'"$name"'", 1, 1)) substr("'"$name"'", 2)}')
  go_code+="{Name: \"http/$capitalized_name\", Run: $name},"
done

for name in "${function_names[@]}"; do
  capitalized_name=$(awk 'BEGIN {print toupper(substr("'"$name"'", 1, 1)) substr("'"$name"'", 2)}')
  go_code+="{Name: \"ws/$capitalized_name\", Run: $name},"
done
go_code+="} //nolint: lll // auto-generated"

# Set the variable in Go
sim_var=$sim_name"Tests"

# Modify the dest file
awk -v name="$sim_var" -v code="$go_code" '$0 ~ name " =" {$0 = "\t" name " = " code} 1' "$dest" > temp.txt
mv temp.txt "$dest"

echo "Go file modified"
